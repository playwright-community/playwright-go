package playwright

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type browserContextImpl struct {
	channelOwner
	timeoutSettings   *timeoutSettings
	isClosedOrClosing bool
	options           *BrowserNewContextOptions
	pages             []Page
	routes            []*routeHandlerEntry
	ownedPage         Page
	browser           *browserImpl
	serviceWorkers    []*workerImpl
	bindings          map[string]BindingCallFunction
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

// NetworkCookie is the return structure of BrowserContext.Cookies()
type NetworkCookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Domain   string `json:"domain"`
	Path     string `json:"path"`
	Expires  int    `json:"expires"`
	HttpOnly bool   `json:"httpOnly"`
	Secure   bool   `json:"secure"`
	SameSite string `json:"sameSite"`
}

func (b *browserContextImpl) Cookies(urls ...string) ([]*NetworkCookie, error) {
	result, err := b.channel.Send("cookies", map[string]interface{}{
		"urls": urls,
	})
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	cookies := make([]*NetworkCookie, len(result.([]interface{})))
	for i, cookie := range result.([]interface{}) {
		cookies[i] = &NetworkCookie{}
		remapMapToStruct(cookie, cookies[i])
	}
	return cookies, nil
}

// SetNetworkCookieParam is used to filter cookies in BrowserContext.AddCookies()
type SetNetworkCookieParam struct {
	Name     string  `json:"name"`
	Value    string  `json:"value"`
	URL      *string `json:"url"`
	Domain   *string `json:"domain"`
	Path     *string `json:"path"`
	Expires  *int    `json:"expires"`
	HttpOnly *bool   `json:"httpOnly"`
	Secure   *bool   `json:"secure"`
	SameSite *string `json:"sameSite"`
}

func (b *browserContextImpl) AddCookies(cookies ...SetNetworkCookieParam) error {
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

// SetGeolocationOptions represents the options for BrowserContext.SetGeolocation()
type SetGeolocationOptions struct {
	Longitude int  `json:"longitude"`
	Latitude  int  `json:"latitude"`
	Accuracy  *int `json:"accuracy"`
}

func (b *browserContextImpl) SetGeolocation(gelocation *SetGeolocationOptions) error {
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
		"headers": serializeHeaders(headers),
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

func (b *browserContextImpl) Route(url interface{}, handler routeHandler) error {
	b.routes = append(b.routes, newRouteHandlerEntry(newURLMatcher(url), handler))
	if len(b.routes) == 1 {
		_, err := b.channel.Send("setNetworkInterceptionEnabled", map[string]interface{}{
			"enabled": true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *browserContextImpl) Unroute(url interface{}, handlers ...routeHandler) error {
	b.Lock()
	defer b.Unlock()

	routes, err := unroute(b.channel, b.routes, url, handlers...)
	if err != nil {
		return err
	}
	b.routes = routes

	return nil
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
	_, err := b.channel.Send("close")
	return err
}

type StorageState struct {
	Cookies []Cookie       `json:"cookies"`
	Origins []OriginsState `json:"origins"`
}

type Cookie struct {
	Name     string  `json:"name"`
	Value    string  `json:"value"`
	URL      string  `json:"url"`
	Domain   string  `json:"domain"`
	Path     string  `json:"path"`
	Expires  float64 `json:"expires"`
	HttpOnly bool    `json:"httpOnly"`
	Secure   bool    `json:"secure"`
	SameSite string  `json:"sameSite"`
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
		file, err := os.Open(paths[0])
		if err != nil {
			return nil, err
		}
		if err := json.NewDecoder(file).Decode(result); err != nil {
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
	b.isClosedOrClosing = true
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
}

func (b *browserContextImpl) onRoute(route *routeImpl, request *requestImpl) {
	go func() {
		for _, handlerEntry := range b.routes {
			if handlerEntry.matcher.Matches(request.URL()) {
				handlerEntry.handler(route, request)
				return
			}
		}
		if err := route.Continue(); err != nil {
			log.Printf("could not continue request: %v", err)
		}
	}()
}

func newBrowserContext(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *browserContextImpl {
	bt := &browserContextImpl{
		timeoutSettings: newTimeoutSettings(nil),
		pages:           make([]Page, 0),
		routes:          make([]*routeHandlerEntry, 0),
		bindings:        make(map[string]BindingCallFunction),
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.channel.On("bindingCall", func(params map[string]interface{}) {
		bt.onBinding(fromChannel(params["binding"]).(*bindingCallImpl))
	})
	bt.channel.On("close", bt.onClose)
	bt.channel.On("page", func(payload map[string]interface{}) {
		bt.onPage(fromChannel(payload["page"]).(*pageImpl))
	})
	bt.channel.On("route", func(params map[string]interface{}) {
		bt.onRoute(fromChannel(params["route"]).(*routeImpl), fromChannel(params["request"]).(*requestImpl))
	})
	return bt
}
