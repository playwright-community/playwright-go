package playwright

type BindingCall interface {
	Call(f BindingCallFunction)
}

// A Browser is created when Playwright connects to a browser instance, either through browserType.launch([options]) or
// browserType.connect(params).
// An example of using a Browser to create a Page:
// See ChromiumBrowser, FirefoxBrowser and WebKitBrowser for browser-specific features. Note that
// browserType.connect(params) and browserType.launch([options]) always return a specific browser instance, based on the browser
// being connected to or launched.
type Browser interface {
	EventEmitter
	// In case this browser is obtained using browserType.launch([options]), closes the browser and all of its pages (if any were
	// opened).
	// In case this browser is obtained using browserType.connect(params), clears all created contexts belonging to this browser
	// and disconnects from the browser server.
	// The Browser object itself is considered to be disposed and cannot be used anymore.
	Close() error
	// Returns an array of all open browser contexts. In a newly created browser, this will return zero browser contexts.
	Contexts() []BrowserContext
	// Indicates that the browser is connected.
	IsConnected() bool
	// Creates a new browser context. It won't share cookies/cache with other browser contexts.
	NewContext(options ...BrowserNewContextOptions) (BrowserContext, error)
	// Creates a new page in a new browser context. Closing this page will close the context as well.
	// This is a convenience API that should only be used for the single-page scenarios and short snippets. Production code and
	// testing frameworks should explicitly create browser.newContext([options]) followed by the browserContext.newPage() to
	// control their exact life times.
	NewPage(options ...BrowserNewContextOptions) (Page, error)
	// Returns the browser version.
	Version() string
}

// BrowserContexts provide a way to operate multiple independent browser sessions.
// If a page opens another page, e.g. with a `window.open` call, the popup will belong to the parent page's browser
// context.
// Playwright allows creation of "incognito" browser contexts with `browser.newContext()` method. "Incognito" browser
// contexts don't write any browsing data to disk.
type BrowserContext interface {
	EventEmitter
	// Adds cookies into this browser context. All pages within this context will have these cookies installed. Cookies can be
	// obtained via browserContext.cookies([urls]).
	AddCookies(cookies ...SetNetworkCookieParam) error
	// Adds a script which would be evaluated in one of the following scenarios:
	// Whenever a page is created in the browser context or is navigated.
	// Whenever a child frame is attached or navigated in any page in the browser context. In this case, the script is evaluated in the context of the newly attached frame.
	// The script is evaluated after the document was created but before any of its scripts were run. This is useful to amend
	// the JavaScript environment, e.g. to seed `Math.random`.
	// An example of overriding `Math.random` before the page loads:
	// **NOTE** The order of evaluation of multiple scripts installed via browserContext.addInitScript(script[, arg]) and
	// page.addInitScript(script[, arg]) is not defined.
	AddInitScript(options BrowserContextAddInitScriptOptions) error
	// Returns the browser instance of the context. If it was launched as a persistent context null gets returned.
	Browser() Browser
	// Clears context cookies.
	ClearCookies() error
	// Clears all permission overrides for the browser context.
	ClearPermissions() error
	// Closes the browser context. All the pages that belong to the browser context will be closed.
	// **NOTE** the default browser context cannot be closed.
	Close() error
	// If no URLs are specified, this method returns all cookies. If URLs are specified, only cookies that affect those URLs
	// are returned.
	Cookies(urls ...string) ([]*NetworkCookie, error)
	ExpectEvent(event string, cb func() error) (interface{}, error)
	// The method adds a function called `name` on the `window` object of every frame in every page in the context. When
	// called, the function executes `playwrightBinding` in Node.js and returns a Promise which resolves to the return value
	// of `playwrightBinding`. If the `playwrightBinding` returns a Promise, it will be awaited.
	// The first argument of the `playwrightBinding` function contains information about the caller: `{ browserContext: BrowserContext, page: Page, frame: Frame }`.
	// See page.exposeBinding(name, playwrightBinding[, options]) for page-only version.
	// An example of exposing page URL to all frames in all pages in the context:
	// An example of passing an element handle:
	ExposeBinding(name string, binding BindingCallFunction, handle ...bool) error
	// The method adds a function called `name` on the `window` object of every frame in every page in the context. When
	// called, the function executes `playwrightFunction` in Node.js and returns a Promise which resolves to the return value
	// of `playwrightFunction`.
	// If the `playwrightFunction` returns a Promise, it will be awaited.
	// See page.exposeFunction(name, playwrightFunction) for page-only version.
	// An example of adding an `md5` function to all pages in the context:
	ExposeFunction(name string, binding ExposedFunction) error
	// Grants specified permissions to the browser context. Only grants corresponding permissions to the given origin if
	// specified.
	GrantPermissions(permissions []string, options ...BrowserContextGrantPermissionsOptions) error
	// Creates a new page in the browser context.
	NewPage(options ...BrowserNewPageOptions) (Page, error)
	// Returns all open pages in the context. Non visible pages, such as `"background_page"`, will not be listed here. You can
	// find them using chromiumBrowserContext.backgroundPages().
	Pages() []Page
	// This setting will change the default maximum navigation time for the following methods and related shortcuts:
	// page.goBack([options])
	// page.goForward([options])
	// page.goto(url[, options])
	// page.reload([options])
	// page.setContent(html[, options])
	// page.waitForNavigation([options])
	// **NOTE** page.setDefaultNavigationTimeout(timeout) and page.setDefaultTimeout(timeout) take priority over
	// browserContext.setDefaultNavigationTimeout(timeout).
	SetDefaultNavigationTimeout(timeout int)
	// This setting will change the default maximum time for all the methods accepting `timeout` option.
	// **NOTE** page.setDefaultNavigationTimeout(timeout), page.setDefaultTimeout(timeout) and
	// browserContext.setDefaultNavigationTimeout(timeout) take priority over browserContext.setDefaultTimeout(timeout).
	SetDefaultTimeout(timeout int)
	// The extra HTTP headers will be sent with every request initiated by any page in the context. These headers are merged
	// with page-specific extra HTTP headers set with page.setExtraHTTPHeaders(headers). If page overrides a particular header,
	// page-specific header value will be used instead of the browser context header value.
	// **NOTE** `browserContext.setExtraHTTPHeaders` does not guarantee the order of headers in the outgoing requests.
	SetExtraHTTPHeaders(headers map[string]string) error
	// Sets the context's geolocation. Passing `null` or `undefined` emulates position unavailable.
	// **NOTE** Consider using browserContext.grantPermissions(permissions[, options]) to grant permissions for the browser context pages to
	// read its geolocation.
	SetGeolocation(gelocation *SetGeolocationOptions) error
	ResetGeolocation() error
	// Routing provides the capability to modify network requests that are made by any page in the browser context. Once route
	// is enabled, every request matching the url pattern will stall unless it's continued, fulfilled or aborted.
	// An example of a naïve handler that aborts all image requests:
	// or the same snippet using a regex pattern instead:
	// Page routes (set up with page.route(url, handler)) take precedence over browser context routes when request matches both
	// handlers.
	// **NOTE** Enabling routing disables http cache.
	Route(url interface{}, handler routeHandler) error
	SetOffline(offline bool) error
	// Removes a route created with browserContext.route(url, handler). When `handler` is not specified, removes all routes for the
	// `url`.
	Unroute(url interface{}, handler routeHandler) error
	// Waits for event to fire and passes its value into the predicate function. Resolves when the predicate returns truthy
	// value. Will throw an error if the context closes before the event is fired. Returns the event data value.
	WaitForEvent(event string, predicate ...interface{}) interface{}
}

