package playwright

import (
	"fmt"
	"io"
	"reflect"
)

type Connection struct {
	transport               *Transport
	waitingForRemoteObjects map[string]chan interface{}
	objects                 map[string]*ChannelOwner
	lastID                  int
	rootObject              *ChannelOwner
	callbacks               map[int]chan interface{}
}

func (c *Connection) Start() error {
	c.transport.SetDispatch(c.Dispatch)
	return c.transport.Start()
}

func (c *Connection) Stop() error {
	return c.transport.Stop()
}

func (c *Connection) CallOnObjectWithKnownName(name string) (interface{}, error) {
	if _, ok := c.waitingForRemoteObjects[name]; !ok {
		c.waitingForRemoteObjects[name] = make(chan interface{})
	}
	return <-c.waitingForRemoteObjects[name], nil
}

func (c *Connection) Dispatch(msg *Message) error {
	method := msg.Method
	if msg.ID != 0 {
		c.callbacks[msg.ID] <- c.replaceGuidsWithChannels(msg.Result)
		// TODO: Error handling
		return nil
	}
	object := c.objects[msg.GUID]
	if method == "__create__" {
		c.createRemoteObject(
			object, msg.Params.Type, msg.Params.GUID, msg.Params.Initializer,
		)
		return nil
	}
	if method == "__dispose__" {
		object.Dispose()
		return nil
	}

	return nil
}

func (c *Connection) createRemoteObject(parent *ChannelOwner, objectType string, guid string, initializer interface{}) interface{} {
	initializer = c.replaceGuidsWithChannels(initializer)
	result := createObjectFactory(parent, objectType, guid, initializer)
	if _, ok := c.waitingForRemoteObjects[guid]; ok {
		c.waitingForRemoteObjects[guid] <- result
	}
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
		"params": c.replaceGuidsWithChannels(params),
	}
	if _, ok := c.callbacks[id]; !ok {
		c.callbacks[id] = make(chan interface{})
	}
	if err := c.transport.Send(message); err != nil {
		return nil, fmt.Errorf("could not send message: %v", err)
	}
	return <-c.callbacks[id], nil
}

func newConnection(stdin io.WriteCloser, stdout io.ReadCloser) *Connection {
	connection := &Connection{
		waitingForRemoteObjects: make(map[string]chan interface{}),
		transport:               newTransport(stdin, stdout),
		objects:                 make(map[string]*ChannelOwner),
		callbacks:               make(map[int]chan interface{}),
	}
	connection.rootObject = newRootChannelOwner(connection)
	return connection
}
