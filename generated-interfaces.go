package playwright

// Exposes API that can be used for the Web API testing. This class is used for creating [APIRequestContext] instance
// which in turn can be used for sending web requests. An instance of this class can be obtained via
// [Playwright.Request]. For more information see [APIRequestContext].
type APIRequest interface {
	// Creates new instances of [APIRequestContext].
	NewContext(options ...APIRequestNewContextOptions) (APIRequestContext, error)
}

// This API is used for the Web API testing. You can use it to trigger API endpoints, configure micro-services,
// prepare environment or the service to your e2e test.
// Each Playwright browser context has associated with it [APIRequestContext] instance which shares cookie storage
// with the browser context and can be accessed via [BrowserContext.Request] or [Page.Request]. It is also possible to
// create a new APIRequestContext instance manually by calling [APIRequest.NewContext].
// **Cookie management**
// [APIRequestContext] returned by [BrowserContext.Request] and [Page.Request] shares cookie storage with the
// corresponding [BrowserContext]. Each API request will have `Cookie` header populated with the values from the
// browser context. If the API response contains `Set-Cookie` header it will automatically update [BrowserContext]
// cookies and requests made from the page will pick them up. This means that if you log in using this API, your e2e
// test will be logged in and vice versa.
// If you want API requests to not interfere with the browser cookies you should create a new [APIRequestContext] by
// calling [APIRequest.NewContext]. Such `APIRequestContext` object will have its own isolated cookie storage.
type APIRequestContext interface {
	// Sends HTTP(S) [DELETE] request and returns its
	// response. The method will populate request cookies from the context and update context cookies from the response.
	// The method will automatically follow redirects.
	//
	//  url: Target URL.
	//
	// [DELETE]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/DELETE
	Delete(url string, options ...APIRequestContextDeleteOptions) (APIResponse, error)

	// All responses returned by [APIRequestContext.Get] and similar methods are stored in the memory, so that you can
	// later call [APIResponse.Body].This method discards all its resources, calling any method on disposed
	// [APIRequestContext] will throw an exception.
	Dispose() error

	// Sends HTTP(S) request and returns its response. The method will populate request cookies from the context and
	// update context cookies from the response. The method will automatically follow redirects. JSON objects can be
	// passed directly to the request.
	//
	//  urlOrRequest: Target URL or Request to get all parameters from.
	Fetch(urlOrRequest interface{}, options ...APIRequestContextFetchOptions) (APIResponse, error)

	// Sends HTTP(S) [GET] request and returns its
	// response. The method will populate request cookies from the context and update context cookies from the response.
	// The method will automatically follow redirects.
	//
	//  url: Target URL.
	//
	// [GET]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET
	Get(url string, options ...APIRequestContextGetOptions) (APIResponse, error)

	// Sends HTTP(S) [HEAD] request and returns its
	// response. The method will populate request cookies from the context and update context cookies from the response.
	// The method will automatically follow redirects.
	//
	//  url: Target URL.
	//
	// [HEAD]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/HEAD
	Head(url string, options ...APIRequestContextHeadOptions) (APIResponse, error)

	// Sends HTTP(S) [PATCH] request and returns its
	// response. The method will populate request cookies from the context and update context cookies from the response.
	// The method will automatically follow redirects.
	//
	//  url: Target URL.
	//
	// [PATCH]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PATCH
	Patch(url string, options ...APIRequestContextPatchOptions) (APIResponse, error)

	// Sends HTTP(S) [POST] request and returns its
	// response. The method will populate request cookies from the context and update context cookies from the response.
	// The method will automatically follow redirects.
	//
	//  url: Target URL.
	//
	// [POST]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST
	Post(url string, options ...APIRequestContextPostOptions) (APIResponse, error)

	// Sends HTTP(S) [PUT] request and returns its
	// response. The method will populate request cookies from the context and update context cookies from the response.
	// The method will automatically follow redirects.
	//
	//  url: Target URL.
	//
	// [PUT]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PUT
	Put(url string, options ...APIRequestContextPutOptions) (APIResponse, error)

	// Returns storage state for this request context, contains current cookies and local storage snapshot if it was
	// passed to the constructor.
	StorageState(path ...string) (*StorageState, error)
}

// [APIResponse] class represents responses returned by [APIRequestContext.Get] and similar methods.
type APIResponse interface {
	// Returns the buffer with response body.
	Body() ([]byte, error)

	// Disposes the body of this response. If not called then the body will stay in memory until the context closes.
	Dispose() error

	// An object with all the response HTTP headers associated with this response.
	Headers() map[string]string

	// An array with all the request HTTP headers associated with this response. Header names are not lower-cased. Headers
	// with multiple entries, such as `Set-Cookie`, appear in the array multiple times.
	HeadersArray() []NameValue

	// Returns the JSON representation of response body.
	// This method will throw if the response body is not parsable via `JSON.parse`.
	JSON(v interface{}) error

	// Contains a boolean stating whether the response was successful (status in the range 200-299) or not.
	Ok() bool

	// Contains the status code of the response (e.g., 200 for a success).
	Status() int

	// Contains the status text of the response (e.g. usually an "OK" for a success).
	StatusText() string

	// Returns the text representation of response body.
	Text() (string, error)

	// Contains the URL of the response.
	URL() string
}

// The [APIResponseAssertions] class provides assertion methods that can be used to make assertions about the
// [APIResponse] in the tests.
type APIResponseAssertions interface {
	// Makes the assertion check for the opposite condition. For example, this code tests that the response status is not
	// successful:
	Not() APIResponseAssertions

	// Ensures the response status code is within `200..299` range.
	ToBeOK() error
}

// A Browser is created via [BrowserType.Launch]. An example of using a [Browser] to create a [Page]:
type Browser interface {
	EventEmitter
	// Emitted when Browser gets disconnected from the browser application. This might happen because of one of the
	// following:
	//  - Browser application is closed or crashed.
	//  - The [Browser.Close] method was called.
	OnDisconnected(fn func(Browser))

	// Get the browser type (chromium, firefox or webkit) that the browser belongs to.
	BrowserType() BrowserType

	// In case this browser is obtained using [BrowserType.Launch], closes the browser and all of its pages (if any were
	// opened).
	// In case this browser is connected to, clears all created contexts belonging to this browser and disconnects from
	// the browser server.
	// **NOTE** This is similar to force quitting the browser. Therefore, you should call [BrowserContext.Close] on any
	// [BrowserContext]'s you explicitly created earlier with [Browser.NewContext] **before** calling [Browser.Close].
	// The [Browser] object itself is considered to be disposed and cannot be used anymore.
	Close(options ...BrowserCloseOptions) error

	// Returns an array of all open browser contexts. In a newly created browser, this will return zero browser contexts.
	Contexts() []BrowserContext

	// Indicates that the browser is connected.
	IsConnected() bool

	// **NOTE** CDP Sessions are only supported on Chromium-based browsers.
	// Returns the newly created browser session.
	NewBrowserCDPSession() (CDPSession, error)

	// Creates a new browser context. It won't share cookies/cache with other browser contexts.
	// **NOTE** If directly using this method to create [BrowserContext]s, it is best practice to explicitly close the
	// returned context via [BrowserContext.Close] when your code is done with the [BrowserContext], and before calling
	// [Browser.Close]. This will ensure the `context` is closed gracefully and any artifacts—like HARs and videos—are
	// fully flushed and saved.
	NewContext(options ...BrowserNewContextOptions) (BrowserContext, error)

	// Creates a new page in a new browser context. Closing this page will close the context as well.
	// This is a convenience API that should only be used for the single-page scenarios and short snippets. Production
	// code and testing frameworks should explicitly create [Browser.NewContext] followed by the [BrowserContext.NewPage]
	// to control their exact life times.
	NewPage(options ...BrowserNewPageOptions) (Page, error)

	// **NOTE** This API controls
	// [Chromium Tracing] which is a low-level
	// chromium-specific debugging tool. API to control [Playwright Tracing] could be found
	// [here].
	// You can use [Browser.StartTracing] and [Browser.StopTracing] to create a trace file that can be opened in Chrome
	// DevTools performance panel.
	//
	// [Chromium Tracing]: https://www.chromium.org/developers/how-tos/trace-event-profiling-tool
	// [Playwright Tracing]: ../trace-viewer
	// [here]: ./class-tracing
	StartTracing(options ...BrowserStartTracingOptions) error

	// **NOTE** This API controls
	// [Chromium Tracing] which is a low-level
	// chromium-specific debugging tool. API to control [Playwright Tracing] could be found
	// [here].
	// Returns the buffer with trace data.
	//
	// [Chromium Tracing]: https://www.chromium.org/developers/how-tos/trace-event-profiling-tool
	// [Playwright Tracing]: ../trace-viewer
	// [here]: ./class-tracing
	StopTracing() ([]byte, error)

	// Returns the browser version.
	Version() string
}

//	BrowserContexts provide a way to operate multiple independent browser sessions.
//
// If a page opens another page, e.g. with a `window.open` call, the popup will belong to the parent page's browser
// context.
// Playwright allows creating "incognito" browser contexts with [Browser.NewContext] method. "Incognito" browser
// contexts don't write any browsing data to disk.
type BrowserContext interface {
	EventEmitter
	// Emitted when Browser context gets closed. This might happen because of one of the following:
	//  - Browser context is closed.
	//  - Browser application is closed or crashed.
	//  - The [Browser.Close] method was called.
	OnClose(fn func(BrowserContext))

	// Emitted when JavaScript within the page calls one of console API methods, e.g. `console.log` or `console.dir`. Also
	// emitted if the page throws an error or a warning.
	// The arguments passed into `console.log` and the page are available on the [ConsoleMessage] event handler argument.
	OnConsole(fn func(ConsoleMessage))

	// Emitted when a JavaScript dialog appears, such as `alert`, `prompt`, `confirm` or `beforeunload`. Listener **must**
	// either [Dialog.Accept] or [Dialog.Dismiss] the dialog - otherwise the page will
	// [freeze] waiting for the dialog,
	// and actions like click will never finish.
	//
	// [freeze]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/EventLoop#never_blocking
	OnDialog(fn func(Dialog))

	// The event is emitted when a new Page is created in the BrowserContext. The page may still be loading. The event
	// will also fire for popup pages. See also [Page.OnPopup] to receive events about popups relevant to a specific page.
	// The earliest moment that page is available is when it has navigated to the initial url. For example, when opening a
	// popup with `window.open('http://example.com')`, this event will fire when the network request to
	// "http://example.com" is done and its response has started loading in the popup.
	// **NOTE** Use [Page.WaitForLoadState] to wait until the page gets to a particular state (you should not need it in
	// most cases).
	OnPage(fn func(Page))

	// Emitted when exception is unhandled in any of the pages in this context. To listen for errors from a particular
	// page, use [Page.OnPageError] instead.
	OnWebError(fn func(WebError))

	// Emitted when a request is issued from any pages created through this context. The [request] object is read-only. To
	// only listen for requests from a particular page, use [Page.OnRequest].
	// In order to intercept and mutate requests, see [BrowserContext.Route] or [Page.Route].
	OnRequest(fn func(Request))

	// Emitted when a request fails, for example by timing out. To only listen for failed requests from a particular page,
	// use [Page.OnRequestFailed].
	// **NOTE** HTTP Error responses, such as 404 or 503, are still successful responses from HTTP standpoint, so request
	// will complete with [BrowserContext.OnRequestFinished] event and not with [BrowserContext.OnRequestFailed].
	OnRequestFailed(fn func(Request))

	// Emitted when a request finishes successfully after downloading the response body. For a successful response, the
	// sequence of events is `request`, `response` and `requestfinished`. To listen for successful requests from a
	// particular page, use [Page.OnRequestFinished].
	OnRequestFinished(fn func(Request))

	// Emitted when [response] status and headers are received for a request. For a successful response, the sequence of
	// events is `request`, `response` and `requestfinished`. To listen for response events from a particular page, use
	// [Page.OnResponse].
	OnResponse(fn func(Response))

	// Adds cookies into this browser context. All pages within this context will have these cookies installed. Cookies
	// can be obtained via [BrowserContext.Cookies].
	//
	//  cookies: Adds cookies to the browser context.
	//
	//    For the cookie to apply to all subdomains as well, prefix domain with a dot, like this: ".example.com".
	AddCookies(cookies []OptionalCookie) error

	// Adds a script which would be evaluated in one of the following scenarios:
	//  - Whenever a page is created in the browser context or is navigated.
	//  - Whenever a child frame is attached or navigated in any page in the browser context. In this case, the script is
	//   evaluated in the context of the newly attached frame.
	// The script is evaluated after the document was created but before any of its scripts were run. This is useful to
	// amend the JavaScript environment, e.g. to seed `Math.random`.
	//
	//  script: Script to be evaluated in all pages in the browser context.
	AddInitScript(script Script) error

	// **NOTE** Background pages are only supported on Chromium-based browsers.
	// All existing background pages in the context.
	BackgroundPages() []Page

	// Returns the browser instance of the context. If it was launched as a persistent context null gets returned.
	Browser() Browser

	// Clears context cookies.
	ClearCookies() error

	// Clears all permission overrides for the browser context.
	ClearPermissions() error

	// Closes the browser context. All the pages that belong to the browser context will be closed.
	// **NOTE** The default browser context cannot be closed.
	Close(options ...BrowserContextCloseOptions) error

	// If no URLs are specified, this method returns all cookies. If URLs are specified, only cookies that affect those
	// URLs are returned.
	Cookies(urls ...string) ([]Cookie, error)

	// The method adds a function called “name” on the `window` object of every frame in every page in the context. When
	// called, the function executes “callback” and returns a [Promise] which resolves to the return value of “callback”.
	// If the “callback” returns a [Promise], it will be awaited.
	// The first argument of the “callback” function contains information about the caller: `{ browserContext:
	// BrowserContext, page: Page, frame: Frame }`.
	// See [Page.ExposeBinding] for page-only version.
	//
	// 1. name: Name of the function on the window object.
	// 2. binding: Callback function that will be called in the Playwright's context.
	ExposeBinding(name string, binding BindingCallFunction, handle ...bool) error

	// The method adds a function called “name” on the `window` object of every frame in every page in the context. When
	// called, the function executes “callback” and returns a [Promise] which resolves to the return value of “callback”.
	// If the “callback” returns a [Promise], it will be awaited.
	// See [Page.ExposeFunction] for page-only version.
	//
	// 1. name: Name of the function on the window object.
	// 2. binding: Callback function that will be called in the Playwright's context.
	ExposeFunction(name string, binding ExposedFunction) error

	// Grants specified permissions to the browser context. Only grants corresponding permissions to the given origin if
	// specified.
	//
	//  permissions: A permission or an array of permissions to grant. Permissions can be one of the following values:
	//    - `'geolocation'`
	//    - `'midi'`
	//    - `'midi-sysex'` (system-exclusive midi)
	//    - `'notifications'`
	//    - `'camera'`
	//    - `'microphone'`
	//    - `'background-sync'`
	//    - `'ambient-light-sensor'`
	//    - `'accelerometer'`
	//    - `'gyroscope'`
	//    - `'magnetometer'`
	//    - `'accessibility-events'`
	//    - `'clipboard-read'`
	//    - `'clipboard-write'`
	//    - `'payment-handler'`
	GrantPermissions(permissions []string, options ...BrowserContextGrantPermissionsOptions) error

	// **NOTE** CDP sessions are only supported on Chromium-based browsers.
	// Returns the newly created session.
	//
	//  page: Target to create new session for. For backwards-compatibility, this parameter is named `page`, but it can be a
	//    `Page` or `Frame` type.
	NewCDPSession(page interface{}) (CDPSession, error)

	// Creates a new page in the browser context.
	NewPage() (Page, error)

	// Returns all open pages in the context.
	Pages() []Page

	// API testing helper associated with this context. Requests made with this API will use context cookies.
	Request() APIRequestContext

	// Routing provides the capability to modify network requests that are made by any page in the browser context. Once
	// route is enabled, every request matching the url pattern will stall unless it's continued, fulfilled or aborted.
	// **NOTE** [BrowserContext.Route] will not intercept requests intercepted by Service Worker. See
	// [this] issue. We recommend disabling Service Workers when
	// using request interception by setting “Browser.newContext.serviceWorkers” to `block`.
	//
	// 1. url: A glob pattern, regex pattern or predicate receiving [URL] to match while routing. When a “baseURL” via the context
	//    options was provided and the passed URL is a path, it gets merged via the
	//    [`new URL()`](https://developer.mozilla.org/en-US/docs/Web/API/URL/URL) constructor.
	// 2. handler: handler function to route the request.
	//
	// [this]: https://github.com/microsoft/playwright/issues/1090
	Route(url interface{}, handler routeHandler, times ...int) error

	// If specified the network requests that are made in the context will be served from the HAR file. Read more about
	// [Replaying from HAR].
	// Playwright will not serve requests intercepted by Service Worker from the HAR file. See
	// [this] issue. We recommend disabling Service Workers when
	// using request interception by setting “Browser.newContext.serviceWorkers” to `block`.
	//
	//  har: Path to a [HAR](http://www.softwareishard.com/blog/har-12-spec) file with prerecorded network data. If `path` is a
	//    relative path, then it is resolved relative to the current working directory.
	//
	// [Replaying from HAR]: https://playwright.dev/docs/mock#replaying-from-har
	// [this]: https://github.com/microsoft/playwright/issues/1090
	RouteFromHAR(har string, options ...BrowserContextRouteFromHAROptions) error

	// **NOTE** Service workers are only supported on Chromium-based browsers.
	// All existing service workers in the context.
	ServiceWorkers() []Worker

	// This setting will change the default maximum navigation time for the following methods and related shortcuts:
	//  - [Page.GoBack]
	//  - [Page.GoForward]
	//  - [Page.Goto]
	//  - [Page.Reload]
	//  - [Page.SetContent]
	//  - [Page.ExpectNavigation]
	// **NOTE** [Page.SetDefaultNavigationTimeout] and [Page.SetDefaultTimeout] take priority over
	// [BrowserContext.SetDefaultNavigationTimeout].
	//
	//  timeout: Maximum navigation time in milliseconds
	SetDefaultNavigationTimeout(timeout float64)

	// This setting will change the default maximum time for all the methods accepting “timeout” option.
	// **NOTE** [Page.SetDefaultNavigationTimeout], [Page.SetDefaultTimeout] and
	// [BrowserContext.SetDefaultNavigationTimeout] take priority over [BrowserContext.SetDefaultTimeout].
	//
	//  timeout: Maximum time in milliseconds
	SetDefaultTimeout(timeout float64)

	// The extra HTTP headers will be sent with every request initiated by any page in the context. These headers are
	// merged with page-specific extra HTTP headers set with [Page.SetExtraHTTPHeaders]. If page overrides a particular
	// header, page-specific header value will be used instead of the browser context header value.
	// **NOTE** [BrowserContext.SetExtraHTTPHeaders] does not guarantee the order of headers in the outgoing requests.
	//
	//  headers: An object containing additional HTTP headers to be sent with every request. All header values must be strings.
	SetExtraHTTPHeaders(headers map[string]string) error

	// Sets the context's geolocation. Passing `null` or `undefined` emulates position unavailable.
	SetGeolocation(geolocation *Geolocation) error

	//
	//  offline: Whether to emulate network being offline for the browser context.
	SetOffline(offline bool) error

	// Returns storage state for this browser context, contains current cookies and local storage snapshot.
	StorageState(path ...string) (*StorageState, error)

	Tracing() Tracing

	// Removes all routes created with [BrowserContext.Route] and [BrowserContext.RouteFromHAR].
	UnrouteAll(options ...BrowserContextUnrouteAllOptions) error

	// Removes a route created with [BrowserContext.Route]. When “handler” is not specified, removes all routes for the
	// “url”.
	//
	// 1. url: A glob pattern, regex pattern or predicate receiving [URL] used to register a routing with [BrowserContext.Route].
	// 2. handler: Optional handler function used to register a routing with [BrowserContext.Route].
	Unroute(url interface{}, handler ...routeHandler) error

	// Performs action and waits for a [ConsoleMessage] to be logged by in the pages in the context. If predicate is
	// provided, it passes [ConsoleMessage] value into the `predicate` function and waits for `predicate(message)` to
	// return a truthy value. Will throw an error if the page is closed before the [BrowserContext.OnConsole] event is
	// fired.
	ExpectConsoleMessage(cb func() error, options ...BrowserContextExpectConsoleMessageOptions) (ConsoleMessage, error)

	// Waits for event to fire and passes its value into the predicate function. Returns when the predicate returns truthy
	// value. Will throw an error if the context closes before the event is fired. Returns the event data value.
	//
	//  event: Event name, same one would pass into `browserContext.on(event)`.
	ExpectEvent(event string, cb func() error, options ...BrowserContextExpectEventOptions) (interface{}, error)

	// Performs action and waits for a new [Page] to be created in the context. If predicate is provided, it passes [Page]
	// value into the `predicate` function and waits for `predicate(event)` to return a truthy value. Will throw an error
	// if the context closes before new [Page] is created.
	ExpectPage(cb func() error, options ...BrowserContextExpectPageOptions) (Page, error)

	// **NOTE** In most cases, you should use [BrowserContext.ExpectEvent].
	// Waits for given `event` to fire. If predicate is provided, it passes event's value into the `predicate` function
	// and waits for `predicate(event)` to return a truthy value. Will throw an error if the browser context is closed
	// before the `event` is fired.
	//
	//  event: Event name, same one typically passed into `*.on(event)`.
	WaitForEvent(event string, options ...BrowserContextWaitForEventOptions) (interface{}, error)
}