// BrowserType provides methods to launch a specific browser instance or connect to an existing one. The following is a
// typical example of using Playwright to drive automation:
type BrowserType interface {
	// A path where Playwright expects to find a bundled browser executable.
	ExecutablePath() string
	// Returns the browser instance.
	// You can use `ignoreDefaultArgs` to filter out `--mute-audio` from default arguments:
	// **Chromium-only** Playwright can also be used to control the Chrome browser, but it works best with the version of
	// Chromium it is bundled with. There is no guarantee it will work with any other version. Use `executablePath` option with
	// extreme caution.
	// If Google Chrome (rather than Chromium) is preferred, a Chrome
	// Canary or Dev
	// Channel build is suggested.
	// In browserType.launch([options]) above, any mention of Chromium also applies to Chrome.
	// See `this article` for
	// a description of the differences between Chromium and Chrome. `This article` describes
	// some differences for Linux users.
	Launch(options ...BrowserTypeLaunchOptions) (Browser, error)
	// Returns the persistent browser context instance.
	// Launches browser that uses persistent storage located at `userDataDir` and returns the only context. Closing this
	// context will automatically close the browser.
	LaunchPersistentContext(userDataDir string, options ...BrowserTypeLaunchPersistentContextOptions) (BrowserContext, error)
	// Returns browser name. For example: `'chromium'`, `'webkit'` or `'firefox'`.
	Name() string
}

// ConsoleMessage objects are dispatched by page via the page.on('console') event.
type ConsoleMessage interface {
	Args() []JSHandle
	Location() ConsoleMessageLocation
	String() string
	Text() string
	// One of the following values: `'log'`, `'debug'`, `'info'`, `'error'`, `'warning'`, `'dir'`, `'dirxml'`, `'table'`,
	// `'trace'`, `'clear'`, `'startGroup'`, `'startGroupCollapsed'`, `'endGroup'`, `'assert'`, `'profile'`, `'profileEnd'`,
	// `'count'`, `'timeEnd'`.
	Type() string
}

// Dialog objects are dispatched by page via the page.on('dialog') event.
// An example of using `Dialog` class:
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

// Download objects are dispatched by page via the page.on('download') event.
// All the downloaded files belonging to the browser context are deleted when the browser context is closed. All downloaded
// files are deleted when the browser closes.
// Download event is emitted once the download starts. Download path becomes available once download completes:
// **NOTE** Browser context **must** be created with the `acceptDownloads` set to `true` when user needs access to the
// downloaded content. If `acceptDownloads` is not set or set to `false`, download events are emitted, but the actual
// download is not performed and user has no access to the downloaded files.
type Download interface {
	// Deletes the downloaded file.
	Delete() error
	// Returns download error if any.
	Failure() error
	// Returns path to the downloaded file in case of successful download.
	Path() (string, error)
	// Saves the download to a user-specified path.
	SaveAs(path string) error
	String() string
	// Returns suggested filename for this download. It is typically computed by the browser from the
	// `Content-Disposition` response header
	// or the `download` attribute. See the spec on whatwg. Different
	// browsers can use different logic for computing it.
	SuggestedFilename() string
	// Returns downloaded url.
	URL() string
}

// ElementHandle represents an in-page DOM element. ElementHandles can be created with the page.$(selector) method.
// ElementHandle prevents DOM element from garbage collection unless the handle is disposed with jsHandle.dispose().
// ElementHandles are auto-disposed when their origin frame gets navigated.
// ElementHandle instances can be used as an argument in page.$eval(selector, pageFunction[, arg]) and page.evaluate(pageFunction[, arg]) methods.
type ElementHandle interface {
	AsElement() ElementHandle
	// This method returns the bounding box of the element, or `null` if the element is not visible. The bounding box is
	// calculated relative to the main frame viewport - which is usually the same as the browser window.
	// Scrolling affects the returned bonding box, similarly to
	// Element.getBoundingClientRect. That
	// means `x` and/or `y` may be negative.
	// Elements from child frames return the bounding box relative to the main frame, unlike the
	// Element.getBoundingClientRect.
	// Assuming the page is static, it is safe to use bounding box coordinates to perform input. For example, the following
	// snippet should click the center of the element.
	BoundingBox() (*Rect, error)
	// This method checks the element by performing the following steps:
	// Ensure that element is a checkbox or a radio input. If not, this method rejects. If the element is already checked, this method returns immediately.
	// Wait for actionability checks on the element, unless `force` option is set.
	// Scroll the element into view if needed.
	// Use page.mouse to click in the center of the element.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// Ensure that the element is now checked. If not, this method rejects.
	// If the element is detached from the DOM at any moment during the action, this method rejects.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	Check(options ...ElementHandleCheckOptions) error
	// This method clicks the element by performing the following steps:
	// Wait for actionability checks on the element, unless `force` option is set.
	// Scroll the element into view if needed.
	// Use page.mouse to click in the center of the element, or the specified `position`.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// If the element is detached from the DOM at any moment during the action, this method rejects.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	Click(options ...ElementHandleClickOptions) error
	// Returns the content frame for element handles referencing iframe nodes, or `null` otherwise
	ContentFrame() (Frame, error)
	// This method double clicks the element by performing the following steps:
	// Wait for actionability checks on the element, unless `force` option is set.
	// Scroll the element into view if needed.
	// Use page.mouse to double click in the center of the element, or the specified `position`.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set. Note that if the first click of the `dblclick()` triggers a navigation event, this method will reject.
	// If the element is detached from the DOM at any moment during the action, this method rejects.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	// **NOTE** `elementHandle.dblclick()` dispatches two `click` events and a single `dblclick` event.
	Dblclick(options ...ElementHandleDblclickOptions) error
	// The snippet below dispatches the `click` event on the element. Regardless of the visibility state of the elment, `click`
	// is dispatched. This is equivalend to calling
	// element.click().
	// Under the hood, it creates an instance of an event based on the given `type`, initializes it with `eventInit` properties
	// and dispatches it on the element. Events are `composed`, `cancelable` and bubble by default.
	// Since `eventInit` is event-specific, please refer to the events documentation for the lists of initial properties:
	// DragEvent
	// FocusEvent
	// KeyboardEvent
	// MouseEvent
	// PointerEvent
	// TouchEvent
	// Event
	// You can also specify `JSHandle` as the property value if you want live objects to be passed into the event:
	DispatchEvent(typ string, initObjects ...interface{}) error
	// Returns the return value of `pageFunction`
	// The method finds an element matching the specified selector in the `ElementHandle`s subtree and passes it as a first
	// argument to `pageFunction`. See Working with selectors for more details. If no elements match
	// the selector, the method throws an error.
	// If `pageFunction` returns a Promise, then `frame.$eval` would wait for the promise to resolve and return its value.
	// Examples:
	EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `pageFunction`
	// The method finds all elements matching the specified selector in the `ElementHandle`'s subtree and passes an array of
	// matched elements as a first argument to `pageFunction`. See Working with selectors for more
	// details.
	// If `pageFunction` returns a Promise, then `frame.$$eval` would wait for the promise to resolve and return its value.
	// Examples:
	// ```html
	// <div class="feed">
	// <div class="tweet">Hello!</div>
	// <div class="tweet">Hi!</div>
	// </div>
	// ```
	EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
	// This method waits for actionability checks, focuses the element, fills it and triggers an `input`
	// event after filling. If the element is not an `<input>`, `<textarea>` or `[contenteditable]` element, this method throws
	// an error. Note that you can pass an empty string to clear the input field.
	Fill(value string, options ...ElementHandleFillOptions) error
	// Calls focus on the element.
	Focus() error
	// Returns element attribute value.
	GetAttribute(name string) (string, error)
	// This method hovers over the element by performing the following steps:
	// Wait for actionability checks on the element, unless `force` option is set.
	// Scroll the element into view if needed.
	// Use page.mouse to hover over the center of the element, or the specified `position`.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// If the element is detached from the DOM at any moment during the action, this method rejects.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	Hover(options ...ElementHandleHoverOptions) error
	// Returns the `element.innerHTML`.
	InnerHTML() (string, error)
	// Returns the `element.innerText`.
	InnerText() (string, error)
	// Returns the frame containing the given element.
	OwnerFrame() (Frame, error)
	// Focuses the element, and then uses keyboard.down(key) and keyboard.up(key).
	// `key` can specify the intended keyboardEvent.key
	// value or a single character to generate the text for. A superset of the `key` values can be found
	// here. Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`, etc.
	// Following modification shortcuts are also suported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When speficied with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	Press(options ...ElementHandlePressOptions) error
	// The method finds an element matching the specified selector in the `ElementHandle`'s subtree. See Working with
	// selectors for more details. If no elements match the selector, the return value resolves to
	// `null`.
	QuerySelector(selector string) (ElementHandle, error)
	// The method finds all elements matching the specified selector in the `ElementHandle`s subtree. See Working with
	// selectors for more details. If no elements match the selector, the return value resolves to
	// `[]`.
	QuerySelectorAll(selector string) ([]ElementHandle, error)
	// Returns the buffer with the captured screenshot.
	// This method waits for the actionability checks, then scrolls element into view before taking a
	// screenshot. If the element is detached from DOM, the method throws an error.
	Screenshot(options ...ElementHandleScreenshotOptions) ([]byte, error)
	// This method waits for actionability checks, then tries to scroll element into view, unless it is
	// completely visible as defined by
	// IntersectionObserver's `ratio`.
	// Throws when `elementHandle` does not point to an element
	// connected to a Document or a ShadowRoot.
	ScrollIntoViewIfNeeded(options ...ElementHandleScrollIntoViewIfNeededOptions) error
	// This method waits for actionability checks, then focuses the element and selects all its text
	// content.
	SelectText(options ...ElementHandleSelectTextOptions) error
	// This method expects `elementHandle` to point to an input
	// element.
	// Sets the value of the file input to these file paths or files. If some of the `filePaths` are relative paths, then they
	// are resolved relative to the current working directory. For
	// empty array, clears the selected files.
	SetInputFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error
	// This method taps the element by performing the following steps:
	// Wait for actionability checks on the element, unless `force` option is set.
	// Scroll the element into view if needed.
	// Use page.touchscreen to tap in the center of the element, or the specified `position`.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// If the element is detached from the DOM at any moment during the action, this method rejects.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	// **NOTE** `elementHandle.tap()` requires that the `hasTouch` option of the browser context be set to true.
	Tap(options ...ElementHandleTapOptions) error
	// Returns the `node.textContent`.
	TextContent() (string, error)
	ToString() string
	// Focuses the element, and then sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the text.
	// To press a special key, like `Control` or `ArrowDown`, use elementHandle.press(key[, options]).
	// An example of typing into a text field and then submitting the form:
	Type(value string, options ...ElementHandleTypeOptions) error
	// This method checks the element by performing the following steps:
	// Ensure that element is a checkbox or a radio input. If not, this method rejects. If the element is already unchecked, this method returns immediately.
	// Wait for actionability checks on the element, unless `force` option is set.
	// Scroll the element into view if needed.
	// Use page.mouse to click in the center of the element.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// Ensure that the element is now unchecked. If not, this method rejects.
	// If the element is detached from the DOM at any moment during the action, this method rejects.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	Uncheck(options ...ElementHandleUncheckOptions) error
}

