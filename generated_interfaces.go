package playwright

type Browser interface {
	EventEmitter
	Close() error
	Contexts() []BrowserContext
	IsConnected() bool
	NewContext(options ...BrowserNewContextOptions) (BrowserContext, error)
	NewPage(options ...BrowserNewContextOptions) (Page, error)
	Version() string
}

type BrowserContext interface {
	EventEmitter
	AddCookies(cookies ...SetNetworkCookieParam) error
	AddInitScript(options BrowserContextAddInitScriptOptions) error
	Browser() Browser
	ClearCookies() error
	ClearPermissions() error
	Close() error
	Cookies(urls ...string) ([]*NetworkCookie, error)
	ExpectEvent(event string, cb func() error) (interface{}, error)
	GrantPermissions(permissions []string, options ...BrowserContextGrantPermissionsOptions) error
	NewPage(options ...BrowserNewPageOptions) (Page, error)
	Pages() []Page
	SetDefaultNavigationTimeout(timeout int)
	SetDefaultTimeout(timeout int)
	SetExtraHTTPHeaders(headers map[string]string) error
	SetGeolocation(gelocation *SetGeolocationOptions) error
	SetOffline(offline bool) error
	WaitForEvent(event string, predicate ...interface{}) interface{}
}

type BrowserType interface {
	ExecutablePath() string
	Launch(options ...BrowserTypeLaunchOptions) (Browser, error)
	LaunchPersistentContext(userDataDir string, options ...BrowserTypeLaunchPersistentContextOptions) (BrowserContext, error)
	Name() string
}

type ConsoleMessage interface {
	Args() []JSHandle
	Location() ConsoleMessageLocation
	String() string
	Text() string
	Type() string
}

type Dialog interface {
	Accept(texts ...string) error
	DefaultValue() string
	Dismiss() error
	Message() string
	Type() string
}

type Download interface {
	Delete() error
	Failure() error
	Path() (string, error)
	SaveAs(path string) error
	String() string
	SuggestedFilename() string
	URL() string
}

type ElementHandle interface {
	AsElement() ElementHandle
	BoundingBox() (*Rect, error)
	Check(options ...ElementHandleCheckOptions) error
	Click(options ...ElementHandleClickOptions) error
	ContentFrame() (Frame, error)
	Dblclick(options ...ElementHandleDblclickOptions) error
	DispatchEvent(typ string, initObjects ...interface{}) error
	EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
	EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
	Fill(value string, options ...ElementHandleFillOptions) error
	Focus() error
	GetAttribute(name string) (string, error)
	Hover(options ...ElementHandleHoverOptions) error
	InnerHTML() (string, error)
	InnerText() (string, error)
	OwnerFrame() (Frame, error)
	Press(options ...ElementHandlePressOptions) error
	QuerySelector(selector string) (ElementHandle, error)
	QuerySelectorAll(selector string) ([]ElementHandle, error)
	Screenshot(options ...ElementHandleScreenshotOptions) ([]byte, error)
	ScrollIntoViewIfNeeded(options ...ElementHandleScrollIntoViewIfNeededOptions) error
	SelectText(options ...ElementHandleSelectTextOptions) error
	SetInputFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error
	TextContent() (string, error)
	Type(value string, options ...ElementHandleTypeOptions) error
	Uncheck(options ...ElementHandleUncheckOptions) error
}

type EventEmitter interface {
	Emit(name string, payload ...interface{})
	ListenerCount(name string) int
	On(name string, handler interface{})
	Once(name string, handler interface{})
	RemoveListener(name string, handler interface{})
}

type FileChooser interface {
	Element() ElementHandle
	IsMultiple() bool
	Page() Page
	SetFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error
}

type Frame interface {
	AddScriptTag(options PageAddScriptTagOptions) (ElementHandle, error)
	AddStyleTag(options PageAddStyleTagOptions) (ElementHandle, error)
	Check(selector string, options ...FrameCheckOptions) error
	ChildFrames() []Frame
	Click(selector string, options ...PageClickOptions) error
	Content() (string, error)
	Dblclick(selector string, options ...FrameDblclickOptions) error
	DispatchEvent(selector, typ string, options ...PageDispatchEventOptions) error
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	EvaluateHandle(expression string, options ...interface{}) (interface{}, error)
	EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
	EvalOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
	Fill(selector string, value string, options ...FrameFillOptions) error
	Focus(selector string, options ...FrameFocusOptions) error
	FrameElement() (ElementHandle, error)
	GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error)
	Goto(url string, options ...PageGotoOptions) (Response, error)
	Hover(selector string, options ...PageHoverOptions) error
	InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error)
	InnerText(selector string, options ...PageInnerTextOptions) (string, error)
	IsDetached() bool
	Name() string
	Page() Page
	ParentFrame() Frame
	Press(selector, key string, options ...PagePressOptions) error
	QuerySelector(selector string) (ElementHandle, error)
	QuerySelectorAll(selector string) ([]ElementHandle, error)
	SetContent(content string, options ...PageSetContentOptions) error
	SetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) error
	TextContent(selector string, options ...FrameTextContentOptions) (string, error)
	Title() (string, error)
	Type(selector, text string, options ...PageTypeOptions) error
	URL() string
	Uncheck(selector string, options ...FrameUncheckOptions) error
	WaitForEvent(event string, predicate ...interface{}) interface{}
	WaitForEventCh(event string, predicate ...interface{}) <-chan interface{}
	WaitForFunction(expression string, options ...FrameWaitForFunctionOptions) (JSHandle, error)
	WaitForLoadState(given ...string)
	WaitForNavigation(options ...PageWaitForNavigationOptions) (Response, error)
	WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandle, error)
	WaitForTimeout(timeout int)
}

