package playwright

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type browserImpl struct {
	channelOwner
	isConnected              bool
	isClosedOrClosing        bool
	isConnectedOverWebSocket bool
	contexts                 []BrowserContext
}

func (b *browserImpl) IsConnected() bool {
	b.RLock()
	defer b.RUnlock()
	return b.isConnected
}

func (b *browserImpl) NewContext(options ...BrowserNewContextOptions) (BrowserContext, error) {
	overrides := map[string]interface{}{}
	if len(options) == 1 {
		if options[0].ExtraHttpHeaders != nil {
			overrides["extraHTTPHeaders"] = serializeMapToNameAndValue(options[0].ExtraHttpHeaders)
			options[0].ExtraHttpHeaders = nil
		}
		if options[0].StorageStatePath != nil {
			var storageState *BrowserNewContextOptionsStorageState
			storageString, err := ioutil.ReadFile(*options[0].StorageStatePath)
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
		if options[0].NoViewport != nil && *options[0].NoViewport {
			overrides["noDefaultViewport"] = true
			options[0].NoViewport = nil
		}
	}
	channel, err := b.channel.Send("newContext", overrides, options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	context := fromChannel(channel).(*browserContextImpl)
	if len(options) == 1 {
		context.options = &options[0]
	}
	context.browser = b
	b.Lock()
	b.contexts = append(b.contexts, context)
	b.Unlock()
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
	_, err := b.channel.Send("close")
	if err != nil {
		return fmt.Errorf("could not send message: %w", err)
	}
	if b.isConnectedOverWebSocket {
		return b.connection.Stop()
	}
	return nil
}

func (b *browserImpl) Version() string {
	return b.initializer["version"].(string)
}

func (b *browserImpl) onClose() {
	b.Lock()
	if !b.isClosedOrClosing {
		b.isConnected = false
		b.isClosedOrClosing = true
		b.Emit("disconnected")
	}
	b.Unlock()
}

func newBrowser(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *browserImpl {
	bt := &browserImpl{
		isConnected: true,
		contexts:    make([]BrowserContext, 0),
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.channel.On("close", bt.onClose)
	return bt
}
