package playwright

type BindingCall interface {
	Call(f BindingCallFunction)
}

// A Browser is created via BrowserType.launch(). An example of using a `Browser` to create a `Page`:
type Browser interface {
	EventEmitter
	// In case this browser is obtained using BrowserType.launch(), closes the browser and all of its pages (if any
	// were opened).
	// In case this browser is connected to, clears all created contexts belonging to this browser and disconnects from the
	// browser server.
	// The `Browser` object itself is considered to be disposed and cannot be used anymore.
	Close() error
	// Returns an array of all open browser contexts. In a newly created browser, this will return zero browser contexts.
	Contexts() []BrowserContext
	// Indicates that the browser is connected.
	IsConnected() bool
	// Creates a new browser context. It won't share cookies/cache with other browser contexts.
	NewContext(options ...BrowserNewContextOptions) (BrowserContext, error)
	// Creates a new page in a new browser context. Closing this page will close the context as well.
	// This is a convenience API that should only be used for the single-page scenarios and short snippets. Production code and
	// testing frameworks should explicitly create Browser.newContext() followed by the
	// BrowserContext.newPage() to control their exact life times.
	NewPage(options ...BrowserNewContextOptions) (Page, error)
	// > NOTE: CDP Sessions are only supported on Chromium-based browsers.
	// Returns the newly created browser session.
	NewBrowserCDPSession() (CDPSession, error)
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
	// Detaches the CDPSession from the target. Once detached, the CDPSession object won't emit any events and can't be used to
	// send messages.
	Detach() error
	Send(method string, params map[string]interface{}) (interface{}, error)
}

// BrowserContexts provide a way to operate multiple independent browser sessions.
// If a page opens another page, e.g. with a `window.open` call, the popup will belong to the parent page's browser
// context.
// Playwright allows creating "incognito" browser contexts with Browser.newContext() method. "Incognito" browser
// contexts don't write any browsing data to disk.
type BrowserContext interface {
	EventEmitter
	// Adds cookies into this browser context. All pages within this context will have these cookies installed. Cookies can be
	// obtained via BrowserContext.cookies().
	AddCookies(cookies ...BrowserContextAddCookiesOptionsCookies) error
	// Adds a script which would be evaluated in one of the following scenarios:
	// - Whenever a page is created in the browser context or is navigated.
	// - Whenever a child frame is attached or navigated in any page in the browser context. In this case, the script is
	// evaluated in the context of the newly attached frame.
	// The script is evaluated after the document was created but before any of its scripts were run. This is useful to amend
	// the JavaScript environment, e.g. to seed `Math.random`.
	// An example of overriding `Math.random` before the page loads:
	// > NOTE: The order of evaluation of multiple scripts installed via BrowserContext.addInitScript() and
	// Page.addInitScript() is not defined.
	AddInitScript(script BrowserContextAddInitScriptOptions) error
	// Returns the browser instance of the context. If it was launched as a persistent context null gets returned.
	Browser() Browser
	// Clears context cookies.
	ClearCookies() error
	// Clears all permission overrides for the browser context.
	ClearPermissions() error
	// Closes the browser context. All the pages that belong to the browser context will be closed.
	// > NOTE: The default browser context cannot be closed.
	Close() error
	// If no URLs are specified, this method returns all cookies. If URLs are specified, only cookies that affect those URLs
	// are returned.
	Cookies(urls ...string) ([]*BrowserContextCookiesResult, error)
	ExpectEvent(event string, cb func() error) (interface{}, error)
	// The method adds a function called `name` on the `window` object of every frame in every page in the context. When
	// called, the function executes `callback` and returns a [Promise] which resolves to the return value of `callback`. If
	// the `callback` returns a [Promise], it will be awaited.
	// The first argument of the `callback` function contains information about the caller: `{ browserContext: BrowserContext,
	// page: Page, frame: Frame }`.
	// See Page.exposeBinding() for page-only version.
	// An example of exposing page URL to all frames in all pages in the context:
	// An example of passing an element handle:
	ExposeBinding(name string, binding BindingCallFunction, handle ...bool) error
	// The method adds a function called `name` on the `window` object of every frame in every page in the context. When
	// called, the function executes `callback` and returns a [Promise] which resolves to the return value of `callback`.
	// If the `callback` returns a [Promise], it will be awaited.
	// See Page.exposeFunction() for page-only version.
	// An example of adding a `sha256` function to all pages in the context:
	ExposeFunction(name string, binding ExposedFunction) error
	// Grants specified permissions to the browser context. Only grants corresponding permissions to the given origin if
	// specified.
	GrantPermissions(permissions []string, options ...BrowserContextGrantPermissionsOptions) error
	// > NOTE: CDP sessions are only supported on Chromium-based browsers.
	// Returns the newly created session.
	NewCDPSession(page Page) (CDPSession, error)
	// Creates a new page in the browser context.
	NewPage(options ...BrowserNewPageOptions) (Page, error)
	// Returns all open pages in the context.
	Pages() []Page
	// This setting will change the default maximum navigation time for the following methods and related shortcuts:
	// - Page.goBack()
	// - Page.goForward()
	// - Page.goto()
	// - Page.reload()
	// - Page.setContent()
	// - Page.waitForNavigation()
	// > NOTE: Page.setDefaultNavigationTimeout`] and [`method: Page.setDefaultTimeout() take priority over
	// BrowserContext.setDefaultNavigationTimeout().
	SetDefaultNavigationTimeout(timeout float64)
	// This setting will change the default maximum time for all the methods accepting `timeout` option.
	// > NOTE: Page.setDefaultNavigationTimeout`], [`method: Page.setDefaultTimeout() and
	// BrowserContext.setDefaultNavigationTimeout`] take priority over [`method: BrowserContext.setDefaultTimeout().
	SetDefaultTimeout(timeout float64)
	// The extra HTTP headers will be sent with every request initiated by any page in the context. These headers are merged
	// with page-specific extra HTTP headers set with Page.setExtraHTTPHeaders(). If page overrides a particular
	// header, page-specific header value will be used instead of the browser context header value.
	// > NOTE: BrowserContext.setExtraHTTPHeaders() does not guarantee the order of headers in the outgoing requests.
	SetExtraHTTPHeaders(headers map[string]string) error
	// Sets the context's geolocation. Passing `null` or `undefined` emulates position unavailable.
	// > NOTE: Consider using BrowserContext.grantPermissions() to grant permissions for the browser context pages to
	// read its geolocation.
	SetGeolocation(gelocation *SetGeolocationOptions) error
	ResetGeolocation() error
	// Routing provides the capability to modify network requests that are made by any page in the browser context. Once route
	// is enabled, every request matching the url pattern will stall unless it's continued, fulfilled or aborted.
	// > NOTE: Page.route() will not intercept requests intercepted by Service Worker. See
	// [this](https://github.com/microsoft/playwright/issues/1090) issue. We recommend disabling Service Workers when using
	// request interception. Via `await context.addInitScript(() => delete window.navigator.serviceWorker);`
	// An example of a naive handler that aborts all image requests:
	// or the same snippet using a regex pattern instead:
	// It is possible to examine the request to decide the route action. For example, mocking all requests that contain some
	// post data, and leaving all other requests as is:
	// Page routes (set up with Page.route()) take precedence over browser context routes when request matches both
	// handlers.
	// To remove a route with its handler you can use BrowserContext.unroute().
	// > NOTE: Enabling routing disables http cache.
	Route(url interface{}, handler routeHandler) error
	SetOffline(offline bool) error
	// Returns storage state for this browser context, contains current cookies and local storage snapshot.
	StorageState(path ...string) (*StorageState, error)
	// Removes a route created with BrowserContext.route(). When `handler` is not specified, removes all routes for
	// the `url`.
	Unroute(url interface{}, handler ...routeHandler) error
	// Waits for event to fire and passes its value into the predicate function. Returns when the predicate returns truthy
	// value. Will throw an error if the context closes before the event is fired. Returns the event data value.
	WaitForEvent(event string, predicate ...interface{}) interface{}
	Tracing() Tracing
	// > NOTE: Background pages are only supported on Chromium-based browsers.
	// All existing background pages in the context.
	BackgroundPages() []Page
}

// API for collecting and saving Playwright traces. Playwright traces can be opened in [Trace Viewer](./trace-viewer.md)
// after Playwright script runs.
// Start recording a trace before performing actions. At the end, stop tracing and save it to a file.
type Tracing interface {
	// Start tracing.
	Start(options ...TracingStartOptions) error
	// Stop tracing.
	Stop(options ...TracingStopOptions) error
	// Start a new trace chunk. If you'd like to record multiple traces on the same `BrowserContext`, use
	// Tracing.start`] once, and then create multiple trace chunks with [`method: Tracing.startChunk() and
	// Tracing.stopChunk().
	StartChunk(options ...TracingStartChunkOptions) error
	// Stop the trace chunk. See Tracing.startChunk() for more details about multiple trace chunks.
	StopChunk(options ...TracingStopChunkOptions) error
}

