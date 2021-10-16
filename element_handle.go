package playwright

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

type elementhandleImpl struct {
	jshandleImpl
}

func (e *elementhandleImpl) AsElement() ElementHandle {
	return e
}

func (e *elementhandleImpl) OwnerFrame() (Frame, error) {
	channel, err := e.channel.Send("ownerFrame")
	if err != nil {
		return nil, err
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*frameImpl), nil
}

func (e *elementhandleImpl) ContentFrame() (Frame, error) {
	channel, err := e.channel.Send("contentFrame")
	if err != nil {
		return nil, err
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*frameImpl), nil
}

func (e *elementhandleImpl) GetAttribute(name string) (string, error) {
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

func (e *elementhandleImpl) TextContent() (string, error) {
	textContent, err := e.channel.Send("textContent")
	if err != nil {
		return "", err
	}
	return textContent.(string), nil
}

func (e *elementhandleImpl) InnerText() (string, error) {
	innerText, err := e.channel.Send("innerText")
	if err != nil {
		return "", err
	}
	return innerText.(string), nil
}

func (e *elementhandleImpl) InnerHTML() (string, error) {
	innerHTML, err := e.channel.Send("innerHTML")
	if err != nil {
		return "", err
	}
	return innerHTML.(string), nil
}

func (e *elementhandleImpl) DispatchEvent(typ string, initObjects ...interface{}) error {
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

func (e *elementhandleImpl) Hover(options ...ElementHandleHoverOptions) error {
	_, err := e.channel.Send("hover", options)
	return err
}

func (e *elementhandleImpl) Click(options ...ElementHandleClickOptions) error {
	_, err := e.channel.Send("click", options)
	return err
}

func (e *elementhandleImpl) Dblclick(options ...ElementHandleDblclickOptions) error {
	_, err := e.channel.Send("dblclick", options)
	return err
}

func (e *elementhandleImpl) QuerySelector(selector string) (ElementHandle, error) {
	channel, err := e.channel.Send("querySelector", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	if channel == nil {
		return nil, nil
	}
	return fromChannel(channel).(*elementhandleImpl), nil
}

func (e *elementhandleImpl) QuerySelectorAll(selector string) ([]ElementHandle, error) {
	channels, err := e.channel.Send("querySelectorAll", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	elements := make([]ElementHandle, 0)
	for _, channel := range channels.([]interface{}) {
		elements = append(elements, fromChannel(channel).(*elementhandleImpl))
	}
	return elements, nil
}

func (e *elementhandleImpl) EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error) {
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

func (e *elementhandleImpl) EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error) {
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

func (e *elementhandleImpl) ScrollIntoViewIfNeeded(options ...ElementHandleScrollIntoViewIfNeededOptions) error {
	_, err := e.channel.Send("scrollIntoViewIfNeeded", options)
	if err != nil {
		return err
	}
	return err
}

func (e *elementhandleImpl) SetInputFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error {
	_, err := e.channel.Send("setInputFiles", map[string]interface{}{
		"files": normalizeFilePayloads(files),
	}, options)
	return err
}

// Rect is the return structure for ElementHandle.BoundingBox()
type Rect struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	X      int `json:"x"`
	Y      int `json:"y"`
}

func (e *elementhandleImpl) BoundingBox() (*Rect, error) {
	boundingBox, err := e.channel.Send("boundingBox")
	if err != nil {
		return nil, err
	}
	out := &Rect{}
	remapMapToStruct(boundingBox, out)
	return out, nil
}

func (e *elementhandleImpl) Check(options ...ElementHandleCheckOptions) error {
	_, err := e.channel.Send("check", options)
	return err
}

func (e *elementhandleImpl) Uncheck(options ...ElementHandleUncheckOptions) error {
	_, err := e.channel.Send("uncheck", options)
	return err
}

func (e *elementhandleImpl) Press(key string, options ...ElementHandlePressOptions) error {
	_, err := e.channel.Send("press", map[string]interface{}{
		"key": key,
	}, options)
	return err
}

func (e *elementhandleImpl) Fill(value string, options ...ElementHandleFillOptions) error {
	_, err := e.channel.Send("fill", map[string]interface{}{
		"value": value,
	}, options)
	return err
}

func (e *elementhandleImpl) Type(value string, options ...ElementHandleTypeOptions) error {
	_, err := e.channel.Send("type", map[string]interface{}{
		"text": value,
	}, options)
	return err
}

func (e *elementhandleImpl) Focus() error {
	_, err := e.channel.Send("focus")
	return err
}

func (e *elementhandleImpl) SelectText(options ...ElementHandleSelectTextOptions) error {
	_, err := e.channel.Send("selectText", options)
	return err
}

func (e *elementhandleImpl) Screenshot(options ...ElementHandleScreenshotOptions) ([]byte, error) {
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

func (e *elementhandleImpl) Tap(options ...ElementHandleTapOptions) error {
	_, err := e.channel.Send("tap", options)
	return err
}

func (e *elementhandleImpl) SelectOption(values SelectOptionValues, options ...ElementHandleSelectOptionOptions) ([]string, error) {
	opts := convertSelectOptionSet(values)
	selected, err := e.channel.Send("selectOption", opts, options)
	if err != nil {
		return nil, err
	}

	return transformToStringList(selected), nil
}

func (e *elementhandleImpl) IsChecked() (bool, error) {
	checked, err := e.channel.Send("isChecked")
	if err != nil {
		return false, err
	}
	return checked.(bool), nil
}

func (e *elementhandleImpl) IsDisabled() (bool, error) {
	disabled, err := e.channel.Send("isDisabled")
	if err != nil {
		return false, err
	}
	return disabled.(bool), nil
}

func (e *elementhandleImpl) IsEditable() (bool, error) {
	editable, err := e.channel.Send("isEditable")
	if err != nil {
		return false, err
	}
	return editable.(bool), nil
}

func (e *elementhandleImpl) IsEnabled() (bool, error) {
	enabled, err := e.channel.Send("isEnabled")
	if err != nil {
		return false, err
	}
	return enabled.(bool), nil
}

func (e *elementhandleImpl) IsHidden() (bool, error) {
	hidden, err := e.channel.Send("isHidden")
	if err != nil {
		return false, err
	}
	return hidden.(bool), nil
}

func (e *elementhandleImpl) IsVisible() (bool, error) {
	visible, err := e.channel.Send("isVisible")
	if err != nil {
		return false, err
	}
	return visible.(bool), nil
}

func (e *elementhandleImpl) WaitForElementState(state string, options ...ElementHandleWaitForElementStateOptions) error {
	_, err := e.channel.Send("waitForElementState", map[string]interface{}{
		"state": state,
	}, options)
	if err != nil {
		return err
	}
	return nil
}

func (e *elementhandleImpl) WaitForSelector(selector string, options ...ElementHandleWaitForSelectorOptions) (ElementHandle, error) {
	ch, err := e.channel.Send("waitForSelector", map[string]interface{}{
		"selector": selector,
	}, options)
	if err != nil {
		return nil, err
	}

	channelOwner := fromNullableChannel(ch)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*elementhandleImpl), nil
}

func (e *elementhandleImpl) InputValue(options ...ElementHandleInputValueOptions) (string, error) {
	result, err := e.channel.Send("inputValue", options)
	return result.(string), err
}

func newElementHandle(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *elementhandleImpl {
	bt := &elementhandleImpl{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}

func normalizeFilePayloads(files []InputFile) []map[string]string {
	out := make([]map[string]string, 0)
	for _, file := range files {
		out = append(out, map[string]string{
			"name":     file.Name,
			"mimeType": file.MimeType,
			"buffer":   base64.StdEncoding.EncodeToString(file.Buffer),
		})
	}
	return out
}

func transformToStringList(in interface{}) []string {
	s := in.([]interface{})

	var out []string
	for _, v := range s {
		out = append(out, v.(string))
	}
	return out
}
