package playwright

import "encoding/json"

type channel struct {
	eventEmitter
	guid       string
	connection *connection
	owner      *channelOwner // to avoid type conversion
	object     interface{}   // retain type info (for fromChannel needed)
}

func (c *channel) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"guid": c.guid,
	})
}

func (c *channel) Send(method string, options ...interface{}) (interface{}, error) {
	return c.connection.WrapAPICall(func() (interface{}, error) {
		return c.innerSend(method, options...).GetResultValue()
	}, false)
}

func (c *channel) SendReturnAsDict(method string, options ...interface{}) (map[string]interface{}, error) {
	ret, err := c.connection.WrapAPICall(func() (interface{}, error) {
		return c.innerSend(method, options...).GetResult()
	}, false)
	return ret.(map[string]interface{}), err
}

func (c *channel) innerSend(method string, options ...interface{}) *protocolCallback {
	params := transformOptions(options...)
	return c.connection.sendMessageToServer(c.owner, method, params, false)
}

func (c *channel) SendNoReply(method string, isInternal bool, options ...interface{}) {
	params := transformOptions(options...)
	_, err := c.connection.WrapAPICall(func() (interface{}, error) {
		return c.connection.sendMessageToServer(c.owner, method, params, true).GetResult()
	}, isInternal)
	if err != nil {
		logger.Printf("SendNoReply failed: %v\n", err)
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
