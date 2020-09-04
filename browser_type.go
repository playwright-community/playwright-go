package playwright

import (
	"fmt"
)

type BrowserType struct {
	ChannelOwner
}

func (b *BrowserType) Name() string {
	return b.initializer["name"].(string)
}

func (b *BrowserType) ExecutablePath() string {
	return b.initializer["executablePath"].(string)
}

func (b *BrowserType) Launch(options ...BrowserTypeLaunchOptions) (*Browser, error) {
	channel, err := b.channel.Send("launch", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	return fromChannel(channel).(*Browser), nil
}

func newBrowserType(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *BrowserType {
	bt := &BrowserType{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