// BrowserType provides methods to launch a specific browser instance or connect to an existing one. The following is
// a typical example of using Playwright to drive automation:
type BrowserType interface {
	// This method attaches Playwright to an existing browser instance. When connecting to another browser launched via
	// `BrowserType.launchServer` in Node.js, the major and minor version needs to match the client version (1.2.3 → is
	// compatible with 1.2.x).
	//
	//  wsEndpoint: A browser websocket endpoint to connect to.
	Connect(wsEndpoint string, options ...BrowserTypeConnectOptions) (Browser, error)

	// This method attaches Playwright to an existing browser instance using the Chrome DevTools Protocol.
	// The default browser context is accessible via [Browser.Contexts].
	// **NOTE** Connecting over the Chrome DevTools Protocol is only supported for Chromium-based browsers.
	//
	//  endpointURL: A CDP websocket endpoint or http url to connect to. For example `http://localhost:9222/` or
	//    `ws://127.0.0.1:9222/devtools/browser/387adf4c-243f-4051-a181-46798f4a46f4`.
	ConnectOverCDP(endpointURL string, options ...BrowserTypeConnectOverCDPOptions) (Browser, error)

	// A path where Playwright expects to find a bundled browser executable.
	ExecutablePath() string

	// Returns the browser instance.
	//
	// [Chrome Canary]: https://www.google.com/chrome/browser/canary.html
	// [Dev Channel]: https://www.chromium.org/getting-involved/dev-channel
	// [this article]: https://www.howtogeek.com/202825/what%E2%80%99s-the-difference-between-chromium-and-chrome/
	// [This article]: https://chromium.googlesource.com/chromium/src/+/lkgr/docs/chromium_browser_vs_google_chrome.md
	Launch(options ...BrowserTypeLaunchOptions) (Browser, error)

	// Returns the persistent browser context instance.
	// Launches browser that uses persistent storage located at “userDataDir” and returns the only context. Closing this
	// context will automatically close the browser.
	//
	//  userDataDir: Path to a User Data Directory, which stores browser session data like cookies and local storage. More details for
	//    [Chromium](https://chromium.googlesource.com/chromium/src/+/master/docs/user_data_dir.md#introduction) and
	//    [Firefox](https://developer.mozilla.org/en-US/docs/Mozilla/Command_Line_Options#User_Profile). Note that Chromium's
	//    user data directory is the **parent** directory of the "Profile Path" seen at `chrome://version`. Pass an empty
	//    string to use a temporary directory instead.
	LaunchPersistentContext(userDataDir string, options ...BrowserTypeLaunchPersistentContextOptions) (BrowserContext, error)

	// Returns browser name. For example: `chromium`, `webkit` or `firefox`.
	Name() string
}

//	The `CDPSession` instances are used to talk raw Chrome Devtools Protocol:
//	- protocol methods can be called with `session.send` method.
//	- protocol events can be subscribed to with `session.on` method.
//
// Useful links:
//   - Documentation on DevTools Protocol can be found here:
//     [DevTools Protocol Viewer].
//   - Getting Started with DevTools Protocol:
//     https://github.com/aslushnikov/getting-started-with-cdp/blob/master/README.md
//
// [DevTools Protocol Viewer]: https://chromedevtools.github.io/devtools-protocol/
type CDPSession interface {
	EventEmitter
	// Detaches the CDPSession from the target. Once detached, the CDPSession object won't emit any events and can't be
	// used to send messages.
	Detach() error

	//
	// 1. method: Protocol method name.
	// 2. params: Optional method parameters.
	Send(method string, params map[string]interface{}) (interface{}, error)
}

// [ConsoleMessage] objects are dispatched by page via the [Page.OnConsole] event. For each console messages logged in
// the page there will be corresponding event in the Playwright context.
type ConsoleMessage interface {
	// List of arguments passed to a `console` function call. See also [Page.OnConsole].
	Args() []JSHandle

	Location() *ConsoleMessageLocation

	// The page that produced this console message, if any.
	Page() Page

	// The text of the console message.
	Text() string

	// The text of the console message.
	String() string

	// One of the following values: `log`, `debug`, `info`, `error`, `warning`, `dir`, `dirxml`, `table`,
	// `trace`, `clear`, `startGroup`, `startGroupCollapsed`, `endGroup`, `assert`, `profile`,
	// `profileEnd`, `count`, `timeEnd`.
	Type() string
}

// [Dialog] objects are dispatched by page via the [Page.OnDialog] event.
// An example of using `Dialog` class:
// **NOTE** Dialogs are dismissed automatically, unless there is a [Page.OnDialog] listener. When listener is present,
// it **must** either [Dialog.Accept] or [Dialog.Dismiss] the dialog - otherwise the page will
// [freeze] waiting for the dialog,
// and actions like click will never finish.
//
// [freeze]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/EventLoop#never_blocking
type Dialog interface {
	// Returns when the dialog has been accepted.
	Accept(promptText ...string) error

	// If dialog is prompt, returns default prompt value. Otherwise, returns empty string.
	DefaultValue() string

	// Returns when the dialog has been dismissed.
	Dismiss() error

	// A message displayed in the dialog.
	Message() string

	// The page that initiated this dialog, if available.
	Page() Page

	// Returns dialog's type, can be one of `alert`, `beforeunload`, `confirm` or `prompt`.
	Type() string
}

// [Download] objects are dispatched by page via the [Page.OnDownload] event.
// All the downloaded files belonging to the browser context are deleted when the browser context is closed.
// Download event is emitted once the download starts. Download path becomes available once download completes.
type Download interface {
	// Cancels a download. Will not fail if the download is already finished or canceled. Upon successful cancellations,
	// `download.failure()` would resolve to `canceled`.
	Cancel() error

	// Deletes the downloaded file. Will wait for the download to finish if necessary.
	Delete() error

	// Returns download error if any. Will wait for the download to finish if necessary.
	Failure() error

	// Get the page that the download belongs to.
	Page() Page

	// Returns path to the downloaded file for a successful download, or throws for a failed/canceled download. The method
	// will wait for the download to finish if necessary. The method throws when connected remotely.
	// Note that the download's file name is a random GUID, use [Download.SuggestedFilename] to get suggested file name.
	Path() (string, error)

	// Copy the download to a user-specified path. It is safe to call this method while the download is still in progress.
	// Will wait for the download to finish if necessary.
	//
	//  path: Path where the download should be copied.
	SaveAs(path string) error

	// Returns suggested filename for this download. It is typically computed by the browser from the
	// [`Content-Disposition`] response
	// header or the `download` attribute. See the spec on [whatwg].
	// Different browsers can use different logic for computing it.
	//
	// [`Content-Disposition`]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition
	// [whatwg]: https://html.spec.whatwg.org/#downloading-resources
	SuggestedFilename() string

	// Returns downloaded url.
	URL() string

	String() string
}

//	ElementHandle represents an in-page DOM element. ElementHandles can be created with the [Page.QuerySelector]
//
// method.
// **NOTE** The use of ElementHandle is discouraged, use [Locator] objects and web-first assertions instead.
// ElementHandle prevents DOM element from garbage collection unless the handle is disposed with [JSHandle.Dispose].
// ElementHandles are auto-disposed when their origin frame gets navigated.
// ElementHandle instances can be used as an argument in [Page.EvalOnSelector] and [Page.Evaluate] methods.
// The difference between the [Locator] and ElementHandle is that the ElementHandle points to a particular element,
// while [Locator] captures the logic of how to retrieve an element.
// In the example below, handle points to a particular DOM element on page. If that element changes text or is used by
// React to render an entirely different component, handle is still pointing to that very DOM element. This can lead
// to unexpected behaviors.
// With the locator, every time the `element` is used, up-to-date DOM element is located in the page using the
// selector. So in the snippet below, underlying DOM element is going to be located twice.
type ElementHandle interface {
	JSHandle
	// This method returns the bounding box of the element, or `null` if the element is not visible. The bounding box is
	// calculated relative to the main frame viewport - which is usually the same as the browser window.
	// Scrolling affects the returned bounding box, similarly to
	// [Element.GetBoundingClientRect].
	// That means `x` and/or `y` may be negative.
	// Elements from child frames return the bounding box relative to the main frame, unlike the
	// [Element.GetBoundingClientRect].
	// Assuming the page is static, it is safe to use bounding box coordinates to perform input. For example, the
	// following snippet should click the center of the element.
	//
	// [Element.GetBoundingClientRect]: https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
	// [Element.GetBoundingClientRect]: https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
	BoundingBox() (*Rect, error)

	// This method checks the element by performing the following steps:
	//  1. Ensure that element is a checkbox or a radio input. If not, this method throws. If the element is already
	//    checked, this method returns immediately.
	//  2. Wait for [actionability] checks on the element, unless “force” option is set.
	//  3. Scroll the element into view if needed.
	//  4. Use [Page.Mouse] to click in the center of the element.
	//  5. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	//  6. Ensure that the element is now checked. If not, this method throws.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.Check] instead. Read more about [locators].
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Check(options ...ElementHandleCheckOptions) error

	// This method clicks the element by performing the following steps:
	//  1. Wait for [actionability] checks on the element, unless “force” option is set.
	//  2. Scroll the element into view if needed.
	//  3. Use [Page.Mouse] to click in the center of the element, or the specified “position”.
	//  4. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.Click] instead. Read more about [locators].
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Click(options ...ElementHandleClickOptions) error

	// Returns the content frame for element handles referencing iframe nodes, or `null` otherwise
	ContentFrame() (Frame, error)

	// This method double clicks the element by performing the following steps:
	//  1. Wait for [actionability] checks on the element, unless “force” option is set.
	//  2. Scroll the element into view if needed.
	//  3. Use [Page.Mouse] to double click in the center of the element, or the specified “position”.
	//  4. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set. Note that if
	//    the first click of the `dblclick()` triggers a navigation event, this method will throw.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	// **NOTE** `elementHandle.dblclick()` dispatches two `click` events and a single `dblclick` event.
	//
	// Deprecated: Use locator-based [Locator.Dblclick] instead. Read more about [locators].
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Dblclick(options ...ElementHandleDblclickOptions) error

	// The snippet below dispatches the `click` event on the element. Regardless of the visibility state of the element,
	// `click` is dispatched. This is equivalent to calling
	// [element.Click()].
	//
	// Deprecated: Use locator-based [Locator.DispatchEvent] instead. Read more about [locators].
	//
	// 1. typ: DOM event type: `"click"`, `"dragstart"`, etc.
	// 2. eventInit: Optional event-specific initialization properties.
	//
	// [element.Click()]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click
	// [DeviceMotionEvent]: https://developer.mozilla.org/en-US/docs/Web/API/DeviceMotionEvent/DeviceMotionEvent
	// [DeviceOrientationEvent]: https://developer.mozilla.org/en-US/docs/Web/API/DeviceOrientationEvent/DeviceOrientationEvent
	// [DragEvent]: https://developer.mozilla.org/en-US/docs/Web/API/DragEvent/DragEvent
	// [Event]: https://developer.mozilla.org/en-US/docs/Web/API/Event/Event
	// [FocusEvent]: https://developer.mozilla.org/en-US/docs/Web/API/FocusEvent/FocusEvent
	// [KeyboardEvent]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/KeyboardEvent
	// [MouseEvent]: https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/MouseEvent
	// [PointerEvent]: https://developer.mozilla.org/en-US/docs/Web/API/PointerEvent/PointerEvent
	// [TouchEvent]: https://developer.mozilla.org/en-US/docs/Web/API/TouchEvent/TouchEvent
	// [WheelEvent]: https://developer.mozilla.org/en-US/docs/Web/API/WheelEvent/WheelEvent
	// [locators]: https://playwright.dev/docs/locators
	DispatchEvent(typ string, eventInit ...interface{}) error

	// Returns the return value of “expression”.
	// The method finds an element matching the specified selector in the `ElementHandle`s subtree and passes it as a
	// first argument to “expression”. If no elements match the selector, the method throws an error.
	// If “expression” returns a [Promise], then [ElementHandle.EvalOnSelector] would wait for the promise to resolve and
	// return its value.
	//
	// Deprecated: This method does not wait for the element to pass actionability checks and therefore can lead to the flaky tests. Use [Locator.Evaluate], other [Locator] helper methods or web-first assertions instead.
	//
	// 1. selector: A selector to query for.
	// 2. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 3. arg: Optional argument to pass to “expression”.
	EvalOnSelector(selector string, expression string, arg ...interface{}) (interface{}, error)

	// Returns the return value of “expression”.
	// The method finds all elements matching the specified selector in the `ElementHandle`'s subtree and passes an array
	// of matched elements as a first argument to “expression”.
	// If “expression” returns a [Promise], then [ElementHandle.EvalOnSelectorAll] would wait for the promise to resolve
	// and return its value.
	//
	// Deprecated: In most cases, [Locator.EvaluateAll], other [Locator] helper methods and web-first assertions do a better job.
	//
	// 1. selector: A selector to query for.
	// 2. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 3. arg: Optional argument to pass to “expression”.
	EvalOnSelectorAll(selector string, expression string, arg ...interface{}) (interface{}, error)

	// This method waits for [actionability] checks, focuses the element, fills it and triggers an
	// `input` event after filling. Note that you can pass an empty string to clear the input field.
	// If the target element is not an `<input>`, `<textarea>` or `[contenteditable]` element, this method throws an
	// error. However, if the element is inside the `<label>` element that has an associated
	// [control], the control will be filled
	// instead.
	// To send fine-grained keyboard events, use [Locator.PressSequentially].
	//
	// Deprecated: Use locator-based [Locator.Fill] instead. Read more about [locators].
	//
	//  value: Value to set for the `<input>`, `<textarea>` or `[contenteditable]` element.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	// [locators]: https://playwright.dev/docs/locators
	Fill(value string, options ...ElementHandleFillOptions) error

	// Calls [focus] on the element.
	//
	// Deprecated: Use locator-based [Locator.Focus] instead. Read more about [locators].
	//
	// [focus]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus
	// [locators]: https://playwright.dev/docs/locators
	Focus() error

	// Returns element attribute value.
	//
	// Deprecated: Use locator-based [Locator.GetAttribute] instead. Read more about [locators].
	//
	//  name: Attribute name to get the value for.
	//
	// [locators]: https://playwright.dev/docs/locators
	GetAttribute(name string) (string, error)

	// This method hovers over the element by performing the following steps:
	//  1. Wait for [actionability] checks on the element, unless “force” option is set.
	//  2. Scroll the element into view if needed.
	//  3. Use [Page.Mouse] to hover over the center of the element, or the specified “position”.
	//  4. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.Hover] instead. Read more about [locators].
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Hover(options ...ElementHandleHoverOptions) error

	// Returns the `element.innerHTML`.
	//
	// Deprecated: Use locator-based [Locator.InnerHTML] instead. Read more about [locators].
	//
	// [locators]: https://playwright.dev/docs/locators
	InnerHTML() (string, error)

	// Returns the `element.innerText`.
	//
	// Deprecated: Use locator-based [Locator.InnerText] instead. Read more about [locators].
	//
	// [locators]: https://playwright.dev/docs/locators
	InnerText() (string, error)

	// Returns `input.value` for the selected `<input>` or `<textarea>` or `<select>` element.
	// Throws for non-input elements. However, if the element is inside the `<label>` element that has an associated
	// [control], returns the value of the
	// control.
	//
	// Deprecated: Use locator-based [Locator.InputValue] instead. Read more about [locators].
	//
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	// [locators]: https://playwright.dev/docs/locators
	InputValue(options ...ElementHandleInputValueOptions) (string, error)

	// Returns whether the element is checked. Throws if the element is not a checkbox or radio input.
	//
	// Deprecated: Use locator-based [Locator.IsChecked] instead. Read more about [locators].
	//
	// [locators]: https://playwright.dev/docs/locators
	IsChecked() (bool, error)

	// Returns whether the element is disabled, the opposite of [enabled].
	//
	// Deprecated: Use locator-based [Locator.IsDisabled] instead. Read more about [locators].
	//
	// [enabled]: https://playwright.dev/docs/actionability#enabled
	// [locators]: https://playwright.dev/docs/locators
	IsDisabled() (bool, error)

	// Returns whether the element is [editable].
	//
	// Deprecated: Use locator-based [Locator.IsEditable] instead. Read more about [locators].
	//
	// [editable]: https://playwright.dev/docs/actionability#editable
	// [locators]: https://playwright.dev/docs/locators
	IsEditable() (bool, error)

	// Returns whether the element is [enabled].
	//
	// Deprecated: Use locator-based [Locator.IsEnabled] instead. Read more about [locators].
	//
	// [enabled]: https://playwright.dev/docs/actionability#enabled
	// [locators]: https://playwright.dev/docs/locators
	IsEnabled() (bool, error)

	// Returns whether the element is hidden, the opposite of [visible].
	//
	// Deprecated: Use locator-based [Locator.IsHidden] instead. Read more about [locators].
	//
	// [visible]: https://playwright.dev/docs/actionability#visible
	// [locators]: https://playwright.dev/docs/locators
	IsHidden() (bool, error)

	// Returns whether the element is [visible].
	//
	// Deprecated: Use locator-based [Locator.IsVisible] instead. Read more about [locators].
	//
	// [visible]: https://playwright.dev/docs/actionability#visible
	// [locators]: https://playwright.dev/docs/locators
	IsVisible() (bool, error)

	// Returns the frame containing the given element.
	OwnerFrame() (Frame, error)

	// Focuses the element, and then uses [Keyboard.Down] and [Keyboard.Up].
	// “key” can specify the intended
	// [keyboardEvent.Key] value or a single character
	// to generate the text for. A superset of the “key” values can be found
	// [here]. Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`,
	// etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the “key” in the upper case.
	// If “key” is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When specified with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	//
	// Deprecated: Use locator-based [Locator.Press] instead. Read more about [locators].
	//
	//  key: Name of the key to press or a character to generate, such as `ArrowLeft` or `a`.
	//
	// [keyboardEvent.Key]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key
	// [here]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values
	// [locators]: https://playwright.dev/docs/locators
	Press(key string, options ...ElementHandlePressOptions) error

	// The method finds an element matching the specified selector in the `ElementHandle`'s subtree. If no elements match
	// the selector, returns `null`.
	//
	// Deprecated: Use locator-based [Page.Locator] instead. Read more about [locators].
	//
	//  selector: A selector to query for.
	//
	// [locators]: https://playwright.dev/docs/locators
	QuerySelector(selector string) (ElementHandle, error)

	// The method finds all elements matching the specified selector in the `ElementHandle`s subtree. If no elements match
	// the selector, returns empty array.
	//
	// Deprecated: Use locator-based [Page.Locator] instead. Read more about [locators].
	//
	//  selector: A selector to query for.
	//
	// [locators]: https://playwright.dev/docs/locators
	QuerySelectorAll(selector string) ([]ElementHandle, error)

	// This method captures a screenshot of the page, clipped to the size and position of this particular element. If the
	// element is covered by other elements, it will not be actually visible on the screenshot. If the element is a
	// scrollable container, only the currently scrolled content will be visible on the screenshot.
	// This method waits for the [actionability] checks, then scrolls element into view before taking
	// a screenshot. If the element is detached from DOM, the method throws an error.
	// Returns the buffer with the captured screenshot.
	//
	// Deprecated: Use locator-based [Locator.Screenshot] instead. Read more about [locators].
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Screenshot(options ...ElementHandleScreenshotOptions) ([]byte, error)

	// This method waits for [actionability] checks, then tries to scroll element into view, unless
	// it is completely visible as defined by
	// [IntersectionObserver]'s `ratio`.
	// Throws when `elementHandle` does not point to an element
	// [connected] to a Document or a ShadowRoot.
	//
	// Deprecated: Use locator-based [Locator.ScrollIntoViewIfNeeded] instead. Read more about [locators].
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [IntersectionObserver]: https://developer.mozilla.org/en-US/docs/Web/API/Intersection_Observer_API
	// [connected]: https://developer.mozilla.org/en-US/docs/Web/API/Node/isConnected
	// [locators]: https://playwright.dev/docs/locators
	ScrollIntoViewIfNeeded(options ...ElementHandleScrollIntoViewIfNeededOptions) error

	// This method waits for [actionability] checks, waits until all specified options are present in
	// the `<select>` element and selects these options.
	// If the target element is not a `<select>` element, this method throws an error. However, if the element is inside
	// the `<label>` element that has an associated
	// [control], the control will be used
	// instead.
	// Returns the array of option values that have been successfully selected.
	// Triggers a `change` and `input` event once all the provided options have been selected.
	//
	// Deprecated: Use locator-based [Locator.SelectOption] instead. Read more about [locators].
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	// [locators]: https://playwright.dev/docs/locators
	SelectOption(values SelectOptionValues, options ...ElementHandleSelectOptionOptions) ([]string, error)

	// This method waits for [actionability] checks, then focuses the element and selects all its
	// text content.
	// If the element is inside the `<label>` element that has an associated
	// [control], focuses and selects text in
	// the control instead.
	//
	// Deprecated: Use locator-based [Locator.SelectText] instead. Read more about [locators].
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	// [locators]: https://playwright.dev/docs/locators
	SelectText(options ...ElementHandleSelectTextOptions) error

	// This method checks or unchecks an element by performing the following steps:
	//  1. Ensure that element is a checkbox or a radio input. If not, this method throws.
	//  2. If the element already has the right checked state, this method returns immediately.
	//  3. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  4. Scroll the element into view if needed.
	//  5. Use [Page.Mouse] to click in the center of the element.
	//  6. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	//  7. Ensure that the element is now checked or unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.SetChecked] instead. Read more about [locators].
	//
	//  checked: Whether to check or uncheck the checkbox.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	SetChecked(checked bool, options ...ElementHandleSetCheckedOptions) error

	// Sets the value of the file input to these file paths or files. If some of the `filePaths` are relative paths, then
	// they are resolved relative to the current working directory. For empty array, clears the selected files.
	// This method expects [ElementHandle] to point to an
	// [input element]. However, if the element is inside
	// the `<label>` element that has an associated
	// [control], targets the control instead.
	//
	// Deprecated: Use locator-based [Locator.SetInputFiles] instead. Read more about [locators].
	//
	// [input element]: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	// [locators]: https://playwright.dev/docs/locators
	SetInputFiles(files interface{}, options ...ElementHandleSetInputFilesOptions) error

	// This method taps the element by performing the following steps:
	//  1. Wait for [actionability] checks on the element, unless “force” option is set.
	//  2. Scroll the element into view if needed.
	//  3. Use [Page.Touchscreen] to tap the center of the element, or the specified “position”.
	//  4. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	// **NOTE** `elementHandle.tap()` requires that the `hasTouch` option of the browser context be set to true.
	//
	// Deprecated: Use locator-based [Locator.Tap] instead. Read more about [locators].
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Tap(options ...ElementHandleTapOptions) error

	// Returns the `node.textContent`.
	//
	// Deprecated: Use locator-based [Locator.TextContent] instead. Read more about [locators].
	//
	// [locators]: https://playwright.dev/docs/locators
	TextContent() (string, error)

	// Focuses the element, and then sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the
	// text.
	// To press a special key, like `Control` or `ArrowDown`, use [ElementHandle.Press].
	//
	// Deprecated: In most cases, you should use [Locator.Fill] instead. You only need to press keys one by one if there is special keyboard handling on the page - in this case use [Locator.PressSequentially].
	//
	//  text: A text to type into a focused element.
	Type(text string, options ...ElementHandleTypeOptions) error

	// This method checks the element by performing the following steps:
	//  1. Ensure that element is a checkbox or a radio input. If not, this method throws. If the element is already
	//    unchecked, this method returns immediately.
	//  2. Wait for [actionability] checks on the element, unless “force” option is set.
	//  3. Scroll the element into view if needed.
	//  4. Use [Page.Mouse] to click in the center of the element.
	//  5. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	//  6. Ensure that the element is now unchecked. If not, this method throws.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.Uncheck] instead. Read more about [locators].
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Uncheck(options ...ElementHandleUncheckOptions) error

	// Returns when the element satisfies the “state”.
	// Depending on the “state” parameter, this method waits for one of the [actionability] checks to
	// pass. This method throws when the element is detached while waiting, unless waiting for the `"hidden"` state.
	//  - `"visible"` Wait until the element is [visible].
	//  - `"hidden"` Wait until the element is [not visible] or not attached. Note that
	//   waiting for hidden does not throw when the element detaches.
	//  - `"stable"` Wait until the element is both [visible] and
	//   [stable].
	//  - `"enabled"` Wait until the element is [enabled].
	//  - `"disabled"` Wait until the element is [not enabled].
	//  - `"editable"` Wait until the element is [editable].
	// If the element does not satisfy the condition for the “timeout” milliseconds, this method will throw.
	//
	//  state: A state to wait for, see below for more details.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [visible]: https://playwright.dev/docs/actionability#visible
	// [not visible]: https://playwright.dev/docs/actionability#visible
	// [visible]: https://playwright.dev/docs/actionability#visible
	// [stable]: https://playwright.dev/docs/actionability#stable
	// [enabled]: https://playwright.dev/docs/actionability#enabled
	// [not enabled]: https://playwright.dev/docs/actionability#enabled
	// [editable]: https://playwright.dev/docs/actionability#editable
	WaitForElementState(state ElementState, options ...ElementHandleWaitForElementStateOptions) error

	// Returns element specified by selector when it satisfies “state” option. Returns `null` if waiting for `hidden` or
	// `detached`.
	// Wait for the “selector” relative to the element handle to satisfy “state” option (either appear/disappear from dom,
	// or become visible/hidden). If at the moment of calling the method “selector” already satisfies the condition, the
	// method will return immediately. If the selector doesn't satisfy the condition for the “timeout” milliseconds, the
	// function will throw.
	//
	// Deprecated: Use web assertions that assert visibility or a locator-based [Locator.WaitFor] instead.
	//
	//  selector: A selector to query for.
	WaitForSelector(selector string, options ...ElementHandleWaitForSelectorOptions) (ElementHandle, error)
}

