package playwright

type Frame struct {
	ChannelOwner
	page Page
	url  string
}

func (b *Frame) URL() string {
	return b.url
}

func newFrame(parent *ChannelOwner, objectType string, guid string, initializer interface{}) *Frame {
	bt := &Frame{
		url: initializer.(map[string]interface{})["url"].(string),
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
