package playwright

type ElementHandle struct {
	ChannelOwner
}

func (e *ElementHandle) QuerySelector(selector string) (*ElementHandle, error) {
	channelOwner, err := e.channel.Send("querySelector", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	return channelOwner.(*Channel).object.(*ElementHandle), nil
}

func (e *ElementHandle) TextContent() (string, error) {
	textContent, err := e.channel.Send("textContent")
	if err != nil {
		return "", err
	}
	return textContent.(string), nil
}

func newElementHandle(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *ElementHandle {
	bt := &ElementHandle{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
