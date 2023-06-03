package playwright

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"time"
)

type pageImpl struct {
	channelOwner
	isClosed        bool
	closedOrCrashed chan bool
	video           *videoImpl
	mouse           *mouseImpl
	keyboard        *keyboardImpl
	touchscreen     *touchscreenImpl
	timeoutSettings *timeoutSettings
	browserContext  *browserContextImpl
	frames          []Frame
	workers         []Worker
	mainFrame       Frame
	routes          []*routeHandlerEntry
	viewportSize    ViewportSize
	ownedContext    BrowserContext
	bindings        map[string]BindingCallFunction
}

func (p *pageImpl) Context() BrowserContext {
	return p.browserContext
}

func (p *pageImpl) Close(options ...PageCloseOptions) error {
	_, err := p.channel.Send("close", options)
	if err != nil {
		return err
	}
	if p.ownedContext != nil {
		return p.ownedContext.Close()
	}
	return nil
}

func (p *pageImpl) InnerText(selector string, options ...PageInnerTextOptions) (string, error) {
	return p.mainFrame.InnerText(selector, options...)
}

func (p *pageImpl) InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error) {
	return p.mainFrame.InnerHTML(selector, options...)
}

func (p *pageImpl) Opener() (Page, error) {
	channel := p.initializer["opener"]
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		// not popup page or opener has been closed
		return nil, nil
	}
	return channelOwner.(*pageImpl), nil
}

func (p *pageImpl) MainFrame() Frame {
	return p.mainFrame
}

// PageFrameOptions is the option struct for Page.Frame()
type PageFrameOptions struct {
	Name *string
	URL  interface{}
}

func (p *pageImpl) Frame(options PageFrameOptions) Frame {
	var matcher *urlMatcher
	if options.URL != nil {
		matcher = newURLMatcher(options.URL, p.browserContext.options.BaseURL)
	}

	for _, f := range p.frames {
		if options.Name != nil && f.Name() == *options.Name {
			return f
		}

		if options.URL != nil && matcher != nil && matcher.Matches(f.URL()) {
			return f
		}
	}

	return nil
}

func (p *pageImpl) Frames() []Frame {
	return p.frames
}

func (p *pageImpl) SetDefaultNavigationTimeout(timeout float64) {
	p.timeoutSettings.SetNavigationTimeout(timeout)
	p.channel.SendNoReply("setDefaultNavigationTimeoutNoReply", map[string]interface{}{
		"timeout": timeout,
	})
}

func (p *pageImpl) SetDefaultTimeout(timeout float64) {
	p.timeoutSettings.SetTimeout(timeout)
	p.channel.SendNoReply("setDefaultTimeoutNoReply", map[string]interface{}{
		"timeout": timeout,
	})
}

func (p *pageImpl) QuerySelector(selector string) (ElementHandle, error) {
	return p.mainFrame.QuerySelector(selector)
}

func (p *pageImpl) QuerySelectorAll(selector string) ([]ElementHandle, error) {
	return p.mainFrame.QuerySelectorAll(selector)
}

func (p *pageImpl) WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandle, error) {
	return p.mainFrame.WaitForSelector(selector, options...)
}

func (p *pageImpl) DispatchEvent(selector string, typ string, options ...PageDispatchEventOptions) error {
	return p.mainFrame.DispatchEvent(selector, typ, nil, options...)
}

func (p *pageImpl) Evaluate(expression string, options ...interface{}) (interface{}, error) {
	return p.mainFrame.Evaluate(expression, options...)
}

func (p *pageImpl) EvaluateHandle(expression string, options ...interface{}) (JSHandle, error) {
	return p.mainFrame.EvaluateHandle(expression, options...)
}

func (p *pageImpl) EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error) {
	return p.mainFrame.EvalOnSelector(selector, expression, options...)
}

func (p *pageImpl) EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error) {
	return p.mainFrame.EvalOnSelectorAll(selector, expression, options...)
}

func (p *pageImpl) AddScriptTag(options PageAddScriptTagOptions) (ElementHandle, error) {
	return p.mainFrame.AddScriptTag(options)
}

func (p *pageImpl) AddStyleTag(options PageAddStyleTagOptions) (ElementHandle, error) {
	return p.mainFrame.AddStyleTag(options)
}