// BrowserType provides methods to launch a specific browser instance or connect to an existing one. The following is a
// typical example of using Playwright to drive automation:
type BrowserType interface {
	// A path where Playwright expects to find a bundled browser executable.
	ExecutablePath() string
	// Returns the browser instance.
	// You can use `ignoreDefaultArgs` to filter out `--mute-audio` from default arguments:
	// > **Chromium-only** Playwright can also be used to control the Google Chrome or Microsoft Edge browsers, but it works
	// best with the version of Chromium it is bundled with. There is no guarantee it will work with any other version. Use
	// `executablePath` option with extreme caution.
	// >
	// > If Google Chrome (rather than Chromium) is preferred, a
	// [Chrome Canary](https://www.google.com/chrome/browser/canary.html) or
	// [Dev Channel](https://www.chromium.org/getting-involved/dev-channel) build is suggested.
	// >
	// > Stock browsers like Google Chrome and Microsoft Edge are suitable for tests that require proprietary media codecs for
	// video playback. See
	// [this article](https://www.howtogeek.com/202825/what%E2%80%99s-the-difference-between-chromium-and-chrome/) for other
	// differences between Chromium and Chrome.
	// [This article](https://chromium.googlesource.com/chromium/src/+/lkgr/docs/chromium_browser_vs_google_chrome.md)
	// describes some differences for Linux users.
	Launch(options ...BrowserTypeLaunchOptions) (Browser, error)
	// Returns the persistent browser context instance.
	// Launches browser that uses persistent storage located at `userDataDir` and returns the only context. Closing this
	// context will automatically close the browser.
	LaunchPersistentContext(userDataDir string, options ...BrowserTypeLaunchPersistentContextOptions) (BrowserContext, error)
	// Returns browser name. For example: `'chromium'`, `'webkit'` or `'firefox'`.
	Name() string
	// This methods attaches Playwright to an existing browser instance.
	Connect(url string, options ...BrowserTypeConnectOptions) (Browser, error)
	// This methods attaches Playwright to an existing browser instance using the Chrome DevTools Protocol.
	// The default browser context is accessible via Browser.contexts().
	// > NOTE: Connecting over the Chrome DevTools Protocol is only supported for Chromium-based browsers.
	ConnectOverCDP(endpointURL string, options ...BrowserTypeConnectOverCDPOptions) (Browser, error)
}

// `ConsoleMessage` objects are dispatched by page via the [`event: Page.console`] event.
type ConsoleMessage interface {
	// List of arguments passed to a `console` function call. See also [`event: Page.console`].
	Args() []JSHandle
	Location() ConsoleMessageLocation
	String() string
	// The text of the console message.
	Text() string
	// One of the following values: `'log'`, `'debug'`, `'info'`, `'error'`, `'warning'`, `'dir'`, `'dirxml'`, `'table'`,
	// `'trace'`, `'clear'`, `'startGroup'`, `'startGroupCollapsed'`, `'endGroup'`, `'assert'`, `'profile'`, `'profileEnd'`,
	// `'count'`, `'timeEnd'`.
	Type() string
}

// `Dialog` objects are dispatched by page via the [`event: Page.dialog`] event.
// An example of using `Dialog` class:
// > NOTE: Dialogs are dismissed automatically, unless there is a [`event: Page.dialog`] listener. When listener is
// present, it **must** either Dialog.accept`] or [`method: Dialog.dismiss() the dialog - otherwise the page will
// [freeze](https://developer.mozilla.org/en-US/docs/Web/JavaScript/EventLoop#never_blocking) waiting for the dialog, and
// actions like click will never finish.
type Dialog interface {
	// Returns when the dialog has been accepted.
	Accept(texts ...string) error
	// If dialog is prompt, returns default prompt value. Otherwise, returns empty string.
	DefaultValue() string
	// Returns when the dialog has been dismissed.
	Dismiss() error
	// A message displayed in the dialog.
	Message() string
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
	// Returns path to the downloaded file in case of successful download. The method will wait for the download to finish if
	// necessary. The method throws when connected remotely.
	// Note that the download's file name is a random GUID, use Download.suggestedFilename() to get suggested file
	// name.
	Path() (string, error)
	// Copy the download to a user-specified path. It is safe to call this method while the download is still in progress. Will
	// wait for the download to finish if necessary.
	SaveAs(path string) error
	String() string
	// Returns suggested filename for this download. It is typically computed by the browser from the
	// [`Content-Disposition`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition) response header
	// or the `download` attribute. See the spec on [whatwg](https://html.spec.whatwg.org/#downloading-resources). Different
	// browsers can use different logic for computing it.
	SuggestedFilename() string
	// Returns downloaded url.
	URL() string
	// Get the page that the download belongs to.
	Page() Page
	// Cancels a download. Will not fail if the download is already finished or canceled. Upon successful cancellations,
	// `download.failure()` would resolve to `'canceled'`.
	Cancel() error
}

