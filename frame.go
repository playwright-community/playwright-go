package playwright

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"sync"
	"time"
)

type Frame struct {
	ChannelOwner
	sync.RWMutex
	detached    bool
	page        *Page
	name        string
	url         string
	parentFrame *Frame
	childFrames []*Frame
	loadStates  *safeStringSet
}

func (f *Frame) URL() string {
	f.RLock()
	defer f.RUnlock()
	return f.url
}

func (f *Frame) Name() string {
	f.RLock()
	defer f.RUnlock()
	return f.name
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

func (f *Frame) Goto(url string, options ...PageGotoOptions) (*Response, error) {
	channel, err := f.channel.Send("goto", map[string]interface{}{
		"url": url,
	}, options)
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

func (f *Frame) Page() *Page {
	return f.page
}

func (f *Frame) WaitForLoadState(given ...string) {
	state := "load"
	if len(given) == 1 {
		state = given[0]
	}
	if f.loadStates.Has(state) {
		return
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

func (f *Frame) WaitForEventCh(event string, predicate ...interface{}) <-chan interface{} {
	evChan := make(chan interface{}, 1)
	f.Once(event, func(ev ...interface{}) {
		if len(predicate) == 0 {
			evChan <- ev[0]
		} else if len(predicate) == 1 {
			result := reflect.ValueOf(predicate[0]).Call([]reflect.Value{reflect.ValueOf(ev[0])})
			if result[0].Bool() {
				evChan <- ev[0]
			}
		}
	})
	return evChan
}

func (f *Frame) WaitForEvent(event string, predicate ...interface{}) interface{} {
	return <-f.WaitForEventCh(event, predicate...)
}

func (f *Frame) WaitForNavigation(options ...PageWaitForNavigationOptions) (*Response, error) {
	option := PageWaitForNavigationOptions{}
	if len(options) == 1 {
		option = options[0]
	}
	if option.WaitUntil == nil {
		option.WaitUntil = String("load")
	}
	if option.Timeout == nil {
		option.Timeout = Int(f.page.timeoutSettings.NavigationTimeout())
	}
	deadline := time.After(time.Duration(*option.Timeout) * time.Millisecond)
	var matcher *urlMatcher
	if option.Url != nil {
		matcher = newURLMatcher(option.Url)
	}
	predicate := func(events ...interface{}) bool {
		ev := events[0].(map[string]interface{})
		if ev["error"] != nil {
			print("error")
		}
		return matcher == nil || matcher.Match(ev["url"].(string))
	}
	select {
	case <-deadline:
		return nil, fmt.Errorf("Timeout %dms exceeded.", *option.Timeout)
	case eventData := <-f.WaitForEventCh("navigated", predicate):
		event := eventData.(map[string]interface{})
		if event["newDocument"] != nil && event["newDocument"].(map[string]interface{})["request"] != nil {
			request := fromChannel(event["newDocument"].(map[string]interface{})["request"]).(*Request)
			return request.Response()
		}
	}
	return nil, nil
}

func (f *Frame) onFrameNavigated(event ...interface{}) {
	ev := event[0].(map[string]interface{})
	f.Lock()
	f.url = ev["url"].(string)
	f.name = ev["name"].(string)
	f.Unlock()
	f.Emit("navigated", event...)
	if f.page != nil {
		f.page.Emit("framenavigated", f)
	}
}

func (f *Frame) QuerySelector(selector string) (*ElementHandle, error) {
	channel, err := f.channel.Send("querySelector", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	if channel == nil {
		return nil, nil
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

func (f *Frame) EvaluateHandle(expression string, options ...interface{}) (interface{}, error) {
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
	return channelOwner, nil
}

func (f *Frame) Click(selector string, options ...PageClickOptions) error {
	_, err := f.channel.Send("click", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}

func (f *Frame) WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (*ElementHandle, error) {
	channel, err := f.channel.Send("waitForSelector", map[string]interface{}{
		"selector": selector,
	}, options)
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
	var loadStates *safeStringSet
	if ls, ok := initializer["loadStates"].([]string); ok {
		loadStates = newSafeStringSet(ls)
	} else {
		loadStates = newSafeStringSet([]string{})
	}
	bt := &Frame{
		name:        initializer["name"].(string),
		url:         initializer["url"].(string),
		loadStates:  loadStates,
		childFrames: make([]*Frame, 0),
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
			bt.loadStates.Add(add)
			bt.Emit("loadstate", add)
		} else if ev["remove"] != nil {
			remove := ev["remove"].(string)
			bt.loadStates.Remove(remove)
		}
	})
	return bt
}

func (f *Frame) InnerText(selector string, options ...PageInnerTextOptions) (string, error) {
	innerText, err := f.channel.Send("innerText", map[string]interface{}{
		"selector": selector,
	}, options)
	if err != nil {
		return "", err
	}
	return innerText.(string), nil
}

func (f *Frame) InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error) {
	innerHTML, err := f.channel.Send("innerHTML", map[string]interface{}{
		"selector": selector,
	}, options)
	if err != nil {
		return "", err
	}
	return innerHTML.(string), nil
}

func (f *Frame) GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error) {
	attribute, err := f.channel.Send("getAttribute", map[string]interface{}{
		"selector": selector,
		"name":     name,
	}, options)
	if err != nil {
		return "", err
	}
	return attribute.(string), nil
}

func (e *Frame) SetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) error {
	_, err := e.channel.Send("setInputFiles", map[string]interface{}{
		"selector": selector,
		"files":    normalizeFilePayloads(files),
	}, options)
	return err
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

func (f *Frame) Check(selector string, options ...FrameCheckOptions) error {
	_, err := f.channel.Send("check", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}

func (f *Frame) Uncheck(selector string, options ...FrameUncheckOptions) error {
	_, err := f.channel.Send("uncheck", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}

func (f *Frame) WaitForTimeout(timeout int) {
	time.Sleep(time.Duration(timeout) * time.Millisecond)
}

func (f *Frame) WaitForFunction(expression string, options ...FrameWaitForFunctionOptions) (*JSHandle, error) {
	var option FrameWaitForFunctionOptions
	if len(options) == 1 {
		option = options[0]
	}
	forceExpression := false
	if !isFunctionBody(expression) {
		forceExpression = true
	}
	result, err := f.channel.Send("evaluateExpression", map[string]interface{}{
		"expression": expression,
		"isFunction": !forceExpression,
		"arg":        serializeArgument(option.Arg),
		"timeout":    option.Timeout,
		"polling":    option.Polling,
	})
	if err != nil {
		return nil, err
	}
	return fromChannel(result).(*JSHandle), nil
}

func (f *Frame) Title() (string, error) {
	title, err := f.channel.Send("title")
	return title.(string), err
}
