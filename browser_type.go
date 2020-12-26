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

func (b *BrowserType) Launch(options ...BrowserTypeLaunchOptions) (BrowserI, error) {
	channel, err := b.channel.Send("launch", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	return fromChannel(channel).(*Browser), nil
}

func (b *BrowserType) LaunchPersistentContext(userDataDir string, options ...BrowserTypeLaunchPersistentContextOptions) (BrowserContextI, error) {
	overrides := map[string]interface{}{
		"userDataDir": userDataDir,
	}
	if len(options) == 1 && options[0].ExtraHTTPHeaders != nil {
		overrides["extraHTTPHeaders"] = serializeHeaders(options[0].ExtraHTTPHeaders)
	}
	channel, err := b.channel.Send("launchPersistentContext", options, overrides)
	if err != nil {
		return nil, fmt.Errorf("could not send message: %w", err)
	}
	return fromChannel(channel).(*BrowserContext), nil
}

func newBrowserType(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *BrowserType {
	bt := &BrowserType{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
