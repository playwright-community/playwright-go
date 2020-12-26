package playwright

type BrowserI interface {
	EventEmitterI
	Close() error
	Contexts() []BrowserContextI
	IsConnected() bool
	NewContext(options ...BrowserNewContextOptions) (BrowserContextI, error)
	NewPage(options ...BrowserNewContextOptions) (PageI, error)
	Version() string
}

type BrowserContextI interface {
	EventEmitterI
	AddCookies(cookies ...SetNetworkCookieParam) error
	AddInitScript(options BrowserContextAddInitScriptOptions) error
	ClearCookies() error
	ClearPermissions() error
	Close() error
	Cookies(urls ...string) ([]*NetworkCookie, error)
	ExpectEvent(event string, cb func() error) (interface{}, error)
	GrantPermissions(permissions []string, options ...BrowserContextGrantPermissionsOptions) error
	NewPage(options ...BrowserNewPageOptions) (PageI, error)
	Pages() []PageI
	SetDefaultNavigationTimeout(timeout int)
	SetDefaultTimeout(timeout int)
	SetExtraHTTPHeaders(headers map[string]string) error
	SetGeolocation(gelocation *SetGeolocationOptions) error
	SetOffline(offline bool) error
	WaitForEvent(event string, predicate ...interface{}) interface{}
}

type BrowserTypeI interface {
	ExecutablePath() string
	Launch(options ...BrowserTypeLaunchOptions) (BrowserI, error)
	LaunchPersistentContext(userDataDir string, options ...BrowserTypeLaunchPersistentContextOptions) (BrowserContextI, error)
	Name() string
}

type ConsoleMessageI interface {
	Args() []JSHandleI
	Location() ConsoleMessageLocation
	String() string
	Text() string
	Type() string
}

type DialogI interface {
	Accept(texts ...string) error
	DefaultValue() string
	Dismiss() error
	Message() string
	Type() string
}

type DownloadI interface {
	Delete() error
	Failure() error
	Path() (string, error)
	SaveAs(path string) error
	String() string
	SuggestedFilename() string
	URL() string
}

type ElementHandleI interface {
	AsElement() ElementHandleI
	BoundingBox() (*Rect, error)
	Check(options ...ElementHandleCheckOptions) error
	Click(options ...ElementHandleClickOptions) error
	ContentFrame() (FrameI, error)
	DblClick(options ...ElementHandleDblclickOptions) error
	DispatchEvent(typ string, initObjects ...interface{}) error
	EvaluateOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
	EvaluateOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
	Fill(value string, options ...ElementHandleFillOptions) error
	Focus() error
	GetAttribute(name string) (string, error)
	Hover(options ...ElementHandleHoverOptions) error
	InnerHTML() (string, error)
	InnerText() (string, error)
	OwnerFrame() (FrameI, error)
	Press(options ...ElementHandlePressOptions) error
	QuerySelector(selector string) (ElementHandleI, error)
	QuerySelectorAll(selector string) ([]ElementHandleI, error)
	Screenshot(options ...ElementHandleScreenshotOptions) ([]byte, error)
	ScrollIntoViewIfNeeded(options ...ElementHandleScrollIntoViewIfNeededOptions) error
	SelectText(options ...ElementHandleSelectTextOptions) error
	SetInputFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error
	TextContent() (string, error)
	Type(value string, options ...ElementHandleTypeOptions) error
	Uncheck(options ...ElementHandleUncheckOptions) error
}

type EventEmitterI interface {
	Emit(name string, payload ...interface{})
	ListenerCount(name string) int
	On(name string, handler interface{})
	Once(name string, handler interface{})
	RemoveListener(name string, handler interface{})
}

type FileChooserI interface {
	Element() ElementHandleI
	IsMultiple() bool
	Page() PageI
	SetFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error
}

type FrameI interface {
	AddScriptTag(options PageAddScriptTagOptions) (ElementHandleI, error)
	AddStyleTag(options PageAddStyleTagOptions) (ElementHandleI, error)
	Check(selector string, options ...FrameCheckOptions) error
	ChildFrames() []FrameI
	Click(selector string, options ...PageClickOptions) error
	Content() (string, error)
	DblClick(selector string, options ...FrameDblclickOptions) error
	DispatchEvent(selector, typ string, options ...PageDispatchEventOptions) error
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	EvaluateHandle(expression string, options ...interface{}) (interface{}, error)
	EvaluateOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
	EvaluateOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
	Fill(selector string, value string, options ...FrameFillOptions) error
	Focus(selector string, options ...FrameFocusOptions) error
	FrameElement() (ElementHandleI, error)
	GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error)
	Goto(url string, options ...PageGotoOptions) (ResponseI, error)
	Hover(selector string, options ...PageHoverOptions) error
	InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error)
	InnerText(selector string, options ...PageInnerTextOptions) (string, error)
	IsDetached() bool
	Name() string
	Page() PageI
	ParentFrame() FrameI
	Press(selector, key string, options ...PagePressOptions) error
	QuerySelector(selector string) (ElementHandleI, error)
	QuerySelectorAll(selector string) ([]ElementHandleI, error)
	SetContent(content string, options ...PageSetContentOptions) error
	SetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) error
	TextContent(selector string, options ...FrameTextContentOptions) (string, error)
	Title() (string, error)
	Type(selector, text string, options ...PageTypeOptions) error
	URL() string
	Uncheck(selector string, options ...FrameUncheckOptions) error
	WaitForEvent(event string, predicate ...interface{}) interface{}
	WaitForEventCh(event string, predicate ...interface{}) <-chan interface{}
	WaitForFunction(expression string, options ...FrameWaitForFunctionOptions) (JSHandleI, error)
	WaitForLoadState(given ...string)
	WaitForNavigation(options ...PageWaitForNavigationOptions) (ResponseI, error)
	WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandleI, error)
	WaitForTimeout(timeout int)
}

