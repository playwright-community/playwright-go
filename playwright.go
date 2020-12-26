// Package playwright is a library to automate Chromium, Firefox and WebKit with
// a single API. Playwright is built to enable cross-browser web automation that
// is ever-green, capable, reliable and fast.
package playwright

type DeviceDescriptor struct {
	UserAgent          string                     `json:"userAgent"`
	Viewport           *BrowserNewContextViewport `json:"viewport"`
	DeviceScaleFactor  int                        `json:"deviceScaleFactor"`
	IsMobile           bool                       `json:"isMobile"`
	HasTouch           bool                       `json:"hasTouch"`
	DefaultBrowserType string                     `json:"defaultBrowserType"`
}

type Playwright struct {
	ChannelOwner
	Chromium BrowserTypeI
	Firefox  BrowserTypeI
	WebKit   BrowserTypeI
	Devices  map[string]*DeviceDescriptor
}

func (p *Playwright) Stop() error {
	return p.connection.Stop()
}

func newPlaywright(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Playwright {
	// TODO: add selectors
	pw := &Playwright{
		Chromium: fromChannel(initializer["chromium"]).(*BrowserType),
		Firefox:  fromChannel(initializer["firefox"]).(*BrowserType),
		WebKit:   fromChannel(initializer["webkit"]).(*BrowserType),
		Devices:  make(map[string]*DeviceDescriptor),
	}
	for _, dd := range initializer["deviceDescriptors"].([]interface{}) {
		entry := dd.(map[string]interface{})
		pw.Devices[entry["name"].(string)] = &DeviceDescriptor{
			Viewport: &BrowserNewContextViewport{},
		}
		remapMapToStruct(entry["descriptor"], pw.Devices[entry["name"].(string)])
	}
	pw.createChannelOwner(pw, parent, objectType, guid, initializer)
	return pw
}