func (p *pageImpl) SetExtraHTTPHeaders(headers map[string]string) error {
	_, err := p.channel.Send("setExtraHTTPHeaders", map[string]interface{}{
		"headers": serializeMapToNameAndValue(headers),
	})
	return err
}

func (p *pageImpl) URL() string {
	return p.mainFrame.URL()
}

func (p *pageImpl) Unroute(url interface{}, handlers ...routeHandler) error {
	p.Lock()
	defer p.Unlock()

	routes, err := unroute(p.channel, p.routes, url, handlers...)
	if err != nil {
		return err
	}
	p.routes = routes

	return nil
}

func (p *pageImpl) Content() (string, error) {
	return p.mainFrame.Content()
}

func (p *pageImpl) SetContent(content string, options ...PageSetContentOptions) error {
	return p.mainFrame.SetContent(content, options...)
}

func (p *pageImpl) Goto(url string, options ...PageGotoOptions) (Response, error) {
	return p.mainFrame.Goto(url, options...)
}

func (p *pageImpl) Reload(options ...PageReloadOptions) (Response, error) {
	channel, err := p.channel.Send("reload", options)
	if err != nil {
		return nil, err
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*responseImpl), nil
}

func (p *pageImpl) WaitForLoadState(options ...PageWaitForLoadStateOptions) error {
	return p.mainFrame.WaitForLoadState(options...)
}

func (p *pageImpl) GoBack(options ...PageGoBackOptions) (Response, error) {
	channel, err := p.channel.Send("goBack", options)
	if err != nil {
		return nil, err
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		// can not go back
		return nil, nil
	}
	return channelOwner.(*responseImpl), nil
}

func (p *pageImpl) GoForward(options ...PageGoForwardOptions) (Response, error) {
	channel, err := p.channel.Send("goForward", options)
	if err != nil {
		return nil, err
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		// can not go forward
		return nil, nil
	}
	return channelOwner.(*responseImpl), nil
}

func (p *pageImpl) EmulateMedia(options ...PageEmulateMediaOptions) error {
	_, err := p.channel.Send("emulateMedia", options)
	if err != nil {
		return err
	}
	return err
}

// ViewportSize represents the viewport size
type ViewportSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (p *pageImpl) SetViewportSize(width, height int) error {
	_, err := p.channel.Send("setViewportSize", map[string]interface{}{
		"viewportSize": map[string]interface{}{
			"width":  width,
			"height": height,
		},
	})
	if err != nil {
		return err
	}
	p.viewportSize.Width = width
	p.viewportSize.Height = height
	return nil
}

func (p *pageImpl) ViewportSize() ViewportSize {
	return p.viewportSize
}

func (p *pageImpl) BringToFront() error {
	_, err := p.channel.Send("bringToFront")
	return err
}

func (p *pageImpl) Type(selector, text string, options ...PageTypeOptions) error {
	return p.mainFrame.Type(selector, text, options...)
}

func (p *pageImpl) Fill(selector, text string, options ...FrameFillOptions) error {
	return p.mainFrame.Fill(selector, text, options...)
}

func (p *pageImpl) Press(selector, key string, options ...PagePressOptions) error {
	return p.mainFrame.Press(selector, key, options...)
}

func (p *pageImpl) Title() (string, error) {
	return p.mainFrame.Title()
}

func (p *pageImpl) Workers() []Worker {
	return p.workers
}

func (p *pageImpl) Screenshot(options ...PageScreenshotOptions) ([]byte, error) {
	var path *string
	if len(options) > 0 {
		path = options[0].Path
	}
	data, err := p.channel.Send("screenshot", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message :%w", err)
	}
	image, err := base64.StdEncoding.DecodeString(data.(string))
	if err != nil {
		return nil, fmt.Errorf("could not decode base64 :%w", err)
	}
	if path != nil {
		if err := os.WriteFile(*path, image, 0644); err != nil {
			return nil, err
		}
	}
	return image, nil
}

func (p *pageImpl) PDF(options ...PagePdfOptions) ([]byte, error) {
	var path *string
	if len(options) > 0 {
		path = options[0].Path
	}
	data, err := p.channel.Send("pdf", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message :%w", err)
	}
	pdf, err := base64.StdEncoding.DecodeString(data.(string))
	if err != nil {
		return nil, fmt.Errorf("could not decode base64 :%w", err)
	}
	if path != nil {
		if err := os.WriteFile(*path, pdf, 0644); err != nil {
			return nil, err
		}
	}
	return pdf, nil
}

