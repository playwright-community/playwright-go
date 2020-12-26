package playwright

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"reflect"
	"sync"
)

type Page struct {
	ChannelOwner
	isClosed        bool
	mouse           *Mouse
	keyboard        *Keyboard
	timeoutSettings *timeoutSettings
	browserContext  BrowserContextI
	frames          []FrameI
	workersLock     sync.Mutex
	workers         []WorkerI
	mainFrame       FrameI
	routesMu        sync.Mutex
	routes          []*routeHandlerEntry
	viewportSize    ViewportSize
	ownedContext    BrowserContextI
}

func (p *Page) Context() BrowserContextI {
	return p.browserContext
}

func (p *Page) Close(options ...PageCloseOptions) error {
	_, err := p.channel.Send("close", options)
	if err != nil {
		return err
	}
	if p.ownedContext != nil {
		return p.ownedContext.Close()
	}
	return nil
}

func (p *Page) InnerText(selector string, options ...PageInnerTextOptions) (string, error) {
	return p.mainFrame.InnerText(selector, options...)
}

func (p *Page) InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error) {
	return p.mainFrame.InnerHTML(selector, options...)
}

func (p *Page) Opener() (PageI, error) {
	channel, err := p.channel.Send("opener")
	if err != nil {
		return nil, err
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*Page), nil
}

func (p *Page) MainFrame() FrameI {
	return p.mainFrame
}

func (p *Page) Frames() []FrameI {
	return p.frames
}

func (p *Page) SetDefaultNavigationTimeout(timeout int) {
	p.timeoutSettings.SetNavigationTimeout(timeout)
	p.channel.SendNoReply("setDefaultNavigationTimeoutNoReply", map[string]interface{}{
		"timeout": timeout,
	})
}

func (p *Page) SetDefaultTimeout(timeout int) {
	p.timeoutSettings.SetTimeout(timeout)
	p.channel.SendNoReply("setDefaultTimeoutNoReply", map[string]interface{}{
		"timeout": timeout,
	})
}

func (p *Page) QuerySelector(selector string) (ElementHandleI, error) {
	return p.mainFrame.QuerySelector(selector)
}

func (p *Page) QuerySelectorAll(selector string) ([]ElementHandleI, error) {
	return p.mainFrame.QuerySelectorAll(selector)
}

func (p *Page) WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandleI, error) {
	return p.mainFrame.WaitForSelector(selector, options...)
}

func (p *Page) DispatchEvent(selector string, typ string, options ...PageDispatchEventOptions) error {
	return p.mainFrame.DispatchEvent(selector, typ, options...)
}

func (p *Page) Evaluate(expression string, options ...interface{}) (interface{}, error) {
	return p.mainFrame.Evaluate(expression, options...)
}

func (p *Page) EvaluateHandle(expression string, options ...interface{}) (interface{}, error) {
	return p.mainFrame.EvaluateHandle(expression, options...)
}

func (p *Page) EvaluateOnSelector(selector string, expression string, options ...interface{}) (interface{}, error) {
	return p.mainFrame.EvaluateOnSelector(selector, expression, options...)
}

func (p *Page) EvaluateOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error) {
	return p.mainFrame.EvaluateOnSelectorAll(selector, expression, options...)
}

func (p *Page) AddScriptTag(options PageAddScriptTagOptions) (ElementHandleI, error) {
	return p.mainFrame.AddScriptTag(options)
}

func (p *Page) AddStyleTag(options PageAddStyleTagOptions) (ElementHandleI, error) {
	return p.mainFrame.AddStyleTag(options)
}

func (p *Page) SetExtraHTTPHeaders(headers map[string]string) error {
	_, err := p.channel.Send("setExtraHTTPHeaders", map[string]interface{}{
		"headers": serializeHeaders(headers),
	})
	return err
}

func (p *Page) URL() string {
	return p.mainFrame.URL()
}

func (p *Page) Content() (string, error) {
	return p.mainFrame.Content()
}

func (p *Page) SetContent(content string, options ...PageSetContentOptions) error {
	return p.mainFrame.SetContent(content, options...)
}

func (p *Page) Goto(url string, options ...PageGotoOptions) (ResponseI, error) {
	return p.mainFrame.Goto(url, options...)
}

func (p *Page) Reload(options ...PageReloadOptions) (ResponseI, error) {
	response, err := p.channel.Send("reload", options)
	if err != nil {
		return nil, err
	}
	return fromChannel(response).(*Response), err
}

func (p *Page) WaitForLoadState(state ...string) {
	p.mainFrame.WaitForLoadState(state...)
}

func (p *Page) GoBack(options ...PageGoBackOptions) (ResponseI, error) {
	channel, err := p.channel.Send("goBack", options)
	if err != nil {
		return nil, err
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*Response), nil
}

func (p *Page) GoForward(options ...PageGoForwardOptions) (ResponseI, error) {
	resp, err := p.channel.Send("goForward", options)
	if err != nil {
		return nil, err
	}
	obj := fromNullableChannel(resp)
	if obj == nil {
		return nil, nil
	}
	return obj.(*Response), nil
}