// ElementHandle represents an in-page DOM element. ElementHandles can be created with the Page.querySelector()
// method.
// > NOTE: The use of ElementHandle is discouraged, use `Locator` objects and web-first assertions instead.
// ElementHandle prevents DOM element from garbage collection unless the handle is disposed with
// JSHandle.dispose(). ElementHandles are auto-disposed when their origin frame gets navigated.
// ElementHandle instances can be used as an argument in Page.evalOnSelector`] and [`method: Page.evaluate()
// methods.
// The difference between the `Locator` and ElementHandle is that the ElementHandle points to a particular element, while
// `Locator` captures the logic of how to retrieve an element.
// In the example below, handle points to a particular DOM element on page. If that element changes text or is used by
// React to render an entirely different component, handle is still pointing to that very DOM element. This can lead to
// unexpected behaviors.
// With the locator, every time the `element` is used, up-to-date DOM element is located in the page using the selector. So
// in the snippet below, underlying DOM element is going to be located twice.
type ElementHandle interface {
	JSHandle
	// This method returns the bounding box of the element, or `null` if the element is not visible. The bounding box is
	// calculated relative to the main frame viewport - which is usually the same as the browser window.
	// Scrolling affects the returned bonding box, similarly to
	// [Element.getBoundingClientRect](https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect). That
	// means `x` and/or `y` may be negative.
	// Elements from child frames return the bounding box relative to the main frame, unlike the
	// [Element.getBoundingClientRect](https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect).
	// Assuming the page is static, it is safe to use bounding box coordinates to perform input. For example, the following
	// snippet should click the center of the element.
	BoundingBox() (*Rect, error)
	// This method checks the element by performing the following steps:
	// 1. Ensure that element is a checkbox or a radio input. If not, this method throws. If the element is already checked,
	// this method returns immediately.
	// 1. Wait for [actionability](./actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now checked. If not, this method throws.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	Check(options ...ElementHandleCheckOptions) error
	// This method checks or unchecks an element by performing the following steps:
	// 1. Ensure that element is a checkbox or a radio input. If not, this method throws.
	// 1. If the element already has the right checked state, this method returns immediately.
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now checked or unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	SetChecked(checked bool, options ...ElementHandleSetCheckedOptions) error
	// This method clicks the element by performing the following steps:
	// 1. Wait for [actionability](./actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to click in the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	Click(options ...ElementHandleClickOptions) error
	// Returns the content frame for element handles referencing iframe nodes, or `null` otherwise
	ContentFrame() (Frame, error)
	// This method double clicks the element by performing the following steps:
	// 1. Wait for [actionability](./actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to double click in the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set. Note that if the
	// first click of the `dblclick()` triggers a navigation event, this method will throw.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	// > NOTE: `elementHandle.dblclick()` dispatches two `click` events and a single `dblclick` event.
	Dblclick(options ...ElementHandleDblclickOptions) error
	// The snippet below dispatches the `click` event on the element. Regardless of the visibility state of the element,
	// `click` is dispatched. This is equivalent to calling
	// [element.click()](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click).
	// Under the hood, it creates an instance of an event based on the given `type`, initializes it with `eventInit` properties
	// and dispatches it on the element. Events are `composed`, `cancelable` and bubble by default.
	// Since `eventInit` is event-specific, please refer to the events documentation for the lists of initial properties:
	// - [DragEvent](https://developer.mozilla.org/en-US/docs/Web/API/DragEvent/DragEvent)
	// - [FocusEvent](https://developer.mozilla.org/en-US/docs/Web/API/FocusEvent/FocusEvent)
	// - [KeyboardEvent](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/KeyboardEvent)
	// - [MouseEvent](https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/MouseEvent)
	// - [PointerEvent](https://developer.mozilla.org/en-US/docs/Web/API/PointerEvent/PointerEvent)
	// - [TouchEvent](https://developer.mozilla.org/en-US/docs/Web/API/TouchEvent/TouchEvent)
	// - [Event](https://developer.mozilla.org/en-US/docs/Web/API/Event/Event)
	// You can also specify `JSHandle` as the property value if you want live objects to be passed into the event:
	DispatchEvent(typ string, initObjects ...interface{}) error
	// Returns the return value of `expression`.
	// The method finds an element matching the specified selector in the `ElementHandle`s subtree and passes it as a first
	// argument to `expression`. See [Working with selectors](./selectors.md) for more details. If no elements match the
	// selector, the method throws an error.
	// If `expression` returns a [Promise], then ElementHandle.evalOnSelector() would wait for the promise to resolve
	// and return its value.
	// Examples:
	EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `expression`.
	// The method finds all elements matching the specified selector in the `ElementHandle`'s subtree and passes an array of
	// matched elements as a first argument to `expression`. See [Working with selectors](./selectors.md) for more details.
	// If `expression` returns a [Promise], then ElementHandle.evalOnSelectorAll() would wait for the promise to
	// resolve and return its value.
	// Examples:
	// ```html
	// <div class="feed">
	// <div class="tweet">Hello!</div>
	// <div class="tweet">Hi!</div>
	// </div>
	// ```
	EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
	// This method waits for [actionability](./actionability.md) checks, focuses the element, fills it and triggers an `input`
	// event after filling. Note that you can pass an empty string to clear the input field.
	// If the target element is not an `<input>`, `<textarea>` or `[contenteditable]` element, this method throws an error.
	// However, if the element is inside the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), the control will be filled
	// instead.
	// To send fine-grained keyboard events, use ElementHandle.type().
	Fill(value string, options ...ElementHandleFillOptions) error
	// Calls [focus](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus) on the element.
	Focus() error
	// Returns element attribute value.
	GetAttribute(name string) (string, error)
	// This method hovers over the element by performing the following steps:
	// 1. Wait for [actionability](./actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to hover over the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	Hover(options ...ElementHandleHoverOptions) error
	// Returns the `element.innerHTML`.
	InnerHTML() (string, error)
	// Returns the `element.innerText`.
	InnerText() (string, error)
	// Returns whether the element is checked. Throws if the element is not a checkbox or radio input.
	IsChecked() (bool, error)
	// Returns whether the element is disabled, the opposite of [enabled](./actionability.md#enabled).
	IsDisabled() (bool, error)
	// Returns whether the element is [editable](./actionability.md#editable).
	IsEditable() (bool, error)
	// Returns whether the element is [enabled](./actionability.md#enabled).
	IsEnabled() (bool, error)
	// Returns whether the element is hidden, the opposite of [visible](./actionability.md#visible).
	IsHidden() (bool, error)
	// Returns whether the element is [visible](./actionability.md#visible).
	IsVisible() (bool, error)
	// Returns the frame containing the given element.
	OwnerFrame() (Frame, error)
	// Focuses the element, and then uses Keyboard.down`] and [`method: Keyboard.up().
	// `key` can specify the intended [keyboardEvent.key](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key)
	// value or a single character to generate the text for. A superset of the `key` values can be found
	// [here](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values). Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`, etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When specified with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	Press(key string, options ...ElementHandlePressOptions) error
	// The method finds an element matching the specified selector in the `ElementHandle`'s subtree. See
	// [Working with selectors](./selectors.md) for more details. If no elements match the selector, returns `null`.
	QuerySelector(selector string) (ElementHandle, error)
	// The method finds all elements matching the specified selector in the `ElementHandle`s subtree. See
	// [Working with selectors](./selectors.md) for more details. If no elements match the selector, returns empty array.
	QuerySelectorAll(selector string) ([]ElementHandle, error)
	// Returns the buffer with the captured screenshot.
	// This method waits for the [actionability](./actionability.md) checks, then scrolls element into view before taking a
	// screenshot. If the element is detached from DOM, the method throws an error.
	Screenshot(options ...ElementHandleScreenshotOptions) ([]byte, error)
	// This method waits for [actionability](./actionability.md) checks, then tries to scroll element into view, unless it is
	// completely visible as defined by
	// [IntersectionObserver](https://developer.mozilla.org/en-US/docs/Web/API/Intersection_Observer_API)'s `ratio`.
	// Throws when `elementHandle` does not point to an element
	// [connected](https://developer.mozilla.org/en-US/docs/Web/API/Node/isConnected) to a Document or a ShadowRoot.
	ScrollIntoViewIfNeeded(options ...ElementHandleScrollIntoViewIfNeededOptions) error
	// This method waits for [actionability](./actionability.md) checks, waits until all specified options are present in the
	// `<select>` element and selects these options.
	// If the target element is not a `<select>` element, this method throws an error. However, if the element is inside the
	// `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), the control will be used instead.
	// Returns the array of option values that have been successfully selected.
	// Triggers a `change` and `input` event once all the provided options have been selected.
	SelectOption(values SelectOptionValues, options ...ElementHandleSelectOptionOptions) ([]string, error)
	// This method waits for [actionability](./actionability.md) checks, then focuses the element and selects all its text
	// content.
	SelectText(options ...ElementHandleSelectTextOptions) error
	// This method expects `elementHandle` to point to an
	// [input element](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input).
	// Sets the value of the file input to these file paths or files. If some of the `filePaths` are relative paths, then they
	// are resolved relative to the the current working directory. For empty array, clears the selected files.
	SetInputFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error
	// This method taps the element by performing the following steps:
	// 1. Wait for [actionability](./actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.touchscreen`] to tap the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	// > NOTE: `elementHandle.tap()` requires that the `hasTouch` option of the browser context be set to true.
	Tap(options ...ElementHandleTapOptions) error
	// Returns the `node.textContent`.
	TextContent() (string, error)
	// Focuses the element, and then sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the text.
	// To press a special key, like `Control` or `ArrowDown`, use ElementHandle.press().
	// An example of typing into a text field and then submitting the form:
	Type(value string, options ...ElementHandleTypeOptions) error
	// This method checks the element by performing the following steps:
	// 1. Ensure that element is a checkbox or a radio input. If not, this method throws. If the element is already
	// unchecked, this method returns immediately.
	// 1. Wait for [actionability](./actionability.md) checks on the element, unless `force` option is set.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now unchecked. If not, this method throws.
	// If the element is detached from the DOM at any moment during the action, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	Uncheck(options ...ElementHandleUncheckOptions) error
	// Returns when the element satisfies the `state`.
	// Depending on the `state` parameter, this method waits for one of the [actionability](./actionability.md) checks to pass.
	// This method throws when the element is detached while waiting, unless waiting for the `"hidden"` state.
	// - `"visible"` Wait until the element is [visible](./actionability.md#visible).
	// - `"hidden"` Wait until the element is [not visible](./actionability.md#visible) or
	// [not attached](./actionability.md#attached). Note that waiting for hidden does not throw when the element detaches.
	// - `"stable"` Wait until the element is both [visible](./actionability.md#visible) and
	// [stable](./actionability.md#stable).
	// - `"enabled"` Wait until the element is [enabled](./actionability.md#enabled).
	// - `"disabled"` Wait until the element is [not enabled](./actionability.md#enabled).
	// - `"editable"` Wait until the element is [editable](./actionability.md#editable).
	// If the element does not satisfy the condition for the `timeout` milliseconds, this method will throw.
	WaitForElementState(state string, options ...ElementHandleWaitForElementStateOptions) error
	// Returns element specified by selector when it satisfies `state` option. Returns `null` if waiting for `hidden` or
	// `detached`.
	// Wait for the `selector` relative to the element handle to satisfy `state` option (either appear/disappear from dom, or
	// become visible/hidden). If at the moment of calling the method `selector` already satisfies the condition, the method
	// will return immediately. If the selector doesn't satisfy the condition for the `timeout` milliseconds, the function will
	// throw.
	// > NOTE: This method does not work across navigations, use Page.waitForSelector() instead.
	WaitForSelector(selector string, options ...ElementHandleWaitForSelectorOptions) (ElementHandle, error)
	// Returns `input.value` for `<input>` or `<textarea>` or `<select>` element. Throws for non-input elements.
	InputValue(options ...ElementHandleInputValueOptions) (string, error)
}

type EventEmitter interface {
	Emit(name string, payload ...interface{})
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
	// Sets the value of the file input this chooser is associated with. If some of the `filePaths` are relative paths, then
	// they are resolved relative to the the current working directory. For empty array, clears the selected files.
	SetFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error
}

