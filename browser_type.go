package playwright

import (
	"fmt"
)

type BrowserType struct {
	ChannelOwner
}

func (b *BrowserType) Name() string {
	return b.initializer.(map[string]interface{})["name"].(string)
}

func (b *BrowserType) ExecutablePath() string {
	return b.initializer.(map[string]interface{})["executablePath"].(string)
}

func (b *BrowserType) Launch() (*Browser, error) {
	channelOwner, err := b.channel.Send("launch", nil)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %v", err)
	}
	return channelOwner.(*Channel).object.(*Browser), nil
}

func newBrowserType(parent *ChannelOwner, objectType string, guid string, initializer interface{}) *BrowserType {
	bt := &BrowserType{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
