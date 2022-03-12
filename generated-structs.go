package playwright

type BrowserNewContextOptions struct {
	// Whether to automatically download all the attachments. Defaults to `true` where
	// all the downloads are accepted.
	AcceptDownloads *bool `json:"acceptDownloads"`
	// When using Page.Goto(), Page.Route(), Page.WaitForURL(), Page.WaitForRequest(),
	// or Page.WaitForResponse() it takes the base URL in consideration by using the [`URL()`](https://developer.mozilla.org/en-US/docs/Web/API/URL/URL)
	// constructor for building the corresponding URL. Examples:
	// baseURL: `http://localhost:3000` and navigating to `/bar.html` results in `http://localhost:3000/bar.html`
	// baseURL: `http://localhost:3000/foo/` and navigating to `./bar.html` results in
	// `http://localhost:3000/foo/bar.html`
	// baseURL: `http://localhost:3000/foo` (without trailing slash) and navigating to
	// `./bar.html` results in `http://localhost:3000/bar.html`
	BaseURL *string `json:"baseURL"`
	// Toggles bypassing page's Content-Security-Policy.
	BypassCSP *bool `json:"bypassCSP"`
	// Emulates `'prefers-colors-scheme'` media feature, supported values are `'light'`,
	// `'dark'`, `'no-preference'`. See Page.EmulateMedia() for more details. Defaults
	// to `'light'`.
	ColorScheme *ColorScheme `json:"colorScheme"`
	// Specify device scale factor (can be thought of as dpr). Defaults to `1`.
	DeviceScaleFactor *float64 `json:"deviceScaleFactor"`
	// An object containing additional HTTP headers to be sent with every request.
	ExtraHttpHeaders map[string]string `json:"extraHTTPHeaders"`
	// Emulates `'forced-colors'` media feature, supported values are `'active'`, `'none'`.
	// See Page.EmulateMedia() for more details. Defaults to `'none'`.
	// It's not supported in WebKit, see [here](https://bugs.webkit.org/show_bug.cgi?id=225281)
	// in their issue tracker.
	ForcedColors *ForcedColors                        `json:"forcedColors"`
	Geolocation  *BrowserNewContextOptionsGeolocation `json:"geolocation"`
	// Specifies if viewport supports touch events. Defaults to false.
	HasTouch *bool `json:"hasTouch"`
	// Credentials for [HTTP authentication](https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication).
	HttpCredentials *BrowserNewContextOptionsHttpCredentials `json:"httpCredentials"`
	// Whether to ignore HTTPS errors when sending network requests. Defaults to `false`.
	IgnoreHttpsErrors *bool `json:"ignoreHTTPSErrors"`
	// Whether the `meta viewport` tag is taken into account and touch events are enabled.
	// Defaults to `false`. Not supported in Firefox.
	IsMobile *bool `json:"isMobile"`
	// Whether or not to enable JavaScript in the context. Defaults to `true`.
	JavaScriptEnabled *bool `json:"javaScriptEnabled"`
	// Specify user locale, for example `en-GB`, `de-DE`, etc. Locale will affect `navigator.language`
	// value, `Accept-Language` request header value as well as number and date formatting
	// rules.
	Locale *string `json:"locale"`
	// Does not enforce fixed viewport, allows resizing window in the headed mode.
	NoViewport *bool `json:"noViewport"`
	// Whether to emulate network being offline. Defaults to `false`.
	Offline *bool `json:"offline"`
	// A list of permissions to grant to all pages in this context. See BrowserContext.GrantPermissions()
	// for more details.
	Permissions []string `json:"permissions"`
	// Network proxy settings to use with this context.
	// For Chromium on Windows the browser needs to be launched with the global proxy for
	// this option to work. If all contexts override the proxy, global proxy will be never
	// used and can be any string, for example `launch({ proxy: { server: 'http://per-context'
	// } })`.
	Proxy *BrowserNewContextOptionsProxy `json:"proxy"`
	// Enables video recording for all pages into `recordVideo.dir` directory. If not specified
	// videos are not recorded. Make sure to await BrowserContext.Close() for videos to
	// be saved.
	RecordVideo *BrowserNewContextOptionsRecordVideo `json:"recordVideo"`
	// Emulates `'prefers-reduced-motion'` media feature, supported values are `'reduce'`,
	// `'no-preference'`. See Page.EmulateMedia() for more details. Defaults to `'no-preference'`.
	ReducedMotion *ReducedMotion `json:"reducedMotion"`
	// Emulates consistent window screen size available inside web page via `window.screen`.
	// Is only used when the `viewport` is set.
	Screen *BrowserNewContextOptionsScreen `json:"screen"`
	// Populates context with given storage state. This option can be used to initialize
	// context with logged-in information obtained via BrowserContext.StorageState(). Either
	// a path to the file with saved storage, or an object with the following fields:
	StorageState *BrowserNewContextOptionsStorageState `json:"storageState"`
	// Populates context with given storage state. This option can be used to initialize
	// context with logged-in information obtained via BrowserContext.StorageState(). Path
	// to the file with saved storage state.
	StorageStatePath *string `json:"storageStatePath"`
	// It specified, enables strict selectors mode for this context. In the strict selectors
	// mode all operations on selectors that imply single target DOM element will throw
	// when more than one element matches the selector. See Locator to learn more about
	// the strict mode.
	StrictSelectors *bool `json:"strictSelectors"`
	// Changes the timezone of the context. See [ICU's metaZones.txt](https://cs.chromium.org/chromium/src/third_party/icu/source/data/misc/metaZones.txt?rcl=faee8bc70570192d82d2978a71e2a615788597d1)
	// for a list of supported timezone IDs.
	TimezoneId *string `json:"timezoneId"`
	// Specific user agent to use in this context.
	UserAgent *string `json:"userAgent"`
	// Sets a consistent viewport for each page. Defaults to an 1280x720 viewport. `no_viewport`
	// disables the fixed viewport.
	Viewport *BrowserNewContextOptionsViewport `json:"viewport"`
}
type BrowserGeolocation struct {
	// Latitude between -90 and 90.
	Latitude *float64 `json:"latitude"`
	// Longitude between -180 and 180.
	Longitude *float64 `json:"longitude"`
	// Non-negative accuracy value. Defaults to `0`.
	Accuracy *float64 `json:"accuracy"`
}
type BrowserHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type BrowserProxy struct {
	// Proxy to be used for all requests. HTTP and SOCKS proxies are supported, for example
	// `http://myproxy.com:3128` or `socks5://myproxy.com:3128`. Short form `myproxy.com:3128`
	// is considered an HTTP proxy.
	Server *string `json:"server"`
	// Optional comma-separated domains to bypass proxy, for example `".com, chromium.org,
	// .domain.com"`.
	Bypass *string `json:"bypass"`
	// Optional username to use if HTTP proxy requires authentication.
	Username *string `json:"username"`
	// Optional password to use if HTTP proxy requires authentication.
	Password *string `json:"password"`
}
type BrowserRecordVideo struct {
	// Path to the directory to put videos into.
	Dir *string `json:"dir"`
	// Optional dimensions of the recorded videos. If not specified the size will be equal
	// to `viewport` scaled down to fit into 800x800. If `viewport` is not configured explicitly
	// the video size defaults to 800x450. Actual picture of each page will be scaled down
	// if necessary to fit the specified size.
	Size *BrowserRecordVideoSize `json:"size"`
}
type BrowserScreen struct {
	// page width in pixels.
	Width *int `json:"width"`
	// page height in pixels.
	Height *int `json:"height"`
}
type BrowserStorageState struct {
	// cookies to set for context
	Cookies []BrowserStorageStateCookies `json:"cookies"`
	// localStorage to set for context
	Origins []BrowserStorageStateOrigins `json:"origins"`
}
type BrowserViewport struct {
	// page width in pixels.
	Width *int `json:"width"`
	// page height in pixels.
	Height *int `json:"height"`
}
type BrowserNewPageOptions struct {
	// Whether to automatically download all the attachments. Defaults to `true` where
	// all the downloads are accepted.
	AcceptDownloads *bool `json:"acceptDownloads"`
	// When using Page.Goto(), Page.Route(), Page.WaitForURL(), Page.WaitForRequest(),
	// or Page.WaitForResponse() it takes the base URL in consideration by using the [`URL()`](https://developer.mozilla.org/en-US/docs/Web/API/URL/URL)
	// constructor for building the corresponding URL. Examples:
	// baseURL: `http://localhost:3000` and navigating to `/bar.html` results in `http://localhost:3000/bar.html`
	// baseURL: `http://localhost:3000/foo/` and navigating to `./bar.html` results in
	// `http://localhost:3000/foo/bar.html`
	// baseURL: `http://localhost:3000/foo` (without trailing slash) and navigating to
	// `./bar.html` results in `http://localhost:3000/bar.html`
	BaseURL *string `json:"baseURL"`
	// Toggles bypassing page's Content-Security-Policy.
	BypassCSP *bool `json:"bypassCSP"`
	// Emulates `'prefers-colors-scheme'` media feature, supported values are `'light'`,
	// `'dark'`, `'no-preference'`. See Page.EmulateMedia() for more details. Defaults
	// to `'light'`.
	ColorScheme *ColorScheme `json:"colorScheme"`
	// Specify device scale factor (can be thought of as dpr). Defaults to `1`.
	DeviceScaleFactor *float64 `json:"deviceScaleFactor"`
	// An object containing additional HTTP headers to be sent with every request.
	ExtraHttpHeaders map[string]string `json:"extraHTTPHeaders"`
	// Emulates `'forced-colors'` media feature, supported values are `'active'`, `'none'`.
	// See Page.EmulateMedia() for more details. Defaults to `'none'`.
	// It's not supported in WebKit, see [here](https://bugs.webkit.org/show_bug.cgi?id=225281)
	// in their issue tracker.
	ForcedColors *ForcedColors                     `json:"forcedColors"`
	Geolocation  *BrowserNewPageOptionsGeolocation `json:"geolocation"`
	// Specifies if viewport supports touch events. Defaults to false.
	HasTouch *bool `json:"hasTouch"`
	// Credentials for [HTTP authentication](https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication).
	HttpCredentials *BrowserNewPageOptionsHttpCredentials `json:"httpCredentials"`
	// Whether to ignore HTTPS errors when sending network requests. Defaults to `false`.
	IgnoreHttpsErrors *bool `json:"ignoreHTTPSErrors"`
	// Whether the `meta viewport` tag is taken into account and touch events are enabled.
	// Defaults to `false`. Not supported in Firefox.
	IsMobile *bool `json:"isMobile"`
	// Whether or not to enable JavaScript in the context. Defaults to `true`.
	JavaScriptEnabled *bool `json:"javaScriptEnabled"`
	// Specify user locale, for example `en-GB`, `de-DE`, etc. Locale will affect `navigator.language`
	// value, `Accept-Language` request header value as well as number and date formatting
	// rules.
	Locale *string `json:"locale"`
	// Does not enforce fixed viewport, allows resizing window in the headed mode.
	NoViewport *bool `json:"noViewport"`
	// Whether to emulate network being offline. Defaults to `false`.
	Offline *bool `json:"offline"`
	// A list of permissions to grant to all pages in this context. See BrowserContext.GrantPermissions()
	// for more details.
	Permissions []string `json:"permissions"`
	// Network proxy settings to use with this context.
	// For Chromium on Windows the browser needs to be launched with the global proxy for
	// this option to work. If all contexts override the proxy, global proxy will be never
	// used and can be any string, for example `launch({ proxy: { server: 'http://per-context'
	// } })`.
	Proxy *BrowserNewPageOptionsProxy `json:"proxy"`
	// Enables video recording for all pages into `recordVideo.dir` directory. If not specified
	// videos are not recorded. Make sure to await BrowserContext.Close() for videos to
	// be saved.
	RecordVideo *BrowserNewPageOptionsRecordVideo `json:"recordVideo"`
	// Emulates `'prefers-reduced-motion'` media feature, supported values are `'reduce'`,
	// `'no-preference'`. See Page.EmulateMedia() for more details. Defaults to `'no-preference'`.
	ReducedMotion *ReducedMotion `json:"reducedMotion"`
	// Emulates consistent window screen size available inside web page via `window.screen`.
	// Is only used when the `viewport` is set.
	Screen *BrowserNewPageOptionsScreen `json:"screen"`
	// Populates context with given storage state. This option can be used to initialize
	// context with logged-in information obtained via BrowserContext.StorageState(). Either
	// a path to the file with saved storage, or an object with the following fields:
	StorageState *BrowserNewPageOptionsStorageState `json:"storageState"`
	// Populates context with given storage state. This option can be used to initialize
	// context with logged-in information obtained via BrowserContext.StorageState(). Path
	// to the file with saved storage state.
	StorageStatePath *string `json:"storageStatePath"`
	// It specified, enables strict selectors mode for this context. In the strict selectors
	// mode all operations on selectors that imply single target DOM element will throw
	// when more than one element matches the selector. See Locator to learn more about
	// the strict mode.
	StrictSelectors *bool `json:"strictSelectors"`
	// Changes the timezone of the context. See [ICU's metaZones.txt](https://cs.chromium.org/chromium/src/third_party/icu/source/data/misc/metaZones.txt?rcl=faee8bc70570192d82d2978a71e2a615788597d1)
	// for a list of supported timezone IDs.
	TimezoneId *string `json:"timezoneId"`
	// Specific user agent to use in this context.
	UserAgent *string `json:"userAgent"`
	// Sets a consistent viewport for each page. Defaults to an 1280x720 viewport. `no_viewport`
	// disables the fixed viewport.
	Viewport *BrowserNewPageOptionsViewport `json:"viewport"`
}
type BrowserContextAddCookiesOptions struct {
	Cookies []BrowserContextAddCookiesOptionsCookies `json:"cookies"`
}
type BrowserContextCookies struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
	// either url or domain / path are required. Optional.
	URL *string `json:"url"`
	// either url or domain / path are required Optional.
	Domain *string `json:"domain"`
	// either url or domain / path are required Optional.
	Path *string `json:"path"`
	// Unix time in seconds. Optional.
	Expires *float64 `json:"expires"`
	// Optional.
	HttpOnly *bool `json:"httpOnly"`
	// Optional.
	Secure *bool `json:"secure"`
	// Optional.
	SameSite *SameSiteAttribute `json:"sameSite"`
}
type BrowserContextAddInitScriptOptions struct {
	// Optional Script source to be evaluated in all pages in the browser context.
	Script *string `json:"script"`
	// Optional Script path to be evaluated in all pages in the browser context.
	Path *string `json:"path"`
}