func (p *pageImpl) Click(selector string, options ...PageClickOptions) error {
	return p.mainFrame.Click(selector, options...)
}

func (p *pageImpl) WaitForEvent(event string, options ...PageWaitForEventOptions) (interface{}, error) {
	return p.waiterForEvent(event, options...).Wait()
}

func (p *pageImpl) waiterForEvent(event string, options ...PageWaitForEventOptions) *waiter {
	timeout := p.timeoutSettings.Timeout()
	var predicate interface{} = nil
	if len(options) == 1 {
		if options[0].Timeout != nil {
			timeout = *options[0].Timeout
		}
		predicate = options[0].Predicate
	}
	waiter := newWaiter().WithTimeout(timeout)
	waiter.RejectOnEvent(p, "close", errors.New("page closed"))
	waiter.RejectOnEvent(p, "crash", errors.New("page crashed"))
	return waiter.WaitForEvent(p, event, predicate)
}

func (p *pageImpl) WaitForNavigation(options ...PageWaitForNavigationOptions) (Response, error) {
	return p.mainFrame.WaitForNavigation(options...)
}

func (p *pageImpl) WaitForRequest(url interface{}, options ...PageWaitForRequestOptions) (Request, error) {
	req, err := p.waiterForRequest(url, options...).Wait()
	if req == nil {
		return nil, err
	}
	return req.(*requestImpl), nil
}

func (p *pageImpl) waiterForRequest(url interface{}, options ...PageWaitForRequestOptions) *waiter {
	option := PageWaitForRequestOptions{}
	if len(options) == 1 {
		option = options[0]
	}
	if option.Timeout == nil {
		option.Timeout = Float(p.timeoutSettings.timeout)
	}
	var matcher *urlMatcher
	if url != nil {
		matcher = newURLMatcher(url, p.browserContext.options.BaseURL)
	}
	predicate := func(req *requestImpl) bool {
		if matcher != nil {
			return matcher.Matches(req.URL())
		}
		return true
	}

	waiter := newWaiter().WithTimeout(*option.Timeout)
	return waiter.WaitForEvent(p, "request", predicate)
}

func (p *pageImpl) WaitForResponse(url interface{}, options ...PageWaitForResponseOptions) (Response, error) {
	res, err := p.waiterForResponse(url, options...).Wait()
	if res == nil {
		return nil, err
	}
	return res.(*responseImpl), nil
}

func (p *pageImpl) waiterForResponse(url interface{}, options ...PageWaitForResponseOptions) *waiter {
	option := PageWaitForResponseOptions{}
	if len(options) == 1 {
		option = options[0]
	}
	if option.Timeout == nil {
		option.Timeout = Float(p.timeoutSettings.timeout)
	}
	var matcher *urlMatcher
	if url != nil {
		matcher = newURLMatcher(url, p.browserContext.options.BaseURL)
	}
	predicate := func(req *responseImpl) bool {
		if matcher != nil {
			return matcher.Matches(req.URL())
		}
		return true
	}

	waiter := newWaiter().WithTimeout(*option.Timeout)
	return waiter.WaitForEvent(p, "response", predicate)
}

func (p *pageImpl) ExpectEvent(event string, cb func() error, options ...PageWaitForEventOptions) (interface{}, error) {
	return p.waiterForEvent(event, options...).Expect(cb)
}

