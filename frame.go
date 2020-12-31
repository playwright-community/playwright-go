package playwright

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"sync"
	"time"
)

type frameImpl struct {
	channelOwner
	sync.RWMutex
	detached    bool
	page        *pageImpl
	name        string
	url         string
	parentFrame Frame
	childFrames []Frame
	loadStates  *safeStringSet
}

func newFrame(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *frameImpl {
	var loadStates *safeStringSet
	if ls, ok := initializer["loadStates"].([]string); ok {
		loadStates = newSafeStringSet(ls)
	} else {
		loadStates = newSafeStringSet([]string{})
	}
	bt := &frameImpl{
		name:        initializer["name"].(string),
		url:         initializer["url"].(string),
		loadStates:  loadStates,
		childFrames: make([]Frame, 0),
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)

	channelOwner := fromNullableChannel(initializer["parentFrame"])
	if channelOwner != nil {
		bt.parentFrame = channelOwner.(*frameImpl)
		bt.parentFrame.(*frameImpl).childFrames = append(bt.parentFrame.(*frameImpl).childFrames, bt)
	}

	bt.channel.On("navigated", bt.onFrameNavigated)
	bt.channel.On("loadstate", func(ev map[string]interface{}) {
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

func (f *frameImpl) URL() string {
	f.RLock()
	defer f.RUnlock()
	return f.url
}

func (f *frameImpl) Name() string {
	f.RLock()
	defer f.RUnlock()
	return f.name
}

func (f *frameImpl) SetContent(content string, options ...PageSetContentOptions) error {
	_, err := f.channel.Send("setContent", map[string]interface{}{
		"html": content,
	}, options)
	return err
}

func (f *frameImpl) Content() (string, error) {
	content, err := f.channel.Send("content")
	return content.(string), err
}

func (f *frameImpl) Goto(url string, options ...PageGotoOptions) (Response, error) {
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
	return channelOwner.(*responseImpl), nil
}

func (f *frameImpl) AddScriptTag(options PageAddScriptTagOptions) (ElementHandle, error) {
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
	return fromChannel(channel).(*elementHandleImpl), nil
}

func (f *frameImpl) AddStyleTag(options PageAddStyleTagOptions) (ElementHandle, error) {
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
	return fromChannel(channel).(*elementHandleImpl), nil
}

func (f *frameImpl) Page() Page {
	return f.page
}

func (f *frameImpl) WaitForLoadState(given ...string) {
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

func (f *frameImpl) WaitForEventCh(event string, predicate ...interface{}) <-chan interface{} {
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

func (f *frameImpl) WaitForEvent(event string, predicate ...interface{}) interface{} {
	return <-f.WaitForEventCh(event, predicate...)
}

func (f *frameImpl) WaitForNavigation(options ...PageWaitForNavigationOptions) (Response, error) {
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
			request := fromChannel(event["newDocument"].(map[string]interface{})["request"]).(*requestImpl)
			return request.Response()
		}
	}
	return nil, nil
}

func (f *frameImpl) onFrameNavigated(ev map[string]interface{}) {
	f.Lock()
	f.url = ev["url"].(string)
	f.name = ev["name"].(string)
	f.Unlock()
	f.Emit("navigated", ev)
	if f.page != nil {
		f.page.Emit("framenavigated", f)
	}
}

func (f *frameImpl) QuerySelector(selector string) (ElementHandle, error) {
	channel, err := f.channel.Send("querySelector", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	if channel == nil {
		return nil, nil
	}
	return fromChannel(channel).(*elementHandleImpl), nil
}

func (f *frameImpl) QuerySelectorAll(selector string) ([]ElementHandle, error) {
	channels, err := f.channel.Send("querySelectorAll", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	elements := make([]ElementHandle, 0)
	for _, channel := range channels.([]interface{}) {
		elements = append(elements, fromChannel(channel).(*elementHandleImpl))
	}
	return elements, nil
}

func (f *frameImpl) Evaluate(expression string, options ...interface{}) (interface{}, error) {
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

func (f *frameImpl) EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error) {
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

func (f *frameImpl) EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error) {
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

func (f *frameImpl) EvaluateHandle(expression string, options ...interface{}) (interface{}, error) {
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

func (f *frameImpl) Click(selector string, options ...PageClickOptions) error {
	_, err := f.channel.Send("click", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}

func (f *frameImpl) WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandle, error) {
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
	return channelOwner.(*elementHandleImpl), nil
}

func (f *frameImpl) DispatchEvent(selector, typ string, options ...PageDispatchEventOptions) error {
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

func (f *frameImpl) InnerText(selector string, options ...PageInnerTextOptions) (string, error) {
	innerText, err := f.channel.Send("innerText", map[string]interface{}{
		"selector": selector,
	}, options)
	if err != nil {
		return "", err
	}
	return innerText.(string), nil
}

func (f *frameImpl) InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error) {
	innerHTML, err := f.channel.Send("innerHTML", map[string]interface{}{
		"selector": selector,
	}, options)
	if err != nil {
		return "", err
	}
	return innerHTML.(string), nil
}

func (f *frameImpl) GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error) {
	attribute, err := f.channel.Send("getAttribute", map[string]interface{}{
		"selector": selector,
		"name":     name,
	}, options)
	if err != nil {
		return "", err
	}
	return attribute.(string), nil
}

func (f *frameImpl) Hover(selector string, options ...PageHoverOptions) error {
	_, err := f.channel.Send("hover", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}

func (e *frameImpl) SetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) error {
	_, err := e.channel.Send("setInputFiles", map[string]interface{}{
		"selector": selector,
		"files":    normalizeFilePayloads(files),
	}, options)
	return err
}

func (f *frameImpl) Type(selector, text string, options ...PageTypeOptions) error {
	_, err := f.channel.Send("type", map[string]interface{}{
		"selector": selector,
		"text":     text,
	}, options)
	return err
}

func (f *frameImpl) Press(selector, key string, options ...PagePressOptions) error {
	_, err := f.channel.Send("press", map[string]interface{}{
		"selector": selector,
		"key":      key,
	}, options)
	return err
}

func (f *frameImpl) Check(selector string, options ...FrameCheckOptions) error {
	_, err := f.channel.Send("check", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}

func (f *frameImpl) Uncheck(selector string, options ...FrameUncheckOptions) error {
	_, err := f.channel.Send("uncheck", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}

func (f *frameImpl) WaitForTimeout(timeout int) {
	time.Sleep(time.Duration(timeout) * time.Millisecond)
}

func (f *frameImpl) WaitForFunction(expression string, options ...FrameWaitForFunctionOptions) (JSHandle, error) {
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
	handle := result.(map[string]interface{})["handle"]
	if handle == nil {
		return nil, nil
	}
	return handle.(*jsHandleImpl), nil
}

func (f *frameImpl) Title() (string, error) {
	title, err := f.channel.Send("title")
	return title.(string), err
}

func (f *frameImpl) ChildFrames() []Frame {
	return f.childFrames
}

func (f *frameImpl) Dblclick(selector string, options ...FrameDblclickOptions) error {
	_, err := f.channel.Send("dblclick", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}

func (f *frameImpl) Fill(selector string, value string, options ...FrameFillOptions) error {
	_, err := f.channel.Send("fill", map[string]interface{}{
		"selector": selector,
		"value":    value,
	}, options)
	return err
}

func (f *frameImpl) Focus(selector string, options ...FrameFocusOptions) error {
	_, err := f.channel.Send("focus", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}

func (f *frameImpl) FrameElement() (ElementHandle, error) {
	elementHandle, err := f.channel.Send("frameElement")
	if err != nil {
		return nil, err
	}
	return elementHandle.(*elementHandleImpl), nil
}

func (f *frameImpl) IsDetached() bool {
	return f.detached
}

func (f *frameImpl) ParentFrame() Frame {
	return f.parentFrame
}

func (f *frameImpl) TextContent(selector string, options ...FrameTextContentOptions) (string, error) {
	textContent, err := f.channel.Send("textContent", map[string]interface{}{
		"selector": selector,
	}, options)
	if err != nil {
		return "", err
	}
	return textContent.(string), nil
}

func (f *frameImpl) Tap(selector string, options ...FrameTapOptions) error {
	_, err := f.channel.Send("tap", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}
