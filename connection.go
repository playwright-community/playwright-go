package playwright

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-stack/stack"
)

var (
	pkgSourcePathPattern = regexp.MustCompile(`.+[\\/]playwright-go[\\/][^\\/]+\.go`)
	apiNameTransform     = regexp.MustCompile(`(?U)\(\*(.+)(Impl)?\)`)
)

type result struct {
	Data  interface{}
	Error error
}

type connection struct {
	transport    transport
	apiZone      sync.Map
	objects      map[string]*channelOwner
	lastID       atomic.Uint32
	rootObject   *rootChannelOwner
	callbacks    sync.Map
	afterClose   func()
	onClose      func() error
	isRemote     bool
	localUtils   *localUtilsImpl
	tracingCount atomic.Int32
	abort        chan struct{}
	abortOnce    sync.Once
	closedError  *safeValue[error]
}

func (c *connection) Start() (*Playwright, error) {
	go func() {
		for {
			msg, err := c.transport.Poll()
			if err != nil {
				_ = c.transport.Close()
				c.cleanup(err)
				return
			}
			c.Dispatch(msg)
		}
	}()

	c.onClose = func() error {
		if err := c.transport.Close(); err != nil {
			return err
		}
		return nil
	}

	return c.rootObject.initialize()
}

func (c *connection) Stop() error {
	if err := c.onClose(); err != nil {
		return err
	}
	c.cleanup()
	return nil
}

func (c *connection) cleanup(cause ...error) {
	if len(cause) > 0 {
		c.closedError.Set(fmt.Errorf("%w: %w", ErrTargetClosed, cause[0]))
	} else {
		c.closedError.Set(ErrTargetClosed)
	}
	if c.afterClose != nil {
		c.afterClose()
	}
	c.abortOnce.Do(func() {
		select {
		case <-c.abort:
		default:
			close(c.abort)
		}
	})
}

func (c *connection) Dispatch(msg *message) {
	if c.closedError.Get() != nil {
		return
	}
	method := msg.Method
	if msg.ID != 0 {
		cb, _ := c.callbacks.LoadAndDelete(uint32(msg.ID))
		if cb.(*protocolCallback).noReply {
			return
		}
		if msg.Error != nil {
			cb.(*protocolCallback).SetResult(result{
				Error: parseError(msg.Error.Error),
			})
		} else {
			cb.(*protocolCallback).SetResult(result{
				Data: c.replaceGuidsWithChannels(msg.Result),
			})
		}
		return
	}
	object := c.objects[msg.GUID]
	if method == "__create__" {
		c.createRemoteObject(
			object, msg.Params["type"].(string), msg.Params["guid"].(string), msg.Params["initializer"],
		)
		return
	}
	if object == nil {
		return
	}
	if method == "__adopt__" {
		child, ok := c.objects[msg.Params["guid"].(string)]
		if !ok {
			return
		}
		object.adopt(child)
		return
	}
	if method == "__dispose__" {
		reason, ok := msg.Params["reason"]
		if ok {
			object.dispose(reason.(string))
		} else {
			object.dispose()
		}
		return
	}
	if object.objectType == "JsonPipe" {
		object.channel.Emit(method, msg.Params)
	} else {
		object.channel.Emit(method, c.replaceGuidsWithChannels(msg.Params))
	}
}

func (c *connection) LocalUtils() *localUtilsImpl {
	return c.localUtils
}

func (c *connection) createRemoteObject(parent *channelOwner, objectType string, guid string, initializer interface{}) interface{} {
	initializer = c.replaceGuidsWithChannels(initializer)
	result := createObjectFactory(parent, objectType, guid, initializer.(map[string]interface{}))
	return result
}

func (c *connection) WrapAPICall(cb func() (interface{}, error), isInternal bool) (interface{}, error) {
	if _, ok := c.apiZone.Load("apiZone"); ok {
		return cb()
	}
	c.apiZone.Store("apiZone", serializeCallStack(isInternal))
	return cb()
}

func (c *connection) replaceChannelsWithGuids(payload interface{}) interface{} {
	if payload == nil {
		return nil
	}
	if channel, isChannel := payload.(*channel); isChannel {
		return map[string]string{
			"guid": channel.guid,
		}
	}
	v := reflect.ValueOf(payload)
	if v.Kind() == reflect.Slice {
		listV := make([]interface{}, 0)
		for i := 0; i < v.Len(); i++ {
			listV = append(listV, c.replaceChannelsWithGuids(v.Index(i).Interface()))
		}
		return listV
	}
	if v.Kind() == reflect.Map {
		mapV := make(map[string]interface{})
		for _, key := range v.MapKeys() {
			mapV[key.String()] = c.replaceChannelsWithGuids(v.MapIndex(key).Interface())
		}
		return mapV
	}
	return payload
}

func (c *connection) replaceGuidsWithChannels(payload interface{}) interface{} {
	if payload == nil {
		return nil
	}
	v := reflect.ValueOf(payload)
	if v.Kind() == reflect.Slice {
		listV := payload.([]interface{})
		for i := 0; i < len(listV); i++ {
			listV[i] = c.replaceGuidsWithChannels(listV[i])
		}
		return listV
	}
	if v.Kind() == reflect.Map {
		mapV := payload.(map[string]interface{})
		if guid, hasGUID := mapV["guid"]; hasGUID {
			if channelOwner, ok := c.objects[guid.(string)]; ok {
				return channelOwner.channel
			}
		}
		for key := range mapV {
			mapV[key] = c.replaceGuidsWithChannels(mapV[key])
		}
		return mapV
	}
	return payload
}

