package playwright

import (
	"io/ioutil"
	"sync"
)

type Frame struct {
	ChannelOwner
	sync.RWMutex
	page        *Page
	name        string
	url         string
	parentFrame *Frame
	childFrames []*Frame
	loadStates  []string
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
	channel, err := f.channel.Send("goto", map[string]interface{}{
		"url": url,
	})
	if err != nil {
		return nil, err
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*Response), nil
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
	channel, err := f.channel.Send("addScriptTag", options)
	if err != nil {
		return nil, err
	}
	return fromChannel(channel).(*ElementHandle), nil
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
	channel, err := f.channel.Send("addStyleTag", options)
	if err != nil {
		return nil, err
	}
	return fromChannel(channel).(*ElementHandle), nil
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

func (f *Frame) WaitForLoadState(given ...string) {
	state := "load"
	if len(given) == 1 {
		state = given[0]
	}
	for _, prevState := range f.loadStates {
		if prevState == state {
			return
		}
	}
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
	channel, err := f.channel.Send("querySelector", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	return fromChannel(channel).(*ElementHandle), nil
}

func (f *Frame) QuerySelectorAll(selector string) ([]*ElementHandle, error) {
	channels, err := f.channel.Send("querySelectorAll", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	elements := make([]*ElementHandle, 0)
	for _, channel := range channels.([]interface{}) {
		elements = append(elements, fromChannel(channel).(*ElementHandle))
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
	channel, err := f.channel.Send("waitForSelector", options, map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*ElementHandle), nil
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
	var loadStates []string
	if ls, ok := initializer["loadStates"].([]string); ok {
		loadStates = ls
	}
	bt := &Frame{
		name:       initializer["name"].(string),
		url:        initializer["url"].(string),
		loadStates: loadStates,
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)

	channelOwner := fromNullableChannel(initializer["parentFrame"])
	if channelOwner != nil {
		bt.parentFrame = channelOwner.(*Frame)
		bt.parentFrame.childFrames = append(bt.parentFrame.childFrames, bt)
	}

	bt.channel.On("navigated", bt.onFrameNavigated)
	bt.channel.On("loadstate", func(event ...interface{}) {
		ev := event[0].(map[string]interface{})
		if ev["add"] != nil {
			add := ev["add"].(string)
			addInLoadStates := false
			for i := 0; i < len(bt.loadStates); i++ {
				if bt.loadStates[i] == add {
					addInLoadStates = true
				}
			}
			if !addInLoadStates {
				bt.loadStates = append(bt.loadStates, add)
			}
			bt.Emit("loadstate", add)
		} else if ev["remove"] != nil {
			remove := ev["remove"].(string)
			newLoadstates := make([]string, 0)
			for i := 0; i < len(bt.loadStates); i++ {
				if bt.loadStates[i] != remove {
					newLoadstates = append(newLoadstates, bt.loadStates[i])
				}
			}
			if len(newLoadstates) != len(bt.loadStates) {
				bt.loadStates = newLoadstates
			}
		}
	})
	return bt
}
