package playwright

type Playwright struct {
	channelOwner ChannelOwner
	Chromium     *BrowserType
	Firefox      *BrowserType
	WebKit       *BrowserType
}

func (p *Playwright) Stop() error {
	return nil
}

func newPlaywright(parent *ChannelOwner, objectType string, guid string, initializer interface{}) *Playwright {
	// TODO: add devices and selectors
	return &Playwright{
		Chromium: (initializer.(map[string]interface{})["chromium"]).(*Channel).object.(*BrowserType),
		Firefox:  (initializer.(map[string]interface{})["firefox"]).(*Channel).object.(*BrowserType),
		WebKit:   (initializer.(map[string]interface{})["webkit"]).(*Channel).object.(*BrowserType),
	}
}
