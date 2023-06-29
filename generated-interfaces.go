package playwright

// Exposes API that can be used for the Web API testing. This class is used for creating `APIRequestContext` instance
// which in turn can be used for sending web requests. An instance of this class can be obtained via
// Playwright.request(). For more information see `APIRequestContext`.
type APIRequest interface {
	EventEmitter
	// Creates new instances of `APIRequestContext`.
	NewContext(options ...APIRequestNewContextOptions) (APIRequestContext, error)
}

// This API is used for the Web API testing. You can use it to trigger API endpoints, configure micro-services,
// prepare environment or the service to your e2e test.
// Each Playwright browser context has associated with it `APIRequestContext` instance which shares cookie storage
// with the browser context and can be accessed via BrowserContext.request() or
// Page.request(). It is also possible to create a new APIRequestContext instance manually by calling
// APIRequest.newContext().
// **Cookie management**
// `APIRequestContext` returned by BrowserContext.request() and Page.request() shares cookie
// storage with the corresponding `BrowserContext`. Each API request will have `Cookie` header populated with the
// values from the browser context. If the API response contains `Set-Cookie` header it will automatically update
// `BrowserContext` cookies and requests made from the page will pick them up. This means that if you log in using
// this API, your e2e test will be logged in and vice versa.
// If you want API requests to not interfere with the browser cookies you should create a new `APIRequestContext` by
// calling APIRequest.newContext(). Such `APIRequestContext` object will have its own isolated cookie
// storage.
type APIRequestContext interface {
	EventEmitter
	// Sends HTTP(S) [DELETE](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/DELETE) request and returns its
	// response. The method will populate request cookies from the context and update context cookies from the response.
	// The method will automatically follow redirects.
	Delete(url string, options ...APIRequestContextDeleteOptions) (APIResponse, error)
	// All responses returned by APIRequestContext.get() and similar methods are stored in the memory, so that
	// you can later call APIResponse.body(). This method discards all stored responses, and makes
	// APIResponse.body() throw "Response disposed" error.
	Dispose() error
	// Sends HTTP(S) request and returns its response. The method will populate request cookies from the context and
	// update context cookies from the response. The method will automatically follow redirects. JSON objects can be
	// passed directly to the request.
	Fetch(urlOrRequest interface{}, options ...APIRequestContextFetchOptions) (APIResponse, error)
	// Sends HTTP(S) [GET](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET) request and returns its
	// response. The method will populate request cookies from the context and update context cookies from the response.
	// The method will automatically follow redirects.
	Get(url string, options ...APIRequestContextGetOptions) (APIResponse, error)
	// Sends HTTP(S) [HEAD](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/HEAD) request and returns its
	// response. The method will populate request cookies from the context and update context cookies from the response.
	// The method will automatically follow redirects.
	Head(url string, options ...APIRequestContextHeadOptions) (APIResponse, error)
	// Sends HTTP(S) [POST](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST) request and returns its
	// response. The method will populate request cookies from the context and update context cookies from the response.
	// The method will automatically follow redirects.
	Post(url string, options ...APIRequestContextPostOptions) (APIResponse, error)
	// Sends HTTP(S) [PUT](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PUT) request and returns its
	// response. The method will populate request cookies from the context and update context cookies from the response.
	// The method will automatically follow redirects.
	Put(url string, options ...APIRequestContextPutOptions) (APIResponse, error)
	// Sends HTTP(S) [PATCH](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PATCH) request and returns its
	// response. The method will populate request cookies from the context and update context cookies from the response.
	// The method will automatically follow redirects.
	Patch(url string, options ...APIRequestContextPatchOptions) (APIResponse, error)
	// Returns storage state for this request context, contains current cookies and local storage snapshot if it was
	// passed to the constructor.
	StorageState(path ...string) (*StorageState, error)
}

