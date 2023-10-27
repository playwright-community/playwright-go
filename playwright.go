// Package playwright is a library to automate Chromium, Firefox and WebKit with
// a single API. Playwright is built to enable cross-browser web automation that
// is ever-green, capable, reliable and fast.
package playwright

// DeviceDescriptor represents a single device
type DeviceDescriptor struct {
	UserAgent          string  `json:"userAgent"`
	Viewport           *Size   `json:"viewport"`
	Screen             *Size   `json:"screen"`
	DeviceScaleFactor  float64 `json:"deviceScaleFactor"`
	IsMobile           bool    `json:"isMobile"`
	HasTouch           bool    `json:"hasTouch"`
	DefaultBrowserType string  `json:"defaultBrowserType"`
}

// Playwright represents a Playwright instance
type Playwright struct {
	channelOwner
	Selectors Selectors
	Chromium  BrowserType
	Firefox   BrowserType
	WebKit    BrowserType
	Request   APIRequest
	Devices   map[string]*DeviceDescriptor
}

// Stop stops the Playwright instance
func (p *Playwright) Stop() error {
	return p.connection.Stop()
}

func (p *Playwright) setSelectors(selectors Selectors) {
	selectorsOwner := fromChannel(p.initializer["selectors"]).(*selectorsOwnerImpl)
	p.Selectors.(*selectorsImpl).removeChannel(selectorsOwner)
	p.Selectors = selectors
	p.Selectors.(*selectorsImpl).addChannel(selectorsOwner)
}

func newPlaywright(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *Playwright {
	pw := &Playwright{
		Selectors: newSelectorsImpl(),
		Chromium:  fromChannel(initializer["chromium"]).(*browserTypeImpl),
		Firefox:   fromChannel(initializer["firefox"]).(*browserTypeImpl),
		WebKit:    fromChannel(initializer["webkit"]).(*browserTypeImpl),
		Devices:   make(map[string]*DeviceDescriptor),
	}
	pw.createChannelOwner(pw, parent, objectType, guid, initializer)
	pw.Request = newApiRequestImpl(pw)
	pw.Chromium.(*browserTypeImpl).playwright = pw
	pw.Firefox.(*browserTypeImpl).playwright = pw
	pw.WebKit.(*browserTypeImpl).playwright = pw
	selectorsOwner := fromChannel(initializer["selectors"]).(*selectorsOwnerImpl)
	pw.Selectors.(*selectorsImpl).addChannel(selectorsOwner)
	pw.connection.afterClose = func() {
		pw.Selectors.(*selectorsImpl).removeChannel(selectorsOwner)
	}
	if pw.connection.localUtils != nil {
		pw.Devices = pw.connection.localUtils.Devices
	}
	return pw
}

//go:generate bash scripts/generate-api.sh