// At every point of time, page exposes its current frame tree via the Page.mainFrame() and
// Frame.childFrames() methods.
// `Frame` object's lifecycle is controlled by three events, dispatched on the page object:
// - [`event: Page.frameAttached`] - fired when the frame gets attached to the page. A Frame can be attached to the page
// only once.
// - [`event: Page.frameNavigated`] - fired when the frame commits navigation to a different URL.
// - [`event: Page.frameDetached`] - fired when the frame gets detached from the page.  A Frame can be detached from the
// page only once.
// An example of dumping frame tree:
type Frame interface {
	// This method checks or unchecks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Ensure that matched element is a checkbox or a radio input. If not, this method throws.
	// 1. If the element already has the right checked state, this method returns immediately.
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now checked or unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
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
	// 1. Ensure that matched element is a checkbox or a radio input. If not, this method throws. If the element is already
	// checked, this method returns immediately.
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now checked. If not, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	Check(selector string, options ...FrameCheckOptions) error
	ChildFrames() []Frame
	// This method clicks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to click in the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	Click(selector string, options ...PageClickOptions) error
	// Gets the full HTML contents of the frame, including the doctype.
	Content() (string, error)
	// This method double clicks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to double click in the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set. Note that if the
	// first click of the `dblclick()` triggers a navigation event, this method will throw.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	// > NOTE: `frame.dblclick()` dispatches two `click` events and a single `dblclick` event.
	Dblclick(selector string, options ...FrameDblclickOptions) error
	// The snippet below dispatches the `click` event on the element. Regardless of the visibility state of the element,
	// `click` is dispatched. This is equivalent to calling
	// [element.click()](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click).
	// Under the hood, it creates an instance of an event based on the given `type`, initializes it with `eventInit` properties
	// and dispatches it on the element. Events are `composed`, `cancelable` and bubble by default.
	// Since `eventInit` is event-specific, please refer to the events documentation for the lists of initial properties:
	// - [DragEvent](https://developer.mozilla.org/en-US/docs/Web/API/DragEvent/DragEvent)
	// - [FocusEvent](https://developer.mozilla.org/en-US/docs/Web/API/FocusEvent/FocusEvent)
	// - [KeyboardEvent](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/KeyboardEvent)
	// - [MouseEvent](https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/MouseEvent)
	// - [PointerEvent](https://developer.mozilla.org/en-US/docs/Web/API/PointerEvent/PointerEvent)
	// - [TouchEvent](https://developer.mozilla.org/en-US/docs/Web/API/TouchEvent/TouchEvent)
	// - [Event](https://developer.mozilla.org/en-US/docs/Web/API/Event/Event)
	// You can also specify `JSHandle` as the property value if you want live objects to be passed into the event:
	DispatchEvent(selector, typ string, eventInit interface{}, options ...PageDispatchEventOptions) error
	// Returns the return value of `expression`.
	// If the function passed to the Frame.evaluate`] returns a [Promise], then [`method: Frame.evaluate() would wait
	// for the promise to resolve and return its value.
	// If the function passed to the Frame.evaluate() returns a non-[Serializable] value, then
	// Frame.evaluate() returns `undefined`. Playwright also supports transferring some additional values that are
	// not serializable by `JSON`: `-0`, `NaN`, `Infinity`, `-Infinity`.
	// A string can also be passed in instead of a function.
	// `ElementHandle` instances can be passed as an argument to the Frame.evaluate():
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `expression` as a `JSHandle`.
	// The only difference between Frame.evaluate`] and [`method: Frame.evaluateHandle() is that
	// Frame.evaluateHandle() returns `JSHandle`.
	// If the function, passed to the Frame.evaluateHandle(), returns a [Promise], then
	// Frame.evaluateHandle() would wait for the promise to resolve and return its value.
	// A string can also be passed in instead of a function.
	// `JSHandle` instances can be passed as an argument to the Frame.evaluateHandle():
	EvaluateHandle(expression string, options ...interface{}) (JSHandle, error)
	// Returns the return value of `expression`.
	// > NOTE: This method does not wait for the element to pass actionability checks and therefore can lead to the flaky
	// tests. Use Locator.evaluate(), other `Locator` helper methods or web-first assertions instead.
	// The method finds an element matching the specified selector within the frame and passes it as a first argument to
	// `expression`. See [Working with selectors](./selectors.md) for more details. If no elements match the selector, the
	// method throws an error.
	// If `expression` returns a [Promise], then Frame.evalOnSelector() would wait for the promise to resolve and
	// return its value.
	// Examples:
	EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `expression`.
	// > NOTE: In most cases, Locator.evaluateAll(), other `Locator` helper methods and web-first assertions do a
	// better job.
	// The method finds all elements matching the specified selector within the frame and passes an array of matched elements
	// as a first argument to `expression`. See [Working with selectors](./selectors.md) for more details.
	// If `expression` returns a [Promise], then Frame.evalOnSelectorAll() would wait for the promise to resolve and
	// return its value.
	// Examples:
	EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
	// This method waits for an element matching `selector`, waits for [actionability](./actionability.md) checks, focuses the
	// element, fills it and triggers an `input` event after filling. Note that you can pass an empty string to clear the input
	// field.
	// If the target element is not an `<input>`, `<textarea>` or `[contenteditable]` element, this method throws an error.
	// However, if the element is inside the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), the control will be filled
	// instead.
	// To send fine-grained keyboard events, use Frame.type().
	Fill(selector string, value string, options ...FrameFillOptions) error
	// This method fetches an element with `selector` and focuses it. If there's no element matching `selector`, the method
	// waits until a matching element appears in the DOM.
	Focus(selector string, options ...FrameFocusOptions) error
	// Returns the `frame` or `iframe` element handle which corresponds to this frame.
	// This is an inverse of ElementHandle.contentFrame(). Note that returned handle actually belongs to the parent
	// frame.
	// This method throws an error if the frame has been detached before `frameElement()` returns.
	FrameElement() (ElementHandle, error)
	// Returns element attribute value.
	GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error)
	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of the
	// last redirect.
	// The method will throw an error if:
	// - there's an SSL error (e.g. in case of self-signed certificates).
	// - target URL is invalid.
	// - the `timeout` is exceeded during navigation.
	// - the remote server does not respond or is unreachable.
	// - the main resource failed to load.
	// The method will not throw an error when any valid HTTP status code is returned by the remote server, including 404 "Not
	// Found" and 500 "Internal Server Error".  The status code for such responses can be retrieved by calling
	// Response.status().
	// > NOTE: The method either throws an error or returns a main resource response. The only exceptions are navigation to
	// `about:blank` or navigation to the same URL with a different hash, which would succeed and return `null`.
	// > NOTE: Headless mode doesn't support navigation to a PDF document. See the
	// [upstream issue](https://bugs.chromium.org/p/chromium/issues/detail?id=761295).
	Goto(url string, options ...PageGotoOptions) (Response, error)
	// This method hovers over an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to hover over the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	Hover(selector string, options ...PageHoverOptions) error
	// Returns `element.innerHTML`.
	InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error)
	// Returns `element.innerText`.
	InnerText(selector string, options ...PageInnerTextOptions) (string, error)
	// Returns `true` if the frame has been detached, or `false` otherwise.
	IsDetached() bool
	// Returns whether the element is checked. Throws if the element is not a checkbox or radio input.
	IsChecked(selector string, options ...FrameIsCheckedOptions) (bool, error)
	// Returns whether the element is disabled, the opposite of [enabled](./actionability.md#enabled).
	IsDisabled(selector string, options ...FrameIsDisabledOptions) (bool, error)
	// Returns whether the element is [editable](./actionability.md#editable).
	IsEditable(selector string, options ...FrameIsEditableOptions) (bool, error)
	// Returns whether the element is [enabled](./actionability.md#enabled).
	IsEnabled(selector string, options ...FrameIsEnabledOptions) (bool, error)
	// Returns whether the element is hidden, the opposite of [visible](./actionability.md#visible).  `selector` that does not
	// match any elements is considered hidden.
	IsHidden(selector string, options ...FrameIsHiddenOptions) (bool, error)
	// Returns whether the element is [visible](./actionability.md#visible). `selector` that does not match any elements is
	// considered not visible.
	IsVisible(selector string, options ...FrameIsVisibleOptions) (bool, error)
	// Returns frame's name attribute as specified in the tag.
	// If the name is empty, returns the id attribute instead.
	// > NOTE: This value is calculated once when the frame is created, and will not update if the attribute is changed later.
	Name() string
	// Returns the page containing this frame.
	Page() Page
	// Parent frame, if any. Detached frames and main frames return `null`.
	ParentFrame() Frame
	// `key` can specify the intended [keyboardEvent.key](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key)
	// value or a single character to generate the text for. A superset of the `key` values can be found
	// [here](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values). Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`, etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When specified with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	Press(selector, key string, options ...PagePressOptions) error
	// Returns the ElementHandle pointing to the frame element.
	// > NOTE: The use of `ElementHandle` is discouraged, use `Locator` objects and web-first assertions instead.
	// The method finds an element matching the specified selector within the frame. See
	// [Working with selectors](./selectors.md) for more details. If no elements match the selector, returns `null`.
	QuerySelector(selector string) (ElementHandle, error)
	// Returns the ElementHandles pointing to the frame elements.
	// > NOTE: The use of `ElementHandle` is discouraged, use `Locator` objects instead.
	// The method finds all elements matching the specified selector within the frame. See
	// [Working with selectors](./selectors.md) for more details. If no elements match the selector, returns empty array.
	QuerySelectorAll(selector string) ([]ElementHandle, error)
	SetContent(content string, options ...PageSetContentOptions) error
	// This method waits for an element matching `selector`, waits for [actionability](./actionability.md) checks, waits until
	// all specified options are present in the `<select>` element and selects these options.
	// If the target element is not a `<select>` element, this method throws an error. However, if the element is inside the
	// `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), the control will be used instead.
	// Returns the array of option values that have been successfully selected.
	// Triggers a `change` and `input` event once all the provided options have been selected.
	SelectOption(selector string, values SelectOptionValues, options ...FrameSelectOptionOptions) ([]string, error)
	// This method expects `selector` to point to an
	// [input element](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input).
	// Sets the value of the file input to these file paths or files. If some of the `filePaths` are relative paths, then they
	// are resolved relative to the the current working directory. For empty array, clears the selected files.
	SetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) error
	// This method taps an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.touchscreen`] to tap the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	// > NOTE: `frame.tap()` requires that the `hasTouch` option of the browser context be set to true.
	Tap(selector string, options ...FrameTapOptions) error
	// Returns `element.textContent`.
	TextContent(selector string, options ...FrameTextContentOptions) (string, error)
	// Returns the page title.
	Title() (string, error)
	// Sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the text. `frame.type` can be used to
	// send fine-grained keyboard events. To fill values in form fields, use Frame.fill().
	// To press a special key, like `Control` or `ArrowDown`, use Keyboard.press().
	Type(selector, text string, options ...PageTypeOptions) error
	// Returns frame's url.
	URL() string
	// This method checks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Ensure that matched element is a checkbox or a radio input. If not, this method throws. If the element is already
	// unchecked, this method returns immediately.
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	Uncheck(selector string, options ...FrameUncheckOptions) error
	WaitForEvent(event string, predicate ...interface{}) interface{}
	// Returns when the `expression` returns a truthy value, returns that value.
	// The Frame.waitForFunction() can be used to observe viewport size change:
	// To pass an argument to the predicate of `frame.waitForFunction` function:
	WaitForFunction(expression string, arg interface{}, options ...FrameWaitForFunctionOptions) (JSHandle, error)
	// Waits for the required load state to be reached.
	// This returns when the frame reaches a required load state, `load` by default. The navigation must have been committed
	// when this method is called. If current document has already reached the required state, resolves immediately.
	WaitForLoadState(given ...string)
	// Waits for the frame navigation and returns the main resource response. In case of multiple redirects, the navigation
	// will resolve with the response of the last redirect. In case of navigation to a different anchor or navigation due to
	// History API usage, the navigation will resolve with `null`.
	// This method waits for the frame to navigate to a new URL. It is useful for when you run code which will indirectly cause
	// the frame to navigate. Consider this example:
	// > NOTE: Usage of the [History API](https://developer.mozilla.org/en-US/docs/Web/API/History_API) to change the URL is
	// considered a navigation.
	WaitForNavigation(options ...PageWaitForNavigationOptions) (Response, error)
	// Waits for the frame to navigate to the given URL.
	WaitForURL(url string, options ...FrameWaitForURLOptions) error
	// Returns when element specified by selector satisfies `state` option. Returns `null` if waiting for `hidden` or
	// `detached`.
	// > NOTE: Playwright automatically waits for element to be ready before performing an action. Using `Locator` objects and
	// web-first assertions make the code wait-for-selector-free.
	// Wait for the `selector` to satisfy `state` option (either appear/disappear from dom, or become visible/hidden). If at
	// the moment of calling the method `selector` already satisfies the condition, the method will return immediately. If the
	// selector doesn't satisfy the condition for the `timeout` milliseconds, the function will throw.
	// This method works across navigations:
	WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandle, error)
	// Waits for the given `timeout` in milliseconds.
	// Note that `frame.waitForTimeout()` should only be used for debugging. Tests using the timer in production are going to
	// be flaky. Use signals such as network events, selectors becoming visible and others instead.
	WaitForTimeout(timeout float64)
	// Returns `input.value` for the selected `<input>` or `<textarea>` or `<select>` element. Throws for non-input elements.
	InputValue(selector string, options ...FrameInputValueOptions) (string, error)
	DragAndDrop(source, target string, options ...FrameDragAndDropOptions) error
}

