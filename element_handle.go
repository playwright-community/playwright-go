package playwright

import "encoding/base64"

type ElementHandle struct {
	JSHandle
}

func (e *ElementHandle) AsElement() *ElementHandle {
	return e
}

func (e *ElementHandle) OwnerFrame() (*Frame, error) {
	channel, err := e.channel.Send("ownerFrame")
	if err != nil {
		return nil, err
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*Frame), nil
}

func (e *ElementHandle) ContentFrame() (*Frame, error) {
	channel, err := e.channel.Send("contentFrame")
	if err != nil {
		return nil, err
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*Frame), nil
}

func (e *ElementHandle) GetAttribute(name string) (string, error) {
	attribute, err := e.channel.Send("getAttribute", map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return "", err
	}
	if attribute == nil {
		return "", nil
	}
	return attribute.(string), nil
}

func (e *ElementHandle) TextContent() (string, error) {
	textContent, err := e.channel.Send("textContent")
	if err != nil {
		return "", err
	}
	return textContent.(string), nil
}

func (e *ElementHandle) InnerText() (string, error) {
	innerText, err := e.channel.Send("innerText")
	if err != nil {
		return "", err
	}
	return innerText.(string), nil
}

func (e *ElementHandle) InnerHTML() (string, error) {
	innerHTML, err := e.channel.Send("innerHTML")
	if err != nil {
		return "", err
	}
	return innerHTML.(string), nil
}

func (e *ElementHandle) DispatchEvent(typ string, initObjects ...interface{}) error {
	var initObject interface{}
	if len(initObjects) == 1 {
		initObject = initObjects[0]
	}
	_, err := e.channel.Send("dispatchEvent", map[string]interface{}{
		"type":      typ,
		"eventInit": serializeArgument(initObject),
	})
	return err
}

func (e *ElementHandle) Hover(options ...ElementHandleHoverOptions) error {
	_, err := e.channel.Send("hover", options)
	return err
}

func (e *ElementHandle) Click(options ...ElementHandleClickOptions) error {
	_, err := e.channel.Send("click", options)
	return err
}

func (e *ElementHandle) DblClick(options ...ElementHandleDblclickOptions) error {
	_, err := e.channel.Send("dblclick", options)
	return err
}

func (e *ElementHandle) QuerySelector(selector string) (*ElementHandle, error) {
	channel, err := e.channel.Send("querySelector", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	return fromChannel(channel).(*ElementHandle), nil
}

func (e *ElementHandle) ScrollIntoViewIfNeeded(options ...ElementHandleScrollIntoViewIfNeededOptions) error {
	_, err := e.channel.Send("scrollIntoViewIfNeeded", options)
	if err != nil {
		return err
	}
	return err
}

func (e *ElementHandle) SetInputFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error {
	_, err := e.channel.Send("setInputFiles", map[string]interface{}{
		"files": normalizeFilePayloads(files),
	}, options)
	return err
}

func newElementHandle(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *ElementHandle {
	bt := &ElementHandle{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}

func normalizeFilePayloads(files []InputFile) []map[string]string {
	out := make([]map[string]string, 0)
	for _, file := range files {
		// file.Buffer
		out = append(out, map[string]string{
			"name":     file.Name,
			"mimeType": file.MimeType,
			"buffer":   base64.StdEncoding.EncodeToString(file.Buffer),
		})
	}
	return out
}