// `APIResponse` class represents responses returned by APIRequestContext.get() and similar methods.
type APIResponse interface {
	// Returns the buffer with response body.
	Body() ([]byte, error)
	// Disposes the body of this response. If not called then the body will stay in memory until the context closes.
	Dispose() error
	// An object with all the response HTTP headers associated with this response.
	Headers() map[string]string
	// An array with all the request HTTP headers associated with this response. Header names are not lower-cased. Headers
	// with multiple entries, such as `Set-Cookie`, appear in the array multiple times.
	HeadersArray() HeadersArray
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

// The `APIResponseAssertions` class provides assertion methods that can be used to make assertions about the
// `APIResponse` in the tests.
type APIResponseAssertions interface {
	// Makes the assertion check for the opposite condition. For example, this code tests that the response status is not
	// successful:
	Not() APIResponseAssertions
	// Ensures the response status code is within `200..299` range.
	ToBeOK() error
	// The opposite of APIResponseAssertions.toBeOK().
	NotToBeOK() error
}

type BindingCall interface {
	Call(f BindingCallFunction)
}

// A Browser is created via BrowserType.launch(). An example of using a `Browser` to create a `Page`:
type Browser interface {
	EventEmitter
	// Get the browser type (chromium, firefox or webkit) that the browser belongs to.
	BrowserType() BrowserType
	// In case this browser is obtained using BrowserType.launch(), closes the browser and all of its pages (if
	// any were opened).
	// In case this browser is connected to, clears all created contexts belonging to this browser and disconnects from
	// the browser server.
	// **NOTE** This is similar to force quitting the browser. Therefore, you should call BrowserContext.close()
	// on any `BrowserContext`'s you explicitly created earlier with Browser.newContext() **before** calling
	// Browser.close().
	// The `Browser` object itself is considered to be disposed and cannot be used anymore.
	Close() error
	// Returns an array of all open browser contexts. In a newly created browser, this will return zero browser contexts.
	Contexts() []BrowserContext
	// Indicates that the browser is connected.
	IsConnected() bool
	// Creates a new browser context. It won't share cookies/cache with other browser contexts.
	// **NOTE** If directly using this method to create `BrowserContext`s, it is best practice to explicitly close the
	// returned context via BrowserContext.close() when your code is done with the `BrowserContext`, and before
	// calling Browser.close(). This will ensure the `context` is closed gracefully and any artifacts—like HARs
	// and videos—are fully flushed and saved.
	NewContext(options ...BrowserNewContextOptions) (BrowserContext, error)
	// Creates a new page in a new browser context. Closing this page will close the context as well.
	// This is a convenience API that should only be used for the single-page scenarios and short snippets. Production
	// code and testing frameworks should explicitly create Browser.newContext() followed by the
	// BrowserContext.newPage() to control their exact life times.
	NewPage(options ...BrowserNewContextOptions) (Page, error)
	// **NOTE** CDP Sessions are only supported on Chromium-based browsers.
	// Returns the newly created browser session.
	NewBrowserCDPSession() (CDPSession, error)
	// **NOTE** This API controls
	// [Chromium Tracing](https://www.chromium.org/developers/how-tos/trace-event-profiling-tool) which is a low-level
	// chromium-specific debugging tool. API to control [Playwright Tracing](../trace-viewer) could be found
	// [here](./class-tracing).
	// You can use Browser.startTracing() and Browser.stopTracing() to create a trace file that can be
	// opened in Chrome DevTools performance panel.
	StartTracing(options ...BrowserStartTracingOptions) error
	// **NOTE** This API controls
	// [Chromium Tracing](https://www.chromium.org/developers/how-tos/trace-event-profiling-tool) which is a low-level
	// chromium-specific debugging tool. API to control [Playwright Tracing](../trace-viewer) could be found
	// [here](./class-tracing).
	// Returns the buffer with trace data.
	StopTracing() ([]byte, error)
	// Returns the browser version.
	Version() string
}

// The `CDPSession` instances are used to talk raw Chrome Devtools Protocol:
// - protocol methods can be called with `session.send` method.
// - protocol events can be subscribed to with `session.on` method.
// Useful links:
// - Documentation on DevTools Protocol can be found here:
// [DevTools Protocol Viewer](https://chromedevtools.github.io/devtools-protocol/).
// - Getting Started with DevTools Protocol:
// https://github.com/aslushnikov/getting-started-with-cdp/blob/master/README.md
type CDPSession interface {
	EventEmitter
	// Detaches the CDPSession from the target. Once detached, the CDPSession object won't emit any events and can't be
	// used to send messages.
	Detach() error
	Send(method string, params map[string]interface{}) (interface{}, error)
}

// BrowserContexts provide a way to operate multiple independent browser sessions.
// If a page opens another page, e.g. with a `window.open` call, the popup will belong to the parent page's browser
// context.
// Playwright allows creating "incognito" browser contexts with Browser.newContext() method. "Incognito"
// browser contexts don't write any browsing data to disk.
type BrowserContext interface {
	EventEmitter
	// Adds cookies into this browser context. All pages within this context will have these cookies installed. Cookies
	// can be obtained via BrowserContext.cookies().
	AddCookies(cookies ...OptionalCookie) error
	// Adds a script which would be evaluated in one of the following scenarios:
	// - Whenever a page is created in the browser context or is navigated.
	// - Whenever a child frame is attached or navigated in any page in the browser context. In this case, the script is
	// evaluated in the context of the newly attached frame.
	// The script is evaluated after the document was created but before any of its scripts were run. This is useful to
	// amend the JavaScript environment, e.g. to seed `Math.random`.
	AddInitScript(script BrowserContextAddInitScriptOptions) error
	// Returns the browser instance of the context. If it was launched as a persistent context null gets returned.
	Browser() Browser
	// Clears context cookies.
	ClearCookies() error
	// Clears all permission overrides for the browser context.
	ClearPermissions() error
	// Closes the browser context. All the pages that belong to the browser context will be closed.
	// **NOTE** The default browser context cannot be closed.
	Close() error
	// If no URLs are specified, this method returns all cookies. If URLs are specified, only cookies that affect those
	// URLs are returned.
	Cookies(urls ...string) ([]*Cookie, error)
	// Performs action and waits for a `ConsoleMessage` to be logged by in the pages in the context. If predicate is
	// provided, it passes `ConsoleMessage` value into the `predicate` function and waits for `predicate(message)` to
	// return a truthy value. Will throw an error if the page is closed before the [`event: BrowserContext.console`] event
	// is fired.
	ExpectConsoleMessage(cb func() error, options ...BrowserContextExpectConsoleMessageOptions) (ConsoleMessage, error)
	// Waits for event to fire and passes its value into the predicate function. Returns when the predicate returns truthy
	// value. Will throw an error if the context closes before the event is fired. Returns the event data value.
	ExpectEvent(event string, cb func() error, options ...BrowserContextWaitForEventOptions) (interface{}, error)
	// Performs action and waits for a new `Page` to be created in the context. If predicate is provided, it passes `Page`
	// value into the `predicate` function and waits for `predicate(event)` to return a truthy value. Will throw an error
	// if the context closes before new `Page` is created.
	ExpectPage(cb func() error, options ...BrowserContextExpectPageOptions) (Page, error)
	// The method adds a function called `name` on the `window` object of every frame in every page in the context. When
	// called, the function executes `callback` and returns a [Promise] which resolves to the return value of `callback`.
	// If the `callback` returns a [Promise], it will be awaited.
	// The first argument of the `callback` function contains information about the caller: `{ browserContext:
	// BrowserContext, page: Page, frame: Frame }`.
	// See Page.exposeBinding() for page-only version.
	ExposeBinding(name string, binding BindingCallFunction, handle ...bool) error
	// The method adds a function called `name` on the `window` object of every frame in every page in the context. When
	// called, the function executes `callback` and returns a [Promise] which resolves to the return value of `callback`.
	// If the `callback` returns a [Promise], it will be awaited.
	// See Page.exposeFunction() for page-only version.
	ExposeFunction(name string, binding ExposedFunction) error
	// Grants specified permissions to the browser context. Only grants corresponding permissions to the given origin if
	// specified.
	GrantPermissions(permissions []string, options ...BrowserContextGrantPermissionsOptions) error
	// **NOTE** CDP sessions are only supported on Chromium-based browsers.
	// Returns the newly created session.
	NewCDPSession(page Page) (CDPSession, error)
	// Creates a new page in the browser context.
	NewPage(options ...BrowserNewPageOptions) (Page, error)
	// Returns all open pages in the context.
	Pages() []Page
	// **NOTE** Service workers are only supported on Chromium-based browsers.
	// All existing service workers in the context.
	ServiceWorkers() []Worker
	// This setting will change the default maximum navigation time for the following methods and related shortcuts:
	// - Page.goBack()
	// - Page.goForward()
	// - Page.goto()
	// - Page.reload()
	// - Page.setContent()
	// - Page.waitForNavigation()
	// **NOTE** Page.setDefaultNavigationTimeout() and Page.setDefaultTimeout() take priority over
	// BrowserContext.setDefaultNavigationTimeout().
	SetDefaultNavigationTimeout(timeout float64)
	// This setting will change the default maximum time for all the methods accepting `timeout` option.
	// **NOTE** Page.setDefaultNavigationTimeout(), Page.setDefaultTimeout() and
	// BrowserContext.setDefaultNavigationTimeout() take priority over
	// BrowserContext.setDefaultTimeout().
	SetDefaultTimeout(timeout float64)
	// The extra HTTP headers will be sent with every request initiated by any page in the context. These headers are
	// merged with page-specific extra HTTP headers set with Page.setExtraHTTPHeaders(). If page overrides a
	// particular header, page-specific header value will be used instead of the browser context header value.
	// **NOTE** BrowserContext.setExtraHTTPHeaders() does not guarantee the order of headers in the outgoing
	// requests.
	SetExtraHTTPHeaders(headers map[string]string) error
	// Sets the context's geolocation. Passing `null` or `undefined` emulates position unavailable.
	SetGeolocation(gelocation *Geolocation) error
	// API testing helper associated with this context. Requests made with this API will use context cookies.
	Request() APIRequestContext
	ResetGeolocation() error
	// Routing provides the capability to modify network requests that are made by any page in the browser context. Once
	// route is enabled, every request matching the url pattern will stall unless it's continued, fulfilled or aborted.
	// **NOTE** BrowserContext.route() will not intercept requests intercepted by Service Worker. See
	// [this](https://github.com/microsoft/playwright/issues/1090) issue. We recommend disabling Service Workers when
	// using request interception by setting `Browser.newContext.serviceWorkers` to `'block'`.
	Route(url interface{}, handler routeHandler, times ...int) error
	SetOffline(offline bool) error
	// If specified the network requests that are made in the context will be served from the HAR file. Read more about
	// [Replaying from HAR](../network.md#replaying-from-har).
	// Playwright will not serve requests intercepted by Service Worker from the HAR file. See
	// [this](https://github.com/microsoft/playwright/issues/1090) issue. We recommend disabling Service Workers when
	// using request interception by setting `Browser.newContext.serviceWorkers` to `'block'`.
	RouteFromHAR(har string, options ...BrowserContextRouteFromHAROptions) error
	// Returns storage state for this browser context, contains current cookies and local storage snapshot.
	StorageState(path ...string) (*StorageState, error)
	// Removes a route created with BrowserContext.route(). When `handler` is not specified, removes all routes
	// for the `url`.
	Unroute(url interface{}, handler ...routeHandler) error
	// **NOTE** In most cases, you should use BrowserContext.ExpectForEvent().
	// Waits for given `event` to fire. If predicate is provided, it passes event's value into the `predicate` function
	// and waits for `predicate(event)` to return a truthy value. Will throw an error if the browser context is closed
	// before the `event` is fired.
	WaitForEvent(event string, options ...BrowserContextWaitForEventOptions) (interface{}, error)
	Tracing() Tracing
	// **NOTE** Background pages are only supported on Chromium-based browsers.
	// All existing background pages in the context.
	BackgroundPages() []Page
}

// API for collecting and saving Playwright traces. Playwright traces can be opened in
// [Trace Viewer](../trace-viewer.md) after Playwright script runs.
// Start recording a trace before performing actions. At the end, stop tracing and save it to a file.
type Tracing interface {
	// Start tracing.
	Start(options ...TracingStartOptions) error
	// Stop tracing.
	Stop(options ...TracingStopOptions) error
	// Start a new trace chunk. If you'd like to record multiple traces on the same `BrowserContext`, use
	// Tracing.start() once, and then create multiple trace chunks with Tracing.startChunk() and
	// Tracing.stopChunk().
	StartChunk(options ...TracingStartChunkOptions) error
	// Stop the trace chunk. See Tracing.startChunk() for more details about multiple trace chunks.
	StopChunk(options ...TracingStopChunkOptions) error
}

// BrowserType provides methods to launch a specific browser instance or connect to an existing one. The following is
// a typical example of using Playwright to drive automation:
type BrowserType interface {
	// A path where Playwright expects to find a bundled browser executable.
	ExecutablePath() string
	// Returns the browser instance.
	Launch(options ...BrowserTypeLaunchOptions) (Browser, error)
	// Returns the persistent browser context instance.
	// Launches browser that uses persistent storage located at `userDataDir` and returns the only context. Closing this
	// context will automatically close the browser.
	LaunchPersistentContext(userDataDir string, options ...BrowserTypeLaunchPersistentContextOptions) (BrowserContext, error)
	// Returns browser name. For example: `'chromium'`, `'webkit'` or `'firefox'`.
	Name() string
	// This method attaches Playwright to an existing browser instance. When connecting to another browser launched via
	// `BrowserType.launchServer` in Node.js, the major and minor version needs to match the client version (1.2.3 → is
	// compatible with 1.2.x).
	Connect(url string, options ...BrowserTypeConnectOptions) (Browser, error)
	// This method attaches Playwright to an existing browser instance using the Chrome DevTools Protocol.
	// The default browser context is accessible via Browser.contexts().
	// **NOTE** Connecting over the Chrome DevTools Protocol is only supported for Chromium-based browsers.
	ConnectOverCDP(endpointURL string, options ...BrowserTypeConnectOverCDPOptions) (Browser, error)
}

// `ConsoleMessage` objects are dispatched by page via the [`event: Page.console`] event. For each console messages
// logged in the page there will be corresponding event in the Playwright context.
type ConsoleMessage interface {
	// List of arguments passed to a `console` function call. See also [`event: Page.console`].
	Args() []JSHandle
	Location() ConsoleMessageLocation
	// The page that produced this console message, if any.
	Page() Page
	String() string
	// The text of the console message.
	Text() string
	// One of the following values: `'log'`, `'debug'`, `'info'`, `'error'`, `'warning'`, `'dir'`, `'dirxml'`, `'table'`,
	// `'trace'`, `'clear'`, `'startGroup'`, `'startGroupCollapsed'`, `'endGroup'`, `'assert'`, `'profile'`,
	// `'profileEnd'`, `'count'`, `'timeEnd'`.
	Type() string
}

// `Dialog` objects are dispatched by page via the [`event: Page.dialog`] event.
// An example of using `Dialog` class:
// **NOTE** Dialogs are dismissed automatically, unless there is a [`event: Page.dialog`] listener. When listener is
// present, it **must** either Dialog.accept() or Dialog.dismiss() the dialog - otherwise the page
// will [freeze](https://developer.mozilla.org/en-US/docs/Web/JavaScript/EventLoop#never_blocking) waiting for the
// dialog, and actions like click will never finish.
type Dialog interface {
	// Returns when the dialog has been accepted.
	Accept(texts ...string) error
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

// `Download` objects are dispatched by page via the [`event: Page.download`] event.
// All the downloaded files belonging to the browser context are deleted when the browser context is closed.
// Download event is emitted once the download starts. Download path becomes available once download completes:
type Download interface {
	// Deletes the downloaded file. Will wait for the download to finish if necessary.
	Delete() error
	// Returns download error if any. Will wait for the download to finish if necessary.
	Failure() (string, error)
	// Returns path to the downloaded file in case of successful download. The method will wait for the download to finish
	// if necessary. The method throws when connected remotely.
	// Note that the download's file name is a random GUID, use Download.suggestedFilename() to get suggested
	// file name.
	Path() (string, error)
	// Copy the download to a user-specified path. It is safe to call this method while the download is still in progress.
	// Will wait for the download to finish if necessary.
	SaveAs(path string) error
	String() string
	// Returns suggested filename for this download. It is typically computed by the browser from the
	// [`Content-Disposition`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition) response
	// header or the `download` attribute. See the spec on [whatwg](https://html.spec.whatwg.org/#downloading-resources).
	// Different browsers can use different logic for computing it.
	SuggestedFilename() string
	// Returns downloaded url.
	URL() string
	// Get the page that the download belongs to.
	Page() Page
	// Cancels a download. Will not fail if the download is already finished or canceled. Upon successful cancellations,
	// `download.failure()` would resolve to `'canceled'`.
	Cancel() error
}

// ElementHandle represents an in-page DOM element. ElementHandles can be created with the
// Page.querySelector() method.
// **NOTE** The use of ElementHandle is discouraged, use `Locator` objects and web-first assertions instead.
// ElementHandle prevents DOM element from garbage collection unless the handle is disposed with
// JSHandle.dispose(). ElementHandles are auto-disposed when their origin frame gets navigated.
// ElementHandle instances can be used as an argument in Page.evalOnSelector() and Page.evaluate()
// methods.
// The difference between the `Locator` and ElementHandle is that the ElementHandle points to a particular element,
// while `Locator` captures the logic of how to retrieve an element.
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
	// [Element.getBoundingClientRect](https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect).
	// That means `x` and/or `y` may be negative.
	// Elements from child frames return the bounding box relative to the main frame, unlike the
	// [Element.getBoundingClientRect](https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect).
	// Assuming the page is static, it is safe to use bounding box coordinates to perform input. For example, the
	// following snippet should click the center of the element.
	BoundingBox() (*Rect, error)
	// This method checks the element by performing the following steps:
	// 1. Ensure that element is a checkbox or a radio input. If not, this method throws. If the element is already
	// checked, this method returns immediately.
	// 1. Wait for [actionability](../actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now checked. If not, this method throws.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	Check(options ...ElementHandleCheckOptions) error
	// This method checks or unchecks an element by performing the following steps:
	// 1. Ensure that element is a checkbox or a radio input. If not, this method throws.
	// 1. If the element already has the right checked state, this method returns immediately.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now checked or unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	SetChecked(checked bool, options ...ElementHandleSetCheckedOptions) error
	// This method clicks the element by performing the following steps:
	// 1. Wait for [actionability](../actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to click in the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	Click(options ...ElementHandleClickOptions) error
	// Returns the content frame for element handles referencing iframe nodes, or `null` otherwise
	ContentFrame() (Frame, error)
	// This method double clicks the element by performing the following steps:
	// 1. Wait for [actionability](../actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to double click in the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set. Note that if
	// the first click of the `dblclick()` triggers a navigation event, this method will throw.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	// **NOTE** `elementHandle.dblclick()` dispatches two `click` events and a single `dblclick` event.
	Dblclick(options ...ElementHandleDblclickOptions) error
	// The snippet below dispatches the `click` event on the element. Regardless of the visibility state of the element,
	// `click` is dispatched. This is equivalent to calling
	// [element.click()](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click).
	DispatchEvent(typ string, initObjects ...interface{}) error
	// Returns the return value of `expression`.
	// The method finds an element matching the specified selector in the `ElementHandle`s subtree and passes it as a
	// first argument to `expression`. If no elements match the selector, the method throws an error.
	// If `expression` returns a [Promise], then ElementHandle.evalOnSelector() would wait for the promise to
	// resolve and return its value.
	EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `expression`.
	// The method finds all elements matching the specified selector in the `ElementHandle`'s subtree and passes an array
	// of matched elements as a first argument to `expression`.
	// If `expression` returns a [Promise], then ElementHandle.evalOnSelectorAll() would wait for the promise to
	// resolve and return its value.
	EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
	// This method waits for [actionability](../actionability.md) checks, focuses the element, fills it and triggers an
	// `input` event after filling. Note that you can pass an empty string to clear the input field.
	// If the target element is not an `<input>`, `<textarea>` or `[contenteditable]` element, this method throws an
	// error. However, if the element is inside the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), the control will be filled
	// instead.
	// To send fine-grained keyboard events, use ElementHandle.type().
	Fill(value string, options ...ElementHandleFillOptions) error
	// Calls [focus](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus) on the element.
	Focus() error
	// Returns element attribute value.
	GetAttribute(name string) (string, error)
	// This method hovers over the element by performing the following steps:
	// 1. Wait for [actionability](../actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to hover over the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	Hover(options ...ElementHandleHoverOptions) error
	// Returns the `element.innerHTML`.
	InnerHTML() (string, error)
	// Returns the `element.innerText`.
	InnerText() (string, error)
	// Returns whether the element is checked. Throws if the element is not a checkbox or radio input.
	IsChecked() (bool, error)
	// Returns whether the element is disabled, the opposite of [enabled](../actionability.md#enabled).
	IsDisabled() (bool, error)
	// Returns whether the element is [editable](../actionability.md#editable).
	IsEditable() (bool, error)
	// Returns whether the element is [enabled](../actionability.md#enabled).
	IsEnabled() (bool, error)
	// Returns whether the element is hidden, the opposite of [visible](../actionability.md#visible).
	IsHidden() (bool, error)
	// Returns whether the element is [visible](../actionability.md#visible).
	IsVisible() (bool, error)
	// Returns the frame containing the given element.
	OwnerFrame() (Frame, error)
	// Focuses the element, and then uses Keyboard.down() and Keyboard.up().
	// `key` can specify the intended
	// [keyboardEvent.key](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key) value or a single character
	// to generate the text for. A superset of the `key` values can be found
	// [here](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values). Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`,
	// etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When specified with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	Press(key string, options ...ElementHandlePressOptions) error
	// The method finds an element matching the specified selector in the `ElementHandle`'s subtree. If no elements match
	// the selector, returns `null`.
	QuerySelector(selector string) (ElementHandle, error)
	// The method finds all elements matching the specified selector in the `ElementHandle`s subtree. If no elements match
	// the selector, returns empty array.
	QuerySelectorAll(selector string) ([]ElementHandle, error)
	// This method captures a screenshot of the page, clipped to the size and position of this particular element. If the
	// element is covered by other elements, it will not be actually visible on the screenshot. If the element is a
	// scrollable container, only the currently scrolled content will be visible on the screenshot.
	// This method waits for the [actionability](../actionability.md) checks, then scrolls element into view before taking
	// a screenshot. If the element is detached from DOM, the method throws an error.
	// Returns the buffer with the captured screenshot.
	Screenshot(options ...ElementHandleScreenshotOptions) ([]byte, error)
	// This method waits for [actionability](../actionability.md) checks, then tries to scroll element into view, unless
	// it is completely visible as defined by
	// [IntersectionObserver](https://developer.mozilla.org/en-US/docs/Web/API/Intersection_Observer_API)'s `ratio`.
	// Throws when `elementHandle` does not point to an element
	// [connected](https://developer.mozilla.org/en-US/docs/Web/API/Node/isConnected) to a Document or a ShadowRoot.
	ScrollIntoViewIfNeeded(options ...ElementHandleScrollIntoViewIfNeededOptions) error
	// This method waits for [actionability](../actionability.md) checks, waits until all specified options are present in
	// the `<select>` element and selects these options.
	// If the target element is not a `<select>` element, this method throws an error. However, if the element is inside
	// the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), the control will be used
	// instead.
	// Returns the array of option values that have been successfully selected.
	// Triggers a `change` and `input` event once all the provided options have been selected.
	SelectOption(values SelectOptionValues, options ...ElementHandleSelectOptionOptions) ([]string, error)
	// This method waits for [actionability](../actionability.md) checks, then focuses the element and selects all its
	// text content.
	// If the element is inside the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), focuses and selects text in
	// the control instead.
	SelectText(options ...ElementHandleSelectTextOptions) error
	// Sets the value of the file input to these file paths or files. If some of the `filePaths` are relative paths, then
	// they are resolved relative to the current working directory. For empty array, clears the selected files.
	// This method expects `ElementHandle` to point to an
	// [input element](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input). However, if the element is inside
	// the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), targets the control instead.
	SetInputFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error
	// This method taps the element by performing the following steps:
	// 1. Wait for [actionability](../actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.touchscreen() to tap the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	// **NOTE** `elementHandle.tap()` requires that the `hasTouch` option of the browser context be set to true.
	Tap(options ...ElementHandleTapOptions) error
	// Returns the `node.textContent`.
	TextContent() (string, error)
	// Focuses the element, and then sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the
	// text.
	// To press a special key, like `Control` or `ArrowDown`, use ElementHandle.press().
	Type(value string, options ...ElementHandleTypeOptions) error
	// This method checks the element by performing the following steps:
	// 1. Ensure that element is a checkbox or a radio input. If not, this method throws. If the element is already
	// unchecked, this method returns immediately.
	// 1. Wait for [actionability](../actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now unchecked. If not, this method throws.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	Uncheck(options ...ElementHandleUncheckOptions) error
	// Returns when the element satisfies the `state`.
	// Depending on the `state` parameter, this method waits for one of the [actionability](../actionability.md) checks to
	// pass. This method throws when the element is detached while waiting, unless waiting for the `"hidden"` state.
	// - `"visible"` Wait until the element is [visible](../actionability.md#visible).
	// - `"hidden"` Wait until the element is [not visible](../actionability.md#visible) or
	// [not attached](../actionability.md#attached). Note that waiting for hidden does not throw when the element
	// detaches.
	// - `"stable"` Wait until the element is both [visible](../actionability.md#visible) and
	// [stable](../actionability.md#stable).
	// - `"enabled"` Wait until the element is [enabled](../actionability.md#enabled).
	// - `"disabled"` Wait until the element is [not enabled](../actionability.md#enabled).
	// - `"editable"` Wait until the element is [editable](../actionability.md#editable).
	// If the element does not satisfy the condition for the `timeout` milliseconds, this method will throw.
	WaitForElementState(state string, options ...ElementHandleWaitForElementStateOptions) error
	// Returns element specified by selector when it satisfies `state` option. Returns `null` if waiting for `hidden` or
	// `detached`.
	// Wait for the `selector` relative to the element handle to satisfy `state` option (either appear/disappear from dom,
	// or become visible/hidden). If at the moment of calling the method `selector` already satisfies the condition, the
	// method will return immediately. If the selector doesn't satisfy the condition for the `timeout` milliseconds, the
	// function will throw.
	WaitForSelector(selector string, options ...ElementHandleWaitForSelectorOptions) (ElementHandle, error)
	// Returns `input.value` for the selected `<input>` or `<textarea>` or `<select>` element.
	// Throws for non-input elements. However, if the element is inside the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), returns the value of the
	// control.
	InputValue(options ...ElementHandleInputValueOptions) (string, error)
}

type EventEmitter interface {
	Emit(name string, payload ...interface{}) bool
	ListenerCount(name string) int
	On(name string, handler interface{})
	Once(name string, handler interface{})
	RemoveListener(name string, handler interface{})
}

// `FileChooser` objects are dispatched by the page in the [`event: Page.fileChooser`] event.
type FileChooser interface {
	// Returns input element associated with this file chooser.
	Element() ElementHandle
	// Returns whether this file chooser accepts multiple files.
	IsMultiple() bool
	// Returns page this file chooser belongs to.
	Page() Page
	// Sets the value of the file input this chooser is associated with. If some of the `filePaths` are relative paths,
	// then they are resolved relative to the current working directory. For empty array, clears the selected files.
	SetFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error
}

// At every point of time, page exposes its current frame tree via the Page.mainFrame() and
// Frame.childFrames() methods.
// `Frame` object's lifecycle is controlled by three events, dispatched on the page object:
// - [`event: Page.frameAttached`] - fired when the frame gets attached to the page. A Frame can be attached to the
// page only once.
// - [`event: Page.frameNavigated`] - fired when the frame commits navigation to a different URL.
// - [`event: Page.frameDetached`] - fired when the frame gets detached from the page.  A Frame can be detached from
// the page only once.
// An example of dumping frame tree:
type Frame interface {
	// This method checks or unchecks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Ensure that matched element is a checkbox or a radio input. If not, this method throws.
	// 1. If the element already has the right checked state, this method returns immediately.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now checked or unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	SetChecked(selector string, checked bool, options ...FrameSetCheckedOptions) error
	// Returns the added tag when the script's onload fires or when the script content was injected into frame.
	// Adds a `<script>` tag into the page with the desired url or content.
	AddScriptTag(options PageAddScriptTagOptions) (ElementHandle, error)
	// Returns the added tag when the stylesheet's onload fires or when the CSS content was injected into frame.
	// Adds a `<link rel="stylesheet">` tag into the page with the desired url or a `<style type="text/css">` tag with the
	// content.
	AddStyleTag(options PageAddStyleTagOptions) (ElementHandle, error)
	// This method checks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Ensure that matched element is a checkbox or a radio input. If not, this method throws. If the element is
	// already checked, this method returns immediately.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now checked. If not, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	Check(selector string, options ...FrameCheckOptions) error
	ChildFrames() []Frame
	// This method clicks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to click in the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	Click(selector string, options ...PageClickOptions) error
	// Gets the full HTML contents of the frame, including the doctype.
	Content() (string, error)
	// This method double clicks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to double click in the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set. Note that if
	// the first click of the `dblclick()` triggers a navigation event, this method will throw.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	// **NOTE** `frame.dblclick()` dispatches two `click` events and a single `dblclick` event.
	Dblclick(selector string, options ...FrameDblclickOptions) error
	// The snippet below dispatches the `click` event on the element. Regardless of the visibility state of the element,
	// `click` is dispatched. This is equivalent to calling
	// [element.click()](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click).
	DispatchEvent(selector, typ string, eventInit interface{}, options ...PageDispatchEventOptions) error
	// Returns the return value of `expression`.
	// If the function passed to the Frame.evaluate() returns a [Promise], then Frame.evaluate() would
	// wait for the promise to resolve and return its value.
	// If the function passed to the Frame.evaluate() returns a non-[Serializable] value, then
	// Frame.evaluate() returns `undefined`. Playwright also supports transferring some additional values that
	// are not serializable by `JSON`: `-0`, `NaN`, `Infinity`, `-Infinity`.
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `expression` as a `JSHandle`.
	// The only difference between Frame.evaluate() and Frame.evaluateHandle() is that
	// Frame.evaluateHandle() returns `JSHandle`.
	// If the function, passed to the Frame.evaluateHandle(), returns a [Promise], then
	// Frame.evaluateHandle() would wait for the promise to resolve and return its value.
	EvaluateHandle(expression string, options ...interface{}) (JSHandle, error)
	// Returns the return value of `expression`.
	// The method finds an element matching the specified selector within the frame and passes it as a first argument to
	// `expression`. If no elements match the selector, the method throws an error.
	// If `expression` returns a [Promise], then Frame.evalOnSelector() would wait for the promise to resolve
	// and return its value.
	EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `expression`.
	// The method finds all elements matching the specified selector within the frame and passes an array of matched
	// elements as a first argument to `expression`.
	// If `expression` returns a [Promise], then Frame.evalOnSelectorAll() would wait for the promise to resolve
	// and return its value.
	EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
	// This method waits for an element matching `selector`, waits for [actionability](../actionability.md) checks,
	// focuses the element, fills it and triggers an `input` event after filling. Note that you can pass an empty string
	// to clear the input field.
	// If the target element is not an `<input>`, `<textarea>` or `[contenteditable]` element, this method throws an
	// error. However, if the element is inside the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), the control will be filled
	// instead.
	// To send fine-grained keyboard events, use Frame.type().
	Fill(selector string, value string, options ...FrameFillOptions) error
	// This method fetches an element with `selector` and focuses it. If there's no element matching `selector`, the
	// method waits until a matching element appears in the DOM.
	Focus(selector string, options ...FrameFocusOptions) error
	// Returns the `frame` or `iframe` element handle which corresponds to this frame.
	// This is an inverse of ElementHandle.contentFrame(). Note that returned handle actually belongs to the
	// parent frame.
	// This method throws an error if the frame has been detached before `frameElement()` returns.
	FrameElement() (ElementHandle, error)
	// When working with iframes, you can create a frame locator that will enter the iframe and allow selecting elements
	// in that iframe.
	FrameLocator(selector string) FrameLocator
	// Returns element attribute value.
	GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error)
	// Allows locating elements by their alt text.
	GetByAltText(text interface{}, options ...LocatorGetByAltTextOptions) Locator
	// Allows locating input elements by the text of the associated `<label>` or `aria-labelledby` element, or by the
	// `aria-label` attribute.
	GetByLabel(text interface{}, options ...LocatorGetByLabelOptions) Locator
	// Allows locating input elements by the placeholder text.
	GetByPlaceholder(text interface{}, options ...LocatorGetByPlaceholderOptions) Locator
	// Allows locating elements by their [ARIA role](https://www.w3.org/TR/wai-aria-1.2/#roles),
	// [ARIA attributes](https://www.w3.org/TR/wai-aria-1.2/#aria-attributes) and
	// [accessible name](https://w3c.github.io/accname/#dfn-accessible-name).
	GetByRole(role AriaRole, options ...LocatorGetByRoleOptions) Locator
	// Locate element by the test id.
	GetByTestId(testId interface{}) Locator
	// Allows locating elements that contain given text.
	// See also Locator.filter() that allows to match by another criteria, like an accessible role, and then
	// filter by the text content.
	GetByText(text interface{}, options ...LocatorGetByTextOptions) Locator
	// Allows locating elements by their title attribute.
	GetByTitle(title interface{}, options ...LocatorGetByTitleOptions) Locator
	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of
	// the last redirect.
	// The method will throw an error if:
	// - there's an SSL error (e.g. in case of self-signed certificates).
	// - target URL is invalid.
	// - the `timeout` is exceeded during navigation.
	// - the remote server does not respond or is unreachable.
	// - the main resource failed to load.
	// The method will not throw an error when any valid HTTP status code is returned by the remote server, including 404
	// "Not Found" and 500 "Internal Server Error".  The status code for such responses can be retrieved by calling
	// Response.status().
	// **NOTE** The method either throws an error or returns a main resource response. The only exceptions are navigation
	// to `about:blank` or navigation to the same URL with a different hash, which would succeed and return `null`.
	// **NOTE** Headless mode doesn't support navigation to a PDF document. See the
	// [upstream issue](https://bugs.chromium.org/p/chromium/issues/detail?id=761295).
	Goto(url string, options ...PageGotoOptions) (Response, error)
	// This method hovers over an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to hover over the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	Hover(selector string, options ...PageHoverOptions) error
	// Returns `element.innerHTML`.
	InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error)
	// Returns `element.innerText`.
	InnerText(selector string, options ...PageInnerTextOptions) (string, error)
	// Returns `true` if the frame has been detached, or `false` otherwise.
	IsDetached() bool
	// Returns whether the element is checked. Throws if the element is not a checkbox or radio input.
	IsChecked(selector string, options ...FrameIsCheckedOptions) (bool, error)
	// Returns whether the element is disabled, the opposite of [enabled](../actionability.md#enabled).
	IsDisabled(selector string, options ...FrameIsDisabledOptions) (bool, error)
	// Returns whether the element is [editable](../actionability.md#editable).
	IsEditable(selector string, options ...FrameIsEditableOptions) (bool, error)
	// Returns whether the element is [enabled](../actionability.md#enabled).
	IsEnabled(selector string, options ...FrameIsEnabledOptions) (bool, error)
	// Returns whether the element is hidden, the opposite of [visible](../actionability.md#visible).  `selector` that
	// does not match any elements is considered hidden.
	IsHidden(selector string, options ...FrameIsHiddenOptions) (bool, error)
	// Returns whether the element is [visible](../actionability.md#visible). `selector` that does not match any elements
	// is considered not visible.
	IsVisible(selector string, options ...FrameIsVisibleOptions) (bool, error)
	// The method returns an element locator that can be used to perform actions on this page / frame. Locator is resolved
	// to the element immediately before performing an action, so a series of actions on the same locator can in fact be
	// performed on different DOM elements. That would happen if the DOM structure between those actions has changed.
	// [Learn more about locators](../locators.md).
	// [Learn more about locators](../locators.md).
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
	// `key` can specify the intended
	// [keyboardEvent.key](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key) value or a single character
	// to generate the text for. A superset of the `key` values can be found
	// [here](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values). Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`,
	// etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When specified with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	Press(selector, key string, options ...PagePressOptions) error
	// Returns the ElementHandle pointing to the frame element.
	// **NOTE** The use of `ElementHandle` is discouraged, use `Locator` objects and web-first assertions instead.
	// The method finds an element matching the specified selector within the frame. If no elements match the selector,
	// returns `null`.
	QuerySelector(selector string) (ElementHandle, error)
	// Returns the ElementHandles pointing to the frame elements.
	// **NOTE** The use of `ElementHandle` is discouraged, use `Locator` objects instead.
	// The method finds all elements matching the specified selector within the frame. If no elements match the selector,
	// returns empty array.
	QuerySelectorAll(selector string) ([]ElementHandle, error)
	SetContent(content string, options ...PageSetContentOptions) error
	// This method waits for an element matching `selector`, waits for [actionability](../actionability.md) checks, waits
	// until all specified options are present in the `<select>` element and selects these options.
	// If the target element is not a `<select>` element, this method throws an error. However, if the element is inside
	// the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), the control will be used
	// instead.
	// Returns the array of option values that have been successfully selected.
	// Triggers a `change` and `input` event once all the provided options have been selected.
	SelectOption(selector string, values SelectOptionValues, options ...FrameSelectOptionOptions) ([]string, error)
	// Sets the value of the file input to these file paths or files. If some of the `filePaths` are relative paths, then
	// they are resolved relative to the current working directory. For empty array, clears the selected files.
	// This method expects `selector` to point to an
	// [input element](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input). However, if the element is inside
	// the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), targets the control instead.
	SetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) error
	// This method taps an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.touchscreen() to tap the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	// **NOTE** `frame.tap()` requires that the `hasTouch` option of the browser context be set to true.
	Tap(selector string, options ...FrameTapOptions) error
	// Returns `element.textContent`.
	TextContent(selector string, options ...FrameTextContentOptions) (string, error)
	// Returns the page title.
	Title() (string, error)
	// Sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the text. `frame.type` can be used
	// to send fine-grained keyboard events. To fill values in form fields, use Frame.fill().
	// To press a special key, like `Control` or `ArrowDown`, use Keyboard.press().
	Type(selector, text string, options ...PageTypeOptions) error
	// Returns frame's url.
	URL() string
	// This method checks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Ensure that matched element is a checkbox or a radio input. If not, this method throws. If the element is
	// already unchecked, this method returns immediately.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	Uncheck(selector string, options ...FrameUncheckOptions) error
	WaitForEvent(event string, options ...PageWaitForEventOptions) (interface{}, error)
	// Returns when the `expression` returns a truthy value, returns that value.
	WaitForFunction(expression string, arg interface{}, options ...FrameWaitForFunctionOptions) (JSHandle, error)
	// Waits for the required load state to be reached.
	// This returns when the frame reaches a required load state, `load` by default. The navigation must have been
	// committed when this method is called. If current document has already reached the required state, resolves
	// immediately.
	WaitForLoadState(options ...PageWaitForLoadStateOptions) error
	// Waits for the frame navigation and returns the main resource response. In case of multiple redirects, the
	// navigation will resolve with the response of the last redirect. In case of navigation to a different anchor or
	// navigation due to History API usage, the navigation will resolve with `null`.
	WaitForNavigation(options ...PageWaitForNavigationOptions) (Response, error)
	// Waits for the frame to navigate to the given URL.
	WaitForURL(url string, options ...FrameWaitForURLOptions) error
	// Returns when element specified by selector satisfies `state` option. Returns `null` if waiting for `hidden` or
	// `detached`.
	// **NOTE** Playwright automatically waits for element to be ready before performing an action. Using `Locator`
	// objects and web-first assertions make the code wait-for-selector-free.
	// Wait for the `selector` to satisfy `state` option (either appear/disappear from dom, or become visible/hidden). If
	// at the moment of calling the method `selector` already satisfies the condition, the method will return immediately.
	// If the selector doesn't satisfy the condition for the `timeout` milliseconds, the function will throw.
	WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandle, error)
	// Waits for the given `timeout` in milliseconds.
	// Note that `frame.waitForTimeout()` should only be used for debugging. Tests using the timer in production are going
	// to be flaky. Use signals such as network events, selectors becoming visible and others instead.
	WaitForTimeout(timeout float64)
	// Returns `input.value` for the selected `<input>` or `<textarea>` or `<select>` element.
	// Throws for non-input elements. However, if the element is inside the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), returns the value of the
	// control.
	InputValue(selector string, options ...FrameInputValueOptions) (string, error)
	DragAndDrop(source, target string, options ...FrameDragAndDropOptions) error
}