// [FileChooser] objects are dispatched by the page in the [Page.OnFileChooser] event.
type FileChooser interface {
	// Returns input element associated with this file chooser.
	Element() ElementHandle

	// Returns whether this file chooser accepts multiple files.
	IsMultiple() bool

	// Returns page this file chooser belongs to.
	Page() Page

	// Sets the value of the file input this chooser is associated with. If some of the `filePaths` are relative paths,
	// then they are resolved relative to the current working directory. For empty array, clears the selected files.
	SetFiles(files interface{}, options ...FileChooserSetFilesOptions) error
}

// At every point of time, page exposes its current frame tree via the [Page.MainFrame] and [Frame.ChildFrames]
// methods.
// [Frame] object's lifecycle is controlled by three events, dispatched on the page object:
//   - [Page.OnFrameAttached] - fired when the frame gets attached to the page. A Frame can be attached to the page
//     only once.
//   - [Page.OnFrameNavigated] - fired when the frame commits navigation to a different URL.
//   - [Page.OnFrameDetached] - fired when the frame gets detached from the page.  A Frame can be detached from the
//     page only once.
//
// An example of dumping frame tree:
type Frame interface {
	// Returns the added tag when the script's onload fires or when the script content was injected into frame.
	// Adds a `<script>` tag into the page with the desired url or content.
	AddScriptTag(options FrameAddScriptTagOptions) (ElementHandle, error)

	// Returns the added tag when the stylesheet's onload fires or when the CSS content was injected into frame.
	// Adds a `<link rel="stylesheet">` tag into the page with the desired url or a `<style type="text/css">` tag with the
	// content.
	AddStyleTag(options FrameAddStyleTagOptions) (ElementHandle, error)

	// This method checks an element matching “selector” by performing the following steps:
	//  1. Find an element matching “selector”. If there is none, wait until a matching element is attached to the DOM.
	//  2. Ensure that matched element is a checkbox or a radio input. If not, this method throws. If the element is
	//    already checked, this method returns immediately.
	//  3. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  4. Scroll the element into view if needed.
	//  5. Use [Page.Mouse] to click in the center of the element.
	//  6. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	//  7. Ensure that the element is now checked. If not, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.Check] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Check(selector string, options ...FrameCheckOptions) error

	ChildFrames() []Frame

	// This method clicks an element matching “selector” by performing the following steps:
	//  1. Find an element matching “selector”. If there is none, wait until a matching element is attached to the DOM.
	//  2. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  3. Scroll the element into view if needed.
	//  4. Use [Page.Mouse] to click in the center of the element, or the specified “position”.
	//  5. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.Click] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Click(selector string, options ...FrameClickOptions) error

	// Gets the full HTML contents of the frame, including the doctype.
	Content() (string, error)

	// This method double clicks an element matching “selector” by performing the following steps:
	//  1. Find an element matching “selector”. If there is none, wait until a matching element is attached to the DOM.
	//  2. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  3. Scroll the element into view if needed.
	//  4. Use [Page.Mouse] to double click in the center of the element, or the specified “position”.
	//  5. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set. Note that if
	//    the first click of the `dblclick()` triggers a navigation event, this method will throw.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	// **NOTE** `frame.dblclick()` dispatches two `click` events and a single `dblclick` event.
	//
	// Deprecated: Use locator-based [Locator.Dblclick] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Dblclick(selector string, options ...FrameDblclickOptions) error

	// The snippet below dispatches the `click` event on the element. Regardless of the visibility state of the element,
	// `click` is dispatched. This is equivalent to calling
	// [element.Click()].
	//
	// Deprecated: Use locator-based [Locator.DispatchEvent] instead. Read more about [locators].
	//
	// 1. selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	// 2. typ: DOM event type: `"click"`, `"dragstart"`, etc.
	// 3. eventInit: Optional event-specific initialization properties.
	//
	// [element.Click()]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click
	// [DeviceMotionEvent]: https://developer.mozilla.org/en-US/docs/Web/API/DeviceMotionEvent/DeviceMotionEvent
	// [DeviceOrientationEvent]: https://developer.mozilla.org/en-US/docs/Web/API/DeviceOrientationEvent/DeviceOrientationEvent
	// [DragEvent]: https://developer.mozilla.org/en-US/docs/Web/API/DragEvent/DragEvent
	// [Event]: https://developer.mozilla.org/en-US/docs/Web/API/Event/Event
	// [FocusEvent]: https://developer.mozilla.org/en-US/docs/Web/API/FocusEvent/FocusEvent
	// [KeyboardEvent]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/KeyboardEvent
	// [MouseEvent]: https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/MouseEvent
	// [PointerEvent]: https://developer.mozilla.org/en-US/docs/Web/API/PointerEvent/PointerEvent
	// [TouchEvent]: https://developer.mozilla.org/en-US/docs/Web/API/TouchEvent/TouchEvent
	// [WheelEvent]: https://developer.mozilla.org/en-US/docs/Web/API/WheelEvent/WheelEvent
	// [locators]: https://playwright.dev/docs/locators
	DispatchEvent(selector string, typ string, eventInit interface{}, options ...FrameDispatchEventOptions) error

	//
	// 1. source: A selector to search for an element to drag. If there are multiple elements satisfying the selector, the first will
	//    be used.
	// 2. target: A selector to search for an element to drop onto. If there are multiple elements satisfying the selector, the first
	//    will be used.
	DragAndDrop(source string, target string, options ...FrameDragAndDropOptions) error

	// Returns the return value of “expression”.
	// The method finds an element matching the specified selector within the frame and passes it as a first argument to
	// “expression”. If no elements match the selector, the method throws an error.
	// If “expression” returns a [Promise], then [Frame.EvalOnSelector] would wait for the promise to resolve and return
	// its value.
	//
	// Deprecated: This method does not wait for the element to pass the actionability checks and therefore can lead to the flaky tests. Use [Locator.Evaluate], other [Locator] helper methods or web-first assertions instead.
	//
	// 1. selector: A selector to query for.
	// 2. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 3. arg: Optional argument to pass to “expression”.
	EvalOnSelector(selector string, expression string, arg interface{}, options ...FrameEvalOnSelectorOptions) (interface{}, error)

	// Returns the return value of “expression”.
	// The method finds all elements matching the specified selector within the frame and passes an array of matched
	// elements as a first argument to “expression”.
	// If “expression” returns a [Promise], then [Frame.EvalOnSelectorAll] would wait for the promise to resolve and
	// return its value.
	//
	// Deprecated: In most cases, [Locator.EvaluateAll], other [Locator] helper methods and web-first assertions do a better job.
	//
	// 1. selector: A selector to query for.
	// 2. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 3. arg: Optional argument to pass to “expression”.
	EvalOnSelectorAll(selector string, expression string, arg ...interface{}) (interface{}, error)

	// Returns the return value of “expression”.
	// If the function passed to the [Frame.Evaluate] returns a [Promise], then [Frame.Evaluate] would wait for the
	// promise to resolve and return its value.
	// If the function passed to the [Frame.Evaluate] returns a non-[Serializable] value, then [Frame.Evaluate] returns
	// `undefined`. Playwright also supports transferring some additional values that are not serializable by `JSON`:
	// `-0`, `NaN`, `Infinity`, `-Infinity`.
	//
	// 1. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 2. arg: Optional argument to pass to “expression”.
	Evaluate(expression string, arg ...interface{}) (interface{}, error)

	// Returns the return value of “expression” as a [JSHandle].
	// The only difference between [Frame.Evaluate] and [Frame.EvaluateHandle] is that [Frame.EvaluateHandle] returns
	// [JSHandle].
	// If the function, passed to the [Frame.EvaluateHandle], returns a [Promise], then [Frame.EvaluateHandle] would wait
	// for the promise to resolve and return its value.
	//
	// 1. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 2. arg: Optional argument to pass to “expression”.
	EvaluateHandle(expression string, arg ...interface{}) (JSHandle, error)

	// This method waits for an element matching “selector”, waits for [actionability] checks,
	// focuses the element, fills it and triggers an `input` event after filling. Note that you can pass an empty string
	// to clear the input field.
	// If the target element is not an `<input>`, `<textarea>` or `[contenteditable]` element, this method throws an
	// error. However, if the element is inside the `<label>` element that has an associated
	// [control], the control will be filled
	// instead.
	// To send fine-grained keyboard events, use [Locator.PressSequentially].
	//
	// Deprecated: Use locator-based [Locator.Fill] instead. Read more about [locators].
	//
	// 1. selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	// 2. value: Value to fill for the `<input>`, `<textarea>` or `[contenteditable]` element.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	// [locators]: https://playwright.dev/docs/locators
	Fill(selector string, value string, options ...FrameFillOptions) error

	// This method fetches an element with “selector” and focuses it. If there's no element matching “selector”, the
	// method waits until a matching element appears in the DOM.
	//
	// Deprecated: Use locator-based [Locator.Focus] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [locators]: https://playwright.dev/docs/locators
	Focus(selector string, options ...FrameFocusOptions) error

	// Returns the `frame` or `iframe` element handle which corresponds to this frame.
	// This is an inverse of [ElementHandle.ContentFrame]. Note that returned handle actually belongs to the parent frame.
	// This method throws an error if the frame has been detached before `frameElement()` returns.
	FrameElement() (ElementHandle, error)

	// When working with iframes, you can create a frame locator that will enter the iframe and allow selecting elements
	// in that iframe.
	//
	//  selector: A selector to use when resolving DOM element.
	FrameLocator(selector string) FrameLocator

	// Returns element attribute value.
	//
	// Deprecated: Use locator-based [Locator.GetAttribute] instead. Read more about [locators].
	//
	// 1. selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	// 2. name: Attribute name to get the value for.
	//
	// [locators]: https://playwright.dev/docs/locators
	GetAttribute(selector string, name string, options ...FrameGetAttributeOptions) (string, error)

	// Allows locating elements by their alt text.
	//
	//  text: Text to locate the element for.
	GetByAltText(text interface{}, options ...FrameGetByAltTextOptions) Locator

	// Allows locating input elements by the text of the associated `<label>` or `aria-labelledby` element, or by the
	// `aria-label` attribute.
	//
	//  text: Text to locate the element for.
	GetByLabel(text interface{}, options ...FrameGetByLabelOptions) Locator

	// Allows locating input elements by the placeholder text.
	//
	//  text: Text to locate the element for.
	GetByPlaceholder(text interface{}, options ...FrameGetByPlaceholderOptions) Locator

	// Allows locating elements by their [ARIA role],
	// [ARIA attributes] and
	// [accessible name].
	//
	// # Details
	//
	// Role selector **does not replace** accessibility audits and conformance tests, but rather gives early feedback
	// about the ARIA guidelines.
	// Many html elements have an implicitly [defined role]
	// that is recognized by the role selector. You can find all the
	// [supported roles here]. ARIA guidelines **do not recommend**
	// duplicating implicit roles and attributes by setting `role` and/or `aria-*` attributes to default values.
	//
	//  role: Required aria role.
	//
	// [ARIA role]: https://www.w3.org/TR/wai-aria-1.2/#roles
	// [ARIA attributes]: https://www.w3.org/TR/wai-aria-1.2/#aria-attributes
	// [accessible name]: https://w3c.github.io/accname/#dfn-accessible-name
	// [defined role]: https://w3c.github.io/html-aam/#html-element-role-mappings
	// [supported roles here]: https://www.w3.org/TR/wai-aria-1.2/#role_definitions
	GetByRole(role AriaRole, options ...FrameGetByRoleOptions) Locator

	// Locate element by the test id.
	//
	// # Details
	//
	// By default, the `data-testid` attribute is used as a test id. Use [Selectors.SetTestIdAttribute] to configure a
	// different test id attribute if necessary.
	//
	//  testId: Id to locate the element by.
	GetByTestId(testId interface{}) Locator

	// Allows locating elements that contain given text.
	// See also [Locator.Filter] that allows to match by another criteria, like an accessible role, and then filter by the
	// text content.
	//
	// # Details
	//
	// Matching by text always normalizes whitespace, even with exact match. For example, it turns multiple spaces into
	// one, turns line breaks into spaces and ignores leading and trailing whitespace.
	// Input elements of the type `button` and `submit` are matched by their `value` instead of the text content. For
	// example, locating by text `"Log in"` matches `<input type=button value="Log in">`.
	//
	//  text: Text to locate the element for.
	GetByText(text interface{}, options ...FrameGetByTextOptions) Locator

	// Allows locating elements by their title attribute.
	//
	//  text: Text to locate the element for.
	GetByTitle(text interface{}, options ...FrameGetByTitleOptions) Locator

	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of
	// the last redirect.
	// The method will throw an error if:
	//  - there's an SSL error (e.g. in case of self-signed certificates).
	//  - target URL is invalid.
	//  - the “timeout” is exceeded during navigation.
	//  - the remote server does not respond or is unreachable.
	//  - the main resource failed to load.
	// The method will not throw an error when any valid HTTP status code is returned by the remote server, including 404
	// "Not Found" and 500 "Internal Server Error".  The status code for such responses can be retrieved by calling
	// [Response.Status].
	// **NOTE** The method either throws an error or returns a main resource response. The only exceptions are navigation
	// to `about:blank` or navigation to the same URL with a different hash, which would succeed and return `null`.
	// **NOTE** Headless mode doesn't support navigation to a PDF document. See the
	// [upstream issue].
	//
	//  url: URL to navigate frame to. The url should include scheme, e.g. `https://`.
	//
	// [upstream issue]: https://bugs.chromium.org/p/chromium/issues/detail?id=761295
	Goto(url string, options ...FrameGotoOptions) (Response, error)

	// This method hovers over an element matching “selector” by performing the following steps:
	//  1. Find an element matching “selector”. If there is none, wait until a matching element is attached to the DOM.
	//  2. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  3. Scroll the element into view if needed.
	//  4. Use [Page.Mouse] to hover over the center of the element, or the specified “position”.
	//  5. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.Hover] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Hover(selector string, options ...FrameHoverOptions) error

	// Returns `element.innerHTML`.
	//
	// Deprecated: Use locator-based [Locator.InnerHTML] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [locators]: https://playwright.dev/docs/locators
	InnerHTML(selector string, options ...FrameInnerHTMLOptions) (string, error)

	// Returns `element.innerText`.
	//
	// Deprecated: Use locator-based [Locator.InnerText] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [locators]: https://playwright.dev/docs/locators
	InnerText(selector string, options ...FrameInnerTextOptions) (string, error)

	// Returns `input.value` for the selected `<input>` or `<textarea>` or `<select>` element.
	// Throws for non-input elements. However, if the element is inside the `<label>` element that has an associated
	// [control], returns the value of the
	// control.
	//
	// Deprecated: Use locator-based [Locator.InputValue] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	// [locators]: https://playwright.dev/docs/locators
	InputValue(selector string, options ...FrameInputValueOptions) (string, error)

	// Returns whether the element is checked. Throws if the element is not a checkbox or radio input.
	//
	// Deprecated: Use locator-based [Locator.IsChecked] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [locators]: https://playwright.dev/docs/locators
	IsChecked(selector string, options ...FrameIsCheckedOptions) (bool, error)

	// Returns `true` if the frame has been detached, or `false` otherwise.
	IsDetached() bool

	// Returns whether the element is disabled, the opposite of [enabled].
	//
	// Deprecated: Use locator-based [Locator.IsDisabled] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [enabled]: https://playwright.dev/docs/actionability#enabled
	// [locators]: https://playwright.dev/docs/locators
	IsDisabled(selector string, options ...FrameIsDisabledOptions) (bool, error)

	// Returns whether the element is [editable].
	//
	// Deprecated: Use locator-based [Locator.IsEditable] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [editable]: https://playwright.dev/docs/actionability#editable
	// [locators]: https://playwright.dev/docs/locators
	IsEditable(selector string, options ...FrameIsEditableOptions) (bool, error)

	// Returns whether the element is [enabled].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [enabled]: https://playwright.dev/docs/actionability#enabled
	IsEnabled(selector string, options ...FrameIsEnabledOptions) (bool, error)

	// Returns whether the element is hidden, the opposite of [visible].  “selector” that
	// does not match any elements is considered hidden.
	//
	// Deprecated: Use locator-based [Locator.IsHidden] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [visible]: https://playwright.dev/docs/actionability#visible
	// [locators]: https://playwright.dev/docs/locators
	IsHidden(selector string, options ...FrameIsHiddenOptions) (bool, error)

	// Returns whether the element is [visible]. “selector” that does not match any elements
	// is considered not visible.
	//
	// Deprecated: Use locator-based [Locator.IsVisible] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [visible]: https://playwright.dev/docs/actionability#visible
	// [locators]: https://playwright.dev/docs/locators
	IsVisible(selector string, options ...FrameIsVisibleOptions) (bool, error)

	// The method returns an element locator that can be used to perform actions on this page / frame. Locator is resolved
	// to the element immediately before performing an action, so a series of actions on the same locator can in fact be
	// performed on different DOM elements. That would happen if the DOM structure between those actions has changed.
	// [Learn more about locators].
	// [Learn more about locators].
	//
	//  selector: A selector to use when resolving DOM element.
	//
	// [Learn more about locators]: https://playwright.dev/docs/locators
	// [Learn more about locators]: https://playwright.dev/docs/locators
	Locator(selector string, options ...FrameLocatorOptions) Locator

	// Returns frame's name attribute as specified in the tag.
	// If the name is empty, returns the id attribute instead.
	// **NOTE** This value is calculated once when the frame is created, and will not update if the attribute is changed
	// later.
	Name() string

	// Returns the page containing this frame.
	Page() Page

	// Parent frame, if any. Detached frames and main frames return `null`.
	ParentFrame() Frame

	// “key” can specify the intended
	// [keyboardEvent.Key] value or a single character
	// to generate the text for. A superset of the “key” values can be found
	// [here]. Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`,
	// etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the “key” in the upper case.
	// If “key” is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When specified with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	//
	// Deprecated: Use locator-based [Locator.Press] instead. Read more about [locators].
	//
	// 1. selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	// 2. key: Name of the key to press or a character to generate, such as `ArrowLeft` or `a`.
	//
	// [keyboardEvent.Key]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key
	// [here]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values
	// [locators]: https://playwright.dev/docs/locators
	Press(selector string, key string, options ...FramePressOptions) error

	// Returns the ElementHandle pointing to the frame element.
	// **NOTE** The use of [ElementHandle] is discouraged, use [Locator] objects and web-first assertions instead.
	// The method finds an element matching the specified selector within the frame. If no elements match the selector,
	// returns `null`.
	//
	// Deprecated: Use locator-based [Frame.Locator] instead. Read more about [locators].
	//
	//  selector: A selector to query for.
	//
	// [locators]: https://playwright.dev/docs/locators
	QuerySelector(selector string, options ...FrameQuerySelectorOptions) (ElementHandle, error)

	// Returns the ElementHandles pointing to the frame elements.
	// **NOTE** The use of [ElementHandle] is discouraged, use [Locator] objects instead.
	// The method finds all elements matching the specified selector within the frame. If no elements match the selector,
	// returns empty array.
	//
	// Deprecated: Use locator-based [Frame.Locator] instead. Read more about [locators].
	//
	//  selector: A selector to query for.
	//
	// [locators]: https://playwright.dev/docs/locators
	QuerySelectorAll(selector string) ([]ElementHandle, error)

	// This method waits for an element matching “selector”, waits for [actionability] checks, waits
	// until all specified options are present in the `<select>` element and selects these options.
	// If the target element is not a `<select>` element, this method throws an error. However, if the element is inside
	// the `<label>` element that has an associated
	// [control], the control will be used
	// instead.
	// Returns the array of option values that have been successfully selected.
	// Triggers a `change` and `input` event once all the provided options have been selected.
	//
	// Deprecated: Use locator-based [Locator.SelectOption] instead. Read more about [locators].
	//
	//  selector: A selector to query for.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	// [locators]: https://playwright.dev/docs/locators
	SelectOption(selector string, values SelectOptionValues, options ...FrameSelectOptionOptions) ([]string, error)

	// This method checks or unchecks an element matching “selector” by performing the following steps:
	//  1. Find an element matching “selector”. If there is none, wait until a matching element is attached to the DOM.
	//  2. Ensure that matched element is a checkbox or a radio input. If not, this method throws.
	//  3. If the element already has the right checked state, this method returns immediately.
	//  4. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  5. Scroll the element into view if needed.
	//  6. Use [Page.Mouse] to click in the center of the element.
	//  7. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	//  8. Ensure that the element is now checked or unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.SetChecked] instead. Read more about [locators].
	//
	// 1. selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	// 2. checked: Whether to check or uncheck the checkbox.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	SetChecked(selector string, checked bool, options ...FrameSetCheckedOptions) error

	// This method internally calls [document.Write()],
	// inheriting all its specific characteristics and behaviors.
	//
	//  html: HTML markup to assign to the page.
	//
	// [document.Write()]: https://developer.mozilla.org/en-US/docs/Web/API/Document/write
	SetContent(html string, options ...FrameSetContentOptions) error

	// Sets the value of the file input to these file paths or files. If some of the `filePaths` are relative paths, then
	// they are resolved relative to the current working directory. For empty array, clears the selected files.
	// This method expects “selector” to point to an
	// [input element]. However, if the element is inside
	// the `<label>` element that has an associated
	// [control], targets the control instead.
	//
	// Deprecated: Use locator-based [Locator.SetInputFiles] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [input element]: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	// [locators]: https://playwright.dev/docs/locators
	SetInputFiles(selector string, files interface{}, options ...FrameSetInputFilesOptions) error

	// This method taps an element matching “selector” by performing the following steps:
	//  1. Find an element matching “selector”. If there is none, wait until a matching element is attached to the DOM.
	//  2. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  3. Scroll the element into view if needed.
	//  4. Use [Page.Touchscreen] to tap the center of the element, or the specified “position”.
	//  5. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	// **NOTE** `frame.tap()` requires that the `hasTouch` option of the browser context be set to true.
	//
	// Deprecated: Use locator-based [Locator.Tap] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Tap(selector string, options ...FrameTapOptions) error

	// Returns `element.textContent`.
	//
	// Deprecated: Use locator-based [Locator.TextContent] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [locators]: https://playwright.dev/docs/locators
	TextContent(selector string, options ...FrameTextContentOptions) (string, error)

	// Returns the page title.
	Title() (string, error)

	// Sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the text. `frame.type` can be used
	// to send fine-grained keyboard events. To fill values in form fields, use [Frame.Fill].
	// To press a special key, like `Control` or `ArrowDown`, use [Keyboard.Press].
	//
	// Deprecated: In most cases, you should use [Locator.Fill] instead. You only need to press keys one by one if there is special keyboard handling on the page - in this case use [Locator.PressSequentially].
	//
	// 1. selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	// 2. text: A text to type into a focused element.
	Type(selector string, text string, options ...FrameTypeOptions) error

	// This method checks an element matching “selector” by performing the following steps:
	//  1. Find an element matching “selector”. If there is none, wait until a matching element is attached to the DOM.
	//  2. Ensure that matched element is a checkbox or a radio input. If not, this method throws. If the element is
	//    already unchecked, this method returns immediately.
	//  3. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  4. Scroll the element into view if needed.
	//  5. Use [Page.Mouse] to click in the center of the element.
	//  6. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	//  7. Ensure that the element is now unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.Uncheck] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Uncheck(selector string, options ...FrameUncheckOptions) error

	// Returns frame's url.
	URL() string

	// Returns when the “expression” returns a truthy value, returns that value.
	//
	// 1. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 2. arg: Optional argument to pass to “expression”.
	WaitForFunction(expression string, arg interface{}, options ...FrameWaitForFunctionOptions) (JSHandle, error)

	// Waits for the required load state to be reached.
	// This returns when the frame reaches a required load state, `load` by default. The navigation must have been
	// committed when this method is called. If current document has already reached the required state, resolves
	// immediately.
	WaitForLoadState(options ...FrameWaitForLoadStateOptions) error

	// Waits for the frame navigation and returns the main resource response. In case of multiple redirects, the
	// navigation will resolve with the response of the last redirect. In case of navigation to a different anchor or
	// navigation due to History API usage, the navigation will resolve with `null`.
	//
	// Deprecated: This method is inherently racy, please use [Frame.WaitForURL] instead.
	//
	// [History API]: https://developer.mozilla.org/en-US/docs/Web/API/History_API
	ExpectNavigation(cb func() error, options ...FrameExpectNavigationOptions) (Response, error)

	// Returns when element specified by selector satisfies “state” option. Returns `null` if waiting for `hidden` or
	// `detached`.
	// **NOTE** Playwright automatically waits for element to be ready before performing an action. Using [Locator]
	// objects and web-first assertions make the code wait-for-selector-free.
	// Wait for the “selector” to satisfy “state” option (either appear/disappear from dom, or become visible/hidden). If
	// at the moment of calling the method “selector” already satisfies the condition, the method will return immediately.
	// If the selector doesn't satisfy the condition for the “timeout” milliseconds, the function will throw.
	//
	// Deprecated: Use web assertions that assert visibility or a locator-based [Locator.WaitFor] instead. Read more about [locators].
	//
	//  selector: A selector to query for.
	//
	// [locators]: https://playwright.dev/docs/locators
	WaitForSelector(selector string, options ...FrameWaitForSelectorOptions) (ElementHandle, error)

	// Waits for the given “timeout” in milliseconds.
	// Note that `frame.waitForTimeout()` should only be used for debugging. Tests using the timer in production are going
	// to be flaky. Use signals such as network events, selectors becoming visible and others instead.
	//
	// Deprecated: Never wait for timeout in production. Tests that wait for time are inherently flaky. Use [Locator] actions and web assertions that wait automatically.
	//
	//  timeout: A timeout to wait for
	WaitForTimeout(timeout float64)

	// Waits for the frame to navigate to the given URL.
	//
	//  url: A glob pattern, regex pattern or predicate receiving [URL] to match while waiting for the navigation. Note that if
	//    the parameter is a string without wildcard characters, the method will wait for navigation to URL that is exactly
	//    equal to the string.
	WaitForURL(url interface{}, options ...FrameWaitForURLOptions) error
}

