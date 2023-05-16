package playwright

import (
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/go-stack/stack"
)

type channel struct {
	eventEmitter
	guid       string
	connection *connection
	object     interface{}
}

func (c *channel) Send(method string, options ...interface{}) (interface{}, error) {
	return c.innerSend(method, false, false, options...)
}

func (c *channel) SendReturnAsDict(method string, options ...interface{}) (interface{}, error) {
	return c.innerSend(method, true, true, options...)
}

func (c *channel) innerSend(method string, returnAsDict bool, isInternal bool, options ...interface{}) (interface{}, error) {
	params := transformOptions(options...)
	skip := 3
	if method == "setNetworkInterceptionPatterns" {
		skip = 4
	}
	metadata := getMetadata(skip, isInternal)
	result, err := c.connection.SendMessageToServer(c.guid, method, metadata, params)
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
	_, err := c.connection.SendMessageToServer(c.guid, method, getMetadata(2, true), params)
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

// getMetadata use skip and isInternal to confirm the api name
func getMetadata(skip int, isInternal bool) map[string]interface{} {
	caller := stack.Caller(skip)
	apiName := ""
	if !isInternal {
		apiName = fmt.Sprintf("%n", caller)
	}
	metadata := make(map[string]interface{})
	metadata["location"] = serializeCallLocation(caller)
	metadata["apiName"] = apiName
	metadata["isInternal"] = isInternal
	return metadata
}

func serializeCallLocation(caller stack.Call) map[string]interface{} {
	line, _ := strconv.Atoi(fmt.Sprintf("%d", caller))
	return map[string]interface{}{
		"file": fmt.Sprintf("%s", caller),
		"line": line,
	}
}
