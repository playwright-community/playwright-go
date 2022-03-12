package playwright

import (
	"fmt"
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
	}
	if len(options) == 1 {
		if options[0].ExtraHttpHeaders != nil {
			overrides["extraHTTPHeaders"] = serializeMapToNameAndValue(options[0].ExtraHttpHeaders)
		}
		if options[0].Env != nil {
			overrides["env"] = serializeMapToNameAndValue(options[0].Env)
			options[0].Env = nil
		}
		if options[0].NoViewport != nil && *options[0].NoViewport {
			overrides["noDefaultViewport"] = true
			options[0].NoViewport = nil
		}
	}
	channel, err := b.channel.Send("launchPersistentContext", overrides, options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	return fromChannel(channel).(*browserContextImpl), nil
}
func (b *browserTypeImpl) Connect(url string, options ...BrowserTypeConnectOptions) (Browser, error) {
	overrides := map[string]interface{}{
		"wsEndpoint": url,
	}
	pipe, err := b.channel.Send("connect", overrides, options)
	if err != nil {
		return nil, err
	}
	jsonPipe := fromChannel(pipe).(*jsonPipe)
	connection := newConnection(jsonPipe.Close)
	connection.isRemote = true
	var browser *browserImpl
	pipeClosed := func() {
		for _, context := range browser.contexts {
			pages := context.(*browserContextImpl).pages
			for _, page := range pages {
				page.(*pageImpl).onClose()
			}
			context.(*browserContextImpl).onClose()
		}
		browser.onClose()
	}
	jsonPipe.On("closed", pipeClosed)
	connection.onmessage = func(message map[string]interface{}) error {
		if err := jsonPipe.Send(message); err != nil {
			pipeClosed()
			return err
		}
		return nil
	}
	jsonPipe.On("message", connection.Dispatch)
	playwright := connection.Start()
	browser = fromChannel(playwright.initializer["preLaunchedBrowser"]).(*browserImpl)
	browser.isConnectedOverWebSocket = true
	return browser, nil
}
func (b *browserTypeImpl) ConnectOverCDP(endpointURL string, options ...BrowserTypeConnectOverCDPOptions) (Browser, error) {
	overrides := map[string]interface{}{
		"endpointURL": endpointURL,
	}
	response, err := b.channel.SendReturnAsDict("connectOverCDP", overrides, options)
	if err != nil {
		return nil, err
	}
	browser := fromChannel(response.(map[string]interface{})["browser"]).(*browserImpl)
	if defaultContext, ok := response.(map[string]interface{})["defaultContext"]; ok {
		context := fromChannel(defaultContext).(*browserContextImpl)
		browser.contexts = append(browser.contexts, context)
		context.browser = browser
	}
	return browser, nil
}

func newBrowserType(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *browserTypeImpl {
	bt := &browserTypeImpl{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