// JSHandle represents an in-page JavaScript object. JSHandles can be created with the Page.evaluateHandle()
// method.
// JSHandle prevents the referenced JavaScript object being garbage collected unless the handle is exposed with
// JSHandle.dispose(). JSHandles are auto-disposed when their origin frame gets navigated or the parent context
// gets destroyed.
// JSHandle instances can be used as an argument in Page.evalOnSelector`], [`method: Page.evaluate() and
// Page.evaluateHandle() methods.
type JSHandle interface {
	// Returns either `null` or the object handle itself, if the object handle is an instance of `ElementHandle`.
	AsElement() ElementHandle
	// The `jsHandle.dispose` method stops referencing the element handle.
	Dispose() error
	// Returns the return value of `expression`.
	// This method passes this handle as the first argument to `expression`.
	// If `expression` returns a [Promise], then `handle.evaluate` would wait for the promise to resolve and return its value.
	// Examples:
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `expression` as a `JSHandle`.
	// This method passes this handle as the first argument to `expression`.
	// The only difference between `jsHandle.evaluate` and `jsHandle.evaluateHandle` is that `jsHandle.evaluateHandle` returns
	// `JSHandle`.
	// If the function passed to the `jsHandle.evaluateHandle` returns a [Promise], then `jsHandle.evaluateHandle` would wait
	// for the promise to resolve and return its value.
	// See Page.evaluateHandle() for more details.
	EvaluateHandle(expression string, options ...interface{}) (JSHandle, error)
	// The method returns a map with **own property names** as keys and JSHandle instances for the property values.
	GetProperties() (map[string]JSHandle, error)
	// Fetches a single property from the referenced object.
	GetProperty(name string) (JSHandle, error)
	// Returns a JSON representation of the object. If the object has a `toJSON` function, it **will not be called**.
	// > NOTE: The method will return an empty JSON object if the referenced object is not stringifiable. It will throw an
	// error if the object has circular references.
	JSONValue() (interface{}, error)
	String() string
}

// Keyboard provides an api for managing a virtual keyboard. The high level api is Keyboard.type(), which takes
// raw characters and generates proper keydown, keypress/input, and keyup events on your page.
// For finer control, you can use Keyboard.down`], [`method: Keyboard.up`], and [`method: Keyboard.insertText()
// to manually fire events as if they were generated from a real keyboard.
// An example of holding down `Shift` in order to select and delete some text:
// An example of pressing uppercase `A`
// An example to trigger select-all with the keyboard
type Keyboard interface {
	// Dispatches a `keydown` event.
	// `key` can specify the intended [keyboardEvent.key](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key)
	// value or a single character to generate the text for. A superset of the `key` values can be found
	// [here](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values). Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`, etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// If `key` is a modifier key, `Shift`, `Meta`, `Control`, or `Alt`, subsequent key presses will be sent with that modifier
	// active. To release the modifier key, use Keyboard.up().
	// After the key is pressed once, subsequent calls to Keyboard.down() will have
	// [repeat](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/repeat) set to true. To release the key, use
	// Keyboard.up().
	// > NOTE: Modifier keys DO influence `keyboard.down`. Holding down `Shift` will type the text in upper case.
	Down(key string) error
	// Dispatches only `input` event, does not emit the `keydown`, `keyup` or `keypress` events.
	// > NOTE: Modifier keys DO NOT effect `keyboard.insertText`. Holding down `Shift` will not type the text in upper case.
	InsertText(text string) error
	// `key` can specify the intended [keyboardEvent.key](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key)
	// value or a single character to generate the text for. A superset of the `key` values can be found
	// [here](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values). Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`, etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When specified with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	// Shortcut for Keyboard.down`] and [`method: Keyboard.up().
	Press(key string, options ...KeyboardPressOptions) error
	// Sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the text.
	// To press a special key, like `Control` or `ArrowDown`, use Keyboard.press().
	// > NOTE: Modifier keys DO NOT effect `keyboard.type`. Holding down `Shift` will not type the text in upper case.
	// > NOTE: For characters that are not on a US keyboard, only an `input` event will be sent.
	Type(text string, options ...KeyboardTypeOptions) error
	// Dispatches a `keyup` event.
	Up(key string) error
}