// FrameLocator represents a view to the `iframe` on the page. It captures the logic sufficient to retrieve the
// `iframe` and locate elements in that iframe. FrameLocator can be created with either Page.frameLocator()
// or Locator.frameLocator() method.
// **Strictness**
// Frame locators are strict. This means that all operations on frame locators will throw if more than one element
// matches a given selector.
// **Converting Locator to FrameLocator**
// If you have a `Locator` object pointing to an `iframe` it can be converted to `FrameLocator` using
// [`:scope`](https://developer.mozilla.org/en-US/docs/Web/CSS/:scope) CSS selector:
type FrameLocator interface {
	// Returns locator to the first matching frame.
	First() FrameLocator
	// When working with iframes, you can create a frame locator that will enter the iframe and allow selecting elements
	// in that iframe.
	FrameLocator(selector string) FrameLocator
	// Returns locator to the last matching frame.
	Last() FrameLocator
	// The method finds an element matching the specified selector in the locator's subtree. It also accepts filter
	// options, similar to Locator.filter() method.
	// [Learn more about locators](../locators.md).
	Locator(selector string, options ...FrameLocatorOptions) Locator
	// Allows locating elements by their alt text.
	GetByAltText(text interface{}, options ...LocatorGetByAltTextOptions) Locator
	// Allows locating input elements by the text of the associated `<label>` or `aria-labelledby` element, or by the
	// `aria-label` attribute.
	GetByLabel(text interface{}, options ...LocatorGetByLabelOptions) Locator
	// Allows locating input elements by the placeholder text.
	GetByPlaceholder(text interface{}, options ...LocatorGetByPlaceholderOptions) Locator
	// Allows locating elements by their [ARIA role](https://www.w3.org/TR/wai-aria-1.2/#roles),
	// [ARIA attributes](https://www.w3.org/TR/wai-aria-1.2/#aria-attributes) and
	// [accessible name](https://w3c.github.io/accname/#dfn-accessible-name).
	GetByRole(role AriaRole, options ...LocatorGetByRoleOptions) Locator
	// Locate element by the test id.
	GetByTestId(testId interface{}) Locator
	// Allows locating elements that contain given text.
	// See also Locator.filter() that allows to match by another criteria, like an accessible role, and then
	// filter by the text content.
	GetByText(text interface{}, options ...LocatorGetByTextOptions) Locator
	// Allows locating elements by their title attribute.
	GetByTitle(title interface{}, options ...LocatorGetByTitleOptions) Locator
	// Returns locator to the n-th matching frame. It's zero based, `nth(0)` selects the first frame.
	Nth(index int) FrameLocator
}