// Result of calling <see cref="BrowserContext.Cookies" />.
type BrowserContextCookiesResult struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Domain string `json:"domain"`
	Path   string `json:"path"`
	// Unix time in seconds.
	Expires  float64           `json:"expires"`
	HttpOnly bool              `json:"httpOnly"`
	Secure   bool              `json:"secure"`
	SameSite SameSiteAttribute `json:"sameSite"`
}
type BrowserContextCookiesOptions struct {
	// Optional list of URLs.
	Urls []string `json:"urls"`
}
type BrowserContextExposeBindingOptions struct {
	// Whether to pass the argument as a handle, instead of passing by value. When passing
	// a handle, only one argument is supported. When passing by value, multiple arguments
	// are supported.
	Handle *bool `json:"handle"`
}
type BrowserContextGrantPermissionsOptions struct {
	// The [origin] to grant permissions to, e.g. "https://example.com".
	Origin *string `json:"origin"`
}
type BrowserContextRouteOptions struct {
	// How often a route should be used. By default it will be used every time.
	Times *int `json:"times"`
}
type BrowserContextGeolocation struct {
	// Latitude between -90 and 90.
	Latitude *float64 `json:"latitude"`
	// Longitude between -180 and 180.
	Longitude *float64 `json:"longitude"`
	// Non-negative accuracy value. Defaults to `0`.
	Accuracy *float64 `json:"accuracy"`
}

