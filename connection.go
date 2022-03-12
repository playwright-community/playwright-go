package playwright

import (
	"fmt"
	"log"
	"reflect"
	"sync"

	"github.com/go-stack/stack"
)

type callback struct {
	Data  interface{}
	Error error
}

type connection struct {
	waitingForRemoteObjectsLock sync.Mutex
	waitingForRemoteObjects     map[string]chan interface{}
	objects                     map[string]*channelOwner
	lastID                      int
	lastIDLock                  sync.Mutex
	rootObject                  *rootChannelOwner
	callbacks                   sync.Map
	onClose                     func() error
	onmessage                   func(map[string]interface{}) error
	isRemote                    bool
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
	return c.onClose()
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

func (c *connection) createRemoteObject(parent *channelOwner, objectType string, guid string, initializer interface{}) interface{} {
	initializer = c.replaceGuidsWithChannels(initializer)
	result := createObjectFactory(parent, objectType, guid, initializer.(map[string]interface{}))
	c.waitingForRemoteObjectsLock.Lock()
	if _, ok := c.waitingForRemoteObjects[guid]; ok {
		c.waitingForRemoteObjects[guid] <- result
		delete(c.waitingForRemoteObjects, guid)
	}
	c.waitingForRemoteObjectsLock.Unlock()
	return result
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

func (c *connection) SendMessageToServer(guid string, method string, params interface{}) (interface{}, error) {
	c.lastIDLock.Lock()
	c.lastID++
	id := c.lastID
	c.lastIDLock.Unlock()
	stack := serializeCallStack(stack.Trace())
	metadata := make(map[string]interface{})
	metadata["stack"] = stack
	metadata["apiName"] = ""
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
	result := <-cb.(chan callback)
	c.callbacks.Delete(id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.Data, nil
}

func serializeCallStack(stack stack.CallStack) []map[string]interface{} {
	callStack := make([]map[string]interface{}, 0)
	for _, s := range stack {
		callStack = append(callStack, map[string]interface{}{
			"file":     s.Frame().File,
			"line":     s.Frame().Line,
			"function": s.Frame().Function,
		})
	}
	return callStack
}

func newConnection(onClose func() error) *connection {
	connection := &connection{
		waitingForRemoteObjects: make(map[string]chan interface{}),
		objects:                 make(map[string]*channelOwner),
		onClose:                 onClose,
		isRemote:                false,
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