// JSHandle represents an in-page JavaScript object. JSHandles can be created with the Page.evaluateHandle()
// method.
// JSHandle prevents the referenced JavaScript object being garbage collected unless the handle is exposed with
// JSHandle.dispose(). JSHandles are auto-disposed when their origin frame gets navigated or the parent
// context gets destroyed.
// JSHandle instances can be used as an argument in Page.evalOnSelector(), Page.evaluate() and
// Page.evaluateHandle() methods.
type JSHandle interface {
	// Returns either `null` or the object handle itself, if the object handle is an instance of `ElementHandle`.
	AsElement() ElementHandle
	// The `jsHandle.dispose` method stops referencing the element handle.
	Dispose() error
	// Returns the return value of `expression`.
	// This method passes this handle as the first argument to `expression`.
	// If `expression` returns a [Promise], then `handle.evaluate` would wait for the promise to resolve and return its
	// value.
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `expression` as a `JSHandle`.
	// This method passes this handle as the first argument to `expression`.
	// The only difference between `jsHandle.evaluate` and `jsHandle.evaluateHandle` is that `jsHandle.evaluateHandle`
	// returns `JSHandle`.
	// If the function passed to the `jsHandle.evaluateHandle` returns a [Promise], then `jsHandle.evaluateHandle` would
	// wait for the promise to resolve and return its value.
	// See Page.evaluateHandle() for more details.
	EvaluateHandle(expression string, options ...interface{}) (JSHandle, error)
	// The method returns a map with **own property names** as keys and JSHandle instances for the property values.
	GetProperties() (map[string]JSHandle, error)
	// Fetches a single property from the referenced object.
	GetProperty(name string) (JSHandle, error)
	// Returns a JSON representation of the object. If the object has a `toJSON` function, it **will not be called**.
	// **NOTE** The method will return an empty JSON object if the referenced object is not stringifiable. It will throw
	// an error if the object has circular references.
	JSONValue() (interface{}, error)
	String() string
}

// Keyboard provides an api for managing a virtual keyboard. The high level api is Keyboard.type(), which
// takes raw characters and generates proper `keydown`, `keypress`/`input`, and `keyup` events on your page.
// For finer control, you can use Keyboard.down(), Keyboard.up(), and
// Keyboard.insertText() to manually fire events as if they were generated from a real keyboard.
// An example of holding down `Shift` in order to select and delete some text:
// An example of pressing uppercase `A`
// An example to trigger select-all with the keyboard
type Keyboard interface {
	// Dispatches a `keydown` event.
	// `key` can specify the intended
	// [keyboardEvent.key](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key) value or a single character
	// to generate the text for. A superset of the `key` values can be found
	// [here](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values). Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`,
	// etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// If `key` is a modifier key, `Shift`, `Meta`, `Control`, or `Alt`, subsequent key presses will be sent with that
	// modifier active. To release the modifier key, use Keyboard.up().
	// After the key is pressed once, subsequent calls to Keyboard.down() will have
	// [repeat](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/repeat) set to true. To release the key,
	// use Keyboard.up().
	// **NOTE** Modifier keys DO influence `keyboard.down`. Holding down `Shift` will type the text in upper case.
	Down(key string) error
	// Dispatches only `input` event, does not emit the `keydown`, `keyup` or `keypress` events.
	InsertText(text string) error
	// `key` can specify the intended
	// [keyboardEvent.key](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key) value or a single character
	// to generate the text for. A superset of the `key` values can be found
	// [here](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values). Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`,
	// etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When specified with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	Press(key string, options ...KeyboardPressOptions) error
	// Sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the text.
	// To press a special key, like `Control` or `ArrowDown`, use Keyboard.press().
	Type(text string, options ...KeyboardTypeOptions) error
	// Dispatches a `keyup` event.
	Up(key string) error
}

