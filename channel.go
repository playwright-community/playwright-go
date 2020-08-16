package playwright

import (
	"fmt"
	"reflect"
)

type Channel struct {
	EventEmitter
	guid       string
	connection *Connection
	object     interface{}
}

func (c *Channel) Send(method string, params interface{}) (interface{}, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	result, err := c.connection.SendMessageToServer(c.guid, method, params)
	if err != nil {
		return nil, fmt.Errorf("could not send message to server: %v", err)
	}
	if result == nil {
		return nil, nil
	}
	if reflect.TypeOf(result).Kind() == reflect.Map {
		for key := range result.(map[string]interface{}) {
			return result.(map[string]interface{})[key], nil
		}
	}
	return result, nil
}

func newChannel(connection *Connection, guid string) *Channel {
	channel := &Channel{
		connection: connection,
		guid:       guid,
	}
	channel.initEventEmitter()
	return channel
}
