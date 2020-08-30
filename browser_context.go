package playwright

import (
	"fmt"
	"reflect"
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
	channel, err := b.channel.Send("newPage", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %v", err)
	}
	return fromChannel(channel).(*Page), nil
}

func (b *BrowserContext) WaitForEvent(event string, predicate ...interface{}) interface{} {
	evChan := make(chan interface{}, 1)
	b.Once(event, func(ev ...interface{}) {
		if len(predicate) == 0 {
			evChan <- ev[0]
		} else if len(predicate) == 1 {
			result := reflect.ValueOf(predicate[0]).Call([]reflect.Value{reflect.ValueOf(ev[0])})
			if result[0].Bool() {
				evChan <- ev[0]
			}
		}
	})
	return <-evChan
}

func (b *BrowserContext) ExpectEvent(event string, cb func() error) (interface{}, error) {
	return newExpectWrapper(b.WaitForEvent, []interface{}{event}, cb)
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
		bt.Emit("page", page)
	})
	return bt
}
