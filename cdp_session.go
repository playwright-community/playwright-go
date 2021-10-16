package playwright

type cdpsessionImpl struct {
	channelOwner
}

func (c *cdpsessionImpl) Detach() error {
	_, err := c.channel.Send("detach")
	return err
}

func (c *cdpsessionImpl) Send(method string, params map[string]interface{}) (interface{}, error) {
	result, err := c.channel.Send("send", map[string]interface{}{
		"method": method,
		"params": params,
	})
	if err != nil {
		return nil, err
	}

	return result, err
}

func (c *cdpsessionImpl) onEvent(params map[string]interface{}) {
	c.Emit(params["method"].(string), params["params"])
}

func newCDPSession(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *cdpsessionImpl {
	bt := &cdpsessionImpl{}

	bt.createChannelOwner(bt, parent, objectType, guid, initializer)

	bt.channel.On("event", func(params map[string]interface{}) {
		bt.onEvent(params)
	})

	return bt
}
