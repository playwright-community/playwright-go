package playwright

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"reflect"
	"sync"
)

type pageImpl struct {
	channelOwner
	isClosed        bool
	video           *videoImpl
	mouse           *mouseImpl
	keyboard        *keyboardImpl
	touchscreen     *touchscreenImpl
	timeoutSettings *timeoutSettings
	browserContext  *browserContextImpl
	frames          []Frame
	workers         []Worker
	mainFrame       Frame
	routesMu        sync.Mutex
	routes          []*routeHandlerEntry
	viewportSize    ViewportSize
	ownedContext    BrowserContext
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
	channel, err := p.channel.Send("opener")
	if err != nil {
		return nil, err
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*pageImpl), nil
}

func (p *pageImpl) MainFrame() Frame {
	return p.mainFrame
}

type PageFrameOptions struct {
	Name *string
	URL  interface{}
}

func (p *pageImpl) Frame(options PageFrameOptions) Frame {
	var matcher *urlMatcher
	if options.URL != nil {
		matcher = newURLMatcher(options.URL)
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

func (p *pageImpl) SetDefaultNavigationTimeout(timeout int) {
	p.timeoutSettings.SetNavigationTimeout(timeout)
	p.channel.SendNoReply("setDefaultNavigationTimeoutNoReply", map[string]interface{}{
		"timeout": timeout,
	})
}

func (p *pageImpl) SetDefaultTimeout(timeout int) {
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
	return p.mainFrame.DispatchEvent(selector, typ, options...)
}

func (p *pageImpl) Evaluate(expression string, options ...interface{}) (interface{}, error) {
	return p.mainFrame.Evaluate(expression, options...)
}

func (p *pageImpl) EvaluateHandle(expression string, options ...interface{}) (interface{}, error) {
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
		"headers": serializeHeaders(headers),
	})
	return err
}

func (p *pageImpl) URL() string {
	return p.mainFrame.URL()
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
	response, err := p.channel.Send("reload", options)
	if err != nil {
		return nil, err
	}
	return fromChannel(response).(*responseImpl), err
}

func (p *pageImpl) WaitForLoadState(state ...string) {
	p.mainFrame.WaitForLoadState(state...)
}

func (p *pageImpl) GoBack(options ...PageGoBackOptions) (Response, error) {
	channel, err := p.channel.Send("goBack", options)
	if err != nil {
		return nil, err
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*responseImpl), nil
}

func (p *pageImpl) GoForward(options ...PageGoForwardOptions) (Response, error) {
	resp, err := p.channel.Send("goForward", options)
	if err != nil {
		return nil, err
	}
	obj := fromNullableChannel(resp)
	if obj == nil {
		return nil, nil
	}
	return obj.(*responseImpl), nil
}

func (p *pageImpl) EmulateMedia(options ...PageEmulateMediaOptions) error {
	_, err := p.channel.Send("emulateMedia", options)
	if err != nil {
		return err
	}
	return err
}

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
		if err := ioutil.WriteFile(*path, image, 0644); err != nil {
			return nil, err
		}
	}
	return image, nil
}

func (p *pageImpl) PDF(options ...PagePDFOptions) ([]byte, error) {
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
		if err := ioutil.WriteFile(*path, pdf, 0644); err != nil {
			return nil, err
		}
	}
	return pdf, nil
}

func (p *pageImpl) Click(selector string, options ...PageClickOptions) error {
	return p.mainFrame.Click(selector, options...)
}

func (p *pageImpl) WaitForEvent(event string, predicate ...interface{}) interface{} {
	return <-waitForEvent(p, event, predicate...)
}

func (p *pageImpl) WaitForNavigation(options ...PageWaitForNavigationOptions) (Response, error) {
	return p.mainFrame.WaitForNavigation(options...)
}

func (p *pageImpl) WaitForRequest(url interface{}, options ...interface{}) Request {
	var matcher *urlMatcher
	if url != nil {
		matcher = newURLMatcher(url)
	}
	predicate := func(req *requestImpl) bool {
		if matcher != nil {
			return matcher.Matches(req.URL())
		}
		if len(options) == 1 {
			return reflect.ValueOf(options[0]).Call([]reflect.Value{reflect.ValueOf(req)})[0].Bool()
		}
		return true
	}
	return p.WaitForEvent("request", predicate).(*requestImpl)
}

func (p *pageImpl) WaitForResponse(url interface{}, options ...interface{}) Response {
	var matcher *urlMatcher
	if url != nil {
		matcher = newURLMatcher(url)
	}
	predicate := func(req *responseImpl) bool {
		if matcher != nil {
			return matcher.Matches(req.URL())
		}
		if len(options) == 1 {
			return reflect.ValueOf(options[0]).Call([]reflect.Value{reflect.ValueOf(req)})[0].Bool()
		}
		return true
	}
	return p.WaitForEvent("response", predicate).(*responseImpl)
}