// FrameLocator represents a view to the `iframe` on the page. It captures the logic sufficient to retrieve the
// `iframe` and locate elements in that iframe. FrameLocator can be created with either [Page.FrameLocator] or
// [Locator.FrameLocator] method.
// **Strictness**
// Frame locators are strict. This means that all operations on frame locators will throw if more than one element
// matches a given selector.
// **Converting Locator to FrameLocator**
// If you have a [Locator] object pointing to an `iframe` it can be converted to [FrameLocator] using
// [`:scope`] CSS selector:
//
// [`:scope`]: https://developer.mozilla.org/en-US/docs/Web/CSS/:scope
type FrameLocator interface {
	// Returns locator to the first matching frame.
	First() FrameLocator

	// When working with iframes, you can create a frame locator that will enter the iframe and allow selecting elements
	// in that iframe.
	//
	//  selector: A selector to use when resolving DOM element.
	FrameLocator(selector string) FrameLocator

	// Allows locating elements by their alt text.
	//
	//  text: Text to locate the element for.
	GetByAltText(text interface{}, options ...FrameLocatorGetByAltTextOptions) Locator

	// Allows locating input elements by the text of the associated `<label>` or `aria-labelledby` element, or by the
	// `aria-label` attribute.
	//
	//  text: Text to locate the element for.
	GetByLabel(text interface{}, options ...FrameLocatorGetByLabelOptions) Locator

	// Allows locating input elements by the placeholder text.
	//
	//  text: Text to locate the element for.
	GetByPlaceholder(text interface{}, options ...FrameLocatorGetByPlaceholderOptions) Locator

	// Allows locating elements by their [ARIA role],
	// [ARIA attributes] and
	// [accessible name].
	//
	// # Details
	//
	// Role selector **does not replace** accessibility audits and conformance tests, but rather gives early feedback
	// about the ARIA guidelines.
	// Many html elements have an implicitly [defined role]
	// that is recognized by the role selector. You can find all the
	// [supported roles here]. ARIA guidelines **do not recommend**
	// duplicating implicit roles and attributes by setting `role` and/or `aria-*` attributes to default values.
	//
	//  role: Required aria role.
	//
	// [ARIA role]: https://www.w3.org/TR/wai-aria-1.2/#roles
	// [ARIA attributes]: https://www.w3.org/TR/wai-aria-1.2/#aria-attributes
	// [accessible name]: https://w3c.github.io/accname/#dfn-accessible-name
	// [defined role]: https://w3c.github.io/html-aam/#html-element-role-mappings
	// [supported roles here]: https://www.w3.org/TR/wai-aria-1.2/#role_definitions
	GetByRole(role AriaRole, options ...FrameLocatorGetByRoleOptions) Locator

	// Locate element by the test id.
	//
	// # Details
	//
	// By default, the `data-testid` attribute is used as a test id. Use [Selectors.SetTestIdAttribute] to configure a
	// different test id attribute if necessary.
	//
	//  testId: Id to locate the element by.
	GetByTestId(testId interface{}) Locator

	// Allows locating elements that contain given text.
	// See also [Locator.Filter] that allows to match by another criteria, like an accessible role, and then filter by the
	// text content.
	//
	// # Details
	//
	// Matching by text always normalizes whitespace, even with exact match. For example, it turns multiple spaces into
	// one, turns line breaks into spaces and ignores leading and trailing whitespace.
	// Input elements of the type `button` and `submit` are matched by their `value` instead of the text content. For
	// example, locating by text `"Log in"` matches `<input type=button value="Log in">`.
	//
	//  text: Text to locate the element for.
	GetByText(text interface{}, options ...FrameLocatorGetByTextOptions) Locator

	// Allows locating elements by their title attribute.
	//
	//  text: Text to locate the element for.
	GetByTitle(text interface{}, options ...FrameLocatorGetByTitleOptions) Locator

	// Returns locator to the last matching frame.
	Last() FrameLocator

	// The method finds an element matching the specified selector in the locator's subtree. It also accepts filter
	// options, similar to [Locator.Filter] method.
	// [Learn more about locators].
	//
	//  selectorOrLocator: A selector or locator to use when resolving DOM element.
	//
	// [Learn more about locators]: https://playwright.dev/docs/locators
	Locator(selectorOrLocator interface{}, options ...FrameLocatorLocatorOptions) Locator

	// Returns locator to the n-th matching frame. It's zero based, `nth(0)` selects the first frame.
	Nth(index int) FrameLocator
}

// JSHandle represents an in-page JavaScript object. JSHandles can be created with the [Page.EvaluateHandle] method.
// JSHandle prevents the referenced JavaScript object being garbage collected unless the handle is exposed with
// [JSHandle.Dispose]. JSHandles are auto-disposed when their origin frame gets navigated or the parent context gets
// destroyed.
// JSHandle instances can be used as an argument in [Page.EvalOnSelector], [Page.Evaluate] and [Page.EvaluateHandle]
// methods.
type JSHandle interface {
	// Returns either `null` or the object handle itself, if the object handle is an instance of [ElementHandle].
	AsElement() ElementHandle

	// The `jsHandle.dispose` method stops referencing the element handle.
	Dispose() error

	// Returns the return value of “expression”.
	// This method passes this handle as the first argument to “expression”.
	// If “expression” returns a [Promise], then `handle.evaluate` would wait for the promise to resolve and return its
	// value.
	//
	// 1. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 2. arg: Optional argument to pass to “expression”.
	Evaluate(expression string, arg ...interface{}) (interface{}, error)

	// Returns the return value of “expression” as a [JSHandle].
	// This method passes this handle as the first argument to “expression”.
	// The only difference between `jsHandle.evaluate` and `jsHandle.evaluateHandle` is that `jsHandle.evaluateHandle`
	// returns [JSHandle].
	// If the function passed to the `jsHandle.evaluateHandle` returns a [Promise], then `jsHandle.evaluateHandle` would
	// wait for the promise to resolve and return its value.
	// See [Page.EvaluateHandle] for more details.
	//
	// 1. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 2. arg: Optional argument to pass to “expression”.
	EvaluateHandle(expression string, arg ...interface{}) (JSHandle, error)

	// The method returns a map with **own property names** as keys and JSHandle instances for the property values.
	GetProperties() (map[string]JSHandle, error)

	// Fetches a single property from the referenced object.
	//
	//  propertyName: property to get
	GetProperty(propertyName string) (JSHandle, error)

	// Returns a JSON representation of the object. If the object has a `toJSON` function, it **will not be called**.
	// **NOTE** The method will return an empty JSON object if the referenced object is not stringifiable. It will throw
	// an error if the object has circular references.
	JSONValue() (interface{}, error)

	String() string
}

// Keyboard provides an api for managing a virtual keyboard. The high level api is [Keyboard.Type], which takes raw
// characters and generates proper `keydown`, `keypress`/`input`, and `keyup` events on your page.
// For finer control, you can use [Keyboard.Down], [Keyboard.Up], and [Keyboard.InsertText] to manually fire events as
// if they were generated from a real keyboard.
// An example of holding down `Shift` in order to select and delete some text:
// An example of pressing uppercase `A`
// An example to trigger select-all with the keyboard
type Keyboard interface {
	// Dispatches a `keydown` event.
	// “key” can specify the intended
	// [keyboardEvent.Key] value or a single character
	// to generate the text for. A superset of the “key” values can be found
	// [here]. Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`,
	// etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the “key” in the upper case.
	// If “key” is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// If “key” is a modifier key, `Shift`, `Meta`, `Control`, or `Alt`, subsequent key presses will be sent with that
	// modifier active. To release the modifier key, use [Keyboard.Up].
	// After the key is pressed once, subsequent calls to [Keyboard.Down] will have
	// [repeat] set to true. To release the key,
	// use [Keyboard.Up].
	// **NOTE** Modifier keys DO influence `keyboard.down`. Holding down `Shift` will type the text in upper case.
	//
	//  key: Name of the key to press or a character to generate, such as `ArrowLeft` or `a`.
	//
	// [keyboardEvent.Key]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key
	// [here]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values
	// [repeat]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/repeat
	Down(key string) error

	// Dispatches only `input` event, does not emit the `keydown`, `keyup` or `keypress` events.
	//
	//  text: Sets input to the specified text value.
	InsertText(text string) error

	// **NOTE** In most cases, you should use [Locator.Press] instead.
	// “key” can specify the intended
	// [keyboardEvent.Key] value or a single character
	// to generate the text for. A superset of the “key” values can be found
	// [here]. Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`,
	// etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the “key” in the upper case.
	// If “key” is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When specified with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	//
	//  key: Name of the key to press or a character to generate, such as `ArrowLeft` or `a`.
	//
	// [keyboardEvent.Key]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key
	// [here]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values
	Press(key string, options ...KeyboardPressOptions) error

	// **NOTE** In most cases, you should use [Locator.Fill] instead. You only need to press keys one by one if there is
	// special keyboard handling on the page - in this case use [Locator.PressSequentially].
	// Sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the text.
	// To press a special key, like `Control` or `ArrowDown`, use [Keyboard.Press].
	//
	//  text: A text to type into a focused element.
	Type(text string, options ...KeyboardTypeOptions) error

	// Dispatches a `keyup` event.
	//
	//  key: Name of the key to press or a character to generate, such as `ArrowLeft` or `a`.
	Up(key string) error
}

