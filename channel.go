package playwright

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
	if mapV, ok := result.(map[string]interface{}); ok && len(mapV) <= 1 {
		for key := range mapV {
			return mapV[key], nil
		}
		return nil, nil
	}
	return result, nil
}

func (c *channel) SendNoReply(method string, options ...interface{}) {
	params := transformOptions(options...)
	_, err := c.connection.WrapAPICall(func() (interface{}, error) {
		return c.connection.sendMessageToServer(c.owner, method, params, true)
	}, false)
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
