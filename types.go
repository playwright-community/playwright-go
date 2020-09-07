package playwright

type BrowserNewContextOptions struct {
	AcceptDownloads   *bool                             `json:"acceptDownloads"`
	IgnoreHTTPSErrors *bool                             `json:"ignoreHTTPSErrors"`
	BypassCSP         *bool                             `json:"bypassCSP"`
	Viewport          *BrowserNewContextViewport        `json:"viewport"`
	UserAgent         *string                           `json:"userAgent"`
	DeviceScaleFactor *int                              `json:"deviceScaleFactor"`
	IsMobile          *bool                             `json:"isMobile"`
	HasTouch          *bool                             `json:"hasTouch"`
	JavaScriptEnabled *bool                             `json:"javaScriptEnabled"`
	TimezoneId        *string                           `json:"timezoneId"`
	Geolocation       *BrowserNewContextGeolocation     `json:"geolocation"`
	Locale            *string                           `json:"locale"`
	Permissions       []string                          `json:"permissions"`
	ExtraHTTPHeaders  map[string]string                 `json:"extraHTTPHeaders"`
	Offline           *bool                             `json:"offline"`
	HttpCredentials   *BrowserNewContextHttpCredentials `json:"httpCredentials"`
	ColorScheme       *string                           `json:"colorScheme"`
	Logger            interface{}                       `json:"logger"`
	RecordVideos      *BrowserNewContextRecordVideos    `json:"_recordVideos"`
}
type BrowserNewPageOptions struct {
	AcceptDownloads   *bool                          `json:"acceptDownloads"`
	IgnoreHTTPSErrors *bool                          `json:"ignoreHTTPSErrors"`
	BypassCSP         *bool                          `json:"bypassCSP"`
	Viewport          *BrowserNewPageViewport        `json:"viewport"`
	UserAgent         *string                        `json:"userAgent"`
	DeviceScaleFactor *int                           `json:"deviceScaleFactor"`
	IsMobile          *bool                          `json:"isMobile"`
	HasTouch          *bool                          `json:"hasTouch"`
	JavaScriptEnabled *bool                          `json:"javaScriptEnabled"`
	TimezoneId        *string                        `json:"timezoneId"`
	Geolocation       *BrowserNewPageGeolocation     `json:"geolocation"`
	Locale            *string                        `json:"locale"`
	Permissions       []string                       `json:"permissions"`
	ExtraHTTPHeaders  map[string]string              `json:"extraHTTPHeaders"`
	Offline           *bool                          `json:"offline"`
	HttpCredentials   *BrowserNewPageHttpCredentials `json:"httpCredentials"`
	ColorScheme       *string                        `json:"colorScheme"`
	Logger            interface{}                    `json:"logger"`
	RecordVideos      *BrowserNewPageRecordVideos    `json:"_recordVideos"`
}
type BrowserContextGrantPermissionsOptions struct {
	Origin *string `json:"origin"`
}
type BrowserContextWaitForEventOptions struct {
	Predicate interface{} `json:"predicate"`
	Timeout   *int        `json:"timeout"`
}
type PageCloseOptions struct {
	RunBeforeUnload *bool `json:"runBeforeUnload"`
}
type PageAddScriptTagOptions struct {
	Url     *string `json:"url"`
	Path    *string `json:"path"`
	Content *string `json:"content"`
	Type    *string `json:"type"`
}
type PageAddStyleTagOptions struct {
	Url     *string `json:"url"`
	Path    *string `json:"path"`
	Content *string `json:"content"`
}
type PageCheckOptions struct {
	Force       *bool `json:"force"`
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type PageClickOptions struct {
	Button      *string            `json:"button"`
	ClickCount  *int               `json:"clickCount"`
	Delay       *int               `json:"delay"`
	Position    *PageClickPosition `json:"position"`
	Modifiers   interface{}        `json:"modifiers"`
	Force       *bool              `json:"force"`
	NoWaitAfter *bool              `json:"noWaitAfter"`
	Timeout     *int               `json:"timeout"`
}
type PageDblclickOptions struct {
	Button      *string               `json:"button"`
	Delay       *int                  `json:"delay"`
	Position    *PageDblclickPosition `json:"position"`
	Modifiers   interface{}           `json:"modifiers"`
	Force       *bool                 `json:"force"`
	NoWaitAfter *bool                 `json:"noWaitAfter"`
	Timeout     *int                  `json:"timeout"`
}
type PageDispatchEventOptions struct {
	EventInit interface{} `json:"eventInit"`
	Timeout   *int        `json:"timeout"`
}
type PageEmulateMediaOptions struct {
	Media       interface{} `json:"media"`
	ColorScheme interface{} `json:"colorScheme"`
}
type PageFillOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type PageFocusOptions struct {
	Timeout *int `json:"timeout"`
}
type PageFrameOptions struct {
	Name *string     `json:"name"`
	Url  interface{} `json:"url"`
}
type PageGetAttributeOptions struct {
	Timeout *int `json:"timeout"`
}
type PageGoBackOptions struct {
	Timeout   *int    `json:"timeout"`
	WaitUntil *string `json:"waitUntil"`
}
type PageGoForwardOptions struct {
	Timeout   *int    `json:"timeout"`
	WaitUntil *string `json:"waitUntil"`
}
type PageGotoOptions struct {
	Timeout   *int    `json:"timeout"`
	WaitUntil *string `json:"waitUntil"`
	Referer   *string `json:"referer"`
}
type PageHoverOptions struct {
	Position  *PageHoverPosition `json:"position"`
	Modifiers interface{}        `json:"modifiers"`
	Force     *bool              `json:"force"`
	Timeout   *int               `json:"timeout"`
}
type PageInnerHTMLOptions struct {
	Timeout *int `json:"timeout"`
}
type PageInnerTextOptions struct {
	Timeout *int `json:"timeout"`
}
type PagePdfOptions struct {
	Path                *string        `json:"path"`
	Scale               *int           `json:"scale"`
	DisplayHeaderFooter *bool          `json:"displayHeaderFooter"`
	HeaderTemplate      *string        `json:"headerTemplate"`
	FooterTemplate      *string        `json:"footerTemplate"`
	PrintBackground     *bool          `json:"printBackground"`
	Landscape           *bool          `json:"landscape"`
	PageRanges          *string        `json:"pageRanges"`
	Format              *string        `json:"format"`
	Width               interface{}    `json:"width"`
	Height              interface{}    `json:"height"`
	Margin              *PagePdfMargin `json:"margin"`
	PreferCSSPageSize   *bool          `json:"preferCSSPageSize"`
}
type PagePressOptions struct {
	Delay       *int  `json:"delay"`
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type PageReloadOptions struct {
	Timeout   *int    `json:"timeout"`
	WaitUntil *string `json:"waitUntil"`
}
type PageScreenshotOptions struct {
	Path           *string             `json:"path"`
	Type           *string             `json:"type"`
	Quality        *int                `json:"quality"`
	FullPage       *bool               `json:"fullPage"`
	Clip           *PageScreenshotClip `json:"clip"`
	OmitBackground *bool               `json:"omitBackground"`
	Timeout        *int                `json:"timeout"`
}
type PageSelectOptionOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type PageSetContentOptions struct {
	Timeout   *int    `json:"timeout"`
	WaitUntil *string `json:"waitUntil"`
}
type PageSetInputFilesOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type PageTextContentOptions struct {
	Timeout *int `json:"timeout"`
}
type PageTypeOptions struct {
	Delay       *int  `json:"delay"`
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type PageUncheckOptions struct {
	Force       *bool `json:"force"`
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type PageWaitForEventOptions struct {
	Predicate interface{} `json:"predicate"`
	Timeout   *int        `json:"timeout"`
}
type PageWaitForFunctionOptions struct {
	Arg     interface{} `json:"arg"`
	Polling interface{} `json:"polling"`
	Timeout *int        `json:"timeout"`
}
type PageWaitForLoadStateOptions struct {
	State   *string `json:"state"`
	Timeout *int    `json:"timeout"`
}
type PageWaitForNavigationOptions struct {
	Timeout   *int        `json:"timeout"`
	Url       interface{} `json:"url"`
	WaitUntil *string     `json:"waitUntil"`
}
type PageWaitForRequestOptions struct {
	Timeout *int `json:"timeout"`
}
type PageWaitForResponseOptions struct {
	Timeout *int `json:"timeout"`
}
type PageWaitForSelectorOptions struct {
	State   *string `json:"state"`
	Timeout *int    `json:"timeout"`
}
type FrameAddScriptTagOptions struct {
	Url     *string `json:"url"`
	Path    *string `json:"path"`
	Content *string `json:"content"`
	Type    *string `json:"type"`
}
type FrameAddStyleTagOptions struct {
	Url     *string `json:"url"`
	Path    *string `json:"path"`
	Content *string `json:"content"`
}
type FrameCheckOptions struct {
	Force       *bool `json:"force"`
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type FrameClickOptions struct {
	Button      *string             `json:"button"`
	ClickCount  *int                `json:"clickCount"`
	Delay       *int                `json:"delay"`
	Position    *FrameClickPosition `json:"position"`
	Modifiers   interface{}         `json:"modifiers"`
	Force       *bool               `json:"force"`
	NoWaitAfter *bool               `json:"noWaitAfter"`
	Timeout     *int                `json:"timeout"`
}
type FrameDblclickOptions struct {
	Button      *string                `json:"button"`
	Delay       *int                   `json:"delay"`
	Position    *FrameDblclickPosition `json:"position"`
	Modifiers   interface{}            `json:"modifiers"`
	Force       *bool                  `json:"force"`
	NoWaitAfter *bool                  `json:"noWaitAfter"`
	Timeout     *int                   `json:"timeout"`
}
type FrameDispatchEventOptions struct {
	EventInit interface{} `json:"eventInit"`
	Timeout   *int        `json:"timeout"`
}
type FrameFillOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type FrameFocusOptions struct {
	Timeout *int `json:"timeout"`
}
type FrameGetAttributeOptions struct {
	Timeout *int `json:"timeout"`
}
type FrameGotoOptions struct {
	Timeout   *int    `json:"timeout"`
	WaitUntil *string `json:"waitUntil"`
	Referer   *string `json:"referer"`
}
type FrameHoverOptions struct {
	Position  *FrameHoverPosition `json:"position"`
	Modifiers interface{}         `json:"modifiers"`
	Force     *bool               `json:"force"`
	Timeout   *int                `json:"timeout"`
}
type FrameInnerHTMLOptions struct {
	Timeout *int `json:"timeout"`
}
type FrameInnerTextOptions struct {
	Timeout *int `json:"timeout"`
}
type FramePressOptions struct {
	Delay       *int  `json:"delay"`
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type FrameSelectOptionOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type FrameSetContentOptions struct {
	Timeout   *int    `json:"timeout"`
	WaitUntil *string `json:"waitUntil"`
}
type FrameSetInputFilesOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type FrameTextContentOptions struct {
	Timeout *int `json:"timeout"`
}
type FrameTypeOptions struct {
	Delay       *int  `json:"delay"`
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type FrameUncheckOptions struct {
	Force       *bool `json:"force"`
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type FrameWaitForFunctionOptions struct {
	Arg     interface{} `json:"arg"`
	Polling interface{} `json:"polling"`
	Timeout *int        `json:"timeout"`
}
type FrameWaitForLoadStateOptions struct {
	State   *string `json:"state"`
	Timeout *int    `json:"timeout"`
}
type FrameWaitForNavigationOptions struct {
	Timeout   *int        `json:"timeout"`
	Url       interface{} `json:"url"`
	WaitUntil *string     `json:"waitUntil"`
}
type FrameWaitForSelectorOptions struct {
	State   *string `json:"state"`
	Timeout *int    `json:"timeout"`
}
type ElementHandleCheckOptions struct {
	Force       *bool `json:"force"`
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type ElementHandleClickOptions struct {
	Button      *string                     `json:"button"`
	ClickCount  *int                        `json:"clickCount"`
	Delay       *int                        `json:"delay"`
	Position    *ElementHandleClickPosition `json:"position"`
	Modifiers   interface{}                 `json:"modifiers"`
	Force       *bool                       `json:"force"`
	NoWaitAfter *bool                       `json:"noWaitAfter"`
	Timeout     *int                        `json:"timeout"`
}
type ElementHandleDblclickOptions struct {
	Button      *string                        `json:"button"`
	Delay       *int                           `json:"delay"`
	Position    *ElementHandleDblclickPosition `json:"position"`
	Modifiers   interface{}                    `json:"modifiers"`
	Force       *bool                          `json:"force"`
	NoWaitAfter *bool                          `json:"noWaitAfter"`
	Timeout     *int                           `json:"timeout"`
}
type ElementHandleFillOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type ElementHandleHoverOptions struct {
	Position  *ElementHandleHoverPosition `json:"position"`
	Modifiers interface{}                 `json:"modifiers"`
	Force     *bool                       `json:"force"`
	Timeout   *int                        `json:"timeout"`
}
type ElementHandlePressOptions struct {
	Delay       *int  `json:"delay"`
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type ElementHandleScreenshotOptions struct {
	Path           *string `json:"path"`
	Type           *string `json:"type"`
	Quality        *int    `json:"quality"`
	OmitBackground *bool   `json:"omitBackground"`
	Timeout        *int    `json:"timeout"`
}
type ElementHandleScrollIntoViewIfNeededOptions struct {
	Timeout *int `json:"timeout"`
}
type ElementHandleSelectOptionOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type ElementHandleSelectTextOptions struct {
	Timeout *int `json:"timeout"`
}
type ElementHandleSetInputFilesOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type ElementHandleTypeOptions struct {
	Delay       *int  `json:"delay"`
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type ElementHandleUncheckOptions struct {
	Force       *bool `json:"force"`
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type ElementHandleWaitForElementStateOptions struct {
	Timeout *int `json:"timeout"`
}
type ElementHandleWaitForSelectorOptions struct {
	State   *string `json:"state"`
	Timeout *int    `json:"timeout"`
}
type FileChooserSetFilesOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter"`
	Timeout     *int  `json:"timeout"`
}
type KeyboardPressOptions struct {
	Delay *int `json:"delay"`
}
type KeyboardTypeOptions struct {
	Delay *int `json:"delay"`
}
type MouseClickOptions struct {
	Button     *string `json:"button"`
	ClickCount *int    `json:"clickCount"`
	Delay      *int    `json:"delay"`
}
type MouseDblclickOptions struct {
	Button *string `json:"button"`
	Delay  *int    `json:"delay"`
}
type MouseDownOptions struct {
	Button     *string `json:"button"`
	ClickCount *int    `json:"clickCount"`
}
type MouseMoveOptions struct {
	Steps *int `json:"steps"`
}
type MouseUpOptions struct {
	Button     *string `json:"button"`
	ClickCount *int    `json:"clickCount"`
}
type SelectorsRegisterOptions struct {
	ContentScript *bool `json:"contentScript"`
}
type RouteContinueOptions struct {
	Method   *string           `json:"method"`
	PostData interface{}       `json:"postData"`
	Headers  map[string]string `json:"headers"`
}
type AccessibilitySnapshotOptions struct {
	InterestingOnly *bool          `json:"interestingOnly"`
	Root            *ElementHandle `json:"root"`
}
type BrowserTypeConnectOptions struct {
	WsEndpoint *string     `json:"wsEndpoint"`
	SlowMo     *int        `json:"slowMo"`
	Logger     interface{} `json:"logger"`
	Timeout    *int        `json:"timeout"`
}
type BrowserTypeLaunchOptions struct {
	Headless          *bool                   `json:"headless"`
	ExecutablePath    *string                 `json:"executablePath"`
	Args              []string                `json:"args"`
	IgnoreDefaultArgs interface{}             `json:"ignoreDefaultArgs"`
	Proxy             *BrowserTypeLaunchProxy `json:"proxy"`
	DownloadsPath     *string                 `json:"downloadsPath"`
	VideosPath        *string                 `json:"_videosPath"`
	ChromiumSandbox   *bool                   `json:"chromiumSandbox"`
	FirefoxUserPrefs  map[string]interface{}  `json:"firefoxUserPrefs"`
	HandleSIGINT      *bool                   `json:"handleSIGINT"`
	HandleSIGTERM     *bool                   `json:"handleSIGTERM"`
	HandleSIGHUP      *bool                   `json:"handleSIGHUP"`
	Logger            interface{}             `json:"logger"`
	Timeout           *int                    `json:"timeout"`
	Env               map[string]interface{}  `json:"env"`
	Devtools          *bool                   `json:"devtools"`
	SlowMo            *int                    `json:"slowMo"`
}
type BrowserTypeLaunchPersistentContextOptions struct {
	Headless          *bool                                              `json:"headless"`
	ExecutablePath    *string                                            `json:"executablePath"`
	Args              []string                                           `json:"args"`
	IgnoreDefaultArgs interface{}                                        `json:"ignoreDefaultArgs"`
	Proxy             *BrowserTypeLaunchPersistentContextProxy           `json:"proxy"`
	AcceptDownloads   *bool                                              `json:"acceptDownloads"`
	DownloadsPath     *string                                            `json:"downloadsPath"`
	ChromiumSandbox   *bool                                              `json:"chromiumSandbox"`
	HandleSIGINT      *bool                                              `json:"handleSIGINT"`
	HandleSIGTERM     *bool                                              `json:"handleSIGTERM"`
	HandleSIGHUP      *bool                                              `json:"handleSIGHUP"`
	Logger            interface{}                                        `json:"logger"`
	Timeout           *int                                               `json:"timeout"`
	Env               map[string]interface{}                             `json:"env"`
	Devtools          *bool                                              `json:"devtools"`
	SlowMo            *int                                               `json:"slowMo"`
	IgnoreHTTPSErrors *bool                                              `json:"ignoreHTTPSErrors"`
	BypassCSP         *bool                                              `json:"bypassCSP"`
	Viewport          *BrowserTypeLaunchPersistentContextViewport        `json:"viewport"`
	UserAgent         *string                                            `json:"userAgent"`
	DeviceScaleFactor *int                                               `json:"deviceScaleFactor"`
	IsMobile          *bool                                              `json:"isMobile"`
	HasTouch          *bool                                              `json:"hasTouch"`
	JavaScriptEnabled *bool                                              `json:"javaScriptEnabled"`
	TimezoneId        *string                                            `json:"timezoneId"`
	Geolocation       *BrowserTypeLaunchPersistentContextGeolocation     `json:"geolocation"`
	Locale            *string                                            `json:"locale"`
	Permissions       []string                                           `json:"permissions"`
	ExtraHTTPHeaders  map[string]string                                  `json:"extraHTTPHeaders"`
	Offline           *bool                                              `json:"offline"`
	HttpCredentials   *BrowserTypeLaunchPersistentContextHttpCredentials `json:"httpCredentials"`
	ColorScheme       *string                                            `json:"colorScheme"`
	VideosPath        *string                                            `json:"_videosPath"`
	RecordVideos      *BrowserTypeLaunchPersistentContextRecordVideos    `json:"_recordVideos"`
}
type BrowserTypeLaunchServerOptions struct {
	Headless          *bool                         `json:"headless"`
	Port              *int                          `json:"port"`
	ExecutablePath    *string                       `json:"executablePath"`
	Args              []string                      `json:"args"`
	IgnoreDefaultArgs interface{}                   `json:"ignoreDefaultArgs"`
	Proxy             *BrowserTypeLaunchServerProxy `json:"proxy"`
	DownloadsPath     *string                       `json:"downloadsPath"`
	VideosPath        *string                       `json:"_videosPath"`
	ChromiumSandbox   *bool                         `json:"chromiumSandbox"`
	FirefoxUserPrefs  map[string]interface{}        `json:"firefoxUserPrefs"`
	HandleSIGINT      *bool                         `json:"handleSIGINT"`
	HandleSIGTERM     *bool                         `json:"handleSIGTERM"`
	HandleSIGHUP      *bool                         `json:"handleSIGHUP"`
	Logger            interface{}                   `json:"logger"`
	Timeout           *int                          `json:"timeout"`
	Env               map[string]interface{}        `json:"env"`
	Devtools          *bool                         `json:"devtools"`
}
type ChromiumBrowserStartTracingOptions struct {
	Page        interface{} `json:"page"`
	Path        *string     `json:"path"`
	Screenshots *bool       `json:"screenshots"`
	Categories  []string    `json:"categories"`
}
type ChromiumBrowserNewContextOptions struct {
	AcceptDownloads   *bool                                     `json:"acceptDownloads"`
	IgnoreHTTPSErrors *bool                                     `json:"ignoreHTTPSErrors"`
	BypassCSP         *bool                                     `json:"bypassCSP"`
	Viewport          *ChromiumBrowserNewContextViewport        `json:"viewport"`
	UserAgent         *string                                   `json:"userAgent"`
	DeviceScaleFactor *int                                      `json:"deviceScaleFactor"`
	IsMobile          *bool                                     `json:"isMobile"`
	HasTouch          *bool                                     `json:"hasTouch"`
	JavaScriptEnabled *bool                                     `json:"javaScriptEnabled"`
	TimezoneId        *string                                   `json:"timezoneId"`
	Geolocation       *ChromiumBrowserNewContextGeolocation     `json:"geolocation"`
	Locale            *string                                   `json:"locale"`
	Permissions       []string                                  `json:"permissions"`
	ExtraHTTPHeaders  map[string]string                         `json:"extraHTTPHeaders"`
	Offline           *bool                                     `json:"offline"`
	HttpCredentials   *ChromiumBrowserNewContextHttpCredentials `json:"httpCredentials"`
	ColorScheme       *string                                   `json:"colorScheme"`
	Logger            interface{}                               `json:"logger"`
	RecordVideos      *ChromiumBrowserNewContextRecordVideos    `json:"_recordVideos"`
}
type ChromiumBrowserNewPageOptions struct {
	AcceptDownloads   *bool                                  `json:"acceptDownloads"`
	IgnoreHTTPSErrors *bool                                  `json:"ignoreHTTPSErrors"`
	BypassCSP         *bool                                  `json:"bypassCSP"`
	Viewport          *ChromiumBrowserNewPageViewport        `json:"viewport"`
	UserAgent         *string                                `json:"userAgent"`
	DeviceScaleFactor *int                                   `json:"deviceScaleFactor"`
	IsMobile          *bool                                  `json:"isMobile"`
	HasTouch          *bool                                  `json:"hasTouch"`
	JavaScriptEnabled *bool                                  `json:"javaScriptEnabled"`
	TimezoneId        *string                                `json:"timezoneId"`
	Geolocation       *ChromiumBrowserNewPageGeolocation     `json:"geolocation"`
	Locale            *string                                `json:"locale"`
	Permissions       []string                               `json:"permissions"`
	ExtraHTTPHeaders  map[string]string                      `json:"extraHTTPHeaders"`
	Offline           *bool                                  `json:"offline"`
	HttpCredentials   *ChromiumBrowserNewPageHttpCredentials `json:"httpCredentials"`
	ColorScheme       *string                                `json:"colorScheme"`
	Logger            interface{}                            `json:"logger"`
	RecordVideos      *ChromiumBrowserNewPageRecordVideos    `json:"_recordVideos"`
}
type ChromiumBrowserContextGrantPermissionsOptions struct {
	Origin *string `json:"origin"`
}
type ChromiumBrowserContextWaitForEventOptions struct {
	Predicate interface{} `json:"predicate"`
	Timeout   *int        `json:"timeout"`
}
type ChromiumCoverageStartCSSCoverageOptions struct {
	ResetOnNavigation *bool `json:"resetOnNavigation"`
}
type ChromiumCoverageStartJSCoverageOptions struct {
	ResetOnNavigation      *bool `json:"resetOnNavigation"`
	ReportAnonymousScripts *bool `json:"reportAnonymousScripts"`
}
type FirefoxBrowserNewContextOptions struct {
	AcceptDownloads   *bool                                    `json:"acceptDownloads"`
	IgnoreHTTPSErrors *bool                                    `json:"ignoreHTTPSErrors"`
	BypassCSP         *bool                                    `json:"bypassCSP"`
	Viewport          *FirefoxBrowserNewContextViewport        `json:"viewport"`
	UserAgent         *string                                  `json:"userAgent"`
	DeviceScaleFactor *int                                     `json:"deviceScaleFactor"`
	IsMobile          *bool                                    `json:"isMobile"`
	HasTouch          *bool                                    `json:"hasTouch"`
	JavaScriptEnabled *bool                                    `json:"javaScriptEnabled"`
	TimezoneId        *string                                  `json:"timezoneId"`
	Geolocation       *FirefoxBrowserNewContextGeolocation     `json:"geolocation"`
	Locale            *string                                  `json:"locale"`
	Permissions       []string                                 `json:"permissions"`
	ExtraHTTPHeaders  map[string]string                        `json:"extraHTTPHeaders"`
	Offline           *bool                                    `json:"offline"`
	HttpCredentials   *FirefoxBrowserNewContextHttpCredentials `json:"httpCredentials"`
	ColorScheme       *string                                  `json:"colorScheme"`
	Logger            interface{}                              `json:"logger"`
	RecordVideos      *FirefoxBrowserNewContextRecordVideos    `json:"_recordVideos"`
}
type FirefoxBrowserNewPageOptions struct {
	AcceptDownloads   *bool                                 `json:"acceptDownloads"`
	IgnoreHTTPSErrors *bool                                 `json:"ignoreHTTPSErrors"`
	BypassCSP         *bool                                 `json:"bypassCSP"`
	Viewport          *FirefoxBrowserNewPageViewport        `json:"viewport"`
	UserAgent         *string                               `json:"userAgent"`
	DeviceScaleFactor *int                                  `json:"deviceScaleFactor"`
	IsMobile          *bool                                 `json:"isMobile"`
	HasTouch          *bool                                 `json:"hasTouch"`
	JavaScriptEnabled *bool                                 `json:"javaScriptEnabled"`
	TimezoneId        *string                               `json:"timezoneId"`
	Geolocation       *FirefoxBrowserNewPageGeolocation     `json:"geolocation"`
	Locale            *string                               `json:"locale"`
	Permissions       []string                              `json:"permissions"`
	ExtraHTTPHeaders  map[string]string                     `json:"extraHTTPHeaders"`
	Offline           *bool                                 `json:"offline"`
	HttpCredentials   *FirefoxBrowserNewPageHttpCredentials `json:"httpCredentials"`
	ColorScheme       *string                               `json:"colorScheme"`
	Logger            interface{}                           `json:"logger"`
	RecordVideos      *FirefoxBrowserNewPageRecordVideos    `json:"_recordVideos"`
}
type WebKitBrowserNewContextOptions struct {
	AcceptDownloads   *bool                                   `json:"acceptDownloads"`
	IgnoreHTTPSErrors *bool                                   `json:"ignoreHTTPSErrors"`
	BypassCSP         *bool                                   `json:"bypassCSP"`
	Viewport          *WebKitBrowserNewContextViewport        `json:"viewport"`
	UserAgent         *string                                 `json:"userAgent"`
	DeviceScaleFactor *int                                    `json:"deviceScaleFactor"`
	IsMobile          *bool                                   `json:"isMobile"`
	HasTouch          *bool                                   `json:"hasTouch"`
	JavaScriptEnabled *bool                                   `json:"javaScriptEnabled"`
	TimezoneId        *string                                 `json:"timezoneId"`
	Geolocation       *WebKitBrowserNewContextGeolocation     `json:"geolocation"`
	Locale            *string                                 `json:"locale"`
	Permissions       []string                                `json:"permissions"`
	ExtraHTTPHeaders  map[string]string                       `json:"extraHTTPHeaders"`
	Offline           *bool                                   `json:"offline"`
	HttpCredentials   *WebKitBrowserNewContextHttpCredentials `json:"httpCredentials"`
	ColorScheme       *string                                 `json:"colorScheme"`
	Logger            interface{}                             `json:"logger"`
	RecordVideos      *WebKitBrowserNewContextRecordVideos    `json:"_recordVideos"`
}
type WebKitBrowserNewPageOptions struct {
	AcceptDownloads   *bool                                `json:"acceptDownloads"`
	IgnoreHTTPSErrors *bool                                `json:"ignoreHTTPSErrors"`
	BypassCSP         *bool                                `json:"bypassCSP"`
	Viewport          *WebKitBrowserNewPageViewport        `json:"viewport"`
	UserAgent         *string                              `json:"userAgent"`
	DeviceScaleFactor *int                                 `json:"deviceScaleFactor"`
	IsMobile          *bool                                `json:"isMobile"`
	HasTouch          *bool                                `json:"hasTouch"`
	JavaScriptEnabled *bool                                `json:"javaScriptEnabled"`
	TimezoneId        *string                              `json:"timezoneId"`
	Geolocation       *WebKitBrowserNewPageGeolocation     `json:"geolocation"`
	Locale            *string                              `json:"locale"`
	Permissions       []string                             `json:"permissions"`
	ExtraHTTPHeaders  map[string]string                    `json:"extraHTTPHeaders"`
	Offline           *bool                                `json:"offline"`
	HttpCredentials   *WebKitBrowserNewPageHttpCredentials `json:"httpCredentials"`
	ColorScheme       *string                              `json:"colorScheme"`
	Logger            interface{}                          `json:"logger"`
	RecordVideos      *WebKitBrowserNewPageRecordVideos    `json:"_recordVideos"`
}
type BrowserNewContextViewport struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type BrowserNewContextGeolocation struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Accuracy  *int     `json:"accuracy"`
}
type BrowserNewContextHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type BrowserNewContextRecordVideos struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type BrowserNewPageViewport struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type BrowserNewPageGeolocation struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Accuracy  *int     `json:"accuracy"`
}
type BrowserNewPageHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type BrowserNewPageRecordVideos struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type PageClickPosition struct {
	X *int `json:"x"`
	Y *int `json:"y"`
}
type PageDblclickPosition struct {
	X *int `json:"x"`
	Y *int `json:"y"`
}
type PageHoverPosition struct {
	X *int `json:"x"`
	Y *int `json:"y"`
}
type PagePdfMargin struct {
	Top    interface{} `json:"top"`
	Right  interface{} `json:"right"`
	Bottom interface{} `json:"bottom"`
	Left   interface{} `json:"left"`
}
type PageScreenshotClip struct {
	X      *int `json:"x"`
	Y      *int `json:"y"`
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type FrameClickPosition struct {
	X *int `json:"x"`
	Y *int `json:"y"`
}
type FrameDblclickPosition struct {
	X *int `json:"x"`
	Y *int `json:"y"`
}
type FrameHoverPosition struct {
	X *int `json:"x"`
	Y *int `json:"y"`
}
type ElementHandleClickPosition struct {
	X *int `json:"x"`
	Y *int `json:"y"`
}
type ElementHandleDblclickPosition struct {
	X *int `json:"x"`
	Y *int `json:"y"`
}
type ElementHandleHoverPosition struct {
	X *int `json:"x"`
	Y *int `json:"y"`
}
type BrowserTypeLaunchProxy struct {
	Server   *string `json:"server"`
	Bypass   *string `json:"bypass"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type BrowserTypeLaunchPersistentContextProxy struct {
	Server   *string `json:"server"`
	Bypass   *string `json:"bypass"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type BrowserTypeLaunchPersistentContextViewport struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type BrowserTypeLaunchPersistentContextGeolocation struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Accuracy  *int     `json:"accuracy"`
}
type BrowserTypeLaunchPersistentContextHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type BrowserTypeLaunchPersistentContextRecordVideos struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type BrowserTypeLaunchServerProxy struct {
	Server   *string `json:"server"`
	Bypass   *string `json:"bypass"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type ChromiumBrowserNewContextViewport struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type ChromiumBrowserNewContextGeolocation struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Accuracy  *int     `json:"accuracy"`
}
type ChromiumBrowserNewContextHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type ChromiumBrowserNewContextRecordVideos struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type ChromiumBrowserNewPageViewport struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type ChromiumBrowserNewPageGeolocation struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Accuracy  *int     `json:"accuracy"`
}
type ChromiumBrowserNewPageHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type ChromiumBrowserNewPageRecordVideos struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type FirefoxBrowserNewContextViewport struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type FirefoxBrowserNewContextGeolocation struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Accuracy  *int     `json:"accuracy"`
}
type FirefoxBrowserNewContextHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type FirefoxBrowserNewContextRecordVideos struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type FirefoxBrowserNewPageViewport struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type FirefoxBrowserNewPageGeolocation struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Accuracy  *int     `json:"accuracy"`
}
type FirefoxBrowserNewPageHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type FirefoxBrowserNewPageRecordVideos struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type WebKitBrowserNewContextViewport struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type WebKitBrowserNewContextGeolocation struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Accuracy  *int     `json:"accuracy"`
}
type WebKitBrowserNewContextHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type WebKitBrowserNewContextRecordVideos struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type WebKitBrowserNewPageViewport struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
type WebKitBrowserNewPageGeolocation struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Accuracy  *int     `json:"accuracy"`
}
type WebKitBrowserNewPageHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type WebKitBrowserNewPageRecordVideos struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}
