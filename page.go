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

func (b *Page) SetContent(content string, options ...PageSetContentOptions) error {
	return b.mainFrame.SetContent(content, options...)
}

func (b *Page) Content() (string, error) {
	return b.mainFrame.Content()
}

func (b *Page) Evaluate(expression string, options ...interface{}) (interface{}, error) {
	return b.mainFrame.Evaluate(expression, options...)
}

func (b *Page) Screenshot(options ...PageScreenshotOptions) ([]byte, error) {
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

func (b *Page) QuerySelector(selector string) (*ElementHandle, error) {
	return b.mainFrame.QuerySelector(selector)
}

func newPage(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Page {
	channelOwner := initializer["mainFrame"].(*Channel).object
	bt := &Page{
		mainFrame: channelOwner.(*Frame),
	}
	bt.frames = []*Frame{bt.mainFrame}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.channel.On("console", func(payload ...interface{}) {
		bt.Emit("console", payload[0].(map[string]interface{})["message"].(*Channel).object)
	})
	return bt
}