type EventEmitter interface {
	Emit(name string, payload ...interface{})
	ListenerCount(name string) int
	On(name string, handler interface{})
	Once(name string, handler interface{})
	RemoveListener(name string, handler interface{})
}

// FileChooser objects are dispatched by the page in the page.on('filechooser') event.
type FileChooser interface {
	// Returns input element associated with this file chooser.
	Element() ElementHandle
	// Returns whether this file chooser accepts multiple files.
	IsMultiple() bool
	// Returns page this file chooser belongs to.
	Page() Page
	// Sets the value of the file input this chooser is associated with. If some of the `filePaths` are relative paths, then
	// they are resolved relative to the current working directory.
	// For empty array, clears the selected files.
	SetFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error
}

// At every point of time, page exposes its current frame tree via the page.mainFrame() and frame.childFrames()
// methods.
// Frame object's lifecycle is controlled by three events, dispatched on the page object:
// page.on('frameattached') - fired when the frame gets attached to the page. A Frame can be attached to the page only once.
// page.on('framenavigated') - fired when the frame commits navigation to a different URL.
// page.on('framedetached') - fired when the frame gets detached from the page.  A Frame can be detached from the page only once.
// An example of dumping frame tree:
// An example of getting text from an iframe element:
type Frame interface {
	// Returns the added tag when the script's onload fires or when the script content was injected into frame.
	// Adds a `<script>` tag into the page with the desired url or content.
	AddScriptTag(options PageAddScriptTagOptions) (ElementHandle, error)
	// Returns the added tag when the stylesheet's onload fires or when the CSS content was injected into frame.
	// Adds a `<link rel="stylesheet">` tag into the page with the desired url or a `<style type="text/css">` tag with the
	// content.
	AddStyleTag(options PageAddStyleTagOptions) (ElementHandle, error)
	// This method checks an element matching `selector` by performing the following steps:
	// Find an element match matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// Ensure that matched element is a checkbox or a radio input. If not, this method rejects. If the element is already checked, this method returns immediately.
	// Wait for actionability checks on the matched element, unless `force` option is set. If the element is detached during the checks, the whole action is retried.
	// Scroll the element into view if needed.
	// Use page.mouse to click in the center of the element.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// Ensure that the element is now checked. If not, this method rejects.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	Check(selector string, options ...FrameCheckOptions) error
	ChildFrames() []Frame
	// This method clicks an element matching `selector` by performing the following steps:
	// Find an element match matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// Wait for actionability checks on the matched element, unless `force` option is set. If the element is detached during the checks, the whole action is retried.
	// Scroll the element into view if needed.
	// Use page.mouse to click in the center of the element, or the specified `position`.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	Click(selector string, options ...PageClickOptions) error
	// Gets the full HTML contents of the frame, including the doctype.
	Content() (string, error)
	// This method double clicks an element matching `selector` by performing the following steps:
	// Find an element match matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// Wait for actionability checks on the matched element, unless `force` option is set. If the element is detached during the checks, the whole action is retried.
	// Scroll the element into view if needed.
	// Use page.mouse to double click in the center of the element, or the specified `position`.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set. Note that if the first click of the `dblclick()` triggers a navigation event, this method will reject.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	// **NOTE** `frame.dblclick()` dispatches two `click` events and a single `dblclick` event.
	Dblclick(selector string, options ...FrameDblclickOptions) error
	// The snippet below dispatches the `click` event on the element. Regardless of the visibility state of the elment, `click`
	// is dispatched. This is equivalend to calling
	// element.click().
	// Under the hood, it creates an instance of an event based on the given `type`, initializes it with `eventInit` properties
	// and dispatches it on the element. Events are `composed`, `cancelable` and bubble by default.
	// Since `eventInit` is event-specific, please refer to the events documentation for the lists of initial properties:
	// DragEvent
	// FocusEvent
	// KeyboardEvent
	// MouseEvent
	// PointerEvent
	// TouchEvent
	// Event
	// You can also specify `JSHandle` as the property value if you want live objects to be passed into the event:
	DispatchEvent(selector, typ string, options ...PageDispatchEventOptions) error
	// Returns the return value of `pageFunction`
	// If the function passed to the `frame.evaluate` returns a Promise, then `frame.evaluate` would wait for the promise to
	// resolve and return its value.
	// If the function passed to the `frame.evaluate` returns a non-Serializable value, then `frame.evaluate` resolves to
	// `undefined`. DevTools Protocol also supports transferring some additional values that are not serializable by `JSON`:
	// `-0`, `NaN`, `Infinity`, `-Infinity`, and bigint literals.
	// A string can also be passed in instead of a function.
	// ElementHandle instances can be passed as an argument to the `frame.evaluate`:
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `pageFunction` as in-page object (JSHandle).
	// The only difference between `frame.evaluate` and `frame.evaluateHandle` is that `frame.evaluateHandle` returns in-page
	// object (JSHandle).
	// If the function, passed to the `frame.evaluateHandle`, returns a Promise, then `frame.evaluateHandle` would wait for
	// the promise to resolve and return its value.
	// A string can also be passed in instead of a function.
	// JSHandle instances can be passed as an argument to the `frame.evaluateHandle`:
	EvaluateHandle(expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `pageFunction`
	// The method finds an element matching the specified selector within the frame and passes it as a first argument to
	// `pageFunction`. See Working with selectors for more details. If no elements match the
	// selector, the method throws an error.
	// If `pageFunction` returns a Promise, then `frame.$eval` would wait for the promise to resolve and return its value.
	// Examples:
	EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `pageFunction`
	// The method finds all elements matching the specified selector within the frame and passes an array of matched elements
	// as a first argument to `pageFunction`. See Working with selectors for more details.
	// If `pageFunction` returns a Promise, then `frame.$$eval` would wait for the promise to resolve and return its value.
	// Examples:
	EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
	// This method waits for an element matching `selector`, waits for actionability checks, focuses the
	// element, fills it and triggers an `input` event after filling. If the element matching `selector` is not an `<input>`,
	// `<textarea>` or `[contenteditable]` element, this method throws an error. Note that you can pass an empty string to
	// clear the input field.
	// To send fine-grained keyboard events, use frame.type(selector, text[, options]).
	Fill(selector string, value string, options ...FrameFillOptions) error
	// This method fetches an element with `selector` and focuses it. If there's no element matching `selector`, the method
	// waits until a matching element appears in the DOM.
	Focus(selector string, options ...FrameFocusOptions) error
	// Returns the `frame` or `iframe` element handle which corresponds to this frame.
	// This is an inverse of elementHandle.contentFrame(). Note that returned handle actually belongs to the parent frame.
	// This method throws an error if the frame has been detached before `frameElement()` returns.
	FrameElement() (ElementHandle, error)
	// Returns element attribute value.
	GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error)
	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of the
	// last redirect.
	// `frame.goto` will throw an error if:
	// there's an SSL error (e.g. in case of self-signed certificates).
	// target URL is invalid.
	// the `timeout` is exceeded during navigation.
	// the remote server does not respond or is unreachable.
	// the main resource failed to load.
	// `frame.goto` will not throw an error when any valid HTTP status code is returned by the remote server, including 404
	// "Not Found" and 500 "Internal Server Error".  The status code for such responses can be retrieved by calling
	// response.status().
	// **NOTE** `frame.goto` either throws an error or returns a main resource response. The only exceptions are navigation
	// to `about:blank` or navigation to the same URL with a different hash, which would succeed and return `null`.
	// **NOTE** Headless mode doesn't support navigation to a PDF document. See the upstream
	// issue.
	Goto(url string, options ...PageGotoOptions) (Response, error)
	// This method hovers over an element matching `selector` by performing the following steps:
	// Find an element match matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// Wait for actionability checks on the matched element, unless `force` option is set. If the element is detached during the checks, the whole action is retried.
	// Scroll the element into view if needed.
	// Use page.mouse to hover over the center of the element, or the specified `position`.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	Hover(selector string, options ...PageHoverOptions) error
	// Resolves to the `element.innerHTML`.
	InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error)
	// Resolves to the `element.innerText`.
	InnerText(selector string, options ...PageInnerTextOptions) (string, error)
	// Returns `true` if the frame has been detached, or `false` otherwise.
	IsDetached() bool
	// Returns frame's name attribute as specified in the tag.
	// If the name is empty, returns the id attribute instead.
	// **NOTE** This value is calculated once when the frame is created, and will not update if the attribute is changed
	// later.
	Name() string
	// Returns the page containing this frame.
	Page() Page
	// Parent frame, if any. Detached frames and main frames return `null`.
	ParentFrame() Frame
	// `key` can specify the intended keyboardEvent.key
	// value or a single character to generate the text for. A superset of the `key` values can be found
	// here. Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`, etc.
	// Following modification shortcuts are also suported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When speficied with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	Press(selector, key string, options ...PagePressOptions) error
	// Returns the ElementHandle pointing to the frame element.
	// The method finds an element matching the specified selector within the frame. See Working with
	// selectors for more details. If no elements match the selector, the return value resolves to
	// `null`.
	QuerySelector(selector string) (ElementHandle, error)
	// Returns the ElementHandles pointing to the frame elements.
	// The method finds all elements matching the specified selector within the frame. See Working with
	// selectors for more details. If no elements match the selector, the return value resolves to
	// `[]`.
	QuerySelectorAll(selector string) ([]ElementHandle, error)
	SetContent(content string, options ...PageSetContentOptions) error
	// This method expects `selector` to point to an input
	// element.
	// Sets the value of the file input to these file paths or files. If some of the `filePaths` are relative paths, then they
	// are resolved relative to the current working directory. For
	// empty array, clears the selected files.
	SetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) error
	// This method taps an element matching `selector` by performing the following steps:
	// Find an element match matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// Wait for actionability checks on the matched element, unless `force` option is set. If the element is detached during the checks, the whole action is retried.
	// Scroll the element into view if needed.
	// Use page.touchscreen to tap the center of the element, or the specified `position`.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	// **NOTE** `frame.tap()` requires that the `hasTouch` option of the browser context be set to true.
	Tap(selector string, options ...FrameTapOptions) error
	// Resolves to the `element.textContent`.
	TextContent(selector string, options ...FrameTextContentOptions) (string, error)
	// Returns the page title.
	Title() (string, error)
	// Sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the text. `frame.type` can be used to
	// send fine-grained keyboard events. To fill values in form fields, use frame.fill(selector, value[, options]).
	// To press a special key, like `Control` or `ArrowDown`, use keyboard.press(key[, options]).
	Type(selector, text string, options ...PageTypeOptions) error
	// Returns frame's url.
	URL() string
	// This method checks an element matching `selector` by performing the following steps:
	// Find an element match matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// Ensure that matched element is a checkbox or a radio input. If not, this method rejects. If the element is already unchecked, this method returns immediately.
	// Wait for actionability checks on the matched element, unless `force` option is set. If the element is detached during the checks, the whole action is retried.
	// Scroll the element into view if needed.
	// Use page.mouse to click in the center of the element.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// Ensure that the element is now unchecked. If not, this method rejects.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	Uncheck(selector string, options ...FrameUncheckOptions) error
	WaitForEvent(event string, predicate ...interface{}) interface{}
	// Returns when the `pageFunction` returns a truthy value. It resolves to a JSHandle of the truthy value.
	// The `waitForFunction` can be used to observe viewport size change:
	// To pass an argument from Node.js to the predicate of `frame.waitForFunction` function:
	WaitForFunction(expression string, options ...FrameWaitForFunctionOptions) (JSHandle, error)
	// Returns when the required load state has been reached.
	// This resolves when the frame reaches a required load state, `load` by default. The navigation must have been committed
	// when this method is called. If current document has already reached the required state, resolves immediately.
	WaitForLoadState(given ...string)
	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of the
	// last redirect. In case of navigation to a different anchor or navigation due to History API usage, the navigation will
	// resolve with `null`.
	// This resolves when the frame navigates to a new URL. It is useful for when you run code which will indirectly cause the
	// frame to navigate. Consider this example:
	// **NOTE** Usage of the History API to change the URL is
	// considered a navigation.
	WaitForNavigation(options ...PageWaitForNavigationOptions) (Response, error)
	// Returns when element specified by selector satisfies `state` option. Resolves to `null` if waiting for `hidden` or
	// `detached`.
	// Wait for the `selector` to satisfy `state` option (either appear/disappear from dom, or become visible/hidden). If at
	// the moment of calling the method `selector` already satisfies the condition, the method will return immediately. If the
	// selector doesn't satisfy the condition for the `timeout` milliseconds, the function will throw.
	// This method works across navigations:
	WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandle, error)
	// Returns a promise that resolves after the timeout.
	// Note that `frame.waitForTimeout()` should only be used for debugging. Tests using the timer in production are going to
	// be flaky. Use signals such as network events, selectors becoming visible and others instead.
	WaitForTimeout(timeout int)
}