// Locators are the central piece of Playwright's auto-waiting and retry-ability. In a nutshell, locators represent a
// way to find element(s) on the page at any moment. A locator can be created with the [Page.Locator] method.
// [Learn more about locators].
//
// [Learn more about locators]: https://playwright.dev/docs/locators
type Locator interface {
	// When the locator points to a list of elements, this returns an array of locators, pointing to their respective
	// elements.
	// **NOTE** [Locator.All] does not wait for elements to match the locator, and instead immediately returns whatever is
	// present in the page.  When the list of elements changes dynamically, [Locator.All] will produce unpredictable and
	// flaky results.  When the list of elements is stable, but loaded dynamically, wait for the full list to finish
	// loading before calling [Locator.All].
	All() ([]Locator, error)

	// Returns an array of `node.innerText` values for all matching nodes.
	// **NOTE** If you need to assert text on the page, prefer [LocatorAssertions.ToHaveText] with “useInnerText” option
	// to avoid flakiness. See [assertions guide] for more details.
	//
	// [assertions guide]: https://playwright.dev/docs/test-assertions
	AllInnerTexts() ([]string, error)

	// Returns an array of `node.textContent` values for all matching nodes.
	// **NOTE** If you need to assert text on the page, prefer [LocatorAssertions.ToHaveText] to avoid flakiness. See
	// [assertions guide] for more details.
	//
	// [assertions guide]: https://playwright.dev/docs/test-assertions
	AllTextContents() ([]string, error)

	// Creates a locator that matches both this locator and the argument locator.
	//
	//  locator: Additional locator to match.
	And(locator Locator) Locator

	// Calls [blur] on the element.
	//
	// [blur]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/blur
	Blur(options ...LocatorBlurOptions) error

	// This method returns the bounding box of the element matching the locator, or `null` if the element is not visible.
	// The bounding box is calculated relative to the main frame viewport - which is usually the same as the browser
	// window.
	//
	// # Details
	//
	// Scrolling affects the returned bounding box, similarly to
	// [Element.GetBoundingClientRect].
	// That means `x` and/or `y` may be negative.
	// Elements from child frames return the bounding box relative to the main frame, unlike the
	// [Element.GetBoundingClientRect].
	// Assuming the page is static, it is safe to use bounding box coordinates to perform input. For example, the
	// following snippet should click the center of the element.
	//
	// [Element.GetBoundingClientRect]: https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
	// [Element.GetBoundingClientRect]: https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
	BoundingBox(options ...LocatorBoundingBoxOptions) (*Rect, error)

	// Ensure that checkbox or radio element is checked.
	//
	// # Details
	//
	// Performs the following steps:
	//  1. Ensure that element is a checkbox or a radio input. If not, this method throws. If the element is already
	//    checked, this method returns immediately.
	//  2. Wait for [actionability] checks on the element, unless “force” option is set.
	//  3. Scroll the element into view if needed.
	//  4. Use [Page.Mouse] to click in the center of the element.
	//  5. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	//  6. Ensure that the element is now checked. If not, this method throws.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	Check(options ...LocatorCheckOptions) error

	// Clear the input field.
	//
	// # Details
	//
	// This method waits for [actionability] checks, focuses the element, clears it and triggers an
	// `input` event after clearing.
	// If the target element is not an `<input>`, `<textarea>` or `[contenteditable]` element, this method throws an
	// error. However, if the element is inside the `<label>` element that has an associated
	// [control], the control will be cleared
	// instead.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	Clear(options ...LocatorClearOptions) error

	// Click an element.
	//
	// # Details
	//
	// This method clicks the element by performing the following steps:
	//  1. Wait for [actionability] checks on the element, unless “force” option is set.
	//  2. Scroll the element into view if needed.
	//  3. Use [Page.Mouse] to click in the center of the element, or the specified “position”.
	//  4. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	Click(options ...LocatorClickOptions) error

	// Returns the number of elements matching the locator.
	// **NOTE** If you need to assert the number of elements on the page, prefer [LocatorAssertions.ToHaveCount] to avoid
	// flakiness. See [assertions guide] for more details.
	//
	// [assertions guide]: https://playwright.dev/docs/test-assertions
	Count() (int, error)

	// Double-click an element.
	//
	// # Details
	//
	// This method double clicks the element by performing the following steps:
	//  1. Wait for [actionability] checks on the element, unless “force” option is set.
	//  2. Scroll the element into view if needed.
	//  3. Use [Page.Mouse] to double click in the center of the element, or the specified “position”.
	//  4. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set. Note that if
	//    the first click of the `dblclick()` triggers a navigation event, this method will throw.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	// **NOTE** `element.dblclick()` dispatches two `click` events and a single `dblclick` event.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	Dblclick(options ...LocatorDblclickOptions) error

	// Programmatically dispatch an event on the matching element.
	//
	// # Details
	//
	// The snippet above dispatches the `click` event on the element. Regardless of the visibility state of the element,
	// `click` is dispatched. This is equivalent to calling
	// [element.Click()].
	// Under the hood, it creates an instance of an event based on the given “type”, initializes it with “eventInit”
	// properties and dispatches it on the element. Events are `composed`, `cancelable` and bubble by default.
	// Since “eventInit” is event-specific, please refer to the events documentation for the lists of initial properties:
	//  - [DeviceMotionEvent]
	//  - [DeviceOrientationEvent]
	//  - [DragEvent]
	//  - [Event]
	//  - [FocusEvent]
	//  - [KeyboardEvent]
	//  - [MouseEvent]
	//  - [PointerEvent]
	//  - [TouchEvent]
	//  - [WheelEvent]
	// You can also specify [JSHandle] as the property value if you want live objects to be passed into the event:
	//
	// 1. typ: DOM event type: `"click"`, `"dragstart"`, etc.
	// 2. eventInit: Optional event-specific initialization properties.
	//
	// [element.Click()]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click
	// [DeviceMotionEvent]: https://developer.mozilla.org/en-US/docs/Web/API/DeviceMotionEvent/DeviceMotionEvent
	// [DeviceOrientationEvent]: https://developer.mozilla.org/en-US/docs/Web/API/DeviceOrientationEvent/DeviceOrientationEvent
	// [DragEvent]: https://developer.mozilla.org/en-US/docs/Web/API/DragEvent/DragEvent
	// [Event]: https://developer.mozilla.org/en-US/docs/Web/API/Event/Event
	// [FocusEvent]: https://developer.mozilla.org/en-US/docs/Web/API/FocusEvent/FocusEvent
	// [KeyboardEvent]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/KeyboardEvent
	// [MouseEvent]: https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/MouseEvent
	// [PointerEvent]: https://developer.mozilla.org/en-US/docs/Web/API/PointerEvent/PointerEvent
	// [TouchEvent]: https://developer.mozilla.org/en-US/docs/Web/API/TouchEvent/TouchEvent
	// [WheelEvent]: https://developer.mozilla.org/en-US/docs/Web/API/WheelEvent/WheelEvent
	DispatchEvent(typ string, eventInit interface{}, options ...LocatorDispatchEventOptions) error

	// Drag the source element towards the target element and drop it.
	//
	// # Details
	//
	// This method drags the locator to another target locator or target position. It will first move to the source
	// element, perform a `mousedown`, then move to the target element or position and perform a `mouseup`.
	//
	//  target: Locator of the element to drag to.
	DragTo(target Locator, options ...LocatorDragToOptions) error

	// Resolves given locator to the first matching DOM element. If there are no matching elements, waits for one. If
	// multiple elements match the locator, throws.
	//
	// Deprecated: Always prefer using [Locator]s and web assertions over [ElementHandle]s because latter are inherently racy.
	ElementHandle(options ...LocatorElementHandleOptions) (ElementHandle, error)

	// Resolves given locator to all matching DOM elements. If there are no matching elements, returns an empty list.
	//
	// Deprecated: Always prefer using [Locator]s and web assertions over [ElementHandle]s because latter are inherently racy.
	ElementHandles() ([]ElementHandle, error)

	// Execute JavaScript code in the page, taking the matching element as an argument.
	//
	// # Details
	//
	// Returns the return value of “expression”, called with the matching element as a first argument, and “arg” as a
	// second argument.
	// If “expression” returns a [Promise], this method will wait for the promise to resolve and return its value.
	// If “expression” throws or rejects, this method throws.
	//
	// 1. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 2. arg: Optional argument to pass to “expression”.
	Evaluate(expression string, arg interface{}, options ...LocatorEvaluateOptions) (interface{}, error)

	// Execute JavaScript code in the page, taking all matching elements as an argument.
	//
	// # Details
	//
	// Returns the return value of “expression”, called with an array of all matching elements as a first argument, and
	// “arg” as a second argument.
	// If “expression” returns a [Promise], this method will wait for the promise to resolve and return its value.
	// If “expression” throws or rejects, this method throws.
	//
	// 1. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 2. arg: Optional argument to pass to “expression”.
	EvaluateAll(expression string, arg ...interface{}) (interface{}, error)

	// Execute JavaScript code in the page, taking the matching element as an argument, and return a [JSHandle] with the
	// result.
	//
	// # Details
	//
	// Returns the return value of “expression” as a[JSHandle], called with the matching element as a first argument, and
	// “arg” as a second argument.
	// The only difference between [Locator.Evaluate] and [Locator.EvaluateHandle] is that [Locator.EvaluateHandle]
	// returns [JSHandle].
	// If “expression” returns a [Promise], this method will wait for the promise to resolve and return its value.
	// If “expression” throws or rejects, this method throws.
	// See [Page.EvaluateHandle] for more details.
	//
	// 1. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 2. arg: Optional argument to pass to “expression”.
	EvaluateHandle(expression string, arg interface{}, options ...LocatorEvaluateHandleOptions) (JSHandle, error)

	// Set a value to the input field.
	//
	// # Details
	//
	// This method waits for [actionability] checks, focuses the element, fills it and triggers an
	// `input` event after filling. Note that you can pass an empty string to clear the input field.
	// If the target element is not an `<input>`, `<textarea>` or `[contenteditable]` element, this method throws an
	// error. However, if the element is inside the `<label>` element that has an associated
	// [control], the control will be filled
	// instead.
	// To send fine-grained keyboard events, use [Locator.PressSequentially].
	//
	//  value: Value to set for the `<input>`, `<textarea>` or `[contenteditable]` element.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	Fill(value string, options ...LocatorFillOptions) error

	// This method narrows existing locator according to the options, for example filters by text. It can be chained to
	// filter multiple times.
	Filter(options ...LocatorFilterOptions) Locator

	// Returns locator to the first matching element.
	First() Locator

	// Calls [focus] on the matching element.
	//
	// [focus]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus
	Focus(options ...LocatorFocusOptions) error

	// When working with iframes, you can create a frame locator that will enter the iframe and allow locating elements in
	// that iframe:
	//
	//  selector: A selector to use when resolving DOM element.
	FrameLocator(selector string) FrameLocator

	// Returns the matching element's attribute value.
	// **NOTE** If you need to assert an element's attribute, prefer [LocatorAssertions.ToHaveAttribute] to avoid
	// flakiness. See [assertions guide] for more details.
	//
	//  name: Attribute name to get the value for.
	//
	// [assertions guide]: https://playwright.dev/docs/test-assertions
	GetAttribute(name string, options ...LocatorGetAttributeOptions) (string, error)

	// Allows locating elements by their alt text.
	//
	//  text: Text to locate the element for.
	GetByAltText(text interface{}, options ...LocatorGetByAltTextOptions) Locator

	// Allows locating input elements by the text of the associated `<label>` or `aria-labelledby` element, or by the
	// `aria-label` attribute.
	//
	//  text: Text to locate the element for.
	GetByLabel(text interface{}, options ...LocatorGetByLabelOptions) Locator

	// Allows locating input elements by the placeholder text.
	//
	//  text: Text to locate the element for.
	GetByPlaceholder(text interface{}, options ...LocatorGetByPlaceholderOptions) Locator

	// Allows locating elements by their [ARIA role],
	// [ARIA attributes] and
	// [accessible name].
	//
	// # Details
	//
	// Role selector **does not replace** accessibility audits and conformance tests, but rather gives early feedback
	// about the ARIA guidelines.
	// Many html elements have an implicitly [defined role]
	// that is recognized by the role selector. You can find all the
	// [supported roles here]. ARIA guidelines **do not recommend**
	// duplicating implicit roles and attributes by setting `role` and/or `aria-*` attributes to default values.
	//
	//  role: Required aria role.
	//
	// [ARIA role]: https://www.w3.org/TR/wai-aria-1.2/#roles
	// [ARIA attributes]: https://www.w3.org/TR/wai-aria-1.2/#aria-attributes
	// [accessible name]: https://w3c.github.io/accname/#dfn-accessible-name
	// [defined role]: https://w3c.github.io/html-aam/#html-element-role-mappings
	// [supported roles here]: https://www.w3.org/TR/wai-aria-1.2/#role_definitions
	GetByRole(role AriaRole, options ...LocatorGetByRoleOptions) Locator

	// Locate element by the test id.
	//
	// # Details
	//
	// By default, the `data-testid` attribute is used as a test id. Use [Selectors.SetTestIdAttribute] to configure a
	// different test id attribute if necessary.
	//
	//  testId: Id to locate the element by.
	GetByTestId(testId interface{}) Locator

	// Allows locating elements that contain given text.
	// See also [Locator.Filter] that allows to match by another criteria, like an accessible role, and then filter by the
	// text content.
	//
	// # Details
	//
	// Matching by text always normalizes whitespace, even with exact match. For example, it turns multiple spaces into
	// one, turns line breaks into spaces and ignores leading and trailing whitespace.
	// Input elements of the type `button` and `submit` are matched by their `value` instead of the text content. For
	// example, locating by text `"Log in"` matches `<input type=button value="Log in">`.
	//
	//  text: Text to locate the element for.
	GetByText(text interface{}, options ...LocatorGetByTextOptions) Locator

	// Allows locating elements by their title attribute.
	//
	//  text: Text to locate the element for.
	GetByTitle(text interface{}, options ...LocatorGetByTitleOptions) Locator

	// Highlight the corresponding element(s) on the screen. Useful for debugging, don't commit the code that uses
	// [Locator.Highlight].
	Highlight() error

	// Hover over the matching element.
	//
	// # Details
	//
	// This method hovers over the element by performing the following steps:
	//  1. Wait for [actionability] checks on the element, unless “force” option is set.
	//  2. Scroll the element into view if needed.
	//  3. Use [Page.Mouse] to hover over the center of the element, or the specified “position”.
	//  4. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	Hover(options ...LocatorHoverOptions) error

	// Returns the [`element.innerHTML`].
	//
	// [`element.innerHTML`]: https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
	InnerHTML(options ...LocatorInnerHTMLOptions) (string, error)

	// Returns the [`element.innerText`].
	// **NOTE** If you need to assert text on the page, prefer [LocatorAssertions.ToHaveText] with “useInnerText” option
	// to avoid flakiness. See [assertions guide] for more details.
	//
	// [`element.innerText`]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
	// [assertions guide]: https://playwright.dev/docs/test-assertions
	InnerText(options ...LocatorInnerTextOptions) (string, error)

	// Returns the value for the matching `<input>` or `<textarea>` or `<select>` element.
	// **NOTE** If you need to assert input value, prefer [LocatorAssertions.ToHaveValue] to avoid flakiness. See
	// [assertions guide] for more details.
	//
	// # Details
	//
	// Throws elements that are not an input, textarea or a select. However, if the element is inside the `<label>`
	// element that has an associated
	// [control], returns the value of the
	// control.
	//
	// [assertions guide]: https://playwright.dev/docs/test-assertions
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	InputValue(options ...LocatorInputValueOptions) (string, error)

	// Returns whether the element is checked. Throws if the element is not a checkbox or radio input.
	// **NOTE** If you need to assert that checkbox is checked, prefer [LocatorAssertions.ToBeChecked] to avoid flakiness.
	// See [assertions guide] for more details.
	//
	// [assertions guide]: https://playwright.dev/docs/test-assertions
	IsChecked(options ...LocatorIsCheckedOptions) (bool, error)

	// Returns whether the element is disabled, the opposite of [enabled].
	// **NOTE** If you need to assert that an element is disabled, prefer [LocatorAssertions.ToBeDisabled] to avoid
	// flakiness. See [assertions guide] for more details.
	//
	// [enabled]: https://playwright.dev/docs/actionability#enabled
	// [assertions guide]: https://playwright.dev/docs/test-assertions
	IsDisabled(options ...LocatorIsDisabledOptions) (bool, error)

	// Returns whether the element is [editable].
	// **NOTE** If you need to assert that an element is editable, prefer [LocatorAssertions.ToBeEditable] to avoid
	// flakiness. See [assertions guide] for more details.
	//
	// [editable]: https://playwright.dev/docs/actionability#editable
	// [assertions guide]: https://playwright.dev/docs/test-assertions
	IsEditable(options ...LocatorIsEditableOptions) (bool, error)

	// Returns whether the element is [enabled].
	// **NOTE** If you need to assert that an element is enabled, prefer [LocatorAssertions.ToBeEnabled] to avoid
	// flakiness. See [assertions guide] for more details.
	//
	// [enabled]: https://playwright.dev/docs/actionability#enabled
	// [assertions guide]: https://playwright.dev/docs/test-assertions
	IsEnabled(options ...LocatorIsEnabledOptions) (bool, error)

	// Returns whether the element is hidden, the opposite of [visible].
	// **NOTE** If you need to assert that element is hidden, prefer [LocatorAssertions.ToBeHidden] to avoid flakiness.
	// See [assertions guide] for more details.
	//
	// [visible]: https://playwright.dev/docs/actionability#visible
	// [assertions guide]: https://playwright.dev/docs/test-assertions
	IsHidden(options ...LocatorIsHiddenOptions) (bool, error)

	// Returns whether the element is [visible].
	// **NOTE** If you need to assert that element is visible, prefer [LocatorAssertions.ToBeVisible] to avoid flakiness.
	// See [assertions guide] for more details.
	//
	// [visible]: https://playwright.dev/docs/actionability#visible
	// [assertions guide]: https://playwright.dev/docs/test-assertions
	IsVisible(options ...LocatorIsVisibleOptions) (bool, error)

	// Returns locator to the last matching element.
	Last() Locator

	// The method finds an element matching the specified selector in the locator's subtree. It also accepts filter
	// options, similar to [Locator.Filter] method.
	// [Learn more about locators].
	//
	//  selectorOrLocator: A selector or locator to use when resolving DOM element.
	//
	// [Learn more about locators]: https://playwright.dev/docs/locators
	Locator(selectorOrLocator interface{}, options ...LocatorLocatorOptions) Locator

	// Returns locator to the n-th matching element. It's zero based, `nth(0)` selects the first element.
	Nth(index int) Locator

	// Creates a locator that matches either of the two locators.
	//
	//  locator: Alternative locator to match.
	Or(locator Locator) Locator

	// A page this locator belongs to.
	Page() (Page, error)

	// Focuses the matching element and presses a combination of the keys.
	//
	// # Details
	//
	// Focuses the element, and then uses [Keyboard.Down] and [Keyboard.Up].
	// “key” can specify the intended
	// [keyboardEvent.Key] value or a single character
	// to generate the text for. A superset of the “key” values can be found
	// [here]. Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`,
	// etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the “key” in the upper case.
	// If “key” is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When specified with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	//
	//  key: Name of the key to press or a character to generate, such as `ArrowLeft` or `a`.
	//
	// [keyboardEvent.Key]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key
	// [here]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values
	Press(key string, options ...LocatorPressOptions) error

	// **NOTE** In most cases, you should use [Locator.Fill] instead. You only need to press keys one by one if there is
	// special keyboard handling on the page.
	// Focuses the element, and then sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the
	// text.
	// To press a special key, like `Control` or `ArrowDown`, use [Locator.Press].
	//
	//  text: String of characters to sequentially press into a focused element.
	PressSequentially(text string, options ...LocatorPressSequentiallyOptions) error

	// Take a screenshot of the element matching the locator.
	//
	// # Details
	//
	// This method captures a screenshot of the page, clipped to the size and position of a particular element matching
	// the locator. If the element is covered by other elements, it will not be actually visible on the screenshot. If the
	// element is a scrollable container, only the currently scrolled content will be visible on the screenshot.
	// This method waits for the [actionability] checks, then scrolls element into view before taking
	// a screenshot. If the element is detached from DOM, the method throws an error.
	// Returns the buffer with the captured screenshot.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	Screenshot(options ...LocatorScreenshotOptions) ([]byte, error)

	// This method waits for [actionability] checks, then tries to scroll element into view, unless
	// it is completely visible as defined by
	// [IntersectionObserver]'s `ratio`.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [IntersectionObserver]: https://developer.mozilla.org/en-US/docs/Web/API/Intersection_Observer_API
	ScrollIntoViewIfNeeded(options ...LocatorScrollIntoViewIfNeededOptions) error

	// Selects option or options in `<select>`.
	//
	// # Details
	//
	// This method waits for [actionability] checks, waits until all specified options are present in
	// the `<select>` element and selects these options.
	// If the target element is not a `<select>` element, this method throws an error. However, if the element is inside
	// the `<label>` element that has an associated
	// [control], the control will be used
	// instead.
	// Returns the array of option values that have been successfully selected.
	// Triggers a `change` and `input` event once all the provided options have been selected.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	SelectOption(values SelectOptionValues, options ...LocatorSelectOptionOptions) ([]string, error)

	// This method waits for [actionability] checks, then focuses the element and selects all its
	// text content.
	// If the element is inside the `<label>` element that has an associated
	// [control], focuses and selects text in
	// the control instead.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	SelectText(options ...LocatorSelectTextOptions) error

	// Set the state of a checkbox or a radio element.
	//
	// # Details
	//
	// This method checks or unchecks an element by performing the following steps:
	//  1. Ensure that matched element is a checkbox or a radio input. If not, this method throws.
	//  2. If the element already has the right checked state, this method returns immediately.
	//  3. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  4. Scroll the element into view if needed.
	//  5. Use [Page.Mouse] to click in the center of the element.
	//  6. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	//  7. Ensure that the element is now checked or unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	//  checked: Whether to check or uncheck the checkbox.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	SetChecked(checked bool, options ...LocatorSetCheckedOptions) error

	// Upload file or multiple files into `<input type=file>`.
	//
	// # Details
	//
	// Sets the value of the file input to these file paths or files. If some of the `filePaths` are relative paths, then
	// they are resolved relative to the current working directory. For empty array, clears the selected files.
	// This method expects [Locator] to point to an
	// [input element]. However, if the element is inside
	// the `<label>` element that has an associated
	// [control], targets the control instead.
	//
	// [input element]: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	SetInputFiles(files interface{}, options ...LocatorSetInputFilesOptions) error

	// Perform a tap gesture on the element matching the locator.
	//
	// # Details
	//
	// This method taps the element by performing the following steps:
	//  1. Wait for [actionability] checks on the element, unless “force” option is set.
	//  2. Scroll the element into view if needed.
	//  3. Use [Page.Touchscreen] to tap the center of the element, or the specified “position”.
	//  4. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	// **NOTE** `element.tap()` requires that the `hasTouch` option of the browser context be set to true.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	Tap(options ...LocatorTapOptions) error

	// Returns the [`node.textContent`].
	// **NOTE** If you need to assert text on the page, prefer [LocatorAssertions.ToHaveText] to avoid flakiness. See
	// [assertions guide] for more details.
	//
	// [`node.textContent`]: https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
	// [assertions guide]: https://playwright.dev/docs/test-assertions
	TextContent(options ...LocatorTextContentOptions) (string, error)

	// Focuses the element, and then sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the
	// text.
	// To press a special key, like `Control` or `ArrowDown`, use [Locator.Press].
	//
	// Deprecated: In most cases, you should use [Locator.Fill] instead. You only need to press keys one by one if there is special keyboard handling on the page - in this case use [Locator.PressSequentially].
	//
	//  text: A text to type into a focused element.
	Type(text string, options ...LocatorTypeOptions) error

	// Ensure that checkbox or radio element is unchecked.
	//
	// # Details
	//
	// This method unchecks the element by performing the following steps:
	//  1. Ensure that element is a checkbox or a radio input. If not, this method throws. If the element is already
	//    unchecked, this method returns immediately.
	//  2. Wait for [actionability] checks on the element, unless “force” option is set.
	//  3. Scroll the element into view if needed.
	//  4. Use [Page.Mouse] to click in the center of the element.
	//  5. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	//  6. Ensure that the element is now unchecked. If not, this method throws.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	Uncheck(options ...LocatorUncheckOptions) error

	// Returns when element specified by locator satisfies the “state” option.
	// If target element already satisfies the condition, the method returns immediately. Otherwise, waits for up to
	// “timeout” milliseconds until the condition is met.
	WaitFor(options ...LocatorWaitForOptions) error

	Err() error
}

