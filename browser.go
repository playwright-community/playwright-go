package playwright

import "fmt"

type Browser struct {
	ChannelOwner
	IsConnected bool
	contexts    []*BrowserContext
}

func (b *Browser) NewContext(options ...BrowserNewContextOptions) (*BrowserContext, error) {
	channel, err := b.channel.Send("newContext", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	context := fromChannel(channel).(*BrowserContext)
	context.browser = b
	b.contexts = append(b.contexts, context)
	return context, nil
}

func (b *Browser) NewPage(options ...BrowserNewContextOptions) (*Page, error) {
	context, err := b.NewContext(options...)
	if err != nil {
		return nil, err
	}
	page, err := context.NewPage()
	if err != nil {
		return nil, err
	}
	page.ownedContext = context
	context.ownedPage = page
	return page, nil
}

func (b *Browser) Contexts() []*BrowserContext {
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
		IsConnected: true,
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
