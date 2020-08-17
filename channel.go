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
	if reflect.TypeOf(params).Kind() == reflect.Slice {
		val := reflect.ValueOf(params)
		if val.Len() == 1 {
			params = val.Index(0).Interface()
		} else if val.Len() == 0 {
			params = make(map[string]interface{})
		}
	}
	result, err := c.connection.SendMessageToServer(c.guid, method, params)
	if err != nil {
		return nil, fmt.Errorf("could not send message to server: %v", err)
	}
	if result == nil {
		return nil, nil
	}
	if reflect.TypeOf(result).Kind() == reflect.Map {
		mapV := result.(map[string]interface{})
		for key := range mapV {
			return mapV[key], nil
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