// The [LocatorAssertions] class provides assertion methods that can be used to make assertions about the [Locator]
// state in the tests.
type LocatorAssertions interface {
	// Makes the assertion check for the opposite condition. For example, this code tests that the Locator doesn't contain
	// text `"error"`:
	Not() LocatorAssertions

	// Ensures that [Locator] points to an element that is
	// [connected] to a Document or a ShadowRoot.
	//
	// [connected]: https://developer.mozilla.org/en-US/docs/Web/API/Node/isConnected
	ToBeAttached(options ...LocatorAssertionsToBeAttachedOptions) error

	// Ensures the [Locator] points to a checked input.
	ToBeChecked(options ...LocatorAssertionsToBeCheckedOptions) error

	// Ensures the [Locator] points to a disabled element. Element is disabled if it has "disabled" attribute or is
	// disabled via
	// ['aria-disabled']. Note
	// that only native control elements such as HTML `button`, `input`, `select`, `textarea`, `option`, `optgroup` can be
	// disabled by setting "disabled" attribute. "disabled" attribute on other elements is ignored by the browser.
	//
	// ['aria-disabled']: https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-disabled
	ToBeDisabled(options ...LocatorAssertionsToBeDisabledOptions) error

	// Ensures the [Locator] points to an editable element.
	ToBeEditable(options ...LocatorAssertionsToBeEditableOptions) error

	// Ensures the [Locator] points to an empty editable element or to a DOM node that has no text.
	ToBeEmpty(options ...LocatorAssertionsToBeEmptyOptions) error

	// Ensures the [Locator] points to an enabled element.
	ToBeEnabled(options ...LocatorAssertionsToBeEnabledOptions) error

	// Ensures the [Locator] points to a focused DOM node.
	ToBeFocused(options ...LocatorAssertionsToBeFocusedOptions) error

	// Ensures that [Locator] either does not resolve to any DOM node, or resolves to a
	// [non-visible] one.
	//
	// [non-visible]: https://playwright.dev/docs/actionability#visible
	ToBeHidden(options ...LocatorAssertionsToBeHiddenOptions) error

	// Ensures the [Locator] points to an element that intersects viewport, according to the
	// [intersection observer API].
	//
	// [intersection observer API]: https://developer.mozilla.org/en-US/docs/Web/API/Intersection_Observer_API
	ToBeInViewport(options ...LocatorAssertionsToBeInViewportOptions) error

	// Ensures that [Locator] points to an attached and [visible] DOM node.
	// To check that at least one element from the list is visible, use [Locator.First].
	//
	// [visible]: https://playwright.dev/docs/actionability#visible
	ToBeVisible(options ...LocatorAssertionsToBeVisibleOptions) error

	// Ensures the [Locator] points to an element that contains the given text. All nested elements will be considered
	// when computing the text content of the element. You can use regular expressions for the value as well.
	//
	// # Details
	//
	// When `expected` parameter is a string, Playwright will normalize whitespaces and line breaks both in the actual
	// text and in the expected string before matching. When regular expression is used, the actual text is matched as is.
	//
	//  expected: Expected substring or RegExp or a list of those.
	ToContainText(expected interface{}, options ...LocatorAssertionsToContainTextOptions) error

	// Ensures the [Locator] points to an element with given attribute.
	//
	// 1. name: Attribute name.
	// 2. value: Expected attribute value.
	ToHaveAttribute(name string, value interface{}, options ...LocatorAssertionsToHaveAttributeOptions) error

	// Ensures the [Locator] points to an element with given CSS classes. This needs to be a full match or using a relaxed
	// regular expression.
	//
	//  expected: Expected class or RegExp or a list of those.
	ToHaveClass(expected interface{}, options ...LocatorAssertionsToHaveClassOptions) error

	// Ensures the [Locator] resolves to an exact number of DOM nodes.
	//
	//  count: Expected count.
	ToHaveCount(count int, options ...LocatorAssertionsToHaveCountOptions) error

	// Ensures the [Locator] resolves to an element with the given computed CSS style.
	//
	// 1. name: CSS property name.
	// 2. value: CSS property value.
	ToHaveCSS(name string, value interface{}, options ...LocatorAssertionsToHaveCSSOptions) error

	// Ensures the [Locator] points to an element with the given DOM Node ID.
	//
	//  id: Element id.
	ToHaveId(id interface{}, options ...LocatorAssertionsToHaveIdOptions) error

	// Ensures the [Locator] points to an element with given JavaScript property. Note that this property can be of a
	// primitive type as well as a plain serializable JavaScript object.
	//
	// 1. name: Property name.
	// 2. value: Property value.
	ToHaveJSProperty(name string, value interface{}, options ...LocatorAssertionsToHaveJSPropertyOptions) error

	// Ensures the [Locator] points to an element with the given text. All nested elements will be considered when
	// computing the text content of the element. You can use regular expressions for the value as well.
	//
	// # Details
	//
	// When `expected` parameter is a string, Playwright will normalize whitespaces and line breaks both in the actual
	// text and in the expected string before matching. When regular expression is used, the actual text is matched as is.
	//
	//  expected: Expected string or RegExp or a list of those.
	ToHaveText(expected interface{}, options ...LocatorAssertionsToHaveTextOptions) error

	// Ensures the [Locator] points to an element with the given input value. You can use regular expressions for the
	// value as well.
	//
	//  value: Expected value.
	ToHaveValue(value interface{}, options ...LocatorAssertionsToHaveValueOptions) error

	// Ensures the [Locator] points to multi-select/combobox (i.e. a `select` with the `multiple` attribute) and the
	// specified values are selected.
	//
	//  values: Expected options currently selected.
	ToHaveValues(values []interface{}, options ...LocatorAssertionsToHaveValuesOptions) error
}

// The Mouse class operates in main-frame CSS pixels relative to the top-left corner of the viewport.
// Every `page` object has its own Mouse, accessible with [Page.Mouse].
type Mouse interface {
	// Shortcut for [Mouse.Move], [Mouse.Down], [Mouse.Up].
	Click(x float64, y float64, options ...MouseClickOptions) error

	// Shortcut for [Mouse.Move], [Mouse.Down], [Mouse.Up], [Mouse.Down] and [Mouse.Up].
	Dblclick(x float64, y float64, options ...MouseDblclickOptions) error

	// Dispatches a `mousedown` event.
	Down(options ...MouseDownOptions) error

	// Dispatches a `mousemove` event.
	Move(x float64, y float64, options ...MouseMoveOptions) error

	// Dispatches a `mouseup` event.
	Up(options ...MouseUpOptions) error

	// Dispatches a `wheel` event.
	// **NOTE** Wheel events may cause scrolling if they are not handled, and this method does not wait for the scrolling
	// to finish before returning.
	//
	// 1. deltaX: Pixels to scroll horizontally.
	// 2. deltaY: Pixels to scroll vertically.
	Wheel(deltaX float64, deltaY float64) error
}

