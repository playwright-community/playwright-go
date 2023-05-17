package playwright

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type browserContextImpl struct {
	channelOwner
	timeoutSettings   *timeoutSettings
	isClosedOrClosing bool
	options           map[string]interface{}
	pages             []Page
	routes            []*routeHandlerEntry
	ownedPage         Page
	browser           *browserImpl
	serviceWorkers    []*workerImpl
	backgroundPages   []Page
	bindings          map[string]BindingCallFunction
	tracing           *tracingImpl
	request           *apiRequestContextImpl
	harRecorders      map[string]harRecordingMetadata
}

func (b *browserContextImpl) SetDefaultNavigationTimeout(timeout float64) {
	b.timeoutSettings.SetNavigationTimeout(timeout)
	b.channel.SendNoReply("setDefaultNavigationTimeoutNoReply", map[string]interface{}{
		"timeout": timeout,
	})
}

func (b *browserContextImpl) SetDefaultTimeout(timeout float64) {
	b.timeoutSettings.SetTimeout(timeout)
	b.channel.SendNoReply("setDefaultTimeoutNoReply", map[string]interface{}{
		"timeout": timeout,
	})
}

func (b *browserContextImpl) Pages() []Page {
	b.Lock()
	defer b.Unlock()
	return b.pages
}

func (b *browserContextImpl) Browser() Browser {
	return b.browser
}
func (b *browserContextImpl) Tracing() Tracing {
	return b.tracing
}

func (b *browserContextImpl) NewCDPSession(page Page) (CDPSession, error) {
	channel, err := b.channel.Send("newCDPSession", map[string]interface{}{
		"page": page.(*pageImpl).channel,
	})
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}

	cdpSession := fromChannel(channel).(*cdpSessionImpl)

	return cdpSession, nil
}