// JSHandle represents an in-page JavaScript object. JSHandles can be created with the page.evaluateHandle(pageFunction[, arg]) method.
// JSHandle prevents the referenced JavaScript object being garbage collected unless the handle is exposed with
// jsHandle.dispose(). JSHandles are auto-disposed when their origin frame gets navigated or the parent context gets
// destroyed.
// JSHandle instances can be used as an argument in page.$eval(selector, pageFunction[, arg]), page.evaluate(pageFunction[, arg]) and page.evaluateHandle(pageFunction[, arg])
// methods.
type JSHandle interface {
	// Returns either `null` or the object handle itself, if the object handle is an instance of ElementHandle.
	AsElement() ElementHandle
	// The `jsHandle.dispose` method stops referencing the element handle.
	Dispose() error
	// Returns the return value of `pageFunction`
	// This method passes this handle as the first argument to `pageFunction`.
	// If `pageFunction` returns a Promise, then `handle.evaluate` would wait for the promise to resolve and return its
	// value.
	// Examples:
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `pageFunction` as in-page object (JSHandle).
	// This method passes this handle as the first argument to `pageFunction`.
	// The only difference between `jsHandle.evaluate` and `jsHandle.evaluateHandle` is that `jsHandle.evaluateHandle` returns
	// in-page object (JSHandle).
	// If the function passed to the `jsHandle.evaluateHandle` returns a Promise, then `jsHandle.evaluateHandle` would wait
	// for the promise to resolve and return its value.
	// See page.evaluateHandle(pageFunction[, arg]) for more details.
	EvaluateHandle(expression string, options ...interface{}) (interface{}, error)
	// The method returns a map with **own property names** as keys and JSHandle instances for the property values.
	GetProperties() (map[string]JSHandle, error)
	// Fetches a single property from the referenced object.
	GetProperty(name string) (JSHandle, error)
	// Returns a JSON representation of the object. If the object has a
	// `toJSON`
	// function, it **will not be called**.
	// **NOTE** The method will return an empty JSON object if the referenced object is not stringifiable. It will throw an
	// error if the object has circular references.
	JSONValue() (interface{}, error)
	String() string
}