//	Page provides methods to interact with a single tab in a [Browser], or an
//
// [extension background page] in Chromium. One [Browser]
// instance might have multiple [Page] instances.
// This example creates a page, navigates it to a URL, and then saves a screenshot:
// The Page class emits various events (described below) which can be handled using any of Node's native
// [`EventEmitter`] methods, such as `on`, `once` or
// `removeListener`.
// This example logs a message for a single page `load` event:
// To unsubscribe from events use the `removeListener` method:
//
// [extension background page]: https://developer.chrome.com/extensions/background_pages
// [`EventEmitter`]: https://nodejs.org/api/events.html#events_class_eventemitter
type Page interface {
	EventEmitter
	// Emitted when the page closes.
	OnClose(fn func(Page))

	// Emitted when JavaScript within the page calls one of console API methods, e.g. `console.log` or `console.dir`. Also
	// emitted if the page throws an error or a warning.
	// The arguments passed into `console.log` are available on the [ConsoleMessage] event handler argument.
	OnConsole(fn func(ConsoleMessage))

	// Emitted when the page crashes. Browser pages might crash if they try to allocate too much memory. When the page
	// crashes, ongoing and subsequent operations will throw.
	// The most common way to deal with crashes is to catch an exception:
	OnCrash(fn func(Page))

	// Emitted when a JavaScript dialog appears, such as `alert`, `prompt`, `confirm` or `beforeunload`. Listener **must**
	// either [Dialog.Accept] or [Dialog.Dismiss] the dialog - otherwise the page will
	// [freeze] waiting for the dialog,
	// and actions like click will never finish.
	//
	// [freeze]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/EventLoop#never_blocking
	OnDialog(fn func(Dialog))

	// Emitted when the JavaScript
	// [`DOMContentLoaded`] event is dispatched.
	//
	// [`DOMContentLoaded`]: https://developer.mozilla.org/en-US/docs/Web/Events/DOMContentLoaded
	OnDOMContentLoaded(fn func(Page))

	// Emitted when attachment download started. User can access basic file operations on downloaded content via the
	// passed [Download] instance.
	OnDownload(fn func(Download))

	// Emitted when a file chooser is supposed to appear, such as after clicking the  `<input type=file>`. Playwright can
	// respond to it via setting the input files using [FileChooser.SetFiles] that can be uploaded after that.
	OnFileChooser(fn func(FileChooser))

	// Emitted when a frame is attached.
	OnFrameAttached(fn func(Frame))

	// Emitted when a frame is detached.
	OnFrameDetached(fn func(Frame))

	// Emitted when a frame is navigated to a new url.
	OnFrameNavigated(fn func(Frame))

	// Emitted when the JavaScript [`load`] event is dispatched.
	//
	// [`load`]: https://developer.mozilla.org/en-US/docs/Web/Events/load
	OnLoad(fn func(Page))

	// Emitted when an uncaught exception happens within the page.
	OnPageError(fn func(error))

	// Emitted when the page opens a new tab or window. This event is emitted in addition to the [BrowserContext.OnPage],
	// but only for popups relevant to this page.
	// The earliest moment that page is available is when it has navigated to the initial url. For example, when opening a
	// popup with `window.open('http://example.com')`, this event will fire when the network request to
	// "http://example.com" is done and its response has started loading in the popup.
	// **NOTE** Use [Page.WaitForLoadState] to wait until the page gets to a particular state (you should not need it in
	// most cases).
	OnPopup(fn func(Page))

	// Emitted when a page issues a request. The [request] object is read-only. In order to intercept and mutate requests,
	// see [Page.Route] or [BrowserContext.Route].
	OnRequest(fn func(Request))

	// Emitted when a request fails, for example by timing out.
	// **NOTE** HTTP Error responses, such as 404 or 503, are still successful responses from HTTP standpoint, so request
	// will complete with [Page.OnRequestFinished] event and not with [Page.OnRequestFailed]. A request will only be
	// considered failed when the client cannot get an HTTP response from the server, e.g. due to network error
	// net::ERR_FAILED.
	OnRequestFailed(fn func(Request))

	// Emitted when a request finishes successfully after downloading the response body. For a successful response, the
	// sequence of events is `request`, `response` and `requestfinished`.
	OnRequestFinished(fn func(Request))

	// Emitted when [response] status and headers are received for a request. For a successful response, the sequence of
	// events is `request`, `response` and `requestfinished`.
	OnResponse(fn func(Response))

	// Emitted when [WebSocket] request is sent.
	OnWebSocket(fn func(WebSocket))

	// Emitted when a dedicated [WebWorker] is spawned
	// by the page.
	//
	// [WebWorker]: https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API
	OnWorker(fn func(Worker))

	// Adds a script which would be evaluated in one of the following scenarios:
	//  - Whenever the page is navigated.
	//  - Whenever the child frame is attached or navigated. In this case, the script is evaluated in the context of the
	//   newly attached frame.
	// The script is evaluated after the document was created but before any of its scripts were run. This is useful to
	// amend the JavaScript environment, e.g. to seed `Math.random`.
	//
	//  script: Script to be evaluated in the page.
	AddInitScript(script Script) error

	// Adds a `<script>` tag into the page with the desired url or content. Returns the added tag when the script's onload
	// fires or when the script content was injected into frame.
	AddScriptTag(options PageAddScriptTagOptions) (ElementHandle, error)

	// Adds a `<link rel="stylesheet">` tag into the page with the desired url or a `<style type="text/css">` tag with the
	// content. Returns the added tag when the stylesheet's onload fires or when the CSS content was injected into frame.
	AddStyleTag(options PageAddStyleTagOptions) (ElementHandle, error)

	// Brings page to front (activates tab).
	BringToFront() error

	// This method checks an element matching “selector” by performing the following steps:
	//  1. Find an element matching “selector”. If there is none, wait until a matching element is attached to the DOM.
	//  2. Ensure that matched element is a checkbox or a radio input. If not, this method throws. If the element is
	//    already checked, this method returns immediately.
	//  3. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  4. Scroll the element into view if needed.
	//  5. Use [Page.Mouse] to click in the center of the element.
	//  6. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	//  7. Ensure that the element is now checked. If not, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.Check] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Check(selector string, options ...PageCheckOptions) error

	// This method clicks an element matching “selector” by performing the following steps:
	//  1. Find an element matching “selector”. If there is none, wait until a matching element is attached to the DOM.
	//  2. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  3. Scroll the element into view if needed.
	//  4. Use [Page.Mouse] to click in the center of the element, or the specified “position”.
	//  5. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.Click] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Click(selector string, options ...PageClickOptions) error

	// If “runBeforeUnload” is `false`, does not run any unload handlers and waits for the page to be closed. If
	// “runBeforeUnload” is `true` the method will run unload handlers, but will **not** wait for the page to close.
	// By default, `page.close()` **does not** run `beforeunload` handlers.
	// **NOTE** if “runBeforeUnload” is passed as true, a `beforeunload` dialog might be summoned and should be handled
	// manually via [Page.OnDialog] event.
	Close(options ...PageCloseOptions) error

	// Gets the full HTML contents of the page, including the doctype.
	Content() (string, error)

	// Get the browser context that the page belongs to.
	Context() BrowserContext

	// This method double clicks an element matching “selector” by performing the following steps:
	//  1. Find an element matching “selector”. If there is none, wait until a matching element is attached to the DOM.
	//  2. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  3. Scroll the element into view if needed.
	//  4. Use [Page.Mouse] to double click in the center of the element, or the specified “position”.
	//  5. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set. Note that if
	//    the first click of the `dblclick()` triggers a navigation event, this method will throw.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	// **NOTE** `page.dblclick()` dispatches two `click` events and a single `dblclick` event.
	//
	// Deprecated: Use locator-based [Locator.Dblclick] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Dblclick(selector string, options ...PageDblclickOptions) error

	// The snippet below dispatches the `click` event on the element. Regardless of the visibility state of the element,
	// `click` is dispatched. This is equivalent to calling
	// [element.Click()].
	//
	// Deprecated: Use locator-based [Locator.DispatchEvent] instead. Read more about [locators].
	//
	// 1. selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	// 2. typ: DOM event type: `"click"`, `"dragstart"`, etc.
	// 3. eventInit: Optional event-specific initialization properties.
	//
	// [element.Click()]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click
	// [DeviceMotionEvent]: https://developer.mozilla.org/en-US/docs/Web/API/DeviceMotionEvent/DeviceMotionEvent
	// [DeviceOrientationEvent]: https://developer.mozilla.org/en-US/docs/Web/API/DeviceOrientationEvent/DeviceOrientationEvent
	// [DragEvent]: https://developer.mozilla.org/en-US/docs/Web/API/DragEvent/DragEvent
	// [Event]: https://developer.mozilla.org/en-US/docs/Web/API/Event/Event
	// [FocusEvent]: https://developer.mozilla.org/en-US/docs/Web/API/FocusEvent/FocusEvent
	// [KeyboardEvent]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/KeyboardEvent
	// [MouseEvent]: https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/MouseEvent
	// [PointerEvent]: https://developer.mozilla.org/en-US/docs/Web/API/PointerEvent/PointerEvent
	// [TouchEvent]: https://developer.mozilla.org/en-US/docs/Web/API/TouchEvent/TouchEvent
	// [WheelEvent]: https://developer.mozilla.org/en-US/docs/Web/API/WheelEvent/WheelEvent
	// [locators]: https://playwright.dev/docs/locators
	DispatchEvent(selector string, typ string, eventInit interface{}, options ...PageDispatchEventOptions) error

	// This method drags the source element to the target element. It will first move to the source element, perform a
	// `mousedown`, then move to the target element and perform a `mouseup`.
	//
	// 1. source: A selector to search for an element to drag. If there are multiple elements satisfying the selector, the first will
	//    be used.
	// 2. target: A selector to search for an element to drop onto. If there are multiple elements satisfying the selector, the first
	//    will be used.
	DragAndDrop(source string, target string, options ...PageDragAndDropOptions) error

	// This method changes the `CSS media type` through the `media` argument, and/or the `prefers-colors-scheme` media
	// feature, using the `colorScheme` argument.
	EmulateMedia(options ...PageEmulateMediaOptions) error

	// The method finds an element matching the specified selector within the page and passes it as a first argument to
	// “expression”. If no elements match the selector, the method throws an error. Returns the value of “expression”.
	// If “expression” returns a [Promise], then [Page.EvalOnSelector] would wait for the promise to resolve and return
	// its value.
	//
	// Deprecated: This method does not wait for the element to pass actionability checks and therefore can lead to the flaky tests. Use [Locator.Evaluate], other [Locator] helper methods or web-first assertions instead.
	//
	// 1. selector: A selector to query for.
	// 2. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 3. arg: Optional argument to pass to “expression”.
	EvalOnSelector(selector string, expression string, arg interface{}, options ...PageEvalOnSelectorOptions) (interface{}, error)

	// The method finds all elements matching the specified selector within the page and passes an array of matched
	// elements as a first argument to “expression”. Returns the result of “expression” invocation.
	// If “expression” returns a [Promise], then [Page.EvalOnSelectorAll] would wait for the promise to resolve and return
	// its value.
	//
	// Deprecated: In most cases, [Locator.EvaluateAll], other [Locator] helper methods and web-first assertions do a better job.
	//
	// 1. selector: A selector to query for.
	// 2. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 3. arg: Optional argument to pass to “expression”.
	EvalOnSelectorAll(selector string, expression string, arg ...interface{}) (interface{}, error)

	// Returns the value of the “expression” invocation.
	// If the function passed to the [Page.Evaluate] returns a [Promise], then [Page.Evaluate] would wait for the promise
	// to resolve and return its value.
	// If the function passed to the [Page.Evaluate] returns a non-[Serializable] value, then [Page.Evaluate] resolves to
	// `undefined`. Playwright also supports transferring some additional values that are not serializable by `JSON`:
	// `-0`, `NaN`, `Infinity`, `-Infinity`.
	//
	// 1. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 2. arg: Optional argument to pass to “expression”.
	Evaluate(expression string, arg ...interface{}) (interface{}, error)

	// Returns the value of the “expression” invocation as a [JSHandle].
	// The only difference between [Page.Evaluate] and [Page.EvaluateHandle] is that [Page.EvaluateHandle] returns
	// [JSHandle].
	// If the function passed to the [Page.EvaluateHandle] returns a [Promise], then [Page.EvaluateHandle] would wait for
	// the promise to resolve and return its value.
	//
	// 1. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 2. arg: Optional argument to pass to “expression”.
	EvaluateHandle(expression string, arg ...interface{}) (JSHandle, error)

	// The method adds a function called “name” on the `window` object of every frame in this page. When called, the
	// function executes “callback” and returns a [Promise] which resolves to the return value of “callback”. If the
	// “callback” returns a [Promise], it will be awaited.
	// The first argument of the “callback” function contains information about the caller: `{ browserContext:
	// BrowserContext, page: Page, frame: Frame }`.
	// See [BrowserContext.ExposeBinding] for the context-wide version.
	// **NOTE** Functions installed via [Page.ExposeBinding] survive navigations.
	//
	// 1. name: Name of the function on the window object.
	// 2. binding: Callback function that will be called in the Playwright's context.
	ExposeBinding(name string, binding BindingCallFunction, handle ...bool) error

	// The method adds a function called “name” on the `window` object of every frame in the page. When called, the
	// function executes “callback” and returns a [Promise] which resolves to the return value of “callback”.
	// If the “callback” returns a [Promise], it will be awaited.
	// See [BrowserContext.ExposeFunction] for context-wide exposed function.
	// **NOTE** Functions installed via [Page.ExposeFunction] survive navigations.
	//
	// 1. name: Name of the function on the window object
	// 2. binding: Callback function which will be called in Playwright's context.
	ExposeFunction(name string, binding ExposedFunction) error

	// This method waits for an element matching “selector”, waits for [actionability] checks,
	// focuses the element, fills it and triggers an `input` event after filling. Note that you can pass an empty string
	// to clear the input field.
	// If the target element is not an `<input>`, `<textarea>` or `[contenteditable]` element, this method throws an
	// error. However, if the element is inside the `<label>` element that has an associated
	// [control], the control will be filled
	// instead.
	// To send fine-grained keyboard events, use [Locator.PressSequentially].
	//
	// Deprecated: Use locator-based [Locator.Fill] instead. Read more about [locators].
	//
	// 1. selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	// 2. value: Value to fill for the `<input>`, `<textarea>` or `[contenteditable]` element.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	// [locators]: https://playwright.dev/docs/locators
	Fill(selector string, value string, options ...PageFillOptions) error

	// This method fetches an element with “selector” and focuses it. If there's no element matching “selector”, the
	// method waits until a matching element appears in the DOM.
	//
	// Deprecated: Use locator-based [Locator.Focus] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [locators]: https://playwright.dev/docs/locators
	Focus(selector string, options ...PageFocusOptions) error

	// Returns frame matching the specified criteria. Either `name` or `url` must be specified.
	Frame(options ...PageFrameOptions) Frame

	// When working with iframes, you can create a frame locator that will enter the iframe and allow selecting elements
	// in that iframe.
	//
	//  selector: A selector to use when resolving DOM element.
	FrameLocator(selector string) FrameLocator

	// An array of all frames attached to the page.
	Frames() []Frame

	// Returns element attribute value.
	//
	// Deprecated: Use locator-based [Locator.GetAttribute] instead. Read more about [locators].
	//
	// 1. selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	// 2. name: Attribute name to get the value for.
	//
	// [locators]: https://playwright.dev/docs/locators
	GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error)

	// Allows locating elements by their alt text.
	//
	//  text: Text to locate the element for.
	GetByAltText(text interface{}, options ...PageGetByAltTextOptions) Locator

	// Allows locating input elements by the text of the associated `<label>` or `aria-labelledby` element, or by the
	// `aria-label` attribute.
	//
	//  text: Text to locate the element for.
	GetByLabel(text interface{}, options ...PageGetByLabelOptions) Locator

	// Allows locating input elements by the placeholder text.
	//
	//  text: Text to locate the element for.
	GetByPlaceholder(text interface{}, options ...PageGetByPlaceholderOptions) Locator

	// Allows locating elements by their [ARIA role],
	// [ARIA attributes] and
	// [accessible name].
	//
	// # Details
	//
	// Role selector **does not replace** accessibility audits and conformance tests, but rather gives early feedback
	// about the ARIA guidelines.
	// Many html elements have an implicitly [defined role]
	// that is recognized by the role selector. You can find all the
	// [supported roles here]. ARIA guidelines **do not recommend**
	// duplicating implicit roles and attributes by setting `role` and/or `aria-*` attributes to default values.
	//
	//  role: Required aria role.
	//
	// [ARIA role]: https://www.w3.org/TR/wai-aria-1.2/#roles
	// [ARIA attributes]: https://www.w3.org/TR/wai-aria-1.2/#aria-attributes
	// [accessible name]: https://w3c.github.io/accname/#dfn-accessible-name
	// [defined role]: https://w3c.github.io/html-aam/#html-element-role-mappings
	// [supported roles here]: https://www.w3.org/TR/wai-aria-1.2/#role_definitions
	GetByRole(role AriaRole, options ...PageGetByRoleOptions) Locator

	// Locate element by the test id.
	//
	// # Details
	//
	// By default, the `data-testid` attribute is used as a test id. Use [Selectors.SetTestIdAttribute] to configure a
	// different test id attribute if necessary.
	//
	//  testId: Id to locate the element by.
	GetByTestId(testId interface{}) Locator

	// Allows locating elements that contain given text.
	// See also [Locator.Filter] that allows to match by another criteria, like an accessible role, and then filter by the
	// text content.
	//
	// # Details
	//
	// Matching by text always normalizes whitespace, even with exact match. For example, it turns multiple spaces into
	// one, turns line breaks into spaces and ignores leading and trailing whitespace.
	// Input elements of the type `button` and `submit` are matched by their `value` instead of the text content. For
	// example, locating by text `"Log in"` matches `<input type=button value="Log in">`.
	//
	//  text: Text to locate the element for.
	GetByText(text interface{}, options ...PageGetByTextOptions) Locator

	// Allows locating elements by their title attribute.
	//
	//  text: Text to locate the element for.
	GetByTitle(text interface{}, options ...PageGetByTitleOptions) Locator

	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of
	// the last redirect. If can not go back, returns `null`.
	// Navigate to the previous page in history.
	GoBack(options ...PageGoBackOptions) (Response, error)

	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of
	// the last redirect. If can not go forward, returns `null`.
	// Navigate to the next page in history.
	GoForward(options ...PageGoForwardOptions) (Response, error)

	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the first
	// non-redirect response.
	// The method will throw an error if:
	//  - there's an SSL error (e.g. in case of self-signed certificates).
	//  - target URL is invalid.
	//  - the “timeout” is exceeded during navigation.
	//  - the remote server does not respond or is unreachable.
	//  - the main resource failed to load.
	// The method will not throw an error when any valid HTTP status code is returned by the remote server, including 404
	// "Not Found" and 500 "Internal Server Error".  The status code for such responses can be retrieved by calling
	// [Response.Status].
	// **NOTE** The method either throws an error or returns a main resource response. The only exceptions are navigation
	// to `about:blank` or navigation to the same URL with a different hash, which would succeed and return `null`.
	// **NOTE** Headless mode doesn't support navigation to a PDF document. See the
	// [upstream issue].
	//
	//  url: URL to navigate page to. The url should include scheme, e.g. `https://`. When a “baseURL” via the context options
	//    was provided and the passed URL is a path, it gets merged via the
	//    [`new URL()`](https://developer.mozilla.org/en-US/docs/Web/API/URL/URL) constructor.
	//
	// [upstream issue]: https://bugs.chromium.org/p/chromium/issues/detail?id=761295
	Goto(url string, options ...PageGotoOptions) (Response, error)

	// This method hovers over an element matching “selector” by performing the following steps:
	//  1. Find an element matching “selector”. If there is none, wait until a matching element is attached to the DOM.
	//  2. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  3. Scroll the element into view if needed.
	//  4. Use [Page.Mouse] to hover over the center of the element, or the specified “position”.
	//  5. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.Hover] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Hover(selector string, options ...PageHoverOptions) error

	// Returns `element.innerHTML`.
	//
	// Deprecated: Use locator-based [Locator.InnerHTML] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [locators]: https://playwright.dev/docs/locators
	InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error)

	// Returns `element.innerText`.
	//
	// Deprecated: Use locator-based [Locator.InnerText] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [locators]: https://playwright.dev/docs/locators
	InnerText(selector string, options ...PageInnerTextOptions) (string, error)

	// Returns `input.value` for the selected `<input>` or `<textarea>` or `<select>` element.
	// Throws for non-input elements. However, if the element is inside the `<label>` element that has an associated
	// [control], returns the value of the
	// control.
	//
	// Deprecated: Use locator-based [Locator.InputValue] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	// [locators]: https://playwright.dev/docs/locators
	InputValue(selector string, options ...PageInputValueOptions) (string, error)

	// Returns whether the element is checked. Throws if the element is not a checkbox or radio input.
	//
	// Deprecated: Use locator-based [Locator.IsChecked] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [locators]: https://playwright.dev/docs/locators
	IsChecked(selector string, options ...PageIsCheckedOptions) (bool, error)

	// Indicates that the page has been closed.
	IsClosed() bool

	// Returns whether the element is disabled, the opposite of [enabled].
	//
	// Deprecated: Use locator-based [Locator.IsDisabled] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [enabled]: https://playwright.dev/docs/actionability#enabled
	// [locators]: https://playwright.dev/docs/locators
	IsDisabled(selector string, options ...PageIsDisabledOptions) (bool, error)

	// Returns whether the element is [editable].
	//
	// Deprecated: Use locator-based [Locator.IsEditable] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [editable]: https://playwright.dev/docs/actionability#editable
	// [locators]: https://playwright.dev/docs/locators
	IsEditable(selector string, options ...PageIsEditableOptions) (bool, error)

	// Returns whether the element is [enabled].
	//
	// Deprecated: Use locator-based [Locator.IsEnabled] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [enabled]: https://playwright.dev/docs/actionability#enabled
	// [locators]: https://playwright.dev/docs/locators
	IsEnabled(selector string, options ...PageIsEnabledOptions) (bool, error)

	// Returns whether the element is hidden, the opposite of [visible].  “selector” that
	// does not match any elements is considered hidden.
	//
	// Deprecated: Use locator-based [Locator.IsHidden] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [visible]: https://playwright.dev/docs/actionability#visible
	// [locators]: https://playwright.dev/docs/locators
	IsHidden(selector string, options ...PageIsHiddenOptions) (bool, error)

	// Returns whether the element is [visible]. “selector” that does not match any elements
	// is considered not visible.
	//
	// Deprecated: Use locator-based [Locator.IsVisible] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [visible]: https://playwright.dev/docs/actionability#visible
	// [locators]: https://playwright.dev/docs/locators
	IsVisible(selector string, options ...PageIsVisibleOptions) (bool, error)

	Keyboard() Keyboard

	// The method returns an element locator that can be used to perform actions on this page / frame. Locator is resolved
	// to the element immediately before performing an action, so a series of actions on the same locator can in fact be
	// performed on different DOM elements. That would happen if the DOM structure between those actions has changed.
	// [Learn more about locators].
	//
	//  selector: A selector to use when resolving DOM element.
	//
	// [Learn more about locators]: https://playwright.dev/docs/locators
	Locator(selector string, options ...PageLocatorOptions) Locator

	// The page's main frame. Page is guaranteed to have a main frame which persists during navigations.
	MainFrame() Frame

	Mouse() Mouse

	// Returns the opener for popup pages and `null` for others. If the opener has been closed already the returns `null`.
	Opener() (Page, error)

	// Pauses script execution. Playwright will stop executing the script and wait for the user to either press 'Resume'
	// button in the page overlay or to call `playwright.resume()` in the DevTools console.
	// User can inspect selectors or perform manual steps while paused. Resume will continue running the original script
	// from the place it was paused.
	// **NOTE** This method requires Playwright to be started in a headed mode, with a falsy “headless” value in the
	// [BrowserType.Launch].
	Pause() error

	// Returns the PDF buffer.
	// **NOTE** Generating a pdf is currently only supported in Chromium headless.
	// `page.pdf()` generates a pdf of the page with `print` css media. To generate a pdf with `screen` media, call
	// [Page.EmulateMedia] before calling `page.pdf()`:
	// **NOTE** By default, `page.pdf()` generates a pdf with modified colors for printing. Use the
	// [`-webkit-print-color-adjust`]
	// property to force rendering of exact colors.
	//
	// [`-webkit-print-color-adjust`]: https://developer.mozilla.org/en-US/docs/Web/CSS/-webkit-print-color-adjust
	PDF(options ...PagePdfOptions) ([]byte, error)

	// Focuses the element, and then uses [Keyboard.Down] and [Keyboard.Up].
	// “key” can specify the intended
	// [keyboardEvent.Key] value or a single character
	// to generate the text for. A superset of the “key” values can be found
	// [here]. Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`,
	// etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the “key” in the upper case.
	// If “key” is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When specified with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	//
	// Deprecated: Use locator-based [Locator.Press] instead. Read more about [locators].
	//
	// 1. selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	// 2. key: Name of the key to press or a character to generate, such as `ArrowLeft` or `a`.
	//
	// [keyboardEvent.Key]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key
	// [here]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values
	// [locators]: https://playwright.dev/docs/locators
	Press(selector string, key string, options ...PagePressOptions) error

	// The method finds an element matching the specified selector within the page. If no elements match the selector, the
	// return value resolves to `null`. To wait for an element on the page, use [Locator.WaitFor].
	//
	// Deprecated: Use locator-based [Page.Locator] instead. Read more about [locators].
	//
	//  selector: A selector to query for.
	//
	// [locators]: https://playwright.dev/docs/locators
	QuerySelector(selector string, options ...PageQuerySelectorOptions) (ElementHandle, error)

	// The method finds all elements matching the specified selector within the page. If no elements match the selector,
	// the return value resolves to `[]`.
	//
	// Deprecated: Use locator-based [Page.Locator] instead. Read more about [locators].
	//
	//  selector: A selector to query for.
	//
	// [locators]: https://playwright.dev/docs/locators
	QuerySelectorAll(selector string) ([]ElementHandle, error)

	// This method reloads the current page, in the same way as if the user had triggered a browser refresh. Returns the
	// main resource response. In case of multiple redirects, the navigation will resolve with the response of the last
	// redirect.
	Reload(options ...PageReloadOptions) (Response, error)

	// API testing helper associated with this page. This method returns the same instance as [BrowserContext.Request] on
	// the page's context. See [BrowserContext.Request] for more details.
	Request() APIRequestContext

	// Routing provides the capability to modify network requests that are made by a page.
	// Once routing is enabled, every request matching the url pattern will stall unless it's continued, fulfilled or
	// aborted.
	// **NOTE** The handler will only be called for the first url if the response is a redirect.
	// **NOTE** [Page.Route] will not intercept requests intercepted by Service Worker. See
	// [this] issue. We recommend disabling Service Workers when
	// using request interception by setting “Browser.newContext.serviceWorkers” to `block`.
	//
	// 1. url: A glob pattern, regex pattern or predicate receiving [URL] to match while routing. When a “baseURL” via the context
	//    options was provided and the passed URL is a path, it gets merged via the
	//    [`new URL()`](https://developer.mozilla.org/en-US/docs/Web/API/URL/URL) constructor.
	// 2. handler: handler function to route the request.
	//
	// [this]: https://github.com/microsoft/playwright/issues/1090
	Route(url interface{}, handler routeHandler, times ...int) error

	// If specified the network requests that are made in the page will be served from the HAR file. Read more about
	// [Replaying from HAR].
	// Playwright will not serve requests intercepted by Service Worker from the HAR file. See
	// [this] issue. We recommend disabling Service Workers when
	// using request interception by setting “Browser.newContext.serviceWorkers” to `block`.
	//
	//  har: Path to a [HAR](http://www.softwareishard.com/blog/har-12-spec) file with prerecorded network data. If `path` is a
	//    relative path, then it is resolved relative to the current working directory.
	//
	// [Replaying from HAR]: https://playwright.dev/docs/mock#replaying-from-har
	// [this]: https://github.com/microsoft/playwright/issues/1090
	RouteFromHAR(har string, options ...PageRouteFromHAROptions) error

	// Returns the buffer with the captured screenshot.
	Screenshot(options ...PageScreenshotOptions) ([]byte, error)

	// This method waits for an element matching “selector”, waits for [actionability] checks, waits
	// until all specified options are present in the `<select>` element and selects these options.
	// If the target element is not a `<select>` element, this method throws an error. However, if the element is inside
	// the `<label>` element that has an associated
	// [control], the control will be used
	// instead.
	// Returns the array of option values that have been successfully selected.
	// Triggers a `change` and `input` event once all the provided options have been selected.
	//
	// Deprecated: Use locator-based [Locator.SelectOption] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	// [locators]: https://playwright.dev/docs/locators
	SelectOption(selector string, values SelectOptionValues, options ...PageSelectOptionOptions) ([]string, error)

	// This method checks or unchecks an element matching “selector” by performing the following steps:
	//  1. Find an element matching “selector”. If there is none, wait until a matching element is attached to the DOM.
	//  2. Ensure that matched element is a checkbox or a radio input. If not, this method throws.
	//  3. If the element already has the right checked state, this method returns immediately.
	//  4. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  5. Scroll the element into view if needed.
	//  6. Use [Page.Mouse] to click in the center of the element.
	//  7. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	//  8. Ensure that the element is now checked or unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.SetChecked] instead. Read more about [locators].
	//
	// 1. selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	// 2. checked: Whether to check or uncheck the checkbox.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	SetChecked(selector string, checked bool, options ...PageSetCheckedOptions) error

	// This method internally calls [document.Write()],
	// inheriting all its specific characteristics and behaviors.
	//
	//  html: HTML markup to assign to the page.
	//
	// [document.Write()]: https://developer.mozilla.org/en-US/docs/Web/API/Document/write
	SetContent(html string, options ...PageSetContentOptions) error

	// This setting will change the default maximum navigation time for the following methods and related shortcuts:
	//  - [Page.GoBack]
	//  - [Page.GoForward]
	//  - [Page.Goto]
	//  - [Page.Reload]
	//  - [Page.SetContent]
	//  - [Page.ExpectNavigation]
	//  - [Page.WaitForURL]
	// **NOTE** [Page.SetDefaultNavigationTimeout] takes priority over [Page.SetDefaultTimeout],
	// [BrowserContext.SetDefaultTimeout] and [BrowserContext.SetDefaultNavigationTimeout].
	//
	//  timeout: Maximum navigation time in milliseconds
	SetDefaultNavigationTimeout(timeout float64)

	// This setting will change the default maximum time for all the methods accepting “timeout” option.
	// **NOTE** [Page.SetDefaultNavigationTimeout] takes priority over [Page.SetDefaultTimeout].
	//
	//  timeout: Maximum time in milliseconds
	SetDefaultTimeout(timeout float64)

	// The extra HTTP headers will be sent with every request the page initiates.
	// **NOTE** [Page.SetExtraHTTPHeaders] does not guarantee the order of headers in the outgoing requests.
	//
	//  headers: An object containing additional HTTP headers to be sent with every request. All header values must be strings.
	SetExtraHTTPHeaders(headers map[string]string) error

	// Sets the value of the file input to these file paths or files. If some of the `filePaths` are relative paths, then
	// they are resolved relative to the current working directory. For empty array, clears the selected files.
	// This method expects “selector” to point to an
	// [input element]. However, if the element is inside
	// the `<label>` element that has an associated
	// [control], targets the control instead.
	//
	// Deprecated: Use locator-based [Locator.SetInputFiles] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [input element]: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
	// [control]: https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control
	// [locators]: https://playwright.dev/docs/locators
	SetInputFiles(selector string, files interface{}, options ...PageSetInputFilesOptions) error

	// In the case of multiple pages in a single browser, each page can have its own viewport size. However,
	// [Browser.NewContext] allows to set viewport size (and more) for all pages in the context at once.
	// [Page.SetViewportSize] will resize the page. A lot of websites don't expect phones to change size, so you should
	// set the viewport size before navigating to the page. [Page.SetViewportSize] will also reset `screen` size, use
	// [Browser.NewContext] with `screen` and `viewport` parameters if you need better control of these properties.
	SetViewportSize(width int, height int) error

	// This method taps an element matching “selector” by performing the following steps:
	//  1. Find an element matching “selector”. If there is none, wait until a matching element is attached to the DOM.
	//  2. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  3. Scroll the element into view if needed.
	//  4. Use [Page.Touchscreen] to tap the center of the element, or the specified “position”.
	//  5. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	// **NOTE** [Page.Tap] the method will throw if “hasTouch” option of the browser context is false.
	//
	// Deprecated: Use locator-based [Locator.Tap] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Tap(selector string, options ...PageTapOptions) error

	// Returns `element.textContent`.
	//
	// Deprecated: Use locator-based [Locator.TextContent] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [locators]: https://playwright.dev/docs/locators
	TextContent(selector string, options ...PageTextContentOptions) (string, error)

	// Returns the page's title.
	Title() (string, error)

	Touchscreen() Touchscreen

	// Sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the text. `page.type` can be used to
	// send fine-grained keyboard events. To fill values in form fields, use [Page.Fill].
	// To press a special key, like `Control` or `ArrowDown`, use [Keyboard.Press].
	//
	// Deprecated: In most cases, you should use [Locator.Fill] instead. You only need to press keys one by one if there is special keyboard handling on the page - in this case use [Locator.PressSequentially].
	//
	// 1. selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	// 2. text: A text to type into a focused element.
	Type(selector string, text string, options ...PageTypeOptions) error

	// This method unchecks an element matching “selector” by performing the following steps:
	//  1. Find an element matching “selector”. If there is none, wait until a matching element is attached to the DOM.
	//  2. Ensure that matched element is a checkbox or a radio input. If not, this method throws. If the element is
	//    already unchecked, this method returns immediately.
	//  3. Wait for [actionability] checks on the matched element, unless “force” option is set. If
	//    the element is detached during the checks, the whole action is retried.
	//  4. Scroll the element into view if needed.
	//  5. Use [Page.Mouse] to click in the center of the element.
	//  6. Wait for initiated navigations to either succeed or fail, unless “noWaitAfter” option is set.
	//  7. Ensure that the element is now unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified “timeout”, this method throws a [TimeoutError].
	// Passing zero timeout disables this.
	//
	// Deprecated: Use locator-based [Locator.Uncheck] instead. Read more about [locators].
	//
	//  selector: A selector to search for an element. If there are multiple elements satisfying the selector, the first will be
	//    used.
	//
	// [actionability]: https://playwright.dev/docs/actionability
	// [locators]: https://playwright.dev/docs/locators
	Uncheck(selector string, options ...PageUncheckOptions) error

	// Removes all routes created with [Page.Route] and [Page.RouteFromHAR].
	UnrouteAll(options ...PageUnrouteAllOptions) error

	// Removes a route created with [Page.Route]. When “handler” is not specified, removes all routes for the “url”.
	//
	// 1. url: A glob pattern, regex pattern or predicate receiving [URL] to match while routing.
	// 2. handler: Optional handler function to route the request.
	Unroute(url interface{}, handler ...routeHandler) error

	URL() string

	// Video object associated with this page.
	Video() Video

	ViewportSize() *Size

	// Performs action and waits for a [ConsoleMessage] to be logged by in the page. If predicate is provided, it passes
	// [ConsoleMessage] value into the `predicate` function and waits for `predicate(message)` to return a truthy value.
	// Will throw an error if the page is closed before the [Page.OnConsole] event is fired.
	ExpectConsoleMessage(cb func() error, options ...PageExpectConsoleMessageOptions) (ConsoleMessage, error)

	// Performs action and waits for a new [Download]. If predicate is provided, it passes [Download] value into the
	// `predicate` function and waits for `predicate(download)` to return a truthy value. Will throw an error if the page
	// is closed before the download event is fired.
	ExpectDownload(cb func() error, options ...PageExpectDownloadOptions) (Download, error)

	// Waits for event to fire and passes its value into the predicate function. Returns when the predicate returns truthy
	// value. Will throw an error if the page is closed before the event is fired. Returns the event data value.
	//
	//  event: Event name, same one typically passed into `*.on(event)`.
	ExpectEvent(event string, cb func() error, options ...PageExpectEventOptions) (interface{}, error)

	// Performs action and waits for a new [FileChooser] to be created. If predicate is provided, it passes [FileChooser]
	// value into the `predicate` function and waits for `predicate(fileChooser)` to return a truthy value. Will throw an
	// error if the page is closed before the file chooser is opened.
	ExpectFileChooser(cb func() error, options ...PageExpectFileChooserOptions) (FileChooser, error)

	// Returns when the “expression” returns a truthy value. It resolves to a JSHandle of the truthy value.
	//
	// 1. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 2. arg: Optional argument to pass to “expression”.
	WaitForFunction(expression string, arg interface{}, options ...PageWaitForFunctionOptions) (JSHandle, error)

	// Returns when the required load state has been reached.
	// This resolves when the page reaches a required load state, `load` by default. The navigation must have been
	// committed when this method is called. If current document has already reached the required state, resolves
	// immediately.
	WaitForLoadState(options ...PageWaitForLoadStateOptions) error

	// Waits for the main frame navigation and returns the main resource response. In case of multiple redirects, the
	// navigation will resolve with the response of the last redirect. In case of navigation to a different anchor or
	// navigation due to History API usage, the navigation will resolve with `null`.
	//
	// Deprecated: This method is inherently racy, please use [Page.WaitForURL] instead.
	//
	// [History API]: https://developer.mozilla.org/en-US/docs/Web/API/History_API
	ExpectNavigation(cb func() error, options ...PageExpectNavigationOptions) (Response, error)

	// Performs action and waits for a popup [Page]. If predicate is provided, it passes [Popup] value into the
	// `predicate` function and waits for `predicate(page)` to return a truthy value. Will throw an error if the page is
	// closed before the popup event is fired.
	ExpectPopup(cb func() error, options ...PageExpectPopupOptions) (Page, error)

	// Waits for the matching request and returns it. See [waiting for event] for more
	// details about events.
	//
	//  urlOrPredicate: Request URL string, regex or predicate receiving [Request] object. When a “baseURL” via the context options was
	//    provided and the passed URL is a path, it gets merged via the
	//    [`new URL()`](https://developer.mozilla.org/en-US/docs/Web/API/URL/URL) constructor.
	//
	// [waiting for event]: https://playwright.dev/docs/events#waiting-for-event
	ExpectRequest(urlOrPredicate interface{}, cb func() error, options ...PageExpectRequestOptions) (Request, error)

	// Performs action and waits for a [Request] to finish loading. If predicate is provided, it passes [Request] value
	// into the `predicate` function and waits for `predicate(request)` to return a truthy value. Will throw an error if
	// the page is closed before the [Page.OnRequestFinished] event is fired.
	ExpectRequestFinished(cb func() error, options ...PageExpectRequestFinishedOptions) (Request, error)

	// Returns the matched response. See [waiting for event] for more details about
	// events.
	//
	//  urlOrPredicate: Request URL string, regex or predicate receiving [Response] object. When a “baseURL” via the context options was
	//    provided and the passed URL is a path, it gets merged via the
	//    [`new URL()`](https://developer.mozilla.org/en-US/docs/Web/API/URL/URL) constructor.
	//
	// [waiting for event]: https://playwright.dev/docs/events#waiting-for-event
	ExpectResponse(urlOrPredicate interface{}, cb func() error, options ...PageExpectResponseOptions) (Response, error)

	// Returns when element specified by selector satisfies “state” option. Returns `null` if waiting for `hidden` or
	// `detached`.
	// **NOTE** Playwright automatically waits for element to be ready before performing an action. Using [Locator]
	// objects and web-first assertions makes the code wait-for-selector-free.
	// Wait for the “selector” to satisfy “state” option (either appear/disappear from dom, or become visible/hidden). If
	// at the moment of calling the method “selector” already satisfies the condition, the method will return immediately.
	// If the selector doesn't satisfy the condition for the “timeout” milliseconds, the function will throw.
	//
	// Deprecated: Use web assertions that assert visibility or a locator-based [Locator.WaitFor] instead. Read more about [locators].
	//
	//  selector: A selector to query for.
	//
	// [locators]: https://playwright.dev/docs/locators
	WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandle, error)

	// Waits for the given “timeout” in milliseconds.
	// Note that `page.waitForTimeout()` should only be used for debugging. Tests using the timer in production are going
	// to be flaky. Use signals such as network events, selectors becoming visible and others instead.
	//
	// Deprecated: Never wait for timeout in production. Tests that wait for time are inherently flaky. Use [Locator] actions and web assertions that wait automatically.
	//
	//  timeout: A timeout to wait for
	WaitForTimeout(timeout float64)

	// Waits for the main frame to navigate to the given URL.
	//
	//  url: A glob pattern, regex pattern or predicate receiving [URL] to match while waiting for the navigation. Note that if
	//    the parameter is a string without wildcard characters, the method will wait for navigation to URL that is exactly
	//    equal to the string.
	WaitForURL(url interface{}, options ...PageWaitForURLOptions) error

	// Performs action and waits for a new [WebSocket]. If predicate is provided, it passes [WebSocket] value into the
	// `predicate` function and waits for `predicate(webSocket)` to return a truthy value. Will throw an error if the page
	// is closed before the WebSocket event is fired.
	ExpectWebSocket(cb func() error, options ...PageExpectWebSocketOptions) (WebSocket, error)

	// Performs action and waits for a new [Worker]. If predicate is provided, it passes [Worker] value into the
	// `predicate` function and waits for `predicate(worker)` to return a truthy value. Will throw an error if the page is
	// closed before the worker event is fired.
	ExpectWorker(cb func() error, options ...PageExpectWorkerOptions) (Worker, error)

	// This method returns all of the dedicated
	// [WebWorkers] associated with the page.
	// **NOTE** This does not contain ServiceWorkers
	//
	// [WebWorkers]: https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API
	Workers() []Worker

	// **NOTE** In most cases, you should use [Page.ExpectEvent].
	// Waits for given `event` to fire. If predicate is provided, it passes event's value into the `predicate` function
	// and waits for `predicate(event)` to return a truthy value. Will throw an error if the page is closed before the
	// `event` is fired.
	//
	//  event: Event name, same one typically passed into `*.on(event)`.
	WaitForEvent(event string, options ...PageWaitForEventOptions) (interface{}, error)
}