// Result of calling <see cref="BrowserContext.StorageState" />.
type BrowserContextStorageStateResult struct {
	Cookies []BrowserContextStorageStateResultCookies `json:"cookies"`
	Origins []BrowserContextStorageStateResultOrigins `json:"origins"`
}
type BrowserContextStorageStateOptions struct {
	// The file path to save the storage state to. If `path` is a relative path, then it
	// is resolved relative to current working directory. If no path is provided, storage
	// state is still returned, but won't be saved to the disk.
	Path *string `json:"path"`
}
type BrowserContextUnrouteOptions struct {
	// Optional handler function used to register a routing with BrowserContext.Route().
	Handler func(Route, Request) `json:"handler"`
}
type BrowserTypeConnectOptions struct {
	// Additional HTTP headers to be sent with web socket connect request. Optional.
	Headers map[string]string `json:"headers"`
	// Slows down Playwright operations by the specified amount of milliseconds. Useful
	// so that you can see what is going on. Defaults to 0.
	SlowMo *float64 `json:"slowMo"`
	// Maximum time in milliseconds to wait for the connection to be established. Defaults
	// to `30000` (30 seconds). Pass `0` to disable timeout.
	Timeout *float64 `json:"timeout"`
}
type BrowserTypeConnectOverCDPOptions struct {
	// Additional HTTP headers to be sent with connect request. Optional.
	Headers map[string]string `json:"headers"`
	// Slows down Playwright operations by the specified amount of milliseconds. Useful
	// so that you can see what is going on. Defaults to 0.
	SlowMo *float64 `json:"slowMo"`
	// Maximum time in milliseconds to wait for the connection to be established. Defaults
	// to `30000` (30 seconds). Pass `0` to disable timeout.
	Timeout *float64 `json:"timeout"`
}
type BrowserTypeLaunchOptions struct {
	// Additional arguments to pass to the browser instance. The list of Chromium flags
	// can be found [here](http://peter.sh/experiments/chromium-command-line-switches/).
	Args []string `json:"args"`
	// Browser distribution channel.  Supported values are "chrome", "chrome-beta", "chrome-dev",
	// "chrome-canary", "msedge", "msedge-beta", "msedge-dev", "msedge-canary". Read more
	// about using [Google Chrome and Microsoft Edge](./browsers.md#google-chrome--microsoft-edge).
	Channel *string `json:"channel"`
	// Enable Chromium sandboxing. Defaults to `false`.
	ChromiumSandbox *bool `json:"chromiumSandbox"`
	// **Chromium-only** Whether to auto-open a Developer Tools panel for each tab. If
	// this option is `true`, the `headless` option will be set `false`.
	Devtools *bool `json:"devtools"`
	// If specified, accepted downloads are downloaded into this directory. Otherwise,
	// temporary directory is created and is deleted when browser is closed. In either
	// case, the downloads are deleted when the browser context they were created in is
	// closed.
	DownloadsPath *string `json:"downloadsPath"`
	// Specify environment variables that will be visible to the browser. Defaults to `process.env`.
	Env map[string]string `json:"env"`
	// Path to a browser executable to run instead of the bundled one. If `executablePath`
	// is a relative path, then it is resolved relative to the current working directory.
	// Note that Playwright only works with the bundled Chromium, Firefox or WebKit, use
	// at your own risk.
	ExecutablePath *string `json:"executablePath"`
	// Firefox user preferences. Learn more about the Firefox user preferences at [`about:config`](https://support.mozilla.org/en-US/kb/about-config-editor-firefox).
	FirefoxUserPrefs map[string]interface{} `json:"firefoxUserPrefs"`
	// Close the browser process on SIGHUP. Defaults to `true`.
	HandleSIGHUP *bool `json:"handleSIGHUP"`
	// Close the browser process on Ctrl-C. Defaults to `true`.
	HandleSIGINT *bool `json:"handleSIGINT"`
	// Close the browser process on SIGTERM. Defaults to `true`.
	HandleSIGTERM *bool `json:"handleSIGTERM"`
	// Whether to run browser in headless mode. More details for [Chromium](https://developers.google.com/web/updates/2017/04/headless-chrome)
	// and [Firefox](https://developer.mozilla.org/en-US/docs/Mozilla/Firefox/Headless_mode).
	// Defaults to `true` unless the `devtools` option is `true`.
	Headless *bool `json:"headless"`
	// Network proxy settings.
	Proxy *BrowserTypeLaunchOptionsProxy `json:"proxy"`
	// Slows down Playwright operations by the specified amount of milliseconds. Useful
	// so that you can see what is going on.
	SlowMo *float64 `json:"slowMo"`
	// Maximum time in milliseconds to wait for the browser instance to start. Defaults
	// to `30000` (30 seconds). Pass `0` to disable timeout.
	Timeout *float64 `json:"timeout"`
	// If specified, traces are saved into this directory.
	TracesDir *string `json:"tracesDir"`
}
type BrowserTypeProxy struct {
	// Proxy to be used for all requests. HTTP and SOCKS proxies are supported, for example
	// `http://myproxy.com:3128` or `socks5://myproxy.com:3128`. Short form `myproxy.com:3128`
	// is considered an HTTP proxy.
	Server *string `json:"server"`
	// Optional comma-separated domains to bypass proxy, for example `".com, chromium.org,
	// .domain.com"`.
	Bypass *string `json:"bypass"`
	// Optional username to use if HTTP proxy requires authentication.
	Username *string `json:"username"`
	// Optional password to use if HTTP proxy requires authentication.
	Password *string `json:"password"`
}
type BrowserTypeLaunchPersistentContextOptions struct {
	// Whether to automatically download all the attachments. Defaults to `true` where
	// all the downloads are accepted.
	AcceptDownloads *bool `json:"acceptDownloads"`
	// Additional arguments to pass to the browser instance. The list of Chromium flags
	// can be found [here](http://peter.sh/experiments/chromium-command-line-switches/).
	Args []string `json:"args"`
	// When using Page.Goto(), Page.Route(), Page.WaitForURL(), Page.WaitForRequest(),
	// or Page.WaitForResponse() it takes the base URL in consideration by using the [`URL()`](https://developer.mozilla.org/en-US/docs/Web/API/URL/URL)
	// constructor for building the corresponding URL. Examples:
	// baseURL: `http://localhost:3000` and navigating to `/bar.html` results in `http://localhost:3000/bar.html`
	// baseURL: `http://localhost:3000/foo/` and navigating to `./bar.html` results in
	// `http://localhost:3000/foo/bar.html`
	// baseURL: `http://localhost:3000/foo` (without trailing slash) and navigating to
	// `./bar.html` results in `http://localhost:3000/bar.html`
	BaseURL *string `json:"baseURL"`
	// Toggles bypassing page's Content-Security-Policy.
	BypassCSP *bool `json:"bypassCSP"`
	// Browser distribution channel.  Supported values are "chrome", "chrome-beta", "chrome-dev",
	// "chrome-canary", "msedge", "msedge-beta", "msedge-dev", "msedge-canary". Read more
	// about using [Google Chrome and Microsoft Edge](./browsers.md#google-chrome--microsoft-edge).
	Channel *string `json:"channel"`
	// Enable Chromium sandboxing. Defaults to `false`.
	ChromiumSandbox *bool `json:"chromiumSandbox"`
	// Emulates `'prefers-colors-scheme'` media feature, supported values are `'light'`,
	// `'dark'`, `'no-preference'`. See Page.EmulateMedia() for more details. Defaults
	// to `'light'`.
	ColorScheme *ColorScheme `json:"colorScheme"`
	// Specify device scale factor (can be thought of as dpr). Defaults to `1`.
	DeviceScaleFactor *float64 `json:"deviceScaleFactor"`
	// **Chromium-only** Whether to auto-open a Developer Tools panel for each tab. If
	// this option is `true`, the `headless` option will be set `false`.
	Devtools *bool `json:"devtools"`
	// If specified, accepted downloads are downloaded into this directory. Otherwise,
	// temporary directory is created and is deleted when browser is closed. In either
	// case, the downloads are deleted when the browser context they were created in is
	// closed.
	DownloadsPath *string `json:"downloadsPath"`
	// Specify environment variables that will be visible to the browser. Defaults to `process.env`.
	Env map[string]string `json:"env"`
	// Path to a browser executable to run instead of the bundled one. If `executablePath`
	// is a relative path, then it is resolved relative to the current working directory.
	// Note that Playwright only works with the bundled Chromium, Firefox or WebKit, use
	// at your own risk.
	ExecutablePath *string `json:"executablePath"`
	// An object containing additional HTTP headers to be sent with every request.
	ExtraHttpHeaders map[string]string `json:"extraHTTPHeaders"`
	// Emulates `'forced-colors'` media feature, supported values are `'active'`, `'none'`.
	// See Page.EmulateMedia() for more details. Defaults to `'none'`.
	// It's not supported in WebKit, see [here](https://bugs.webkit.org/show_bug.cgi?id=225281)
	// in their issue tracker.
	ForcedColors *ForcedColors                                         `json:"forcedColors"`
	Geolocation  *BrowserTypeLaunchPersistentContextOptionsGeolocation `json:"geolocation"`
	// Close the browser process on SIGHUP. Defaults to `true`.
	HandleSIGHUP *bool `json:"handleSIGHUP"`
	// Close the browser process on Ctrl-C. Defaults to `true`.
	HandleSIGINT *bool `json:"handleSIGINT"`
	// Close the browser process on SIGTERM. Defaults to `true`.
	HandleSIGTERM *bool `json:"handleSIGTERM"`
	// Specifies if viewport supports touch events. Defaults to false.
	HasTouch *bool `json:"hasTouch"`
	// Whether to run browser in headless mode. More details for [Chromium](https://developers.google.com/web/updates/2017/04/headless-chrome)
	// and [Firefox](https://developer.mozilla.org/en-US/docs/Mozilla/Firefox/Headless_mode).
	// Defaults to `true` unless the `devtools` option is `true`.
	Headless *bool `json:"headless"`
	// Credentials for [HTTP authentication](https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication).
	HttpCredentials *BrowserTypeLaunchPersistentContextOptionsHttpCredentials `json:"httpCredentials"`
	// Whether to ignore HTTPS errors when sending network requests. Defaults to `false`.
	IgnoreHttpsErrors *bool `json:"ignoreHTTPSErrors"`
	// Whether the `meta viewport` tag is taken into account and touch events are enabled.
	// Defaults to `false`. Not supported in Firefox.
	IsMobile *bool `json:"isMobile"`
	// Whether or not to enable JavaScript in the context. Defaults to `true`.
	JavaScriptEnabled *bool `json:"javaScriptEnabled"`
	// Specify user locale, for example `en-GB`, `de-DE`, etc. Locale will affect `navigator.language`
	// value, `Accept-Language` request header value as well as number and date formatting
	// rules.
	Locale *string `json:"locale"`
	// Does not enforce fixed viewport, allows resizing window in the headed mode.
	NoViewport *bool `json:"noViewport"`
	// Whether to emulate network being offline. Defaults to `false`.
	Offline *bool `json:"offline"`
	// A list of permissions to grant to all pages in this context. See BrowserContext.GrantPermissions()
	// for more details.
	Permissions []string `json:"permissions"`
	// Network proxy settings.
	Proxy *BrowserTypeLaunchPersistentContextOptionsProxy `json:"proxy"`
	// Enables video recording for all pages into `recordVideo.dir` directory. If not specified
	// videos are not recorded. Make sure to await BrowserContext.Close() for videos to
	// be saved.
	RecordVideo *BrowserTypeLaunchPersistentContextOptionsRecordVideo `json:"recordVideo"`
	// Emulates `'prefers-reduced-motion'` media feature, supported values are `'reduce'`,
	// `'no-preference'`. See Page.EmulateMedia() for more details. Defaults to `'no-preference'`.
	ReducedMotion *ReducedMotion `json:"reducedMotion"`
	// Emulates consistent window screen size available inside web page via `window.screen`.
	// Is only used when the `viewport` is set.
	Screen *BrowserTypeLaunchPersistentContextOptionsScreen `json:"screen"`
	// Slows down Playwright operations by the specified amount of milliseconds. Useful
	// so that you can see what is going on.
	SlowMo *float64 `json:"slowMo"`
	// It specified, enables strict selectors mode for this context. In the strict selectors
	// mode all operations on selectors that imply single target DOM element will throw
	// when more than one element matches the selector. See Locator to learn more about
	// the strict mode.
	StrictSelectors *bool `json:"strictSelectors"`
	// Maximum time in milliseconds to wait for the browser instance to start. Defaults
	// to `30000` (30 seconds). Pass `0` to disable timeout.
	Timeout *float64 `json:"timeout"`
	// Changes the timezone of the context. See [ICU's metaZones.txt](https://cs.chromium.org/chromium/src/third_party/icu/source/data/misc/metaZones.txt?rcl=faee8bc70570192d82d2978a71e2a615788597d1)
	// for a list of supported timezone IDs.
	TimezoneId *string `json:"timezoneId"`
	// If specified, traces are saved into this directory.
	TracesDir *string `json:"tracesDir"`
	// Specific user agent to use in this context.
	UserAgent *string `json:"userAgent"`
	// Sets a consistent viewport for each page. Defaults to an 1280x720 viewport. `no_viewport`
	// disables the fixed viewport.
	Viewport *BrowserTypeLaunchPersistentContextOptionsViewport `json:"viewport"`
}
type BrowserTypeGeolocation struct {
	// Latitude between -90 and 90.
	Latitude *float64 `json:"latitude"`
	// Longitude between -180 and 180.
	Longitude *float64 `json:"longitude"`
	// Non-negative accuracy value. Defaults to `0`.
	Accuracy *float64 `json:"accuracy"`
}
type BrowserTypeHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type BrowserTypeRecordVideo struct {
	// Path to the directory to put videos into.
	Dir *string `json:"dir"`
	// Optional dimensions of the recorded videos. If not specified the size will be equal
	// to `viewport` scaled down to fit into 800x800. If `viewport` is not configured explicitly
	// the video size defaults to 800x450. Actual picture of each page will be scaled down
	// if necessary to fit the specified size.
	Size *BrowserTypeRecordVideoSize `json:"size"`
}
type BrowserTypeScreen struct {
	// page width in pixels.
	Width *int `json:"width"`
	// page height in pixels.
	Height *int `json:"height"`
}
type BrowserTypeViewport struct {
	// page width in pixels.
	Width *int `json:"width"`
	// page height in pixels.
	Height *int `json:"height"`
}
type DialogAcceptOptions struct {
	// A text to enter in prompt. Does not cause any effects if the dialog's `type` is
	// not prompt. Optional.
	PromptText *string `json:"promptText"`
}