type JSHandle interface {
	AsElement() ElementHandle
	Dispose() error
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	EvaluateHandle(expression string, options ...interface{}) (interface{}, error)
	GetProperties() (map[string]JSHandle, error)
	GetProperty(name string) (JSHandle, error)
	JSONValue() (interface{}, error)
	String() string
}

type Keyboard interface {
	Down(key string) error
	InsertText(text string) error
	Press(key string, options ...KeyboardPressOptions) error
	Type(text string, options ...KeyboardTypeOptions) error
	Up(key string) error
}

type Mouse interface {
	Click(x, y float64, options ...MouseClickOptions) error
	Dblclick(x, y float64, options ...MouseDblclickOptions) error
	Down(options ...MouseDownOptions) error
	Move(x float64, y float64, options ...MouseMoveOptions) error
	Up(options ...MouseUpOptions) error
}

type Page interface {
	Mouse() Mouse
	Keyboard() Keyboard
	EventEmitter
	AddInitScript(options BrowserContextAddInitScriptOptions) error
	AddScriptTag(options PageAddScriptTagOptions) (ElementHandle, error)
	AddStyleTag(options PageAddStyleTagOptions) (ElementHandle, error)
	BringToFront() error
	Check(selector string, options ...FrameCheckOptions) error
	Click(selector string, options ...PageClickOptions) error
	Close(options ...PageCloseOptions) error
	Content() (string, error)
	Context() BrowserContext
	Dblclick(expression string, options ...FrameDblclickOptions) error
	DispatchEvent(selector string, typ string, options ...PageDispatchEventOptions) error
	EmulateMedia(options ...PageEmulateMediaOptions) error
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	EvaluateHandle(expression string, options ...interface{}) (interface{}, error)
	EvalOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
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
	Fill(selector, text string, options ...FrameFillOptions) error
	Focus(expression string, options ...FrameFocusOptions) error
	Frames() []Frame
	GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error)
	GoBack(options ...PageGoBackOptions) (Response, error)
	GoForward(options ...PageGoForwardOptions) (Response, error)
	Goto(url string, options ...PageGotoOptions) (Response, error)
	Hover(selector string, options ...PageHoverOptions) error
	InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error)
	InnerText(selector string, options ...PageInnerTextOptions) (string, error)
	Isclosed() bool
	MainFrame() Frame
	Opener() (Page, error)
	PDF(options ...PagePDFOptions) ([]byte, error)
	Press(selector, key string, options ...PagePressOptions) error
	QuerySelector(selector string) (ElementHandle, error)
	QuerySelectorAll(selector string) ([]ElementHandle, error)
	Reload(options ...PageReloadOptions) (Response, error)
	Route(url interface{}, handler routeHandler) error
	Screenshot(options ...PageScreenshotOptions) ([]byte, error)
	SetContent(content string, options ...PageSetContentOptions) error
	SetDefaultNavigationTimeout(timeout int)
	SetDefaultTimeout(timeout int)
	SetExtraHTTPHeaders(headers map[string]string) error
	SetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) error
	SetViewportSize(width, height int) error
	TextContent(selector string, options ...FrameTextContentOptions) (string, error)
	Title() (string, error)
	Type(selector, text string, options ...PageTypeOptions) error
	URL() string
	Uncheck(selector string, options ...FrameUncheckOptions) error
	ViewportSize() ViewportSize
	WaitForEvent(event string, predicate ...interface{}) interface{}
	WaitForFunction(expression string, options ...FrameWaitForFunctionOptions) (JSHandle, error)
	WaitForLoadState(state ...string)
	WaitForNavigation(options ...PageWaitForNavigationOptions) (Response, error)
	WaitForRequest(url interface{}, options ...interface{}) Request
	WaitForResponse(url interface{}, options ...interface{}) Response
	WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandle, error)
	WaitForTimeout(timeout int)
	Workers() []Worker
}

type Request interface {
	Failure() *RequestFailure
	Frame() Frame
	Headers() map[string]string
	IsNavigationRequest() bool
	Method() string
	PostData() (string, error)
	PostDataBuffer() ([]byte, error)
	PostDataJSON(v interface{}) error
	RedirectedFrom() Request
	RedirectedTo() Request
	ResourceType() string
	Response() (Response, error)
	URL() string
}

type Response interface {
	Body() ([]byte, error)
	Finished() error
	Frame() Frame
	Headers() map[string]string
	JSON(v interface{}) error
	Ok() bool
	Request() Request
	Status() int
	StatusText() string
	Text() (string, error)
	URL() string
}

type Route interface {
	Abort(errorCode *string) error
	Continue(options ...RouteContinueOptions) error
	Fulfill(options RouteFulfillOptions) error
	Request() Request
}

type WebSocket interface {
	URL() string
}

type Worker interface {
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	EvaluateHandle(expression string, options ...interface{}) (JSHandle, error)
	URL() string
}