// The [PageAssertions] class provides assertion methods that can be used to make assertions about the [Page] state in
// the tests.
type PageAssertions interface {
	// Makes the assertion check for the opposite condition. For example, this code tests that the page URL doesn't
	// contain `"error"`:
	Not() PageAssertions

	// Ensures the page has the given title.
	//
	//  titleOrRegExp: Expected title or RegExp.
	ToHaveTitle(titleOrRegExp interface{}, options ...PageAssertionsToHaveTitleOptions) error

	// Ensures the page is navigated to the given URL.
	//
	//  urlOrRegExp: Expected URL string or RegExp.
	ToHaveURL(urlOrRegExp interface{}, options ...PageAssertionsToHaveURLOptions) error
}

// Playwright gives you Web-First Assertions with convenience methods for creating assertions that will wait and retry
// until the expected condition is met.
// Consider the following example:
// Playwright will be re-testing the node with the selector `.status` until fetched Node has the `"Submitted"` text.
// It will be re-fetching the node and checking it over and over, until the condition is met or until the timeout is
// reached. You can pass this timeout as an option.
// By default, the timeout for assertions is set to 5 seconds.
type PlaywrightAssertions interface {
	// Creates a [APIResponseAssertions] object for the given [APIResponse].
	//
	//  response: [APIResponse] object to use for assertions.
	APIResponse(response APIResponse) APIResponseAssertions

	// Creates a [LocatorAssertions] object for the given [Locator].
	//
	//  locator: [Locator] object to use for assertions.
	Locator(locator Locator) LocatorAssertions

	// Creates a [PageAssertions] object for the given [Page].
	//
	//  page: [Page] object to use for assertions.
	Page(page Page) PageAssertions
}

// Whenever the page sends a request for a network resource the following sequence of events are emitted by [Page]:
//   - [Page.OnRequest] emitted when the request is issued by the page.
//   - [Page.OnResponse] emitted when/if the response status and headers are received for the request.
//   - [Page.OnRequestFinished] emitted when the response body is downloaded and the request is complete.
//
// If request fails at some point, then instead of `requestfinished` event (and possibly instead of 'response'
// event), the  [Page.OnRequestFailed] event is emitted.
// **NOTE** HTTP Error responses, such as 404 or 503, are still successful responses from HTTP standpoint, so request
// will complete with `requestfinished` event.
// If request gets a 'redirect' response, the request is successfully finished with the `requestfinished` event, and a
// new request is  issued to a redirected url.
type Request interface {
	// An object with all the request HTTP headers associated with this request. The header names are lower-cased.
	AllHeaders() (map[string]string, error)

	// The method returns `null` unless this request has failed, as reported by `requestfailed` event.
	Failure() error

	// Returns the [Frame] that initiated this request.
	//
	// # Details
	//
	// Note that in some cases the frame is not available, and this method will throw.
	//  - When request originates in the Service Worker. You can use `request.serviceWorker()` to check that.
	//  - When navigation request is issued before the corresponding frame is created. You can use
	//   [Request.IsNavigationRequest] to check that.
	// Here is an example that handles all the cases:
	Frame() Frame

	// An object with the request HTTP headers. The header names are lower-cased. Note that this method does not return
	// security-related headers, including cookie-related ones. You can use [Request.AllHeaders] for complete list of
	// headers that include `cookie` information.
	Headers() map[string]string

	// An array with all the request HTTP headers associated with this request. Unlike [Request.AllHeaders], header names
	// are NOT lower-cased. Headers with multiple entries, such as `Set-Cookie`, appear in the array multiple times.
	HeadersArray() ([]NameValue, error)

	// Returns the value of the header matching the name. The name is case insensitive.
	//
	//  name: Name of the header.
	HeaderValue(name string) (string, error)

	// Whether this request is driving frame's navigation.
	// Some navigation requests are issued before the corresponding frame is created, and therefore do not have
	// [Request.Frame] available.
	IsNavigationRequest() bool

	// Request's method (GET, POST, etc.)
	Method() string

	// Request's post body, if any.
	PostData() (string, error)

	// Request's post body in a binary form, if any.
	PostDataBuffer() ([]byte, error)

	// Returns parsed request's body for `form-urlencoded` and JSON as a fallback if any.
	// When the response is `application/x-www-form-urlencoded` then a key/value object of the values will be returned.
	// Otherwise it will be parsed as JSON.
	PostDataJSON(v interface{}) error

	// Request that was redirected by the server to this one, if any.
	// When the server responds with a redirect, Playwright creates a new [Request] object. The two requests are connected
	// by `redirectedFrom()` and `redirectedTo()` methods. When multiple server redirects has happened, it is possible to
	// construct the whole redirect chain by repeatedly calling `redirectedFrom()`.
	RedirectedFrom() Request

	// New request issued by the browser if the server responded with redirect.
	RedirectedTo() Request

	// Contains the request's resource type as it was perceived by the rendering engine. ResourceType will be one of the
	// following: `document`, `stylesheet`, `image`, `media`, `font`, `script`, `texttrack`, `xhr`, `fetch`,
	// `eventsource`, `websocket`, `manifest`, `other`.
	ResourceType() string

	// Returns the matching [Response] object, or `null` if the response was not received due to error.
	Response() (Response, error)

	// Returns resource size information for given request.
	Sizes() (*RequestSizesResult, error)

	// Returns resource timing information for given request. Most of the timing values become available upon the
	// response, `responseEnd` becomes available when request finishes. Find more information at
	// [Resource Timing API].
	//
	// [Resource Timing API]: https://developer.mozilla.org/en-US/docs/Web/API/PerformanceResourceTiming
	Timing() *RequestTiming

	// URL of the request.
	URL() string
}

// [Response] class represents responses which are received by page.
type Response interface {
	// An object with all the response HTTP headers associated with this response.
	AllHeaders() (map[string]string, error)

	// Returns the buffer with response body.
	Body() ([]byte, error)

	// Waits for this response to finish, returns always `null`.
	Finished() error

	// Returns the [Frame] that initiated this response.
	Frame() Frame

	// Indicates whether this Response was fulfilled by a Service Worker's Fetch Handler (i.e. via
	// [FetchEvent.RespondWith].
	//
	// [FetchEvent.RespondWith]: https://developer.mozilla.org/en-US/docs/Web/API/FetchEvent/respondWith)
	FromServiceWorker() bool

	// An object with the response HTTP headers. The header names are lower-cased. Note that this method does not return
	// security-related headers, including cookie-related ones. You can use [Response.AllHeaders] for complete list of
	// headers that include `cookie` information.
	Headers() map[string]string

	// An array with all the request HTTP headers associated with this response. Unlike [Response.AllHeaders], header
	// names are NOT lower-cased. Headers with multiple entries, such as `Set-Cookie`, appear in the array multiple times.
	HeadersArray() ([]NameValue, error)

	// Returns the value of the header matching the name. The name is case insensitive. If multiple headers have the same
	// name (except `set-cookie`), they are returned as a list separated by `, `. For `set-cookie`, the `\n` separator is
	// used. If no headers are found, `null` is returned.
	//
	//  name: Name of the header.
	HeaderValue(name string) (string, error)

	// Returns all values of the headers matching the name, for example `set-cookie`. The name is case insensitive.
	//
	//  name: Name of the header.
	HeaderValues(name string) ([]string, error)

	// Returns the JSON representation of response body.
	// This method will throw if the response body is not parsable via `JSON.parse`.
	JSON(v interface{}) error

	// Contains a boolean stating whether the response was successful (status in the range 200-299) or not.
	Ok() bool

	// Returns the matching [Request] object.
	Request() Request

	// Returns SSL and other security information.
	SecurityDetails() (*ResponseSecurityDetailsResult, error)

	// Returns the IP address and port of the server.
	ServerAddr() (*ResponseServerAddrResult, error)

	// Contains the status code of the response (e.g., 200 for a success).
	Status() int

	// Contains the status text of the response (e.g. usually an "OK" for a success).
	StatusText() string

	// Returns the text representation of response body.
	Text() (string, error)

	// Contains the URL of the response.
	URL() string
}

// Whenever a network route is set up with [Page.Route] or [BrowserContext.Route], the `Route` object allows to handle
// the route.
// Learn more about [networking].
//
// [networking]: https://playwright.dev/docs/network
type Route interface {
	// Aborts the route's request.
	Abort(errorCode ...string) error

	// Continues route's request with optional overrides.
	//
	// # Details
	//
	// Note that any overrides such as “url” or “headers” only apply to the request being routed. If this request results
	// in a redirect, overrides will not be applied to the new redirected request. If you want to propagate a header
	// through redirects, use the combination of [Route.Fetch] and [Route.Fulfill] instead.
	Continue(options ...RouteContinueOptions) error

	// When several routes match the given pattern, they run in the order opposite to their registration. That way the
	// last registered route can always override all the previous ones. In the example below, request will be handled by
	// the bottom-most handler first, then it'll fall back to the previous one and in the end will be aborted by the first
	// registered route.
	Fallback(options ...RouteFallbackOptions) error

	// Performs the request and fetches result without fulfilling it, so that the response could be modified and then
	// fulfilled.
	//
	// # Details
	//
	// Note that “headers” option will apply to the fetched request as well as any redirects initiated by it. If you want
	// to only apply “headers” to the original request, but not to redirects, look into [Route.Continue] instead.
	Fetch(options ...RouteFetchOptions) (APIResponse, error)

	// Fulfills route's request with given response.
	Fulfill(options ...RouteFulfillOptions) error

	// A request to be routed.
	Request() Request
}

// Selectors can be used to install custom selector engines. See [extensibility] for more
// information.
//
// [extensibility]: https://playwright.dev/docs/extensibility
type Selectors interface {
	// Selectors must be registered before creating the page.
	//
	// 1. name: Name that is used in selectors as a prefix, e.g. `{name: 'foo'}` enables `foo=myselectorbody` selectors. May only
	//    contain `[a-zA-Z0-9_]` characters.
	// 2. script: Script that evaluates to a selector engine instance. The script is evaluated in the page context.
	Register(name string, script Script, options ...SelectorsRegisterOptions) error

	// Defines custom attribute name to be used in [Page.GetByTestId]. `data-testid` is used by default.
	//
	//  attributeName: Test id attribute name.
	SetTestIdAttribute(attributeName string)
}

// The Touchscreen class operates in main-frame CSS pixels relative to the top-left corner of the viewport. Methods on
// the touchscreen can only be used in browser contexts that have been initialized with `hasTouch` set to true.
type Touchscreen interface {
	// Dispatches a `touchstart` and `touchend` event with a single touch at the position (“x”,“y”).
	// **NOTE** [Page.Tap] the method will throw if “hasTouch” option of the browser context is false.
	Tap(x int, y int) error
}

// API for collecting and saving Playwright traces. Playwright traces can be opened in
// [Trace Viewer] after Playwright script runs.
// Start recording a trace before performing actions. At the end, stop tracing and save it to a file.
//
// [Trace Viewer]: https://playwright.dev/docs/trace-viewer
type Tracing interface {
	// Start tracing.
	Start(options ...TracingStartOptions) error

	// Start a new trace chunk. If you'd like to record multiple traces on the same [BrowserContext], use [Tracing.Start]
	// once, and then create multiple trace chunks with [Tracing.StartChunk] and [Tracing.StopChunk].
	StartChunk(options ...TracingStartChunkOptions) error

	// Stop tracing.
	Stop(path ...string) error

	// Stop the trace chunk. See [Tracing.StartChunk] for more details about multiple trace chunks.
	StopChunk(path ...string) error
}

// When browser context is created with the `recordVideo` option, each page has a video object associated with it.
type Video interface {
	// Deletes the video file. Will wait for the video to finish if necessary.
	Delete() error

	// Returns the file system path this video will be recorded to. The video is guaranteed to be written to the
	// filesystem upon closing the browser context. This method throws when connected remotely.
	Path() (string, error)

	// Saves the video to a user-specified path. It is safe to call this method while the video is still in progress, or
	// after the page has closed. This method waits until the page is closed and the video is fully saved.
	//
	//  path: Path where the video should be saved.
	SaveAs(path string) error
}

// [WebError] class represents an unhandled exception thrown in the page. It is dispatched via the
// [BrowserContext.OnWebError] event.
type WebError interface {
	// The page that produced this unhandled exception, if any.
	Page() Page

	// Unhandled error that was thrown.
	Error() error
}

// The [WebSocket] class represents websocket connections in the page.
type WebSocket interface {
	// Fired when the websocket closes.
	OnClose(fn func(WebSocket))

	// Fired when the websocket receives a frame.
	OnFrameReceived(fn func([]byte))

	// Fired when the websocket sends a frame.
	OnFrameSent(fn func([]byte))

	// Fired when the websocket has an error.
	OnSocketError(fn func(string))

	// Indicates that the web socket has been closed.
	IsClosed() bool

	// Contains the URL of the WebSocket.
	URL() string

	// Waits for event to fire and passes its value into the predicate function. Returns when the predicate returns truthy
	// value. Will throw an error if the webSocket is closed before the event is fired. Returns the event data value.
	//
	//  event: Event name, same one would pass into `webSocket.on(event)`.
	ExpectEvent(event string, cb func() error, options ...WebSocketExpectEventOptions) (interface{}, error)

	// **NOTE** In most cases, you should use [WebSocket.ExpectEvent].
	// Waits for given `event` to fire. If predicate is provided, it passes event's value into the `predicate` function
	// and waits for `predicate(event)` to return a truthy value. Will throw an error if the socket is closed before the
	// `event` is fired.
	//
	//  event: Event name, same one typically passed into `*.on(event)`.
	WaitForEvent(event string, options ...WebSocketWaitForEventOptions) (interface{}, error)
}

// The Worker class represents a [WebWorker].
// `worker` event is emitted on the page object to signal a worker creation. `close` event is emitted on the worker
// object when the worker is gone.
//
// [WebWorker]: https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API
type Worker interface {
	// Emitted when this dedicated [WebWorker] is
	// terminated.
	//
	// [WebWorker]: https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API
	OnClose(fn func(Worker))

	// Returns the return value of “expression”.
	// If the function passed to the [Worker.Evaluate] returns a [Promise], then [Worker.Evaluate] would wait for the
	// promise to resolve and return its value.
	// If the function passed to the [Worker.Evaluate] returns a non-[Serializable] value, then [Worker.Evaluate] returns
	// `undefined`. Playwright also supports transferring some additional values that are not serializable by `JSON`:
	// `-0`, `NaN`, `Infinity`, `-Infinity`.
	//
	// 1. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 2. arg: Optional argument to pass to “expression”.
	Evaluate(expression string, arg ...interface{}) (interface{}, error)

	// Returns the return value of “expression” as a [JSHandle].
	// The only difference between [Worker.Evaluate] and [Worker.EvaluateHandle] is that [Worker.EvaluateHandle] returns
	// [JSHandle].
	// If the function passed to the [Worker.EvaluateHandle] returns a [Promise], then [Worker.EvaluateHandle] would wait
	// for the promise to resolve and return its value.
	//
	// 1. expression: JavaScript expression to be evaluated in the browser context. If the expression evaluates to a function, the
	//    function is automatically invoked.
	// 2. arg: Optional argument to pass to “expression”.
	EvaluateHandle(expression string, arg ...interface{}) (JSHandle, error)

	URL() string
}
