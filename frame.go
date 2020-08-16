package playwright

type Frame struct {
	ChannelOwner
	page Page
	url  string
}

func (b *Frame) URL() string {
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

func newFrame(parent *ChannelOwner, objectType string, guid string, initializer interface{}) *Frame {
	bt := &Frame{
		url: initializer.(map[string]interface{})["url"].(string),
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