func (b *browserContextImpl) NewPage(options ...BrowserNewPageOptions) (Page, error) {
	if b.ownedPage != nil {
		return nil, errors.New("Please use browser.NewContext()")
	}
	channel, err := b.channel.Send("newPage", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	return fromChannel(channel).(*pageImpl), nil
}

func (b *browserContextImpl) Cookies(urls ...string) ([]*Cookie, error) {
	result, err := b.channel.Send("cookies", map[string]interface{}{
		"urls": urls,
	})
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	cookies := make([]*Cookie, len(result.([]interface{})))
	for i, cookie := range result.([]interface{}) {
		cookies[i] = &Cookie{}
		remapMapToStruct(cookie, cookies[i])
	}
	return cookies, nil
}

func (b *browserContextImpl) AddCookies(cookies ...OptionalCookie) error {
	_, err := b.channel.Send("addCookies", map[string]interface{}{
		"cookies": cookies,
	})
	return err
}

func (b *browserContextImpl) ClearCookies() error {
	_, err := b.channel.Send("clearCookies")
	return err
}

func (b *browserContextImpl) GrantPermissions(permissions []string, options ...BrowserContextGrantPermissionsOptions) error {
	_, err := b.channel.Send("grantPermissions", map[string]interface{}{
		"permissions": permissions,
	}, options)
	return err
}

func (b *browserContextImpl) ClearPermissions() error {
	_, err := b.channel.Send("clearPermissions")
	return err
}

// Geolocation represents the options for BrowserContext.SetGeolocation()
type Geolocation struct {
	Longitude int  `json:"longitude"`
	Latitude  int  `json:"latitude"`
	Accuracy  *int `json:"accuracy"`
}

func (b *browserContextImpl) SetGeolocation(gelocation *Geolocation) error {
	_, err := b.channel.Send("setGeolocation", map[string]interface{}{
		"geolocation": gelocation,
	})
	return err
}

func (b *browserContextImpl) ResetGeolocation() error {
	_, err := b.channel.Send("setGeolocation", map[string]interface{}{})
	return err
}

func (b *browserContextImpl) SetExtraHTTPHeaders(headers map[string]string) error {
	_, err := b.channel.Send("setExtraHTTPHeaders", map[string]interface{}{
		"headers": serializeMapToNameAndValue(headers),
	})
	return err
}

func (b *browserContextImpl) SetOffline(offline bool) error {
	_, err := b.channel.Send("setOffline", map[string]interface{}{
		"offline": offline,
	})
	return err
}

func (b *browserContextImpl) AddInitScript(options BrowserContextAddInitScriptOptions) error {
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
	_, err := b.channel.Send("addInitScript", map[string]interface{}{
		"source": source,
	})
	return err
}

func (b *browserContextImpl) ExposeBinding(name string, binding BindingCallFunction, handle ...bool) error {
	needsHandle := false
	if len(handle) == 1 {
		needsHandle = handle[0]
	}
	for _, page := range b.pages {
		if _, ok := page.(*pageImpl).bindings[name]; ok {
			return fmt.Errorf("Function '%s' has been already registered in one of the pages", name)
		}
	}
	if _, ok := b.bindings[name]; ok {
		return fmt.Errorf("Function '%s' has been already registered", name)
	}
	b.bindings[name] = binding
	_, err := b.channel.Send("exposeBinding", map[string]interface{}{
		"name":        name,
		"needsHandle": needsHandle,
	})
	return err
}

func (b *browserContextImpl) ExposeFunction(name string, binding ExposedFunction) error {
	return b.ExposeBinding(name, func(source *BindingSource, args ...interface{}) interface{} {
		return binding(args...)
	})
}

func (b *browserContextImpl) Route(url interface{}, handler routeHandler, times ...int) error {
	b.routes = append(b.routes, newRouteHandlerEntry(newURLMatcher(url), handler, times...))
	return b.updateInterceptionPatterns()
}

func (b *browserContextImpl) Unroute(url interface{}, handlers ...routeHandler) error {
	b.Lock()
	defer b.Unlock()

	routes, err := unroute(b.routes, url, handlers...)
	if err != nil {
		return err
	}
	b.routes = routes
	return b.updateInterceptionPatterns()
}

func (b *browserContextImpl) Request() APIRequestContext {
	return b.request
}

func (b *browserContextImpl) RouteFromHAR(har string, options ...BrowserContextRouteFromHAROptions) error {
	opt := BrowserContextRouteFromHAROptions{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Update != nil && *opt.Update {
		var updateContent *HarContentPolicy
		switch opt.UpdateContent {
		case RouteFromHarUpdateContentPolicyAttach:
			updateContent = HarContentPolicyAttach
		case RouteFromHarUpdateContentPolicyEmbed:
			updateContent = HarContentPolicyEmbed
		}
		return b.recordIntoHar(har, browserContextRecordIntoHarOptions{
			URL:           opt.URL,
			UpdateContent: updateContent,
			UpdateMode:    opt.UpdateMode,
		})
	}
	notFound := opt.NotFound
	if notFound == nil {
		notFound = HarNotFoundAbort
	}
	router := newHarRouter(b.connection.localUtils, har, *notFound, opt.URL)
	return router.addContextRoute(b)
}

func (b *browserContextImpl) WaitForEvent(event string, predicate ...interface{}) interface{} {
	return <-waitForEvent(b, event, predicate...)
}

func (b *browserContextImpl) ExpectEvent(event string, cb func() error) (interface{}, error) {
	return newExpectWrapper(b.WaitForEvent, []interface{}{event}, cb)
}

func (b *browserContextImpl) Close() error {
	if b.isClosedOrClosing {
		return nil
	}
	b.Lock()
	b.isClosedOrClosing = true
	b.Unlock()

	for harId, harMetaData := range b.harRecorders {
		response, err := b.channel.Send("harExport", map[string]interface{}{
			"harId": harId,
		})
		if err != nil {
			return err
		}
		artifact := fromChannel(response).(*artifactImpl)
		// Server side will compress artifact if content is attach or if file is .zip.
		needCompressed := strings.HasSuffix(strings.ToLower(harMetaData.Path), ".zip")
		if !needCompressed && harMetaData.Content == HarContentPolicyAttach {
			tmpPath := harMetaData.Path + ".tmp"
			if err := artifact.SaveAs(tmpPath); err != nil {
				return err
			}
			_, err = b.connection.localUtils.HarUnzip(tmpPath, harMetaData.Path)
			if err != nil {
				return err
			}
		} else {
			if err := artifact.SaveAs(harMetaData.Path); err != nil {
				return err
			}
		}
		if err := artifact.Delete(); err != nil {
			return err
		}
	}
	_, err := b.channel.Send("close")
	return err
}

type browserContextRecordIntoHarOptions struct {
	Page          Page
	URL           interface{}
	UpdateContent *HarContentPolicy
	UpdateMode    *HarMode
}

func (b *browserContextImpl) recordIntoHar(har string, options ...browserContextRecordIntoHarOptions) error {
	overrides := map[string]interface{}{}
	harOptions := recordHarInputOptions{
		Path: har,
	}
	content := HarContentPolicyAttach
	if len(options) == 1 {
		if options[0].UpdateContent != nil {
			content = options[0].UpdateContent
		}
		harOptions.Content = content
		if options[0].UpdateMode != nil {
			harOptions.Mode = options[0].UpdateMode
		} else {
			harOptions.Mode = HarModeMinimal
		}
		harOptions.URL = options[0].URL
		overrides["options"] = prepareRecordHarOptions(harOptions)
		if options[0].Page != nil {
			overrides["page"] = options[0].Page.(*pageImpl).channel
		}
	}
	harId, err := b.channel.Send("harStart", overrides)
	if err != nil {
		return err
	}
	b.harRecorders[harId.(string)] = harRecordingMetadata{
		Path:    har,
		Content: content,
	}
	return nil
}

type StorageState struct {
	Cookies []Cookie       `json:"cookies"`
	Origins []OriginsState `json:"origins"`
}

type Cookie struct {
	Name     string             `json:"name"`
	Value    string             `json:"value"`
	URL      string             `json:"url"`
	Domain   string             `json:"domain"`
	Path     string             `json:"path"`
	Expires  float64            `json:"expires"`
	HttpOnly bool               `json:"httpOnly"`
	Secure   bool               `json:"secure"`
	SameSite *SameSiteAttribute `json:"sameSite"`
}
type OriginsState struct {
	Origin       string              `json:"origin"`
	LocalStorage []LocalStorageEntry `json:"localStorage"`
}

type LocalStorageEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (b *browserContextImpl) StorageState(paths ...string) (*StorageState, error) {
	result, err := b.channel.SendReturnAsDict("storageState")
	if err != nil {
		return nil, err
	}
	if len(paths) == 1 {
		file, err := os.Create(paths[0])
		if err != nil {
			return nil, err
		}
		if err := json.NewEncoder(file).Encode(result); err != nil {
			return nil, err
		}
		if err := file.Close(); err != nil {
			return nil, err
		}
	}
	var storageState StorageState
	remapMapToStruct(result, &storageState)
	return &storageState, nil
}

func (b *browserContextImpl) onBinding(binding *bindingCallImpl) {
	function := b.bindings[binding.initializer["name"].(string)]
	if function == nil {
		return
	}
	go binding.Call(function)
}

func (b *browserContextImpl) onClose() {
	b.Lock()
	b.isClosedOrClosing = true
	b.Unlock()
	if b.browser != nil {
		contexts := make([]BrowserContext, 0)
		b.browser.Lock()
		for _, context := range b.browser.contexts {
			if context != b {
				contexts = append(contexts, context)
			}
		}
		b.browser.contexts = contexts
		b.browser.Unlock()
	}
	b.Emit("close")
}

func (b *browserContextImpl) onPage(page *pageImpl) {
	page.setBrowserContext(b)
	b.Lock()
	b.pages = append(b.pages, page)
	b.Unlock()
	b.Emit("page", page)
	opener, _ := page.Opener()
	if opener != nil && !opener.IsClosed() {
		opener.Emit("popup", page)
	}
}

func (b *browserContextImpl) onRoute(route *routeImpl) {
	go func() {
		b.Lock()
		defer b.Unlock()

		routes := make([]*routeHandlerEntry, len(b.routes))
		copy(routes, b.routes)
		url := route.Request().URL()
		for i, handlerEntry := range routes {
			if handlerEntry.Matches(url) {
				if handlerEntry.WillExceed() {
					b.routes = append(b.routes[:i], b.routes[i+1:]...)
				}
				handled := handlerEntry.Handle(route)
				if len(b.routes) == 0 {
					err := b.updateInterceptionPatterns()
					if err != nil {
						log.Printf("could not update interception patterns: %v", err)
					}
				}
				yes := <-handled
				if yes {
					return
				}
			}
		}
		if err := route.internalContinue(true); err != nil {
			log.Printf("could not continue request: %v", err)
		}
	}()
}

func (b *browserContextImpl) updateInterceptionPatterns() error {
	patterns := prepareInterceptionPatterns(b.routes)
	_, err := b.channel.Send("setNetworkInterceptionPatterns", map[string]interface{}{
		"patterns": patterns,
	})
	return err
}

func (p *browserContextImpl) Pause() error {
	_, err := p.channel.Send("pause")
	return err
}

func (b *browserContextImpl) OnBackgroundPage(ev map[string]interface{}) {
	b.Lock()
	p := fromChannel(ev["page"]).(*pageImpl)
	p.browserContext = b
	b.backgroundPages = append(b.backgroundPages, p)
	b.Unlock()
	b.Emit("backgroundPage", p)
}

func (b *browserContextImpl) BackgroundPages() []Page {
	b.Lock()
	defer b.Unlock()
	return b.backgroundPages
}

func (b *browserContextImpl) setOptions(options, browserOptions map[string]interface{}) {
	b.options = options
	recordHar, ok := b.options["recordHar"]
	if ok {
		b.harRecorders[""] = harRecordingMetadata{
			Path:    recordHar.(recordHarOptions).Path,
			Content: recordHar.(recordHarOptions).Content,
		}
	}
	tracesDir, ok := browserOptions["tracesDir"].(string)
	if ok {
		b.tracing.tracesDir = tracesDir
	}
}

func newBrowserContext(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *browserContextImpl {
	bt := &browserContextImpl{
		timeoutSettings: newTimeoutSettings(nil),
		pages:           make([]Page, 0),
		routes:          make([]*routeHandlerEntry, 0),
		bindings:        make(map[string]BindingCallFunction),
		harRecorders:    make(map[string]harRecordingMetadata),
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	if parent.objectType == "Browser" {
		var browser interface{} = parent
		b, ok := browser.(*browserImpl)
		if ok {
			bt.browser = b
			bt.browser.contexts = append(bt.browser.contexts, bt)
		}
	}
	bt.tracing = fromChannel(initializer["tracing"]).(*tracingImpl)
	bt.request = fromChannel(initializer["requestContext"]).(*apiRequestContextImpl)
	bt.channel.On("bindingCall", func(params map[string]interface{}) {
		bt.onBinding(fromChannel(params["binding"]).(*bindingCallImpl))
	})
	bt.channel.On("request", func(ev map[string]interface{}) {
		request := fromChannel(ev["request"]).(*requestImpl)
		page := fromNullableChannel(ev["page"])
		bt.Emit("request", request)
		if page != nil {
			page.(*pageImpl).Emit("request", request)
		}
	})
	bt.channel.On("requestFailed", func(ev map[string]interface{}) {
		request := fromChannel(ev["request"]).(*requestImpl)
		failureText := ev["failureText"]
		if failureText != nil {
			request.failureText = failureText.(string)
		}
		page := fromNullableChannel(ev["page"])
		if request.timing != nil {
			request.timing.ResponseEnd = ev["responseEndTiming"].(float64)
		}
		bt.Emit("requestfailed", request)
		if page != nil {
			page.(*pageImpl).Emit("requestfailed", request)
		}
	})

	bt.channel.On("requestFinished", func(ev map[string]interface{}) {
		request := fromChannel(ev["request"]).(*requestImpl)
		response := fromNullableChannel(ev["response"])
		if response != nil {
			response.(*responseImpl).finished <- true
		}
		page := fromNullableChannel(ev["page"])
		if request.timing != nil {
			request.timing.ResponseEnd = ev["responseEndTiming"].(float64)
		}
		bt.Emit("requestfinished", request)
		if page != nil {
			page.(*pageImpl).Emit("requestfinished", request)
		}
	})
	bt.channel.On("response", func(ev map[string]interface{}) {
		response := fromChannel(ev["response"]).(*responseImpl)
		page := fromNullableChannel(ev["page"])
		bt.Emit("response", response)
		if page != nil {
			page.(*pageImpl).Emit("response", response)
		}
	})
	bt.channel.On("close", bt.onClose)
	bt.channel.On("page", func(payload map[string]interface{}) {
		bt.onPage(fromChannel(payload["page"]).(*pageImpl))
	})
	bt.channel.On("route", func(params map[string]interface{}) {
		bt.onRoute(fromChannel(params["route"]).(*routeImpl))
	})
	bt.channel.On("backgroundPage", bt.OnBackgroundPage)
	bt.setEventSubscriptionMapping(map[string]string{
		"request":         "request",
		"response":        "response",
		"requestfinished": "requestFinished",
		"responsefailed":  "responseFailed",
	})
	return bt
}