// The Mouse class operates in main-frame CSS pixels relative to the top-left corner of the viewport.
// Every `page` object has its own Mouse, accessible with [`property: Page.mouse`].
type Mouse interface {
	// Shortcut for Mouse.move`], [`method: Mouse.down`], [`method: Mouse.up().
	Click(x, y float64, options ...MouseClickOptions) error
	// Shortcut for Mouse.move`], [`method: Mouse.down`], [`method: Mouse.up`], [`method: Mouse.down() and
	// Mouse.up().
	Dblclick(x, y float64, options ...MouseDblclickOptions) error
	// Dispatches a `mousedown` event.
	Down(options ...MouseDownOptions) error
	// Dispatches a `mousemove` event.
	Move(x float64, y float64, options ...MouseMoveOptions) error
	// Dispatches a `mouseup` event.
	Up(options ...MouseUpOptions) error
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
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now checked or unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	// Shortcut for main frame's Frame.setChecked().
	SetChecked(selector string, checked bool, options ...FrameSetCheckedOptions) error
	Mouse() Mouse
	Keyboard() Keyboard
	Touchscreen() Touchscreen
	// Adds a script which would be evaluated in one of the following scenarios:
	// - Whenever the page is navigated.
	// - Whenever the child frame is attached or navigated. In this case, the script is evaluated in the context of the newly
	// attached frame.
	// The script is evaluated after the document was created but before any of its scripts were run. This is useful to amend
	// the JavaScript environment, e.g. to seed `Math.random`.
	// An example of overriding `Math.random` before the page loads:
	// > NOTE: The order of evaluation of multiple scripts installed via BrowserContext.addInitScript() and
	// Page.addInitScript() is not defined.
	AddInitScript(script PageAddInitScriptOptions) error
	// Adds a `<script>` tag into the page with the desired url or content. Returns the added tag when the script's onload
	// fires or when the script content was injected into frame.
	// Shortcut for main frame's Frame.addScriptTag().
	AddScriptTag(options PageAddScriptTagOptions) (ElementHandle, error)
	// Adds a `<link rel="stylesheet">` tag into the page with the desired url or a `<style type="text/css">` tag with the
	// content. Returns the added tag when the stylesheet's onload fires or when the CSS content was injected into frame.
	// Shortcut for main frame's Frame.addStyleTag().
	AddStyleTag(options PageAddStyleTagOptions) (ElementHandle, error)
	// Brings page to front (activates tab).
	BringToFront() error
	// This method checks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Ensure that matched element is a checkbox or a radio input. If not, this method throws. If the element is already
	// checked, this method returns immediately.
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now checked. If not, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	// Shortcut for main frame's Frame.check().
	Check(selector string, options ...FrameCheckOptions) error
	// This method clicks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to click in the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	// Shortcut for main frame's Frame.click().
	Click(selector string, options ...PageClickOptions) error
	// If `runBeforeUnload` is `false`, does not run any unload handlers and waits for the page to be closed. If
	// `runBeforeUnload` is `true` the method will run unload handlers, but will **not** wait for the page to close.
	// By default, `page.close()` **does not** run `beforeunload` handlers.
	// > NOTE: if `runBeforeUnload` is passed as true, a `beforeunload` dialog might be summoned and should be handled manually
	// via [`event: Page.dialog`] event.
	Close(options ...PageCloseOptions) error
	// Gets the full HTML contents of the page, including the doctype.
	Content() (string, error)
	// Get the browser context that the page belongs to.
	Context() BrowserContext
	// This method double clicks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to double click in the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set. Note that if the
	// first click of the `dblclick()` triggers a navigation event, this method will throw.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	// > NOTE: `page.dblclick()` dispatches two `click` events and a single `dblclick` event.
	// Shortcut for main frame's Frame.dblclick().
	Dblclick(expression string, options ...FrameDblclickOptions) error
	// The snippet below dispatches the `click` event on the element. Regardless of the visibility state of the element,
	// `click` is dispatched. This is equivalent to calling
	// [element.click()](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click).
	// Under the hood, it creates an instance of an event based on the given `type`, initializes it with `eventInit` properties
	// and dispatches it on the element. Events are `composed`, `cancelable` and bubble by default.
	// Since `eventInit` is event-specific, please refer to the events documentation for the lists of initial properties:
	// - [DragEvent](https://developer.mozilla.org/en-US/docs/Web/API/DragEvent/DragEvent)
	// - [FocusEvent](https://developer.mozilla.org/en-US/docs/Web/API/FocusEvent/FocusEvent)
	// - [KeyboardEvent](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/KeyboardEvent)
	// - [MouseEvent](https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/MouseEvent)
	// - [PointerEvent](https://developer.mozilla.org/en-US/docs/Web/API/PointerEvent/PointerEvent)
	// - [TouchEvent](https://developer.mozilla.org/en-US/docs/Web/API/TouchEvent/TouchEvent)
	// - [Event](https://developer.mozilla.org/en-US/docs/Web/API/Event/Event)
	// You can also specify `JSHandle` as the property value if you want live objects to be passed into the event:
	DispatchEvent(selector string, typ string, options ...PageDispatchEventOptions) error
	// The method adds a function called `name` on the `window` object of every frame in this page. When called, the function
	// executes `callback` and returns a [Promise] which resolves to the return value of `callback`. If the `callback` returns
	// a [Promise], it will be awaited.
	// The first argument of the `callback` function contains information about the caller: `{ browserContext: BrowserContext,
	// page: Page, frame: Frame }`.
	// See BrowserContext.exposeBinding() for the context-wide version.
	// > NOTE: Functions installed via Page.exposeBinding() survive navigations.
	// An example of exposing page URL to all frames in a page:
	// An example of passing an element handle:
	ExposeBinding(name string, binding BindingCallFunction, handle ...bool) error
	// The method adds a function called `name` on the `window` object of every frame in the page. When called, the function
	// executes `callback` and returns a [Promise] which resolves to the return value of `callback`.
	// If the `callback` returns a [Promise], it will be awaited.
	// See BrowserContext.exposeFunction() for context-wide exposed function.
	// > NOTE: Functions installed via Page.exposeFunction() survive navigations.
	// An example of adding a `sha256` function to the page:
	ExposeFunction(name string, binding ExposedFunction) error
	// This method changes the `CSS media type` through the `media` argument, and/or the `'prefers-colors-scheme'` media
	// feature, using the `colorScheme` argument.
	EmulateMedia(options ...PageEmulateMediaOptions) error
	// Returns the value of the `expression` invocation.
	// If the function passed to the Page.evaluate`] returns a [Promise], then [`method: Page.evaluate() would wait
	// for the promise to resolve and return its value.
	// If the function passed to the Page.evaluate() returns a non-[Serializable] value, then
	// Page.evaluate() resolves to `undefined`. Playwright also supports transferring some additional values that are
	// not serializable by `JSON`: `-0`, `NaN`, `Infinity`, `-Infinity`.
	// Passing argument to `expression`:
	// A string can also be passed in instead of a function:
	// `ElementHandle` instances can be passed as an argument to the Page.evaluate():
	// Shortcut for main frame's Frame.evaluate().
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	// Returns the value of the `expression` invocation as a `JSHandle`.
	// The only difference between Page.evaluate`] and [`method: Page.evaluateHandle() is that
	// Page.evaluateHandle() returns `JSHandle`.
	// If the function passed to the Page.evaluateHandle`] returns a [Promise], then [`method: Page.evaluateHandle()
	// would wait for the promise to resolve and return its value.
	// A string can also be passed in instead of a function:
	// `JSHandle` instances can be passed as an argument to the Page.evaluateHandle():
	EvaluateHandle(expression string, options ...interface{}) (JSHandle, error)
	// > NOTE: This method does not wait for the element to pass actionability checks and therefore can lead to the flaky
	// tests. Use Locator.evaluate(), other `Locator` helper methods or web-first assertions instead.
	// The method finds an element matching the specified selector within the page and passes it as a first argument to
	// `expression`. If no elements match the selector, the method throws an error. Returns the value of `expression`.
	// If `expression` returns a [Promise], then Page.evalOnSelector() would wait for the promise to resolve and
	// return its value.
	// Examples:
	// Shortcut for main frame's Frame.evalOnSelector().
	EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
	// > NOTE: In most cases, Locator.evaluateAll(), other `Locator` helper methods and web-first assertions do a
	// better job.
	// The method finds all elements matching the specified selector within the page and passes an array of matched elements as
	// a first argument to `expression`. Returns the result of `expression` invocation.
	// If `expression` returns a [Promise], then Page.evalOnSelectorAll() would wait for the promise to resolve and
	// return its value.
	// Examples:
	EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
	ExpectConsoleMessage(cb func() error) (ConsoleMessage, error)
	ExpectDownload(cb func() error) (Download, error)
	ExpectEvent(event string, cb func() error, predicates ...interface{}) (interface{}, error)
	ExpectFileChooser(cb func() error) (FileChooser, error)
	ExpectLoadState(state string, cb func() error) error
	ExpectNavigation(cb func() error, options ...PageWaitForNavigationOptions) (Response, error)
	ExpectPopup(cb func() error) (Page, error)
	ExpectRequest(url interface{}, cb func() error, options ...interface{}) (Request, error)
	ExpectResponse(url interface{}, cb func() error, options ...interface{}) (Response, error)
	ExpectWorker(cb func() error) (Worker, error)
	ExpectedDialog(cb func() error) (Dialog, error)
	// This method waits for an element matching `selector`, waits for [actionability](./actionability.md) checks, focuses the
	// element, fills it and triggers an `input` event after filling. Note that you can pass an empty string to clear the input
	// field.
	// If the target element is not an `<input>`, `<textarea>` or `[contenteditable]` element, this method throws an error.
	// However, if the element is inside the `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), the control will be filled
	// instead.
	// To send fine-grained keyboard events, use Page.type().
	// Shortcut for main frame's Frame.fill().
	Fill(selector, text string, options ...FrameFillOptions) error
	// This method fetches an element with `selector` and focuses it. If there's no element matching `selector`, the method
	// waits until a matching element appears in the DOM.
	// Shortcut for main frame's Frame.focus().
	Focus(expression string, options ...FrameFocusOptions) error
	// Returns frame matching the specified criteria. Either `name` or `url` must be specified.
	Frame(options PageFrameOptions) Frame
	// An array of all frames attached to the page.
	Frames() []Frame
	// Returns element attribute value.
	GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error)
	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of the
	// last redirect. If can not go back, returns `null`.
	// Navigate to the previous page in history.
	GoBack(options ...PageGoBackOptions) (Response, error)
	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of the
	// last redirect. If can not go forward, returns `null`.
	// Navigate to the next page in history.
	GoForward(options ...PageGoForwardOptions) (Response, error)
	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of the
	// last redirect.
	// The method will throw an error if:
	// - there's an SSL error (e.g. in case of self-signed certificates).
	// - target URL is invalid.
	// - the `timeout` is exceeded during navigation.
	// - the remote server does not respond or is unreachable.
	// - the main resource failed to load.
	// The method will not throw an error when any valid HTTP status code is returned by the remote server, including 404 "Not
	// Found" and 500 "Internal Server Error".  The status code for such responses can be retrieved by calling
	// Response.status().
	// > NOTE: The method either throws an error or returns a main resource response. The only exceptions are navigation to
	// `about:blank` or navigation to the same URL with a different hash, which would succeed and return `null`.
	// > NOTE: Headless mode doesn't support navigation to a PDF document. See the
	// [upstream issue](https://bugs.chromium.org/p/chromium/issues/detail?id=761295).
	// Shortcut for main frame's Frame.goto()
	Goto(url string, options ...PageGotoOptions) (Response, error)
	// This method hovers over an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to hover over the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	// Shortcut for main frame's Frame.hover().
	Hover(selector string, options ...PageHoverOptions) error
	// Returns `element.innerHTML`.
	InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error)
	// Returns `element.innerText`.
	InnerText(selector string, options ...PageInnerTextOptions) (string, error)
	// Indicates that the page has been closed.
	IsClosed() bool
	// Returns whether the element is checked. Throws if the element is not a checkbox or radio input.
	IsChecked(selector string, options ...FrameIsCheckedOptions) (bool, error)
	// Returns whether the element is disabled, the opposite of [enabled](./actionability.md#enabled).
	IsDisabled(selector string, options ...FrameIsDisabledOptions) (bool, error)
	// Returns whether the element is [editable](./actionability.md#editable).
	IsEditable(selector string, options ...FrameIsEditableOptions) (bool, error)
	// Returns whether the element is [enabled](./actionability.md#enabled).
	IsEnabled(selector string, options ...FrameIsEnabledOptions) (bool, error)
	// Returns whether the element is hidden, the opposite of [visible](./actionability.md#visible).  `selector` that does not
	// match any elements is considered hidden.
	IsHidden(selector string, options ...FrameIsHiddenOptions) (bool, error)
	// Returns whether the element is [visible](./actionability.md#visible). `selector` that does not match any elements is
	// considered not visible.
	IsVisible(selector string, options ...FrameIsVisibleOptions) (bool, error)
	// The page's main frame. Page is guaranteed to have a main frame which persists during navigations.
	MainFrame() Frame
	// Returns the opener for popup pages and `null` for others. If the opener has been closed already the returns `null`.
	Opener() (Page, error)
	// Returns the PDF buffer.
	// > NOTE: Generating a pdf is currently only supported in Chromium headless.
	// `page.pdf()` generates a pdf of the page with `print` css media. To generate a pdf with `screen` media, call
	// Page.emulateMedia() before calling `page.pdf()`:
	// > NOTE: By default, `page.pdf()` generates a pdf with modified colors for printing. Use the
	// [`-webkit-print-color-adjust`](https://developer.mozilla.org/en-US/docs/Web/CSS/-webkit-print-color-adjust) property to
	// force rendering of exact colors.
	// The `width`, `height`, and `margin` options accept values labeled with units. Unlabeled values are treated as pixels.
	// A few examples:
	// - `page.pdf({width: 100})` - prints with width set to 100 pixels
	// - `page.pdf({width: '100px'})` - prints with width set to 100 pixels
	// - `page.pdf({width: '10cm'})` - prints with width set to 10 centimeters.
	// All possible units are:
	// - `px` - pixel
	// - `in` - inch
	// - `cm` - centimeter
	// - `mm` - millimeter
	// The `format` options are:
	// - `Letter`: 8.5in x 11in
	// - `Legal`: 8.5in x 14in
	// - `Tabloid`: 11in x 17in
	// - `Ledger`: 17in x 11in
	// - `A0`: 33.1in x 46.8in
	// - `A1`: 23.4in x 33.1in
	// - `A2`: 16.54in x 23.4in
	// - `A3`: 11.7in x 16.54in
	// - `A4`: 8.27in x 11.7in
	// - `A5`: 5.83in x 8.27in
	// - `A6`: 4.13in x 5.83in
	// > NOTE: `headerTemplate` and `footerTemplate` markup have the following limitations: > 1. Script tags inside templates
	// are not evaluated. > 2. Page styles are not visible inside templates.
	PDF(options ...PagePdfOptions) ([]byte, error)
	// Focuses the element, and then uses Keyboard.down`] and [`method: Keyboard.up().
	// `key` can specify the intended [keyboardEvent.key](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key)
	// value or a single character to generate the text for. A superset of the `key` values can be found
	// [here](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values). Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`, etc.
	// Following modification shortcuts are also supported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When specified with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	Press(selector, key string, options ...PagePressOptions) error
	// > NOTE: The use of `ElementHandle` is discouraged, use `Locator` objects and web-first assertions instead.
	// The method finds an element matching the specified selector within the page. If no elements match the selector, the
	// return value resolves to `null`. To wait for an element on the page, use Locator.waitFor().
	// Shortcut for main frame's Frame.querySelector().
	QuerySelector(selector string) (ElementHandle, error)
	// > NOTE: The use of `ElementHandle` is discouraged, use `Locator` objects and web-first assertions instead.
	// The method finds all elements matching the specified selector within the page. If no elements match the selector, the
	// return value resolves to `[]`.
	// Shortcut for main frame's Frame.querySelectorAll().
	QuerySelectorAll(selector string) ([]ElementHandle, error)
	// This method reloads the current page, in the same way as if the user had triggered a browser refresh. Returns the main
	// resource response. In case of multiple redirects, the navigation will resolve with the response of the last redirect.
	Reload(options ...PageReloadOptions) (Response, error)
	// Routing provides the capability to modify network requests that are made by a page.
	// Once routing is enabled, every request matching the url pattern will stall unless it's continued, fulfilled or aborted.
	// > NOTE: The handler will only be called for the first url if the response is a redirect.
	// > NOTE: Page.route() will not intercept requests intercepted by Service Worker. See
	// [this](https://github.com/microsoft/playwright/issues/1090) issue. We recommend disabling Service Workers when using
	// request interception. Via `await context.addInitScript(() => delete window.navigator.serviceWorker);`
	// An example of a naive handler that aborts all image requests:
	// or the same snippet using a regex pattern instead:
	// It is possible to examine the request to decide the route action. For example, mocking all requests that contain some
	// post data, and leaving all other requests as is:
	// Page routes take precedence over browser context routes (set up with BrowserContext.route()) when request
	// matches both handlers.
	// To remove a route with its handler you can use Page.unroute().
	// > NOTE: Enabling routing disables http cache.
	Route(url interface{}, handler routeHandler) error
	// Returns the buffer with the captured screenshot.
	Screenshot(options ...PageScreenshotOptions) ([]byte, error)
	// This method waits for an element matching `selector`, waits for [actionability](./actionability.md) checks, waits until
	// all specified options are present in the `<select>` element and selects these options.
	// If the target element is not a `<select>` element, this method throws an error. However, if the element is inside the
	// `<label>` element that has an associated
	// [control](https://developer.mozilla.org/en-US/docs/Web/API/HTMLLabelElement/control), the control will be used instead.
	// Returns the array of option values that have been successfully selected.
	// Triggers a `change` and `input` event once all the provided options have been selected.
	// Shortcut for main frame's Frame.selectOption().
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
	// > NOTE: Page.setDefaultNavigationTimeout`] takes priority over [`method: Page.setDefaultTimeout(),
	// BrowserContext.setDefaultTimeout`] and [`method: BrowserContext.setDefaultNavigationTimeout().
	SetDefaultNavigationTimeout(timeout float64)
	// This setting will change the default maximum time for all the methods accepting `timeout` option.
	// > NOTE: Page.setDefaultNavigationTimeout`] takes priority over [`method: Page.setDefaultTimeout().
	SetDefaultTimeout(timeout float64)
	// The extra HTTP headers will be sent with every request the page initiates.
	// > NOTE: Page.setExtraHTTPHeaders() does not guarantee the order of headers in the outgoing requests.
	SetExtraHTTPHeaders(headers map[string]string) error
	// This method expects `selector` to point to an
	// [input element](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input).
	// Sets the value of the file input to these file paths or files. If some of the `filePaths` are relative paths, then they
	// are resolved relative to the the current working directory. For empty array, clears the selected files.
	SetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) error
	// In the case of multiple pages in a single browser, each page can have its own viewport size. However,
	// Browser.newContext() allows to set viewport size (and more) for all pages in the context at once.
	// Page.setViewportSize() will resize the page. A lot of websites don't expect phones to change size, so you
	// should set the viewport size before navigating to the page. Page.setViewportSize() will also reset `screen`
	// size, use Browser.newContext() with `screen` and `viewport` parameters if you need better control of these
	// properties.
	SetViewportSize(width, height int) error
	// This method taps an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.touchscreen`] to tap the center of the element, or the specified `position`.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	// > NOTE: Page.tap() requires that the `hasTouch` option of the browser context be set to true.
	// Shortcut for main frame's Frame.tap().
	Tap(selector string, options ...FrameTapOptions) error
	// Returns `element.textContent`.
	TextContent(selector string, options ...FrameTextContentOptions) (string, error)
	// Returns the page's title. Shortcut for main frame's Frame.title().
	Title() (string, error)
	// Sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the text. `page.type` can be used to send
	// fine-grained keyboard events. To fill values in form fields, use Page.fill().
	// To press a special key, like `Control` or `ArrowDown`, use Keyboard.press().
	// Shortcut for main frame's Frame.type().
	Type(selector, text string, options ...PageTypeOptions) error
	// Shortcut for main frame's Frame.url().
	URL() string
	// This method unchecks an element matching `selector` by performing the following steps:
	// 1. Find an element matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// 1. Ensure that matched element is a checkbox or a radio input. If not, this method throws. If the element is already
	// unchecked, this method returns immediately.
	// 1. Wait for [actionability](./actionability.md) checks on the matched element, unless `force` option is set. If the
	// element is detached during the checks, the whole action is retried.
	// 1. Scroll the element into view if needed.
	// 1. Use [`property: Page.mouse`] to click in the center of the element.
	// 1. Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// 1. Ensure that the element is now unchecked. If not, this method throws.
	// When all steps combined have not finished during the specified `timeout`, this method throws a `TimeoutError`. Passing
	// zero timeout disables this.
	// Shortcut for main frame's Frame.uncheck().
	Uncheck(selector string, options ...FrameUncheckOptions) error
	// Removes a route created with Page.route(). When `handler` is not specified, removes all routes for the `url`.
	Unroute(url interface{}, handler ...routeHandler) error
	// Video object associated with this page.
	Video() Video
	ViewportSize() ViewportSize
	// Waits for event to fire and passes its value into the predicate function. Returns when the predicate returns truthy
	// value. Will throw an error if the page is closed before the event is fired. Returns the event data value.
	WaitForEvent(event string, predicate ...interface{}) interface{}
	// Returns when the `expression` returns a truthy value. It resolves to a JSHandle of the truthy value.
	// The Page.waitForFunction() can be used to observe viewport size change:
	// To pass an argument to the predicate of Page.waitForFunction() function:
	// Shortcut for main frame's Frame.waitForFunction().
	WaitForFunction(expression string, arg interface{}, options ...FrameWaitForFunctionOptions) (JSHandle, error)
	// Returns when the required load state has been reached.
	// This resolves when the page reaches a required load state, `load` by default. The navigation must have been committed
	// when this method is called. If current document has already reached the required state, resolves immediately.
	// Shortcut for main frame's Frame.waitForLoadState().
	WaitForLoadState(state ...string)
	// Waits for the main frame navigation and returns the main resource response. In case of multiple redirects, the
	// navigation will resolve with the response of the last redirect. In case of navigation to a different anchor or
	// navigation due to History API usage, the navigation will resolve with `null`.
	// This resolves when the page navigates to a new URL or reloads. It is useful for when you run code which will indirectly
	// cause the page to navigate. e.g. The click target has an `onclick` handler that triggers navigation from a `setTimeout`.
	// Consider this example:
	// > NOTE: Usage of the [History API](https://developer.mozilla.org/en-US/docs/Web/API/History_API) to change the URL is
	// considered a navigation.
	// Shortcut for main frame's Frame.waitForNavigation().
	WaitForNavigation(options ...PageWaitForNavigationOptions) (Response, error)
	// Waits for the matching request and returns it. See [waiting for event](./events.md#waiting-for-event) for more details
	// about events.
	WaitForRequest(url interface{}, options ...interface{}) Request
	// Returns the matched response. See [waiting for event](./events.md#waiting-for-event) for more details about events.
	WaitForResponse(url interface{}, options ...interface{}) Response
	// Returns when element specified by selector satisfies `state` option. Returns `null` if waiting for `hidden` or
	// `detached`.
	// > NOTE: Playwright automatically waits for element to be ready before performing an action. Using `Locator` objects and
	// web-first assertions make the code wait-for-selector-free.
	// Wait for the `selector` to satisfy `state` option (either appear/disappear from dom, or become visible/hidden). If at
	// the moment of calling the method `selector` already satisfies the condition, the method will return immediately. If the
	// selector doesn't satisfy the condition for the `timeout` milliseconds, the function will throw.
	// This method works across navigations:
	WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandle, error)
	// Waits for the given `timeout` in milliseconds.
	// Note that `page.waitForTimeout()` should only be used for debugging. Tests using the timer in production are going to be
	// flaky. Use signals such as network events, selectors becoming visible and others instead.
	// Shortcut for main frame's Frame.waitForTimeout().
	WaitForTimeout(timeout float64)
	// This method returns all of the dedicated [WebWorkers](https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API)
	// associated with the page.
	// > NOTE: This does not contain ServiceWorkers
	Workers() []Worker
	DragAndDrop(source, target string, options ...FrameDragAndDropOptions) error
	// Pauses script execution. Playwright will stop executing the script and wait for the user to either press 'Resume' button
	// in the page overlay or to call `playwright.resume()` in the DevTools console.
	// User can inspect selectors or perform manual steps while paused. Resume will continue running the original script from
	// the place it was paused.
	// > NOTE: This method requires Playwright to be started in a headed mode, with a falsy `headless` value in the
	// BrowserType.launch().
	Pause() error
	// Returns `input.value` for the selected `<input>` or `<textarea>` or `<select>` element. Throws for non-input elements.
	InputValue(selector string, options ...FrameInputValueOptions) (string, error)
	// Waits for the main frame to navigate to the given URL.
	// Shortcut for main frame's Frame.waitForURL().
	WaitForURL(url string, options ...FrameWaitForURLOptions) error
}

