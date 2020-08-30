package playwright

import (
	"io/ioutil"
	"sync"
)

type Frame struct {
	ChannelOwner
	sync.RWMutex
	page *Page
	url  string
}

func (f *Frame) URL() string {
	f.RLock()
	defer f.RUnlock()
	return f.url
}

func (f *Frame) SetContent(content string, options ...PageSetContentOptions) error {
	_, err := f.channel.Send("setContent", map[string]interface{}{
		"html": content,
	}, options)
	return err
}

func (f *Frame) Content() (string, error) {
	content, err := f.channel.Send("content")
	return content.(string), err
}

func (f *Frame) Title() (string, error) {
	title, err := f.channel.Send("title")
	return title.(string), err
}

func (f *Frame) Goto(url string) (*Response, error) {
	response, err := f.channel.Send("goto", map[string]interface{}{
		"url": url,
	})
	if err != nil {
		return nil, err
	}
	obj := fromNullableChannel(response)
	if obj == nil {
		return nil, nil
	}
	return obj.(*Response), nil
}

func (f *Frame) AddScriptTag(options PageAddScriptTagOptions) (*ElementHandle, error) {
	if options.Path != nil {
		file, err := ioutil.ReadFile(*options.Path)
		if err != nil {
			return nil, err
		}
		options.Content = String(string(file))
		options.Path = nil
	}
	channelOwner, err := f.channel.Send("addScriptTag", options)
	if err != nil {
		return nil, err
	}
	return fromChannel(channelOwner).(*ElementHandle), nil
}

func (f *Frame) AddStyleTag(options PageAddStyleTagOptions) (*ElementHandle, error) {
	if options.Path != nil {
		file, err := ioutil.ReadFile(*options.Path)
		if err != nil {
			return nil, err
		}
		options.Content = String(string(file))
		options.Path = nil
	}
	channelOwner, err := f.channel.Send("addStyleTag", options)
	if err != nil {
		return nil, err
	}
	return fromChannel(channelOwner).(*ElementHandle), nil
}

func (f *Frame) Type(selector, text string, options ...PageTypeOptions) error {
	_, err := f.channel.Send("type", map[string]interface{}{
		"selector": selector,
		"text":     text,
	}, options)
	return err
}

func (f *Frame) Press(selector, key string, options ...PagePressOptions) error {
	_, err := f.channel.Send("press", map[string]interface{}{
		"selector": selector,
		"key":      key,
	}, options)
	return err
}

func (f *Frame) Page() *Page {
	return f.page
}

func (f *Frame) WaitForLoadState(state string) {
	succeed := make(chan bool, 1)
	f.Once("loadstate", func(ev ...interface{}) {
		gotState := ev[0].(string)
		if gotState == state {
			succeed <- true
		}
	})
	<-succeed
}

func (f *Frame) onFrameNavigated(event ...interface{}) {
	f.Lock()
	f.url = event[0].(map[string]interface{})["url"].(string)
	f.Unlock()
}

func (f *Frame) QuerySelector(selector string) (*ElementHandle, error) {
	channelOwner, err := f.channel.Send("querySelector", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	return fromChannel(channelOwner).(*ElementHandle), nil
}

func (f *Frame) QuerySelectorAll(selector string) ([]*ElementHandle, error) {
	channelOwner, err := f.channel.Send("querySelectorAll", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	elements := make([]*ElementHandle, 0)
	for _, channelOwner := range channelOwner.([]interface{}) {
		elements = append(elements, fromChannel(channelOwner).(*ElementHandle), nil)
	}
	return elements, nil
}

func (f *Frame) Evaluate(expression string, options ...interface{}) (interface{}, error) {
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
	result, err := f.channel.Send("evaluateExpression", map[string]interface{}{
		"expression": expression,
		"isFunction": !forceExpression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	return parseResult(result), nil
}

func (f *Frame) EvaluateOnSelector(selector string, expression string, options ...interface{}) (interface{}, error) {
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
	result, err := f.channel.Send("evalOnSelector", map[string]interface{}{
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

func (f *Frame) EvaluateOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error) {
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
	result, err := f.channel.Send("evalOnSelectorAll", map[string]interface{}{
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

func (f *Frame) EvaluateHandle(expression string, options ...interface{}) (*JSHandle, error) {
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
	result, err := f.channel.Send("evaluateExpressionHandle", map[string]interface{}{
		"expression": expression,
		"isFunction": !forceExpression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	channelOwner := fromChannel(result)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*JSHandle), nil
}

func (f *Frame) Click(selector string, options ...PageClickOptions) error {
	_, err := f.channel.Send("click", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}

func (f *Frame) WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (*ElementHandle, error) {
	channelOwner, err := f.channel.Send("waitForSelector", options)
	if err != nil {
		return nil, err
	}
	obj := fromNullableChannel(channelOwner)
	if obj == nil {
		return nil, nil
	}
	return obj.(*ElementHandle), nil
}

func (f *Frame) DispatchEvent(selector, typ string, options ...PageDispatchEventOptions) error {
	var eventInit interface{}
	if len(options) == 1 {
		eventInit = options[0].EventInit
	}
	_, err := f.channel.Send("dispatchEvent", map[string]interface{}{
		"selector":  selector,
		"type":      typ,
		"eventInit": serializeArgument(eventInit),
	})
	return err
}

func newFrame(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Frame {
	bt := &Frame{
		url: initializer["url"].(string),
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.channel.On("navigated", bt.onFrameNavigated)
	return bt
}