func (p *Page) EmulateMedia(options ...PageEmulateMediaOptions) error {
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

func (p *Page) SetViewportSize(width, height int) error {
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

func (p *Page) ViewportSize() ViewportSize {
	return p.viewportSize
}

func (p *Page) BringToFront() error {
	_, err := p.channel.Send("bringToFront")
	return err
}

func (p *Page) Type(selector, text string, options ...PageTypeOptions) error {
	return p.mainFrame.Type(selector, text, options...)
}

func (p *Page) Fill(selector, text string, options ...FrameFillOptions) error {
	return p.mainFrame.Fill(selector, text, options...)
}

func (p *Page) Press(selector, key string, options ...PagePressOptions) error {
	return p.mainFrame.Press(selector, key, options...)
}

func (p *Page) Title() (string, error) {
	return p.mainFrame.Title()
}

func (p *Page) Workers() []WorkerI {
	p.workersLock.Lock()
	defer p.workersLock.Unlock()
	return p.workers
}

func (p *Page) Screenshot(options ...PageScreenshotOptions) ([]byte, error) {
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

func (p *Page) PDF(options ...PagePdfOptions) ([]byte, error) {
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

func (p *Page) Click(selector string, options ...PageClickOptions) error {
	return p.mainFrame.Click(selector, options...)
}

func (p *Page) WaitForEvent(event string, predicate ...interface{}) interface{} {
	evChan := make(chan interface{})
	handler := func(ev ...interface{}) {
		if len(predicate) == 0 {
			evChan <- ev[0]
		} else if len(predicate) == 1 {
			result := reflect.ValueOf(predicate[0]).Call([]reflect.Value{reflect.ValueOf(ev[0])})
			if result[0].Bool() {
				evChan <- ev[0]
			}
		}
	}
	p.On(event, handler)
	defer close(evChan)
	defer p.RemoveListener(event, handler)
	return <-evChan
}

func (p *Page) WaitForNavigation(options ...PageWaitForNavigationOptions) (ResponseI, error) {
	return p.mainFrame.WaitForNavigation(options...)
}

func (p *Page) WaitForRequest(url interface{}, options ...interface{}) RequestI {
	var matcher *urlMatcher
	if url != nil {
		matcher = newURLMatcher(url)
	}
	predicate := func(req *Request) bool {
		if matcher != nil {
			return matcher.Match(req.URL())
		}
		if len(options) == 1 {
			return reflect.ValueOf(options[0]).Call([]reflect.Value{reflect.ValueOf(req)})[0].Bool()
		}
		return true
	}
	return p.WaitForEvent("request", predicate).(*Request)
}

func (p *Page) WaitForResponse(url interface{}, options ...interface{}) ResponseI {
	var matcher *urlMatcher
	if url != nil {
		matcher = newURLMatcher(url)
	}
	predicate := func(req *Response) bool {
		if matcher != nil {
			return matcher.Match(req.URL())
		}
		if len(options) == 1 {
			return reflect.ValueOf(options[0]).Call([]reflect.Value{reflect.ValueOf(req)})[0].Bool()
		}
		return true
	}
	return p.WaitForEvent("response", predicate).(*Response)
}

func (p *Page) ExpectEvent(event string, cb func() error, predicates ...interface{}) (interface{}, error) {
	var predicate interface{}
	if len(predicates) == 1 {
		predicate = predicates[0]
	}
	return newExpectWrapper(p.WaitForEvent, []interface{}{event, predicate}, cb)
}

func (p *Page) ExpectNavigation(cb func() error, options ...PageWaitForNavigationOptions) (ResponseI, error) {
	navigationOptions := make([]interface{}, 0)
	for _, option := range options {
		navigationOptions = append(navigationOptions, option)
	}
	response, err := newExpectWrapper(p.WaitForNavigation, navigationOptions, cb)
	if response == nil {
		return nil, err
	}
	return response.(*Response), err
}

func (p *Page) ExpectConsoleMessage(cb func() error) (ConsoleMessageI, error) {
	consoleMessage, err := newExpectWrapper(p.WaitForEvent, []interface{}{"console"}, cb)
	return consoleMessage.(*ConsoleMessage), err
}

func (p *Page) ExpectedDialog(cb func() error) (DialogI, error) {
	dialog, err := newExpectWrapper(p.WaitForEvent, []interface{}{"dialog"}, cb)
	return dialog.(*Dialog), err
}

func (p *Page) ExpectDownload(cb func() error) (DownloadI, error) {
	download, err := newExpectWrapper(p.WaitForEvent, []interface{}{"download"}, cb)
	return download.(*Download), err
}

func (p *Page) ExpectFileChooser(cb func() error) (FileChooserI, error) {
	response, err := newExpectWrapper(p.WaitForEvent, []interface{}{"filechooser"}, cb)
	return response.(*FileChooser), err
}

func (p *Page) ExpectLoadState(state string, cb func() error) (ConsoleMessageI, error) {
	response, err := newExpectWrapper(p.mainFrame.WaitForLoadState, []interface{}{state}, cb)
	return response.(*ConsoleMessage), err
}

func (p *Page) ExpectPopup(cb func() error) (PageI, error) {
	popup, err := newExpectWrapper(p.WaitForEvent, []interface{}{"popup"}, cb)
	return popup.(*Page), err
}

func (p *Page) ExpectResponse(url interface{}, cb func() error, options ...interface{}) (ResponseI, error) {
	response, err := newExpectWrapper(p.WaitForResponse, append([]interface{}{url}, options...), cb)
	if err != nil {
		return nil, err
	}
	return response.(*Response), err
}

func (p *Page) ExpectRequest(url interface{}, cb func() error, options ...interface{}) (RequestI, error) {
	popup, err := newExpectWrapper(p.WaitForRequest, append([]interface{}{url}, options...), cb)
	if err != nil {
		return nil, err
	}
	return popup.(*Request), err
}

func (p *Page) ExpectWorker(cb func() error) (WorkerI, error) {
	response, err := newExpectWrapper(p.WaitForEvent, []interface{}{"worker"}, cb)
	return response.(*Worker), err
}

func (p *Page) Route(url interface{}, handler routeHandler) error {
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

func (p *Page) GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error) {
	return p.mainFrame.GetAttribute(selector, name, options...)
}

func (p *Page) Hover(selector string, options ...PageHoverOptions) error {
	return p.mainFrame.Hover(selector, options...)
}

func (p *Page) Isclosed() bool {
	return p.isClosed
}

func (b *Page) AddInitScript(options BrowserContextAddInitScriptOptions) error {
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
	_, err := b.channel.Send("addInitScript", map[string]interface{}{
		"source": source,
	})
	return err
}

func (p *Page) Keyboard() *Keyboard {
	return p.keyboard
}
func (p *Page) Mouse() *Mouse {
	return p.mouse
}

func newPage(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Page {
	bt := &Page{
		mainFrame: fromChannel(initializer["mainFrame"]).(*Frame),
		workers:   make([]WorkerI, 0),
		routes:    make([]*routeHandlerEntry, 0),
		viewportSize: ViewportSize{
			Height: int(initializer["viewportSize"].(map[string]interface{})["height"].(float64)),
			Width:  int(initializer["viewportSize"].(map[string]interface{})["width"].(float64)),
		},
		timeoutSettings: newTimeoutSettings(nil),
	}
	bt.frames = []FrameI{bt.mainFrame}
	bt.mainFrame.(*Frame).page = bt
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.mouse = newMouse(bt.channel)
	bt.keyboard = newKeyboard(bt.channel)
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
		bt.Emit("filechooser", newFileChooser(bt, fromChannel(ev["element"]).(*ElementHandle), ev["isMultiple"].(bool)))
	})
	bt.channel.On("frameAttached", func(ev map[string]interface{}) {
		frame := fromChannel(ev["frame"]).(*Frame)
		frame.page = bt
		bt.frames = append(bt.frames, frame)
		bt.Emit("frameAttached", frame)
	})
	bt.channel.On("frameDetached", func(ev map[string]interface{}) {
		frame := fromChannel(ev["frame"]).(*Frame)
		frame.detached = true
		frames := make([]FrameI, 0)
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
	bt.channel.On("popup", func(ev map[string]interface{}) {
		bt.Emit("popup", fromChannel(ev["page"]))
	})
	bt.channel.On("request", func(ev map[string]interface{}) {
		bt.Emit("request", fromChannel(ev["request"]))
	})
	bt.channel.On("requestFailed", func(ev map[string]interface{}) {
		req := fromChannel(ev["request"]).(*Request)
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
		route := fromChannel(ev["route"]).(*Route)
		request := fromChannel(ev["request"]).(*Request)
		go func() {
			bt.routesMu.Lock()
			for _, handlerEntry := range bt.routes {
				if handlerEntry.matcher.Match(request.URL()) {
					handlerEntry.handler(route, request)
					break
				}
			}
			bt.routesMu.Unlock()
		}()
	})
	bt.channel.On("worker", func(ev map[string]interface{}) {
		worker := fromChannel(ev["worker"]).(*Worker)
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

func (p *Page) SetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) error {
	return p.mainFrame.SetInputFiles(selector, files, options...)
}

func (p *Page) Check(selector string, options ...FrameCheckOptions) error {
	return p.mainFrame.Check(selector, options...)
}

func (p *Page) Uncheck(selector string, options ...FrameUncheckOptions) error {
	return p.mainFrame.Uncheck(selector, options...)
}

func (p *Page) WaitForTimeout(timeout int) {
	p.mainFrame.WaitForTimeout(timeout)
}

func (p *Page) WaitForFunction(expression string, options ...FrameWaitForFunctionOptions) (JSHandleI, error) {
	return p.mainFrame.WaitForFunction(expression, options...)
}

func (p *Page) DblClick(expression string, options ...FrameDblclickOptions) error {
	return p.mainFrame.DblClick(expression, options...)
}

func (p *Page) Focus(expression string, options ...FrameFocusOptions) error {
	return p.mainFrame.Focus(expression, options...)
}

func (p *Page) TextContent(selector string, options ...FrameTextContentOptions) (string, error) {
	return p.mainFrame.TextContent(selector, options...)
}