// Keyboard provides an api for managing a virtual keyboard. The high level api is keyboard.type(text[, options]), which takes raw
// characters and generates proper keydown, keypress/input, and keyup events on your page.
// For finer control, you can use keyboard.down(key), keyboard.up(key), and keyboard.insertText(text) to manually fire
// events as if they were generated from a real keyboard.
// An example of holding down `Shift` in order to select and delete some text:
// An example of pressing uppercase `A`
// An example to trigger select-all with the keyboard
type Keyboard interface {
	// Dispatches a `keydown` event.
	// `key` can specify the intended keyboardEvent.key
	// value or a single character to generate the text for. A superset of the `key` values can be found
	// here. Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`, etc.
	// Following modification shortcuts are also suported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// If `key` is a modifier key, `Shift`, `Meta`, `Control`, or `Alt`, subsequent key presses will be sent with that modifier
	// active. To release the modifier key, use keyboard.up(key).
	// After the key is pressed once, subsequent calls to keyboard.down(key) will have
	// repeat set to true. To release the key, use
	// keyboard.up(key).
	// **NOTE** Modifier keys DO influence `keyboard.down`. Holding down `Shift` will type the text in upper case.
	Down(key string) error
	// Dispatches only `input` event, does not emit the `keydown`, `keyup` or `keypress` events.
	// **NOTE** Modifier keys DO NOT effect `keyboard.insertText`. Holding down `Shift` will not type the text in upper case.
	InsertText(text string) error
	// `key` can specify the intended keyboardEvent.key
	// value or a single character to generate the text for. A superset of the `key` values can be found
	// here. Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`, etc.
	// Following modification shortcuts are also suported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When speficied with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	// Shortcut for keyboard.down(key) and keyboard.up(key).
	Press(key string, options ...KeyboardPressOptions) error
	// Sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the text.
	// To press a special key, like `Control` or `ArrowDown`, use keyboard.press(key[, options]).
	// **NOTE** Modifier keys DO NOT effect `keyboard.type`. Holding down `Shift` will not type the text in upper case.
	Type(text string, options ...KeyboardTypeOptions) error
	// Dispatches a `keyup` event.
	Up(key string) error
}

// The Mouse class operates in main-frame CSS pixels relative to the top-left corner of the viewport.
// Every `page` object has its own Mouse, accessible with page.mouse.
type Mouse interface {
	// Shortcut for mouse.move(x, y[, options]), mouse.down([options]), mouse.up([options]).
	Click(x, y float64, options ...MouseClickOptions) error
	// Shortcut for mouse.move(x, y[, options]), mouse.down([options]), mouse.up([options]), mouse.down([options]) and mouse.up([options]).
	Dblclick(x, y float64, options ...MouseDblclickOptions) error
	// Dispatches a `mousedown` event.
	Down(options ...MouseDownOptions) error
	// Dispatches a `mousemove` event.
	Move(x float64, y float64, options ...MouseMoveOptions) error
	// Dispatches a `mouseup` event.
	Up(options ...MouseUpOptions) error
}