// Locators are the central piece of Playwright's auto-waiting and retry-ability. In a nutshell, locators represent a
// way to find element(s) on the page at any moment. Locator can be created with the Page.locator() method.
// [Learn more about locators](../locators.md).
type Locator interface {
	// When locator points to a list of elements, returns array of locators, pointing to respective elements.
	// **NOTE** Locator.all() does not wait for elements to match the locator, and instead immediately returns
	// whatever is present in the page.  When the list of elements changes dynamically, Locator.all() will
	// produce unpredictable and flaky results.  When the list of elements is stable, but loaded dynamically, wait for the
	// full list to finish loading before calling Locator.all().
	All() ([]Locator, error)
	// Returns an array of `node.innerText` values for all matching nodes.
	AllInnerTexts() ([]string, error)
	// Returns an array of `node.textContent` values for all matching nodes.
	AllTextContents() ([]string, error)
	// Creates a locator that matches both this locator and the argument locator.
	And(locator Locator) Locator
	// Calls [blur](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/blur) on the element.
	Blur(options ...LocatorBlurOptions) error
	// This method returns the bounding box of the element matching the locator, or `null` if the element is not visible.
	// The bounding box is calculated relative to the main frame viewport - which is usually the same as the browser
	// window.
	// **Details**
	// Scrolling affects the returned bounding box, similarly to
	// [Element.getBoundingClientRect](https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect).
	// That means `x` and/or `y` may be negative.
	// Elements from child frames return the bounding box relative to the main frame, unlike the
	// [Element.getBoundingClientRect](https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect).
	// Assuming the page is static, it is safe to use bounding box coordinates to perform input. For example, the
	// following snippet should click the center of the element.
	BoundingBox(options ...LocatorBoundingBoxOptions) (*Rect, error)
	// Ensure that checkbox or radio element is checked.
	// **Details**
	// Performs the following steps:
	// 1. Ensure that element is a checkbox or a radio input. If not, this method throws. If the element is already
	// checked, this method returns immediately.
	// 1. Wait for [actionability](../actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now checked. If not, this method throws.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	Check(options ...FrameCheckOptions) error
	// Clear the input field.
	// **Details**
	// This method waits for [actionability](../actionability.md) checks, focuses the element, clears it and triggers an
	// `input` event after clearing.
	// If the target element is not an `<input>`, `<textarea>` or `[contenteditable]` element, this method throws an
	// error. However, if the element is inside the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), the control will be cleared
	// instead.
	Clear(options ...LocatorClearOptions) error
	// Click an element.
	// **Details**
	// This method clicks the element by performing the following steps:
	// 1. Wait for [actionability](../actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to click in the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	Click(options ...PageClickOptions) error
	// Returns the number of elements matching the locator.
	Count() (int, error)
	// Double-click an element.
	// **Details**
	// This method double clicks the element by performing the following steps:
	// 1. Wait for [actionability](../actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to double click in the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set. Note that if
	// the first click of the `dblclick()` triggers a navigation event, this method will throw.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	// **NOTE** `element.dblclick()` dispatches two `click` events and a single `dblclick` event.
	Dblclick(options ...FrameDblclickOptions) error
	// Programmatically dispatch an event on the matching element.
	DispatchEvent(typ string, eventInit interface{}, options ...PageDispatchEventOptions) error
	// Drag the source element towards the target element and drop it.
	// **Details**
	// This method drags the locator to another target locator or target position. It will first move to the source
	// element, perform a `mousedown`, then move to the target element or position and perform a `mouseup`.
	DragTo(target Locator, options ...FrameDragAndDropOptions) error
	// Resolves given locator to the first matching DOM element. If there are no matching elements, waits for one. If
	// multiple elements match the locator, throws.
	ElementHandle(options ...LocatorElementHandleOptions) (ElementHandle, error)
	// Resolves given locator to all matching DOM elements. If there are no matching elements, returns an empty list.
	ElementHandles() ([]ElementHandle, error)
	Err() error
	// Execute JavaScript code in the page, taking the matching element as an argument.
	// **Details**
	// Returns the return value of `expression`, called with the matching element as a first argument, and `arg` as a
	// second argument.
	// If `expression` returns a [Promise], this method will wait for the promise to resolve and return its value.
	// If `expression` throws or rejects, this method throws.
	Evaluate(expression string, arg interface{}, options ...LocatorEvaluateOptions) (interface{}, error)
	// Execute JavaScript code in the page, taking all matching elements as an argument.
	// **Details**
	// Returns the return value of `expression`, called with an array of all matching elements as a first argument, and
	// `arg` as a second argument.
	// If `expression` returns a [Promise], this method will wait for the promise to resolve and return its value.
	// If `expression` throws or rejects, this method throws.
	EvaluateAll(expression string, options ...interface{}) (interface{}, error)
	// Execute JavaScript code in the page, taking the matching element as an argument, and return a `JSHandle` with the
	// result.
	// **Details**
	// Returns the return value of `expression` as a`JSHandle`, called with the matching element as a first argument, and
	// `arg` as a second argument.
	// The only difference between Locator.evaluate() and Locator.evaluateHandle() is that
	// Locator.evaluateHandle() returns `JSHandle`.
	// If `expression` returns a [Promise], this method will wait for the promise to resolve and return its value.
	// If `expression` throws or rejects, this method throws.
	// See Page.evaluateHandle() for more details.
	EvaluateHandle(expression string, arg interface{}, options ...LocatorEvaluateHandleOptions) (interface{}, error)
	// Set a value to the input field.
	Fill(value string, options ...FrameFillOptions) error
	// This method narrows existing locator according to the options, for example filters by text. It can be chained to
	// filter multiple times.
	Filter(options ...LocatorLocatorOptions) Locator
	// Returns locator to the first matching element.
	First() Locator
	// Calls [focus](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus) on the matching element.
	Focus(options ...FrameFocusOptions) error
	// When working with iframes, you can create a frame locator that will enter the iframe and allow locating elements in
	// that iframe:
	FrameLocator(selector string) FrameLocator
	// Returns the matching element's attribute value.
	GetAttribute(name string, options ...PageGetAttributeOptions) (string, error)
	// Allows locating elements by their alt text.
	GetByAltText(text interface{}, options ...LocatorGetByAltTextOptions) Locator
	// Allows locating input elements by the text of the associated `<label>` or `aria-labelledby` element, or by the
	// `aria-label` attribute.
	GetByLabel(text interface{}, options ...LocatorGetByLabelOptions) Locator
	// Allows locating input elements by the placeholder text.
	GetByPlaceholder(text interface{}, options ...LocatorGetByPlaceholderOptions) Locator
	// Allows locating elements by their [ARIA role](https://www.w3.org/TR/wai-aria-1.2/#roles),
	// [ARIA attributes](https://www.w3.org/TR/wai-aria-1.2/#aria-attributes) and
	// [accessible name](https://w3c.github.io/accname/#dfn-accessible-name).
	GetByRole(role AriaRole, options ...LocatorGetByRoleOptions) Locator
	// Locate element by the test id.
	GetByTestId(testId interface{}) Locator
	// Allows locating elements that contain given text.
	// See also Locator.filter() that allows to match by another criteria, like an accessible role, and then
	// filter by the text content.
	GetByText(text interface{}, options ...LocatorGetByTextOptions) Locator
	// Allows locating elements by their title attribute.
	GetByTitle(title interface{}, options ...LocatorGetByTitleOptions) Locator
	// Highlight the corresponding element(s) on the screen. Useful for debugging, don't commit the code that uses
	// Locator.highlight().
	Highlight() error
	// Hover over the matching element.
	Hover(options ...PageHoverOptions) error
	// Returns the [`element.innerHTML`](https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML).
	InnerHTML(options ...PageInnerHTMLOptions) (string, error)
	// Returns the [`element.innerText`](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText).
	InnerText(options ...PageInnerTextOptions) (string, error)
	// Returns the value for the matching `<input>` or `<textarea>` or `<select>` element.
	InputValue(options ...FrameInputValueOptions) (string, error)
	// Returns whether the element is checked. Throws if the element is not a checkbox or radio input.
	IsChecked(options ...FrameIsCheckedOptions) (bool, error)
	// Returns whether the element is disabled, the opposite of [enabled](../actionability.md#enabled).
	IsDisabled(options ...FrameIsDisabledOptions) (bool, error)
	// Returns whether the element is [editable](../actionability.md#editable).
	IsEditable(options ...FrameIsEditableOptions) (bool, error)
	// Returns whether the element is [enabled](../actionability.md#enabled).
	IsEnabled(options ...FrameIsEnabledOptions) (bool, error)
	// Returns whether the element is hidden, the opposite of [visible](../actionability.md#visible).
	IsHidden(options ...FrameIsHiddenOptions) (bool, error)
	// Returns whether the element is [visible](../actionability.md#visible).
	IsVisible(options ...FrameIsVisibleOptions) (bool, error)
	// Returns locator to the last matching element.
	Last() Locator
	// The method finds an element matching the specified selector in the locator's subtree. It also accepts filter
	// options, similar to Locator.filter() method.
	// [Learn more about locators](../locators.md).
	Locator(selector string, options ...LocatorLocatorOptions) Locator
	// Returns locator to the n-th matching element. It's zero based, `nth(0)` selects the first element.
	Nth(index int) Locator
	// Creates a locator that matches either of the two locators.
	Or(locator Locator) Locator
	// A page this locator belongs to.
	Page() (Page, error)
	// Focuses the matching element and presses a combination of the keys.
	Press(key string, options ...PagePressOptions) error
	// Take a screenshot of the element matching the locator.
	Screenshot(options ...LocatorScreenshotOptions) ([]byte, error)
	// This method waits for [actionability](../actionability.md) checks, then tries to scroll element into view, unless
	// it is completely visible as defined by
	// [IntersectionObserver](https://developer.mozilla.org/en-US/docs/Web/API/Intersection_Observer_API)'s `ratio`.
	ScrollIntoViewIfNeeded(options ...LocatorScrollIntoViewIfNeededOptions) error
	// Selects option or options in `<select>`.
	// **Details**
	// This method waits for [actionability](../actionability.md) checks, waits until all specified options are present in
	// the `<select>` element and selects these options.
	// If the target element is not a `<select>` element, this method throws an error. However, if the element is inside
	// the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), the control will be used
	// instead.
	// Returns the array of option values that have been successfully selected.
	// Triggers a `change` and `input` event once all the provided options have been selected.
	SelectOption(values SelectOptionValues, options ...FrameSelectOptionOptions) ([]string, error)
	// This method waits for [actionability](../actionability.md) checks, then focuses the element and selects all its
	// text content.
	// If the element is inside the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), focuses and selects text in
	// the control instead.
	SelectText(options ...LocatorSelectTextOptions) error
	// Set the state of a checkbox or a radio element.
	SetChecked(checked bool, options ...FrameSetCheckedOptions) error
	// Upload file or multiple files into `<input type=file>`.
	SetInputFiles(files []InputFile, options ...FrameSetInputFilesOptions) error
	// Perform a tap gesture on the element matching the locator.
	// **Details**
	// This method taps the element by performing the following steps:
	// 1. Wait for [actionability](../actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.touchscreen() to tap the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	// **NOTE** `element.tap()` requires that the `hasTouch` option of the browser context be set to true.
	Tap(options ...FrameTapOptions) error
	// Returns the [`node.textContent`](https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent).
	TextContent(options ...FrameTextContentOptions) (string, error)
	// Focuses the element, and then sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the
	// text.
	// To press a special key, like `Control` or `ArrowDown`, use Locator.press().
	Type(text string, options ...PageTypeOptions) error
	// Ensure that checkbox or radio element is unchecked.
	Uncheck(options ...FrameUncheckOptions) error
	// Returns when element specified by locator satisfies the `state` option.
	// If target element already satisfies the condition, the method returns immediately. Otherwise, waits for up to
	// `timeout` milliseconds until the condition is met.
	WaitFor(options ...PageWaitForSelectorOptions) error
}

// The `LocatorAssertions` class provides assertion methods that can be used to make assertions about the `Locator`
// state in the tests.
type LocatorAssertions interface {
	// Makes the assertion check for the opposite condition. For example, this code tests that the Locator doesn't contain
	// text `"error"`:
	Not() LocatorAssertions
	// Ensures that `Locator` points to an [attached](../actionability.md#attached) DOM node.
	ToBeAttached(options ...LocatorAssertionsToBeAttachedOptions) error
	// Ensures the `Locator` points to a checked input.
	ToBeChecked(options ...LocatorAssertionsToBeCheckedOptions) error
	// Ensures the `Locator` points to a disabled element. Element is disabled if it has "disabled" attribute or is
	// disabled via
	// ['aria-disabled'](https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Attributes/aria-disabled). Note
	// that only native control elements such as HTML `button`, `input`, `select`, `textarea`, `option`, `optgroup` can be
	// disabled by setting "disabled" attribute. "disabled" attribute on other elements is ignored by the browser.
	ToBeDisabled(options ...LocatorAssertionsToBeDisabledOptions) error
	// Ensures the `Locator` points to an editable element.
	ToBeEditable(options ...LocatorAssertionsToBeEditableOptions) error
	// Ensures the `Locator` points to an empty editable element or to a DOM node that has no text.
	ToBeEmpty(options ...LocatorAssertionsToBeEmptyOptions) error
	// Ensures the `Locator` points to an enabled element.
	ToBeEnabled(options ...LocatorAssertionsToBeEnabledOptions) error
	// Ensures the `Locator` points to a focused DOM node.
	ToBeFocused(options ...LocatorAssertionsToBeFocusedOptions) error
	// Ensures that `Locator` either does not resolve to any DOM node, or resolves to a
	// [non-visible](../actionability.md#visible) one.
	ToBeHidden(options ...LocatorAssertionsToBeHiddenOptions) error
	// Ensures the `Locator` points to an element that intersects viewport, according to the
	// [intersection observer API](https://developer.mozilla.org/en-US/docs/Web/API/Intersection_Observer_API).
	ToBeInViewport(options ...LocatorAssertionsToBeInViewportOptions) error
	// Ensures that `Locator` points to an [attached](../actionability.md#attached) and
	// [visible](../actionability.md#visible) DOM node.
	ToBeVisible(options ...LocatorAssertionsToBeVisibleOptions) error
	// Ensures the `Locator` points to an element that contains the given text. You can use regular expressions for the
	// value as well.
	ToContainText(expected interface{}, options ...LocatorAssertionsToContainTextOptions) error
	// Ensures the `Locator` points to an element with given attribute.
	ToHaveAttribute(name string, value interface{}, options ...LocatorAssertionsToHaveAttributeOptions) error
	// Ensures the `Locator` points to an element with given CSS classes. This needs to be a full match or using a relaxed
	// regular expression.
	ToHaveClass(expected interface{}, options ...LocatorAssertionsToHaveClassOptions) error
	// Ensures the `Locator` resolves to an exact number of DOM nodes.
	ToHaveCount(count int, options ...LocatorAssertionsToHaveCountOptions) error
	// Ensures the `Locator` resolves to an element with the given computed CSS style.
	ToHaveCSS(name string, value interface{}, options ...LocatorAssertionsToHaveCSSOptions) error
	// Ensures the `Locator` points to an element with the given DOM Node ID.
	ToHaveId(id interface{}, options ...LocatorAssertionsToHaveIdOptions) error
	// Ensures the `Locator` points to an element with given JavaScript property. Note that this property can be of a
	// primitive type as well as a plain serializable JavaScript object.
	ToHaveJSProperty(name string, value interface{}, options ...LocatorAssertionsToHaveJSPropertyOptions) error
	// Ensures the `Locator` points to an element with the given text. You can use regular expressions for the value as
	// well.
	ToHaveText(expected interface{}, options ...LocatorAssertionsToHaveTextOptions) error
	// Ensures the `Locator` points to an element with the given input value. You can use regular expressions for the
	// value as well.
	ToHaveValue(value interface{}, options ...LocatorAssertionsToHaveValueOptions) error
	// Ensures the `Locator` points to multi-select/combobox (i.e. a `select` with the `multiple` attribute) and the
	// specified values are selected.
	ToHaveValues(values []interface{}, options ...LocatorAssertionsToHaveValuesOptions) error
	// The opposite of LocatorAssertions.toBeAttached().
	NotToBeAttached(options ...LocatorAssertionsToBeAttachedOptions) error
	// The opposite of LocatorAssertions.toBeChecked().
	NotToBeChecked(options ...LocatorAssertionsToBeCheckedOptions) error
	// The opposite of LocatorAssertions.toBeDisabled().
	NotToBeDisabled(options ...LocatorAssertionsToBeDisabledOptions) error
	// The opposite of LocatorAssertions.toBeEditable().
	NotToBeEditable(options ...LocatorAssertionsToBeEditableOptions) error
	// The opposite of LocatorAssertions.toBeEmpty().
	NotToBeEmpty(options ...LocatorAssertionsToBeEmptyOptions) error
	// The opposite of LocatorAssertions.toBeEnabled().
	NotToBeEnabled(options ...LocatorAssertionsToBeEnabledOptions) error
	// The opposite of LocatorAssertions.toBeFocused().
	NotToBeFocused(options ...LocatorAssertionsToBeFocusedOptions) error
	// The opposite of LocatorAssertions.toBeHidden().
	NotToBeHidden(options ...LocatorAssertionsToBeHiddenOptions) error
	// The opposite of LocatorAssertions.toBeInViewport().
	NotToBeInViewport(options ...LocatorAssertionsToBeInViewportOptions) error
	// The opposite of LocatorAssertions.toBeVisible().
	NotToBeVisible(options ...LocatorAssertionsToBeVisibleOptions) error
	// The opposite of LocatorAssertions.toContainText().
	NotToContainText(expected interface{}, options ...LocatorAssertionsToContainTextOptions) error
	// The opposite of LocatorAssertions.toHaveAttribute().
	NotToHaveAttribute(name string, value interface{}, options ...LocatorAssertionsToHaveAttributeOptions) error
	// The opposite of LocatorAssertions.toHaveClass().
	NotToHaveClass(expected interface{}, options ...LocatorAssertionsToHaveClassOptions) error
	// The opposite of LocatorAssertions.toHaveCount().
	NotToHaveCount(count int, options ...LocatorAssertionsToHaveCountOptions) error
	// The opposite of LocatorAssertions.toHaveCSS().
	NotToHaveCSS(name string, value interface{}, options ...LocatorAssertionsToHaveCSSOptions) error
	// The opposite of LocatorAssertions.toHaveId().
	NotToHaveId(id interface{}, options ...LocatorAssertionsToHaveIdOptions) error
	// The opposite of LocatorAssertions.toHaveJSProperty().
	NotToHaveJSProperty(name string, value interface{}, options ...LocatorAssertionsToHaveJSPropertyOptions) error
	// The opposite of LocatorAssertions.toHaveText().
	NotToHaveText(expected interface{}, options ...LocatorAssertionsToHaveTextOptions) error
	// The opposite of LocatorAssertions.toHaveValue().
	NotToHaveValue(value interface{}, options ...LocatorAssertionsToHaveValueOptions) error
	// The opposite of LocatorAssertions.toHaveValues().
	NotToHaveValues(values []interface{}, options ...LocatorAssertionsToHaveValuesOptions) error
}