// Whenever the page sends a request for a network resource the following sequence of events are emitted by `Page`:
// - [`event: Page.request`] emitted when the request is issued by the page.
// - [`event: Page.response`] emitted when/if the response status and headers are received for the request.
// - [`event: Page.requestFinished`] emitted when the response body is downloaded and the request is complete.
// If request fails at some point, then instead of `'requestfinished'` event (and possibly instead of 'response' event),
// the  [`event: Page.requestFailed`] event is emitted.
// > NOTE: HTTP Error responses, such as 404 or 503, are still successful responses from HTTP standpoint, so request will
// complete with `'requestfinished'` event.
// If request gets a 'redirect' response, the request is successfully finished with the 'requestfinished' event, and a new
// request is  issued to a redirected url.
type Request interface {
	// An object with all the request HTTP headers associated with this request. The header names are lower-cased.
	AllHeaders() (map[string]string, error)
	// An array with all the request HTTP headers associated with this request. Unlike Request.allHeaders(), header
	// names are NOT lower-cased. Headers with multiple entries, such as `Set-Cookie`, appear in the array multiple times.
	HeadersArray() (HeadersArray, error)
	// Returns the value of the header matching the name. The name is case insensitive.
	HeaderValue(name string) (string, error)
	HeaderValues(name string) ([]string, error)
	// The method returns `null` unless this request has failed, as reported by `requestfailed` event.
	// Example of logging of all the failed requests:
	Failure() *RequestFailure
	// Returns the `Frame` that initiated this request.
	Frame() Frame
	// **DEPRECATED** Incomplete list of headers as seen by the rendering engine. Use Request.allHeaders() instead.
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
	// When the server responds with a redirect, Playwright creates a new `Request` object. The two requests are connected by
	// `redirectedFrom()` and `redirectedTo()` methods. When multiple server redirects has happened, it is possible to
	// construct the whole redirect chain by repeatedly calling `redirectedFrom()`.
	// For example, if the website `http://example.com` redirects to `https://example.com`:
	// If the website `https://google.com` has no redirects:
	RedirectedFrom() Request
	// New request issued by the browser if the server responded with redirect.
	// This method is the opposite of Request.redirectedFrom():
	RedirectedTo() Request
	// Contains the request's resource type as it was perceived by the rendering engine. ResourceType will be one of the
	// following: `document`, `stylesheet`, `image`, `media`, `font`, `script`, `texttrack`, `xhr`, `fetch`, `eventsource`,
	// `websocket`, `manifest`, `other`.
	ResourceType() string
	// Returns the matching `Response` object, or `null` if the response was not received due to error.
	Response() (Response, error)
	// Returns resource timing information for given request. Most of the timing values become available upon the response,
	// `responseEnd` becomes available when request finishes. Find more information at
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
	// An array with all the request HTTP headers associated with this response. Unlike Response.allHeaders(), header
	// names are NOT lower-cased. Headers with multiple entries, such as `Set-Cookie`, appear in the array multiple times.
	HeadersArray() (HeadersArray, error)
	// Returns the value of the header matching the name. The name is case insensitive. If multiple headers have the same name
	// (except `set-cookie`), they are returned as a list separated by `, `. For `set-cookie`, the `\n` separator is used. If
	// no headers are found, `null` is returned.
	HeaderValue(name string) (string, error)
	// Returns all values of the headers matching the name, for example `set-cookie`. The name is case insensitive.
	HeaderValues(name string) ([]string, error)
	// Returns the buffer with response body.
	Body() ([]byte, error)
	// Waits for this response to finish, returns always `null`.
	Finished()
	// Returns the `Frame` that initiated this response.
	Frame() Frame
	// **DEPRECATED** Incomplete list of headers as seen by the rendering engine. Use Response.allHeaders() instead.
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

// Whenever a network route is set up with Page.route`] or [`method: BrowserContext.route(), the `Route` object
// allows to handle the route.
type Route interface {
	// Aborts the route's request.
	Abort(errorCode ...string) error
	// Continues route's request with optional overrides.
	Continue(options ...RouteContinueOptions) error
	// Fulfills route's request with given response.
	// An example of fulfilling all requests with 404 responses:
	// An example of serving static file:
	Fulfill(options RouteFulfillOptions) error
	// A request to be routed.
	Request() Request
}

// The Touchscreen class operates in main-frame CSS pixels relative to the top-left corner of the viewport. Methods on the
// touchscreen can only be used in browser contexts that have been initialized with `hasTouch` set to true.
type Touchscreen interface {
	// Dispatches a `touchstart` and `touchend` event with a single touch at the position (`x`,`y`).
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
	WaitForEvent(event string, predicate ...interface{}) interface{}
}

// When browser context is created with the `recordVideo` option, each page has a video object associated with it.
type Video interface {
	// Returns the file system path this video will be recorded to. The video is guaranteed to be written to the filesystem
	// upon closing the browser context. This method throws when connected remotely.
	Path() (string, error)
	// Deletes the video file. Will wait for the video to finish if necessary.
	Delete() error
	// Saves the video to a user-specified path. It is safe to call this method while the video is still in progress, or after
	// the page has closed. This method waits until the page is closed and the video is fully saved.
	SaveAs(path string) error
}

// The Worker class represents a [WebWorker](https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API). `worker`
// event is emitted on the page object to signal a worker creation. `close` event is emitted on the worker object when the
// worker is gone.
type Worker interface {
	EventEmitter
	// Returns the return value of `expression`.
	// If the function passed to the Worker.evaluate`] returns a [Promise], then [`method: Worker.evaluate() would
	// wait for the promise to resolve and return its value.
	// If the function passed to the Worker.evaluate() returns a non-[Serializable] value, then
	// Worker.evaluate() returns `undefined`. Playwright also supports transferring some additional values that are
	// not serializable by `JSON`: `-0`, `NaN`, `Infinity`, `-Infinity`.
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `expression` as a `JSHandle`.
	// The only difference between Worker.evaluate`] and [`method: Worker.evaluateHandle() is that
	// Worker.evaluateHandle() returns `JSHandle`.
	// If the function passed to the Worker.evaluateHandle() returns a [Promise], then
	// Worker.evaluateHandle() would wait for the promise to resolve and return its value.
	EvaluateHandle(expression string, options ...interface{}) (JSHandle, error)
	URL() string
	WaitForEvent(event string, predicate ...interface{}) interface{}
	ExpectEvent(event string, cb func() error, predicates ...interface{}) (interface{}, error)
}