func (p *pageImpl) ExpectNavigation(cb func() error, options ...PageWaitForNavigationOptions) (Response, error) {
	option := PageWaitForNavigationOptions{}
	if len(options) == 1 {
		option = options[0]
	}
	if option.WaitUntil == nil {
		option.WaitUntil = WaitUntilStateLoad
	}
	if option.Timeout == nil {
		option.Timeout = Float(p.timeoutSettings.NavigationTimeout())
	}
	deadline := time.Now().Add(time.Duration(*option.Timeout) * time.Millisecond)
	var matcher *urlMatcher
	if option.URL != nil {
		matcher = newURLMatcher(option.URL, p.browserContext.options.BaseURL)
	}
	predicate := func(events ...interface{}) bool {
		ev := events[0].(map[string]interface{})
		if ev["error"] != nil {
			print("error")
		}
		return matcher == nil || matcher.Matches(ev["url"].(string))
	}
	waiter := p.mainFrame.(*frameImpl).setNavigationWaiter(option.Timeout)

	eventData, err := waiter.WaitForEvent(p.mainFrame.(*frameImpl), "navigated", predicate).Expect(cb)
	if err != nil || eventData == nil {
		return nil, err
	}

	t := time.Until(deadline).Milliseconds()
	if t > 0 {
		err = p.mainFrame.(*frameImpl).waitForLoadStateImpl(string(*option.WaitUntil), Float(float64(t)), nil)
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

func (p *pageImpl) ExpectConsoleMessage(cb func() error, options ...PageExpectConsoleMessageOptions) (ConsoleMessage, error) {
	option := PageWaitForEventOptions{}
	if len(options) == 1 {
		option.Timeout = options[0].Timeout
		option.Predicate = options[0].Predicate
	}
	ret, err := p.waiterForEvent("console", option).Expect(cb)
	if ret == nil {
		return nil, err
	}
	return ret.(*consoleMessageImpl), err
}

func (p *pageImpl) ExpectedDialog(cb func() error) (Dialog, error) {
	ret, err := newWaiter().WaitForEvent(p, "dialog", nil).Expect(cb)
	if ret == nil {
		return nil, err
	}
	return ret.(*dialogImpl), err
}

func (p *pageImpl) ExpectDownload(cb func() error, options ...PageExpectDownloadOptions) (Download, error) {
	option := PageWaitForEventOptions{}
	if len(options) == 1 {
		option.Timeout = options[0].Timeout
		option.Predicate = options[0].Predicate
	}
	ret, err := p.waiterForEvent("download", option).Expect(cb)
	if ret == nil {
		return nil, err
	}
	return ret.(*downloadImpl), err
}

func (p *pageImpl) ExpectFileChooser(cb func() error, options ...PageExpectFileChooserOptions) (FileChooser, error) {
	option := PageWaitForEventOptions{}
	if len(options) == 1 {
		option.Timeout = options[0].Timeout
		option.Predicate = options[0].Predicate
	}
	ret, err := p.waiterForEvent("filechooser", option).Expect(cb)
	if ret == nil {
		return nil, err
	}
	return ret.(*fileChooserImpl), err
}

func (p *pageImpl) ExpectLoadState(cb func() error, options ...PageWaitForLoadStateOptions) error {
	option := PageWaitForLoadStateOptions{}
	if len(options) == 1 {
		option = options[0]
	}
	if option.State == nil {
		option.State = LoadStateLoad
	}
	if option.Timeout == nil {
		option.Timeout = Float(p.timeoutSettings.NavigationTimeout())
	}
	return p.mainFrame.(*frameImpl).waitForLoadStateImpl(string(*option.State), option.Timeout, cb)
}

func (p *pageImpl) ExpectPopup(cb func() error, options ...PageExpectPopupOptions) (Page, error) {
	option := PageWaitForEventOptions{}
	if len(options) == 1 {
		option.Timeout = options[0].Timeout
		option.Predicate = options[0].Predicate
	}
	ret, err := p.waiterForEvent("popup", option).Expect(cb)
	if ret == nil {
		return nil, err
	}
	return ret.(*pageImpl), err
}

func (p *pageImpl) ExpectResponse(url interface{}, cb func() error, options ...PageWaitForResponseOptions) (Response, error) {
	ret, err := p.waiterForResponse(url, options...).Expect(cb)
	if ret == nil {
		return nil, err
	}
	return ret.(*responseImpl), err
}

func (p *pageImpl) ExpectRequest(url interface{}, cb func() error, options ...PageWaitForRequestOptions) (Request, error) {
	ret, err := p.waiterForRequest(url, options...).Expect(cb)
	if ret == nil {
		return nil, err
	}
	return ret.(*requestImpl), err
}

func (p *pageImpl) ExpectWorker(cb func() error, options ...PageExpectWorkerOptions) (Worker, error) {
	option := PageWaitForEventOptions{}
	if len(options) == 1 {
		option.Timeout = options[0].Timeout
		option.Predicate = options[0].Predicate
	}
	ret, err := p.waiterForEvent("worker", option).Expect(cb)
	if ret == nil {
		return nil, err
	}
	return ret.(*workerImpl), err
}

func (p *pageImpl) Route(url interface{}, handler routeHandler, times ...int) error {
	p.Lock()
	defer p.Unlock()
	p.routes = append(p.routes, newRouteHandlerEntry(newURLMatcher(url, p.browserContext.options.BaseURL), handler, times...))
	if len(p.routes) == 1 {
		_, err := p.channel.Send("setNetworkInterceptionEnabled", map[string]interface{}{
			"enabled": true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *pageImpl) GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error) {
	return p.mainFrame.GetAttribute(selector, name, options...)
}

func (p *pageImpl) Hover(selector string, options ...PageHoverOptions) error {
	return p.mainFrame.Hover(selector, options...)
}

func (p *pageImpl) IsClosed() bool {
	return p.isClosed
}

func (p *pageImpl) AddInitScript(options PageAddInitScriptOptions) error {
	var source string
	if options.Script != nil {
		source = *options.Script
	}
	if options.Path != nil {
		content, err := os.ReadFile(*options.Path)
		if err != nil {
			return err
		}
		source = string(content)
	}
	_, err := p.channel.Send("addInitScript", map[string]interface{}{
		"source": source,
	})
	return err
}

func (p *pageImpl) Keyboard() Keyboard {
	return p.keyboard
}
func (p *pageImpl) Mouse() Mouse {
	return p.mouse
}

func (p *pageImpl) Touchscreen() Touchscreen {
	return p.touchscreen
}

func (p *pageImpl) setBrowserContext(browserContext *browserContextImpl) {
	p.browserContext = browserContext
	p.timeoutSettings = newTimeoutSettings(browserContext.timeoutSettings)
}

func newPage(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *pageImpl {
	viewportSize := &ViewportSize{}
	if _, ok := initializer["viewportSize"].(map[string]interface{}); ok {
		viewportSize.Height = int(initializer["viewportSize"].(map[string]interface{})["height"].(float64))
		viewportSize.Width = int(initializer["viewportSize"].(map[string]interface{})["width"].(float64))
	}
	bt := &pageImpl{
		mainFrame:       fromChannel(initializer["mainFrame"]).(*frameImpl),
		workers:         make([]Worker, 0),
		routes:          make([]*routeHandlerEntry, 0),
		bindings:        make(map[string]BindingCallFunction),
		viewportSize:    *viewportSize,
		timeoutSettings: newTimeoutSettings(nil),
	}
	bt.frames = []Frame{bt.mainFrame}
	bt.mainFrame.(*frameImpl).page = bt
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.mouse = newMouse(bt.channel)
	bt.keyboard = newKeyboard(bt.channel)
	bt.touchscreen = newTouchscreen(bt.channel)
	bt.channel.On("bindingCall", func(params map[string]interface{}) {
		bt.onBinding(fromChannel(params["binding"]).(*bindingCallImpl))
	})
	bt.channel.On("close", bt.onClose)
	bt.channel.On("console", func(ev map[string]interface{}) {
		bt.Emit("console", fromChannel(ev["message"]))
	})
	bt.channel.On("crash", func() {
		bt.Emit("crash")
	})
	bt.channel.On("dialog", func(ev map[string]interface{}) {
		go func() {
			bt.Emit("dialog", fromChannel(ev["dialog"]))
		}()
	})
	bt.channel.On("domcontentloaded", func() {
		bt.Emit("domcontentloaded")
	})
	bt.channel.On("fileChooser", func(ev map[string]interface{}) {
		bt.Emit("filechooser", newFileChooser(bt, fromChannel(ev["element"]).(*elementHandleImpl), ev["isMultiple"].(bool)))
	})
	bt.channel.On("frameAttached", func(ev map[string]interface{}) {
		bt.onFrameAttached(fromChannel(ev["frame"]).(*frameImpl))
	})
	bt.channel.On("frameDetached", func(ev map[string]interface{}) {
		bt.onFrameDetached(fromChannel(ev["frame"]).(*frameImpl))
	})
	bt.channel.On(
		"load", func(ev map[string]interface{}) {
			bt.Emit("load")
		},
	)
	bt.channel.On(
		"pageError", func(ev map[string]interface{}) {
			err := errorPayload{}
			remapMapToStruct(ev["error"].(map[string]interface{})["error"], &err)
			bt.Emit("pageerror", parseError(err))
		},
	)
	bt.channel.On("popup", func(ev map[string]interface{}) {
		bt.Emit("popup", fromChannel(ev["page"]))
	})
	bt.channel.On("route", func(ev map[string]interface{}) {
		bt.onRoute(fromChannel(ev["route"]).(*routeImpl))
	})
	bt.channel.On("download", func(ev map[string]interface{}) {
		url := ev["url"].(string)
		suggestedFilename := ev["suggestedFilename"].(string)
		artifact := fromChannel(ev["artifact"]).(*artifactImpl)
		bt.Emit("download", newDownload(bt, url, suggestedFilename, artifact))
	})
	bt.channel.On("video", func(params map[string]interface{}) {
		artifact := fromChannel(params["artifact"]).(*artifactImpl)
		bt.Video().(*videoImpl).artifactReady(artifact)
	})
	bt.channel.On("webSocket", func(ev map[string]interface{}) {
		bt.Emit("websocket", fromChannel(ev["webSocket"]).(*webSocketImpl))
	})

	bt.channel.On("worker", func(ev map[string]interface{}) {
		bt.onWorker(fromChannel(ev["worker"]).(*workerImpl))
	})
	bt.closedOrCrashed = make(chan bool, 1)
	bt.On("close", func() {
		select {
		case bt.closedOrCrashed <- true:
		default:
		}
	})
	bt.On("crash", func() {
		select {
		case bt.closedOrCrashed <- true:
		default:
		}
	})
	bt.setEventSubscriptionMapping(map[string]string{
		"request":         "request",
		"response":        "response",
		"requestfinished": "requestFinished",
		"responsefailed":  "responseFailed",
		"filechooser":     "fileChooser",
	})

	return bt
}

func (p *pageImpl) onBinding(binding *bindingCallImpl) {
	function := p.bindings[binding.initializer["name"].(string)]
	if function == nil {
		return
	}
	go binding.Call(function)
}

func (p *pageImpl) onFrameAttached(frame *frameImpl) {
	frame.page = p
	p.frames = append(p.frames, frame)
	p.Emit("frameattached", frame)
}

func (p *pageImpl) onFrameDetached(frame *frameImpl) {
	frame.detached = true
	frames := make([]Frame, 0)
	for i := 0; i < len(p.frames); i++ {
		if p.frames[i] != frame {
			frames = append(frames, frame)
		}
	}
	if len(frames) != len(p.frames) {
		p.frames = frames
	}
	p.Emit("framedetached", frame)
}

func (p *pageImpl) onRoute(route *routeImpl) {
	go func() {
		p.Lock()
		defer p.Unlock()
		url := route.Request().URL()
		for _, handlerEntry := range p.routes {
			if handlerEntry.matcher.Matches(url) {
				handlerEntry.handler(route)
				return
			}
		}
		p.browserContext.onRoute(route)
	}()
}

func (p *pageImpl) onWorker(worker *workerImpl) {
	p.workers = append(p.workers, worker)
	worker.page = p
	p.Emit("worker", worker)
}

func (p *pageImpl) onClose() {
	p.isClosed = true
	newPages := []Page{}
	newBackgoundPages := []Page{}
	if p.browserContext != nil {
		p.browserContext.Lock()
		for _, page := range p.browserContext.pages {
			if page != p {
				newPages = append(newPages, page)
			}
		}
		for _, page := range p.browserContext.backgroundPages {
			if page != p {
				newBackgoundPages = append(newBackgoundPages, page)
			}
		}
		p.browserContext.pages = newPages
		p.browserContext.backgroundPages = newBackgoundPages
		p.browserContext.Unlock()
	}
	p.Emit("close")
}

func (p *pageImpl) SetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) error {
	return p.mainFrame.SetInputFiles(selector, files, options...)
}

func (p *pageImpl) Check(selector string, options ...FrameCheckOptions) error {
	return p.mainFrame.Check(selector, options...)
}

func (p *pageImpl) Uncheck(selector string, options ...FrameUncheckOptions) error {
	return p.mainFrame.Uncheck(selector, options...)
}

func (p *pageImpl) WaitForTimeout(timeout float64) {
	p.mainFrame.WaitForTimeout(timeout)
}

func (p *pageImpl) WaitForFunction(expression string, arg interface{}, options ...FrameWaitForFunctionOptions) (JSHandle, error) {
	return p.mainFrame.WaitForFunction(expression, arg, options...)
}

func (p *pageImpl) Dblclick(expression string, options ...FrameDblclickOptions) error {
	return p.mainFrame.Dblclick(expression, options...)
}

func (p *pageImpl) Focus(expression string, options ...FrameFocusOptions) error {
	return p.mainFrame.Focus(expression, options...)
}

func (p *pageImpl) TextContent(selector string, options ...FrameTextContentOptions) (string, error) {
	return p.mainFrame.TextContent(selector, options...)
}

func (p *pageImpl) Video() Video {
	p.Lock()
	defer p.Unlock()

	if p.video == nil {
		p.video = newVideo(p)
	}
	return p.video
}

func (p *pageImpl) Tap(selector string, options ...FrameTapOptions) error {
	return p.mainFrame.Tap(selector, options...)
}

func (p *pageImpl) ExposeFunction(name string, binding ExposedFunction) error {
	return p.ExposeBinding(name, func(source *BindingSource, args ...interface{}) interface{} {
		return binding(args...)
	})
}
func (p *pageImpl) ExposeBinding(name string, binding BindingCallFunction, handle ...bool) error {
	needsHandle := false
	if len(handle) == 1 {
		needsHandle = handle[0]
	}
	if _, ok := p.bindings[name]; ok {
		return fmt.Errorf("Function '%s' has been already registered", name)
	}
	if _, ok := p.browserContext.bindings[name]; ok {
		return fmt.Errorf("Function '%s' has been already registered in the browser context", name)
	}
	p.bindings[name] = binding
	_, err := p.channel.Send("exposeBinding", map[string]interface{}{
		"name":        name,
		"needsHandle": needsHandle,
	})
	return err
}

func (p *pageImpl) SelectOption(selector string, values SelectOptionValues, options ...FrameSelectOptionOptions) ([]string, error) {
	return p.mainFrame.SelectOption(selector, values, options...)
}

func (p *pageImpl) IsChecked(selector string, options ...FrameIsCheckedOptions) (bool, error) {
	return p.mainFrame.IsChecked(selector, options...)
}

func (p *pageImpl) IsDisabled(selector string, options ...FrameIsDisabledOptions) (bool, error) {
	return p.mainFrame.IsDisabled(selector, options...)
}

func (p *pageImpl) IsEditable(selector string, options ...FrameIsEditableOptions) (bool, error) {
	return p.mainFrame.IsEditable(selector, options...)
}

func (p *pageImpl) IsEnabled(selector string, options ...FrameIsEnabledOptions) (bool, error) {
	return p.mainFrame.IsEnabled(selector, options...)
}

func (p *pageImpl) IsHidden(selector string, options ...FrameIsHiddenOptions) (bool, error) {
	return p.mainFrame.IsHidden(selector, options...)
}

func (p *pageImpl) IsVisible(selector string, options ...FrameIsVisibleOptions) (bool, error) {
	return p.mainFrame.IsVisible(selector, options...)
}

func (p *pageImpl) DragAndDrop(source, target string, options ...FrameDragAndDropOptions) error {
	return p.mainFrame.DragAndDrop(source, target, options...)
}

func (p *pageImpl) Pause() error {
	select {
	case <-p.closedOrCrashed:
		return fmt.Errorf("Page is closed or crashed")
	case err := <-p.browserContext.pause():
		return err
	}
}

func (p *pageImpl) InputValue(selector string, options ...FrameInputValueOptions) (string, error) {
	return p.mainFrame.InputValue(selector, options...)
}

func (p *pageImpl) WaitForURL(url string, options ...FrameWaitForURLOptions) error {
	return p.mainFrame.WaitForURL(url, options...)
}

func (p *pageImpl) SetChecked(selector string, checked bool, options ...FrameSetCheckedOptions) error {
	return p.mainFrame.SetChecked(selector, checked, options...)
}

func (p *pageImpl) Locator(selector string, options ...PageLocatorOptions) (Locator, error) {
	var option FrameLocatorOptions
	if len(options) == 1 {
		option = FrameLocatorOptions(options[0])
	}
	return p.mainFrame.Locator(selector, option)
}

func (p *pageImpl) FrameLocator(selector string) FrameLocator {
	return p.mainFrame.FrameLocator(selector)
}