// The Mouse class operates in main-frame CSS pixels relative to the top-left corner of the viewport.
// Every `page` object has its own Mouse, accessible with Page.mouse().
type Mouse interface {
	// Shortcut for Mouse.move(), Mouse.down(), Mouse.up().
	Click(x, y float64, options ...MouseClickOptions) error
	// Shortcut for Mouse.move(), Mouse.down(), Mouse.up(), Mouse.down() and
	// Mouse.up().
	Dblclick(x, y float64, options ...MouseDblclickOptions) error
	// Dispatches a `mousedown` event.
	Down(options ...MouseDownOptions) error
	// Dispatches a `mousemove` event.
	Move(x float64, y float64, options ...MouseMoveOptions) error
	// Dispatches a `mouseup` event.
	Up(options ...MouseUpOptions) error
	// Dispatches a `wheel` event.
	// **NOTE** Wheel events may cause scrolling if they are not handled, and this method does not wait for the scrolling
	// to finish before returning.
	Wheel(deltaX, deltaY float64) error
}

// Page provides methods to interact with a single tab in a `Browser`, or an
// [extension background page](https://developer.chrome.com/extensions/background_pages) in Chromium. One `Browser`
// instance might have multiple `Page` instances.
// This example creates a page, navigates it to a URL, and then saves a screenshot:
// The Page class emits various events (described below) which can be handled using any of Node's native
// [`EventEmitter`](https://nodejs.org/api/events.html#events_class_eventemitter) methods, such as `on`, `once` or
// `removeListener`.
// This example logs a message for a single page `load` event:
// To unsubscribe from events use the `removeListener` method:
type Page interface {
	EventEmitter
	// This method checks or unchecks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Ensure that matched element is a checkbox or a radio input. If not, this method throws.
	// 1. If the element already has the right checked state, this method returns immediately.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now checked or unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	SetChecked(selector string, checked bool, options ...FrameSetCheckedOptions) error
	Mouse() Mouse
	Keyboard() Keyboard
	Touchscreen() Touchscreen
	// Adds a script which would be evaluated in one of the following scenarios:
	// - Whenever the page is navigated.
	// - Whenever the child frame is attached or navigated. In this case, the script is evaluated in the context of the
	// newly attached frame.
	// The script is evaluated after the document was created but before any of its scripts were run. This is useful to
	// amend the JavaScript environment, e.g. to seed `Math.random`.
	AddInitScript(script PageAddInitScriptOptions) error
	// Adds a `<script>` tag into the page with the desired url or content. Returns the added tag when the script's onload
	// fires or when the script content was injected into frame.
	AddScriptTag(options PageAddScriptTagOptions) (ElementHandle, error)
	// Adds a `<link rel="stylesheet">` tag into the page with the desired url or a `<style type="text/css">` tag with the
	// content. Returns the added tag when the stylesheet's onload fires or when the CSS content was injected into frame.
	AddStyleTag(options PageAddStyleTagOptions) (ElementHandle, error)
	// Brings page to front (activates tab).
	BringToFront() error
	// This method checks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Ensure that matched element is a checkbox or a radio input. If not, this method throws. If the element is
	// already checked, this method returns immediately.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now checked. If not, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	Check(selector string, options ...FrameCheckOptions) error
	// This method clicks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to click in the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	Click(selector string, options ...PageClickOptions) error
	// If `runBeforeUnload` is `false`, does not run any unload handlers and waits for the page to be closed. If
	// `runBeforeUnload` is `true` the method will run unload handlers, but will **not** wait for the page to close.
	// By default, `page.close()` **does not** run `beforeunload` handlers.
	// **NOTE** if `runBeforeUnload` is passed as true, a `beforeunload` dialog might be summoned and should be handled
	// manually via [`event: Page.dialog`] event.
	Close(options ...PageCloseOptions) error
	// Gets the full HTML contents of the page, including the doctype.
	Content() (string, error)
	// Get the browser context that the page belongs to.
	Context() BrowserContext
	// This method double clicks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to double click in the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set. Note that if
	// the first click of the `dblclick()` triggers a navigation event, this method will throw.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	// **NOTE** `page.dblclick()` dispatches two `click` events and a single `dblclick` event.
	Dblclick(expression string, options ...FrameDblclickOptions) error
	// The snippet below dispatches the `click` event on the element. Regardless of the visibility state of the element,
	// `click` is dispatched. This is equivalent to calling
	// [element.click()](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click).
	DispatchEvent(selector string, typ string, options ...PageDispatchEventOptions) error
	// The method adds a function called `name` on the `window` object of every frame in this page. When called, the
	// function executes `callback` and returns a [Promise] which resolves to the return value of `callback`. If the
	// `callback` returns a [Promise], it will be awaited.
	// The first argument of the `callback` function contains information about the caller: `{ browserContext:
	// BrowserContext, page: Page, frame: Frame }`.
	// See BrowserContext.exposeBinding() for the context-wide version.
	// **NOTE** Functions installed via Page.exposeBinding() survive navigations.
	ExposeBinding(name string, binding BindingCallFunction, handle ...bool) error
	// The method adds a function called `name` on the `window` object of every frame in the page. When called, the
	// function executes `callback` and returns a [Promise] which resolves to the return value of `callback`.
	// If the `callback` returns a [Promise], it will be awaited.
	// See BrowserContext.exposeFunction() for context-wide exposed function.
	// **NOTE** Functions installed via Page.exposeFunction() survive navigations.
	ExposeFunction(name string, binding ExposedFunction) error
	// This method changes the `CSS media type` through the `media` argument, and/or the `'prefers-colors-scheme'` media
	// feature, using the `colorScheme` argument.
	EmulateMedia(options ...PageEmulateMediaOptions) error
	// Returns the value of the `expression` invocation.
	// If the function passed to the Page.evaluate() returns a [Promise], then Page.evaluate() would
	// wait for the promise to resolve and return its value.
	// If the function passed to the Page.evaluate() returns a non-[Serializable] value, then
	// Page.evaluate() resolves to `undefined`. Playwright also supports transferring some additional values
	// that are not serializable by `JSON`: `-0`, `NaN`, `Infinity`, `-Infinity`.
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	// Returns the value of the `expression` invocation as a `JSHandle`.
	// The only difference between Page.evaluate() and Page.evaluateHandle() is that
	// Page.evaluateHandle() returns `JSHandle`.
	// If the function passed to the Page.evaluateHandle() returns a [Promise], then
	// Page.evaluateHandle() would wait for the promise to resolve and return its value.
	EvaluateHandle(expression string, options ...interface{}) (JSHandle, error)
	// The method finds an element matching the specified selector within the page and passes it as a first argument to
	// `expression`. If no elements match the selector, the method throws an error. Returns the value of `expression`.
	// If `expression` returns a [Promise], then Page.evalOnSelector() would wait for the promise to resolve and
	// return its value.
	EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
	// The method finds all elements matching the specified selector within the page and passes an array of matched
	// elements as a first argument to `expression`. Returns the result of `expression` invocation.
	// If `expression` returns a [Promise], then Page.evalOnSelectorAll() would wait for the promise to resolve
	// and return its value.
	EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
	// Performs action and waits for a `ConsoleMessage` to be logged by in the page. If predicate is provided, it passes
	// `ConsoleMessage` value into the `predicate` function and waits for `predicate(message)` to return a truthy value.
	// Will throw an error if the page is closed before the [`event: Page.console`] event is fired.
	ExpectConsoleMessage(cb func() error, options ...PageExpectConsoleMessageOptions) (ConsoleMessage, error)
	// Performs action and waits for a new `Download`. If predicate is provided, it passes `Download` value into the
	// `predicate` function and waits for `predicate(download)` to return a truthy value. Will throw an error if the page
	// is closed before the download event is fired.
	ExpectDownload(cb func() error, options ...PageExpectDownloadOptions) (Download, error)
	// Waits for event to fire and passes its value into the predicate function. Returns when the predicate returns truthy
	// value. Will throw an error if the page is closed before the event is fired. Returns the event data value.
	ExpectEvent(event string, cb func() error, options ...PageWaitForEventOptions) (interface{}, error)
	// Performs action and waits for a new `FileChooser` to be created. If predicate is provided, it passes `FileChooser`
	// value into the `predicate` function and waits for `predicate(fileChooser)` to return a truthy value. Will throw an
	// error if the page is closed before the file chooser is opened.
	ExpectFileChooser(cb func() error, options ...PageExpectFileChooserOptions) (FileChooser, error)
	ExpectLoadState(cb func() error, options ...PageWaitForLoadStateOptions) error
	ExpectNavigation(cb func() error, options ...PageWaitForNavigationOptions) (Response, error)
	// Performs action and waits for a popup `Page`. If predicate is provided, it passes [Popup] value into the
	// `predicate` function and waits for `predicate(page)` to return a truthy value. Will throw an error if the page is
	// closed before the popup event is fired.
	ExpectPopup(cb func() error, options ...PageExpectPopupOptions) (Page, error)
	ExpectRequest(url interface{}, cb func() error, options ...PageWaitForRequestOptions) (Request, error)
	// Performs action and waits for a `Request` to finish loading. If predicate is provided, it passes `Request` value
	// into the `predicate` function and waits for `predicate(request)` to return a truthy value. Will throw an error if
	// the page is closed before the [`event: Page.requestFinished`] event is fired.
	ExpectRequestFinished(cb func() error, options ...PageExpectRequestFinishedOptions) (Request, error)
	ExpectResponse(url interface{}, cb func() error, options ...PageWaitForResponseOptions) (Response, error)
	// Performs action and waits for a new `WebSocket`. If predicate is provided, it passes `WebSocket` value into the
	// `predicate` function and waits for `predicate(webSocket)` to return a truthy value. Will throw an error if the page
	// is closed before the WebSocket event is fired.
	ExpectWebSocket(cb func() error, options ...PageExpectWebSocketOptions) (WebSocket, error)
	// Performs action and waits for a new `Worker`. If predicate is provided, it passes `Worker` value into the
	// `predicate` function and waits for `predicate(worker)` to return a truthy value. Will throw an error if the page is
	// closed before the worker event is fired.
	ExpectWorker(cb func() error, options ...PageExpectWorkerOptions) (Worker, error)
	ExpectedDialog(cb func() error) (Dialog, error)
	// This method waits for an element matching `selector`, waits for [actionability](../actionability.md) checks,
	// focuses the element, fills it and triggers an `input` event after filling. Note that you can pass an empty string
	// to clear the input field.
	// If the target element is not an `<input>`, `<textarea>` or `[contenteditable]` element, this method throws an
	// error. However, if the element is inside the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), the control will be filled
	// instead.
	// To send fine-grained keyboard events, use Page.type().
	Fill(selector, text string, options ...FrameFillOptions) error
	// This method fetches an element with `selector` and focuses it. If there's no element matching `selector`, the
	// method waits until a matching element appears in the DOM.
	Focus(expression string, options ...FrameFocusOptions) error
	// Returns frame matching the specified criteria. Either `name` or `url` must be specified.
	Frame(options PageFrameOptions) Frame
	// An array of all frames attached to the page.
	Frames() []Frame
	// When working with iframes, you can create a frame locator that will enter the iframe and allow selecting elements
	// in that iframe.
	FrameLocator(selector string) FrameLocator
	// Returns element attribute value.
	GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error)
	// Allows locating elements by their alt text.
	GetByAltText(text interface{}, options ...LocatorGetByAltTextOptions) Locator
	// Allows locating input elements by the text of the associated `<label>` or `aria-labelledby` element, or by the
	// `aria-label` attribute.
	GetByLabel(text interface{}, options ...LocatorGetByLabelOptions) Locator
	// Allows locating input elements by the placeholder text.
	GetByPlaceholder(text interface{}, options ...LocatorGetByPlaceholderOptions) Locator
	// Allows locating elements by their [ARIA role](https://www.w3.org/TR/wai-aria-1.2/#roles),
	// [ARIA attributes](https://www.w3.org/TR/wai-aria-1.2/#aria-attributes) and
	// [accessible name](https://w3c.github.io/accname/#dfn-accessible-name).
	GetByRole(role AriaRole, options ...LocatorGetByRoleOptions) Locator
	// Locate element by the test id.
	GetByTestId(testId interface{}) Locator
	// Allows locating elements that contain given text.
	// See also Locator.filter() that allows to match by another criteria, like an accessible role, and then
	// filter by the text content.
	GetByText(text interface{}, options ...LocatorGetByTextOptions) Locator
	// Allows locating elements by their title attribute.
	GetByTitle(title interface{}, options ...LocatorGetByTitleOptions) Locator
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
	// - there's an SSL error (e.g. in case of self-signed certificates).
	// - target URL is invalid.
	// - the `timeout` is exceeded during navigation.
	// - the remote server does not respond or is unreachable.
	// - the main resource failed to load.
	// The method will not throw an error when any valid HTTP status code is returned by the remote server, including 404
	// "Not Found" and 500 "Internal Server Error".  The status code for such responses can be retrieved by calling
	// Response.status().
	// **NOTE** The method either throws an error or returns a main resource response. The only exceptions are navigation
	// to `about:blank` or navigation to the same URL with a different hash, which would succeed and return `null`.
	// **NOTE** Headless mode doesn't support navigation to a PDF document. See the
	// [upstream issue](https://bugs.chromium.org/p/chromium/issues/detail?id=761295).
	Goto(url string, options ...PageGotoOptions) (Response, error)
	// This method hovers over an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to hover over the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	Hover(selector string, options ...PageHoverOptions) error
	// Returns `element.innerHTML`.
	InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error)
	// Returns `element.innerText`.
	InnerText(selector string, options ...PageInnerTextOptions) (string, error)
	// Indicates that the page has been closed.
	IsClosed() bool
	// Returns whether the element is checked. Throws if the element is not a checkbox or radio input.
	IsChecked(selector string, options ...FrameIsCheckedOptions) (bool, error)
	// Returns whether the element is disabled, the opposite of [enabled](../actionability.md#enabled).
	IsDisabled(selector string, options ...FrameIsDisabledOptions) (bool, error)
	// Returns whether the element is [editable](../actionability.md#editable).
	IsEditable(selector string, options ...FrameIsEditableOptions) (bool, error)
	// Returns whether the element is [enabled](../actionability.md#enabled).
	IsEnabled(selector string, options ...FrameIsEnabledOptions) (bool, error)
	// Returns whether the element is hidden, the opposite of [visible](../actionability.md#visible).  `selector` that
	// does not match any elements is considered hidden.
	IsHidden(selector string, options ...FrameIsHiddenOptions) (bool, error)
	// Returns whether the element is [visible](../actionability.md#visible). `selector` that does not match any elements
	// is considered not visible.
	IsVisible(selector string, options ...FrameIsVisibleOptions) (bool, error)
	// The method returns an element locator that can be used to perform actions on this page / frame. Locator is resolved
	// to the element immediately before performing an action, so a series of actions on the same locator can in fact be
	// performed on different DOM elements. That would happen if the DOM structure between those actions has changed.
	// [Learn more about locators](../locators.md).
	Locator(selector string, options ...PageLocatorOptions) Locator
	// The page's main frame. Page is guaranteed to have a main frame which persists during navigations.
	MainFrame() Frame
	// Returns the opener for popup pages and `null` for others. If the opener has been closed already the returns `null`.
	Opener() (Page, error)
	// Returns the PDF buffer.
	// **NOTE** Generating a pdf is currently only supported in Chromium headless.
	// `page.pdf()` generates a pdf of the page with `print` css media. To generate a pdf with `screen` media, call
	// Page.emulateMedia() before calling `page.pdf()`:
	// **NOTE** By default, `page.pdf()` generates a pdf with modified colors for printing. Use the
	// [`-webkit-print-color-adjust`](https://developer.mozilla.org/en-US/docs/Web/CSS/-webkit-print-color-adjust)
	// property to force rendering of exact colors.
	PDF(options ...PagePdfOptions) ([]byte, error)
	// Focuses the element, and then uses Keyboard.down() and Keyboard.up().
	// `key` can specify the intended
	// [keyboardEvent.key](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key) value or a single character
	// to generate the text for. A superset of the `key` values can be found
	// [here](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values). Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`,
	// etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When specified with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	Press(selector, key string, options ...PagePressOptions) error
	// The method finds an element matching the specified selector within the page. If no elements match the selector, the
	// return value resolves to `null`. To wait for an element on the page, use Locator.waitFor().
	QuerySelector(selector string) (ElementHandle, error)
	// The method finds all elements matching the specified selector within the page. If no elements match the selector,
	// the return value resolves to `[]`.
	QuerySelectorAll(selector string) ([]ElementHandle, error)
	// This method reloads the current page, in the same way as if the user had triggered a browser refresh. Returns the
	// main resource response. In case of multiple redirects, the navigation will resolve with the response of the last
	// redirect.
	Reload(options ...PageReloadOptions) (Response, error)
	// API testing helper associated with this page. This method returns the same instance as
	// BrowserContext.request() on the page's context. See BrowserContext.request() for more
	// details.
	Request() APIRequestContext
	// Routing provides the capability to modify network requests that are made by a page.
	// Once routing is enabled, every request matching the url pattern will stall unless it's continued, fulfilled or
	// aborted.
	// **NOTE** The handler will only be called for the first url if the response is a redirect.
	// **NOTE** Page.route() will not intercept requests intercepted by Service Worker. See
	// [this](https://github.com/microsoft/playwright/issues/1090) issue. We recommend disabling Service Workers when
	// using request interception by setting `Browser.newContext.serviceWorkers` to `'block'`.
	Route(url interface{}, handler routeHandler, times ...int) error
	// If specified the network requests that are made in the page will be served from the HAR file. Read more about
	// [Replaying from HAR](../network.md#replaying-from-har).
	// Playwright will not serve requests intercepted by Service Worker from the HAR file. See
	// [this](https://github.com/microsoft/playwright/issues/1090) issue. We recommend disabling Service Workers when
	// using request interception by setting `Browser.newContext.serviceWorkers` to `'block'`.
	RouteFromHAR(har string, options ...PageRouteFromHAROptions) error
	// Returns the buffer with the captured screenshot.
	Screenshot(options ...PageScreenshotOptions) ([]byte, error)
	// This method waits for an element matching `selector`, waits for [actionability](../actionability.md) checks, waits
	// until all specified options are present in the `<select>` element and selects these options.
	// If the target element is not a `<select>` element, this method throws an error. However, if the element is inside
	// the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), the control will be used
	// instead.
	// Returns the array of option values that have been successfully selected.
	// Triggers a `change` and `input` event once all the provided options have been selected.
	SelectOption(selector string, values SelectOptionValues, options ...FrameSelectOptionOptions) ([]string, error)
	SetContent(content string, options ...PageSetContentOptions) error
	// This setting will change the default maximum navigation time for the following methods and related shortcuts:
	// - Page.goBack()
	// - Page.goForward()
	// - Page.goto()
	// - Page.reload()
	// - Page.setContent()
	// - Page.waitForNavigation()
	// - Page.waitForURL()
	// **NOTE** Page.setDefaultNavigationTimeout() takes priority over Page.setDefaultTimeout(),
	// BrowserContext.setDefaultTimeout() and BrowserContext.setDefaultNavigationTimeout().
	SetDefaultNavigationTimeout(timeout float64)
	// This setting will change the default maximum time for all the methods accepting `timeout` option.
	// **NOTE** Page.setDefaultNavigationTimeout() takes priority over Page.setDefaultTimeout().
	SetDefaultTimeout(timeout float64)
	// The extra HTTP headers will be sent with every request the page initiates.
	// **NOTE** Page.setExtraHTTPHeaders() does not guarantee the order of headers in the outgoing requests.
	SetExtraHTTPHeaders(headers map[string]string) error
	// Sets the value of the file input to these file paths or files. If some of the `filePaths` are relative paths, then
	// they are resolved relative to the current working directory. For empty array, clears the selected files.
	// This method expects `selector` to point to an
	// [input element](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input). However, if the element is inside
	// the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), targets the control instead.
	SetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) error
	// In the case of multiple pages in a single browser, each page can have its own viewport size. However,
	// Browser.newContext() allows to set viewport size (and more) for all pages in the context at once.
	// Page.setViewportSize() will resize the page. A lot of websites don't expect phones to change size, so you
	// should set the viewport size before navigating to the page. Page.setViewportSize() will also reset
	// `screen` size, use Browser.newContext() with `screen` and `viewport` parameters if you need better
	// control of these properties.
	SetViewportSize(width, height int) error
	// This method taps an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.touchscreen() to tap the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	// **NOTE** Page.tap() the method will throw if `hasTouch` option of the browser context is false.
	Tap(selector string, options ...FrameTapOptions) error
	// Returns `element.textContent`.
	TextContent(selector string, options ...FrameTextContentOptions) (string, error)
	// Returns the page's title.
	Title() (string, error)
	// Sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the text. `page.type` can be used to
	// send fine-grained keyboard events. To fill values in form fields, use Page.fill().
	// To press a special key, like `Control` or `ArrowDown`, use Keyboard.press().
	Type(selector, text string, options ...PageTypeOptions) error
	URL() string
	// This method unchecks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Ensure that matched element is a checkbox or a radio input. If not, this method throws. If the element is
	// already unchecked, this method returns immediately.
	// 1. Wait for [actionability](../actionability.md) checks on the matched element, unless `force` option is set. If
	// the element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use Page.mouse() to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`.
	// Passing zero timeout disables this.
	Uncheck(selector string, options ...FrameUncheckOptions) error
	// Removes a route created with Page.route(). When `handler` is not specified, removes all routes for the
	// `url`.
	Unroute(url interface{}, handler ...routeHandler) error
	// Video object associated with this page.
	Video() Video
	ViewportSize() ViewportSize
	// **NOTE** In most cases, you should use Page.ExpectForEvent().
	// Waits for given `event` to fire. If predicate is provided, it passes event's value into the `predicate` function
	// and waits for `predicate(event)` to return a truthy value. Will throw an error if the page is closed before the
	// `event` is fired.
	WaitForEvent(event string, options ...PageWaitForEventOptions) (interface{}, error)
	// Returns when the `expression` returns a truthy value. It resolves to a JSHandle of the truthy value.
	WaitForFunction(expression string, arg interface{}, options ...FrameWaitForFunctionOptions) (JSHandle, error)
	// Returns when the required load state has been reached.
	// This resolves when the page reaches a required load state, `load` by default. The navigation must have been
	// committed when this method is called. If current document has already reached the required state, resolves
	// immediately.
	WaitForLoadState(options ...PageWaitForLoadStateOptions) error
	// Waits for the main frame navigation and returns the main resource response. In case of multiple redirects, the
	// navigation will resolve with the response of the last redirect. In case of navigation to a different anchor or
	// navigation due to History API usage, the navigation will resolve with `null`.
	WaitForNavigation(options ...PageWaitForNavigationOptions) (Response, error)
	// Waits for the matching request and returns it. See [waiting for event](../events.md#waiting-for-event) for more
	// details about events.
	WaitForRequest(url interface{}, options ...PageWaitForRequestOptions) (Request, error)
	// Returns the matched response. See [waiting for event](../events.md#waiting-for-event) for more details about
	// events.
	WaitForResponse(url interface{}, options ...PageWaitForResponseOptions) (Response, error)
	// Returns when element specified by selector satisfies `state` option. Returns `null` if waiting for `hidden` or
	// `detached`.
	// **NOTE** Playwright automatically waits for element to be ready before performing an action. Using `Locator`
	// objects and web-first assertions makes the code wait-for-selector-free.
	// Wait for the `selector` to satisfy `state` option (either appear/disappear from dom, or become visible/hidden). If
	// at the moment of calling the method `selector` already satisfies the condition, the method will return immediately.
	// If the selector doesn't satisfy the condition for the `timeout` milliseconds, the function will throw.
	WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandle, error)
	// Waits for the given `timeout` in milliseconds.
	// Note that `page.waitForTimeout()` should only be used for debugging. Tests using the timer in production are going
	// to be flaky. Use signals such as network events, selectors becoming visible and others instead.
	WaitForTimeout(timeout float64)
	// This method returns all of the dedicated
	// [WebWorkers](https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API) associated with the page.
	// **NOTE** This does not contain ServiceWorkers
	Workers() []Worker
	// This method drags the source element to the target element. It will first move to the source element, perform a
	// `mousedown`, then move to the target element and perform a `mouseup`.
	DragAndDrop(source, target string, options ...FrameDragAndDropOptions) error
	// Pauses script execution. Playwright will stop executing the script and wait for the user to either press 'Resume'
	// button in the page overlay or to call `playwright.resume()` in the DevTools console.
	// User can inspect selectors or perform manual steps while paused. Resume will continue running the original script
	// from the place it was paused.
	// **NOTE** This method requires Playwright to be started in a headed mode, with a falsy `headless` value in the
	// BrowserType.launch().
	Pause() error
	// Returns `input.value` for the selected `<input>` or `<textarea>` or `<select>` element.
	// Throws for non-input elements. However, if the element is inside the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), returns the value of the
	// control.
	InputValue(selector string, options ...FrameInputValueOptions) (string, error)
	// Waits for the main frame to navigate to the given URL.
	WaitForURL(url string, options ...FrameWaitForURLOptions) error
}

