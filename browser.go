package playwright

import "fmt"

type Browser struct {
	ChannelOwner
	IsConnected bool
	Contexts    []*BrowserContext
}

func (b *Browser) NewContext(options ...BrowserNewContextOptions) (*BrowserContext, error) {
	channelOwner, err := b.channel.Send("newContext", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %v", err)
	}
	context := fromChannel(channelOwner).(*BrowserContext)
	b.Contexts = append(b.Contexts, context)
	return context, nil
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
		IsConnected: true,
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