// Page provides methods to interact with a single tab in a Browser, or an extension background
// page in Chromium. One Browser instance might have multiple
// Page instances.
// This example creates a page, navigates it to a URL, and then saves a screenshot:
// The Page class emits various events (described below) which can be handled using any of Node's native
// `EventEmitter` methods, such as `on`, `once` or
// `removeListener`.
// This example logs a message for a single page `load` event:
// To unsubscribe from events use the `removeListener` method:
type Page interface {
	Mouse() Mouse
	Keyboard() Keyboard
	Touchscreen() Touchscreen
	EventEmitter
	// Adds a script which would be evaluated in one of the following scenarios:
	// Whenever the page is navigated.
	// Whenever the child frame is attached or navigated. In this case, the script is evaluated in the context of the newly attached frame.
	// The script is evaluated after the document was created but before any of its scripts were run. This is useful to amend
	// the JavaScript environment, e.g. to seed `Math.random`.
	// An example of overriding `Math.random` before the page loads:
	// **NOTE** The order of evaluation of multiple scripts installed via browserContext.addInitScript(script[, arg]) and
	// page.addInitScript(script[, arg]) is not defined.
	AddInitScript(options BrowserContextAddInitScriptOptions) error
	// Adds a `<script>` tag into the page with the desired url or content. Returns the added tag when the script's onload
	// fires or when the script content was injected into frame.
	// Shortcut for main frame's frame.addScriptTag(script).
	AddScriptTag(options PageAddScriptTagOptions) (ElementHandle, error)
	// Adds a `<link rel="stylesheet">` tag into the page with the desired url or a `<style type="text/css">` tag with the
	// content. Returns the added tag when the stylesheet's onload fires or when the CSS content was injected into frame.
	// Shortcut for main frame's frame.addStyleTag(style).
	AddStyleTag(options PageAddStyleTagOptions) (ElementHandle, error)
	// Brings page to front (activates tab).
	BringToFront() error
	// This method checks an element matching `selector` by performing the following steps:
	// Find an element match matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// Ensure that matched element is a checkbox or a radio input. If not, this method rejects. If the element is already checked, this method returns immediately.
	// Wait for actionability checks on the matched element, unless `force` option is set. If the element is detached during the checks, the whole action is retried.
	// Scroll the element into view if needed.
	// Use page.mouse to click in the center of the element.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// Ensure that the element is now checked. If not, this method rejects.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	// Shortcut for main frame's frame.check(selector[, options]).
	Check(selector string, options ...FrameCheckOptions) error
	// This method clicks an element matching `selector` by performing the following steps:
	// Find an element match matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// Wait for actionability checks on the matched element, unless `force` option is set. If the element is detached during the checks, the whole action is retried.
	// Scroll the element into view if needed.
	// Use page.mouse to click in the center of the element, or the specified `position`.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	// Shortcut for main frame's frame.click(selector[, options]).
	Click(selector string, options ...PageClickOptions) error
	// If `runBeforeUnload` is `false` the result will resolve only after the page has been closed. If `runBeforeUnload` is
	// `true` the method will **not** wait for the page to close. By default, `page.close()` **does not** run beforeunload
	// handlers.
	// **NOTE** if `runBeforeUnload` is passed as true, a `beforeunload` dialog might be summoned
	// and should be handled manually via page.on('dialog') event.
	Close(options ...PageCloseOptions) error
	// Gets the full HTML contents of the page, including the doctype.
	Content() (string, error)
	// Get the browser context that the page belongs to.
	Context() BrowserContext
	// This method double clicks an element matching `selector` by performing the following steps:
	// Find an element match matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// Wait for actionability checks on the matched element, unless `force` option is set. If the element is detached during the checks, the whole action is retried.
	// Scroll the element into view if needed.
	// Use page.mouse to double click in the center of the element, or the specified `position`.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set. Note that if the first click of the `dblclick()` triggers a navigation event, this method will reject.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	// **NOTE** `page.dblclick()` dispatches two `click` events and a single `dblclick` event.
	// Shortcut for main frame's frame.dblclick(selector[, options]).
	Dblclick(expression string, options ...FrameDblclickOptions) error
	// The snippet below dispatches the `click` event on the element. Regardless of the visibility state of the elment, `click`
	// is dispatched. This is equivalend to calling
	// element.click().
	// Under the hood, it creates an instance of an event based on the given `type`, initializes it with `eventInit` properties
	// and dispatches it on the element. Events are `composed`, `cancelable` and bubble by default.
	// Since `eventInit` is event-specific, please refer to the events documentation for the lists of initial properties:
	// DragEvent
	// FocusEvent
	// KeyboardEvent
	// MouseEvent
	// PointerEvent
	// TouchEvent
	// Event
	// You can also specify `JSHandle` as the property value if you want live objects to be passed into the event:
	DispatchEvent(selector string, typ string, options ...PageDispatchEventOptions) error
	// The method adds a function called `name` on the `window` object of every frame in this page. When called, the function
	// executes `playwrightBinding` in Node.js and returns a Promise which resolves to the return value of
	// `playwrightBinding`. If the `playwrightBinding` returns a Promise, it will be awaited.
	// The first argument of the `playwrightBinding` function contains information about the caller: `{ browserContext: BrowserContext, page: Page, frame: Frame }`.
	// See browserContext.exposeBinding(name, playwrightBinding[, options]) for the context-wide version.
	// **NOTE** Functions installed via `page.exposeBinding` survive navigations.
	// An example of exposing page URL to all frames in a page:
	// An example of passing an element handle:
	ExposeBinding(name string, binding BindingCallFunction, handle ...bool) error
	// The method adds a function called `name` on the `window` object of every frame in the page. When called, the function
	// executes `playwrightFunction` in Node.js and returns a Promise which resolves to the return value of
	// `playwrightFunction`.
	// If the `playwrightFunction` returns a Promise, it will be awaited.
	// See browserContext.exposeFunction(name, playwrightFunction) for context-wide exposed function.
	// **NOTE** Functions installed via `page.exposeFunction` survive navigations.
	// An example of adding an `md5` function to the page:
	// An example of adding a `window.readfile` function to the page:
	ExposeFunction(name string, binding ExposedFunction) error
	EmulateMedia(options ...PageEmulateMediaOptions) error
	// Returns the value of the `pageFunction` invacation.
	// If the function passed to the `page.evaluate` returns a Promise, then `page.evaluate` would wait for the promise to
	// resolve and return its value.
	// If the function passed to the `page.evaluate` returns a non-Serializable value, then `page.evaluate` resolves to
	// `undefined`. DevTools Protocol also supports transferring some additional values that are not serializable by `JSON`:
	// `-0`, `NaN`, `Infinity`, `-Infinity`, and bigint literals.
	// Passing argument to `pageFunction`:
	// A string can also be passed in instead of a function:
	// ElementHandle instances can be passed as an argument to the `page.evaluate`:
	// Shortcut for main frame's frame.evaluate(pageFunction[, arg]).
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	// Returns the value of the `pageFunction` invacation as in-page object (JSHandle).
	// The only difference between `page.evaluate` and `page.evaluateHandle` is that `page.evaluateHandle` returns in-page
	// object (JSHandle).
	// If the function passed to the `page.evaluateHandle` returns a Promise, then `page.evaluateHandle` would wait for the
	// promise to resolve and return its value.
	// A string can also be passed in instead of a function:
	// JSHandle instances can be passed as an argument to the `page.evaluateHandle`:
	EvaluateHandle(expression string, options ...interface{}) (interface{}, error)
	// The method finds an element matching the specified selector within the page and passes it as a first argument to
	// `pageFunction`. If no elements match the selector, the method throws an error. Returns the value of `pageFunction`.
	// If `pageFunction` returns a Promise, then `page.$eval` would wait for the promise to resolve and return its value.
	// Examples:
	// Shortcut for main frame's frame.$eval(selector, pageFunction[, arg]).
	EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
	// The method finds all elements matching the specified selector within the page and passes an array of matched elements as
	// a first argument to `pageFunction`. Returns the result of `pageFunction` invocation.
	// If `pageFunction` returns a Promise, then `page.$$eval` would wait for the promise to resolve and return its value.
	// Examples:
	EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
	ExpectConsoleMessage(cb func() error) (ConsoleMessage, error)
	ExpectDownload(cb func() error) (Download, error)
	ExpectEvent(event string, cb func() error, predicates ...interface{}) (interface{}, error)
	ExpectFileChooser(cb func() error) (FileChooser, error)
	ExpectLoadState(state string, cb func() error) (ConsoleMessage, error)
	ExpectNavigation(cb func() error, options ...PageWaitForNavigationOptions) (Response, error)
	ExpectPopup(cb func() error) (Page, error)
	ExpectRequest(url interface{}, cb func() error, options ...interface{}) (Request, error)
	ExpectResponse(url interface{}, cb func() error, options ...interface{}) (Response, error)
	ExpectWorker(cb func() error) (Worker, error)
	ExpectedDialog(cb func() error) (Dialog, error)
	// This method waits for an element matching `selector`, waits for actionability checks, focuses the
	// element, fills it and triggers an `input` event after filling. If the element matching `selector` is not an `<input>`,
	// `<textarea>` or `[contenteditable]` element, this method throws an error. Note that you can pass an empty string to
	// clear the input field.
	// To send fine-grained keyboard events, use page.type(selector, text[, options]).
	// Shortcut for main frame's frame.fill(selector, value[, options])
	Fill(selector, text string, options ...FrameFillOptions) error
	// This method fetches an element with `selector` and focuses it. If there's no element matching `selector`, the method
	// waits until a matching element appears in the DOM.
	// Shortcut for main frame's frame.focus(selector[, options]).
	Focus(expression string, options ...FrameFocusOptions) error
	// Returns frame matching the specified criteria. Either `name` or `url` must be specified.
	Frame(options PageFrameOptions) Frame
	// An array of all frames attached to the page.
	Frames() []Frame
	// Returns element attribute value.
	GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error)
	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of the
	// last redirect. If can not go back, resolves to `null`.
	// Navigate to the previous page in history.
	GoBack(options ...PageGoBackOptions) (Response, error)
	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of the
	// last redirect. If can not go forward, resolves to `null`.
	// Navigate to the next page in history.
	GoForward(options ...PageGoForwardOptions) (Response, error)
	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of the
	// last redirect.
	// `page.goto` will throw an error if:
	// there's an SSL error (e.g. in case of self-signed certificates).
	// target URL is invalid.
	// the `timeout` is exceeded during navigation.
	// the remote server does not respond or is unreachable.
	// the main resource failed to load.
	// `page.goto` will not throw an error when any valid HTTP status code is returned by the remote server, including 404 "Not
	// Found" and 500 "Internal Server Error".  The status code for such responses can be retrieved by calling
	// response.status().
	// **NOTE** `page.goto` either throws an error or returns a main resource response. The only exceptions are navigation to
	// `about:blank` or navigation to the same URL with a different hash, which would succeed and return `null`.
	// **NOTE** Headless mode doesn't support navigation to a PDF document. See the upstream
	// issue.
	// Shortcut for main frame's frame.goto(url[, options])
	Goto(url string, options ...PageGotoOptions) (Response, error)
	// This method hovers over an element matching `selector` by performing the following steps:
	// Find an element match matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// Wait for actionability checks on the matched element, unless `force` option is set. If the element is detached during the checks, the whole action is retried.
	// Scroll the element into view if needed.
	// Use page.mouse to hover over the center of the element, or the specified `position`.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	// Shortcut for main frame's frame.hover(selector[, options]).
	Hover(selector string, options ...PageHoverOptions) error
	// Returns `element.innerHTML`.
	InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error)
	// Returns `element.innerText`.
	InnerText(selector string, options ...PageInnerTextOptions) (string, error)
	// Indicates that the page has been closed.
	IsClosed() bool
	// The page's main frame. Page is guaranteed to have a main frame which persists during navigations.
	MainFrame() Frame
	// Returns the opener for popup pages and `null` for others. If the opener has been closed already the promise may resolve
	// to `null`.
	Opener() (Page, error)
	// Returns the PDF buffer.
	// **NOTE** Generating a pdf is currently only supported in Chromium headless.
	// `page.pdf()` generates a pdf of the page with `print` css media. To generate a pdf with `screen` media, call
	// page.emulateMedia(params) before calling `page.pdf()`:
	// **NOTE** By default, `page.pdf()` generates a pdf with modified colors for printing. Use the
	// `-webkit-print-color-adjust` property to
	// force rendering of exact colors.
	// The `width`, `height`, and `margin` options accept values labeled with units. Unlabeled values are treated as pixels.
	// A few examples:
	// `page.pdf({width: 100})` - prints with width set to 100 pixels
	// `page.pdf({width: '100px'})` - prints with width set to 100 pixels
	// `page.pdf({width: '10cm'})` - prints with width set to 10 centimeters.
	// All possible units are:
	// `px` - pixel
	// `in` - inch
	// `cm` - centimeter
	// `mm` - millimeter
	// The `format` options are:
	// `Letter`: 8.5in x 11in
	// `Legal`: 8.5in x 14in
	// `Tabloid`: 11in x 17in
	// `Ledger`: 17in x 11in
	// `A0`: 33.1in x 46.8in
	// `A1`: 23.4in x 33.1in
	// `A2`: 16.54in x 23.4in
	// `A3`: 11.7in x 16.54in
	// `A4`: 8.27in x 11.7in
	// `A5`: 5.83in x 8.27in
	// `A6`: 4.13in x 5.83in
	// **NOTE** `headerTemplate` and `footerTemplate` markup have the following limitations:
	// Script tags inside templates are not evaluated.
	// Page styles are not visible inside templates.
	PDF(options ...PagePDFOptions) ([]byte, error)
	// Focuses the element, and then uses keyboard.down(key) and keyboard.up(key).
	// `key` can specify the intended keyboardEvent.key
	// value or a single character to generate the text for. A superset of the `key` values can be found
	// here. Examples of the keys are:
	// `F1` - `F12`, `Digit0`- `Digit9`, `KeyA`- `KeyZ`, `Backquote`, `Minus`, `Equal`, `Backslash`, `Backspace`, `Tab`,
	// `Delete`, `Escape`, `ArrowDown`, `End`, `Enter`, `Home`, `Insert`, `PageDown`, `PageUp`, `ArrowRight`, `ArrowUp`, etc.
	// Following modification shortcuts are also suported: `Shift`, `Control`, `Alt`, `Meta`, `ShiftLeft`.
	// Holding down `Shift` will type the text that corresponds to the `key` in the upper case.
	// If `key` is a single character, it is case-sensitive, so the values `a` and `A` will generate different respective
	// texts.
	// Shortcuts such as `key: "Control+o"` or `key: "Control+Shift+T"` are supported as well. When speficied with the
	// modifier, modifier is pressed and being held while the subsequent key is being pressed.
	Press(selector, key string, options ...PagePressOptions) error
	// The method finds an element matching the specified selector within the page. If no elements match the selector, the
	// return value resolves to `null`.
	// Shortcut for main frame's frame.$(selector).
	QuerySelector(selector string) (ElementHandle, error)
	// The method finds all elements matching the specified selector within the page. If no elements match the selector, the
	// return value resolves to `[]`.
	// Shortcut for main frame's frame.$$(selector).
	QuerySelectorAll(selector string) ([]ElementHandle, error)
	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of the
	// last redirect.
	Reload(options ...PageReloadOptions) (Response, error)
	// Routing provides the capability to modify network requests that are made by a page.
	// Once routing is enabled, every request matching the url pattern will stall unless it's continued, fulfilled or aborted.
	// **NOTE** The handler will only be called for the first url if the response is a redirect.
	// An example of a naïve handler that aborts all image requests:
	// or the same snippet using a regex pattern instead:
	// Page routes take precedence over browser context routes (set up with browserContext.route(url, handler)) when request matches
	// both handlers.
	// **NOTE** Enabling routing disables http cache.
	Route(url interface{}, handler routeHandler) error
	// Returns the buffer with the captured screenshot.
	// **NOTE** Screenshots take at least 1/6 second on Chromium OS X and Chromium Windows. See https://crbug.com/741689 for
	// discussion.
	Screenshot(options ...PageScreenshotOptions) ([]byte, error)
	SetContent(content string, options ...PageSetContentOptions) error
	// This setting will change the default maximum navigation time for the following methods and related shortcuts:
	// page.goBack([options])
	// page.goForward([options])
	// page.goto(url[, options])
	// page.reload([options])
	// page.setContent(html[, options])
	// page.waitForNavigation([options])
	// **NOTE** page.setDefaultNavigationTimeout(timeout) takes priority over page.setDefaultTimeout(timeout),
	// browserContext.setDefaultTimeout(timeout) and browserContext.setDefaultNavigationTimeout(timeout).
	SetDefaultNavigationTimeout(timeout int)
	// This setting will change the default maximum time for all the methods accepting `timeout` option.
	// **NOTE** page.setDefaultNavigationTimeout(timeout) takes priority over page.setDefaultTimeout(timeout).
	SetDefaultTimeout(timeout int)
	// The extra HTTP headers will be sent with every request the page initiates.
	// **NOTE** page.setExtraHTTPHeaders does not guarantee the order of headers in the outgoing requests.
	SetExtraHTTPHeaders(headers map[string]string) error
	// This method expects `selector` to point to an input
	// element.
	// Sets the value of the file input to these file paths or files. If some of the `filePaths` are relative paths, then they
	// are resolved relative to the current working directory. For
	// empty array, clears the selected files.
	SetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) error
	// In the case of multiple pages in a single browser, each page can have its own viewport size. However,
	// browser.newContext([options]) allows to set viewport size (and more) for all pages in the context at once.
	// `page.setViewportSize` will resize the page. A lot of websites don't expect phones to change size, so you should set the
	// viewport size before navigating to the page.
	SetViewportSize(width, height int) error
	// This method taps an element matching `selector` by performing the following steps:
	// Find an element match matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// Wait for actionability checks on the matched element, unless `force` option is set. If the element is detached during the checks, the whole action is retried.
	// Scroll the element into view if needed.
	// Use page.touchscreen to tap the center of the element, or the specified `position`.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	// **NOTE** `page.tap()` requires that the `hasTouch` option of the browser context be set to true.
	// Shortcut for main frame's frame.tap(selector[, options]).
	Tap(selector string, options ...FrameTapOptions) error
	// Returns `element.textContent`.
	TextContent(selector string, options ...FrameTextContentOptions) (string, error)
	// Returns the page's title. Shortcut for main frame's frame.title().
	Title() (string, error)
	// Sends a `keydown`, `keypress`/`input`, and `keyup` event for each character in the text. `page.type` can be used to send
	// fine-grained keyboard events. To fill values in form fields, use page.fill(selector, value[, options]).
	// To press a special key, like `Control` or `ArrowDown`, use keyboard.press(key[, options]).
	// Shortcut for main frame's frame.type(selector, text[, options]).
	Type(selector, text string, options ...PageTypeOptions) error
	// Shortcut for main frame's frame.url().
	URL() string
	// This method unchecks an element matching `selector` by performing the following steps:
	// Find an element match matching `selector`. If there is none, wait until a matching element is attached to the DOM.
	// Ensure that matched element is a checkbox or a radio input. If not, this method rejects. If the element is already unchecked, this method returns immediately.
	// Wait for actionability checks on the matched element, unless `force` option is set. If the element is detached during the checks, the whole action is retried.
	// Scroll the element into view if needed.
	// Use page.mouse to click in the center of the element.
	// Wait for initiated navigations to either succeed or fail, unless `noWaitAfter` option is set.
	// Ensure that the element is now unchecked. If not, this method rejects.
	// When all steps combined have not finished during the specified `timeout`, this method rejects with a TimeoutError.
	// Passing zero timeout disables this.
	// Shortcut for main frame's frame.uncheck(selector[, options]).
	Uncheck(selector string, options ...FrameUncheckOptions) error
	// Video object associated with this page.
	Video() Video
	ViewportSize() ViewportSize
	// Returns the event data value.
	// Waits for event to fire and passes its value into the predicate function. Resolves when the predicate returns truthy
	// value. Will throw an error if the page is closed before the event is fired.
	WaitForEvent(event string, predicate ...interface{}) interface{}
	// Returns when the `pageFunction` returns a truthy value. It resolves to a JSHandle of the truthy value.
	// The `waitForFunction` can be used to observe viewport size change:
	// To pass an argument from Node.js to the predicate of `page.waitForFunction` function:
	// Shortcut for main frame's frame.waitForFunction(pageFunction[, arg, options]).
	WaitForFunction(expression string, options ...FrameWaitForFunctionOptions) (JSHandle, error)
	// Returns when the required load state has been reached.
	// This resolves when the page reaches a required load state, `load` by default. The navigation must have been committed
	// when this method is called. If current document has already reached the required state, resolves immediately.
	// Shortcut for main frame's frame.waitForLoadState([state, options]).
	WaitForLoadState(state ...string)
	// Returns the main resource response. In case of multiple redirects, the navigation will resolve with the response of the
	// last redirect. In case of navigation to a different anchor or navigation due to History API usage, the navigation will
	// resolve with `null`.
	// This resolves when the page navigates to a new URL or reloads. It is useful for when you run code which will indirectly
	// cause the page to navigate. e.g. The click target has an `onclick` handler that triggers navigation from a `setTimeout`.
	// Consider this example:
	// **NOTE** Usage of the History API to change the URL is
	// considered a navigation.
	// Shortcut for main frame's frame.waitForNavigation([options]).
	WaitForNavigation(options ...PageWaitForNavigationOptions) (Response, error)
	// Returns promise that resolves to the matched request.
	WaitForRequest(url interface{}, options ...interface{}) Request
	// Returns the matched response.
	WaitForResponse(url interface{}, options ...interface{}) Response
	// Returns when element specified by selector satisfies `state` option. Resolves to `null` if waiting for `hidden` or
	// `detached`.
	// Wait for the `selector` to satisfy `state` option (either appear/disappear from dom, or become visible/hidden). If at
	// the moment of calling the method `selector` already satisfies the condition, the method will return immediately. If the
	// selector doesn't satisfy the condition for the `timeout` milliseconds, the function will throw.
	// This method works across navigations:
	WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandle, error)
	// Returns a promise that resolves after the timeout.
	// Note that `page.waitForTimeout()` should only be used for debugging. Tests using the timer in production are going to be
	// flaky. Use signals such as network events, selectors becoming visible and others instead.
	// Shortcut for main frame's frame.waitForTimeout(timeout).
	WaitForTimeout(timeout int)
	// This method returns all of the dedicated WebWorkers
	// associated with the page.
	// **NOTE** This does not contain ServiceWorkers
	Workers() []Worker
}

