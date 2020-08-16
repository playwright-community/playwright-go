package playwright

import (
	"encoding/base64"
	"fmt"
)

type Page struct {
	ChannelOwner
	frames    []*Frame
	mainFrame *Frame
}

func (b *Page) Goto(url string) error {
	return b.mainFrame.Goto(url)
}

func (b *Page) URL() string {
	return b.mainFrame.URL()
}

func (b *Page) SetContent(content string) error {
	return b.mainFrame.SetContent(content)
}

func (b *Page) Content() (string, error) {
	return b.mainFrame.Content()
}

func (b *Page) Screenshot() ([]byte, error) {
	data, err := b.channel.Send("screenshot", nil)
	if err != nil {
		return nil, fmt.Errorf("could not send message :%v", err)
	}
	image, err := base64.StdEncoding.DecodeString(data.(string))
	if err != nil {
		return nil, fmt.Errorf("could not decode base64 :%v", err)
	}
	return image, nil
}

func newPage(parent *ChannelOwner, objectType string, guid string, initializer interface{}) *Page {
	channelOwner := (initializer.(map[string]interface{})["mainFrame"]).(*Channel).object
	bt := &Page{
		mainFrame: channelOwner.(*Frame),
	}
	bt.frames = []*Frame{bt.mainFrame}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