// Result of calling <see cref="ElementHandle.BoundingBox" />.
type ElementHandleBoundingBoxResult struct {
	// the x coordinate of the element in pixels.
	X float64 `json:"x"`
	// the y coordinate of the element in pixels.
	Y float64 `json:"y"`
	// the width of the element in pixels.
	Width float64 `json:"width"`
	// the height of the element in pixels.
	Height float64 `json:"height"`
}
type ElementHandleCheckOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *ElementHandleCheckOptionsPosition `json:"position"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type ElementHandlePosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type ElementHandleClickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// defaults to 1. See [UIEvent.detail].
	ClickCount *int `json:"clickCount"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *ElementHandleClickOptionsPosition `json:"position"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type ElementHandleDblclickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *ElementHandleDblclickOptionsPosition `json:"position"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type ElementHandleDispatchEventOptions struct {
	// Optional event-specific initialization properties.
	EventInit interface{} `json:"eventInit"`
}
type ElementHandleEvalOnSelectorOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type ElementHandleEvalOnSelectorAllOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type ElementHandleFillOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandleHoverOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *ElementHandleHoverOptionsPosition `json:"position"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type ElementHandleInputValueOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandlePressOptions struct {
	// Time to wait between `keydown` and `keyup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandleScreenshotOptions struct {
	// When set to `"disabled"`, stops CSS animations, CSS transitions and Web Animations.
	// Animations get different treatment depending on their duration:
	// finite animations are fast-forwarded to completion, so they'll fire `transitionend`
	// event.
	// infinite animations are canceled to initial state, and then played over after the
	// screenshot.
	// Defaults to `"allow"` that leaves animations untouched.
	Animations *ScreenshotAnimations `json:"animations"`
	// When set to `"ready"`, screenshot will wait for [`document.fonts.ready`](https://developer.mozilla.org/en-US/docs/Web/API/FontFaceSet/ready)
	// promise to resolve in all frames. Defaults to `"nowait"`.
	Fonts *ScreenshotFonts `json:"fonts"`
	// Hides default white background and allows capturing screenshots with transparency.
	// Not applicable to `jpeg` images. Defaults to `false`.
	OmitBackground *bool `json:"omitBackground"`
	// The file path to save the image to. The screenshot type will be inferred from file
	// extension. If `path` is a relative path, then it is resolved relative to the current
	// working directory. If no path is provided, the image won't be saved to the disk.
	Path *string `json:"path"`
	// The quality of the image, between 0-100. Not applicable to `png` images.
	Quality *int `json:"quality"`
	// When set to `"css"`, screenshot will have a single pixel per each css pixel on the
	// page. For high-dpi devices, this will keep screenshots small. Using `"device"` option
	// will produce a single pixel per each device pixel, so screenhots of high-dpi devices
	// will be twice as large or even larger. Defaults to `"device"`.
	Size *ScreenshotSize `json:"size"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// Specify screenshot type, defaults to `png`.
	Type *ScreenshotType `json:"type"`
}
type ElementHandleScrollIntoViewIfNeededOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandleSelectOptionOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandleSelectTextOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandleSetCheckedOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *ElementHandleSetCheckedOptionsPosition `json:"position"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type ElementHandleSetInputFilesOptions struct {
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandleTapOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *ElementHandleTapOptionsPosition `json:"position"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type ElementHandleTypeOptions struct {
	// Time to wait between key presses in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandleUncheckOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *ElementHandleUncheckOptionsPosition `json:"position"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type ElementHandleWaitForElementStateOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandleWaitForSelectorOptions struct {
	// Defaults to `'visible'`. Can be either:
	// `'attached'` - wait for element to be present in DOM.
	// `'detached'` - wait for element to not be present in DOM.
	// `'visible'` - wait for element to have non-empty bounding box and no `visibility:hidden`.
	// Note that element without any content or with `display:none` has an empty bounding
	// box and is not considered visible.
	// `'hidden'` - wait for element to be either detached from DOM, or have an empty bounding
	// box or `visibility:hidden`. This is opposite to the `'visible'` option.
	State *WaitForSelectorState `json:"state"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FileChooserSetFilesOptions struct {
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameAddScriptTagOptions struct {
	// Raw JavaScript content to be injected into frame.
	Content *string `json:"content"`
	// Path to the JavaScript file to be injected into frame. If `path` is a relative path,
	// then it is resolved relative to the current working directory.
	Path *string `json:"path"`
	// Script type. Use 'module' in order to load a Javascript ES6 module. See [script](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script)
	// for more details.
	Type *string `json:"type"`
	// URL of a script to be added.
	URL *string `json:"url"`
}
type FrameAddStyleTagOptions struct {
	// Raw CSS content to be injected into frame.
	Content *string `json:"content"`
	// Path to the CSS file to be injected into frame. If `path` is a relative path, then
	// it is resolved relative to the current working directory.
	Path *string `json:"path"`
	// URL of the `<link>` tag.
	URL *string `json:"url"`
}
type FrameCheckOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *FrameCheckOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FramePosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type FrameClickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// defaults to 1. See [UIEvent.detail].
	ClickCount *int `json:"clickCount"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *FrameClickOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FrameDblclickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *FrameDblclickOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FrameDispatchEventOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameDragAndDropOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Clicks on the source element at this point relative to the top-left corner of the
	// element's padding box. If not specified, some visible point of the element is used.
	SourcePosition *FrameDragAndDropOptionsSourcePosition `json:"sourcePosition"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Drops on the target element at this point relative to the top-left corner of the
	// element's padding box. If not specified, some visible point of the element is used.
	TargetPosition *FrameDragAndDropOptionsTargetPosition `json:"targetPosition"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FrameSourcePosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type FrameTargetPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type FrameEvalOnSelectorOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
}
type FrameEvalOnSelectorAllOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type FrameEvaluateOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type FrameEvaluateHandleOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type FrameFillOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameFocusOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameGetAttributeOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameGotoOptions struct {
	// Referer header value. If provided it will take preference over the referer header
	// value set by Page.SetExtraHttpHeaders().
	Referer *string `json:"referer"`
	// Maximum operation time in milliseconds, defaults to 30 seconds, pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultNavigationTimeout(),
	// BrowserContext.SetDefaultTimeout(), Page.SetDefaultNavigationTimeout() or Page.SetDefaultTimeout()
	// methods.
	Timeout *float64 `json:"timeout"`
	// When to consider operation succeeded, defaults to `load`. Events can be either:
	// `'domcontentloaded'` - consider operation to be finished when the `DOMContentLoaded`
	// event is fired.
	// `'load'` - consider operation to be finished when the `load` event is fired.
	// `'networkidle'` - consider operation to be finished when there are no network connections
	// for at least `500` ms.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type FrameHoverOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *FrameHoverOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FrameInnerHTMLOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameInnerTextOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameInputValueOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameIsCheckedOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameIsDisabledOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameIsEditableOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameIsEnabledOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameIsHiddenOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// **DEPRECATED** This option is ignored. Frame.IsHidden() does not wait for the element
	// to become hidden and returns immediately.
	Timeout *float64 `json:"timeout"`
}
type FrameIsVisibleOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// **DEPRECATED** This option is ignored. Frame.IsVisible() does not wait for the element
	// to become visible and returns immediately.
	Timeout *float64 `json:"timeout"`
}
type FramePressOptions struct {
	// Time to wait between `keydown` and `keyup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameQuerySelectorOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
}
type FrameSelectOptionOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameSetCheckedOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *FrameSetCheckedOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FrameSetContentOptions struct {
	// Maximum operation time in milliseconds, defaults to 30 seconds, pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultNavigationTimeout(),
	// BrowserContext.SetDefaultTimeout(), Page.SetDefaultNavigationTimeout() or Page.SetDefaultTimeout()
	// methods.
	Timeout *float64 `json:"timeout"`
	// When to consider operation succeeded, defaults to `load`. Events can be either:
	// `'domcontentloaded'` - consider operation to be finished when the `DOMContentLoaded`
	// event is fired.
	// `'load'` - consider operation to be finished when the `load` event is fired.
	// `'networkidle'` - consider operation to be finished when there are no network connections
	// for at least `500` ms.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type FrameSetInputFilesOptions struct {
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameTapOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *FrameTapOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FrameTextContentOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameTypeOptions struct {
	// Time to wait between key presses in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameUncheckOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *FrameUncheckOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FrameWaitForFunctionOptions struct {
	// If `polling` is `'raf'`, then `expression` is constantly executed in `requestAnimationFrame`
	// callback. If `polling` is a number, then it is treated as an interval in milliseconds
	// at which the function would be executed. Defaults to `raf`.
	Polling interface{} `json:"polling"`
	// maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type FrameWaitForLoadStateOptions struct {
	// Maximum operation time in milliseconds, defaults to 30 seconds, pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultNavigationTimeout(),
	// BrowserContext.SetDefaultTimeout(), Page.SetDefaultNavigationTimeout() or Page.SetDefaultTimeout()
	// methods.
	Timeout *float64 `json:"timeout"`
}
type FrameWaitForNavigationOptions struct {
	// Maximum operation time in milliseconds, defaults to 30 seconds, pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultNavigationTimeout(),
	// BrowserContext.SetDefaultTimeout(), Page.SetDefaultNavigationTimeout() or Page.SetDefaultTimeout()
	// methods.
	Timeout *float64 `json:"timeout"`
	// A glob pattern, regex pattern or predicate receiving [URL] to match while waiting
	// for the navigation. Note that if the parameter is a string without wilcard characters,
	// the method will wait for navigation to URL that is exactly equal to the string.
	URL interface{} `json:"url"`
	// When to consider operation succeeded, defaults to `load`. Events can be either:
	// `'domcontentloaded'` - consider operation to be finished when the `DOMContentLoaded`
	// event is fired.
	// `'load'` - consider operation to be finished when the `load` event is fired.
	// `'networkidle'` - consider operation to be finished when there are no network connections
	// for at least `500` ms.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type FrameWaitForSelectorOptions struct {
	// Defaults to `'visible'`. Can be either:
	// `'attached'` - wait for element to be present in DOM.
	// `'detached'` - wait for element to not be present in DOM.
	// `'visible'` - wait for element to have non-empty bounding box and no `visibility:hidden`.
	// Note that element without any content or with `display:none` has an empty bounding
	// box and is not considered visible.
	// `'hidden'` - wait for element to be either detached from DOM, or have an empty bounding
	// box or `visibility:hidden`. This is opposite to the `'visible'` option.
	State *WaitForSelectorState `json:"state"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameWaitForURLOptions struct {
	// Maximum operation time in milliseconds, defaults to 30 seconds, pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultNavigationTimeout(),
	// BrowserContext.SetDefaultTimeout(), Page.SetDefaultNavigationTimeout() or Page.SetDefaultTimeout()
	// methods.
	Timeout *float64 `json:"timeout"`
	// When to consider operation succeeded, defaults to `load`. Events can be either:
	// `'domcontentloaded'` - consider operation to be finished when the `DOMContentLoaded`
	// event is fired.
	// `'load'` - consider operation to be finished when the `load` event is fired.
	// `'networkidle'` - consider operation to be finished when there are no network connections
	// for at least `500` ms.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type JSHandleEvaluateOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type JSHandleEvaluateHandleOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type KeyboardPressOptions struct {
	// Time to wait between `keydown` and `keyup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
}
type KeyboardTypeOptions struct {
	// Time to wait between key presses in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
}

// Result of calling <see cref="Locator.BoundingBox" />.
type LocatorBoundingBoxResult struct {
	// the x coordinate of the element in pixels.
	X float64 `json:"x"`
	// the y coordinate of the element in pixels.
	Y float64 `json:"y"`
	// the width of the element in pixels.
	Width float64 `json:"width"`
	// the height of the element in pixels.
	Height float64 `json:"height"`
}
type LocatorBoundingBoxOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorCheckOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *LocatorCheckOptionsPosition `json:"position"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type LocatorPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type LocatorClickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// defaults to 1. See [UIEvent.detail].
	ClickCount *int `json:"clickCount"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *LocatorClickOptionsPosition `json:"position"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type LocatorDblclickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *LocatorDblclickOptionsPosition `json:"position"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type LocatorDispatchEventOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorDragToOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Clicks on the source element at this point relative to the top-left corner of the
	// element's padding box. If not specified, some visible point of the element is used.
	SourcePosition *LocatorDragToOptionsSourcePosition `json:"sourcePosition"`
	// Drops on the target element at this point relative to the top-left corner of the
	// element's padding box. If not specified, some visible point of the element is used.
	TargetPosition *LocatorDragToOptionsTargetPosition `json:"targetPosition"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type LocatorSourcePosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type LocatorTargetPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type LocatorElementHandleOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorEvaluateOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorEvaluateAllOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type LocatorEvaluateHandleOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorFillOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorFocusOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorGetAttributeOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorHoverOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *LocatorHoverOptionsPosition `json:"position"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type LocatorInnerHTMLOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorInnerTextOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorInputValueOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorIsCheckedOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorIsDisabledOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorIsEditableOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorIsEnabledOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorIsHiddenOptions struct {
	// **DEPRECATED** This option is ignored. Locator.IsHidden() does not wait for the
	// element to become hidden and returns immediately.
	Timeout *float64 `json:"timeout"`
}
type LocatorIsVisibleOptions struct {
	// **DEPRECATED** This option is ignored. Locator.IsVisible() does not wait for the
	// element to become visible and returns immediately.
	Timeout *float64 `json:"timeout"`
}
type LocatorPressOptions struct {
	// Time to wait between `keydown` and `keyup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorScreenshotOptions struct {
	// When set to `"disabled"`, stops CSS animations, CSS transitions and Web Animations.
	// Animations get different treatment depending on their duration:
	// finite animations are fast-forwarded to completion, so they'll fire `transitionend`
	// event.
	// infinite animations are canceled to initial state, and then played over after the
	// screenshot.
	// Defaults to `"allow"` that leaves animations untouched.
	Animations *ScreenshotAnimations `json:"animations"`
	// When set to `"ready"`, screenshot will wait for [`document.fonts.ready`](https://developer.mozilla.org/en-US/docs/Web/API/FontFaceSet/ready)
	// promise to resolve in all frames. Defaults to `"nowait"`.
	Fonts *ScreenshotFonts `json:"fonts"`
	// Hides default white background and allows capturing screenshots with transparency.
	// Not applicable to `jpeg` images. Defaults to `false`.
	OmitBackground *bool `json:"omitBackground"`
	// The file path to save the image to. The screenshot type will be inferred from file
	// extension. If `path` is a relative path, then it is resolved relative to the current
	// working directory. If no path is provided, the image won't be saved to the disk.
	Path *string `json:"path"`
	// The quality of the image, between 0-100. Not applicable to `png` images.
	Quality *int `json:"quality"`
	// When set to `"css"`, screenshot will have a single pixel per each css pixel on the
	// page. For high-dpi devices, this will keep screenshots small. Using `"device"` option
	// will produce a single pixel per each device pixel, so screenhots of high-dpi devices
	// will be twice as large or even larger. Defaults to `"device"`.
	Size *ScreenshotSize `json:"size"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// Specify screenshot type, defaults to `png`.
	Type *ScreenshotType `json:"type"`
}
type LocatorScrollIntoViewIfNeededOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorSelectOptionOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorSelectTextOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorSetCheckedOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *LocatorSetCheckedOptionsPosition `json:"position"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type LocatorSetInputFilesOptions struct {
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorTapOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *LocatorTapOptionsPosition `json:"position"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type LocatorTextContentOptions struct {
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorTypeOptions struct {
	// Time to wait between key presses in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorUncheckOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *LocatorUncheckOptionsPosition `json:"position"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type LocatorWaitForOptions struct {
	// Defaults to `'visible'`. Can be either:
	// `'attached'` - wait for element to be present in DOM.
	// `'detached'` - wait for element to not be present in DOM.
	// `'visible'` - wait for element to have non-empty bounding box and no `visibility:hidden`.
	// Note that element without any content or with `display:none` has an empty bounding
	// box and is not considered visible.
	// `'hidden'` - wait for element to be either detached from DOM, or have an empty bounding
	// box or `visibility:hidden`. This is opposite to the `'visible'` option.
	State *WaitForSelectorState `json:"state"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToBeCheckedOptions struct {
	Checked *bool `json:"checked"`
}
type LocatorAssertionsToContainTextOptions struct {
	// Whether to use `element.innerText` instead of `element.textContent` when retrieving
	// DOM node text.
	UseInnerText *bool `json:"useInnerText"`
}
type LocatorAssertionsToHaveTextOptions struct {
	// Whether to use `element.innerText` instead of `element.textContent` when retrieving
	// DOM node text.
	UseInnerText *bool `json:"useInnerText"`
}
type MouseClickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// defaults to 1. See [UIEvent.detail].
	ClickCount *int `json:"clickCount"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
}
type MouseDblclickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
}
type MouseDownOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// defaults to 1. See [UIEvent.detail].
	ClickCount *int `json:"clickCount"`
}
type MouseMoveOptions struct {
	// defaults to 1. Sends intermediate `mousemove` events.
	Steps *int `json:"steps"`
}
type MouseUpOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// defaults to 1. See [UIEvent.detail].
	ClickCount *int `json:"clickCount"`
}
type PageAddInitScriptOptions struct {
	// Optional Script source to be evaluated in all pages in the browser context.
	Script *string `json:"script"`
	// Optional Script path to be evaluated in all pages in the browser context.
	Path *string `json:"path"`
}
type PageAddScriptTagOptions struct {
	// Raw JavaScript content to be injected into frame.
	Content *string `json:"content"`
	// Path to the JavaScript file to be injected into frame. If `path` is a relative path,
	// then it is resolved relative to the current working directory.
	Path *string `json:"path"`
	// Script type. Use 'module' in order to load a Javascript ES6 module. See [script](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script)
	// for more details.
	Type *string `json:"type"`
	// URL of a script to be added.
	URL *string `json:"url"`
}
type PageAddStyleTagOptions struct {
	// Raw CSS content to be injected into frame.
	Content *string `json:"content"`
	// Path to the CSS file to be injected into frame. If `path` is a relative path, then
	// it is resolved relative to the current working directory.
	Path *string `json:"path"`
	// URL of the `<link>` tag.
	URL *string `json:"url"`
}
type PageCheckOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *PageCheckOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type PagePosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type PageClickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// defaults to 1. See [UIEvent.detail].
	ClickCount *int `json:"clickCount"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *PageClickOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type PageCloseOptions struct {
	// Defaults to `false`. Whether to run the [before unload](https://developer.mozilla.org/en-US/docs/Web/Events/beforeunload)
	// page handlers.
	RunBeforeUnload *bool `json:"runBeforeUnload"`
}
type PageDblclickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *PageDblclickOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type PageDispatchEventOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageDragAndDropOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Clicks on the source element at this point relative to the top-left corner of the
	// element's padding box. If not specified, some visible point of the element is used.
	SourcePosition *PageDragAndDropOptionsSourcePosition `json:"sourcePosition"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Drops on the target element at this point relative to the top-left corner of the
	// element's padding box. If not specified, some visible point of the element is used.
	TargetPosition *PageDragAndDropOptionsTargetPosition `json:"targetPosition"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type PageSourcePosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type PageTargetPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type PageEmulateMediaOptions struct {
	// Emulates `'prefers-colors-scheme'` media feature, supported values are `'light'`,
	// `'dark'`, `'no-preference'`. Passing `'Null'` disables color scheme emulation.
	ColorScheme *ColorScheme `json:"colorScheme"`
	// Changes the CSS media type of the page. The only allowed values are `'Screen'`,
	// `'Print'` and `'Null'`. Passing `'Null'` disables CSS media emulation.
	Media *Media `json:"media"`
	// Emulates `'prefers-reduced-motion'` media feature, supported values are `'reduce'`,
	// `'no-preference'`. Passing `null` disables reduced motion emulation.
	ReducedMotion *ReducedMotion `json:"reducedMotion"`
}
type PageEvalOnSelectorOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
}
type PageEvalOnSelectorAllOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type PageEvaluateOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type PageEvaluateHandleOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type PageExposeBindingOptions struct {
	// Whether to pass the argument as a handle, instead of passing by value. When passing
	// a handle, only one argument is supported. When passing by value, multiple arguments
	// are supported.
	Handle *bool `json:"handle"`
}
type PageFillOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageFocusOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageGetAttributeOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageGoBackOptions struct {
	// Maximum operation time in milliseconds, defaults to 30 seconds, pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultNavigationTimeout(),
	// BrowserContext.SetDefaultTimeout(), Page.SetDefaultNavigationTimeout() or Page.SetDefaultTimeout()
	// methods.
	Timeout *float64 `json:"timeout"`
	// When to consider operation succeeded, defaults to `load`. Events can be either:
	// `'domcontentloaded'` - consider operation to be finished when the `DOMContentLoaded`
	// event is fired.
	// `'load'` - consider operation to be finished when the `load` event is fired.
	// `'networkidle'` - consider operation to be finished when there are no network connections
	// for at least `500` ms.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type PageGoForwardOptions struct {
	// Maximum operation time in milliseconds, defaults to 30 seconds, pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultNavigationTimeout(),
	// BrowserContext.SetDefaultTimeout(), Page.SetDefaultNavigationTimeout() or Page.SetDefaultTimeout()
	// methods.
	Timeout *float64 `json:"timeout"`
	// When to consider operation succeeded, defaults to `load`. Events can be either:
	// `'domcontentloaded'` - consider operation to be finished when the `DOMContentLoaded`
	// event is fired.
	// `'load'` - consider operation to be finished when the `load` event is fired.
	// `'networkidle'` - consider operation to be finished when there are no network connections
	// for at least `500` ms.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type PageGotoOptions struct {
	// Referer header value. If provided it will take preference over the referer header
	// value set by Page.SetExtraHttpHeaders().
	Referer *string `json:"referer"`
	// Maximum operation time in milliseconds, defaults to 30 seconds, pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultNavigationTimeout(),
	// BrowserContext.SetDefaultTimeout(), Page.SetDefaultNavigationTimeout() or Page.SetDefaultTimeout()
	// methods.
	Timeout *float64 `json:"timeout"`
	// When to consider operation succeeded, defaults to `load`. Events can be either:
	// `'domcontentloaded'` - consider operation to be finished when the `DOMContentLoaded`
	// event is fired.
	// `'load'` - consider operation to be finished when the `load` event is fired.
	// `'networkidle'` - consider operation to be finished when there are no network connections
	// for at least `500` ms.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type PageHoverOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *PageHoverOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type PageInnerHTMLOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageInnerTextOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageInputValueOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageIsCheckedOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageIsDisabledOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageIsEditableOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageIsEnabledOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageIsHiddenOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// **DEPRECATED** This option is ignored. Page.IsHidden() does not wait for the element
	// to become hidden and returns immediately.
	Timeout *float64 `json:"timeout"`
}
type PageIsVisibleOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// **DEPRECATED** This option is ignored. Page.IsVisible() does not wait for the element
	// to become visible and returns immediately.
	Timeout *float64 `json:"timeout"`
}
type PagePdfOptions struct {
	// Display header and footer. Defaults to `false`.
	DisplayHeaderFooter *bool `json:"displayHeaderFooter"`
	// HTML template for the print footer. Should use the same format as the `headerTemplate`.
	FooterTemplate *string `json:"footerTemplate"`
	// Paper format. If set, takes priority over `width` or `height` options. Defaults
	// to 'Letter'.
	Format *string `json:"format"`
	// HTML template for the print header. Should be valid HTML markup with following classes
	// used to inject printing values into them:
	// `'date'` formatted print date
	// `'title'` document title
	// `'url'` document location
	// `'pageNumber'` current page number
	// `'totalPages'` total pages in the document
	HeaderTemplate *string `json:"headerTemplate"`
	// Paper height, accepts values labeled with units.
	Height *string `json:"height"`
	// Paper orientation. Defaults to `false`.
	Landscape *bool `json:"landscape"`
	// Paper margins, defaults to none.
	Margin *PagePdfOptionsMargin `json:"margin"`
	// Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which
	// means print all pages.
	PageRanges *string `json:"pageRanges"`
	// The file path to save the PDF to. If `path` is a relative path, then it is resolved
	// relative to the current working directory. If no path is provided, the PDF won't
	// be saved to the disk.
	Path *string `json:"path"`
	// Give any CSS `@page` size declared in the page priority over what is declared in
	// `width` and `height` or `format` options. Defaults to `false`, which will scale
	// the content to fit the paper size.
	PreferCSSPageSize *bool `json:"preferCSSPageSize"`
	// Print background graphics. Defaults to `false`.
	PrintBackground *bool `json:"printBackground"`
	// Scale of the webpage rendering. Defaults to `1`. Scale amount must be between 0.1
	// and 2.
	Scale *float64 `json:"scale"`
	// Paper width, accepts values labeled with units.
	Width *string `json:"width"`
}
type PageMargin struct {
	// Top margin, accepts values labeled with units. Defaults to `0`.
	Top *string `json:"top"`
	// Right margin, accepts values labeled with units. Defaults to `0`.
	Right *string `json:"right"`
	// Bottom margin, accepts values labeled with units. Defaults to `0`.
	Bottom *string `json:"bottom"`
	// Left margin, accepts values labeled with units. Defaults to `0`.
	Left *string `json:"left"`
}
type PagePressOptions struct {
	// Time to wait between `keydown` and `keyup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageQuerySelectorOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
}
type PageReloadOptions struct {
	// Maximum operation time in milliseconds, defaults to 30 seconds, pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultNavigationTimeout(),
	// BrowserContext.SetDefaultTimeout(), Page.SetDefaultNavigationTimeout() or Page.SetDefaultTimeout()
	// methods.
	Timeout *float64 `json:"timeout"`
	// When to consider operation succeeded, defaults to `load`. Events can be either:
	// `'domcontentloaded'` - consider operation to be finished when the `DOMContentLoaded`
	// event is fired.
	// `'load'` - consider operation to be finished when the `load` event is fired.
	// `'networkidle'` - consider operation to be finished when there are no network connections
	// for at least `500` ms.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type PageRouteOptions struct {
	// How often a route should be used. By default it will be used every time.
	Times *int `json:"times"`
}
type PageScreenshotOptions struct {
	// When set to `"disabled"`, stops CSS animations, CSS transitions and Web Animations.
	// Animations get different treatment depending on their duration:
	// finite animations are fast-forwarded to completion, so they'll fire `transitionend`
	// event.
	// infinite animations are canceled to initial state, and then played over after the
	// screenshot.
	// Defaults to `"allow"` that leaves animations untouched.
	Animations *ScreenshotAnimations `json:"animations"`
	// An object which specifies clipping of the resulting image. Should have the following
	// fields:
	Clip *PageScreenshotOptionsClip `json:"clip"`
	// When set to `"ready"`, screenshot will wait for [`document.fonts.ready`](https://developer.mozilla.org/en-US/docs/Web/API/FontFaceSet/ready)
	// promise to resolve in all frames. Defaults to `"nowait"`.
	Fonts *ScreenshotFonts `json:"fonts"`
	// When true, takes a screenshot of the full scrollable page, instead of the currently
	// visible viewport. Defaults to `false`.
	FullPage *bool `json:"fullPage"`
	// Hides default white background and allows capturing screenshots with transparency.
	// Not applicable to `jpeg` images. Defaults to `false`.
	OmitBackground *bool `json:"omitBackground"`
	// The file path to save the image to. The screenshot type will be inferred from file
	// extension. If `path` is a relative path, then it is resolved relative to the current
	// working directory. If no path is provided, the image won't be saved to the disk.
	Path *string `json:"path"`
	// The quality of the image, between 0-100. Not applicable to `png` images.
	Quality *int `json:"quality"`
	// When set to `"css"`, screenshot will have a single pixel per each css pixel on the
	// page. For high-dpi devices, this will keep screenshots small. Using `"device"` option
	// will produce a single pixel per each device pixel, so screenhots of high-dpi devices
	// will be twice as large or even larger. Defaults to `"device"`.
	Size *ScreenshotSize `json:"size"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// Specify screenshot type, defaults to `png`.
	Type *ScreenshotType `json:"type"`
}
type PageClip struct {
	// x-coordinate of top-left corner of clip area
	X *float64 `json:"x"`
	// y-coordinate of top-left corner of clip area
	Y *float64 `json:"y"`
	// width of clipping area
	Width *float64 `json:"width"`
	// height of clipping area
	Height *float64 `json:"height"`
}
type PageSelectOptionOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageSetCheckedOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *PageSetCheckedOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type PageSetContentOptions struct {
	// Maximum operation time in milliseconds, defaults to 30 seconds, pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultNavigationTimeout(),
	// BrowserContext.SetDefaultTimeout(), Page.SetDefaultNavigationTimeout() or Page.SetDefaultTimeout()
	// methods.
	Timeout *float64 `json:"timeout"`
	// When to consider operation succeeded, defaults to `load`. Events can be either:
	// `'domcontentloaded'` - consider operation to be finished when the `DOMContentLoaded`
	// event is fired.
	// `'load'` - consider operation to be finished when the `load` event is fired.
	// `'networkidle'` - consider operation to be finished when there are no network connections
	// for at least `500` ms.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type PageSetInputFilesOptions struct {
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageTapOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *PageTapOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type PageTextContentOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageTypeOptions struct {
	// Time to wait between key presses in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageUncheckOptions struct {
	// Whether to bypass the [actionability](./actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *PageUncheckOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](./actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type PageUnrouteOptions struct {
	// Optional handler function to route the request.
	Handler func(Route, Request) `json:"handler"`
}

// Result of calling <see cref="Page.ViewportSize" />.
type PageViewportSizeResult struct {
	// page width in pixels.
	Width int `json:"width"`
	// page height in pixels.
	Height int `json:"height"`
}
type PageWaitForFunctionOptions struct {
	// If `polling` is `'raf'`, then `expression` is constantly executed in `requestAnimationFrame`
	// callback. If `polling` is a number, then it is treated as an interval in milliseconds
	// at which the function would be executed. Defaults to `raf`.
	Polling interface{} `json:"polling"`
	// maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type PageWaitForLoadStateOptions struct {
	// Maximum operation time in milliseconds, defaults to 30 seconds, pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultNavigationTimeout(),
	// BrowserContext.SetDefaultTimeout(), Page.SetDefaultNavigationTimeout() or Page.SetDefaultTimeout()
	// methods.
	Timeout *float64 `json:"timeout"`
}
type PageWaitForNavigationOptions struct {
	// Maximum operation time in milliseconds, defaults to 30 seconds, pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultNavigationTimeout(),
	// BrowserContext.SetDefaultTimeout(), Page.SetDefaultNavigationTimeout() or Page.SetDefaultTimeout()
	// methods.
	Timeout *float64 `json:"timeout"`
	// A glob pattern, regex pattern or predicate receiving [URL] to match while waiting
	// for the navigation. Note that if the parameter is a string without wilcard characters,
	// the method will wait for navigation to URL that is exactly equal to the string.
	URL interface{} `json:"url"`
	// When to consider operation succeeded, defaults to `load`. Events can be either:
	// `'domcontentloaded'` - consider operation to be finished when the `DOMContentLoaded`
	// event is fired.
	// `'load'` - consider operation to be finished when the `load` event is fired.
	// `'networkidle'` - consider operation to be finished when there are no network connections
	// for at least `500` ms.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type PageWaitForRequestOptions struct {
	// Maximum wait time in milliseconds, defaults to 30 seconds, pass `0` to disable the
	// timeout. The default value can be changed by using the Page.SetDefaultTimeout()
	// method.
	Timeout *float64 `json:"timeout"`
}
type PageWaitForResponseOptions struct {
	// Maximum wait time in milliseconds, defaults to 30 seconds, pass `0` to disable the
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageWaitForSelectorOptions struct {
	// Defaults to `'visible'`. Can be either:
	// `'attached'` - wait for element to be present in DOM.
	// `'detached'` - wait for element to not be present in DOM.
	// `'visible'` - wait for element to have non-empty bounding box and no `visibility:hidden`.
	// Note that element without any content or with `display:none` has an empty bounding
	// box and is not considered visible.
	// `'hidden'` - wait for element to be either detached from DOM, or have an empty bounding
	// box or `visibility:hidden`. This is opposite to the `'visible'` option.
	State *WaitForSelectorState `json:"state"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageWaitForURLOptions struct {
	// Maximum operation time in milliseconds, defaults to 30 seconds, pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultNavigationTimeout(),
	// BrowserContext.SetDefaultTimeout(), Page.SetDefaultNavigationTimeout() or Page.SetDefaultTimeout()
	// methods.
	Timeout *float64 `json:"timeout"`
	// When to consider operation succeeded, defaults to `load`. Events can be either:
	// `'domcontentloaded'` - consider operation to be finished when the `DOMContentLoaded`
	// event is fired.
	// `'load'` - consider operation to be finished when the `load` event is fired.
	// `'networkidle'` - consider operation to be finished when there are no network connections
	// for at least `500` ms.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}

// Result of calling <see cref="Request.HeadersArray" />.
type RequestHeadersArrayResult struct {
	// Name of the header.
	Name string `json:"name"`
	// Value of the header.
	Value string `json:"value"`
}

// Result of calling <see cref="Request.Sizes" />.
type RequestSizesResult struct {
	// Size of the request body (POST data payload) in bytes. Set to 0 if there was no
	// body.
	RequestBodySize int `json:"requestBodySize"`
	// Total number of bytes from the start of the HTTP request message until (and including)
	// the double CRLF before the body.
	RequestHeadersSize int `json:"requestHeadersSize"`
	// Size of the received response body (encoded) in bytes.
	ResponseBodySize int `json:"responseBodySize"`
	// Total number of bytes from the start of the HTTP response message until (and including)
	// the double CRLF before the body.
	ResponseHeadersSize int `json:"responseHeadersSize"`
}

// Result of calling <see cref="Request.Timing" />.
type RequestTimingResult struct {
	// Request start time in milliseconds elapsed since January 1, 1970 00:00:00 UTC
	StartTime float64 `json:"startTime"`
	// Time immediately before the browser starts the domain name lookup for the resource.
	// The value is given in milliseconds relative to `startTime`, -1 if not available.
	DomainLookupStart float64 `json:"domainLookupStart"`
	// Time immediately after the browser starts the domain name lookup for the resource.
	// The value is given in milliseconds relative to `startTime`, -1 if not available.
	DomainLookupEnd float64 `json:"domainLookupEnd"`
	// Time immediately before the user agent starts establishing the connection to the
	// server to retrieve the resource. The value is given in milliseconds relative to
	// `startTime`, -1 if not available.
	ConnectStart float64 `json:"connectStart"`
	// Time immediately before the browser starts the handshake process to secure the current
	// connection. The value is given in milliseconds relative to `startTime`, -1 if not
	// available.
	SecureConnectionStart float64 `json:"secureConnectionStart"`
	// Time immediately before the user agent starts establishing the connection to the
	// server to retrieve the resource. The value is given in milliseconds relative to
	// `startTime`, -1 if not available.
	ConnectEnd float64 `json:"connectEnd"`
	// Time immediately before the browser starts requesting the resource from the server,
	// cache, or local resource. The value is given in milliseconds relative to `startTime`,
	// -1 if not available.
	RequestStart float64 `json:"requestStart"`
	// Time immediately after the browser starts requesting the resource from the server,
	// cache, or local resource. The value is given in milliseconds relative to `startTime`,
	// -1 if not available.
	ResponseStart float64 `json:"responseStart"`
	// Time immediately after the browser receives the last byte of the resource or immediately
	// before the transport connection is closed, whichever comes first. The value is given
	// in milliseconds relative to `startTime`, -1 if not available.
	ResponseEnd float64 `json:"responseEnd"`
}

// Result of calling <see cref="Response.HeadersArray" />.
type ResponseHeadersArrayResult struct {
	// Name of the header.
	Name string `json:"name"`
	// Value of the header.
	Value string `json:"value"`
}

// Result of calling <see cref="Response.SecurityDetails" />.
type ResponseSecurityDetailsResult struct {
	// Common Name component of the Issuer field. from the certificate. This should only
	// be used for informational purposes. Optional.
	Issuer string `json:"issuer"`
	// The specific TLS protocol used. (e.g. `TLS 1.3`). Optional.
	Protocol string `json:"protocol"`
	// Common Name component of the Subject field from the certificate. This should only
	// be used for informational purposes. Optional.
	SubjectName string `json:"subjectName"`
	// Unix timestamp (in seconds) specifying when this cert becomes valid. Optional.
	ValidFrom float64 `json:"validFrom"`
	// Unix timestamp (in seconds) specifying when this cert becomes invalid. Optional.
	ValidTo float64 `json:"validTo"`
}

// Result of calling <see cref="Response.ServerAddr" />.
type ResponseServerAddrResult struct {
	// IPv4 or IPV6 address of the server.
	IpAddress string `json:"ipAddress"`
	Port      int    `json:"port"`
}
type RouteAbortOptions struct {
	// Optional error code. Defaults to `failed`, could be one of the following:
	// `'aborted'` - An operation was aborted (due to user action)
	// `'accessdenied'` - Permission to access a resource, other than the network, was
	// denied
	// `'addressunreachable'` - The IP address is unreachable. This usually means that
	// there is no route to the specified host or network.
	// `'blockedbyclient'` - The client chose to block the request.
	// `'blockedbyresponse'` - The request failed because the response was delivered along
	// with requirements which are not met ('X-Frame-Options' and 'Content-Security-Policy'
	// ancestor checks, for instance).
	// `'connectionaborted'` - A connection timed out as a result of not receiving an ACK
	// for data sent.
	// `'connectionclosed'` - A connection was closed (corresponding to a TCP FIN).
	// `'connectionfailed'` - A connection attempt failed.
	// `'connectionrefused'` - A connection attempt was refused.
	// `'connectionreset'` - A connection was reset (corresponding to a TCP RST).
	// `'internetdisconnected'` - The Internet connection has been lost.
	// `'namenotresolved'` - The host name could not be resolved.
	// `'timedout'` - An operation timed out.
	// `'failed'` - A generic failure occurred.
	ErrorCode *string `json:"errorCode"`
}
type RouteContinueOptions struct {
	// If set changes the request HTTP headers. Header values will be converted to a string.
	Headers map[string]string `json:"headers"`
	// If set changes the request method (e.g. GET or POST)
	Method *string `json:"method"`
	// If set changes the post data of request
	PostData interface{} `json:"postData"`
	// If set changes the request URL. New URL must have same protocol as original one.
	URL *string `json:"url"`
}
type RouteFulfillOptions struct {
	// Response body.
	Body interface{} `json:"body"`
	// If set, equals to setting `Content-Type` response header.
	ContentType *string `json:"contentType"`
	// Response headers. Header values will be converted to a string.
	Headers map[string]string `json:"headers"`
	// File path to respond with. The content type will be inferred from file extension.
	// If `path` is a relative path, then it is resolved relative to the current working
	// directory.
	Path *string `json:"path"`
	// Response status code, defaults to `200`.
	Status *int `json:"status"`
}
type SelectorsRegisterOptions struct {
	// Whether to run this selector engine in isolated JavaScript environment. This environment
	// has access to the same DOM, but not any JavaScript objects from the frame's scripts.
	// Defaults to `false`. Note that running as a content script is not guaranteed when
	// this engine is used together with other registered engines.
	ContentScript *bool `json:"contentScript"`
}
type TracingStartOptions struct {
	// If specified, the trace is going to be saved into the file with the given name inside
	// the `tracesDir` folder specified in BrowserType.Launch().
	Name *string `json:"name"`
	// Whether to capture screenshots during tracing. Screenshots are used to build a timeline
	// preview.
	Screenshots *bool `json:"screenshots"`
	// If this option is true tracing will
	// capture DOM snapshot on every action
	// record network activity
	Snapshots *bool `json:"snapshots"`
	// Trace name to be shown in the Trace Viewer.
	Title *string `json:"title"`
}
type TracingStartChunkOptions struct {
	// Trace name to be shown in the Trace Viewer.
	Title *string `json:"title"`
}
type TracingStopOptions struct {
	// Export trace into the file with the given path.
	Path *string `json:"path"`
}
type TracingStopChunkOptions struct {
	// Export trace collected since the last Tracing.StartChunk() call into the file with
	// the given path.
	Path *string `json:"path"`
}
type FrameReceivedPayload struct {
	// frame payload
	Payload []byte `json:"payload"`
}
type FrameSentPayload struct {
	// frame payload
	Payload []byte `json:"payload"`
}
type WorkerEvaluateOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type WorkerEvaluateHandleOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type BrowserNewContextOptionsGeolocation struct {
	// Latitude between -90 and 90.
	Latitude *float64 `json:"latitude"`
	// Longitude between -180 and 180.
	Longitude *float64 `json:"longitude"`
	// Non-negative accuracy value. Defaults to `0`.
	Accuracy *float64 `json:"accuracy"`
}
type BrowserNewContextOptionsHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type BrowserNewContextOptionsProxy struct {
	// Proxy to be used for all requests. HTTP and SOCKS proxies are supported, for example
	// `http://myproxy.com:3128` or `socks5://myproxy.com:3128`. Short form `myproxy.com:3128`
	// is considered an HTTP proxy.
	Server *string `json:"server"`
	// Optional comma-separated domains to bypass proxy, for example `".com, chromium.org,
	// .domain.com"`.
	Bypass *string `json:"bypass"`
	// Optional username to use if HTTP proxy requires authentication.
	Username *string `json:"username"`
	// Optional password to use if HTTP proxy requires authentication.
	Password *string `json:"password"`
}
type BrowserNewContextOptionsRecordVideo struct {
	// Path to the directory to put videos into.
	Dir *string `json:"dir"`
	// Optional dimensions of the recorded videos. If not specified the size will be equal
	// to `viewport` scaled down to fit into 800x800. If `viewport` is not configured explicitly
	// the video size defaults to 800x450. Actual picture of each page will be scaled down
	// if necessary to fit the specified size.
	Size *BrowserNewContextOptionsRecordVideoSize `json:"size"`
}
type BrowserNewContextOptionsScreen struct {
	// page width in pixels.
	Width *int `json:"width"`
	// page height in pixels.
	Height *int `json:"height"`
}
type BrowserNewContextOptionsStorageState struct {
	// cookies to set for context
	Cookies []BrowserNewContextOptionsStorageStateCookies `json:"cookies"`
	// localStorage to set for context
	Origins []BrowserNewContextOptionsStorageStateOrigins `json:"origins"`
}
type BrowserNewContextOptionsViewport struct {
	// page width in pixels.
	Width *int `json:"width"`
	// page height in pixels.
	Height *int `json:"height"`
}
type BrowserRecordVideoSize struct {
	// Video frame width.
	Width *int `json:"width"`
	// Video frame height.
	Height *int `json:"height"`
}
type BrowserStorageStateCookies struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
	// domain and path are required
	Domain *string `json:"domain"`
	// domain and path are required
	Path *string `json:"path"`
	// Unix time in seconds.
	Expires  *float64 `json:"expires"`
	HttpOnly *bool    `json:"httpOnly"`
	Secure   *bool    `json:"secure"`
	// sameSite flag
	SameSite *SameSiteAttribute `json:"sameSite"`
}
type BrowserStorageStateOrigins struct {
	Origin       *string                                  `json:"origin"`
	LocalStorage []BrowserStorageStateOriginsLocalStorage `json:"localStorage"`
}
type BrowserNewPageOptionsGeolocation struct {
	// Latitude between -90 and 90.
	Latitude *float64 `json:"latitude"`
	// Longitude between -180 and 180.
	Longitude *float64 `json:"longitude"`
	// Non-negative accuracy value. Defaults to `0`.
	Accuracy *float64 `json:"accuracy"`
}
type BrowserNewPageOptionsHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type BrowserNewPageOptionsProxy struct {
	// Proxy to be used for all requests. HTTP and SOCKS proxies are supported, for example
	// `http://myproxy.com:3128` or `socks5://myproxy.com:3128`. Short form `myproxy.com:3128`
	// is considered an HTTP proxy.
	Server *string `json:"server"`
	// Optional comma-separated domains to bypass proxy, for example `".com, chromium.org,
	// .domain.com"`.
	Bypass *string `json:"bypass"`
	// Optional username to use if HTTP proxy requires authentication.
	Username *string `json:"username"`
	// Optional password to use if HTTP proxy requires authentication.
	Password *string `json:"password"`
}
type BrowserNewPageOptionsRecordVideo struct {
	// Path to the directory to put videos into.
	Dir *string `json:"dir"`
	// Optional dimensions of the recorded videos. If not specified the size will be equal
	// to `viewport` scaled down to fit into 800x800. If `viewport` is not configured explicitly
	// the video size defaults to 800x450. Actual picture of each page will be scaled down
	// if necessary to fit the specified size.
	Size *BrowserNewPageOptionsRecordVideoSize `json:"size"`
}
type BrowserNewPageOptionsScreen struct {
	// page width in pixels.
	Width *int `json:"width"`
	// page height in pixels.
	Height *int `json:"height"`
}
type BrowserNewPageOptionsStorageState struct {
	// cookies to set for context
	Cookies []BrowserNewPageOptionsStorageStateCookies `json:"cookies"`
	// localStorage to set for context
	Origins []BrowserNewPageOptionsStorageStateOrigins `json:"origins"`
}
type BrowserNewPageOptionsViewport struct {
	// page width in pixels.
	Width *int `json:"width"`
	// page height in pixels.
	Height *int `json:"height"`
}
type BrowserContextAddCookiesOptionsCookies struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
	// either url or domain / path are required. Optional.
	URL *string `json:"url"`
	// either url or domain / path are required Optional.
	Domain *string `json:"domain"`
	// either url or domain / path are required Optional.
	Path *string `json:"path"`
	// Unix time in seconds. Optional.
	Expires *float64 `json:"expires"`
	// Optional.
	HttpOnly *bool `json:"httpOnly"`
	// Optional.
	Secure *bool `json:"secure"`
	// Optional.
	SameSite *SameSiteAttribute `json:"sameSite"`
}
type BrowserContextStorageStateResultCookies struct {
	Name   *string `json:"name"`
	Value  *string `json:"value"`
	Domain *string `json:"domain"`
	Path   *string `json:"path"`
	// Unix time in seconds.
	Expires  *float64           `json:"expires"`
	HttpOnly *bool              `json:"httpOnly"`
	Secure   *bool              `json:"secure"`
	SameSite *SameSiteAttribute `json:"sameSite"`
}
type BrowserContextStorageStateResultOrigins struct {
	Origin       *string                                               `json:"origin"`
	LocalStorage []BrowserContextStorageStateResultOriginsLocalStorage `json:"localStorage"`
}
type BrowserTypeLaunchOptionsProxy struct {
	// Proxy to be used for all requests. HTTP and SOCKS proxies are supported, for example
	// `http://myproxy.com:3128` or `socks5://myproxy.com:3128`. Short form `myproxy.com:3128`
	// is considered an HTTP proxy.
	Server *string `json:"server"`
	// Optional comma-separated domains to bypass proxy, for example `".com, chromium.org,
	// .domain.com"`.
	Bypass *string `json:"bypass"`
	// Optional username to use if HTTP proxy requires authentication.
	Username *string `json:"username"`
	// Optional password to use if HTTP proxy requires authentication.
	Password *string `json:"password"`
}
type BrowserTypeLaunchPersistentContextOptionsGeolocation struct {
	// Latitude between -90 and 90.
	Latitude *float64 `json:"latitude"`
	// Longitude between -180 and 180.
	Longitude *float64 `json:"longitude"`
	// Non-negative accuracy value. Defaults to `0`.
	Accuracy *float64 `json:"accuracy"`
}
type BrowserTypeLaunchPersistentContextOptionsHttpCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type BrowserTypeLaunchPersistentContextOptionsProxy struct {
	// Proxy to be used for all requests. HTTP and SOCKS proxies are supported, for example
	// `http://myproxy.com:3128` or `socks5://myproxy.com:3128`. Short form `myproxy.com:3128`
	// is considered an HTTP proxy.
	Server *string `json:"server"`
	// Optional comma-separated domains to bypass proxy, for example `".com, chromium.org,
	// .domain.com"`.
	Bypass *string `json:"bypass"`
	// Optional username to use if HTTP proxy requires authentication.
	Username *string `json:"username"`
	// Optional password to use if HTTP proxy requires authentication.
	Password *string `json:"password"`
}
type BrowserTypeLaunchPersistentContextOptionsRecordVideo struct {
	// Path to the directory to put videos into.
	Dir *string `json:"dir"`
	// Optional dimensions of the recorded videos. If not specified the size will be equal
	// to `viewport` scaled down to fit into 800x800. If `viewport` is not configured explicitly
	// the video size defaults to 800x450. Actual picture of each page will be scaled down
	// if necessary to fit the specified size.
	Size *BrowserTypeLaunchPersistentContextOptionsRecordVideoSize `json:"size"`
}
type BrowserTypeLaunchPersistentContextOptionsScreen struct {
	// page width in pixels.
	Width *int `json:"width"`
	// page height in pixels.
	Height *int `json:"height"`
}
type BrowserTypeLaunchPersistentContextOptionsViewport struct {
	// page width in pixels.
	Width *int `json:"width"`
	// page height in pixels.
	Height *int `json:"height"`
}
type BrowserTypeRecordVideoSize struct {
	// Video frame width.
	Width *int `json:"width"`
	// Video frame height.
	Height *int `json:"height"`
}
type ElementHandleCheckOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type ElementHandleClickOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type ElementHandleDblclickOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type ElementHandleHoverOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type ElementHandleSetCheckedOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type ElementHandleTapOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type ElementHandleUncheckOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type FrameCheckOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type FrameClickOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type FrameDblclickOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type FrameDragAndDropOptionsSourcePosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type FrameDragAndDropOptionsTargetPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type FrameHoverOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type FrameSetCheckedOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type FrameTapOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type FrameUncheckOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type LocatorCheckOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type LocatorClickOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type LocatorDblclickOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type LocatorDragToOptionsSourcePosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type LocatorDragToOptionsTargetPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type LocatorHoverOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type LocatorSetCheckedOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type LocatorTapOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type LocatorUncheckOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type PageCheckOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type PageClickOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type PageDblclickOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type PageDragAndDropOptionsSourcePosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type PageDragAndDropOptionsTargetPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type PageHoverOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type PagePdfOptionsMargin struct {
	// Top margin, accepts values labeled with units. Defaults to `0`.
	Top *string `json:"top"`
	// Right margin, accepts values labeled with units. Defaults to `0`.
	Right *string `json:"right"`
	// Bottom margin, accepts values labeled with units. Defaults to `0`.
	Bottom *string `json:"bottom"`
	// Left margin, accepts values labeled with units. Defaults to `0`.
	Left *string `json:"left"`
}
type PageScreenshotOptionsClip struct {
	// x-coordinate of top-left corner of clip area
	X *float64 `json:"x"`
	// y-coordinate of top-left corner of clip area
	Y *float64 `json:"y"`
	// width of clipping area
	Width *float64 `json:"width"`
	// height of clipping area
	Height *float64 `json:"height"`
}
type PageSetCheckedOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type PageTapOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type PageUncheckOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}
type BrowserNewContextOptionsRecordVideoSize struct {
	// Video frame width.
	Width *int `json:"width"`
	// Video frame height.
	Height *int `json:"height"`
}
type BrowserNewContextOptionsStorageStateCookies struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
	// domain and path are required
	Domain *string `json:"domain"`
	// domain and path are required
	Path *string `json:"path"`
	// Unix time in seconds.
	Expires  *float64 `json:"expires"`
	HttpOnly *bool    `json:"httpOnly"`
	Secure   *bool    `json:"secure"`
	// sameSite flag
	SameSite *SameSiteAttribute `json:"sameSite"`
}
type BrowserNewContextOptionsStorageStateOrigins struct {
	Origin       *string                                                   `json:"origin"`
	LocalStorage []BrowserNewContextOptionsStorageStateOriginsLocalStorage `json:"localStorage"`
}
type BrowserStorageStateOriginsLocalStorage struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}
type BrowserNewPageOptionsRecordVideoSize struct {
	// Video frame width.
	Width *int `json:"width"`
	// Video frame height.
	Height *int `json:"height"`
}
type BrowserNewPageOptionsStorageStateCookies struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
	// domain and path are required
	Domain *string `json:"domain"`
	// domain and path are required
	Path *string `json:"path"`
	// Unix time in seconds.
	Expires  *float64 `json:"expires"`
	HttpOnly *bool    `json:"httpOnly"`
	Secure   *bool    `json:"secure"`
	// sameSite flag
	SameSite *SameSiteAttribute `json:"sameSite"`
}
type BrowserNewPageOptionsStorageStateOrigins struct {
	Origin       *string                                                `json:"origin"`
	LocalStorage []BrowserNewPageOptionsStorageStateOriginsLocalStorage `json:"localStorage"`
}
type BrowserContextStorageStateResultOriginsLocalStorage struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}
type BrowserTypeLaunchPersistentContextOptionsRecordVideoSize struct {
	// Video frame width.
	Width *int `json:"width"`
	// Video frame height.
	Height *int `json:"height"`
}
type BrowserNewContextOptionsStorageStateOriginsLocalStorage struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}
type BrowserNewPageOptionsStorageStateOriginsLocalStorage struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}
