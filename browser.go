package playwright

import (
	"encoding/json"
	"fmt"
	"os"
)

type browserImpl struct {
	channelOwner
	isConnected                  bool
	isClosedOrClosing            bool
	shouldCloseConnectionOnClose bool
	contexts                     []BrowserContext
	browserType                  BrowserType
}

func (b *browserImpl) BrowserType() BrowserType {
	return b.browserType
}

func (b *browserImpl) IsConnected() bool {
	b.RLock()
	defer b.RUnlock()
	return b.isConnected
}

func (b *browserImpl) NewContext(options ...BrowserNewContextOptions) (BrowserContext, error) {
	overrides := map[string]interface{}{}
	option := BrowserNewContextOptions{}
	if len(options) == 1 {
		option = options[0]
	}
	if option.ExtraHttpHeaders != nil {
		overrides["extraHTTPHeaders"] = serializeMapToNameAndValue(options[0].ExtraHttpHeaders)
		options[0].ExtraHttpHeaders = nil
	}
	if option.StorageStatePath != nil {
		var storageState *OptionalStorageState
		storageString, err := os.ReadFile(*options[0].StorageStatePath)
		if err != nil {
			return nil, fmt.Errorf("could not read storage state file: %w", err)
		}
		err = json.Unmarshal(storageString, &storageState)
		if err != nil {
			return nil, fmt.Errorf("could not parse storage state file: %w", err)
		}
		options[0].StorageState = storageState
		options[0].StorageStatePath = nil
	}
	if option.NoViewport != nil && *options[0].NoViewport {
		overrides["noDefaultViewport"] = true
		options[0].NoViewport = nil
	}
	if option.RecordHarPath != nil {
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
	channel, err := b.channel.Send("newContext", overrides, options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	context := fromChannel(channel).(*browserContextImpl)
	context.options = &option
	b.Lock()
	b.contexts = append(b.contexts, context)
	b.Unlock()
	context.browser = b
	context.setBrowserType(b.browserType.(*browserTypeImpl))
	return context, nil
}

func (b *browserImpl) NewPage(options ...BrowserNewContextOptions) (Page, error) {
	context, err := b.NewContext(options...)
	if err != nil {
		return nil, err
	}
	page, err := context.NewPage()
	if err != nil {
		return nil, err
	}
	page.(*pageImpl).ownedContext = context
	context.(*browserContextImpl).ownedPage = page
	return page, nil
}

func (b *browserImpl) NewBrowserCDPSession() (CDPSession, error) {
	channel, err := b.channel.Send("newBrowserCDPSession")
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}

	cdpSession := fromChannel(channel).(*cdpSessionImpl)

	return cdpSession, nil
}

func (b *browserImpl) Contexts() []BrowserContext {
	b.Lock()
	defer b.Unlock()
	return b.contexts
}

func (b *browserImpl) Close() error {
	if b.isClosedOrClosing {
		return nil
	}
	b.Lock()
	b.isClosedOrClosing = true
	b.Unlock()
	_, err := b.channel.Send("close")
	if err != nil {
		return fmt.Errorf("could not send message: %w", err)
	}
	if b.shouldCloseConnectionOnClose {
		return b.connection.Stop()
	}
	return nil
}

func (b *browserImpl) Version() string {
	return b.initializer["version"].(string)
}

func (b *browserImpl) onClose() {
	b.Lock()
	b.isClosedOrClosing = true
	if b.isConnected {
		b.isConnected = false
		b.Emit("disconnected")
	}
	b.Unlock()
}

func (b *browserImpl) setBrowserType(bt *browserTypeImpl) {
	b.browserType = bt
	for _, c := range b.contexts {
		c.(*browserContextImpl).setBrowserType(bt)
	}
}

func newBrowser(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *browserImpl {
	b := &browserImpl{
		isConnected: true,
		contexts:    make([]BrowserContext, 0),
	}
	b.createChannelOwner(b, parent, objectType, guid, initializer)
	// convert parent to *browserTypeImpl
	b.browserType = newBrowserType(parent.parent, parent.objectType, parent.guid, parent.initializer)
	b.channel.On("close", b.onClose)
	return b
}