// The `PageAssertions` class provides assertion methods that can be used to make assertions about the `Page` state in
// the tests.
type PageAssertions interface {
	// Makes the assertion check for the opposite condition. For example, this code tests that the page URL doesn't
	// contain `"error"`:
	Not() PageAssertions
	// Ensures the page has the given title.
	ToHaveTitle(titleOrRegExp interface{}, options ...PageAssertionsToHaveTitleOptions) error
	// Ensures the page is navigated to the given URL.
	ToHaveURL(urlOrRegExp interface{}, options ...PageAssertionsToHaveURLOptions) error
	// The opposite of PageAssertions.toHaveTitle().
	NotToHaveTitle(titleOrRegExp interface{}, options ...PageAssertionsToHaveTitleOptions) error
	// The opposite of PageAssertions.toHaveURL().
	NotToHaveURL(urlOrRegExp interface{}, options ...PageAssertionsToHaveURLOptions) error
}

// Playwright gives you Web-First Assertions with convenience methods for creating assertions that will wait and retry
// until the expected condition is met.
// Consider the following example:
// Playwright will be re-testing the node with the selector `.status` until fetched Node has the `"Submitted"` text.
// It will be re-fetching the node and checking it over and over, until the condition is met or until the timeout is
// reached. You can pass this timeout as an option.
// By default, the timeout for assertions is set to 5 seconds.
type PlaywrightAssertions interface {
	// Creates a `APIResponseAssertions` object for the given `APIResponse`.
	APIResponse(response APIResponse) APIResponseAssertions
	// Creates a `LocatorAssertions` object for the given `Locator`.
	Locator(locator Locator) LocatorAssertions
	// Creates a `PageAssertions` object for the given `Page`.
	Page(page Page) PageAssertions
}

