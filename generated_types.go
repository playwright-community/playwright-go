package playwright

type AccessibilitySnapshotOptions struct {
	InterestingOnly *bool          `json:"interestingOnly"`
	Root            *ElementHandle `json:"root"`
}
type BrowserNewContextOptions struct {
	AcceptDownloads   *bool                             `json:"acceptDownloads"`
	BypassCSP         *bool                             `json:"bypassCSP"`
	ColorScheme       *string                           `json:"colorScheme"`
	DeviceScaleFactor *float64                          `json:"deviceScaleFactor"`
	ExtraHTTPHeaders  map[string]string                 `json:"extraHTTPHeaders"`
	Geolocation       *BrowserNewContextGeolocation     `json:"geolocation"`
	HasTouch          *bool                             `json:"hasTouch"`
	HttpCredentials   *BrowserNewContextHttpCredentials `json:"httpCredentials"`
	IgnoreHTTPSErrors *bool                             `json:"ignoreHTTPSErrors"`
	IsMobile          *bool                             `json:"isMobile"`
	JavaScriptEnabled *bool                             `json:"javaScriptEnabled"`
	Locale            *string                           `json:"locale"`
	Logger            interface{}                       `json:"logger"`
	Offline           *bool                             `json:"offline"`
	Permissions       interface{}                       `json:"permissions"`
	Proxy             *BrowserNewContextProxy           `json:"proxy"`
	RecordHar         *BrowserNewContextRecordHar       `json:"recordHar"`
	RecordVideo       *BrowserNewContextRecordVideo     `json:"recordVideo"`
	StorageState      *BrowserNewContextStorageState    `json:"storageState"`
	TimezoneId        *string                           `json:"timezoneId"`
	UserAgent         *string                           `json:"userAgent"`
	VideoSize         *BrowserNewContextVideoSize       `json:"videoSize"`
	VideosPath        *string                           `json:"videosPath"`
	Viewport          *BrowserNewContextViewport        `json:"viewport"`
}
type BrowserNewPageOptions struct {
	AcceptDownloads   *bool                          `json:"acceptDownloads"`
	BypassCSP         *bool                          `json:"bypassCSP"`
	ColorScheme       *string                        `json:"colorScheme"`
	DeviceScaleFactor *float64                       `json:"deviceScaleFactor"`
	ExtraHTTPHeaders  map[string]string              `json:"extraHTTPHeaders"`
	Geolocation       *BrowserNewPageGeolocation     `json:"geolocation"`
	HasTouch          *bool                          `json:"hasTouch"`
	HttpCredentials   *BrowserNewPageHttpCredentials `json:"httpCredentials"`
	IgnoreHTTPSErrors *bool                          `json:"ignoreHTTPSErrors"`
	IsMobile          *bool                          `json:"isMobile"`
	JavaScriptEnabled *bool                          `json:"javaScriptEnabled"`
	Locale            *string                        `json:"locale"`
	Logger            interface{}                    `json:"logger"`
	Offline           *bool                          `json:"offline"`
	Permissions       interface{}                    `json:"permissions"`
	Proxy             *BrowserNewPageProxy           `json:"proxy"`
	RecordHar         *BrowserNewPageRecordHar       `json:"recordHar"`
	RecordVideo       *BrowserNewPageRecordVideo     `json:"recordVideo"`
	StorageState      *BrowserNewPageStorageState    `json:"storageState"`
	TimezoneId        *string                        `json:"timezoneId"`
	UserAgent         *string                        `json:"userAgent"`
	VideoSize         *BrowserNewPageVideoSize       `json:"videoSize"`
	VideosPath        *string                        `json:"videosPath"`
	Viewport          *BrowserNewPageViewport        `json:"viewport"`
}
type BrowserContextAddInitScriptOptions struct {
	Arg    interface{} `json:"arg"`
	Path   *string     `json:"path"`
	Script *string     `json:"script"`
}
type BrowserContextCookiesOptions struct {
	URLs interface{} `json:"urls"`
}
type BrowserContextExposeBindingOptions struct {
	Handle *bool `json:"handle"`
}
type BrowserContextGrantPermissionsOptions struct {
	Origin *string `json:"origin"`
}
type BrowserContextStorageStateOptions struct {
	Path *string `json:"path"`
}
type BrowserContextUnrouteOptions struct {
	Handler interface{} `json:"handler"`
}
type BrowserContextWaitForEventOptions struct {
	OptionsOrPredicate *BrowserContextWaitForEventOptionsOrPredicate `json:"optionsOrPredicate"`
}
type BrowserTypeLaunchOptions struct {
	Args              interface{}             `json:"args"`
	ChromiumSandbox   *bool                   `json:"chromiumSandbox"`
	Devtools          *bool                   `json:"devtools"`
	DownloadsPath     *string                 `json:"downloadsPath"`
	Env               interface{}             `json:"env"`
	ExecutablePath    *string                 `json:"executablePath"`
	FirefoxUserPrefs  interface{}             `json:"firefoxUserPrefs"`
	HandleSIGHUP      *bool                   `json:"handleSIGHUP"`
	HandleSIGINT      *bool                   `json:"handleSIGINT"`
	HandleSIGTERM     *bool                   `json:"handleSIGTERM"`
	Headless          *bool                   `json:"headless"`
	IgnoreDefaultArgs interface{}             `json:"ignoreDefaultArgs"`
	Logger            interface{}             `json:"logger"`
	Proxy             *BrowserTypeLaunchProxy `json:"proxy"`
	SlowMo            *float64                `json:"slowMo"`
	Timeout           *float64                `json:"timeout"`
}
type BrowserTypeLaunchPersistentContextOptions struct {
	AcceptDownloads   *bool                                              `json:"acceptDownloads"`
	Args              interface{}                                        `json:"args"`
	BypassCSP         *bool                                              `json:"bypassCSP"`
	ChromiumSandbox   *bool                                              `json:"chromiumSandbox"`
	ColorScheme       *string                                            `json:"colorScheme"`
	DeviceScaleFactor *float64                                           `json:"deviceScaleFactor"`
	Devtools          *bool                                              `json:"devtools"`
	DownloadsPath     *string                                            `json:"downloadsPath"`
	Env               interface{}                                        `json:"env"`
	ExecutablePath    *string                                            `json:"executablePath"`
	ExtraHTTPHeaders  map[string]string                                  `json:"extraHTTPHeaders"`
	Geolocation       *BrowserTypeLaunchPersistentContextGeolocation     `json:"geolocation"`
	HandleSIGHUP      *bool                                              `json:"handleSIGHUP"`
	HandleSIGINT      *bool                                              `json:"handleSIGINT"`
	HandleSIGTERM     *bool                                              `json:"handleSIGTERM"`
	HasTouch          *bool                                              `json:"hasTouch"`
	Headless          *bool                                              `json:"headless"`
	HttpCredentials   *BrowserTypeLaunchPersistentContextHttpCredentials `json:"httpCredentials"`
	IgnoreDefaultArgs interface{}                                        `json:"ignoreDefaultArgs"`
	IgnoreHTTPSErrors *bool                                              `json:"ignoreHTTPSErrors"`
	IsMobile          *bool                                              `json:"isMobile"`
	JavaScriptEnabled *bool                                              `json:"javaScriptEnabled"`
	Locale            *string                                            `json:"locale"`
	Logger            interface{}                                        `json:"logger"`
	Offline           *bool                                              `json:"offline"`
	Permissions       interface{}                                        `json:"permissions"`
	Proxy             *BrowserTypeLaunchPersistentContextProxy           `json:"proxy"`
	RecordHar         *BrowserTypeLaunchPersistentContextRecordHar       `json:"recordHar"`
	RecordVideo       *BrowserTypeLaunchPersistentContextRecordVideo     `json:"recordVideo"`
	SlowMo            *float64                                           `json:"slowMo"`
	Timeout           *float64                                           `json:"timeout"`
	TimezoneId        *string                                            `json:"timezoneId"`
	UserAgent         *string                                            `json:"userAgent"`
	VideoSize         *BrowserTypeLaunchPersistentContextVideoSize       `json:"videoSize"`
	VideosPath        *string                                            `json:"videosPath"`
	Viewport          *BrowserTypeLaunchPersistentContextViewport        `json:"viewport"`
}
type BrowserTypeLaunchServerOptions struct {
	Args              interface{}                   `json:"args"`
	ChromiumSandbox   *bool                         `json:"chromiumSandbox"`
	Devtools          *bool                         `json:"devtools"`
	DownloadsPath     *string                       `json:"downloadsPath"`
	Env               interface{}                   `json:"env"`
	ExecutablePath    *string                       `json:"executablePath"`
	FirefoxUserPrefs  interface{}                   `json:"firefoxUserPrefs"`
	HandleSIGHUP      *bool                         `json:"handleSIGHUP"`
	HandleSIGINT      *bool                         `json:"handleSIGINT"`
	HandleSIGTERM     *bool                         `json:"handleSIGTERM"`
	Headless          *bool                         `json:"headless"`
	IgnoreDefaultArgs interface{}                   `json:"ignoreDefaultArgs"`
	Logger            interface{}                   `json:"logger"`
	Port              *int                          `json:"port"`
	Proxy             *BrowserTypeLaunchServerProxy `json:"proxy"`
	Timeout           *float64                      `json:"timeout"`
}
type CDPSessionSendOptions struct {
	Params interface{} `json:"params"`
}
type ChromiumBrowserStartTracingOptions struct {
	Page        interface{} `json:"page"`
	Categories  interface{} `json:"categories"`
	Path        *string     `json:"path"`
	Screenshots *bool       `json:"screenshots"`
}
type ChromiumCoverageStartCSSCoverageOptions struct {
	ResetOnNavigation *bool `json:"resetOnNavigation"`
}
type ChromiumCoverageStartJSCoverageOptions struct {
	ReportAnonymousScripts *bool `json:"reportAnonymousScripts"`
	ResetOnNavigation      *bool `json:"resetOnNavigation"`
}
type DialogAcceptOptions struct {
	PromptText *string `json:"promptText"`
}
type ElementHandleEvalOnSelectorOptions struct {
	Arg       interface{} `json:"arg"`
	Forceexpr *bool       `json:"force_expr"`
}
type ElementHandleEvalOnSelectorAllOptions struct {
	Arg       interface{} `json:"arg"`
	Forceexpr *bool       `json:"force_expr"`
}
type ElementHandleCheckOptions struct {
	Force       *bool    `json:"force"`
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type ElementHandleClickOptions struct {
	Button      *string                     `json:"button"`
	ClickCount  *int                        `json:"clickCount"`
	Delay       *float64                    `json:"delay"`
	Force       *bool                       `json:"force"`
	Modifiers   interface{}                 `json:"modifiers"`
	NoWaitAfter *bool                       `json:"noWaitAfter"`
	Position    *ElementHandleClickPosition `json:"position"`
	Timeout     *float64                    `json:"timeout"`
}
type ElementHandleDblclickOptions struct {
	Button      *string                        `json:"button"`
	Delay       *float64                       `json:"delay"`
	Force       *bool                          `json:"force"`
	Modifiers   interface{}                    `json:"modifiers"`
	NoWaitAfter *bool                          `json:"noWaitAfter"`
	Position    *ElementHandleDblclickPosition `json:"position"`
	Timeout     *float64                       `json:"timeout"`
}
type ElementHandleDispatchEventOptions struct {
	EventInit interface{} `json:"eventInit"`
}
type ElementHandleFillOptions struct {
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type ElementHandleHoverOptions struct {
	Force     *bool                       `json:"force"`
	Modifiers interface{}                 `json:"modifiers"`
	Position  *ElementHandleHoverPosition `json:"position"`
	Timeout   *float64                    `json:"timeout"`
}
type ElementHandlePressOptions struct {
	Delay       *float64 `json:"delay"`
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type ElementHandleScreenshotOptions struct {
	OmitBackground *bool    `json:"omitBackground"`
	Path           *string  `json:"path"`
	Quality        *int     `json:"quality"`
	Timeout        *float64 `json:"timeout"`
	Type           *string  `json:"type"`
}
type ElementHandleScrollIntoViewIfNeededOptions struct {
	Timeout *float64 `json:"timeout"`
}
type ElementHandleSelectOptionOptions struct {
	NoWaitAfter *bool       `json:"noWaitAfter"`
	Timeout     *float64    `json:"timeout"`
	Element     interface{} `json:"element"`
	Index       interface{} `json:"index"`
	Value       interface{} `json:"value"`
	Label       interface{} `json:"label"`
}
type ElementHandleSelectTextOptions struct {
	Timeout *float64 `json:"timeout"`
}
type ElementHandleSetInputFilesOptions struct {
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type ElementHandleTapOptions struct {
	Force       *bool                     `json:"force"`
	Modifiers   interface{}               `json:"modifiers"`
	NoWaitAfter *bool                     `json:"noWaitAfter"`
	Position    *ElementHandleTapPosition `json:"position"`
	Timeout     *float64                  `json:"timeout"`
}
type ElementHandleTypeOptions struct {
	Delay       *float64 `json:"delay"`
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type ElementHandleUncheckOptions struct {
	Force       *bool    `json:"force"`
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type ElementHandleWaitForElementStateOptions struct {
	Timeout *float64 `json:"timeout"`
}
type ElementHandleWaitForSelectorOptions struct {
	State   *string  `json:"state"`
	Timeout *float64 `json:"timeout"`
}
type FileChooserSetFilesOptions struct {
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type FrameEvalOnSelectorOptions struct {
	Arg       interface{} `json:"arg"`
	Forceexpr *bool       `json:"force_expr"`
}
type FrameEvalOnSelectorAllOptions struct {
	Arg       interface{} `json:"arg"`
	Forceexpr *bool       `json:"force_expr"`
}
type FrameAddScriptTagOptions struct {
	Content *string `json:"content"`
	Path    *string `json:"path"`
	Type    *string `json:"type"`
	URL     *string `json:"url"`
}
type FrameAddStyleTagOptions struct {
	Content *string `json:"content"`
	Path    *string `json:"path"`
	URL     *string `json:"url"`
}
type FrameCheckOptions struct {
	Force       *bool    `json:"force"`
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type FrameClickOptions struct {
	Button      *string             `json:"button"`
	ClickCount  *int                `json:"clickCount"`
	Delay       *float64            `json:"delay"`
	Force       *bool               `json:"force"`
	Modifiers   interface{}         `json:"modifiers"`
	NoWaitAfter *bool               `json:"noWaitAfter"`
	Position    *FrameClickPosition `json:"position"`
	Timeout     *float64            `json:"timeout"`
}
type FrameDblclickOptions struct {
	Button      *string                `json:"button"`
	Delay       *float64               `json:"delay"`
	Force       *bool                  `json:"force"`
	Modifiers   interface{}            `json:"modifiers"`
	NoWaitAfter *bool                  `json:"noWaitAfter"`
	Position    *FrameDblclickPosition `json:"position"`
	Timeout     *float64               `json:"timeout"`
}
type FrameDispatchEventOptions struct {
	EventInit interface{} `json:"eventInit"`
	Timeout   *float64    `json:"timeout"`
}
type FrameEvaluateOptions struct {
	Arg       interface{} `json:"arg"`
	Forceexpr *bool       `json:"force_expr"`
}
type FrameEvaluateHandleOptions struct {
	Arg       interface{} `json:"arg"`
	Forceexpr *bool       `json:"force_expr"`
}
type FrameFillOptions struct {
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type FrameFocusOptions struct {
	Timeout *float64 `json:"timeout"`
}
type FrameGetAttributeOptions struct {
	Timeout *float64 `json:"timeout"`
}
type FrameGotoOptions struct {
	Referer   *string  `json:"referer"`
	Timeout   *float64 `json:"timeout"`
	WaitUntil *string  `json:"waitUntil"`
}
type FrameHoverOptions struct {
	Force     *bool               `json:"force"`
	Modifiers interface{}         `json:"modifiers"`
	Position  *FrameHoverPosition `json:"position"`
	Timeout   *float64            `json:"timeout"`
}
type FrameInnerHTMLOptions struct {
	Timeout *float64 `json:"timeout"`
}
type FrameInnerTextOptions struct {
	Timeout *float64 `json:"timeout"`
}
type FrameIsCheckedOptions struct {
	Timeout *float64 `json:"timeout"`
}
type FrameIsDisabledOptions struct {
	Timeout *float64 `json:"timeout"`
}
type FrameIsEditableOptions struct {
	Timeout *float64 `json:"timeout"`
}
type FrameIsEnabledOptions struct {
	Timeout *float64 `json:"timeout"`
}
type FrameIsHiddenOptions struct {
	Timeout *float64 `json:"timeout"`
}
type FrameIsVisibleOptions struct {
	Timeout *float64 `json:"timeout"`
}
type FramePressOptions struct {
	Delay       *float64 `json:"delay"`
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type FrameSelectOptionOptions struct {
	NoWaitAfter *bool       `json:"noWaitAfter"`
	Timeout     *float64    `json:"timeout"`
	Element     interface{} `json:"element"`
	Index       interface{} `json:"index"`
	Value       interface{} `json:"value"`
	Label       interface{} `json:"label"`
}
type FrameSetContentOptions struct {
	Timeout   *float64 `json:"timeout"`
	WaitUntil *string  `json:"waitUntil"`
}
type FrameSetInputFilesOptions struct {
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type FrameTapOptions struct {
	Force       *bool             `json:"force"`
	Modifiers   interface{}       `json:"modifiers"`
	NoWaitAfter *bool             `json:"noWaitAfter"`
	Position    *FrameTapPosition `json:"position"`
	Timeout     *float64          `json:"timeout"`
}
type FrameTextContentOptions struct {
	Timeout *float64 `json:"timeout"`
}
type FrameTypeOptions struct {
	Delay       *float64 `json:"delay"`
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type FrameUncheckOptions struct {
	Force       *bool    `json:"force"`
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type FrameWaitForFunctionOptions struct {
	Arg       interface{} `json:"arg"`
	Polling   interface{} `json:"polling"`
	Timeout   *float64    `json:"timeout"`
	Forceexpr *bool       `json:"force_expr"`
}
type FrameWaitForLoadStateOptions struct {
	State   *string  `json:"state"`
	Timeout *float64 `json:"timeout"`
}
type FrameWaitForNavigationOptions struct {
	Timeout   *float64    `json:"timeout"`
	URL       interface{} `json:"url"`
	WaitUntil *string     `json:"waitUntil"`
}
type FrameWaitForSelectorOptions struct {
	State   *string  `json:"state"`
	Timeout *float64 `json:"timeout"`
}
type JSHandleEvaluateOptions struct {
	Arg       interface{} `json:"arg"`
	Forceexpr *bool       `json:"force_expr"`
}
type JSHandleEvaluateHandleOptions struct {
	Arg       interface{} `json:"arg"`
	Forceexpr *bool       `json:"force_expr"`
}
type KeyboardPressOptions struct {
	Delay *float64 `json:"delay"`
}
type KeyboardTypeOptions struct {
	Delay *float64 `json:"delay"`
}
type MouseClickOptions struct {
	Button     *string  `json:"button"`
	ClickCount *int     `json:"clickCount"`
	Delay      *float64 `json:"delay"`
}
type MouseDblclickOptions struct {
	Button *string  `json:"button"`
	Delay  *float64 `json:"delay"`
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
type PageEvalOnSelectorOptions struct {
	Arg       interface{} `json:"arg"`
	Forceexpr *bool       `json:"force_expr"`
}
type PageEvalOnSelectorAllOptions struct {
	Arg       interface{} `json:"arg"`
	Forceexpr *bool       `json:"force_expr"`
}
type PageAddInitScriptOptions struct {
	Arg    interface{} `json:"arg"`
	Path   *string     `json:"path"`
	Script *string     `json:"script"`
}
type PageAddScriptTagOptions struct {
	Content *string `json:"content"`
	Path    *string `json:"path"`
	Type    *string `json:"type"`
	URL     *string `json:"url"`
}
type PageAddStyleTagOptions struct {
	Content *string `json:"content"`
	Path    *string `json:"path"`
	URL     *string `json:"url"`
}
type PageCheckOptions struct {
	Force       *bool    `json:"force"`
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type PageClickOptions struct {
	Button      *string            `json:"button"`
	ClickCount  *int               `json:"clickCount"`
	Delay       *float64           `json:"delay"`
	Force       *bool              `json:"force"`
	Modifiers   interface{}        `json:"modifiers"`
	NoWaitAfter *bool              `json:"noWaitAfter"`
	Position    *PageClickPosition `json:"position"`
	Timeout     *float64           `json:"timeout"`
}
type PageCloseOptions struct {
	RunBeforeUnload *bool `json:"runBeforeUnload"`
}
type PageDblclickOptions struct {
	Button      *string               `json:"button"`
	Delay       *float64              `json:"delay"`
	Force       *bool                 `json:"force"`
	Modifiers   interface{}           `json:"modifiers"`
	NoWaitAfter *bool                 `json:"noWaitAfter"`
	Position    *PageDblclickPosition `json:"position"`
	Timeout     *float64              `json:"timeout"`
}
type PageDispatchEventOptions struct {
	EventInit interface{} `json:"eventInit"`
	Timeout   *float64    `json:"timeout"`
}
type PageEmulateMediaOptions struct {
	Params      *PageEmulateMediaParams `json:"params"`
	Media       interface{}             `json:"media"`
	ColorScheme interface{}             `json:"colorScheme"`
}
type PageEvaluateOptions struct {
	Arg       interface{} `json:"arg"`
	Forceexpr *bool       `json:"force_expr"`
}
type PageEvaluateHandleOptions struct {
	Arg       interface{} `json:"arg"`
	Forceexpr *bool       `json:"force_expr"`
}
type PageExposeBindingOptions struct {
	Handle *bool `json:"handle"`
}
type PageFillOptions struct {
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type PageFocusOptions struct {
	Timeout *float64 `json:"timeout"`
}
type PageGetAttributeOptions struct {
	Timeout *float64 `json:"timeout"`
}
type PageGoBackOptions struct {
	Timeout   *float64 `json:"timeout"`
	WaitUntil *string  `json:"waitUntil"`
}
type PageGoForwardOptions struct {
	Timeout   *float64 `json:"timeout"`
	WaitUntil *string  `json:"waitUntil"`
}
type PageGotoOptions struct {
	Referer   *string  `json:"referer"`
	Timeout   *float64 `json:"timeout"`
	WaitUntil *string  `json:"waitUntil"`
}
type PageHoverOptions struct {
	Force     *bool              `json:"force"`
	Modifiers interface{}        `json:"modifiers"`
	Position  *PageHoverPosition `json:"position"`
	Timeout   *float64           `json:"timeout"`
}
type PageInnerHTMLOptions struct {
	Timeout *float64 `json:"timeout"`
}
type PageInnerTextOptions struct {
	Timeout *float64 `json:"timeout"`
}
type PageIsCheckedOptions struct {
	Timeout *float64 `json:"timeout"`
}
type PageIsDisabledOptions struct {
	Timeout *float64 `json:"timeout"`
}
type PageIsEditableOptions struct {
	Timeout *float64 `json:"timeout"`
}
type PageIsEnabledOptions struct {
	Timeout *float64 `json:"timeout"`
}
type PageIsHiddenOptions struct {
	Timeout *float64 `json:"timeout"`
}
type PageIsVisibleOptions struct {
	Timeout *float64 `json:"timeout"`
}
type PagePDFOptions struct {
	DisplayHeaderFooter *bool          `json:"displayHeaderFooter"`
	FooterTemplate      *string        `json:"footerTemplate"`
	Format              *string        `json:"format"`
	HeaderTemplate      *string        `json:"headerTemplate"`
	Height              interface{}    `json:"height"`
	Landscape           *bool          `json:"landscape"`
	Margin              *PagePDFMargin `json:"margin"`
	PageRanges          *string        `json:"pageRanges"`
	Path                *string        `json:"path"`
	PreferCSSPageSize   *bool          `json:"preferCSSPageSize"`
	PrintBackground     *bool          `json:"printBackground"`
	Scale               *float64       `json:"scale"`
	Width               interface{}    `json:"width"`
}
type PagePressOptions struct {
	Delay       *float64 `json:"delay"`
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type PageReloadOptions struct {
	Timeout   *float64 `json:"timeout"`
	WaitUntil *string  `json:"waitUntil"`
}
type PageScreenshotOptions struct {
	Clip           *PageScreenshotClip `json:"clip"`
	FullPage       *bool               `json:"fullPage"`
	OmitBackground *bool               `json:"omitBackground"`
	Path           *string             `json:"path"`
	Quality        *int                `json:"quality"`
	Timeout        *float64            `json:"timeout"`
	Type           *string             `json:"type"`
}
type PageSelectOptionOptions struct {
	NoWaitAfter *bool       `json:"noWaitAfter"`
	Timeout     *float64    `json:"timeout"`
	Element     interface{} `json:"element"`
	Index       interface{} `json:"index"`
	Value       interface{} `json:"value"`
	Label       interface{} `json:"label"`
}
type PageSetContentOptions struct {
	Timeout   *float64 `json:"timeout"`
	WaitUntil *string  `json:"waitUntil"`
}
type PageSetInputFilesOptions struct {
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type PageTapOptions struct {
	Force       *bool            `json:"force"`
	Modifiers   interface{}      `json:"modifiers"`
	NoWaitAfter *bool            `json:"noWaitAfter"`
	Position    *PageTapPosition `json:"position"`
	Timeout     *float64         `json:"timeout"`
}
type PageTextContentOptions struct {
	Timeout *float64 `json:"timeout"`
}
type PageTypeOptions struct {
	Delay       *float64 `json:"delay"`
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type PageUncheckOptions struct {
	Force       *bool    `json:"force"`
	NoWaitAfter *bool    `json:"noWaitAfter"`
	Timeout     *float64 `json:"timeout"`
}
type PageUnrouteOptions struct {
	Handler interface{} `json:"handler"`
}
type PageWaitForEventOptions struct {
	OptionsOrPredicate *PageWaitForEventOptionsOrPredicate `json:"optionsOrPredicate"`
}
type PageWaitForFunctionOptions struct {
	Arg       interface{} `json:"arg"`
	Polling   interface{} `json:"polling"`
	Timeout   *float64    `json:"timeout"`
	Forceexpr *bool       `json:"force_expr"`
}
type PageWaitForLoadStateOptions struct {
	State   *string  `json:"state"`
	Timeout *float64 `json:"timeout"`
}
type PageWaitForNavigationOptions struct {
	Timeout   *float64    `json:"timeout"`
	URL       interface{} `json:"url"`
	WaitUntil *string     `json:"waitUntil"`
}
type PageWaitForRequestOptions struct {
	Timeout *float64 `json:"timeout"`
}
type PageWaitForResponseOptions struct {
	Timeout *float64 `json:"timeout"`
}
type PageWaitForSelectorOptions struct {
	State   *string  `json:"state"`
	Timeout *float64 `json:"timeout"`
}
type RouteAbortOptions struct {
	ErrorCode *string `json:"errorCode"`
}
type RouteContinueOptions struct {
	Headers  map[string]string `json:"headers"`
	Method   *string           `json:"method"`
	PostData interface{}       `json:"postData"`
	URL      *string           `json:"url"`
}
type RouteFulfillOptions struct {
	Body        interface{}       `json:"body"`
	ContentType *string           `json:"contentType"`
	Headers     map[string]string `json:"headers"`
	Path        *string           `json:"path"`
	Status      *int              `json:"status"`
}
type SelectorsRegisterOptions struct {
	ContentScript *bool `json:"contentScript"`
}
type WebSocketWaitForEventOptions struct {
	OptionsOrPredicate *WebSocketWaitForEventOptionsOrPredicate `json:"optionsOrPredicate"`
}
type WorkerEvaluateOptions struct {
	Arg       interface{} `json:"arg"`
	Forceexpr *bool       `json:"force_expr"`
}
type WorkerEvaluateHandleOptions struct {
	Arg       interface{} `json:"arg"`
	Forceexpr *bool       `json:"force_expr"`
}
type BrowserNewContextGeolocation struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Accuracy  *float64 `json:"accuracy"`
}

type BrowserNewContextHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type BrowserNewContextProxy struct {
	Server   *string `json:"server"`
	Bypass   *string `json:"bypass"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type BrowserNewContextRecordHar struct {
	OmitContent *bool   `json:"omitContent"`
	Path        *string `json:"path"`
}

type RecordVideoSize struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}

type BrowserNewContextRecordVideo struct {
	Dir  *string          `json:"dir"`
	Size *RecordVideoSize `json:"size"`
}

type BrowserNewContextStorageState struct {
	Cookies interface{} `json:"cookies"`
	Origins interface{} `json:"origins"`
}

type BrowserNewContextVideoSize struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}

type BrowserNewContextViewport struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}

type BrowserNewPageGeolocation struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Accuracy  *float64 `json:"accuracy"`
}

type BrowserNewPageHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type BrowserNewPageProxy struct {
	Server   *string `json:"server"`
	Bypass   *string `json:"bypass"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type BrowserNewPageRecordHar struct {
	OmitContent *bool   `json:"omitContent"`
	Path        *string `json:"path"`
}

type BrowserNewPageRecordVideo struct {
	Dir  *string          `json:"dir"`
	Size *RecordVideoSize `json:"size"`
}

type BrowserNewPageStorageState struct {
	Cookies interface{} `json:"cookies"`
	Origins interface{} `json:"origins"`
}

type BrowserNewPageVideoSize struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}

type BrowserNewPageViewport struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}

type BrowserContextWaitForEventOptionsOrPredicate struct {
	Predicate interface{} `json:"predicate"`
	Timeout   *float64    `json:"timeout"`
}

type BrowserTypeLaunchProxy struct {
	Server   *string `json:"server"`
	Bypass   *string `json:"bypass"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type BrowserTypeLaunchPersistentContextGeolocation struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Accuracy  *float64 `json:"accuracy"`
}

type BrowserTypeLaunchPersistentContextHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type BrowserTypeLaunchPersistentContextProxy struct {
	Server   *string `json:"server"`
	Bypass   *string `json:"bypass"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type BrowserTypeLaunchPersistentContextRecordHar struct {
	OmitContent *bool   `json:"omitContent"`
	Path        *string `json:"path"`
}

type BrowserTypeLaunchPersistentContextRecordVideo struct {
	Dir  *string          `json:"dir"`
	Size *RecordVideoSize `json:"size"`
}

type BrowserTypeLaunchPersistentContextVideoSize struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}

type BrowserTypeLaunchPersistentContextViewport struct {
	Width  *int `json:"width"`
	Height *int `json:"height"`
}

type BrowserTypeLaunchServerProxy struct {
	Server   *string `json:"server"`
	Bypass   *string `json:"bypass"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type ElementHandleClickPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}

type ElementHandleDblclickPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}

type ElementHandleHoverPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}

type ElementHandleTapPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}

type FrameClickPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}

type FrameDblclickPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}

type FrameHoverPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}

type FrameTapPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}

type PageClickPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}

type PageDblclickPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}

type PageEmulateMediaParams struct {
	Media       interface{} `json:"media"`
	ColorScheme interface{} `json:"colorScheme"`
}

type PageHoverPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}

type PagePDFMargin struct {
	Top    interface{} `json:"top"`
	Right  interface{} `json:"right"`
	Bottom interface{} `json:"bottom"`
	Left   interface{} `json:"left"`
}

type PageScreenshotClip struct {
	X      *float64 `json:"x"`
	Y      *float64 `json:"y"`
	Width  *float64 `json:"width"`
	Height *float64 `json:"height"`
}

type PageTapPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}

type PageWaitForEventOptionsOrPredicate struct {
	Predicate interface{} `json:"predicate"`
	Timeout   *float64    `json:"timeout"`
}

type WebSocketWaitForEventOptionsOrPredicate struct {
	Predicate interface{} `json:"predicate"`
	Timeout   *float64    `json:"timeout"`
}
