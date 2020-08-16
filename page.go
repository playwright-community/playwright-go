package playwright

type Page struct {
	ChannelOwner
	frames    []*Frame
	mainFrame *Frame
}

func (b *Page) Goto(url string) error {
	return b.mainFrame.Goto(url)
}

func (b *Page) Screenshot(path string) error {
	return nil
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

func newPage(parent *ChannelOwner, objectType string, guid string, initializer interface{}) *Page {
	channelOwner := (initializer.(map[string]interface{})["mainFrame"]).(*Channel).object
	bt := &Page{
		mainFrame: channelOwner.(*Frame),
	}
	bt.frames = []*Frame{bt.mainFrame}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
