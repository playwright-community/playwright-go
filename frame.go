package playwright

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
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
	loadStates  mapset.Set[string]
}

func newFrame(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *frameImpl {
	var loadStates mapset.Set[string]

	if ls, ok := initializer["loadStates"].([]string); ok {
		loadStates = mapset.NewSet[string](ls...)
	} else {
		loadStates = mapset.NewSet[string]()
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
	bt.channel.On("loadstate", bt.onLoadState)
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

func (f *frameImpl) SetContent(content string, options ...FrameSetContentOptions) error {
	_, err := f.channel.Send("setContent", map[string]interface{}{
		"html": content,
	}, options)
	return err
}

func (f *frameImpl) Content() (string, error) {
	content, err := f.channel.Send("content")
	if content == nil {
		return "", err
	}
	return content.(string), err
}

func (f *frameImpl) Goto(url string, options ...FrameGotoOptions) (Response, error) {
	channel, err := f.channel.Send("goto", map[string]interface{}{
		"url": url,
	}, options)
	if err != nil {
		return nil, fmt.Errorf("Frame.Goto %s: %w", url, err)
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		// navigation to about:blank or navigation to the same URL with a different hash
		return nil, nil
	}
	return channelOwner.(*responseImpl), nil
}

func (f *frameImpl) AddScriptTag(options FrameAddScriptTagOptions) (ElementHandle, error) {
	if options.Path != nil {
		file, err := os.ReadFile(*options.Path)
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

func (f *frameImpl) AddStyleTag(options FrameAddStyleTagOptions) (ElementHandle, error) {
	if options.Path != nil {
		file, err := os.ReadFile(*options.Path)
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

func (f *frameImpl) WaitForLoadState(options ...FrameWaitForLoadStateOptions) error {
	option := FrameWaitForLoadStateOptions{}
	if len(options) == 1 {
		option = options[0]
	}
	if option.State == nil {
		option.State = LoadStateLoad
	}
	return f.waitForLoadStateImpl(string(*option.State), option.Timeout, nil)
}

func (f *frameImpl) waitForLoadStateImpl(state string, timeout *float64, cb func() error) error {
	if f.loadStates.ContainsOne(state) {
		return nil
	}
	waiter, err := f.setNavigationWaiter(timeout)
	if err != nil {
		return err
	}
	waiter.WaitForEvent(f, "loadstate", func(payload interface{}) bool {
		gotState := payload.(string)
		return gotState == state
	})
	if cb == nil {
		_, err := waiter.Wait()
		return err
	} else {
		_, err := waiter.RunAndWait(cb)
		return err
	}
}

func (f *frameImpl) WaitForURL(url interface{}, options ...FrameWaitForURLOptions) error {
	if f.page == nil {
		return errors.New("frame is detached")
	}
	matcher := newURLMatcher(url, f.page.browserContext.options.BaseURL)
	if matcher.Matches(f.URL()) {
		state := "load"
		timeout := Float(f.page.timeoutSettings.NavigationTimeout())
		if len(options) == 1 {
			if options[0].WaitUntil != nil {
				state = string(*options[0].WaitUntil)
			}
			if options[0].Timeout != nil {
				timeout = options[0].Timeout
			}
		}
		return f.waitForLoadStateImpl(state, timeout, nil)
	}
	navigationOptions := FrameExpectNavigationOptions{URL: url}
	if len(options) > 0 {
		navigationOptions.Timeout = options[0].Timeout
		navigationOptions.WaitUntil = options[0].WaitUntil
	}
	if _, err := f.ExpectNavigation(nil, navigationOptions); err != nil {
		return err
	}
	return nil
}

func (f *frameImpl) ExpectNavigation(cb func() error, options ...FrameExpectNavigationOptions) (Response, error) {
	if f.page == nil {
		return nil, errors.New("frame is detached")
	}
	option := FrameExpectNavigationOptions{}
	if len(options) == 1 {
		option = options[0]
	}
	if option.WaitUntil == nil {
		option.WaitUntil = WaitUntilStateLoad
	}
	if option.Timeout == nil {
		option.Timeout = Float(f.page.timeoutSettings.NavigationTimeout())
	}
	deadline := time.Now().Add(time.Duration(*option.Timeout) * time.Millisecond)
	var matcher *urlMatcher
	if option.URL != nil {
		matcher = newURLMatcher(option.URL, f.page.browserContext.options.BaseURL)
	}
	predicate := func(events ...interface{}) bool {
		ev := events[0].(map[string]interface{})
		if ev["error"] != nil {
			print("error")
		}
		return matcher == nil || matcher.Matches(ev["url"].(string))
	}
	waiter, err := f.setNavigationWaiter(option.Timeout)
	if err != nil {
		return nil, err
	}

	eventData, err := waiter.WaitForEvent(f, "navigated", predicate).RunAndWait(cb)
	if err != nil || eventData == nil {
		return nil, err
	}

	t := time.Until(deadline).Milliseconds()
	if t > 0 {
		err = f.waitForLoadStateImpl(string(*option.WaitUntil), Float(float64(t)), nil)
		if err != nil {
			return nil, err
		}
	}
	event := eventData.(map[string]interface{})
	if event["newDocument"] != nil && event["newDocument"].(map[string]interface{})["request"] != nil {
		request := fromChannel(event["newDocument"].(map[string]interface{})["request"]).(*requestImpl)
		return request.Response()
	}
	return nil, nil
}

func (f *frameImpl) setNavigationWaiter(timeout *float64) (*waiter, error) {
	if f.page == nil {
		return nil, errors.New("page does not exist")
	}
	waiter := newWaiter()
	if timeout != nil {
		waiter.WithTimeout(*timeout)
	} else {
		waiter.WithTimeout(f.page.timeoutSettings.NavigationTimeout())
	}
	waiter.RejectOnEvent(f.page, "close", f.page.closeErrorWithReason())
	waiter.RejectOnEvent(f.page, "crash", fmt.Errorf("Navigation failed because page crashed!"))
	waiter.RejectOnEvent(f.page, "framedetached", fmt.Errorf("Navigating frame was detached!"), func(payload interface{}) bool {
		frame, ok := payload.(*frameImpl)
		if ok && frame == f {
			return true
		}
		return false
	})
	return waiter, nil
}

func (f *frameImpl) onFrameNavigated(ev map[string]interface{}) {
	f.Lock()
	f.url = ev["url"].(string)
	f.name = ev["name"].(string)
	f.Unlock()
	f.Emit("navigated", ev)
	_, ok := ev["error"]
	if !ok && f.page != nil {
		f.page.Emit("framenavigated", f)
	}
}

func (f *frameImpl) onLoadState(ev map[string]interface{}) {
	if ev["add"] != nil {
		add := ev["add"].(string)
		f.loadStates.Add(add)
		f.Emit("loadstate", add)
		if f.parentFrame == nil && f.page != nil {
			if add == "load" || add == "domcontentloaded" {
				f.Page().Emit(add, f.page)
			}
		}
	} else if ev["remove"] != nil {
		remove := ev["remove"].(string)
		f.loadStates.Remove(remove)
	}
}

func (f *frameImpl) QuerySelector(selector string, options ...FrameQuerySelectorOptions) (ElementHandle, error) {
	params := map[string]interface{}{
		"selector": selector,
	}
	if len(options) == 1 {
		params["strict"] = options[0].Strict
	}
	channel, err := f.channel.Send("querySelector", params)
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
	if len(options) == 1 {
		arg = options[0]
	}
	result, err := f.channel.Send("evaluateExpression", map[string]interface{}{
		"expression": expression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	return parseResult(result), nil
}

func (f *frameImpl) EvalOnSelector(selector string, expression string, arg interface{}, options ...FrameEvalOnSelectorOptions) (interface{}, error) {
	params := map[string]interface{}{
		"selector":   selector,
		"expression": expression,
		"arg":        serializeArgument(arg),
	}
	if len(options) == 1 && options[0].Strict != nil {
		params["strict"] = *options[0].Strict
	}

	result, err := f.channel.Send("evalOnSelector", params)
	if err != nil {
		return nil, err
	}
	return parseResult(result), nil
}

func (f *frameImpl) EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error) {
	var arg interface{}
	if len(options) == 1 {
		arg = options[0]
	}
	result, err := f.channel.Send("evalOnSelectorAll", map[string]interface{}{
		"selector":   selector,
		"expression": expression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	return parseResult(result), nil
}

func (f *frameImpl) EvaluateHandle(expression string, options ...interface{}) (JSHandle, error) {
	var arg interface{}
	if len(options) == 1 {
		arg = options[0]
	}
	result, err := f.channel.Send("evaluateExpressionHandle", map[string]interface{}{
		"expression": expression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	channelOwner := fromChannel(result)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(JSHandle), nil
}

func (f *frameImpl) Click(selector string, options ...FrameClickOptions) error {
	_, err := f.channel.Send("click", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}

func (f *frameImpl) WaitForSelector(selector string, options ...FrameWaitForSelectorOptions) (ElementHandle, error) {
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

func (f *frameImpl) DispatchEvent(selector, typ string, eventInit interface{}, options ...FrameDispatchEventOptions) error {
	_, err := f.channel.Send("dispatchEvent", map[string]interface{}{
		"selector":  selector,
		"type":      typ,
		"eventInit": serializeArgument(eventInit),
	})
	return err
}

func (f *frameImpl) InnerText(selector string, options ...FrameInnerTextOptions) (string, error) {
	innerText, err := f.channel.Send("innerText", map[string]interface{}{
		"selector": selector,
	}, options)
	if innerText == nil {
		return "", err
	}
	return innerText.(string), err
}

func (f *frameImpl) InnerHTML(selector string, options ...FrameInnerHTMLOptions) (string, error) {
	innerHTML, err := f.channel.Send("innerHTML", map[string]interface{}{
		"selector": selector,
	}, options)
	if innerHTML == nil {
		return "", err
	}
	return innerHTML.(string), err
}

func (f *frameImpl) GetAttribute(selector string, name string, options ...FrameGetAttributeOptions) (string, error) {
	attribute, err := f.channel.Send("getAttribute", map[string]interface{}{
		"selector": selector,
		"name":     name,
	}, options)
	if attribute == nil {
		return "", err
	}
	return attribute.(string), err
}

func (f *frameImpl) Hover(selector string, options ...FrameHoverOptions) error {
	_, err := f.channel.Send("hover", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}

func (f *frameImpl) SetInputFiles(selector string, files interface{}, options ...FrameSetInputFilesOptions) error {
	params, err := convertInputFiles(files, f.page.browserContext)
	if err != nil {
		return err
	}
	params.Selector = &selector
	_, err = f.channel.Send("setInputFiles", params, options)
	return err
}

func (f *frameImpl) Type(selector, text string, options ...FrameTypeOptions) error {
	_, err := f.channel.Send("type", map[string]interface{}{
		"selector": selector,
		"text":     text,
	}, options)
	return err
}

func (f *frameImpl) Press(selector, key string, options ...FramePressOptions) error {
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

func (f *frameImpl) WaitForTimeout(timeout float64) {
	time.Sleep(time.Duration(timeout) * time.Millisecond)
}

func (f *frameImpl) WaitForFunction(expression string, arg interface{}, options ...FrameWaitForFunctionOptions) (JSHandle, error) {
	var option FrameWaitForFunctionOptions
	if len(options) == 1 {
		option = options[0]
	}
	result, err := f.channel.Send("waitForFunction", map[string]interface{}{
		"expression": expression,
		"arg":        serializeArgument(arg),
		"timeout":    option.Timeout,
		"polling":    option.Polling,
	})
	if err != nil {
		return nil, err
	}
	handle := fromChannel(result)
	if handle == nil {
		return nil, nil
	}
	return handle.(*jsHandleImpl), nil
}

func (f *frameImpl) Title() (string, error) {
	title, err := f.channel.Send("title")
	if title == nil {
		return "", err
	}
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
	channel, err := f.channel.Send("frameElement")
	if err != nil {
		return nil, err
	}
	return fromChannel(channel).(*elementHandleImpl), nil
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
	if textContent == nil {
		return "", err
	}
	return textContent.(string), err
}

func (f *frameImpl) Tap(selector string, options ...FrameTapOptions) error {
	_, err := f.channel.Send("tap", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}

func (f *frameImpl) SelectOption(selector string, values SelectOptionValues, options ...FrameSelectOptionOptions) ([]string, error) {
	opts := convertSelectOptionSet(values)

	m := make(map[string]interface{})
	m["selector"] = selector
	for k, v := range opts {
		m[k] = v
	}
	selected, err := f.channel.Send("selectOption", m, options)
	if err != nil {
		return nil, err
	}

	return transformToStringList(selected), nil
}

func (f *frameImpl) IsChecked(selector string, options ...FrameIsCheckedOptions) (bool, error) {
	checked, err := f.channel.Send("isChecked", map[string]interface{}{
		"selector": selector,
	}, options)
	if err != nil {
		return false, err
	}
	return checked.(bool), nil
}

func (f *frameImpl) IsDisabled(selector string, options ...FrameIsDisabledOptions) (bool, error) {
	disabled, err := f.channel.Send("isDisabled", map[string]interface{}{
		"selector": selector,
	}, options)
	if err != nil {
		return false, err
	}
	return disabled.(bool), nil
}

func (f *frameImpl) IsEditable(selector string, options ...FrameIsEditableOptions) (bool, error) {
	editable, err := f.channel.Send("isEditable", map[string]interface{}{
		"selector": selector,
	}, options)
	if err != nil {
		return false, err
	}
	return editable.(bool), nil
}

func (f *frameImpl) IsEnabled(selector string, options ...FrameIsEnabledOptions) (bool, error) {
	enabled, err := f.channel.Send("isEnabled", map[string]interface{}{
		"selector": selector,
	}, options)
	if err != nil {
		return false, err
	}
	return enabled.(bool), nil
}

func (f *frameImpl) IsHidden(selector string, options ...FrameIsHiddenOptions) (bool, error) {
	hidden, err := f.channel.Send("isHidden", map[string]interface{}{
		"selector": selector,
	}, options)
	if err != nil {
		return false, err
	}
	return hidden.(bool), nil
}

func (f *frameImpl) IsVisible(selector string, options ...FrameIsVisibleOptions) (bool, error) {
	visible, err := f.channel.Send("isVisible", map[string]interface{}{
		"selector": selector,
	}, options)
	if err != nil {
		return false, err
	}
	return visible.(bool), nil
}

func (f *frameImpl) InputValue(selector string, options ...FrameInputValueOptions) (string, error) {
	value, err := f.channel.Send("inputValue", map[string]interface{}{
		"selector": selector,
	}, options)
	if value == nil {
		return "", err
	}
	return value.(string), err
}

func (f *frameImpl) DragAndDrop(source, target string, options ...FrameDragAndDropOptions) error {
	_, err := f.channel.Send("dragAndDrop", map[string]interface{}{
		"source": source,
		"target": target,
	}, options)
	return err
}

func (f *frameImpl) SetChecked(selector string, checked bool, options ...FrameSetCheckedOptions) error {
	if checked {
		_, err := f.channel.Send("check", map[string]interface{}{
			"selector": selector,
		}, options)
		return err
	} else {
		_, err := f.channel.Send("uncheck", map[string]interface{}{
			"selector": selector,
		}, options)
		return err
	}
}

func (f *frameImpl) Locator(selector string, options ...FrameLocatorOptions) Locator {
	var option LocatorLocatorOptions
	if len(options) == 1 {
		option = LocatorLocatorOptions(options[0])
	}
	return newLocator(f, selector, option)
}

func (f *frameImpl) GetByAltText(text interface{}, options ...FrameGetByAltTextOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return f.Locator(getByAltTextSelector(text, exact))
}

func (f *frameImpl) GetByLabel(text interface{}, options ...FrameGetByLabelOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return f.Locator(getByLabelSelector(text, exact))
}

func (f *frameImpl) GetByPlaceholder(text interface{}, options ...FrameGetByPlaceholderOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return f.Locator(getByPlaceholderSelector(text, exact))
}

func (f *frameImpl) GetByRole(role AriaRole, options ...FrameGetByRoleOptions) Locator {
	if len(options) == 1 {
		return f.Locator(getByRoleSelector(role, LocatorGetByRoleOptions(options[0])))
	}
	return f.Locator(getByRoleSelector(role))
}

func (f *frameImpl) GetByTestId(testId interface{}) Locator {
	return f.Locator(getByTestIdSelector(getTestIdAttributeName(), testId))
}

func (f *frameImpl) GetByText(text interface{}, options ...FrameGetByTextOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return f.Locator(getByTextSelector(text, exact))
}

func (f *frameImpl) GetByTitle(text interface{}, options ...FrameGetByTitleOptions) Locator {
	exact := false
	if len(options) == 1 {
		if *options[0].Exact {
			exact = true
		}
	}
	return f.Locator(getByTitleSelector(text, exact))
}

func (f *frameImpl) FrameLocator(selector string) FrameLocator {
	return newFrameLocator(f, selector)
}

func (f *frameImpl) highlight(selector string) error {
	_, err := f.channel.Send("highlight", map[string]interface{}{
		"selector": selector,
	})
	return err
}

func (f *frameImpl) queryCount(selector string) (int, error) {
	response, err := f.channel.Send("queryCount", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return 0, err
	}
	return int(response.(float64)), nil
}
