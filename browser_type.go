package playwright

import (
	"fmt"
	"log"
)

type browserTypeImpl struct {
	channelOwner
}

func (b *browserTypeImpl) Name() string {
	return b.initializer["name"].(string)
}

func (b *browserTypeImpl) ExecutablePath() string {
	return b.initializer["executablePath"].(string)
}

func (b *browserTypeImpl) Launch(options ...BrowserTypeLaunchOptions) (Browser, error) {
	overrides := map[string]interface{}{}
	if len(options) == 1 && options[0].Env != nil {
		overrides["env"] = serializeMapToNameAndValue(options[0].Env)
		options[0].Env = nil
	}
	channel, err := b.channel.Send("launch", overrides, options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	return fromChannel(channel).(*browserImpl), nil
}

func (b *browserTypeImpl) LaunchPersistentContext(userDataDir string, options ...BrowserTypeLaunchPersistentContextOptions) (BrowserContext, error) {
	overrides := map[string]interface{}{
		"userDataDir": userDataDir,
		"sdkLanguage": "javascript",
	}
	if len(options) == 1 {
		if options[0].ExtraHttpHeaders != nil {
			overrides["extraHTTPHeaders"] = serializeMapToNameAndValue(options[0].ExtraHttpHeaders)
		}
		if options[0].Env != nil {
			overrides["env"] = serializeMapToNameAndValue(options[0].Env)
			options[0].Env = nil
		}
	}
	channel, err := b.channel.Send("launchPersistentContext", overrides, options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	return fromChannel(channel).(*browserContextImpl), nil
}
func (b *browserTypeImpl) Connect(url string) (Browser, error) {
	transport := newWebSocketTransport(url)
	connection := newConnection(transport, transport.Stop)
	go func() {
		err := connection.Start()
		if err != nil {
			log.Fatalf("could not start connection: %v", err)
		}
	}()
	obj, err := connection.CallOnObjectWithKnownName("Playwright")
	if err != nil {
		return nil, fmt.Errorf("could not call object: %w", err)
	}
	playwright := obj.(*Playwright)
	browser := fromChannel(playwright.initializer["preLaunchedBrowser"]).(*browserImpl)
	browser.isConnectedOverWebSocket = true
	transport.(*webSocketTransport).OnClose = func() {
		for _, context := range browser.contexts {
			pages := context.(*browserContextImpl).pages
			for _, page := range pages {
				page.(*pageImpl).onClose()
			}
			context.(*browserContextImpl).onClose()
		}
		browser.onClose()
	}
	return browser, nil
}

func newBrowserType(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *browserTypeImpl {
	bt := &browserTypeImpl{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
