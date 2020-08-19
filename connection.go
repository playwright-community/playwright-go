package playwright

import (
	"fmt"
	"io"
	"reflect"
	"sync"
)

type callback struct {
	Data  interface{}
	Error error
}

type Connection struct {
	transport                   *Transport
	waitingForRemoteObjectsLock sync.Mutex
	waitingForRemoteObjects     map[string]chan interface{}
	objects                     map[string]*ChannelOwner
	lastID                      int
	rootObject                  *ChannelOwner
	callbacks                   map[int]chan callback
	stopDriver                  func() error
}

func (c *Connection) Start() error {
	c.transport.SetDispatch(c.Dispatch)
	return c.transport.Start()
}

func (c *Connection) Stop() error {
	if err := c.transport.Stop(); err != nil {
		return fmt.Errorf("could not stop transport: %v", err)
	}
	return c.stopDriver()
}

func (c *Connection) CallOnObjectWithKnownName(name string) (interface{}, error) {
	if _, ok := c.waitingForRemoteObjects[name]; !ok {
		c.waitingForRemoteObjectsLock.Lock()
		c.waitingForRemoteObjects[name] = make(chan interface{})
		c.waitingForRemoteObjectsLock.Unlock()
	}
	return <-c.waitingForRemoteObjects[name], nil
}

func (c *Connection) Dispatch(msg *Message) {
	method := msg.Method
	if msg.ID != 0 {
		if msg.Error != nil {
			c.callbacks[msg.ID] <- callback{
				Error: parseError(msg.Error.Error),
			}
		} else {
			c.callbacks[msg.ID] <- callback{
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
	if method == "__dispose__" {
		object.Dispose()
		return
	}
	object.channel.Emit(method, c.replaceGuidsWithChannels(msg.Params))
}

func (c *Connection) createRemoteObject(parent *ChannelOwner, objectType string, guid string, initializer interface{}) interface{} {
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

func (c *Connection) replaceChannelsWithGuids(payload interface{}) interface{} {
	if payload == nil {
		return nil
	}
	if valA, isChannel := payload.(Channel); isChannel {
		return map[string]string{
			"guid": valA.guid,
		}
	}
	v := reflect.ValueOf(payload)
	if v.Kind() == reflect.Slice {
		listV := payload.([]interface{})
		for i := 0; i < len(listV); i++ {
			listV[i] = c.replaceChannelsWithGuids(listV[i])
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
			mapV[key] = c.replaceChannelsWithGuids(mapV[key])
		}
		return mapV
	}
	return payload
}

func (c *Connection) replaceGuidsWithChannels(payload interface{}) interface{} {
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

func (c *Connection) SendMessageToServer(guid string, method string, params interface{}) (interface{}, error) {
	c.lastID++
	id := c.lastID
	message := map[string]interface{}{
		"id":     id,
		"guid":   guid,
		"method": method,
		"params": c.replaceChannelsWithGuids(params),
	}
	if _, ok := c.callbacks[id]; !ok {
		c.callbacks[id] = make(chan callback)
	}
	if err := c.transport.Send(message); err != nil {
		return nil, fmt.Errorf("could not send message: %v", err)
	}
	result := <-c.callbacks[id]
	delete(c.callbacks, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.Data, nil
}

func newConnection(stdin io.WriteCloser, stdout io.ReadCloser, stopDriver func() error) *Connection {
	connection := &Connection{
		waitingForRemoteObjects: make(map[string]chan interface{}),
		transport:               newTransport(stdin, stdout),
		objects:                 make(map[string]*ChannelOwner),
		callbacks:               make(map[int]chan callback),
		stopDriver:              stopDriver,
	}
	connection.rootObject = newRootChannelOwner(connection)
	return connection
}