func (p *pageImpl) ExpectEvent(event string, cb func() error, predicates ...interface{}) (interface{}, error) {
	args := []interface{}{event}
	if len(predicates) == 1 {
		args = append(args, predicates[0])
	}
	return newExpectWrapper(p.WaitForEvent, args, cb)
}

func (p *pageImpl) ExpectNavigation(cb func() error, options ...PageWaitForNavigationOptions) (Response, error) {
	navigationOptions := make([]interface{}, 0)
	for _, option := range options {
		navigationOptions = append(navigationOptions, option)
	}
	response, err := newExpectWrapper(p.WaitForNavigation, navigationOptions, cb)
	if response == nil {
		return nil, err
	}
	return response.(*responseImpl), err
}

func (p *pageImpl) ExpectConsoleMessage(cb func() error) (ConsoleMessage, error) {
	consoleMessage, err := newExpectWrapper(p.WaitForEvent, []interface{}{"console"}, cb)
	return consoleMessage.(*consoleMessageImpl), err
}

func (p *pageImpl) ExpectedDialog(cb func() error) (Dialog, error) {
	dialog, err := newExpectWrapper(p.WaitForEvent, []interface{}{"dialog"}, cb)
	return dialog.(*dialogImpl), err
}

func (p *pageImpl) ExpectDownload(cb func() error) (Download, error) {
	download, err := newExpectWrapper(p.WaitForEvent, []interface{}{"download"}, cb)
	return download.(*downloadImpl), err
}

func (p *pageImpl) ExpectFileChooser(cb func() error) (FileChooser, error) {
	response, err := newExpectWrapper(p.WaitForEvent, []interface{}{"filechooser"}, cb)
	return response.(*fileChooserImpl), err
}

func (p *pageImpl) ExpectLoadState(state string, cb func() error) (ConsoleMessage, error) {
	response, err := newExpectWrapper(p.mainFrame.WaitForLoadState, []interface{}{state}, cb)
	return response.(*consoleMessageImpl), err
}

func (p *pageImpl) ExpectPopup(cb func() error) (Page, error) {
	popup, err := newExpectWrapper(p.WaitForEvent, []interface{}{"popup"}, cb)
	return popup.(*pageImpl), err
}

func (p *pageImpl) ExpectResponse(url interface{}, cb func() error, options ...interface{}) (Response, error) {
	response, err := newExpectWrapper(p.WaitForResponse, append([]interface{}{url}, options...), cb)
	if err != nil {
		return nil, err
	}
	return response.(*responseImpl), err
}

func (p *pageImpl) ExpectRequest(url interface{}, cb func() error, options ...interface{}) (Request, error) {
	popup, err := newExpectWrapper(p.WaitForRequest, append([]interface{}{url}, options...), cb)
	if err != nil {
		return nil, err
	}
	return popup.(*requestImpl), err
}

func (p *pageImpl) ExpectWorker(cb func() error) (Worker, error) {
	response, err := newExpectWrapper(p.WaitForEvent, []interface{}{"worker"}, cb)
	return response.(*workerImpl), err
}

