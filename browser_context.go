package playwright

import (
	"fmt"
	"sync"
)

type BrowserContext struct {
	ChannelOwner
	pagesMutex sync.Mutex
	pages      []*Page
}

func (b *BrowserContext) Pages() []*Page {
	b.pagesMutex.Lock()
	defer b.pagesMutex.Unlock()
	return b.pages
}

func (b *BrowserContext) NewPage(options ...BrowserNewPageOptions) (*Page, error) {
	channelOwner, err := b.channel.Send("newPage", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %v", err)
	}
	return fromChannel(channelOwner).(*Page), nil
}

func newBrowserContext(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *BrowserContext {
	bt := &BrowserContext{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.channel.On("page", func(payload ...interface{}) {
		page := fromChannel(payload[0].(map[string]interface{})["page"]).(*Page)
		page.browserContext = bt
		bt.pagesMutex.Lock()
		bt.pages = append(bt.pages, page)
		bt.pagesMutex.Unlock()
		bt.Emit("request", page)
	})
	return bt
}