// Whenever the page sends a request for a network resource the following sequence of events are emitted by Page:
// page.on('request') emitted when the request is issued by the page.
// page.on('response') emitted when/if the response status and headers are received for the request.
// page.on('requestfinished') emitted when the response body is downloaded and the request is complete.
// If request fails at some point, then instead of `'requestfinished'` event (and possibly instead of 'response' event),
// the  page.on('requestfailed') event is emitted.
// **NOTE** HTTP Error responses, such as 404 or 503, are still successful responses from HTTP standpoint, so request
// will complete with `'requestfinished'` event.
// If request gets a 'redirect' response, the request is successfully finished with the 'requestfinished' event, and a new
// request is  issued to a redirected url.
type Request interface {
	// The method returns `null` unless this request has failed, as reported by `requestfailed` event.
	// Example of logging of all the failed requests:
	Failure() *RequestFailure
	// Returns the Frame that initiated this request.
	Frame() Frame
	// An object with HTTP headers associated with the request. All header names are lower-case.
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
	// When the server responds with a redirect, Playwright creates a new Request object. The two requests are connected by
	// `redirectedFrom()` and `redirectedTo()` methods. When multiple server redirects has happened, it is possible to
	// construct the whole redirect chain by repeatedly calling `redirectedFrom()`.
	// For example, if the website `http://example.com` redirects to `https://example.com`:
	// If the website `https://google.com` has no redirects:
	RedirectedFrom() Request
	// New request issued by the browser if the server responded with redirect.
	// This method is the opposite of request.redirectedFrom():
	RedirectedTo() Request
	// Contains the request's resource type as it was perceived by the rendering engine. ResourceType will be one of the
	// following: `document`, `stylesheet`, `image`, `media`, `font`, `script`, `texttrack`, `xhr`, `fetch`, `eventsource`,
	// `websocket`, `manifest`, `other`.
	ResourceType() string
	// Returns the matching Response object, or `null` if the response was not received due to error.
	Response() (Response, error)
	// Returns resource timing information for given request. Most of the timing values become available upon the response,
	// `responseEnd` becomes available when request finishes. Find more information at Resource Timing
	// API.
	Timing() *ResourceTiming
	// URL of the request.
	URL() string
}