func (c *connection) sendMessageToServer(object *channelOwner, method string, params interface{}, noReply bool) (*protocolCallback, error) {
	if err := c.closedError.Get(); err != nil {
		return nil, err
	}
	if object.wasCollected {
		return nil, errors.New("The object has been collected to prevent unbounded heap growth.")
	}

	id := c.lastID.Add(1)
	cb, _ := c.callbacks.LoadOrStore(id, newProtocolCallback(noReply, c.abort))
	var (
		metadata = make(map[string]interface{}, 0)
		stack    = make([]map[string]interface{}, 0)
	)
	apiZone, ok := c.apiZone.LoadAndDelete("apiZone")
	if ok {
		for k, v := range apiZone.(parsedStackTrace).metadata {
			metadata[k] = v
		}
		stack = append(stack, apiZone.(parsedStackTrace).frames...)
	}
	metadata["wallTime"] = time.Now().Nanosecond()
	message := map[string]interface{}{
		"id":       id,
		"guid":     object.guid,
		"method":   method,
		"params":   c.replaceChannelsWithGuids(params),
		"metadata": metadata,
	}
	if c.tracingCount.Load() > 0 && len(stack) > 0 && object.guid != "localUtils" {
		c.LocalUtils().AddStackToTracingNoReply(id, stack)
	}

	if err := c.transport.Send(message); err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}

	return cb.(*protocolCallback), nil
}

func (c *connection) setInTracing(isTracing bool) {
	if isTracing {
		c.tracingCount.Add(1)
	} else {
		c.tracingCount.Add(-1)
	}
}

type parsedStackTrace struct {
	frames   []map[string]interface{}
	metadata map[string]interface{}
}

func serializeCallStack(isInternal bool) parsedStackTrace {
	st := stack.Trace().TrimRuntime()
	if len(st) == 0 { // https://github.com/go-stack/stack/issues/27
		st = stack.Trace()
	}

	lastInternalIndex := 0
	for i, s := range st {
		if pkgSourcePathPattern.MatchString(s.Frame().File) {
			lastInternalIndex = i
		}
	}
	apiName := ""
	if !isInternal {
		apiName = fmt.Sprintf("%n", st[lastInternalIndex])
	}
	st = st.TrimBelow(st[lastInternalIndex])

	callStack := make([]map[string]interface{}, 0)
	for i, s := range st {
		if i == 0 {
			continue
		}
		callStack = append(callStack, map[string]interface{}{
			"file":     s.Frame().File,
			"line":     s.Frame().Line,
			"column":   0,
			"function": s.Frame().Function,
		})
	}
	metadata := make(map[string]interface{})
	if len(st) > 1 {
		metadata["location"] = serializeCallLocation(st[1])
	}
	apiName = apiNameTransform.ReplaceAllString(apiName, "$1")
	if len(apiName) > 1 {
		apiName = strings.ToUpper(apiName[:1]) + apiName[1:]
	}
	metadata["apiName"] = apiName
	metadata["isInternal"] = isInternal
	return parsedStackTrace{
		metadata: metadata,
		frames:   callStack,
	}
}

func serializeCallLocation(caller stack.Call) map[string]interface{} {
	line, _ := strconv.Atoi(fmt.Sprintf("%d", caller))
	return map[string]interface{}{
		"file": fmt.Sprintf("%s", caller),
		"line": line,
	}
}

func newConnection(transport transport, localUtils ...*localUtilsImpl) *connection {
	connection := &connection{
		abort:       make(chan struct{}, 1),
		objects:     make(map[string]*channelOwner),
		transport:   transport,
		isRemote:    false,
		closedError: &safeValue[error]{},
	}
	if len(localUtils) > 0 {
		connection.localUtils = localUtils[0]
		connection.isRemote = true
	}
	connection.rootObject = newRootChannelOwner(connection)
	return connection
}

func fromChannel(v interface{}) interface{} {
	return v.(*channel).object
}

func fromNullableChannel(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return fromChannel(v)
}

type protocolCallback struct {
	callback chan result
	noReply  bool
	abort    <-chan struct{}
}

func (pc *protocolCallback) SetResult(r result) {
	if pc.noReply {
		return
	}
	select {
	case <-pc.abort:
		select {
		case pc.callback <- r:
		default:
		}
		return
	case pc.callback <- r:
	}
}

func (pc *protocolCallback) GetResult() (interface{}, error) {
	if pc.noReply {
		return nil, nil
	}
	select {
	case result := <-pc.callback:
		return result.Data, result.Error
	case <-pc.abort:
		select {
		case result := <-pc.callback:
			return result.Data, result.Error
		default:
			return nil, errors.New("Connection closed")
		}
	}
}

func newProtocolCallback(noReply bool, abort <-chan struct{}) *protocolCallback {
	if noReply {
		return &protocolCallback{
			noReply: true,
			abort:   abort,
		}
	}
	return &protocolCallback{
		callback: make(chan result, 1),
		abort:    abort,
	}
}
