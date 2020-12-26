package playwright

import (
	"fmt"
	"sync"
)

type Browser struct {
	ChannelOwner
	isConnected bool
	contexts    []BrowserContextI
	contextsMu  sync.Mutex
}

func (b *Browser) IsConnected() bool {
	b.Lock()
	defer b.Unlock()
	return b.isConnected
}

func (b *Browser) NewContext(options ...BrowserNewContextOptions) (BrowserContextI, error) {
	channel, err := b.channel.Send("newContext", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	context := fromChannel(channel).(*BrowserContext)
	context.browser = b
	b.contextsMu.Lock()
	b.contexts = append(b.contexts, context)
	b.contextsMu.Unlock()
	return context, nil
}

func (b *Browser) NewPage(options ...BrowserNewContextOptions) (PageI, error) {
	context, err := b.NewContext(options...)
	if err != nil {
		return nil, err
	}
	page, err := context.NewPage()
	if err != nil {
		return nil, err
	}
	page.(*Page).ownedContext = context
	context.(*BrowserContext).ownedPage = page
	return page, nil
}

func (b *Browser) Contexts() []BrowserContextI {
	b.contextsMu.Lock()
	defer b.contextsMu.Unlock()
	return b.contexts
}

func (b *Browser) Close() error {
	_, err := b.channel.Send("close")
	return err
}

func (b *Browser) Version() string {
	return b.initializer["version"].(string)
}

func newBrowser(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Browser {
	bt := &Browser{
		isConnected: true,
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.channel.On("close", func(ev map[string]interface{}) {
		bt.Lock()
		bt.isConnected = false
		bt.Unlock()
		bt.Emit("close")
	})
	return bt
}
