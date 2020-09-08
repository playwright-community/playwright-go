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
	Mouse           *Mouse
	Keyboard        *Keyboard
	timeoutSettings *timeoutSettings
	browserContext  *BrowserContext
	frames          []*Frame
	workersLock     sync.Mutex
	workers         []*Worker
	mainFrame       *Frame
	routesMu        sync.Mutex
	routes          []*routeHandlerEntry
	viewportSize    ViewportSize
	ownedContext    *BrowserContext
}

func (p *Page) Context() *BrowserContext {
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

func (p *Page) Opener() (*Page, error) {
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

func (p *Page) MainFrame() *Frame {
	return p.mainFrame
}

func (p *Page) Frames() []*Frame {
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

func (p *Page) QuerySelector(selector string) (*ElementHandle, error) {
	return p.mainFrame.QuerySelector(selector)
}

func (p *Page) QuerySelectorAll(selector string) ([]*ElementHandle, error) {
	return p.mainFrame.QuerySelectorAll(selector)
}

func (p *Page) WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (*ElementHandle, error) {
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

func (p *Page) AddScriptTag(options PageAddScriptTagOptions) (*ElementHandle, error) {
	return p.mainFrame.AddScriptTag(options)
}

func (p *Page) AddStyleTag(options PageAddStyleTagOptions) (*ElementHandle, error) {
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

func (p *Page) Goto(url string, options ...PageGotoOptions) (*Response, error) {
	return p.mainFrame.Goto(url, options...)
}

func (p *Page) Reload(options ...PageReloadOptions) (*Response, error) {
	response, err := p.channel.Send("reload", options)
	if err != nil {
		return nil, err
	}
	return fromChannel(response).(*Response), err
}

func (p *Page) WaitForLoadState(state ...string) {
	p.mainFrame.WaitForLoadState(state...)
}

func (p *Page) GoBack(options ...PageGoBackOptions) (*Response, error) {
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

func (p *Page) GoForward(options ...PageGoForwardOptions) (*Response, error) {
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
		"width":  width,
		"height": height,
	})
	if err != nil {
		return err
	}
	p.viewportSize.Height = height
	p.viewportSize.Width = width
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

func (p *Page) Press(selector, key string, options ...PagePressOptions) error {
	return p.mainFrame.Press(selector, key, options...)
}

func (p *Page) Title() (string, error) {
	return p.mainFrame.Title()
}

func (p *Page) Workers() []*Worker {
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

func (p *Page) WaitForNavigation(options ...PageWaitForNavigationOptions) (*Response, error) {
	return p.mainFrame.WaitForNavigation(options...)
}

func (p *Page) WaitForRequest(url interface{}, options ...interface{}) *Request {
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

func (p *Page) WaitForResponse(url interface{}, options ...interface{}) *Response {
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

func (p *Page) ExpectEvent(event string, cb func() error) (interface{}, error) {
	return newExpectWrapper(p.WaitForEvent, []interface{}{event}, cb)
}

func (p *Page) ExpectNavigation(cb func() error, options ...PageWaitForNavigationOptions) (*Response, error) {
	navigationOptions := make([]interface{}, 0)
	for _, option := range options {
		navigationOptions = append(navigationOptions, option)
	}
	response, err := newExpectWrapper(p.WaitForNavigation, navigationOptions, cb)
	return response.(*Response), err
}

func (p *Page) ExpectConsoleMessage(cb func() error) (*ConsoleMessage, error) {
	consoleMessage, err := newExpectWrapper(p.WaitForEvent, []interface{}{"console"}, cb)
	return consoleMessage.(*ConsoleMessage), err
}

func (p *Page) ExpectedDialog(cb func() error) (*Download, error) {
	dialog, err := newExpectWrapper(p.WaitForEvent, []interface{}{"download"}, cb)
	return dialog.(*Download), err
}

func (p *Page) ExpectDownload(cb func() error) (*Download, error) {
	download, err := newExpectWrapper(p.WaitForEvent, []interface{}{"download"}, cb)
	return download.(*Download), err
}

func (p *Page) ExpectFileChooser(cb func() error) (*FileChooser, error) {
	response, err := newExpectWrapper(p.WaitForEvent, []interface{}{"filechooser"}, cb)
	return response.(*FileChooser), err
}

func (p *Page) ExpectLoadState(state string, cb func() error) (*ConsoleMessage, error) {
	response, err := newExpectWrapper(p.mainFrame.WaitForLoadState, []interface{}{state}, cb)
	return response.(*ConsoleMessage), err
}

func (p *Page) ExpectPopup(cb func() error) (*Page, error) {
	popup, err := newExpectWrapper(p.WaitForEvent, []interface{}{"popup"}, cb)
	return popup.(*Page), err
}

func (p *Page) ExpectResponse(url interface{}, cb func() error, options ...interface{}) (*Response, error) {
	response, err := newExpectWrapper(p.WaitForResponse, append([]interface{}{url}, options...), cb)
	if err != nil {
		return nil, err
	}
	return response.(*Response), err
}

func (p *Page) ExpectRequest(url interface{}, cb func() error, options ...interface{}) (*Request, error) {
	popup, err := newExpectWrapper(p.WaitForRequest, append([]interface{}{url}, options...), cb)
	if err != nil {
		return nil, err
	}
	return popup.(*Request), err
}

func (p *Page) ExpectWorker(cb func() error) (*Worker, error) {
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

func (b *Page) AddInitScript(options BrowserContextAddInitScriptOptions) error {
	source := *options.Script
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

func newPage(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Page {
	bt := &Page{
		mainFrame: fromChannel(initializer["mainFrame"]).(*Frame),
		workers:   make([]*Worker, 0),
		routes:    make([]*routeHandlerEntry, 0),
		viewportSize: ViewportSize{
			Height: int(initializer["viewportSize"].(map[string]interface{})["height"].(float64)),
			Width:  int(initializer["viewportSize"].(map[string]interface{})["height"].(float64)),
		},
		timeoutSettings: newTimeoutSettings(nil),
	}
	bt.frames = []*Frame{bt.mainFrame}
	bt.mainFrame.page = bt
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.Mouse = newMouse(bt.channel)
	bt.Keyboard = newKeyboard(bt.channel)
	bt.channel.On("console", func(payload ...interface{}) {
		bt.Emit("console", fromChannel(payload[0].(map[string]interface{})["message"]))
	})
	bt.channel.On("crash", func(payload ...interface{}) {
		bt.Emit("crash")
	})
	bt.channel.On("dialog", func(payload ...interface{}) {
		bt.Emit("dialog", fromChannel(payload[0].(map[string]interface{})["dialog"]))
	})
	bt.channel.On("domcontentloaded", func(payload ...interface{}) {
		bt.Emit("domcontentloaded")
	})
	bt.channel.On("download", func(payload ...interface{}) {
		bt.Emit("download", fromChannel(payload[0].(map[string]interface{})["download"]))
	})
	bt.channel.On("fileChooser", func(payload ...interface{}) {
		bt.Emit("filechooser", newFileChooser(bt, fromChannel(payload[0].(map[string]interface{})["element"]).(*ElementHandle), payload[0].(map[string]interface{})["isMultiple"].(bool)))
	})
	bt.channel.On("frameAttached", func(payload ...interface{}) {
		frame := fromChannel(payload[0].(map[string]interface{})["frame"]).(*Frame)
		frame.page = bt
		bt.frames = append(bt.frames, frame)
		bt.Emit("frameAttached", frame)
	})
	bt.channel.On("frameDetached", func(payload ...interface{}) {
		frame := fromChannel(payload[0].(map[string]interface{})["frame"]).(*Frame)
		frame.detached = true
		frames := make([]*Frame, 0)
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
	bt.channel.On("popup", func(payload ...interface{}) {
		bt.Emit("popup", fromChannel(payload[0].(map[string]interface{})["page"]))
	})
	bt.channel.On("request", func(payload ...interface{}) {
		bt.Emit("request", fromChannel(payload[0].(map[string]interface{})["request"]))
	})
	bt.channel.On("requestFailed", func(payload ...interface{}) {
		req := fromChannel(payload[0].(map[string]interface{})["request"]).(*Request)
		req.failureText = payload[0].(map[string]interface{})["failureText"].(string)
		bt.Emit("requestFailed", req)
	})
	bt.channel.On("requestFinished", func(payload ...interface{}) {
		bt.Emit("requestFinished", fromChannel(payload[0].(map[string]interface{})["request"]))
	})
	bt.channel.On("response", func(payload ...interface{}) {
		bt.Emit("response", fromChannel(payload[0].(map[string]interface{})["response"]))
	})
	bt.channel.On("route", func(payload ...interface{}) {
		route := fromChannel(payload[0].(map[string]interface{})["route"]).(*Route)
		request := fromChannel(payload[0].(map[string]interface{})["request"]).(*Request)
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
	bt.channel.On("worker", func(payload ...interface{}) {
		worker := fromChannel(payload[0].(map[string]interface{})["worker"]).(*Worker)
		worker.page = bt
		bt.workers = append(bt.workers, worker)
		bt.Emit("worker", worker)
	})
	bt.addEventHandler(func(name string, handler eventHandler) {
		if name == "filechooser" && bt.ListenerCount(name) == 0 {
			bt.channel.SendNoReply("setFileChooserInterceptedNoReply", map[string]interface{}{
				"intercepted": true,
			})
		}
	})
	bt.removeEventHandler(func(name string, handler eventHandler) {
		if name == "filechooser" && bt.ListenerCount(name) == 0 {
			bt.channel.SendNoReply("setFileChooserInterceptedNoReply", map[string]interface{}{
				"intercepted": false,
			})
		}
	})

	return bt
}
