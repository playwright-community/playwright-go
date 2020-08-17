package playwright

import "sync"

type Frame struct {
	ChannelOwner
	sync.RWMutex
	page Page
	url  string
}

func (b *Frame) URL() string {
	b.RLock()
	defer b.RUnlock()
	return b.url
}

func (b *Frame) SetContent(content string) error {
	_, err := b.channel.Send("setContent", map[string]interface{}{
		"html": content,
	})
	return err
}

func (b *Frame) Content() (string, error) {
	content, err := b.channel.Send("content", nil)
	return content.(string), err
}

func (b *Frame) Goto(url string) error {
	_, err := b.channel.Send("goto", map[string]interface{}{
		"url": url,
	})
	return err
}

func (b *Frame) onFrameNavigated(event ...interface{}) {
	b.Lock()
	b.url = event[0].(map[string]interface{})["url"].(string)
	b.Unlock()
}

func newFrame(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Frame {
	bt := &Frame{
		url: initializer["url"].(string),
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.channel.On("navigated", bt.onFrameNavigated)
	return bt
}
