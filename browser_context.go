package playwright

import "fmt"

type BrowserContext struct {
	ChannelOwner
	Pages []Page
}

func (b *BrowserContext) NewPage() (*Page, error) {
	channelOwner, err := b.channel.Send("newPage", nil)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %v", err)
	}
	return channelOwner.(*Channel).object.(*Page), nil
}

func newBrowserContext(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *BrowserContext {
	bt := &BrowserContext{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