type JSHandleI interface {
	AsElement() ElementHandleI
	Dispose() error
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	EvaluateHandle(expression string, options ...interface{}) (interface{}, error)
	GetProperties() (map[string]JSHandleI, error)
	GetProperty(name string) (JSHandleI, error)
	JSONValue() (interface{}, error)
	String() string
}

type KeyboardI interface {
	Down(key string) error
	InsertText(text string) error
	Press(key string, options ...KeyboardPressOptions) error
	Type(text string, options ...KeyboardTypeOptions) error
	Up(key string) error
}

type MouseI interface {
	Click(x, y float64, options ...MouseClickOptions) error
	DblClick(x, y float64, options ...MouseDblclickOptions) error
	Down(options ...MouseDownOptions) error
	Move(x float64, y float64, options ...MouseMoveOptions) error
	Up(options ...MouseUpOptions) error
}

type PageI interface {
	Mouse() *Mouse
	Keyboard() *Keyboard
	EventEmitterI
	AddInitScript(options BrowserContextAddInitScriptOptions) error
	AddScriptTag(options PageAddScriptTagOptions) (ElementHandleI, error)
	AddStyleTag(options PageAddStyleTagOptions) (ElementHandleI, error)
	BringToFront() error
	Check(selector string, options ...FrameCheckOptions) error
	Click(selector string, options ...PageClickOptions) error
	Close(options ...PageCloseOptions) error
	Content() (string, error)
	Context() BrowserContextI
	DblClick(expression string, options ...FrameDblclickOptions) error
	DispatchEvent(selector string, typ string, options ...PageDispatchEventOptions) error
	EmulateMedia(options ...PageEmulateMediaOptions) error
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	EvaluateHandle(expression string, options ...interface{}) (interface{}, error)
	EvaluateOnSelector(selector string, expression string, options ...interface{}) (interface{}, error)
	EvaluateOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error)
	ExpectConsoleMessage(cb func() error) (ConsoleMessageI, error)
	ExpectDownload(cb func() error) (DownloadI, error)
	ExpectEvent(event string, cb func() error, predicates ...interface{}) (interface{}, error)
	ExpectFileChooser(cb func() error) (FileChooserI, error)
	ExpectLoadState(state string, cb func() error) (ConsoleMessageI, error)
	ExpectNavigation(cb func() error, options ...PageWaitForNavigationOptions) (ResponseI, error)
	ExpectPopup(cb func() error) (PageI, error)
	ExpectRequest(url interface{}, cb func() error, options ...interface{}) (RequestI, error)
	ExpectResponse(url interface{}, cb func() error, options ...interface{}) (ResponseI, error)
	ExpectWorker(cb func() error) (WorkerI, error)
	ExpectedDialog(cb func() error) (DialogI, error)
	Fill(selector, text string, options ...FrameFillOptions) error
	Focus(expression string, options ...FrameFocusOptions) error
	Frames() []FrameI
	GetAttribute(selector string, name string, options ...PageGetAttributeOptions) (string, error)
	GoBack(options ...PageGoBackOptions) (ResponseI, error)
	GoForward(options ...PageGoForwardOptions) (ResponseI, error)
	Goto(url string, options ...PageGotoOptions) (ResponseI, error)
	Hover(selector string, options ...PageHoverOptions) error
	InnerHTML(selector string, options ...PageInnerHTMLOptions) (string, error)
	InnerText(selector string, options ...PageInnerTextOptions) (string, error)
	Isclosed() bool
	MainFrame() FrameI
	Opener() (PageI, error)
	PDF(options ...PagePdfOptions) ([]byte, error)
	Press(selector, key string, options ...PagePressOptions) error
	QuerySelector(selector string) (ElementHandleI, error)
	QuerySelectorAll(selector string) ([]ElementHandleI, error)
	Reload(options ...PageReloadOptions) (ResponseI, error)
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
	WaitForFunction(expression string, options ...FrameWaitForFunctionOptions) (JSHandleI, error)
	WaitForLoadState(state ...string)
	WaitForNavigation(options ...PageWaitForNavigationOptions) (ResponseI, error)
	WaitForRequest(url interface{}, options ...interface{}) RequestI
	WaitForResponse(url interface{}, options ...interface{}) ResponseI
	WaitForSelector(selector string, options ...PageWaitForSelectorOptions) (ElementHandleI, error)
	WaitForTimeout(timeout int)
	Workers() []WorkerI
}

type RequestI interface {
	Failure() *RequestFailure
	Frame() FrameI
	Headers() map[string]string
	IsNavigationRequest() bool
	Method() string
	PostData() (string, error)
	PostDataBuffer() ([]byte, error)
	PostDataJSON(v interface{}) error
	RedirectedFrom() RequestI
	RedirectedTo() RequestI
	ResourceType() string
	Response() (ResponseI, error)
	URL() string
}

type ResponseI interface {
	Body() ([]byte, error)
	Finished() error
	Frame() FrameI
	Headers() map[string]string
	JSON(v interface{}) error
	Ok() bool
	Request() RequestI
	Status() int
	StatusText() string
	Text() (string, error)
	URL() string
}

type RouteI interface {
	Abort(errorCode *string) error
	Continue(options ...RouteContinueOptions) error
	Fulfill(options RouteFulfillOptions) error
	Request() RequestI
}

type WebSocketI interface {
	URL() string
}

type WorkerI interface {
	Evaluate(expression string, options ...interface{}) (interface{}, error)
	EvaluateHandle(expression string, options ...interface{}) (JSHandleI, error)
	URL() string
}
