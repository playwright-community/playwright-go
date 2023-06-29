package playwright

type APIRequestNewContextOptions struct {
	// Methods like APIRequestContext.Get() take the base URL into consideration by using
	// the [`URL()`](https://developer.mozilla.org/en-US/docs/Web/API/URL/URL) constructor
	// for building the corresponding URL. Examples:
	// baseURL: `http://localhost:3000` and sending request to `/bar.html` results in `http://localhost:3000/bar.html`
	// baseURL: `http://localhost:3000/foo/` and sending request to `./bar.html` results
	// in `http://localhost:3000/foo/bar.html`
	// baseURL: `http://localhost:3000/foo` (without trailing slash) and navigating to
	// `./bar.html` results in `http://localhost:3000/bar.html`
	BaseURL *string `json:"baseURL"`
	// An object containing additional HTTP headers to be sent with every request. Defaults
	// to none.
	ExtraHttpHeaders map[string]string `json:"extraHTTPHeaders"`
	// Credentials for [HTTP authentication](https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication).
	// If no origin is specified, the username and password are sent to any servers upon
	// unauthorized responses.
	HttpCredentials *HttpCredentials `json:"httpCredentials"`
	// Whether to ignore HTTPS errors when sending network requests. Defaults to `false`.
	IgnoreHttpsErrors *bool `json:"ignoreHTTPSErrors"`
	// Network proxy settings.
	Proxy *Proxy `json:"proxy"`
	// Populates context with given storage state. This option can be used to initialize
	// context with logged-in information obtained via BrowserContext.StorageState() or
	// APIRequestContext.StorageState(). Either a path to the file with saved storage,
	// or the value returned by one of BrowserContext.StorageState() or APIRequestContext.StorageState()
	// methods.
	StorageState *StorageState `json:"storageState"`
	// Populates context with given storage state. This option can be used to initialize
	// context with logged-in information obtained via BrowserContext.StorageState(). Path
	// to the file with saved storage state.
	StorageStatePath *string `json:"storageStatePath"`
	// Maximum time in milliseconds to wait for the response. Defaults to `30000` (30 seconds).
	// Pass `0` to disable timeout.
	Timeout *float64 `json:"timeout"`
	// Specific user agent to use in this context.
	UserAgent *string `json:"userAgent"`
}
type Proxy struct {
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
type APIRequestContextDeleteOptions struct {
	// Allows to set post data of the request. If the data parameter is an object, it will
	// be serialized to json string and `content-type` header will be set to `application/json`
	// if not explicitly set. Otherwise the `content-type` header will be set to `application/octet-stream`
	// if not explicitly set.
	Data interface{} `json:"data"`
	// Whether to throw on response codes other than 2xx and 3xx. By default response object
	// is returned for all status codes.
	FailOnStatusCode *bool `json:"failOnStatusCode"`
	// Provides an object that will be serialized as html form using `application/x-www-form-urlencoded`
	// encoding and sent as this request body. If this parameter is specified `content-type`
	// header will be set to `application/x-www-form-urlencoded` unless explicitly provided.
	Form interface{} `json:"form"`
	// Allows to set HTTP headers. These headers will apply to the fetched request as well
	// as any redirects initiated by it.
	Headers map[string]string `json:"headers"`
	// Whether to ignore HTTPS errors when sending network requests. Defaults to `false`.
	IgnoreHttpsErrors *bool `json:"ignoreHTTPSErrors"`
	// Maximum number of request redirects that will be followed automatically. An error
	// will be thrown if the number is exceeded. Defaults to `20`. Pass `0` to not follow
	// redirects.
	MaxRedirects *int `json:"maxRedirects"`
	// Provides an object that will be serialized as html form using `multipart/form-data`
	// encoding and sent as this request body. If this parameter is specified `content-type`
	// header will be set to `multipart/form-data` unless explicitly provided. File values
	// can be passed either as [`fs.ReadStream`](https://nodejs.org/api/fs.html#fs_class_fs_readstream)
	// or as file-like object containing file name, mime-type and its content.
	Multipart interface{} `json:"multipart"`
	// Query parameters to be sent with the URL.
	Params map[string]interface{} `json:"params"`
	// Request timeout in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout.
	Timeout *float64 `json:"timeout"`
}
type APIRequestContextFetchOptions struct {
	// Allows to set post data of the request. If the data parameter is an object, it will
	// be serialized to json string and `content-type` header will be set to `application/json`
	// if not explicitly set. Otherwise the `content-type` header will be set to `application/octet-stream`
	// if not explicitly set.
	Data interface{} `json:"data"`
	// Whether to throw on response codes other than 2xx and 3xx. By default response object
	// is returned for all status codes.
	FailOnStatusCode *bool `json:"failOnStatusCode"`
	// Provides an object that will be serialized as html form using `application/x-www-form-urlencoded`
	// encoding and sent as this request body. If this parameter is specified `content-type`
	// header will be set to `application/x-www-form-urlencoded` unless explicitly provided.
	Form interface{} `json:"form"`
	// Allows to set HTTP headers. These headers will apply to the fetched request as well
	// as any redirects initiated by it.
	Headers map[string]string `json:"headers"`
	// Whether to ignore HTTPS errors when sending network requests. Defaults to `false`.
	IgnoreHttpsErrors *bool `json:"ignoreHTTPSErrors"`
	// Maximum number of request redirects that will be followed automatically. An error
	// will be thrown if the number is exceeded. Defaults to `20`. Pass `0` to not follow
	// redirects.
	MaxRedirects *int `json:"maxRedirects"`
	// If set changes the fetch method (e.g. [PUT](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PUT)
	// or [POST](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST)). If not
	// specified, GET method is used.
	Method *string `json:"method"`
	// Provides an object that will be serialized as html form using `multipart/form-data`
	// encoding and sent as this request body. If this parameter is specified `content-type`
	// header will be set to `multipart/form-data` unless explicitly provided. File values
	// can be passed either as [`fs.ReadStream`](https://nodejs.org/api/fs.html#fs_class_fs_readstream)
	// or as file-like object containing file name, mime-type and its content.
	Multipart interface{} `json:"multipart"`
	// Query parameters to be sent with the URL.
	Params map[string]interface{} `json:"params"`
	// Request timeout in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout.
	Timeout *float64 `json:"timeout"`
}
type APIRequestContextGetOptions struct {
	// Allows to set post data of the request. If the data parameter is an object, it will
	// be serialized to json string and `content-type` header will be set to `application/json`
	// if not explicitly set. Otherwise the `content-type` header will be set to `application/octet-stream`
	// if not explicitly set.
	Data interface{} `json:"data"`
	// Whether to throw on response codes other than 2xx and 3xx. By default response object
	// is returned for all status codes.
	FailOnStatusCode *bool `json:"failOnStatusCode"`
	// Provides an object that will be serialized as html form using `application/x-www-form-urlencoded`
	// encoding and sent as this request body. If this parameter is specified `content-type`
	// header will be set to `application/x-www-form-urlencoded` unless explicitly provided.
	Form interface{} `json:"form"`
	// Allows to set HTTP headers. These headers will apply to the fetched request as well
	// as any redirects initiated by it.
	Headers map[string]string `json:"headers"`
	// Whether to ignore HTTPS errors when sending network requests. Defaults to `false`.
	IgnoreHttpsErrors *bool `json:"ignoreHTTPSErrors"`
	// Maximum number of request redirects that will be followed automatically. An error
	// will be thrown if the number is exceeded. Defaults to `20`. Pass `0` to not follow
	// redirects.
	MaxRedirects *int `json:"maxRedirects"`
	// Provides an object that will be serialized as html form using `multipart/form-data`
	// encoding and sent as this request body. If this parameter is specified `content-type`
	// header will be set to `multipart/form-data` unless explicitly provided. File values
	// can be passed either as [`fs.ReadStream`](https://nodejs.org/api/fs.html#fs_class_fs_readstream)
	// or as file-like object containing file name, mime-type and its content.
	Multipart interface{} `json:"multipart"`
	// Query parameters to be sent with the URL.
	Params map[string]interface{} `json:"params"`
	// Request timeout in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout.
	Timeout *float64 `json:"timeout"`
}
type APIRequestContextHeadOptions struct {
	// Allows to set post data of the request. If the data parameter is an object, it will
	// be serialized to json string and `content-type` header will be set to `application/json`
	// if not explicitly set. Otherwise the `content-type` header will be set to `application/octet-stream`
	// if not explicitly set.
	Data interface{} `json:"data"`
	// Whether to throw on response codes other than 2xx and 3xx. By default response object
	// is returned for all status codes.
	FailOnStatusCode *bool `json:"failOnStatusCode"`
	// Provides an object that will be serialized as html form using `application/x-www-form-urlencoded`
	// encoding and sent as this request body. If this parameter is specified `content-type`
	// header will be set to `application/x-www-form-urlencoded` unless explicitly provided.
	Form interface{} `json:"form"`
	// Allows to set HTTP headers. These headers will apply to the fetched request as well
	// as any redirects initiated by it.
	Headers map[string]string `json:"headers"`
	// Whether to ignore HTTPS errors when sending network requests. Defaults to `false`.
	IgnoreHttpsErrors *bool `json:"ignoreHTTPSErrors"`
	// Maximum number of request redirects that will be followed automatically. An error
	// will be thrown if the number is exceeded. Defaults to `20`. Pass `0` to not follow
	// redirects.
	MaxRedirects *int `json:"maxRedirects"`
	// Provides an object that will be serialized as html form using `multipart/form-data`
	// encoding and sent as this request body. If this parameter is specified `content-type`
	// header will be set to `multipart/form-data` unless explicitly provided. File values
	// can be passed either as [`fs.ReadStream`](https://nodejs.org/api/fs.html#fs_class_fs_readstream)
	// or as file-like object containing file name, mime-type and its content.
	Multipart interface{} `json:"multipart"`
	// Query parameters to be sent with the URL.
	Params map[string]interface{} `json:"params"`
	// Request timeout in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout.
	Timeout *float64 `json:"timeout"`
}
type APIRequestContextPatchOptions struct {
	// Allows to set post data of the request. If the data parameter is an object, it will
	// be serialized to json string and `content-type` header will be set to `application/json`
	// if not explicitly set. Otherwise the `content-type` header will be set to `application/octet-stream`
	// if not explicitly set.
	Data interface{} `json:"data"`
	// Whether to throw on response codes other than 2xx and 3xx. By default response object
	// is returned for all status codes.
	FailOnStatusCode *bool `json:"failOnStatusCode"`
	// Provides an object that will be serialized as html form using `application/x-www-form-urlencoded`
	// encoding and sent as this request body. If this parameter is specified `content-type`
	// header will be set to `application/x-www-form-urlencoded` unless explicitly provided.
	Form interface{} `json:"form"`
	// Allows to set HTTP headers. These headers will apply to the fetched request as well
	// as any redirects initiated by it.
	Headers map[string]string `json:"headers"`
	// Whether to ignore HTTPS errors when sending network requests. Defaults to `false`.
	IgnoreHttpsErrors *bool `json:"ignoreHTTPSErrors"`
	// Maximum number of request redirects that will be followed automatically. An error
	// will be thrown if the number is exceeded. Defaults to `20`. Pass `0` to not follow
	// redirects.
	MaxRedirects *int `json:"maxRedirects"`
	// Provides an object that will be serialized as html form using `multipart/form-data`
	// encoding and sent as this request body. If this parameter is specified `content-type`
	// header will be set to `multipart/form-data` unless explicitly provided. File values
	// can be passed either as [`fs.ReadStream`](https://nodejs.org/api/fs.html#fs_class_fs_readstream)
	// or as file-like object containing file name, mime-type and its content.
	Multipart interface{} `json:"multipart"`
	// Query parameters to be sent with the URL.
	Params map[string]interface{} `json:"params"`
	// Request timeout in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout.
	Timeout *float64 `json:"timeout"`
}
type APIRequestContextPostOptions struct {
	// Allows to set post data of the request. If the data parameter is an object, it will
	// be serialized to json string and `content-type` header will be set to `application/json`
	// if not explicitly set. Otherwise the `content-type` header will be set to `application/octet-stream`
	// if not explicitly set.
	Data interface{} `json:"data"`
	// Whether to throw on response codes other than 2xx and 3xx. By default response object
	// is returned for all status codes.
	FailOnStatusCode *bool `json:"failOnStatusCode"`
	// Provides an object that will be serialized as html form using `application/x-www-form-urlencoded`
	// encoding and sent as this request body. If this parameter is specified `content-type`
	// header will be set to `application/x-www-form-urlencoded` unless explicitly provided.
	Form interface{} `json:"form"`
	// Allows to set HTTP headers. These headers will apply to the fetched request as well
	// as any redirects initiated by it.
	Headers map[string]string `json:"headers"`
	// Whether to ignore HTTPS errors when sending network requests. Defaults to `false`.
	IgnoreHttpsErrors *bool `json:"ignoreHTTPSErrors"`
	// Maximum number of request redirects that will be followed automatically. An error
	// will be thrown if the number is exceeded. Defaults to `20`. Pass `0` to not follow
	// redirects.
	MaxRedirects *int `json:"maxRedirects"`
	// Provides an object that will be serialized as html form using `multipart/form-data`
	// encoding and sent as this request body. If this parameter is specified `content-type`
	// header will be set to `multipart/form-data` unless explicitly provided. File values
	// can be passed either as [`fs.ReadStream`](https://nodejs.org/api/fs.html#fs_class_fs_readstream)
	// or as file-like object containing file name, mime-type and its content.
	Multipart interface{} `json:"multipart"`
	// Query parameters to be sent with the URL.
	Params map[string]interface{} `json:"params"`
	// Request timeout in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout.
	Timeout *float64 `json:"timeout"`
}
type APIRequestContextPutOptions struct {
	// Allows to set post data of the request. If the data parameter is an object, it will
	// be serialized to json string and `content-type` header will be set to `application/json`
	// if not explicitly set. Otherwise the `content-type` header will be set to `application/octet-stream`
	// if not explicitly set.
	Data interface{} `json:"data"`
	// Whether to throw on response codes other than 2xx and 3xx. By default response object
	// is returned for all status codes.
	FailOnStatusCode *bool `json:"failOnStatusCode"`
	// Provides an object that will be serialized as html form using `application/x-www-form-urlencoded`
	// encoding and sent as this request body. If this parameter is specified `content-type`
	// header will be set to `application/x-www-form-urlencoded` unless explicitly provided.
	Form interface{} `json:"form"`
	// Allows to set HTTP headers. These headers will apply to the fetched request as well
	// as any redirects initiated by it.
	Headers map[string]string `json:"headers"`
	// Whether to ignore HTTPS errors when sending network requests. Defaults to `false`.
	IgnoreHttpsErrors *bool `json:"ignoreHTTPSErrors"`
	// Maximum number of request redirects that will be followed automatically. An error
	// will be thrown if the number is exceeded. Defaults to `20`. Pass `0` to not follow
	// redirects.
	MaxRedirects *int `json:"maxRedirects"`
	// Provides an object that will be serialized as html form using `multipart/form-data`
	// encoding and sent as this request body. If this parameter is specified `content-type`
	// header will be set to `multipart/form-data` unless explicitly provided. File values
	// can be passed either as [`fs.ReadStream`](https://nodejs.org/api/fs.html#fs_class_fs_readstream)
	// or as file-like object containing file name, mime-type and its content.
	Multipart interface{} `json:"multipart"`
	// Query parameters to be sent with the URL.
	Params map[string]interface{} `json:"params"`
	// Request timeout in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout.
	Timeout *float64 `json:"timeout"`
}
type APIRequestContextStorageStateOptions struct {
	// The file path to save the storage state to. If `path` is a relative path, then it
	// is resolved relative to current working directory. If no path is provided, storage
	// state is still returned, but won't be saved to the disk.
	Path *string `json:"path"`
}
type BrowserNewContextOptions struct {
	// Whether to automatically download all the attachments. Defaults to `true` where
	// all the downloads are accepted.
	AcceptDownloads *bool `json:"acceptDownloads"`
	// When using Page.Goto(), Page.Route(), Page.WaitForURL(), Page.WaitForRequest(),
	// or Page.WaitForResponse() it takes the base URL in consideration by using the [`URL()`](https://developer.mozilla.org/en-US/docs/Web/API/URL/URL)
	// constructor for building the corresponding URL. Unset by default. Examples:
	// baseURL: `http://localhost:3000` and navigating to `/bar.html` results in `http://localhost:3000/bar.html`
	// baseURL: `http://localhost:3000/foo/` and navigating to `./bar.html` results in
	// `http://localhost:3000/foo/bar.html`
	// baseURL: `http://localhost:3000/foo` (without trailing slash) and navigating to
	// `./bar.html` results in `http://localhost:3000/bar.html`
	BaseURL *string `json:"baseURL"`
	// Toggles bypassing page's Content-Security-Policy. Defaults to `false`.
	BypassCSP *bool `json:"bypassCSP"`
	// Emulates `'prefers-colors-scheme'` media feature, supported values are `'light'`,
	// `'dark'`, `'no-preference'`. See Page.EmulateMedia() for more details. Passing `'no-override'`
	// resets emulation to system defaults. Defaults to `'light'`.
	ColorScheme *ColorScheme `json:"colorScheme"`
	// Specify device scale factor (can be thought of as dpr). Defaults to `1`. Learn more
	// about [emulating devices with device scale factor](../emulation.md#devices).
	DeviceScaleFactor *float64 `json:"deviceScaleFactor"`
	// An object containing additional HTTP headers to be sent with every request. Defaults
	// to none.
	ExtraHttpHeaders map[string]string `json:"extraHTTPHeaders"`
	// Emulates `'forced-colors'` media feature, supported values are `'active'`, `'none'`.
	// See Page.EmulateMedia() for more details. Passing `'no-override'` resets emulation
	// to system defaults. Defaults to `'none'`.
	ForcedColors *ForcedColors `json:"forcedColors"`
	Geolocation  *Geolocation  `json:"geolocation"`
	// Specifies if viewport supports touch events. Defaults to false. Learn more about
	// [mobile emulation](../emulation.md#devices).
	HasTouch *bool `json:"hasTouch"`
	// Credentials for [HTTP authentication](https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication).
	// If no origin is specified, the username and password are sent to any servers upon
	// unauthorized responses.
	HttpCredentials *HttpCredentials `json:"httpCredentials"`
	// Whether to ignore HTTPS errors when sending network requests. Defaults to `false`.
	IgnoreHttpsErrors *bool `json:"ignoreHTTPSErrors"`
	// Whether the `meta viewport` tag is taken into account and touch events are enabled.
	// isMobile is a part of device, so you don't actually need to set it manually. Defaults
	// to `false` and is not supported in Firefox. Learn more about [mobile emulation](../emulation.md#isMobile).
	IsMobile *bool `json:"isMobile"`
	// Whether or not to enable JavaScript in the context. Defaults to `true`. Learn more
	// about [disabling JavaScript](../emulation.md#javascript-enabled).
	JavaScriptEnabled *bool `json:"javaScriptEnabled"`
	// Specify user locale, for example `en-GB`, `de-DE`, etc. Locale will affect `navigator.language`
	// value, `Accept-Language` request header value as well as number and date formatting
	// rules. Defaults to the system default locale. Learn more about emulation in our
	// [emulation guide](../emulation.md#locale--timezone).
	Locale *string `json:"locale"`
	// Does not enforce fixed viewport, allows resizing window in the headed mode.
	NoViewport *bool `json:"noViewport"`
	// Whether to emulate network being offline. Defaults to `false`. Learn more about
	// [network emulation](../emulation.md#offline).
	Offline *bool `json:"offline"`
	// A list of permissions to grant to all pages in this context. See BrowserContext.GrantPermissions()
	// for more details. Defaults to none.
	Permissions []string `json:"permissions"`
	// Network proxy settings to use with this context. Defaults to none.
	// For Chromium on Windows the browser needs to be launched with the global proxy for
	// this option to work. If all contexts override the proxy, global proxy will be never
	// used and can be any string, for example `launch({ proxy: { server: 'http://per-context'
	// } })`.
	Proxy *Proxy `json:"proxy"`
	// Optional setting to control resource content management. If `omit` is specified,
	// content is not persisted. If `attach` is specified, resources are persisted as separate
	// files and all of these files are archived along with the HAR file. Defaults to `embed`,
	// which stores content inline the HAR file as per HAR specification.
	RecordHarContent *HarContentPolicy `json:"recordHarContent"`
	// When set to `minimal`, only record information necessary for routing from HAR. This
	// omits sizes, timing, page, cookies, security and other types of HAR information
	// that are not used when replaying from HAR. Defaults to `full`.
	RecordHarMode *HarMode `json:"recordHarMode"`
	// Optional setting to control whether to omit request content from the HAR. Defaults
	// to `false`.
	RecordHarOmitContent *bool `json:"recordHarOmitContent"`
	// Enables [HAR](http://www.softwareishard.com/blog/har-12-spec) recording for all
	// pages into the specified HAR file on the filesystem. If not specified, the HAR is
	// not recorded. Make sure to call BrowserContext.Close() for the HAR to be saved.
	RecordHarPath      *string     `json:"recordHarPath"`
	RecordHarUrlFilter interface{} `json:"recordHarUrlFilter"`
	// Enables video recording for all pages into `recordVideo.dir` directory. If not specified
	// videos are not recorded. Make sure to await BrowserContext.Close() for videos to
	// be saved.
	RecordVideo *RecordVideo `json:"recordVideo"`
	// Emulates `'prefers-reduced-motion'` media feature, supported values are `'reduce'`,
	// `'no-preference'`. See Page.EmulateMedia() for more details. Passing `'no-override'`
	// resets emulation to system defaults. Defaults to `'no-preference'`.
	ReducedMotion *ReducedMotion `json:"reducedMotion"`
	// Emulates consistent window screen size available inside web page via `window.screen`.
	// Is only used when the `viewport` is set.
	Screen *ScreenSize `json:"screen"`
	// Whether to allow sites to register Service workers. Defaults to `'allow'`.
	// `'allow'`: [Service Workers](https://developer.mozilla.org/en-US/docs/Web/API/Service_Worker_API)
	// can be registered.
	// `'block'`: Playwright will block all registration of Service Workers.
	ServiceWorkers *ServiceWorkerPolicy `json:"serviceWorkers"`
	// Learn more about [storage state and auth](../auth.md).
	// Populates context with given storage state. This option can be used to initialize
	// context with logged-in information obtained via BrowserContext.StorageState().
	StorageState *OptionalStorageState `json:"storageState"`
	// Populates context with given storage state. This option can be used to initialize
	// context with logged-in information obtained via BrowserContext.StorageState(). Path
	// to the file with saved storage state.
	StorageStatePath *string `json:"storageStatePath"`
	// If set to true, enables strict selectors mode for this context. In the strict selectors
	// mode all operations on selectors that imply single target DOM element will throw
	// when more than one element matches the selector. This option does not affect any
	// Locator APIs (Locators are always strict). Defaults to `false`. See Locator to learn
	// more about the strict mode.
	StrictSelectors *bool `json:"strictSelectors"`
	// Changes the timezone of the context. See [ICU's metaZones.txt](https://cs.chromium.org/chromium/src/third_party/icu/source/data/misc/metaZones.txt?rcl=faee8bc70570192d82d2978a71e2a615788597d1)
	// for a list of supported timezone IDs. Defaults to the system timezone.
	TimezoneId *string `json:"timezoneId"`
	// Specific user agent to use in this context.
	UserAgent *string `json:"userAgent"`
	// Sets a consistent viewport for each page. Defaults to an 1280x720 viewport. `no_viewport`
	// disables the fixed viewport. Learn more about [viewport emulation](../emulation.md#viewport).
	Viewport *ViewportSize `json:"viewport"`
}
type RecordVideo struct {
	// Path to the directory to put videos into.
	Dir *string `json:"dir"`
	// Optional dimensions of the recorded videos. If not specified the size will be equal
	// to `viewport` scaled down to fit into 800x800. If `viewport` is not configured explicitly
	// the video size defaults to 800x450. Actual picture of each page will be scaled down
	// if necessary to fit the specified size.
	Size *RecordVideoSize `json:"size"`
}
type ScreenSize struct {
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
	// constructor for building the corresponding URL. Unset by default. Examples:
	// baseURL: `http://localhost:3000` and navigating to `/bar.html` results in `http://localhost:3000/bar.html`
	// baseURL: `http://localhost:3000/foo/` and navigating to `./bar.html` results in
	// `http://localhost:3000/foo/bar.html`
	// baseURL: `http://localhost:3000/foo` (without trailing slash) and navigating to
	// `./bar.html` results in `http://localhost:3000/bar.html`
	BaseURL *string `json:"baseURL"`
	// Toggles bypassing page's Content-Security-Policy. Defaults to `false`.
	BypassCSP *bool `json:"bypassCSP"`
	// Emulates `'prefers-colors-scheme'` media feature, supported values are `'light'`,
	// `'dark'`, `'no-preference'`. See Page.EmulateMedia() for more details. Passing `'no-override'`
	// resets emulation to system defaults. Defaults to `'light'`.
	ColorScheme *ColorScheme `json:"colorScheme"`
	// Specify device scale factor (can be thought of as dpr). Defaults to `1`. Learn more
	// about [emulating devices with device scale factor](../emulation.md#devices).
	DeviceScaleFactor *float64 `json:"deviceScaleFactor"`
	// An object containing additional HTTP headers to be sent with every request. Defaults
	// to none.
	ExtraHttpHeaders map[string]string `json:"extraHTTPHeaders"`
	// Emulates `'forced-colors'` media feature, supported values are `'active'`, `'none'`.
	// See Page.EmulateMedia() for more details. Passing `'no-override'` resets emulation
	// to system defaults. Defaults to `'none'`.
	ForcedColors *ForcedColors `json:"forcedColors"`
	Geolocation  *Geolocation  `json:"geolocation"`
	// Specifies if viewport supports touch events. Defaults to false. Learn more about
	// [mobile emulation](../emulation.md#devices).
	HasTouch *bool `json:"hasTouch"`
	// Credentials for [HTTP authentication](https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication).
	// If no origin is specified, the username and password are sent to any servers upon
	// unauthorized responses.
	HttpCredentials *HttpCredentials `json:"httpCredentials"`
	// Whether to ignore HTTPS errors when sending network requests. Defaults to `false`.
	IgnoreHttpsErrors *bool `json:"ignoreHTTPSErrors"`
	// Whether the `meta viewport` tag is taken into account and touch events are enabled.
	// isMobile is a part of device, so you don't actually need to set it manually. Defaults
	// to `false` and is not supported in Firefox. Learn more about [mobile emulation](../emulation.md#isMobile).
	IsMobile *bool `json:"isMobile"`
	// Whether or not to enable JavaScript in the context. Defaults to `true`. Learn more
	// about [disabling JavaScript](../emulation.md#javascript-enabled).
	JavaScriptEnabled *bool `json:"javaScriptEnabled"`
	// Specify user locale, for example `en-GB`, `de-DE`, etc. Locale will affect `navigator.language`
	// value, `Accept-Language` request header value as well as number and date formatting
	// rules. Defaults to the system default locale. Learn more about emulation in our
	// [emulation guide](../emulation.md#locale--timezone).
	Locale *string `json:"locale"`
	// Does not enforce fixed viewport, allows resizing window in the headed mode.
	NoViewport *bool `json:"noViewport"`
	// Whether to emulate network being offline. Defaults to `false`. Learn more about
	// [network emulation](../emulation.md#offline).
	Offline *bool `json:"offline"`
	// A list of permissions to grant to all pages in this context. See BrowserContext.GrantPermissions()
	// for more details. Defaults to none.
	Permissions []string `json:"permissions"`
	// Network proxy settings to use with this context. Defaults to none.
	// For Chromium on Windows the browser needs to be launched with the global proxy for
	// this option to work. If all contexts override the proxy, global proxy will be never
	// used and can be any string, for example `launch({ proxy: { server: 'http://per-context'
	// } })`.
	Proxy *Proxy `json:"proxy"`
	// Optional setting to control resource content management. If `omit` is specified,
	// content is not persisted. If `attach` is specified, resources are persisted as separate
	// files and all of these files are archived along with the HAR file. Defaults to `embed`,
	// which stores content inline the HAR file as per HAR specification.
	RecordHarContent *HarContentPolicy `json:"recordHarContent"`
	// When set to `minimal`, only record information necessary for routing from HAR. This
	// omits sizes, timing, page, cookies, security and other types of HAR information
	// that are not used when replaying from HAR. Defaults to `full`.
	RecordHarMode *HarMode `json:"recordHarMode"`
	// Optional setting to control whether to omit request content from the HAR. Defaults
	// to `false`.
	RecordHarOmitContent *bool `json:"recordHarOmitContent"`
	// Enables [HAR](http://www.softwareishard.com/blog/har-12-spec) recording for all
	// pages into the specified HAR file on the filesystem. If not specified, the HAR is
	// not recorded. Make sure to call BrowserContext.Close() for the HAR to be saved.
	RecordHarPath      *string     `json:"recordHarPath"`
	RecordHarUrlFilter interface{} `json:"recordHarUrlFilter"`
	// Enables video recording for all pages into `recordVideo.dir` directory. If not specified
	// videos are not recorded. Make sure to await BrowserContext.Close() for videos to
	// be saved.
	RecordVideo *RecordVideo `json:"recordVideo"`
	// Emulates `'prefers-reduced-motion'` media feature, supported values are `'reduce'`,
	// `'no-preference'`. See Page.EmulateMedia() for more details. Passing `'no-override'`
	// resets emulation to system defaults. Defaults to `'no-preference'`.
	ReducedMotion *ReducedMotion `json:"reducedMotion"`
	// Emulates consistent window screen size available inside web page via `window.screen`.
	// Is only used when the `viewport` is set.
	Screen *ScreenSize `json:"screen"`
	// Whether to allow sites to register Service workers. Defaults to `'allow'`.
	// `'allow'`: [Service Workers](https://developer.mozilla.org/en-US/docs/Web/API/Service_Worker_API)
	// can be registered.
	// `'block'`: Playwright will block all registration of Service Workers.
	ServiceWorkers *ServiceWorkerPolicy `json:"serviceWorkers"`
	// Learn more about [storage state and auth](../auth.md).
	// Populates context with given storage state. This option can be used to initialize
	// context with logged-in information obtained via BrowserContext.StorageState().
	StorageState *OptionalStorageState `json:"storageState"`
	// Populates context with given storage state. This option can be used to initialize
	// context with logged-in information obtained via BrowserContext.StorageState(). Path
	// to the file with saved storage state.
	StorageStatePath *string `json:"storageStatePath"`
	// If set to true, enables strict selectors mode for this context. In the strict selectors
	// mode all operations on selectors that imply single target DOM element will throw
	// when more than one element matches the selector. This option does not affect any
	// Locator APIs (Locators are always strict). Defaults to `false`. See Locator to learn
	// more about the strict mode.
	StrictSelectors *bool `json:"strictSelectors"`
	// Changes the timezone of the context. See [ICU's metaZones.txt](https://cs.chromium.org/chromium/src/third_party/icu/source/data/misc/metaZones.txt?rcl=faee8bc70570192d82d2978a71e2a615788597d1)
	// for a list of supported timezone IDs. Defaults to the system timezone.
	TimezoneId *string `json:"timezoneId"`
	// Specific user agent to use in this context.
	UserAgent *string `json:"userAgent"`
	// Sets a consistent viewport for each page. Defaults to an 1280x720 viewport. `no_viewport`
	// disables the fixed viewport. Learn more about [viewport emulation](../emulation.md#viewport).
	Viewport *ViewportSize `json:"viewport"`
}
type BrowserStartTracingOptions struct {
	// specify custom categories to use instead of default.
	Categories []string `json:"categories"`
	// Optional, if specified, tracing includes screenshots of the given page.
	Page Page `json:"page"`
	// A path to write the trace file to.
	Path *string `json:"path"`
	// captures screenshots in the trace.
	Screenshots *bool `json:"screenshots"`
}
type BrowserContextAddCookiesOptions struct {
	// Adds cookies to the browser context.
	// For the cookie to apply to all subdomains as well, prefix domain with a dot, like
	// this: ".example.com".
	Cookies []OptionalCookie `json:"cookies"`
}
type BrowserContextAddInitScriptOptions struct {
	// Optional Script source to be evaluated in all pages in the browser context.
	Script *string `json:"script"`
	// Optional Script path to be evaluated in all pages in the browser context.
	Path *string `json:"path"`
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
type BrowserContextRouteFromHAROptions struct {
	// If set to 'abort' any request not found in the HAR file will be aborted.
	// If set to 'fallback' falls through to the next route handler in the handler chain.
	// Defaults to abort.
	NotFound *HarNotFound `json:"notFound"`
	// If specified, updates the given HAR with the actual network information instead
	// of serving from file. The file is written to disk when BrowserContext.Close() is
	// called.
	Update *bool `json:"update"`
	// Optional setting to control resource content management. If `attach` is specified,
	// resources are persisted as separate files or entries in the ZIP archive. If `embed`
	// is specified, content is stored inline the HAR file.
	UpdateContent *RouteFromHarUpdateContentPolicy `json:"updateContent"`
	// When set to `minimal`, only record information necessary for routing from HAR. This
	// omits sizes, timing, page, cookies, security and other types of HAR information
	// that are not used when replaying from HAR. Defaults to `minimal`.
	UpdateMode *HarMode `json:"updateMode"`
	// A glob pattern, regular expression or predicate to match the request URL. Only requests
	// with URL matching the pattern will be served from the HAR file. If not specified,
	// all requests are served from the HAR file.
	URL interface{} `json:"url"`
}
type BrowserContextStorageStateOptions struct {
	// The file path to save the storage state to. If `path` is a relative path, then it
	// is resolved relative to current working directory. If no path is provided, storage
	// state is still returned, but won't be saved to the disk.
	Path *string `json:"path"`
}
type BrowserContextUnrouteOptions struct {
	// Optional handler function used to register a routing with BrowserContext.Route().
	Handler func(Route) `json:"handler"`
}
type BrowserContextExpectConsoleMessageOptions struct {
	// Receives the ConsoleMessage object and resolves to truthy value when the waiting
	// should resolve.
	Predicate func(ConsoleMessage) bool `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type BrowserContextExpectEventOptions struct {
	// Receives the event data and resolves to truthy value when the waiting should resolve.
	Predicate interface{} `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type BrowserContextExpectPageOptions struct {
	// Receives the Page object and resolves to truthy value when the waiting should resolve.
	Predicate func(Page) bool `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type BrowserContextWaitForEventOptions struct {
	// Receives the event data and resolves to truthy value when the waiting should resolve.
	Predicate interface{} `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type BrowserTypeConnectOptions struct {
	// Additional HTTP headers to be sent with web socket connect request. Optional.
	Headers map[string]string `json:"headers"`
	// Slows down Playwright operations by the specified amount of milliseconds. Useful
	// so that you can see what is going on. Defaults to 0.
	SlowMo *float64 `json:"slowMo"`
	// Maximum time in milliseconds to wait for the connection to be established. Defaults
	// to `0` (no timeout).
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
	// about using [Google Chrome and Microsoft Edge](../browsers.md#google-chrome--microsoft-edge).
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
	// If `true`, Playwright does not pass its own configurations args and only uses the
	// ones from `args`. Dangerous option; use with care. Defaults to `false`.
	IgnoreAllDefaultArgs *bool `json:"ignoreAllDefaultArgs"`
	// If `true`, Playwright does not pass its own configurations args and only uses the
	// ones from `args`. Dangerous option; use with care.
	IgnoreDefaultArgs []string `json:"ignoreDefaultArgs"`
	// Network proxy settings.
	Proxy *Proxy `json:"proxy"`
	// Slows down Playwright operations by the specified amount of milliseconds. Useful
	// so that you can see what is going on.
	SlowMo *float64 `json:"slowMo"`
	// Maximum time in milliseconds to wait for the browser instance to start. Defaults
	// to `30000` (30 seconds). Pass `0` to disable timeout.
	Timeout *float64 `json:"timeout"`
	// If specified, traces are saved into this directory.
	TracesDir *string `json:"tracesDir"`
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
	// constructor for building the corresponding URL. Unset by default. Examples:
	// baseURL: `http://localhost:3000` and navigating to `/bar.html` results in `http://localhost:3000/bar.html`
	// baseURL: `http://localhost:3000/foo/` and navigating to `./bar.html` results in
	// `http://localhost:3000/foo/bar.html`
	// baseURL: `http://localhost:3000/foo` (without trailing slash) and navigating to
	// `./bar.html` results in `http://localhost:3000/bar.html`
	BaseURL *string `json:"baseURL"`
	// Toggles bypassing page's Content-Security-Policy. Defaults to `false`.
	BypassCSP *bool `json:"bypassCSP"`
	// Browser distribution channel.  Supported values are "chrome", "chrome-beta", "chrome-dev",
	// "chrome-canary", "msedge", "msedge-beta", "msedge-dev", "msedge-canary". Read more
	// about using [Google Chrome and Microsoft Edge](../browsers.md#google-chrome--microsoft-edge).
	Channel *string `json:"channel"`
	// Enable Chromium sandboxing. Defaults to `false`.
	ChromiumSandbox *bool `json:"chromiumSandbox"`
	// Emulates `'prefers-colors-scheme'` media feature, supported values are `'light'`,
	// `'dark'`, `'no-preference'`. See Page.EmulateMedia() for more details. Passing `'no-override'`
	// resets emulation to system defaults. Defaults to `'light'`.
	ColorScheme *ColorScheme `json:"colorScheme"`
	// Specify device scale factor (can be thought of as dpr). Defaults to `1`. Learn more
	// about [emulating devices with device scale factor](../emulation.md#devices).
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
	// An object containing additional HTTP headers to be sent with every request. Defaults
	// to none.
	ExtraHttpHeaders map[string]string `json:"extraHTTPHeaders"`
	// Emulates `'forced-colors'` media feature, supported values are `'active'`, `'none'`.
	// See Page.EmulateMedia() for more details. Passing `'no-override'` resets emulation
	// to system defaults. Defaults to `'none'`.
	ForcedColors *ForcedColors `json:"forcedColors"`
	Geolocation  *Geolocation  `json:"geolocation"`
	// Close the browser process on SIGHUP. Defaults to `true`.
	HandleSIGHUP *bool `json:"handleSIGHUP"`
	// Close the browser process on Ctrl-C. Defaults to `true`.
	HandleSIGINT *bool `json:"handleSIGINT"`
	// Close the browser process on SIGTERM. Defaults to `true`.
	HandleSIGTERM *bool `json:"handleSIGTERM"`
	// Specifies if viewport supports touch events. Defaults to false. Learn more about
	// [mobile emulation](../emulation.md#devices).
	HasTouch *bool `json:"hasTouch"`
	// Whether to run browser in headless mode. More details for [Chromium](https://developers.google.com/web/updates/2017/04/headless-chrome)
	// and [Firefox](https://developer.mozilla.org/en-US/docs/Mozilla/Firefox/Headless_mode).
	// Defaults to `true` unless the `devtools` option is `true`.
	Headless *bool `json:"headless"`
	// Credentials for [HTTP authentication](https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication).
	// If no origin is specified, the username and password are sent to any servers upon
	// unauthorized responses.
	HttpCredentials *HttpCredentials `json:"httpCredentials"`
	// If `true`, Playwright does not pass its own configurations args and only uses the
	// ones from `args`. Dangerous option; use with care. Defaults to `false`.
	IgnoreAllDefaultArgs *bool `json:"ignoreAllDefaultArgs"`
	// If `true`, Playwright does not pass its own configurations args and only uses the
	// ones from `args`. Dangerous option; use with care.
	IgnoreDefaultArgs []string `json:"ignoreDefaultArgs"`
	// Whether to ignore HTTPS errors when sending network requests. Defaults to `false`.
	IgnoreHttpsErrors *bool `json:"ignoreHTTPSErrors"`
	// Whether the `meta viewport` tag is taken into account and touch events are enabled.
	// isMobile is a part of device, so you don't actually need to set it manually. Defaults
	// to `false` and is not supported in Firefox. Learn more about [mobile emulation](../emulation.md#isMobile).
	IsMobile *bool `json:"isMobile"`
	// Whether or not to enable JavaScript in the context. Defaults to `true`. Learn more
	// about [disabling JavaScript](../emulation.md#javascript-enabled).
	JavaScriptEnabled *bool `json:"javaScriptEnabled"`
	// Specify user locale, for example `en-GB`, `de-DE`, etc. Locale will affect `navigator.language`
	// value, `Accept-Language` request header value as well as number and date formatting
	// rules. Defaults to the system default locale. Learn more about emulation in our
	// [emulation guide](../emulation.md#locale--timezone).
	Locale *string `json:"locale"`
	// Does not enforce fixed viewport, allows resizing window in the headed mode.
	NoViewport *bool `json:"noViewport"`
	// Whether to emulate network being offline. Defaults to `false`. Learn more about
	// [network emulation](../emulation.md#offline).
	Offline *bool `json:"offline"`
	// A list of permissions to grant to all pages in this context. See BrowserContext.GrantPermissions()
	// for more details. Defaults to none.
	Permissions []string `json:"permissions"`
	// Network proxy settings.
	Proxy *Proxy `json:"proxy"`
	// Optional setting to control resource content management. If `omit` is specified,
	// content is not persisted. If `attach` is specified, resources are persisted as separate
	// files and all of these files are archived along with the HAR file. Defaults to `embed`,
	// which stores content inline the HAR file as per HAR specification.
	RecordHarContent *HarContentPolicy `json:"recordHarContent"`
	// When set to `minimal`, only record information necessary for routing from HAR. This
	// omits sizes, timing, page, cookies, security and other types of HAR information
	// that are not used when replaying from HAR. Defaults to `full`.
	RecordHarMode *HarMode `json:"recordHarMode"`
	// Optional setting to control whether to omit request content from the HAR. Defaults
	// to `false`.
	RecordHarOmitContent *bool `json:"recordHarOmitContent"`
	// Enables [HAR](http://www.softwareishard.com/blog/har-12-spec) recording for all
	// pages into the specified HAR file on the filesystem. If not specified, the HAR is
	// not recorded. Make sure to call BrowserContext.Close() for the HAR to be saved.
	RecordHarPath      *string     `json:"recordHarPath"`
	RecordHarUrlFilter interface{} `json:"recordHarUrlFilter"`
	// Enables video recording for all pages into `recordVideo.dir` directory. If not specified
	// videos are not recorded. Make sure to await BrowserContext.Close() for videos to
	// be saved.
	RecordVideo *RecordVideo `json:"recordVideo"`
	// Emulates `'prefers-reduced-motion'` media feature, supported values are `'reduce'`,
	// `'no-preference'`. See Page.EmulateMedia() for more details. Passing `'no-override'`
	// resets emulation to system defaults. Defaults to `'no-preference'`.
	ReducedMotion *ReducedMotion `json:"reducedMotion"`
	// Emulates consistent window screen size available inside web page via `window.screen`.
	// Is only used when the `viewport` is set.
	Screen *ScreenSize `json:"screen"`
	// Whether to allow sites to register Service workers. Defaults to `'allow'`.
	// `'allow'`: [Service Workers](https://developer.mozilla.org/en-US/docs/Web/API/Service_Worker_API)
	// can be registered.
	// `'block'`: Playwright will block all registration of Service Workers.
	ServiceWorkers *ServiceWorkerPolicy `json:"serviceWorkers"`
	// Slows down Playwright operations by the specified amount of milliseconds. Useful
	// so that you can see what is going on.
	SlowMo *float64 `json:"slowMo"`
	// If set to true, enables strict selectors mode for this context. In the strict selectors
	// mode all operations on selectors that imply single target DOM element will throw
	// when more than one element matches the selector. This option does not affect any
	// Locator APIs (Locators are always strict). Defaults to `false`. See Locator to learn
	// more about the strict mode.
	StrictSelectors *bool `json:"strictSelectors"`
	// Maximum time in milliseconds to wait for the browser instance to start. Defaults
	// to `30000` (30 seconds). Pass `0` to disable timeout.
	Timeout *float64 `json:"timeout"`
	// Changes the timezone of the context. See [ICU's metaZones.txt](https://cs.chromium.org/chromium/src/third_party/icu/source/data/misc/metaZones.txt?rcl=faee8bc70570192d82d2978a71e2a615788597d1)
	// for a list of supported timezone IDs. Defaults to the system timezone.
	TimezoneId *string `json:"timezoneId"`
	// If specified, traces are saved into this directory.
	TracesDir *string `json:"tracesDir"`
	// Specific user agent to use in this context.
	UserAgent *string `json:"userAgent"`
	// Sets a consistent viewport for each page. Defaults to an 1280x720 viewport. `no_viewport`
	// disables the fixed viewport. Learn more about [viewport emulation](../emulation.md#viewport).
	Viewport *ViewportSize `json:"viewport"`
}
type DialogAcceptOptions struct {
	// A text to enter in prompt. Does not cause any effects if the dialog's `type` is
	// not prompt. Optional.
	PromptText *string `json:"promptText"`
}
type ElementHandleCheckOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *Position `json:"position"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type Position struct {
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
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type ElementHandleDblclickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
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
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandleHoverOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type ElementHandleInputValueOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
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
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
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
	// When set to `"hide"`, screenshot will hide text caret. When set to `"initial"`,
	// text caret behavior will not be changed.  Defaults to `"hide"`.
	Caret *ScreenshotCaret `json:"caret"`
	// Specify the color of the overlay box for masked elements, in [CSS color format](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value).
	// Default color is pink `#FF00FF`.
	MaskColor *string `json:"maskColor"`
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
	// will produce a single pixel per each device pixel, so screenshots of high-dpi devices
	// will be twice as large or even larger.
	// Defaults to `"device"`.
	Scale *ScreenshotScale `json:"scale"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// Specify screenshot type, defaults to `png`.
	Type *ScreenshotType `json:"type"`
}
type ElementHandleScrollIntoViewIfNeededOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandleSelectOptionOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandleSelectTextOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandleSetCheckedOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *Position `json:"position"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
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
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandleTapOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
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
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type ElementHandleUncheckOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *Position `json:"position"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type ElementHandleWaitForElementStateOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
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
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FileChooserSetFilesOptions struct {
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
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
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *Position `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FrameClickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// defaults to 1. See [UIEvent.detail].
	ClickCount *int `json:"clickCount"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FrameDblclickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FrameDispatchEventOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameDragAndDropOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Clicks on the source element at this point relative to the top-left corner of the
	// element's padding box. If not specified, some visible point of the element is used.
	SourcePosition *Position `json:"sourcePosition"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Drops on the target element at this point relative to the top-left corner of the
	// element's padding box. If not specified, some visible point of the element is used.
	TargetPosition *Position `json:"targetPosition"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FrameEvalOnSelectorOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
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
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameFocusOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameGetAttributeOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameGetByAltTextOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type FrameGetByLabelOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type FrameGetByPlaceholderOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type FrameGetByRoleOptions struct {
	// An attribute that is usually set by `aria-checked` or native `<input type=checkbox>`
	// controls.
	// Learn more about [`aria-checked`](https://www.w3.org/TR/wai-aria-1.2/#aria-checked).
	Checked *bool `json:"checked"`
	// An attribute that is usually set by `aria-disabled` or `disabled`.
	// Unlike most other attributes, `disabled` is inherited through the DOM hierarchy.
	// Learn more about [`aria-disabled`](https://www.w3.org/TR/wai-aria-1.2/#aria-disabled).
	Disabled *bool `json:"disabled"`
	// Whether `name` is matched exactly: case-sensitive and whole-string. Defaults to
	// false. Ignored when `name` is a regular expression. Note that exact match still
	// trims whitespace.
	Exact *bool `json:"exact"`
	// An attribute that is usually set by `aria-expanded`.
	// Learn more about [`aria-expanded`](https://www.w3.org/TR/wai-aria-1.2/#aria-expanded).
	Expanded *bool `json:"expanded"`
	// Option that controls whether hidden elements are matched. By default, only non-hidden
	// elements, as [defined by ARIA](https://www.w3.org/TR/wai-aria-1.2/#tree_exclusion),
	// are matched by role selector.
	// Learn more about [`aria-hidden`](https://www.w3.org/TR/wai-aria-1.2/#aria-hidden).
	IncludeHidden *bool `json:"includeHidden"`
	// A number attribute that is usually present for roles `heading`, `listitem`, `row`,
	// `treeitem`, with default values for `<h1>-<h6>` elements.
	// Learn more about [`aria-level`](https://www.w3.org/TR/wai-aria-1.2/#aria-level).
	Level *int `json:"level"`
	// Option to match the [accessible name](https://w3c.github.io/accname/#dfn-accessible-name).
	// By default, matching is case-insensitive and searches for a substring, use `exact`
	// to control this behavior.
	// Learn more about [accessible name](https://w3c.github.io/accname/#dfn-accessible-name).
	Name interface{} `json:"name"`
	// An attribute that is usually set by `aria-pressed`.
	// Learn more about [`aria-pressed`](https://www.w3.org/TR/wai-aria-1.2/#aria-pressed).
	Pressed *bool `json:"pressed"`
	// An attribute that is usually set by `aria-selected`.
	// Learn more about [`aria-selected`](https://www.w3.org/TR/wai-aria-1.2/#aria-selected).
	Selected *bool `json:"selected"`
}
type FrameGetByTextOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type FrameGetByTitleOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
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
	// `'networkidle'` - **DISCOURAGED** consider operation to be finished when there are
	// no network connections for at least `500` ms. Don't use this method for testing,
	// rely on web assertions to assess readiness instead.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type FrameHoverOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FrameInnerHTMLOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameInnerTextOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameInputValueOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameIsCheckedOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameIsDisabledOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameIsEditableOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameIsEnabledOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameIsHiddenOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict  *bool    `json:"strict"`
	Timeout *float64 `json:"timeout"`
}
type FrameIsVisibleOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict  *bool    `json:"strict"`
	Timeout *float64 `json:"timeout"`
}
type FrameLocatorOptions struct {
	// Matches elements containing an element that matches an inner locator. Inner locator
	// is queried against the outer one. For example, `article` that has `text=Playwright`
	// matches `<article><div>Playwright</div></article>`.
	// Note that outer and inner locators must belong to the same frame. Inner locator
	// must not contain FrameLocators.
	Has Locator `json:"has"`
	// Matches elements that do not contain an element that matches an inner locator. Inner
	// locator is queried against the outer one. For example, `article` that does not have
	// `div` matches `<article><span>Playwright</span></article>`.
	// Note that outer and inner locators must belong to the same frame. Inner locator
	// must not contain FrameLocators.
	HasNot Locator `json:"hasNot"`
	// Matches elements that do not contain specified text somewhere inside, possibly in
	// a child or a descendant element. When passed a [string], matching is case-insensitive
	// and searches for a substring.
	HasNotText interface{} `json:"hasNotText"`
	// Matches elements containing specified text somewhere inside, possibly in a child
	// or a descendant element. When passed a [string], matching is case-insensitive and
	// searches for a substring. For example, `"Playwright"` matches `<article><div>Playwright</div></article>`.
	HasText interface{} `json:"hasText"`
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
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameQuerySelectorOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
}
type FrameSelectOptionOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameSetCheckedOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *Position `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
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
	// `'networkidle'` - **DISCOURAGED** consider operation to be finished when there are
	// no network connections for at least `500` ms. Don't use this method for testing,
	// rely on web assertions to assess readiness instead.
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
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameTapOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FrameTextContentOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
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
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameUncheckOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *Position `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type FrameWaitForFunctionOptions struct {
	// If `polling` is `'raf'`, then `expression` is constantly executed in `requestAnimationFrame`
	// callback. If `polling` is a number, then it is treated as an interval in milliseconds
	// at which the function would be executed. Defaults to `raf`.
	Polling interface{} `json:"polling"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type FrameWaitForLoadStateOptions struct {
	// Optional load state to wait for, defaults to `load`. If the state has been already
	// reached while loading current document, the method resolves immediately. Can be
	// one of:
	// `'load'` - wait for the `load` event to be fired.
	// `'domcontentloaded'` - wait for the `DOMContentLoaded` event to be fired.
	// `'networkidle'` - **DISCOURAGED** wait until there are no network connections for
	// at least `500` ms. Don't use this method for testing, rely on web assertions to
	// assess readiness instead.
	State *LoadState `json:"state"`
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
	// for the navigation. Note that if the parameter is a string without wildcard characters,
	// the method will wait for navigation to URL that is exactly equal to the string.
	URL interface{} `json:"url"`
	// When to consider operation succeeded, defaults to `load`. Events can be either:
	// `'domcontentloaded'` - consider operation to be finished when the `DOMContentLoaded`
	// event is fired.
	// `'load'` - consider operation to be finished when the `load` event is fired.
	// `'networkidle'` - **DISCOURAGED** consider operation to be finished when there are
	// no network connections for at least `500` ms. Don't use this method for testing,
	// rely on web assertions to assess readiness instead.
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
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
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
	// `'networkidle'` - **DISCOURAGED** consider operation to be finished when there are
	// no network connections for at least `500` ms. Don't use this method for testing,
	// rely on web assertions to assess readiness instead.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type FrameLocatorGetByAltTextOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type FrameLocatorGetByLabelOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type FrameLocatorGetByPlaceholderOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type FrameLocatorGetByRoleOptions struct {
	// An attribute that is usually set by `aria-checked` or native `<input type=checkbox>`
	// controls.
	// Learn more about [`aria-checked`](https://www.w3.org/TR/wai-aria-1.2/#aria-checked).
	Checked *bool `json:"checked"`
	// An attribute that is usually set by `aria-disabled` or `disabled`.
	// Unlike most other attributes, `disabled` is inherited through the DOM hierarchy.
	// Learn more about [`aria-disabled`](https://www.w3.org/TR/wai-aria-1.2/#aria-disabled).
	Disabled *bool `json:"disabled"`
	// Whether `name` is matched exactly: case-sensitive and whole-string. Defaults to
	// false. Ignored when `name` is a regular expression. Note that exact match still
	// trims whitespace.
	Exact *bool `json:"exact"`
	// An attribute that is usually set by `aria-expanded`.
	// Learn more about [`aria-expanded`](https://www.w3.org/TR/wai-aria-1.2/#aria-expanded).
	Expanded *bool `json:"expanded"`
	// Option that controls whether hidden elements are matched. By default, only non-hidden
	// elements, as [defined by ARIA](https://www.w3.org/TR/wai-aria-1.2/#tree_exclusion),
	// are matched by role selector.
	// Learn more about [`aria-hidden`](https://www.w3.org/TR/wai-aria-1.2/#aria-hidden).
	IncludeHidden *bool `json:"includeHidden"`
	// A number attribute that is usually present for roles `heading`, `listitem`, `row`,
	// `treeitem`, with default values for `<h1>-<h6>` elements.
	// Learn more about [`aria-level`](https://www.w3.org/TR/wai-aria-1.2/#aria-level).
	Level *int `json:"level"`
	// Option to match the [accessible name](https://w3c.github.io/accname/#dfn-accessible-name).
	// By default, matching is case-insensitive and searches for a substring, use `exact`
	// to control this behavior.
	// Learn more about [accessible name](https://w3c.github.io/accname/#dfn-accessible-name).
	Name interface{} `json:"name"`
	// An attribute that is usually set by `aria-pressed`.
	// Learn more about [`aria-pressed`](https://www.w3.org/TR/wai-aria-1.2/#aria-pressed).
	Pressed *bool `json:"pressed"`
	// An attribute that is usually set by `aria-selected`.
	// Learn more about [`aria-selected`](https://www.w3.org/TR/wai-aria-1.2/#aria-selected).
	Selected *bool `json:"selected"`
}
type FrameLocatorGetByTextOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type FrameLocatorGetByTitleOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
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
type LocatorBlurOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorBoundingBoxOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorCheckOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *Position `json:"position"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type LocatorClearOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorClickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// defaults to 1. See [UIEvent.detail].
	ClickCount *int `json:"clickCount"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type LocatorDblclickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type LocatorDispatchEventOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorDragToOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Clicks on the source element at this point relative to the top-left corner of the
	// element's padding box. If not specified, some visible point of the element is used.
	SourcePosition *Position `json:"sourcePosition"`
	// Drops on the target element at this point relative to the top-left corner of the
	// element's padding box. If not specified, some visible point of the element is used.
	TargetPosition *Position `json:"targetPosition"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type LocatorElementHandleOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorEvaluateOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorEvaluateAllOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type LocatorEvaluateHandleOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorFillOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorFilterOptions struct {
	// Matches elements containing an element that matches an inner locator. Inner locator
	// is queried against the outer one. For example, `article` that has `text=Playwright`
	// matches `<article><div>Playwright</div></article>`.
	// Note that outer and inner locators must belong to the same frame. Inner locator
	// must not contain FrameLocators.
	Has Locator `json:"has"`
	// Matches elements that do not contain an element that matches an inner locator. Inner
	// locator is queried against the outer one. For example, `article` that does not have
	// `div` matches `<article><span>Playwright</span></article>`.
	// Note that outer and inner locators must belong to the same frame. Inner locator
	// must not contain FrameLocators.
	HasNot Locator `json:"hasNot"`
	// Matches elements that do not contain specified text somewhere inside, possibly in
	// a child or a descendant element. When passed a [string], matching is case-insensitive
	// and searches for a substring.
	HasNotText interface{} `json:"hasNotText"`
	// Matches elements containing specified text somewhere inside, possibly in a child
	// or a descendant element. When passed a [string], matching is case-insensitive and
	// searches for a substring. For example, `"Playwright"` matches `<article><div>Playwright</div></article>`.
	HasText interface{} `json:"hasText"`
}
type LocatorFocusOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorGetAttributeOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorGetByAltTextOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type LocatorGetByLabelOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type LocatorGetByPlaceholderOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type LocatorGetByRoleOptions struct {
	// An attribute that is usually set by `aria-checked` or native `<input type=checkbox>`
	// controls.
	// Learn more about [`aria-checked`](https://www.w3.org/TR/wai-aria-1.2/#aria-checked).
	Checked *bool `json:"checked"`
	// An attribute that is usually set by `aria-disabled` or `disabled`.
	// Unlike most other attributes, `disabled` is inherited through the DOM hierarchy.
	// Learn more about [`aria-disabled`](https://www.w3.org/TR/wai-aria-1.2/#aria-disabled).
	Disabled *bool `json:"disabled"`
	// Whether `name` is matched exactly: case-sensitive and whole-string. Defaults to
	// false. Ignored when `name` is a regular expression. Note that exact match still
	// trims whitespace.
	Exact *bool `json:"exact"`
	// An attribute that is usually set by `aria-expanded`.
	// Learn more about [`aria-expanded`](https://www.w3.org/TR/wai-aria-1.2/#aria-expanded).
	Expanded *bool `json:"expanded"`
	// Option that controls whether hidden elements are matched. By default, only non-hidden
	// elements, as [defined by ARIA](https://www.w3.org/TR/wai-aria-1.2/#tree_exclusion),
	// are matched by role selector.
	// Learn more about [`aria-hidden`](https://www.w3.org/TR/wai-aria-1.2/#aria-hidden).
	IncludeHidden *bool `json:"includeHidden"`
	// A number attribute that is usually present for roles `heading`, `listitem`, `row`,
	// `treeitem`, with default values for `<h1>-<h6>` elements.
	// Learn more about [`aria-level`](https://www.w3.org/TR/wai-aria-1.2/#aria-level).
	Level *int `json:"level"`
	// Option to match the [accessible name](https://w3c.github.io/accname/#dfn-accessible-name).
	// By default, matching is case-insensitive and searches for a substring, use `exact`
	// to control this behavior.
	// Learn more about [accessible name](https://w3c.github.io/accname/#dfn-accessible-name).
	Name interface{} `json:"name"`
	// An attribute that is usually set by `aria-pressed`.
	// Learn more about [`aria-pressed`](https://www.w3.org/TR/wai-aria-1.2/#aria-pressed).
	Pressed *bool `json:"pressed"`
	// An attribute that is usually set by `aria-selected`.
	// Learn more about [`aria-selected`](https://www.w3.org/TR/wai-aria-1.2/#aria-selected).
	Selected *bool `json:"selected"`
}
type LocatorGetByTextOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type LocatorGetByTitleOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type LocatorHoverOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type LocatorInnerHTMLOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorInnerTextOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorInputValueOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorIsCheckedOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorIsDisabledOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorIsEditableOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorIsEnabledOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorIsHiddenOptions struct {
	Timeout *float64 `json:"timeout"`
}
type LocatorIsVisibleOptions struct {
	Timeout *float64 `json:"timeout"`
}
type LocatorLocatorOptions struct {
	// Matches elements containing an element that matches an inner locator. Inner locator
	// is queried against the outer one. For example, `article` that has `text=Playwright`
	// matches `<article><div>Playwright</div></article>`.
	// Note that outer and inner locators must belong to the same frame. Inner locator
	// must not contain FrameLocators.
	Has Locator `json:"has"`
	// Matches elements that do not contain an element that matches an inner locator. Inner
	// locator is queried against the outer one. For example, `article` that does not have
	// `div` matches `<article><span>Playwright</span></article>`.
	// Note that outer and inner locators must belong to the same frame. Inner locator
	// must not contain FrameLocators.
	HasNot Locator `json:"hasNot"`
	// Matches elements that do not contain specified text somewhere inside, possibly in
	// a child or a descendant element. When passed a [string], matching is case-insensitive
	// and searches for a substring.
	HasNotText interface{} `json:"hasNotText"`
	// Matches elements containing specified text somewhere inside, possibly in a child
	// or a descendant element. When passed a [string], matching is case-insensitive and
	// searches for a substring. For example, `"Playwright"` matches `<article><div>Playwright</div></article>`.
	HasText interface{} `json:"hasText"`
}
type LocatorPressOptions struct {
	// Time to wait between `keydown` and `keyup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
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
	// When set to `"hide"`, screenshot will hide text caret. When set to `"initial"`,
	// text caret behavior will not be changed.  Defaults to `"hide"`.
	Caret *ScreenshotCaret `json:"caret"`
	// Specify the color of the overlay box for masked elements, in [CSS color format](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value).
	// Default color is pink `#FF00FF`.
	MaskColor *string `json:"maskColor"`
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
	// will produce a single pixel per each device pixel, so screenshots of high-dpi devices
	// will be twice as large or even larger.
	// Defaults to `"device"`.
	Scale *ScreenshotScale `json:"scale"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// Specify screenshot type, defaults to `png`.
	Type *ScreenshotType `json:"type"`
}
type LocatorScrollIntoViewIfNeededOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorSelectOptionOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorSelectTextOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorSetCheckedOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *Position `json:"position"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
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
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorTapOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type LocatorTextContentOptions struct {
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
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
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorUncheckOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *Position `json:"position"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
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
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToBeAttachedOptions struct {
	Attached *bool `json:"attached"`
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToBeCheckedOptions struct {
	Checked *bool `json:"checked"`
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToBeDisabledOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToBeEditableOptions struct {
	Editable *bool `json:"editable"`
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToBeEmptyOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToBeEnabledOptions struct {
	Enabled *bool `json:"enabled"`
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToBeFocusedOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToBeHiddenOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToBeInViewportOptions struct {
	// The minimal ratio of the element to intersect viewport. If equals to `0`, then element
	// should intersect viewport at any positive ratio. Defaults to `0`.
	Ratio *float64 `json:"ratio"`
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToBeVisibleOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
	Visible *bool    `json:"visible"`
}
type LocatorAssertionsToContainTextOptions struct {
	// Whether to perform case-insensitive match. `ignoreCase` option takes precedence
	// over the corresponding regular expression flag if specified.
	IgnoreCase *bool `json:"ignoreCase"`
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
	// Whether to use `element.innerText` instead of `element.textContent` when retrieving
	// DOM node text.
	UseInnerText *bool `json:"useInnerText"`
}
type LocatorAssertionsToHaveAttributeOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToHaveClassOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToHaveCountOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToHaveCSSOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToHaveIdOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToHaveJSPropertyOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToHaveTextOptions struct {
	// Whether to perform case-insensitive match. `ignoreCase` option takes precedence
	// over the corresponding regular expression flag if specified.
	IgnoreCase *bool `json:"ignoreCase"`
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
	// Whether to use `element.innerText` instead of `element.textContent` when retrieving
	// DOM node text.
	UseInnerText *bool `json:"useInnerText"`
}
type LocatorAssertionsToHaveValueOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type LocatorAssertionsToHaveValuesOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
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
	// Defaults to 1. Sends intermediate `mousemove` events.
	Steps *int `json:"steps"`
}
type MouseUpOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// defaults to 1. See [UIEvent.detail].
	ClickCount *int `json:"clickCount"`
}
type PageAddInitScriptOptions struct {
	// Path to the JavaScript file. If `path` is a relative path, then it is resolved relative
	// to the current working directory. Optional.
	Path *string `json:"path"`
	// Script to be evaluated in all pages in the browser context. Optional.
	Script *string `json:"script"`
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
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *Position `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type PageClickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// defaults to 1. See [UIEvent.detail].
	ClickCount *int `json:"clickCount"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
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
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type PageDispatchEventOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageDragAndDropOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// Clicks on the source element at this point relative to the top-left corner of the
	// element's padding box. If not specified, some visible point of the element is used.
	SourcePosition *Position `json:"sourcePosition"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Drops on the target element at this point relative to the top-left corner of the
	// element's padding box. If not specified, some visible point of the element is used.
	TargetPosition *Position `json:"targetPosition"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type PageEmulateMediaOptions struct {
	// Emulates `'prefers-colors-scheme'` media feature, supported values are `'light'`,
	// `'dark'`, `'no-preference'`. Passing `'no-override'` disables color scheme emulation.
	ColorScheme  *ColorScheme  `json:"colorScheme"`
	ForcedColors *ForcedColors `json:"forcedColors"`
	// Changes the CSS media type of the page. The only allowed values are `'screen'`,
	// `'print'` and `'no-override'`. Passing `'no-override'` disables CSS media emulation.
	Media *Media `json:"media"`
	// Emulates `'prefers-reduced-motion'` media feature, supported values are `'reduce'`,
	// `'no-preference'`. Passing `no-override` disables reduced motion emulation.
	ReducedMotion *ReducedMotion `json:"reducedMotion"`
}
type PageEvalOnSelectorOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
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
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageFocusOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageGetAttributeOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageGetByAltTextOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type PageGetByLabelOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type PageGetByPlaceholderOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type PageGetByRoleOptions struct {
	// An attribute that is usually set by `aria-checked` or native `<input type=checkbox>`
	// controls.
	// Learn more about [`aria-checked`](https://www.w3.org/TR/wai-aria-1.2/#aria-checked).
	Checked *bool `json:"checked"`
	// An attribute that is usually set by `aria-disabled` or `disabled`.
	// Unlike most other attributes, `disabled` is inherited through the DOM hierarchy.
	// Learn more about [`aria-disabled`](https://www.w3.org/TR/wai-aria-1.2/#aria-disabled).
	Disabled *bool `json:"disabled"`
	// Whether `name` is matched exactly: case-sensitive and whole-string. Defaults to
	// false. Ignored when `name` is a regular expression. Note that exact match still
	// trims whitespace.
	Exact *bool `json:"exact"`
	// An attribute that is usually set by `aria-expanded`.
	// Learn more about [`aria-expanded`](https://www.w3.org/TR/wai-aria-1.2/#aria-expanded).
	Expanded *bool `json:"expanded"`
	// Option that controls whether hidden elements are matched. By default, only non-hidden
	// elements, as [defined by ARIA](https://www.w3.org/TR/wai-aria-1.2/#tree_exclusion),
	// are matched by role selector.
	// Learn more about [`aria-hidden`](https://www.w3.org/TR/wai-aria-1.2/#aria-hidden).
	IncludeHidden *bool `json:"includeHidden"`
	// A number attribute that is usually present for roles `heading`, `listitem`, `row`,
	// `treeitem`, with default values for `<h1>-<h6>` elements.
	// Learn more about [`aria-level`](https://www.w3.org/TR/wai-aria-1.2/#aria-level).
	Level *int `json:"level"`
	// Option to match the [accessible name](https://w3c.github.io/accname/#dfn-accessible-name).
	// By default, matching is case-insensitive and searches for a substring, use `exact`
	// to control this behavior.
	// Learn more about [accessible name](https://w3c.github.io/accname/#dfn-accessible-name).
	Name interface{} `json:"name"`
	// An attribute that is usually set by `aria-pressed`.
	// Learn more about [`aria-pressed`](https://www.w3.org/TR/wai-aria-1.2/#aria-pressed).
	Pressed *bool `json:"pressed"`
	// An attribute that is usually set by `aria-selected`.
	// Learn more about [`aria-selected`](https://www.w3.org/TR/wai-aria-1.2/#aria-selected).
	Selected *bool `json:"selected"`
}
type PageGetByTextOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
}
type PageGetByTitleOptions struct {
	// Whether to find an exact match: case-sensitive and whole-string. Default to false.
	// Ignored when locating by a regular expression. Note that exact match still trims
	// whitespace.
	Exact *bool `json:"exact"`
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
	// `'networkidle'` - **DISCOURAGED** consider operation to be finished when there are
	// no network connections for at least `500` ms. Don't use this method for testing,
	// rely on web assertions to assess readiness instead.
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
	// `'networkidle'` - **DISCOURAGED** consider operation to be finished when there are
	// no network connections for at least `500` ms. Don't use this method for testing,
	// rely on web assertions to assess readiness instead.
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
	// `'networkidle'` - **DISCOURAGED** consider operation to be finished when there are
	// no network connections for at least `500` ms. Don't use this method for testing,
	// rely on web assertions to assess readiness instead.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type PageHoverOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type PageInnerHTMLOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageInnerTextOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageInputValueOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageIsCheckedOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageIsDisabledOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageIsEditableOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageIsEnabledOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageIsHiddenOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict  *bool    `json:"strict"`
	Timeout *float64 `json:"timeout"`
}
type PageIsVisibleOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict  *bool    `json:"strict"`
	Timeout *float64 `json:"timeout"`
}
type PageLocatorOptions struct {
	// Matches elements containing an element that matches an inner locator. Inner locator
	// is queried against the outer one. For example, `article` that has `text=Playwright`
	// matches `<article><div>Playwright</div></article>`.
	// Note that outer and inner locators must belong to the same frame. Inner locator
	// must not contain FrameLocators.
	Has Locator `json:"has"`
	// Matches elements that do not contain an element that matches an inner locator. Inner
	// locator is queried against the outer one. For example, `article` that does not have
	// `div` matches `<article><span>Playwright</span></article>`.
	// Note that outer and inner locators must belong to the same frame. Inner locator
	// must not contain FrameLocators.
	HasNot Locator `json:"hasNot"`
	// Matches elements that do not contain specified text somewhere inside, possibly in
	// a child or a descendant element. When passed a [string], matching is case-insensitive
	// and searches for a substring.
	HasNotText interface{} `json:"hasNotText"`
	// Matches elements containing specified text somewhere inside, possibly in a child
	// or a descendant element. When passed a [string], matching is case-insensitive and
	// searches for a substring. For example, `"Playwright"` matches `<article><div>Playwright</div></article>`.
	HasText interface{} `json:"hasText"`
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
	Margin *Margin `json:"margin"`
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
type Margin struct {
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
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageQuerySelectorOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
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
	// `'networkidle'` - **DISCOURAGED** consider operation to be finished when there are
	// no network connections for at least `500` ms. Don't use this method for testing,
	// rely on web assertions to assess readiness instead.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type PageRouteOptions struct {
	// How often a route should be used. By default it will be used every time.
	Times *int `json:"times"`
}
type PageRouteFromHAROptions struct {
	// If set to 'abort' any request not found in the HAR file will be aborted.
	// If set to 'fallback' missing requests will be sent to the network.
	// Defaults to abort.
	NotFound *HarNotFound `json:"notFound"`
	// If specified, updates the given HAR with the actual network information instead
	// of serving from file. The file is written to disk when BrowserContext.Close() is
	// called.
	Update *bool `json:"update"`
	// Optional setting to control resource content management. If `attach` is specified,
	// resources are persisted as separate files or entries in the ZIP archive. If `embed`
	// is specified, content is stored inline the HAR file.
	UpdateContent *RouteFromHarUpdateContentPolicy `json:"updateContent"`
	// When set to `minimal`, only record information necessary for routing from HAR. This
	// omits sizes, timing, page, cookies, security and other types of HAR information
	// that are not used when replaying from HAR. Defaults to `full`.
	UpdateMode *HarMode `json:"updateMode"`
	// A glob pattern, regular expression or predicate to match the request URL. Only requests
	// with URL matching the pattern will be served from the HAR file. If not specified,
	// all requests are served from the HAR file.
	URL interface{} `json:"url"`
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
	// When set to `"hide"`, screenshot will hide text caret. When set to `"initial"`,
	// text caret behavior will not be changed.  Defaults to `"hide"`.
	Caret *ScreenshotCaret `json:"caret"`
	// An object which specifies clipping of the resulting image.
	Clip *Rect `json:"clip"`
	// When true, takes a screenshot of the full scrollable page, instead of the currently
	// visible viewport. Defaults to `false`.
	FullPage *bool `json:"fullPage"`
	// Specify the color of the overlay box for masked elements, in [CSS color format](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value).
	// Default color is pink `#FF00FF`.
	MaskColor *string `json:"maskColor"`
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
	// will produce a single pixel per each device pixel, so screenshots of high-dpi devices
	// will be twice as large or even larger.
	// Defaults to `"device"`.
	Scale *ScreenshotScale `json:"scale"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// Specify screenshot type, defaults to `png`.
	Type *ScreenshotType `json:"type"`
}
type PageSelectOptionOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageSetCheckedOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *Position `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
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
	// `'networkidle'` - **DISCOURAGED** consider operation to be finished when there are
	// no network connections for at least `500` ms. Don't use this method for testing,
	// rely on web assertions to assess readiness instead.
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
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageTapOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
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
	Position *Position `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type PageTextContentOptions struct {
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
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
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageUncheckOptions struct {
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *Position `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}
type PageUnrouteOptions struct {
	// Optional handler function to route the request.
	Handler func(Route) `json:"handler"`
}
type PageExpectConsoleMessageOptions struct {
	// Receives the ConsoleMessage object and resolves to truthy value when the waiting
	// should resolve.
	Predicate func(ConsoleMessage) bool `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type PageExpectDownloadOptions struct {
	// Receives the Download object and resolves to truthy value when the waiting should
	// resolve.
	Predicate func(Download) bool `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type PageExpectEventOptions struct {
	// Receives the event data and resolves to truthy value when the waiting should resolve.
	Predicate interface{} `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type PageExpectFileChooserOptions struct {
	// Receives the FileChooser object and resolves to truthy value when the waiting should
	// resolve.
	Predicate func(FileChooser) bool `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type PageWaitForFunctionOptions struct {
	// If `polling` is `'raf'`, then `expression` is constantly executed in `requestAnimationFrame`
	// callback. If `polling` is a number, then it is treated as an interval in milliseconds
	// at which the function would be executed. Defaults to `raf`.
	Polling interface{} `json:"polling"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
}
type PageWaitForLoadStateOptions struct {
	// Optional load state to wait for, defaults to `load`. If the state has been already
	// reached while loading current document, the method resolves immediately. Can be
	// one of:
	// `'load'` - wait for the `load` event to be fired.
	// `'domcontentloaded'` - wait for the `DOMContentLoaded` event to be fired.
	// `'networkidle'` - **DISCOURAGED** wait until there are no network connections for
	// at least `500` ms. Don't use this method for testing, rely on web assertions to
	// assess readiness instead.
	State *LoadState `json:"state"`
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
	// for the navigation. Note that if the parameter is a string without wildcard characters,
	// the method will wait for navigation to URL that is exactly equal to the string.
	URL interface{} `json:"url"`
	// When to consider operation succeeded, defaults to `load`. Events can be either:
	// `'domcontentloaded'` - consider operation to be finished when the `DOMContentLoaded`
	// event is fired.
	// `'load'` - consider operation to be finished when the `load` event is fired.
	// `'networkidle'` - **DISCOURAGED** consider operation to be finished when there are
	// no network connections for at least `500` ms. Don't use this method for testing,
	// rely on web assertions to assess readiness instead.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type PageExpectPopupOptions struct {
	// Receives the Page object and resolves to truthy value when the waiting should resolve.
	Predicate func(Page) bool `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type PageWaitForRequestOptions struct {
	// Maximum wait time in milliseconds, defaults to 30 seconds, pass `0` to disable the
	// timeout. The default value can be changed by using the Page.SetDefaultTimeout()
	// method.
	Timeout *float64 `json:"timeout"`
}
type PageExpectRequestFinishedOptions struct {
	// Receives the Request object and resolves to truthy value when the waiting should
	// resolve.
	Predicate func(Request) bool `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
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
	// resolves to more than one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout()
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
	// `'networkidle'` - **DISCOURAGED** consider operation to be finished when there are
	// no network connections for at least `500` ms. Don't use this method for testing,
	// rely on web assertions to assess readiness instead.
	// `'commit'` - consider operation to be finished when network response is received
	// and the document started loading.
	WaitUntil *WaitUntilState `json:"waitUntil"`
}
type PageExpectWebSocketOptions struct {
	// Receives the WebSocket object and resolves to truthy value when the waiting should
	// resolve.
	Predicate func(WebSocket) bool `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type PageExpectWorkerOptions struct {
	// Receives the Worker object and resolves to truthy value when the waiting should
	// resolve.
	Predicate func(Worker) bool `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type PageWaitForEventOptions struct {
	// Receives the event data and resolves to truthy value when the waiting should resolve.
	Predicate interface{} `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type PageAssertionsToHaveTitleOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
}
type PageAssertionsToHaveURLOptions struct {
	// Time to retry the assertion for in milliseconds. Defaults to `5000`.
	Timeout *float64 `json:"timeout"`
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
	// Time immediately after the browser receives the first byte of the response from
	// the server, cache, or local resource. The value is given in milliseconds relative
	// to `startTime`, -1 if not available.
	ResponseStart float64 `json:"responseStart"`
	// Time immediately after the browser receives the last byte of the resource or immediately
	// before the transport connection is closed, whichever comes first. The value is given
	// in milliseconds relative to `startTime`, -1 if not available.
	ResponseEnd float64 `json:"responseEnd"`
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
	// If set changes the request method (e.g. GET or POST).
	Method *string `json:"method"`
	// If set changes the post data of request.
	PostData interface{} `json:"postData"`
	// If set changes the request URL. New URL must have same protocol as original one.
	URL *string `json:"url"`
}
type RouteFallbackOptions struct {
	// If set changes the request HTTP headers. Header values will be converted to a string.
	Headers map[string]string `json:"headers"`
	// If set changes the request method (e.g. GET or POST).
	Method *string `json:"method"`
	// If set changes the post data of request.
	PostData interface{} `json:"postData"`
	// If set changes the request URL. New URL must have same protocol as original one.
	// Changing the URL won't affect the route matching, all the routes are matched using
	// the original request URL.
	URL *string `json:"url"`
}
type RouteFetchOptions struct {
	// If set changes the request HTTP headers. Header values will be converted to a string.
	Headers map[string]string `json:"headers"`
	// Maximum number of request redirects that will be followed automatically. An error
	// will be thrown if the number is exceeded. Defaults to `20`. Pass `0` to not follow
	// redirects.
	MaxRedirects *int `json:"maxRedirects"`
	// If set changes the request method (e.g. GET or POST).
	Method *string `json:"method"`
	// If set changes the post data of request.
	PostData interface{} `json:"postData"`
	// Request timeout in milliseconds. Defaults to `30000` (30 seconds). Pass `0` to disable
	// timeout.
	Timeout *float64 `json:"timeout"`
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
	// APIResponse to fulfill route's request with. Individual fields of the response (such
	// as headers) can be overridden using fulfill options.
	Response APIResponse `json:"response"`
	// Response status code, defaults to `200`.
	Status *int `json:"status"`
}
type SelectorsRegisterOptions struct {
	// Whether to run this selector engine in isolated JavaScript environment. This environment
	// has access to the same DOM, but not any JavaScript objects from the frame's scripts.
	// Defaults to `false`. Note that running as a content script is not guaranteed when
	// this engine is used together with other registered engines.
	ContentScript *bool `json:"contentScript"`
	// Script that evaluates to a selector engine instance. The script is evaluated in
	// the page context.
	Path *string `json:"path"`
	// Script that evaluates to a selector engine instance. The script is evaluated in
	// the page context.
	Script *string `json:"script"`
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
	// Whether to include source files for trace actions.
	Sources *bool `json:"sources"`
	// Trace name to be shown in the Trace Viewer.
	Title *string `json:"title"`
}
type TracingStartChunkOptions struct {
	// If specified, the trace is going to be saved into the file with the given name inside
	// the `tracesDir` folder specified in BrowserType.Launch().
	Name *string `json:"name"`
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
	Payload interface{} `json:"payload"`
}
type FrameSentPayload struct {
	// frame payload
	Payload interface{} `json:"payload"`
}
type WebSocketExpectEventOptions struct {
	// Receives the event data and resolves to truthy value when the waiting should resolve.
	Predicate interface{} `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type WebSocketWaitForEventOptions struct {
	// Receives the event data and resolves to truthy value when the waiting should resolve.
	Predicate interface{} `json:"predicate"`
	// Maximum time to wait for in milliseconds. Defaults to `30000` (30 seconds). Pass
	// `0` to disable timeout. The default value can be changed by using the BrowserContext.SetDefaultTimeout().
	Timeout *float64 `json:"timeout"`
}
type WorkerEvaluateOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type WorkerEvaluateHandleOptions struct {
	// Optional argument to pass to `expression`.
	Arg interface{} `json:"arg"`
}
type OptionalStorageState struct {
	// Cookies to set for context
	Cookies []OptionalCookie `json:"cookies"`
	// localStorage to set for context
	Origins []OriginsState `json:"origins"`
}
type RecordVideoSize struct {
	// Video frame width.
	Width *int `json:"width"`
	// Video frame height.
	Height *int `json:"height"`
}
type OptionalCookie struct {
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
