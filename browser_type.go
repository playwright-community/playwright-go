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
	browser := fromChannel(channel).(*browserImpl)
	browser.setBrowserType(b)
	return browser, nil
}

func (b *browserTypeImpl) LaunchPersistentContext(userDataDir string, options ...BrowserTypeLaunchPersistentContextOptions) (BrowserContext, error) {
	overrides := map[string]interface{}{
		"userDataDir": userDataDir,
	}
	option := &BrowserNewContextOptions{}
	if len(options) == 1 {
		err := assignStructFields(option, options[0], true)
		if err != nil {
			return nil, fmt.Errorf("can not convert options: %w", err)
		}
		if options[0].ExtraHttpHeaders != nil {
			overrides["extraHTTPHeaders"] = serializeMapToNameAndValue(options[0].ExtraHttpHeaders)
			options[0].ExtraHttpHeaders = nil
		}
		if options[0].Env != nil {
			overrides["env"] = serializeMapToNameAndValue(options[0].Env)
			options[0].Env = nil
		}
		if options[0].NoViewport != nil && *options[0].NoViewport {
			overrides["noDefaultViewport"] = true
			options[0].NoViewport = nil
		}
		if options[0].RecordHarPath != nil {
			overrides["recordHar"] = prepareRecordHarOptions(recordHarInputOptions{
				Path:        *options[0].RecordHarPath,
				URL:         options[0].RecordHarUrlFilter,
				Mode:        options[0].RecordHarMode,
				Content:     options[0].RecordHarContent,
				OmitContent: options[0].RecordHarOmitContent,
			})
			options[0].RecordHarPath = nil
			options[0].RecordHarUrlFilter = nil
			options[0].RecordHarMode = nil
			options[0].RecordHarContent = nil
			options[0].RecordHarOmitContent = nil
		}
	}
	channel, err := b.channel.Send("launchPersistentContext", overrides, options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	context := fromChannel(channel).(*browserContextImpl)
	context.options = option
	context.setBrowserType(b)
	return context, nil
}
func (b *browserTypeImpl) Connect(url string, options ...BrowserTypeConnectOptions) (Browser, error) {
	overrides := map[string]interface{}{
		"wsEndpoint": url,
	}
	localUtils := b.connection.LocalUtils()
	pipe, err := localUtils.channel.Send("connect", overrides, options)
	if err != nil {
		return nil, err
	}
	jsonPipe := fromChannel(pipe).(*jsonPipe)
	connection := newConnection(jsonPipe.Close, localUtils)
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
		connection.cleanup()
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
	browser.shouldCloseConnectionOnClose = true
	browser.setBrowserType(b)
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
	browser.setBrowserType(b)
	return browser, nil
}

func newBrowserType(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *browserTypeImpl {
	bt := &browserTypeImpl{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
