package playwright

type BrowserNewContextOptions struct {
	AcceptDownloads   *bool `json:"acceptDownloads,omitempty"`
	IgnoreHTTPSErrors *bool `json:"ignoreHTTPSErrors,omitempty"`
	BypassCSP         *bool `json:"bypassCSP,omitempty"`
	Viewport          *struct {
		Width  *int `json:"width,omitempty"`
		Height *int `json:"height,omitempty"`
	} `json:"viewport,omitempty"`
	UserAgent         *string `json:"userAgent,omitempty"`
	DeviceScaleFactor *int    `json:"deviceScaleFactor,omitempty"`
	IsMobile          *bool   `json:"isMobile,omitempty"`
	HasTouch          *bool   `json:"hasTouch,omitempty"`
	JavaScriptEnabled *bool   `json:"javaScriptEnabled,omitempty"`
	TimezoneId        *string `json:"timezoneId,omitempty"`
	Geolocation       *struct {
		Latitude  *int `json:"latitude,omitempty"`
		Longitude *int `json:"longitude,omitempty"`
		Accuracy  *int `json:"accuracy,omitempty"`
	} `json:"geolocation,omitempty"`
	Locale           *string           `json:"locale,omitempty"`
	Permissions      []string          `json:"permissions,omitempty"`
	ExtraHTTPHeaders map[string]string `json:"extraHTTPHeaders,omitempty"`
	Offline          *bool             `json:"offline,omitempty"`
	HttpCredentials  *struct {
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
	} `json:"httpCredentials,omitempty"`
	ColorScheme *string     `json:"colorScheme,omitempty"`
	Logger      interface{} `json:"logger,omitempty"`
}
type BrowserNewPageOptions struct {
	AcceptDownloads   *bool `json:"acceptDownloads,omitempty"`
	IgnoreHTTPSErrors *bool `json:"ignoreHTTPSErrors,omitempty"`
	BypassCSP         *bool `json:"bypassCSP,omitempty"`
	Viewport          *struct {
		Width  *int `json:"width,omitempty"`
		Height *int `json:"height,omitempty"`
	} `json:"viewport,omitempty"`
	UserAgent         *string `json:"userAgent,omitempty"`
	DeviceScaleFactor *int    `json:"deviceScaleFactor,omitempty"`
	IsMobile          *bool   `json:"isMobile,omitempty"`
	HasTouch          *bool   `json:"hasTouch,omitempty"`
	JavaScriptEnabled *bool   `json:"javaScriptEnabled,omitempty"`
	TimezoneId        *string `json:"timezoneId,omitempty"`
	Geolocation       *struct {
		Latitude  *int `json:"latitude,omitempty"`
		Longitude *int `json:"longitude,omitempty"`
		Accuracy  *int `json:"accuracy,omitempty"`
	} `json:"geolocation,omitempty"`
	Locale           *string           `json:"locale,omitempty"`
	Permissions      []string          `json:"permissions,omitempty"`
	ExtraHTTPHeaders map[string]string `json:"extraHTTPHeaders,omitempty"`
	Offline          *bool             `json:"offline,omitempty"`
	HttpCredentials  *struct {
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
	} `json:"httpCredentials,omitempty"`
	ColorScheme *string     `json:"colorScheme,omitempty"`
	Logger      interface{} `json:"logger,omitempty"`
}
type BrowserContextGrantPermissionsOptions struct {
	Origin *string `json:"origin,omitempty"`
}
type PageCloseOptions struct {
	RunBeforeUnload *bool `json:"runBeforeUnload,omitempty"`
}
type PageAddScriptTagOptions struct {
	Url     *string `json:"url,omitempty"`
	Path    *string `json:"path,omitempty"`
	Content *string `json:"content,omitempty"`
	Type    *string `json:"type,omitempty"`
}
type PageAddStyleTagOptions struct {
	Url     *string `json:"url,omitempty"`
	Path    *string `json:"path,omitempty"`
	Content *string `json:"content,omitempty"`
}
type PageCheckOptions struct {
	Force       *bool `json:"force,omitempty"`
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type PageClickOptions struct {
	Button     *string `json:"button,omitempty"`
	ClickCount *int    `json:"clickCount,omitempty"`
	Delay      *int    `json:"delay,omitempty"`
	Position   *struct {
		X *int `json:"x,omitempty"`
		Y *int `json:"y,omitempty"`
	} `json:"position,omitempty"`
	Modifiers   interface{} `json:"modifiers,omitempty"`
	Force       *bool       `json:"force,omitempty"`
	NoWaitAfter *bool       `json:"noWaitAfter,omitempty"`
	Timeout     *int        `json:"timeout,omitempty"`
}
type PageDblclickOptions struct {
	Button   *string `json:"button,omitempty"`
	Delay    *int    `json:"delay,omitempty"`
	Position *struct {
		X *int `json:"x,omitempty"`
		Y *int `json:"y,omitempty"`
	} `json:"position,omitempty"`
	Modifiers   interface{} `json:"modifiers,omitempty"`
	Force       *bool       `json:"force,omitempty"`
	NoWaitAfter *bool       `json:"noWaitAfter,omitempty"`
	Timeout     *int        `json:"timeout,omitempty"`
}
type PageDispatchEventOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type PageEmulateMediaOptions struct {
	Media       interface{} `json:"media,omitempty"`
	ColorScheme interface{} `json:"colorScheme,omitempty"`
}
type PageFillOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type PageFocusOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type PageFrameOptions struct {
	Name *string     `json:"name,omitempty"`
	Url  interface{} `json:"url,omitempty"`
}
type PageGetAttributeOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type PageGoBackOptions struct {
	Timeout   *int    `json:"timeout,omitempty"`
	WaitUntil *string `json:"waitUntil,omitempty"`
}
type PageGoForwardOptions struct {
	Timeout   *int    `json:"timeout,omitempty"`
	WaitUntil *string `json:"waitUntil,omitempty"`
}
type PageGotoOptions struct {
	Timeout   *int    `json:"timeout,omitempty"`
	WaitUntil *string `json:"waitUntil,omitempty"`
	Referer   *string `json:"referer,omitempty"`
}
type PageHoverOptions struct {
	Position *struct {
		X *int `json:"x,omitempty"`
		Y *int `json:"y,omitempty"`
	} `json:"position,omitempty"`
	Modifiers interface{} `json:"modifiers,omitempty"`
	Force     *bool       `json:"force,omitempty"`
	Timeout   *int        `json:"timeout,omitempty"`
}
type PageInnerHTMLOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type PageInnerTextOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type PagePdfOptions struct {
	Path                *string     `json:"path,omitempty"`
	Scale               *int        `json:"scale,omitempty"`
	DisplayHeaderFooter *bool       `json:"displayHeaderFooter,omitempty"`
	HeaderTemplate      *string     `json:"headerTemplate,omitempty"`
	FooterTemplate      *string     `json:"footerTemplate,omitempty"`
	PrintBackground     *bool       `json:"printBackground,omitempty"`
	Landscape           *bool       `json:"landscape,omitempty"`
	PageRanges          *string     `json:"pageRanges,omitempty"`
	Format              *string     `json:"format,omitempty"`
	Width               interface{} `json:"width,omitempty"`
	Height              interface{} `json:"height,omitempty"`
	Margin              *struct {
		Top    interface{} `json:"top,omitempty"`
		Right  interface{} `json:"right,omitempty"`
		Bottom interface{} `json:"bottom,omitempty"`
		Left   interface{} `json:"left,omitempty"`
	} `json:"margin,omitempty"`
	PreferCSSPageSize *bool `json:"preferCSSPageSize,omitempty"`
}
type PagePressOptions struct {
	Delay       *int  `json:"delay,omitempty"`
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type PageReloadOptions struct {
	Timeout   *int    `json:"timeout,omitempty"`
	WaitUntil *string `json:"waitUntil,omitempty"`
}
type PageScreenshotOptions struct {
	Path     *string `json:"path,omitempty"`
	Type     *string `json:"type,omitempty"`
	Quality  *int    `json:"quality,omitempty"`
	FullPage *bool   `json:"fullPage,omitempty"`
	Clip     *struct {
		X      *int `json:"x,omitempty"`
		Y      *int `json:"y,omitempty"`
		Width  *int `json:"width,omitempty"`
		Height *int `json:"height,omitempty"`
	} `json:"clip,omitempty"`
	OmitBackground *bool `json:"omitBackground,omitempty"`
	Timeout        *int  `json:"timeout,omitempty"`
}
type PageSelectOptionOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type PageSetContentOptions struct {
	Timeout   *int    `json:"timeout,omitempty"`
	WaitUntil *string `json:"waitUntil,omitempty"`
}
type PageSetInputFilesOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type PageTextContentOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type PageTypeOptions struct {
	Delay       *int  `json:"delay,omitempty"`
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type PageUncheckOptions struct {
	Force       *bool `json:"force,omitempty"`
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type PageWaitForFunctionOptions struct {
	Polling interface{} `json:"polling,omitempty"`
	Timeout *int        `json:"timeout,omitempty"`
}
type PageWaitForLoadStateOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type PageWaitForNavigationOptions struct {
	Timeout   *int        `json:"timeout,omitempty"`
	Url       interface{} `json:"url,omitempty"`
	WaitUntil *string     `json:"waitUntil,omitempty"`
}
type PageWaitForRequestOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type PageWaitForResponseOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type PageWaitForSelectorOptions struct {
	State   *string `json:"state,omitempty"`
	Timeout *int    `json:"timeout,omitempty"`
}
type FrameAddScriptTagOptions struct {
	Url     *string `json:"url,omitempty"`
	Path    *string `json:"path,omitempty"`
	Content *string `json:"content,omitempty"`
	Type    *string `json:"type,omitempty"`
}
type FrameAddStyleTagOptions struct {
	Url     *string `json:"url,omitempty"`
	Path    *string `json:"path,omitempty"`
	Content *string `json:"content,omitempty"`
}
type FrameCheckOptions struct {
	Force       *bool `json:"force,omitempty"`
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type FrameClickOptions struct {
	Button     *string `json:"button,omitempty"`
	ClickCount *int    `json:"clickCount,omitempty"`
	Delay      *int    `json:"delay,omitempty"`
	Position   *struct {
		X *int `json:"x,omitempty"`
		Y *int `json:"y,omitempty"`
	} `json:"position,omitempty"`
	Modifiers   interface{} `json:"modifiers,omitempty"`
	Force       *bool       `json:"force,omitempty"`
	NoWaitAfter *bool       `json:"noWaitAfter,omitempty"`
	Timeout     *int        `json:"timeout,omitempty"`
}
type FrameDblclickOptions struct {
	Button   *string `json:"button,omitempty"`
	Delay    *int    `json:"delay,omitempty"`
	Position *struct {
		X *int `json:"x,omitempty"`
		Y *int `json:"y,omitempty"`
	} `json:"position,omitempty"`
	Modifiers   interface{} `json:"modifiers,omitempty"`
	Force       *bool       `json:"force,omitempty"`
	NoWaitAfter *bool       `json:"noWaitAfter,omitempty"`
	Timeout     *int        `json:"timeout,omitempty"`
}
type FrameDispatchEventOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type FrameFillOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type FrameFocusOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type FrameGetAttributeOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type FrameGotoOptions struct {
	Timeout   *int    `json:"timeout,omitempty"`
	WaitUntil *string `json:"waitUntil,omitempty"`
	Referer   *string `json:"referer,omitempty"`
}
type FrameHoverOptions struct {
	Position *struct {
		X *int `json:"x,omitempty"`
		Y *int `json:"y,omitempty"`
	} `json:"position,omitempty"`
	Modifiers interface{} `json:"modifiers,omitempty"`
	Force     *bool       `json:"force,omitempty"`
	Timeout   *int        `json:"timeout,omitempty"`
}
type FrameInnerHTMLOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type FrameInnerTextOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type FramePressOptions struct {
	Delay       *int  `json:"delay,omitempty"`
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type FrameSelectOptionOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type FrameSetContentOptions struct {
	Timeout   *int    `json:"timeout,omitempty"`
	WaitUntil *string `json:"waitUntil,omitempty"`
}
type FrameSetInputFilesOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type FrameTextContentOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type FrameTypeOptions struct {
	Delay       *int  `json:"delay,omitempty"`
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type FrameUncheckOptions struct {
	Force       *bool `json:"force,omitempty"`
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type FrameWaitForFunctionOptions struct {
	Polling interface{} `json:"polling,omitempty"`
	Timeout *int        `json:"timeout,omitempty"`
}
type FrameWaitForLoadStateOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type FrameWaitForNavigationOptions struct {
	Timeout   *int        `json:"timeout,omitempty"`
	Url       interface{} `json:"url,omitempty"`
	WaitUntil *string     `json:"waitUntil,omitempty"`
}
type FrameWaitForSelectorOptions struct {
	State   *string `json:"state,omitempty"`
	Timeout *int    `json:"timeout,omitempty"`
}
type ElementHandleCheckOptions struct {
	Force       *bool `json:"force,omitempty"`
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type ElementHandleClickOptions struct {
	Button     *string `json:"button,omitempty"`
	ClickCount *int    `json:"clickCount,omitempty"`
	Delay      *int    `json:"delay,omitempty"`
	Position   *struct {
		X *int `json:"x,omitempty"`
		Y *int `json:"y,omitempty"`
	} `json:"position,omitempty"`
	Modifiers   interface{} `json:"modifiers,omitempty"`
	Force       *bool       `json:"force,omitempty"`
	NoWaitAfter *bool       `json:"noWaitAfter,omitempty"`
	Timeout     *int        `json:"timeout,omitempty"`
}
type ElementHandleDblclickOptions struct {
	Button   *string `json:"button,omitempty"`
	Delay    *int    `json:"delay,omitempty"`
	Position *struct {
		X *int `json:"x,omitempty"`
		Y *int `json:"y,omitempty"`
	} `json:"position,omitempty"`
	Modifiers   interface{} `json:"modifiers,omitempty"`
	Force       *bool       `json:"force,omitempty"`
	NoWaitAfter *bool       `json:"noWaitAfter,omitempty"`
	Timeout     *int        `json:"timeout,omitempty"`
}
type ElementHandleFillOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type ElementHandleHoverOptions struct {
	Position *struct {
		X *int `json:"x,omitempty"`
		Y *int `json:"y,omitempty"`
	} `json:"position,omitempty"`
	Modifiers interface{} `json:"modifiers,omitempty"`
	Force     *bool       `json:"force,omitempty"`
	Timeout   *int        `json:"timeout,omitempty"`
}
type ElementHandlePressOptions struct {
	Delay       *int  `json:"delay,omitempty"`
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type ElementHandleScreenshotOptions struct {
	Path           *string `json:"path,omitempty"`
	Type           *string `json:"type,omitempty"`
	Quality        *int    `json:"quality,omitempty"`
	OmitBackground *bool   `json:"omitBackground,omitempty"`
	Timeout        *int    `json:"timeout,omitempty"`
}
type ElementHandleScrollIntoViewIfNeededOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type ElementHandleSelectOptionOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type ElementHandleSelectTextOptions struct {
	Timeout *int `json:"timeout,omitempty"`
}
type ElementHandleSetInputFilesOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type ElementHandleTypeOptions struct {
	Delay       *int  `json:"delay,omitempty"`
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type ElementHandleUncheckOptions struct {
	Force       *bool `json:"force,omitempty"`
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type FileChooserSetFilesOptions struct {
	NoWaitAfter *bool `json:"noWaitAfter,omitempty"`
	Timeout     *int  `json:"timeout,omitempty"`
}
type KeyboardPressOptions struct {
	Delay *int `json:"delay,omitempty"`
}
type KeyboardTypeOptions struct {
	Delay *int `json:"delay,omitempty"`
}
type MouseClickOptions struct {
	Button     *string `json:"button,omitempty"`
	ClickCount *int    `json:"clickCount,omitempty"`
	Delay      *int    `json:"delay,omitempty"`
}
type MouseDblclickOptions struct {
	Button *string `json:"button,omitempty"`
	Delay  *int    `json:"delay,omitempty"`
}
type MouseDownOptions struct {
	Button     *string `json:"button,omitempty"`
	ClickCount *int    `json:"clickCount,omitempty"`
}
type MouseMoveOptions struct {
	Steps *int `json:"steps,omitempty"`
}
type MouseUpOptions struct {
	Button     *string `json:"button,omitempty"`
	ClickCount *int    `json:"clickCount,omitempty"`
}
type SelectorsRegisterOptions struct {
	ContentScript *bool `json:"contentScript,omitempty"`
}
type AccessibilitySnapshotOptions struct {
	InterestingOnly *bool          `json:"interestingOnly,omitempty"`
	Root            *ElementHandle `json:"root,omitempty"`
}
type BrowserTypeConnectOptions struct {
	WsEndpoint *string     `json:"wsEndpoint,omitempty"`
	SlowMo     *int        `json:"slowMo,omitempty"`
	Logger     interface{} `json:"logger,omitempty"`
	Timeout    *int        `json:"timeout,omitempty"`
}
type BrowserTypeLaunchOptions struct {
	Headless          *bool       `json:"headless,omitempty"`
	ExecutablePath    *string     `json:"executablePath,omitempty"`
	Args              []string    `json:"args,omitempty"`
	IgnoreDefaultArgs interface{} `json:"ignoreDefaultArgs,omitempty"`
	Proxy             *struct {
		Server   *string `json:"server,omitempty"`
		Bypass   *string `json:"bypass,omitempty"`
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
	} `json:"proxy,omitempty"`
	DownloadsPath    *string `json:"downloadsPath,omitempty"`
	ChromiumSandbox  *bool   `json:"chromiumSandbox,omitempty"`
	FirefoxUserPrefs *struct {
	} `json:"firefoxUserPrefs,omitempty"`
	HandleSIGINT  *bool                  `json:"handleSIGINT,omitempty"`
	HandleSIGTERM *bool                  `json:"handleSIGTERM,omitempty"`
	HandleSIGHUP  *bool                  `json:"handleSIGHUP,omitempty"`
	Logger        interface{}            `json:"logger,omitempty"`
	Timeout       *int                   `json:"timeout,omitempty"`
	Env           map[string]interface{} `json:"env,omitempty"`
	Devtools      *bool                  `json:"devtools,omitempty"`
	SlowMo        *int                   `json:"slowMo,omitempty"`
}
type BrowserTypeLaunchPersistentContextOptions struct {
	Headless          *bool       `json:"headless,omitempty"`
	ExecutablePath    *string     `json:"executablePath,omitempty"`
	Args              []string    `json:"args,omitempty"`
	IgnoreDefaultArgs interface{} `json:"ignoreDefaultArgs,omitempty"`
	Proxy             *struct {
		Server   *string `json:"server,omitempty"`
		Bypass   *string `json:"bypass,omitempty"`
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
	} `json:"proxy,omitempty"`
	AcceptDownloads   *bool                  `json:"acceptDownloads,omitempty"`
	DownloadsPath     *string                `json:"downloadsPath,omitempty"`
	ChromiumSandbox   *bool                  `json:"chromiumSandbox,omitempty"`
	HandleSIGINT      *bool                  `json:"handleSIGINT,omitempty"`
	HandleSIGTERM     *bool                  `json:"handleSIGTERM,omitempty"`
	HandleSIGHUP      *bool                  `json:"handleSIGHUP,omitempty"`
	Logger            interface{}            `json:"logger,omitempty"`
	Timeout           *int                   `json:"timeout,omitempty"`
	Env               map[string]interface{} `json:"env,omitempty"`
	Devtools          *bool                  `json:"devtools,omitempty"`
	SlowMo            *int                   `json:"slowMo,omitempty"`
	IgnoreHTTPSErrors *bool                  `json:"ignoreHTTPSErrors,omitempty"`
	BypassCSP         *bool                  `json:"bypassCSP,omitempty"`
	Viewport          *struct {
		Width  *int `json:"width,omitempty"`
		Height *int `json:"height,omitempty"`
	} `json:"viewport,omitempty"`
	UserAgent         *string `json:"userAgent,omitempty"`
	DeviceScaleFactor *int    `json:"deviceScaleFactor,omitempty"`
	IsMobile          *bool   `json:"isMobile,omitempty"`
	HasTouch          *bool   `json:"hasTouch,omitempty"`
	JavaScriptEnabled *bool   `json:"javaScriptEnabled,omitempty"`
	TimezoneId        *string `json:"timezoneId,omitempty"`
	Geolocation       *struct {
		Latitude  *int `json:"latitude,omitempty"`
		Longitude *int `json:"longitude,omitempty"`
		Accuracy  *int `json:"accuracy,omitempty"`
	} `json:"geolocation,omitempty"`
	Locale           *string           `json:"locale,omitempty"`
	Permissions      []string          `json:"permissions,omitempty"`
	ExtraHTTPHeaders map[string]string `json:"extraHTTPHeaders,omitempty"`
	Offline          *bool             `json:"offline,omitempty"`
	HttpCredentials  *struct {
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
	} `json:"httpCredentials,omitempty"`
	ColorScheme *string `json:"colorScheme,omitempty"`
}
type BrowserTypeLaunchServerOptions struct {
	Headless          *bool       `json:"headless,omitempty"`
	Port              *int        `json:"port,omitempty"`
	ExecutablePath    *string     `json:"executablePath,omitempty"`
	Args              []string    `json:"args,omitempty"`
	IgnoreDefaultArgs interface{} `json:"ignoreDefaultArgs,omitempty"`
	Proxy             *struct {
		Server   *string `json:"server,omitempty"`
		Bypass   *string `json:"bypass,omitempty"`
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
	} `json:"proxy,omitempty"`
	DownloadsPath    *string `json:"downloadsPath,omitempty"`
	ChromiumSandbox  *bool   `json:"chromiumSandbox,omitempty"`
	FirefoxUserPrefs *struct {
	} `json:"firefoxUserPrefs,omitempty"`
	HandleSIGINT  *bool                  `json:"handleSIGINT,omitempty"`
	HandleSIGTERM *bool                  `json:"handleSIGTERM,omitempty"`
	HandleSIGHUP  *bool                  `json:"handleSIGHUP,omitempty"`
	Logger        interface{}            `json:"logger,omitempty"`
	Timeout       *int                   `json:"timeout,omitempty"`
	Env           map[string]interface{} `json:"env,omitempty"`
	Devtools      *bool                  `json:"devtools,omitempty"`
}
type ChromiumBrowserStartTracingOptions struct {
	Path        *string  `json:"path,omitempty"`
	Screenshots *bool    `json:"screenshots,omitempty"`
	Categories  []string `json:"categories,omitempty"`
}
type ChromiumBrowserNewContextOptions struct {
	AcceptDownloads   *bool `json:"acceptDownloads,omitempty"`
	IgnoreHTTPSErrors *bool `json:"ignoreHTTPSErrors,omitempty"`
	BypassCSP         *bool `json:"bypassCSP,omitempty"`
	Viewport          *struct {
		Width  *int `json:"width,omitempty"`
		Height *int `json:"height,omitempty"`
	} `json:"viewport,omitempty"`
	UserAgent         *string `json:"userAgent,omitempty"`
	DeviceScaleFactor *int    `json:"deviceScaleFactor,omitempty"`
	IsMobile          *bool   `json:"isMobile,omitempty"`
	HasTouch          *bool   `json:"hasTouch,omitempty"`
	JavaScriptEnabled *bool   `json:"javaScriptEnabled,omitempty"`
	TimezoneId        *string `json:"timezoneId,omitempty"`
	Geolocation       *struct {
		Latitude  *int `json:"latitude,omitempty"`
		Longitude *int `json:"longitude,omitempty"`
		Accuracy  *int `json:"accuracy,omitempty"`
	} `json:"geolocation,omitempty"`
	Locale           *string           `json:"locale,omitempty"`
	Permissions      []string          `json:"permissions,omitempty"`
	ExtraHTTPHeaders map[string]string `json:"extraHTTPHeaders,omitempty"`
	Offline          *bool             `json:"offline,omitempty"`
	HttpCredentials  *struct {
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
	} `json:"httpCredentials,omitempty"`
	ColorScheme *string     `json:"colorScheme,omitempty"`
	Logger      interface{} `json:"logger,omitempty"`
}
type ChromiumBrowserNewPageOptions struct {
	AcceptDownloads   *bool `json:"acceptDownloads,omitempty"`
	IgnoreHTTPSErrors *bool `json:"ignoreHTTPSErrors,omitempty"`
	BypassCSP         *bool `json:"bypassCSP,omitempty"`
	Viewport          *struct {
		Width  *int `json:"width,omitempty"`
		Height *int `json:"height,omitempty"`
	} `json:"viewport,omitempty"`
	UserAgent         *string `json:"userAgent,omitempty"`
	DeviceScaleFactor *int    `json:"deviceScaleFactor,omitempty"`
	IsMobile          *bool   `json:"isMobile,omitempty"`
	HasTouch          *bool   `json:"hasTouch,omitempty"`
	JavaScriptEnabled *bool   `json:"javaScriptEnabled,omitempty"`
	TimezoneId        *string `json:"timezoneId,omitempty"`
	Geolocation       *struct {
		Latitude  *int `json:"latitude,omitempty"`
		Longitude *int `json:"longitude,omitempty"`
		Accuracy  *int `json:"accuracy,omitempty"`
	} `json:"geolocation,omitempty"`
	Locale           *string           `json:"locale,omitempty"`
	Permissions      []string          `json:"permissions,omitempty"`
	ExtraHTTPHeaders map[string]string `json:"extraHTTPHeaders,omitempty"`
	Offline          *bool             `json:"offline,omitempty"`
	HttpCredentials  *struct {
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
	} `json:"httpCredentials,omitempty"`
	ColorScheme *string     `json:"colorScheme,omitempty"`
	Logger      interface{} `json:"logger,omitempty"`
}
type ChromiumBrowserContextGrantPermissionsOptions struct {
	Origin *string `json:"origin,omitempty"`
}
type ChromiumCoverageStartCSSCoverageOptions struct {
	ResetOnNavigation *bool `json:"resetOnNavigation,omitempty"`
}
type ChromiumCoverageStartJSCoverageOptions struct {
	ResetOnNavigation      *bool `json:"resetOnNavigation,omitempty"`
	ReportAnonymousScripts *bool `json:"reportAnonymousScripts,omitempty"`
}
type FirefoxBrowserNewContextOptions struct {
	AcceptDownloads   *bool `json:"acceptDownloads,omitempty"`
	IgnoreHTTPSErrors *bool `json:"ignoreHTTPSErrors,omitempty"`
	BypassCSP         *bool `json:"bypassCSP,omitempty"`
	Viewport          *struct {
		Width  *int `json:"width,omitempty"`
		Height *int `json:"height,omitempty"`
	} `json:"viewport,omitempty"`
	UserAgent         *string `json:"userAgent,omitempty"`
	DeviceScaleFactor *int    `json:"deviceScaleFactor,omitempty"`
	IsMobile          *bool   `json:"isMobile,omitempty"`
	HasTouch          *bool   `json:"hasTouch,omitempty"`
	JavaScriptEnabled *bool   `json:"javaScriptEnabled,omitempty"`
	TimezoneId        *string `json:"timezoneId,omitempty"`
	Geolocation       *struct {
		Latitude  *int `json:"latitude,omitempty"`
		Longitude *int `json:"longitude,omitempty"`
		Accuracy  *int `json:"accuracy,omitempty"`
	} `json:"geolocation,omitempty"`
	Locale           *string           `json:"locale,omitempty"`
	Permissions      []string          `json:"permissions,omitempty"`
	ExtraHTTPHeaders map[string]string `json:"extraHTTPHeaders,omitempty"`
	Offline          *bool             `json:"offline,omitempty"`
	HttpCredentials  *struct {
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
	} `json:"httpCredentials,omitempty"`
	ColorScheme *string     `json:"colorScheme,omitempty"`
	Logger      interface{} `json:"logger,omitempty"`
}
type FirefoxBrowserNewPageOptions struct {
	AcceptDownloads   *bool `json:"acceptDownloads,omitempty"`
	IgnoreHTTPSErrors *bool `json:"ignoreHTTPSErrors,omitempty"`
	BypassCSP         *bool `json:"bypassCSP,omitempty"`
	Viewport          *struct {
		Width  *int `json:"width,omitempty"`
		Height *int `json:"height,omitempty"`
	} `json:"viewport,omitempty"`
	UserAgent         *string `json:"userAgent,omitempty"`
	DeviceScaleFactor *int    `json:"deviceScaleFactor,omitempty"`
	IsMobile          *bool   `json:"isMobile,omitempty"`
	HasTouch          *bool   `json:"hasTouch,omitempty"`
	JavaScriptEnabled *bool   `json:"javaScriptEnabled,omitempty"`
	TimezoneId        *string `json:"timezoneId,omitempty"`
	Geolocation       *struct {
		Latitude  *int `json:"latitude,omitempty"`
		Longitude *int `json:"longitude,omitempty"`
		Accuracy  *int `json:"accuracy,omitempty"`
	} `json:"geolocation,omitempty"`
	Locale           *string           `json:"locale,omitempty"`
	Permissions      []string          `json:"permissions,omitempty"`
	ExtraHTTPHeaders map[string]string `json:"extraHTTPHeaders,omitempty"`
	Offline          *bool             `json:"offline,omitempty"`
	HttpCredentials  *struct {
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
	} `json:"httpCredentials,omitempty"`
	ColorScheme *string     `json:"colorScheme,omitempty"`
	Logger      interface{} `json:"logger,omitempty"`
}
type WebKitBrowserNewContextOptions struct {
	AcceptDownloads   *bool `json:"acceptDownloads,omitempty"`
	IgnoreHTTPSErrors *bool `json:"ignoreHTTPSErrors,omitempty"`
	BypassCSP         *bool `json:"bypassCSP,omitempty"`
	Viewport          *struct {
		Width  *int `json:"width,omitempty"`
		Height *int `json:"height,omitempty"`
	} `json:"viewport,omitempty"`
	UserAgent         *string `json:"userAgent,omitempty"`
	DeviceScaleFactor *int    `json:"deviceScaleFactor,omitempty"`
	IsMobile          *bool   `json:"isMobile,omitempty"`
	HasTouch          *bool   `json:"hasTouch,omitempty"`
	JavaScriptEnabled *bool   `json:"javaScriptEnabled,omitempty"`
	TimezoneId        *string `json:"timezoneId,omitempty"`
	Geolocation       *struct {
		Latitude  *int `json:"latitude,omitempty"`
		Longitude *int `json:"longitude,omitempty"`
		Accuracy  *int `json:"accuracy,omitempty"`
	} `json:"geolocation,omitempty"`
	Locale           *string           `json:"locale,omitempty"`
	Permissions      []string          `json:"permissions,omitempty"`
	ExtraHTTPHeaders map[string]string `json:"extraHTTPHeaders,omitempty"`
	Offline          *bool             `json:"offline,omitempty"`
	HttpCredentials  *struct {
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
	} `json:"httpCredentials,omitempty"`
	ColorScheme *string     `json:"colorScheme,omitempty"`
	Logger      interface{} `json:"logger,omitempty"`
}
type WebKitBrowserNewPageOptions struct {
	AcceptDownloads   *bool `json:"acceptDownloads,omitempty"`
	IgnoreHTTPSErrors *bool `json:"ignoreHTTPSErrors,omitempty"`
	BypassCSP         *bool `json:"bypassCSP,omitempty"`
	Viewport          *struct {
		Width  *int `json:"width,omitempty"`
		Height *int `json:"height,omitempty"`
	} `json:"viewport,omitempty"`
	UserAgent         *string `json:"userAgent,omitempty"`
	DeviceScaleFactor *int    `json:"deviceScaleFactor,omitempty"`
	IsMobile          *bool   `json:"isMobile,omitempty"`
	HasTouch          *bool   `json:"hasTouch,omitempty"`
	JavaScriptEnabled *bool   `json:"javaScriptEnabled,omitempty"`
	TimezoneId        *string `json:"timezoneId,omitempty"`
	Geolocation       *struct {
		Latitude  *int `json:"latitude,omitempty"`
		Longitude *int `json:"longitude,omitempty"`
		Accuracy  *int `json:"accuracy,omitempty"`
	} `json:"geolocation,omitempty"`
	Locale           *string           `json:"locale,omitempty"`
	Permissions      []string          `json:"permissions,omitempty"`
	ExtraHTTPHeaders map[string]string `json:"extraHTTPHeaders,omitempty"`
	Offline          *bool             `json:"offline,omitempty"`
	HttpCredentials  *struct {
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
	} `json:"httpCredentials,omitempty"`
	ColorScheme *string     `json:"colorScheme,omitempty"`
	Logger      interface{} `json:"logger,omitempty"`
}