// Response class represents responses which are received by page.
type Response interface {
	// Returns the buffer with response body.
	Body() ([]byte, error)
	// Waits for this response to finish, returns failure error if request failed.
	Finished() error
	// Returns the Frame that initiated this response.
	Frame() Frame
	// Returns the object with HTTP headers associated with the response. All header names are lower-case.
	Headers() map[string]string
	// Returns the JSON representation of response body.
	// This method will throw if the response body is not parsable via `JSON.parse`.
	JSON(v interface{}) error
	// Contains a boolean stating whether the response was successful (status in the range 200-299) or not.
	Ok() bool
	// Returns the matching Request object.
	Request() Request
	// Contains the status code of the response (e.g., 200 for a success).
	Status() int
	// Contains the status text of the response (e.g. usually an "OK" for a success).
	StatusText() string
	// Returns the text representation of response body.
	Text() (string, error)
	// Contains the URL of the response.
	URL() string
}

// Whenever a network route is set up with page.route(url, handler) or browserContext.route(url, handler), the `Route` object allows to
// handle the route.
type Route interface {
	// Aborts the route's request.
	Abort(options ...RouteAbortOptions) error
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
// touchscreen can only be used in browser contexts that have been intialized with `hasTouch` set to true.
type Touchscreen interface {
	// Dispatches a `touchstart` and `touchend` event with a single touch at the position (`x`,`y`).
	Tap(x int, y int) error
}

// The WebSocket class represents websocket connections in the page.
type WebSocket interface {
	EventEmitter
	// Indicates that the web socket has been closed.
	IsClosed() bool
	// Contains the URL of the WebSocket.
	URL() string
	// Returns the event data value.
	// Waits for event to fire and passes its value into the predicate function. Resolves when the predicate returns truthy
	// value. Will throw an error if the webSocket is closed before the event is fired.
	WaitForEvent(event string, predicate ...interface{}) interface{}
}

// When browser context is created with the `videosPath` option, each page has a video object associated with it.
type Video interface {
	// Returns the file system path this video will be recorded to. The video is guaranteed to be written to the filesystem
	// upon closing the browser context.
	Path() string
}

// The Worker class represents a WebWorker. `worker`
// event is emitted on the page object to signal a worker creation. `close` event is emitted on the worker object when the
// worker is gone.
type Worker interface {
	// Returns the return value of `pageFunction`
	// If the function passed to the `worker.evaluate` returns a Promise, then `worker.evaluate` would wait for the promise
	// to resolve and return its value.
	// If the function passed to the `worker.evaluate` returns a non-Serializable value, then `worker.evaluate` resolves to
	// `undefined`. DevTools Protocol also supports transferring some additional values that are not serializable by `JSON`:
	// `-0`, `NaN`, `Infinity`, `-Infinity`, and bigint literals.
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	// Returns the return value of `pageFunction` as in-page object (JSHandle).
	// The only difference between `worker.evaluate` and `worker.evaluateHandle` is that `worker.evaluateHandle` returns
	// in-page object (JSHandle).
	// If the function passed to the `worker.evaluateHandle` returns a Promise, then `worker.evaluateHandle` would wait for
	// the promise to resolve and return its value.
	EvaluateHandle(expression string, options ...interface{}) (JSHandle, error)
	URL() string
}
