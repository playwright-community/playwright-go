package playwright

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
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

type ScreenshotOptions struct {
	Path     *string `json:"path,-"`
	Type     *string `json:"type,omitempty"`
	Quality  *int    `json:"quality,omitempty"`
	FullPage *bool   `json:"fullPage,omitempty"`
	Clip     *struct {
		X *int `json:"x,omitempty"`
		Y *int `json:"y,omitempty"`
	} `json:"clip,omitempty"`
	Width          *int  `json:"width,omitempty"`
	Height         *int  `json:"height,omitempty"`
	OmitBackground *bool `json:"omitBackground,omitempty"`
	Timeout        *int  `json:"timeout,omitempty"`
}

func (b *Page) Screenshot(options ...*ScreenshotOptions) ([]byte, error) {
	var path *string
	if len(options) > 0 {
		path = options[0].Path
	}
	data, err := b.channel.Send("screenshot", options)
	if err != nil {
		return nil, fmt.Errorf("could not send message :%v", err)
	}
	image, err := base64.StdEncoding.DecodeString(data.(string))
	if err != nil {
		return nil, fmt.Errorf("could not decode base64 :%v", err)
	}
	if path != nil {
		if err := ioutil.WriteFile(*path, image, 0644); err != nil {
			return nil, err
		}
	}
	return image, nil
}

func newPage(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Page {
	channelOwner := initializer["mainFrame"].(*Channel).object
	bt := &Page{
		mainFrame: channelOwner.(*Frame),
	}
	bt.frames = []*Frame{bt.mainFrame}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
