package playwright

import (
	"fmt"
	"log"
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

type callback struct {
	Data  interface{}
	Error error
}

type connection struct {
	apiZone      sync.Map
	objects      map[string]*channelOwner
	lastID       int
	lastIDLock   sync.Mutex
	rootObject   *rootChannelOwner
	callbacks    sync.Map
	afterClose   func()
	onClose      func() error
	onmessage    func(map[string]interface{}) error
	isRemote     bool
	localUtils   *localUtilsImpl
	tracingCount atomic.Int32
}

func (c *connection) Start() *Playwright {
	playwright := make(chan *Playwright, 1)
	go func() {
		pw, err := c.rootObject.initialize()
		if err != nil {
			log.Fatal(err)
			return
		}
		playwright <- pw
	}()
	return <-playwright
}

func (c *connection) Stop() error {
	err := c.onClose()
	if err != nil {
		return err
	}
	c.cleanup()
	return nil
}

func (c *connection) cleanup() {
	if c.afterClose != nil {
		c.afterClose()
	}
	c.callbacks.Range(func(key, value any) bool {
		select {
		case value.(chan callback) <- callback{
			Error: &Error{
				Name:    "Error",
				Message: "Connection closed",
			}}:
		default:
		}
		return true
	})
}

func (c *connection) Dispatch(msg *message) {
	method := msg.Method
	if msg.ID != 0 {
		cb, _ := c.callbacks.Load(msg.ID)
		if msg.Error != nil {
			cb.(chan callback) <- callback{
				Error: parseError(msg.Error.Error),
			}
		} else {
			cb.(chan callback) <- callback{
				Data: c.replaceGuidsWithChannels(msg.Result),
			}
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
		object.dispose()
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
	call := func() (interface{}, error) {
		result := reflect.ValueOf(cb).Call(nil)
		// accept 0 to 2 return values
		switch len(result) {
		case 2:
			val := result[0].Interface()
			err, ok := result[1].Interface().(error)
			if ok && err != nil {
				return val, err
			}
			return val, nil
		case 1:
			return result[0].Interface(), nil
		default:
			return nil, nil
		}
	}
	if _, ok := c.apiZone.Load("apiZone"); !ok {
		return call()
	}
	defer func() {
		c.apiZone.Delete("apiZone")
	}()
	c.apiZone.Store("apiZone", serializeCallStack(isInternal))
	return call()
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

func (c *connection) sendMessageToServer(guid string, method string, params interface{}) (interface{}, error) {
	c.lastIDLock.Lock()
	c.lastID++
	id := c.lastID
	c.lastIDLock.Unlock()
	var (
		metadata = make(map[string]interface{}, 0)
		stack    = make([]map[string]interface{}, 0)
	)
	apiZone, ok := c.apiZone.Load("apiZone")
	if ok {
		metadata = apiZone.(*parsedStackTrace).metadata
		stack = apiZone.(*parsedStackTrace).frames
	}
	metadata["wallTime"] = time.Now().Nanosecond()
	message := map[string]interface{}{
		"id":       id,
		"guid":     guid,
		"method":   method,
		"params":   c.replaceChannelsWithGuids(params),
		"metadata": metadata,
	}
	cb, _ := c.callbacks.LoadOrStore(id, make(chan callback))
	if err := c.onmessage(message); err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}

	if c.tracingCount.Load() > 0 && len(stack) > 0 && guid != "localUtils" {
		c.LocalUtils().AddStackToTracingNoReply(id, stack)
	}
	result := <-cb.(chan callback)
	c.callbacks.Delete(id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.Data, nil
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

func serializeCallStack(isInternal bool) *parsedStackTrace {
	st := stack.Trace().TrimRuntime()

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
	return &parsedStackTrace{
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

func newConnection(onClose func() error, localUtils ...*localUtilsImpl) *connection {
	connection := &connection{
		objects:  make(map[string]*channelOwner),
		onClose:  onClose,
		isRemote: false,
	}
	if len(localUtils) > 0 {
		connection.localUtils = localUtils[0]
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
