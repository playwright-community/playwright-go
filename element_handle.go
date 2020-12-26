package playwright

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

type ElementHandle struct {
	JSHandle
}

func (e *ElementHandle) AsElement() ElementHandleI {
	return e
}

func (e *ElementHandle) OwnerFrame() (FrameI, error) {
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

func (e *ElementHandle) ContentFrame() (FrameI, error) {
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

func (e *ElementHandle) QuerySelector(selector string) (ElementHandleI, error) {
	channel, err := e.channel.Send("querySelector", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	return fromChannel(channel).(*ElementHandle), nil
}

func (e *ElementHandle) QuerySelectorAll(selector string) ([]ElementHandleI, error) {
	channels, err := e.channel.Send("querySelectorAll", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	elements := make([]ElementHandleI, 0)
	for _, channel := range channels.([]interface{}) {
		elements = append(elements, fromChannel(channel).(*ElementHandle))
	}
	return elements, nil
}

func (e *ElementHandle) EvaluateOnSelector(selector string, expression string, options ...interface{}) (interface{}, error) {
	var arg interface{}
	forceExpression := false
	if !isFunctionBody(expression) {
		forceExpression = true
	}
	if len(options) == 1 {
		arg = options[0]
	} else if len(options) == 2 {
		arg = options[0]
		forceExpression = options[1].(bool)
	}
	result, err := e.channel.Send("evalOnSelector", map[string]interface{}{
		"selector":   selector,
		"expression": expression,
		"isFunction": !forceExpression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	return parseResult(result), nil
}

func (e *ElementHandle) EvaluateOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error) {
	var arg interface{}
	forceExpression := false
	if !isFunctionBody(expression) {
		forceExpression = true
	}
	if len(options) == 1 {
		arg = options[0]
	} else if len(options) == 2 {
		arg = options[0]
		forceExpression = options[1].(bool)
	}
	result, err := e.channel.Send("evalOnSelectorAll", map[string]interface{}{
		"selector":   selector,
		"expression": expression,
		"isFunction": !forceExpression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	return parseResult(result), nil
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

func (e *ElementHandle) BoundingBox() (*Rect, error) {
	boundingBox, err := e.channel.Send("boundingBox")
	if err != nil {
		return nil, err
	}
	out := &Rect{}
	remapMapToStruct(boundingBox, out)
	return out, nil
}

func (e *ElementHandle) Check(options ...ElementHandleCheckOptions) error {
	_, err := e.channel.Send("check", options)
	return err
}

func (e *ElementHandle) Uncheck(options ...ElementHandleUncheckOptions) error {
	_, err := e.channel.Send("uncheck", options)
	return err
}

func (e *ElementHandle) Press(options ...ElementHandlePressOptions) error {
	_, err := e.channel.Send("press", options)
	return err
}

func (e *ElementHandle) Fill(value string, options ...ElementHandleFillOptions) error {
	_, err := e.channel.Send("fill", map[string]interface{}{
		"value": value,
	}, options)
	return err
}

func (e *ElementHandle) Type(value string, options ...ElementHandleTypeOptions) error {
	_, err := e.channel.Send("type", map[string]interface{}{
		"value": value,
	}, options)
	return err
}

func (e *ElementHandle) Focus() error {
	_, err := e.channel.Send("focus")
	return err
}

func (e *ElementHandle) SelectText(options ...ElementHandleSelectTextOptions) error {
	_, err := e.channel.Send("selectText", options)
	return err
}

func (e *ElementHandle) Screenshot(options ...ElementHandleScreenshotOptions) ([]byte, error) {
	var path *string
	if len(options) > 0 {
		path = options[0].Path
	}
	data, err := e.channel.Send("screenshot", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message :%w", err)
	}
	image, err := base64.StdEncoding.DecodeString(data.(string))
	if err != nil {
		return nil, fmt.Errorf("could not decode base64 :%w", err)
	}
	if path != nil {
		if err := ioutil.WriteFile(*path, image, 0644); err != nil {
			return nil, err
		}
	}
	return image, nil
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