// Whenever the page sends a request for a network resource the following sequence of events are emitted by `Page`:
// - [`event: Page.request`] emitted when the request is issued by the page.
// - [`event: Page.response`] emitted when/if the response status and headers are received for the request.
// - [`event: Page.requestFinished`] emitted when the response body is downloaded and the request is complete.
// If request fails at some point, then instead of `'requestfinished'` event (and possibly instead of 'response'
// event), the  [`event: Page.requestFailed`] event is emitted.
// **NOTE** HTTP Error responses, such as 404 or 503, are still successful responses from HTTP standpoint, so request
// will complete with `'requestfinished'` event.
// If request gets a 'redirect' response, the request is successfully finished with the `requestfinished` event, and a
// new request is  issued to a redirected url.
type Request interface {
	// An object with all the request HTTP headers associated with this request. The header names are lower-cased.
	AllHeaders() (map[string]string, error)
	// An array with all the request HTTP headers associated with this request. Unlike Request.allHeaders(),
	// header names are NOT lower-cased. Headers with multiple entries, such as `Set-Cookie`, appear in the array multiple
	// times.
	HeadersArray() (HeadersArray, error)
	// Returns the value of the header matching the name. The name is case insensitive.
	HeaderValue(name string) (string, error)
	HeaderValues(name string) ([]string, error)
	// The method returns `null` unless this request has failed, as reported by `requestfailed` event.
	Failure() *RequestFailure
	// Returns the `Frame` that initiated this request.
	Frame() Frame
	// An object with the request HTTP headers. The header names are lower-cased. Note that this method does not return
	// security-related headers, including cookie-related ones. You can use Request.allHeaders() for complete
	// list of headers that include `cookie` information.
	Headers() map[string]string
	// Whether this request is driving frame's navigation.
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
	// When the server responds with a redirect, Playwright creates a new `Request` object. The two requests are connected
	// by `redirectedFrom()` and `redirectedTo()` methods. When multiple server redirects has happened, it is possible to
	// construct the whole redirect chain by repeatedly calling `redirectedFrom()`.
	RedirectedFrom() Request
	// New request issued by the browser if the server responded with redirect.
	RedirectedTo() Request
	// Contains the request's resource type as it was perceived by the rendering engine. ResourceType will be one of the
	// following: `document`, `stylesheet`, `image`, `media`, `font`, `script`, `texttrack`, `xhr`, `fetch`,
	// `eventsource`, `websocket`, `manifest`, `other`.
	ResourceType() string
	// Returns the matching `Response` object, or `null` if the response was not received due to error.
	Response() (Response, error)
	// Returns resource timing information for given request. Most of the timing values become available upon the
	// response, `responseEnd` becomes available when request finishes. Find more information at
	// [Resource Timing API](https://developer.mozilla.org/en-US/docs/Web/API/PerformanceResourceTiming).
	Timing() *ResourceTiming
	// URL of the request.
	URL() string
	// Returns resource size information for given request.
	Sizes() (*RequestSizesResult, error)
}

// `Response` class represents responses which are received by page.
type Response interface {
	// An object with all the response HTTP headers associated with this response.
	AllHeaders() (map[string]string, error)
	// An array with all the request HTTP headers associated with this response. Unlike Response.allHeaders(),
	// header names are NOT lower-cased. Headers with multiple entries, such as `Set-Cookie`, appear in the array multiple
	// times.
	HeadersArray() (HeadersArray, error)
	// Returns the value of the header matching the name. The name is case insensitive. If multiple headers have the same
	// name (except `set-cookie`), they are returned as a list separated by `, `. For `set-cookie`, the `\n` separator is
	// used. If no headers are found, `null` is returned.
	HeaderValue(name string) (string, error)
	// Returns all values of the headers matching the name, for example `set-cookie`. The name is case insensitive.
	HeaderValues(name string) ([]string, error)
	// Returns the buffer with response body.
	Body() ([]byte, error)
	// Waits for this response to finish, returns always `null`.
	Finished() error
	// Returns the `Frame` that initiated this response.
	Frame() Frame
	// Indicates whether this Response was fulfilled by a Service Worker's Fetch Handler (i.e. via
	// [FetchEvent.respondWith](https://developer.mozilla.org/en-US/docs/Web/API/FetchEvent/respondWith)).
	FromServiceWorker() bool
	// An object with the response HTTP headers. The header names are lower-cased. Note that this method does not return
	// security-related headers, including cookie-related ones. You can use Response.allHeaders() for complete
	// list of headers that include `cookie` information.
	Headers() map[string]string
	// Returns the JSON representation of response body.
	// This method will throw if the response body is not parsable via `JSON.parse`.
	JSON(v interface{}) error
	// Contains a boolean stating whether the response was successful (status in the range 200-299) or not.
	Ok() bool
	// Returns the matching `Request` object.
	Request() Request
	// Contains the status code of the response (e.g., 200 for a success).
	Status() int
	// Contains the status text of the response (e.g. usually an "OK" for a success).
	StatusText() string
	// Returns the text representation of response body.
	Text() (string, error)
	// Contains the URL of the response.
	URL() string
	// Returns SSL and other security information.
	SecurityDetails() (*ResponseSecurityDetailsResult, error)
	// Returns the IP address and port of the server.
	ServerAddr() (*ResponseServerAddrResult, error)
}

// Whenever a network route is set up with Page.route() or BrowserContext.route(), the `Route`
// object allows to handle the route.
// Learn more about [networking](../network.md).
type Route interface {
	// Aborts the route's request.
	Abort(errorCode ...string) error
	// Continues route's request with optional overrides.
	Continue(options ...RouteContinueOptions) error
	// When several routes match the given pattern, they run in the order opposite to their registration. That way the
	// last registered route can always override all the previous ones. In the example below, request will be handled by
	// the bottom-most handler first, then it'll fall back to the previous one and in the end will be aborted by the first
	// registered route.
	Fallback(options ...RouteFallbackOptions) error
	// Performs the request and fetches result without fulfilling it, so that the response could be modified and then
	// fulfilled.
	Fetch(options ...RouteFetchOptions) (APIResponse, error)
	// Fulfills route's request with given response.
	Fulfill(options RouteFulfillOptions) error
	// A request to be routed.
	Request() Request
}

// Selectors can be used to install custom selector engines. See [extensibility](../extensibility.md) for more
// information.
type Selectors interface {
	// Selectors must be registered before creating the page.
	Register(name string, option SelectorsRegisterOptions) error
	// Defines custom attribute name to be used in Page.getByTestId(). `data-testid` is used by default.
	SetTestIdAttribute(name string)
}

// The Touchscreen class operates in main-frame CSS pixels relative to the top-left corner of the viewport. Methods on
// the touchscreen can only be used in browser contexts that have been initialized with `hasTouch` set to true.
type Touchscreen interface {
	// Dispatches a `touchstart` and `touchend` event with a single touch at the position (`x`,`y`).
	// **NOTE** Page.tap() the method will throw if `hasTouch` option of the browser context is false.
	Tap(x int, y int) error
}

// The `WebSocket` class represents websocket connections in the page.
type WebSocket interface {
	EventEmitter
	// Indicates that the web socket has been closed.
	IsClosed() bool
	// Contains the URL of the WebSocket.
	URL() string
	// Waits for event to fire and passes its value into the predicate function. Returns when the predicate returns truthy
	// value. Will throw an error if the webSocket is closed before the event is fired. Returns the event data value.
	ExpectEvent(event string, cb func() error, options ...WebSocketWaitForEventOptions) (interface{}, error)
	// **NOTE** In most cases, you should use WebSocket.ExpectForEvent().
	// Waits for given `event` to fire. If predicate is provided, it passes event's value into the `predicate` function
	// and waits for `predicate(event)` to return a truthy value. Will throw an error if the socket is closed before the
	// `event` is fired.
	WaitForEvent(event string, options ...WebSocketWaitForEventOptions) (interface{}, error)
}

// When browser context is created with the `recordVideo` option, each page has a video object associated with it.
type Video interface {
	// Returns the file system path this video will be recorded to. The video is guaranteed to be written to the
	// filesystem upon closing the browser context. This method throws when connected remotely.
	Path() (string, error)
	// Deletes the video file. Will wait for the video to finish if necessary.
	Delete() error
	// Saves the video to a user-specified path. It is safe to call this method while the video is still in progress, or
	// after the page has closed. This method waits until the page is closed and the video is fully saved.
	SaveAs(path string) error
}

// The Worker class represents a [WebWorker](https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API).
// `worker` event is emitted on the page object to signal a worker creation. `close` event is emitted on the worker
// object when the worker is gone.
type Worker interface {
	EventEmitter
	// Returns the return value of `expression`.
	// If the function passed to the Worker.evaluate() returns a [Promise], then Worker.evaluate()
	// would wait for the promise to resolve and return its value.
	// If the function passed to the Worker.evaluate() returns a non-[Serializable] value, then
	// Worker.evaluate() returns `undefined`. Playwright also supports transferring some additional values that
	// are not serializable by `JSON`: `-0`, `NaN`, `Infinity`, `-Infinity`.
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `expression` as a `JSHandle`.
	// The only difference between Worker.evaluate() and Worker.evaluateHandle() is that
	// Worker.evaluateHandle() returns `JSHandle`.
	// If the function passed to the Worker.evaluateHandle() returns a [Promise], then
	// Worker.evaluateHandle() would wait for the promise to resolve and return its value.
	EvaluateHandle(expression string, options ...interface{}) (JSHandle, error)
	URL() string
	WaitForEvent(event string, predicate ...interface{}) (interface{}, error)
	ExpectEvent(event string, cb func() error, predicates ...interface{}) (interface{}, error)
}