func (p *pageImpl) Route(url interface{}, handler routeHandler) error {
	p.routesMu.Lock()
	defer p.routesMu.Unlock()
	p.routes = append(p.routes, newRouteHandlerEntry(newURLMatcher(url), handler))
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

func (p *pageImpl) AddInitScript(options BrowserContextAddInitScriptOptions) error {
	var source string
	if options.Script != nil {
		source = *options.Script
	}
	if options.Path != nil {
		content, err := ioutil.ReadFile(*options.Path)
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
	bt := &pageImpl{
		mainFrame: fromChannel(initializer["mainFrame"]).(*frameImpl),
		workers:   make([]Worker, 0),
		routes:    make([]*routeHandlerEntry, 0),
		viewportSize: ViewportSize{
			Height: int(initializer["viewportSize"].(map[string]interface{})["height"].(float64)),
			Width:  int(initializer["viewportSize"].(map[string]interface{})["width"].(float64)),
		},
		timeoutSettings: newTimeoutSettings(nil),
	}
	bt.frames = []Frame{bt.mainFrame}
	bt.mainFrame.(*frameImpl).page = bt
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.mouse = newMouse(bt.channel)
	bt.keyboard = newKeyboard(bt.channel)
	bt.touchscreen = newTouchscreen(bt.channel)
	bt.channel.On("close", func(ev map[string]interface{}) {
		bt.isClosed = true
		bt.Emit("close")
	})
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
	bt.channel.On("download", func(ev map[string]interface{}) {
		bt.Emit("download", fromChannel(ev["download"]))
	})
	bt.channel.On("fileChooser", func(ev map[string]interface{}) {
		bt.Emit("filechooser", newFileChooser(bt, fromChannel(ev["element"]).(*elementHandleImpl), ev["isMultiple"].(bool)))
	})
	bt.channel.On("frameAttached", func(ev map[string]interface{}) {
		frame := fromChannel(ev["frame"]).(*frameImpl)
		frame.page = bt
		bt.frames = append(bt.frames, frame)
		bt.Emit("frameAttached", frame)
	})
	bt.channel.On("frameDetached", func(ev map[string]interface{}) {
		frame := fromChannel(ev["frame"]).(*frameImpl)
		frame.detached = true
		frames := make([]Frame, 0)
		for i := 0; i < len(bt.frames); i++ {
			if bt.frames[i] != frame {
				frames = append(frames, frame)
			}
		}
		if len(frames) != len(bt.frames) {
			bt.frames = frames
		}
		bt.Emit("frameDetached", frame)
	})
	bt.channel.On(
		"pageError",
		func(params map[string]interface{}) {
			err := errorPayload{}
			remapMapToStruct(params["error"].(map[string]interface{})["error"], &err)
			bt.Emit("pageerror", parseError(err))
		},
	)
	bt.channel.On("popup", func(ev map[string]interface{}) {
		bt.Emit("popup", fromChannel(ev["page"]))
	})
	bt.channel.On("request", func(ev map[string]interface{}) {
		bt.Emit("request", fromChannel(ev["request"]))
	})
	bt.channel.On("requestFailed", func(ev map[string]interface{}) {
		req := fromChannel(ev["request"]).(*requestImpl)
		req.failureText = ev["failureText"].(string)
		bt.Emit("requestfailed", req)
	})
	bt.channel.On("requestFinished", func(ev map[string]interface{}) {
		bt.Emit("requestfinished", fromChannel(ev["request"]))
	})
	bt.channel.On("response", func(ev map[string]interface{}) {
		bt.Emit("response", fromChannel(ev["response"]))
	})
	bt.channel.On("route", func(ev map[string]interface{}) {
		route := fromChannel(ev["route"]).(*routeImpl)
		request := fromChannel(ev["request"]).(*requestImpl)
		go func() {
			bt.routesMu.Lock()
			for _, handlerEntry := range bt.routes {
				if handlerEntry.matcher.Matches(request.URL()) {
					handlerEntry.handler(route, request)
					break
				}
			}
			bt.routesMu.Unlock()
		}()
	})
	bt.channel.On("video", func(params map[string]interface{}) {
		bt.Video().(*videoImpl).setRelativePath(params["relativePath"].(string))
	})
	bt.channel.On("webSocket", func(ev map[string]interface{}) {
		bt.Emit("websocket", fromChannel(ev["webSocket"]).(*webSocketImpl))
	})

	bt.channel.On("worker", func(ev map[string]interface{}) {
		worker := fromChannel(ev["worker"]).(*workerImpl)
		worker.page = bt
		bt.workers = append(bt.workers, worker)
		bt.Emit("worker", worker)
	})
	bt.addEventHandler(func(name string, handler interface{}) {
		if name == "filechooser" && bt.ListenerCount(name) == 0 {
			bt.channel.SendNoReply("setFileChooserInterceptedNoReply", map[string]interface{}{
				"intercepted": true,
			})
		}
	})
	bt.removeEventHandler(func(name string, handler interface{}) {
		if name == "filechooser" && bt.ListenerCount(name) == 0 {
			bt.channel.SendNoReply("setFileChooserInterceptedNoReply", map[string]interface{}{
				"intercepted": false,
			})
		}
	})

	return bt
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

func (p *pageImpl) WaitForTimeout(timeout int) {
	p.mainFrame.WaitForTimeout(timeout)
}

func (p *pageImpl) WaitForFunction(expression string, options ...FrameWaitForFunctionOptions) (JSHandle, error) {
	return p.mainFrame.WaitForFunction(expression, options...)
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
	contextOptions := p.browserContext.options
	if contextOptions.RecordVideo == nil {
		return nil
	}
	if p.video == nil {
		p.video = newVideo(p)
		if videoRelativePath, ok := p.initializer["videoRelativePath"]; ok {
			p.video.setRelativePath(videoRelativePath.(string))
		}
	}
	return p.video
}

func (p *pageImpl) Tap(selector string, options ...FrameTapOptions) error {
	return p.mainFrame.Tap(selector, options...)
}
