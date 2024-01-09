package playwright

import (
	"log"
	"reflect"
)

type channel struct {
	eventEmitter
	guid       string
	connection *connection
	owner      *channelOwner // to avoid type conversion
	object     interface{}   // retain type info (for fromChannel needed)
}

func (c *channel) Send(method string, options ...interface{}) (interface{}, error) {
	return c.connection.WrapAPICall(func() (interface{}, error) {
		return c.innerSend(method, false, options...)
	}, false)
}

func (c *channel) SendReturnAsDict(method string, options ...interface{}) (interface{}, error) {
	return c.connection.WrapAPICall(func() (interface{}, error) {
		return c.innerSend(method, true, options...)
	}, true)
}

func (c *channel) innerSend(method string, returnAsDict bool, options ...interface{}) (interface{}, error) {
	params := transformOptions(options...)
	callback, err := c.connection.sendMessageToServer(c.owner, method, params, false)
	if err != nil {
		return nil, err
	}
	result, err := callback.GetResult()
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	if returnAsDict {
		return result, nil
	}
	if reflect.TypeOf(result).Kind() == reflect.Map {
		mapV := result.(map[string]interface{})
		if len(mapV) == 0 {
			return nil, nil
		}
		for key := range mapV {
			return mapV[key], nil
		}
	}
	return result, nil
}

func (c *channel) SendNoReply(method string, options ...interface{}) {
	params := transformOptions(options...)
	_, err := c.connection.WrapAPICall(func() (interface{}, error) {
		return c.connection.sendMessageToServer(c.owner, method, params, true)
	}, false)
	if err != nil {
		log.Printf("SendNoReply failed: %v", err)
	}
}

func newChannel(owner *channelOwner, object interface{}) *channel {
	channel := &channel{
		connection: owner.connection,
		guid:       owner.guid,
		owner:      owner,
		object:     object,
	}
	return channel
}
