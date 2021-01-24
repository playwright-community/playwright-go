package playwright

import (
	"fmt"
	"log"
	"reflect"
)

type channel struct {
	eventEmitter
	guid       string
	connection *connection
	object     interface{}
}

func (c *channel) Send(method string, options ...interface{}) (interface{}, error) {
	return c.innerSend(method, false, options...)
}

func (c *channel) SendReturnAsDict(method string, options ...interface{}) (interface{}, error) {
	return c.innerSend(method, true, options...)
}

func (c *channel) innerSend(method string, returnAsDict bool, options ...interface{}) (interface{}, error) {
	params := transformOptions(options...)
	result, err := c.connection.SendMessageToServer(c.guid, method, params)
	if err != nil {
		return nil, fmt.Errorf("could not send message to server: %w", err)
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
	_, err := c.connection.SendMessageToServer(c.guid, method, params)
	if err != nil {
		log.Printf("could not send message to server from noreply: %v", err)
	}
}

func newChannel(connection *connection, guid string) *channel {
	channel := &channel{
		connection: connection,
		guid:       guid,
	}
	channel.initEventEmitter()
	return channel
}
