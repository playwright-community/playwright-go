package playwright

import (
	"fmt"
	"io/ioutil"
	"sync"
)

type browserContextImpl struct {
	channelOwner
	timeoutSettings   *timeoutSettings
	isClosedOrClosing bool
	options           *BrowserNewContextOptions
	pagesMutex        sync.Mutex
	pages             []Page
	ownedPage         Page
	browser           *browserImpl
}

func (b *browserContextImpl) Browser() Browser {
	return b.browser
}

func (b *browserContextImpl) SetDefaultNavigationTimeout(timeout int) {
	b.timeoutSettings.SetNavigationTimeout(timeout)
	b.channel.SendNoReply("setDefaultNavigationTimeoutNoReply", map[string]interface{}{
		"timeout": timeout,
	})
}

func (b *browserContextImpl) SetDefaultTimeout(timeout int) {
	b.timeoutSettings.SetTimeout(timeout)
	b.channel.SendNoReply("setDefaultTimeoutNoReply", map[string]interface{}{
		"timeout": timeout,
	})
}

func (b *browserContextImpl) Pages() []Page {
	b.pagesMutex.Lock()
	defer b.pagesMutex.Unlock()
	return b.pages
}

func (b *browserContextImpl) NewPage(options ...BrowserNewPageOptions) (Page, error) {
	channel, err := b.channel.Send("newPage", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	return fromChannel(channel).(*pageImpl), nil
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

type BrowserContextAddInitScriptOptions struct {
	Path   *string
	Script *string
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

func (b *browserContextImpl) WaitForEvent(event string, predicate ...interface{}) interface{} {
	return <-waitForEvent(b, event, predicate...)
}

func (b *browserContextImpl) ExpectEvent(event string, cb func() error) (interface{}, error) {
	return newExpectWrapper(b.WaitForEvent, []interface{}{event}, cb)
}

func (b *browserContextImpl) Close() error {
	b.Lock()
	if b.isClosedOrClosing {
		return nil
	}
	b.isClosedOrClosing = true
	b.Unlock()
	_, err := b.channel.Send("close")
	return err
}

func newBrowserContext(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *browserContextImpl {
	bt := &browserContextImpl{
		timeoutSettings: newTimeoutSettings(nil),
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.channel.On("page", func(payload map[string]interface{}) {
		page := fromChannel(payload["page"]).(*pageImpl)
		page.browserContext = bt
		bt.pagesMutex.Lock()
		bt.pages = append(bt.pages, page)
		bt.pagesMutex.Unlock()
		bt.Emit("page", page)
	})
	bt.channel.On("close", func() {
		if bt.browser != nil {
			contexts := make([]BrowserContext, 0)
			bt.browser.contextsMu.Lock()
			for _, context := range bt.browser.contexts {
				if context != bt {
					contexts = append(contexts, context)
				}
			}
			bt.browser.contexts = contexts
			bt.browser.contextsMu.Unlock()
		}
		bt.Emit("close")
	})
	return bt
}
