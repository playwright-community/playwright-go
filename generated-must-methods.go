package playwright

func (t *browserImpl) MustClose() {
	err := t.Close()
	if err != nil {
		panic(err)
	}
}

func (t *browserImpl) MustNewContext(options ...BrowserNewContextOptions) BrowserContext {
	result, err := t.NewContext(options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *browserImpl) MustNewPage(options ...BrowserNewContextOptions) Page {
	result, err := t.NewPage(options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *browserImpl) MustNewBrowserCDPSession() CDPSession {
	result, err := t.NewBrowserCDPSession()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *cdpsessionImpl) MustDetach() {
	err := t.Detach()
	if err != nil {
		panic(err)
	}
}

func (t *cdpsessionImpl) MustSend(method string, params map[string]interface{}) interface{} {
	result, err := t.Send(method, params)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *browsercontextImpl) MustAddCookies(cookies ...BrowserContextAddCookiesOptionsCookies) {
	err := t.AddCookies(cookies...)
	if err != nil {
		panic(err)
	}
}

func (t *browsercontextImpl) MustAddInitScript(script BrowserContextAddInitScriptOptions) {
	err := t.AddInitScript(script)
	if err != nil {
		panic(err)
	}
}

func (t *browsercontextImpl) MustClearCookies() {
	err := t.ClearCookies()
	if err != nil {
		panic(err)
	}
}

func (t *browsercontextImpl) MustClearPermissions() {
	err := t.ClearPermissions()
	if err != nil {
		panic(err)
	}
}

func (t *browsercontextImpl) MustClose() {
	err := t.Close()
	if err != nil {
		panic(err)
	}
}

func (t *browsercontextImpl) MustCookies(urls ...string) []*BrowserContextCookiesResult {
	result, err := t.Cookies(urls...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *browsercontextImpl) MustExpectEvent(event string, cb func() error) interface{} {
	result, err := t.ExpectEvent(event, cb)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *browsercontextImpl) MustExposeBinding(name string, binding BindingCallFunction, handle ...bool) {
	err := t.ExposeBinding(name, binding, handle...)
	if err != nil {
		panic(err)
	}
}

func (t *browsercontextImpl) MustExposeFunction(name string, binding ExposedFunction) {
	err := t.ExposeFunction(name, binding)
	if err != nil {
		panic(err)
	}
}

func (t *browsercontextImpl) MustGrantPermissions(permissions []string, options ...BrowserContextGrantPermissionsOptions) {
	err := t.GrantPermissions(permissions, options...)
	if err != nil {
		panic(err)
	}
}

func (t *browsercontextImpl) MustNewCDPSession(page Page) CDPSession {
	result, err := t.NewCDPSession(page)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *browsercontextImpl) MustNewPage(options ...BrowserNewPageOptions) Page {
	result, err := t.NewPage(options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *browsercontextImpl) MustSetExtraHTTPHeaders(headers map[string]string) {
	err := t.SetExtraHTTPHeaders(headers)
	if err != nil {
		panic(err)
	}
}

func (t *browsercontextImpl) MustSetGeolocation(gelocation *SetGeolocationOptions) {
	err := t.SetGeolocation(gelocation)
	if err != nil {
		panic(err)
	}
}

func (t *browsercontextImpl) MustResetGeolocation() {
	err := t.ResetGeolocation()
	if err != nil {
		panic(err)
	}
}

func (t *browsercontextImpl) MustRoute(url interface{}, handler routeHandler) {
	err := t.Route(url, handler)
	if err != nil {
		panic(err)
	}
}

func (t *browsercontextImpl) MustSetOffline(offline bool) {
	err := t.SetOffline(offline)
	if err != nil {
		panic(err)
	}
}

func (t *browsercontextImpl) MustStorageState(path ...string) *StorageState {
	result, err := t.StorageState(path...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *browsercontextImpl) MustUnroute(url interface{}, handler ...routeHandler) {
	err := t.Unroute(url, handler...)
	if err != nil {
		panic(err)
	}
}

func (t *tracingImpl) MustStart(options ...TracingStartOptions) {
	err := t.Start(options...)
	if err != nil {
		panic(err)
	}
}

func (t *tracingImpl) MustStop(options ...TracingStopOptions) {
	err := t.Stop(options...)
	if err != nil {
		panic(err)
	}
}

func (t *tracingImpl) MustStartChunk() {
	err := t.StartChunk()
	if err != nil {
		panic(err)
	}
}

func (t *tracingImpl) MustStopChunk(options ...TracingStopChunkOptions) {
	err := t.StopChunk(options...)
	if err != nil {
		panic(err)
	}
}

func (t *browsertypeImpl) MustLaunch(options ...BrowserTypeLaunchOptions) Browser {
	result, err := t.Launch(options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *browsertypeImpl) MustLaunchPersistentContext(userDataDir string, options ...BrowserTypeLaunchPersistentContextOptions) BrowserContext {
	result, err := t.LaunchPersistentContext(userDataDir, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *browsertypeImpl) MustConnect(url string, options ...BrowserTypeConnectOptions) Browser {
	result, err := t.Connect(url, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *browsertypeImpl) MustConnectOverCDP(endpointURL string, options ...BrowserTypeConnectOverCDPOptions) Browser {
	result, err := t.ConnectOverCDP(endpointURL, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *dialogImpl) MustAccept(texts ...string) {
	err := t.Accept(texts...)
	if err != nil {
		panic(err)
	}
}

func (t *dialogImpl) MustDismiss() {
	err := t.Dismiss()
	if err != nil {
		panic(err)
	}
}

func (t *downloadImpl) MustDelete() {
	err := t.Delete()
	if err != nil {
		panic(err)
	}
}

func (t *downloadImpl) MustFailure() string {
	result, err := t.Failure()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *downloadImpl) MustPath() string {
	result, err := t.Path()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *downloadImpl) MustSaveAs(path string) {
	err := t.SaveAs(path)
	if err != nil {
		panic(err)
	}
}

func (t *downloadImpl) MustCancel() {
	err := t.Cancel()
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustBoundingBox() *Rect {
	result, err := t.BoundingBox()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustCheck(options ...ElementHandleCheckOptions) {
	err := t.Check(options...)
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustClick(options ...ElementHandleClickOptions) {
	err := t.Click(options...)
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustContentFrame() Frame {
	result, err := t.ContentFrame()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustDblclick(options ...ElementHandleDblclickOptions) {
	err := t.Dblclick(options...)
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustDispatchEvent(typ string, initObjects ...interface{}) {
	err := t.DispatchEvent(typ, initObjects...)
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustEvalOnSelector(selector string, expression string, options ...interface{}) interface{} {
	result, err := t.EvalOnSelector(selector, expression, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustEvalOnSelectorAll(selector string, expression string, options ...interface{}) interface{} {
	result, err := t.EvalOnSelectorAll(selector, expression, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustFill(value string, options ...ElementHandleFillOptions) {
	err := t.Fill(value, options...)
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustFocus() {
	err := t.Focus()
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustGetAttribute(name string) string {
	result, err := t.GetAttribute(name)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustHover(options ...ElementHandleHoverOptions) {
	err := t.Hover(options...)
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustInnerHTML() string {
	result, err := t.InnerHTML()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustInnerText() string {
	result, err := t.InnerText()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustIsChecked() bool {
	result, err := t.IsChecked()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustIsDisabled() bool {
	result, err := t.IsDisabled()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustIsEditable() bool {
	result, err := t.IsEditable()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustIsEnabled() bool {
	result, err := t.IsEnabled()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustIsHidden() bool {
	result, err := t.IsHidden()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustIsVisible() bool {
	result, err := t.IsVisible()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustOwnerFrame() Frame {
	result, err := t.OwnerFrame()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustPress(key string, options ...ElementHandlePressOptions) {
	err := t.Press(key, options...)
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustQuerySelector(selector string) ElementHandle {
	result, err := t.QuerySelector(selector)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustQuerySelectorAll(selector string) []ElementHandle {
	result, err := t.QuerySelectorAll(selector)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustScreenshot(options ...ElementHandleScreenshotOptions) []byte {
	result, err := t.Screenshot(options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustScrollIntoViewIfNeeded(options ...ElementHandleScrollIntoViewIfNeededOptions) {
	err := t.ScrollIntoViewIfNeeded(options...)
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustSelectOption(values SelectOptionValues, options ...ElementHandleSelectOptionOptions) []string {
	result, err := t.SelectOption(values, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustSelectText(options ...ElementHandleSelectTextOptions) {
	err := t.SelectText(options...)
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustSetInputFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) {
	err := t.SetInputFiles(files, options...)
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustTap(options ...ElementHandleTapOptions) {
	err := t.Tap(options...)
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustTextContent() string {
	result, err := t.TextContent()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustType(value string, options ...ElementHandleTypeOptions) {
	err := t.Type(value, options...)
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustUncheck(options ...ElementHandleUncheckOptions) {
	err := t.Uncheck(options...)
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustWaitForElementState(state string, options ...ElementHandleWaitForElementStateOptions) {
	err := t.WaitForElementState(state, options...)
	if err != nil {
		panic(err)
	}
}

func (t *elementhandleImpl) MustWaitForSelector(selector string, options ...ElementHandleWaitForSelectorOptions) ElementHandle {
	result, err := t.WaitForSelector(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *elementhandleImpl) MustInputValue(options ...ElementHandleInputValueOptions) string {
	result, err := t.InputValue(options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *filechooserImpl) MustSetFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) {
	err := t.SetFiles(files, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustAddScriptTag(options PageAddScriptTagOptions) ElementHandle {
	result, err := t.AddScriptTag(options)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustAddStyleTag(options PageAddStyleTagOptions) ElementHandle {
	result, err := t.AddStyleTag(options)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustCheck(selector string, options ...FrameCheckOptions) {
	err := t.Check(selector, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustClick(selector string, options ...PageClickOptions) {
	err := t.Click(selector, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustContent() string {
	result, err := t.Content()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustDblclick(selector string, options ...FrameDblclickOptions) {
	err := t.Dblclick(selector, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustDispatchEvent(selector, typ string, eventInit interface{}, options ...PageDispatchEventOptions) {
	err := t.DispatchEvent(selector, typ, eventInit, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustEvaluate(expression string, options ...interface{}) interface{} {
	result, err := t.Evaluate(expression, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustEvaluateHandle(expression string, options ...interface{}) JSHandle {
	result, err := t.EvaluateHandle(expression, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustEvalOnSelector(selector string, expression string, options ...interface{}) interface{} {
	result, err := t.EvalOnSelector(selector, expression, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustEvalOnSelectorAll(selector string, expression string, options ...interface{}) interface{} {
	result, err := t.EvalOnSelectorAll(selector, expression, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustFill(selector string, value string, options ...FrameFillOptions) {
	err := t.Fill(selector, value, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustFocus(selector string, options ...FrameFocusOptions) {
	err := t.Focus(selector, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustFrameElement() ElementHandle {
	result, err := t.FrameElement()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustGetAttribute(selector string, name string, options ...PageGetAttributeOptions) string {
	result, err := t.GetAttribute(selector, name, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustGoto(url string, options ...PageGotoOptions) Response {
	result, err := t.Goto(url, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustHover(selector string, options ...PageHoverOptions) {
	err := t.Hover(selector, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustInnerHTML(selector string, options ...PageInnerHTMLOptions) string {
	result, err := t.InnerHTML(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustInnerText(selector string, options ...PageInnerTextOptions) string {
	result, err := t.InnerText(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustIsChecked(selector string, options ...FrameIsCheckedOptions) bool {
	result, err := t.IsChecked(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustIsDisabled(selector string, options ...FrameIsDisabledOptions) bool {
	result, err := t.IsDisabled(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustIsEditable(selector string, options ...FrameIsEditableOptions) bool {
	result, err := t.IsEditable(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustIsEnabled(selector string, options ...FrameIsEnabledOptions) bool {
	result, err := t.IsEnabled(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustIsHidden(selector string, options ...FrameIsHiddenOptions) bool {
	result, err := t.IsHidden(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustIsVisible(selector string, options ...FrameIsVisibleOptions) bool {
	result, err := t.IsVisible(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustPress(selector, key string, options ...PagePressOptions) {
	err := t.Press(selector, key, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustQuerySelector(selector string) ElementHandle {
	result, err := t.QuerySelector(selector)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustQuerySelectorAll(selector string) []ElementHandle {
	result, err := t.QuerySelectorAll(selector)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustSetContent(content string, options ...PageSetContentOptions) {
	err := t.SetContent(content, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustSelectOption(selector string, values SelectOptionValues, options ...FrameSelectOptionOptions) []string {
	result, err := t.SelectOption(selector, values, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustSetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) {
	err := t.SetInputFiles(selector, files, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustTap(selector string, options ...FrameTapOptions) {
	err := t.Tap(selector, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustTextContent(selector string, options ...FrameTextContentOptions) string {
	result, err := t.TextContent(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustTitle() string {
	result, err := t.Title()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustType(selector, text string, options ...PageTypeOptions) {
	err := t.Type(selector, text, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustUncheck(selector string, options ...FrameUncheckOptions) {
	err := t.Uncheck(selector, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustWaitForFunction(expression string, arg interface{}, options ...FrameWaitForFunctionOptions) JSHandle {
	result, err := t.WaitForFunction(expression, arg, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustWaitForNavigation(options ...PageWaitForNavigationOptions) Response {
	result, err := t.WaitForNavigation(options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustWaitForURL(url string, options ...FrameWaitForURLOptions) {
	err := t.WaitForURL(url, options...)
	if err != nil {
		panic(err)
	}
}

func (t *frameImpl) MustWaitForSelector(selector string, options ...PageWaitForSelectorOptions) ElementHandle {
	result, err := t.WaitForSelector(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustInputValue(selector string, options ...FrameInputValueOptions) string {
	result, err := t.InputValue(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *frameImpl) MustDragAndDrop(source, target string, options ...FrameDragAndDropOptions) {
	err := t.DragAndDrop(source, target, options...)
	if err != nil {
		panic(err)
	}
}

func (t *jshandleImpl) MustDispose() {
	err := t.Dispose()
	if err != nil {
		panic(err)
	}
}

func (t *jshandleImpl) MustEvaluate(expression string, options ...interface{}) interface{} {
	result, err := t.Evaluate(expression, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *jshandleImpl) MustEvaluateHandle(expression string, options ...interface{}) JSHandle {
	result, err := t.EvaluateHandle(expression, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *jshandleImpl) MustGetProperties() map[string]JSHandle {
	result, err := t.GetProperties()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *jshandleImpl) MustGetProperty(name string) JSHandle {
	result, err := t.GetProperty(name)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *jshandleImpl) MustJSONValue() interface{} {
	result, err := t.JSONValue()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *keyboardImpl) MustDown(key string) {
	err := t.Down(key)
	if err != nil {
		panic(err)
	}
}

func (t *keyboardImpl) MustInsertText(text string) {
	err := t.InsertText(text)
	if err != nil {
		panic(err)
	}
}

func (t *keyboardImpl) MustPress(key string, options ...KeyboardPressOptions) {
	err := t.Press(key, options...)
	if err != nil {
		panic(err)
	}
}

func (t *keyboardImpl) MustType(text string, options ...KeyboardTypeOptions) {
	err := t.Type(text, options...)
	if err != nil {
		panic(err)
	}
}

func (t *keyboardImpl) MustUp(key string) {
	err := t.Up(key)
	if err != nil {
		panic(err)
	}
}

func (t *mouseImpl) MustClick(x, y float64, options ...MouseClickOptions) {
	err := t.Click(x, y, options...)
	if err != nil {
		panic(err)
	}
}

func (t *mouseImpl) MustDblclick(x, y float64, options ...MouseDblclickOptions) {
	err := t.Dblclick(x, y, options...)
	if err != nil {
		panic(err)
	}
}

func (t *mouseImpl) MustDown(options ...MouseDownOptions) {
	err := t.Down(options...)
	if err != nil {
		panic(err)
	}
}

func (t *mouseImpl) MustMove(x float64, y float64, options ...MouseMoveOptions) {
	err := t.Move(x, y, options...)
	if err != nil {
		panic(err)
	}
}

func (t *mouseImpl) MustUp(options ...MouseUpOptions) {
	err := t.Up(options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustAddInitScript(script PageAddInitScriptOptions) {
	err := t.AddInitScript(script)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustAddScriptTag(options PageAddScriptTagOptions) ElementHandle {
	result, err := t.AddScriptTag(options)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustAddStyleTag(options PageAddStyleTagOptions) ElementHandle {
	result, err := t.AddStyleTag(options)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustBringToFront() {
	err := t.BringToFront()
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustCheck(selector string, options ...FrameCheckOptions) {
	err := t.Check(selector, options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustClick(selector string, options ...PageClickOptions) {
	err := t.Click(selector, options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustClose(options ...PageCloseOptions) {
	err := t.Close(options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustContent() string {
	result, err := t.Content()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustDblclick(expression string, options ...FrameDblclickOptions) {
	err := t.Dblclick(expression, options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustDispatchEvent(selector string, typ string, options ...PageDispatchEventOptions) {
	err := t.DispatchEvent(selector, typ, options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustExposeBinding(name string, binding BindingCallFunction, handle ...bool) {
	err := t.ExposeBinding(name, binding, handle...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustExposeFunction(name string, binding ExposedFunction) {
	err := t.ExposeFunction(name, binding)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustEmulateMedia(options ...PageEmulateMediaOptions) {
	err := t.EmulateMedia(options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustEvaluate(expression string, options ...interface{}) interface{} {
	result, err := t.Evaluate(expression, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustEvaluateHandle(expression string, options ...interface{}) JSHandle {
	result, err := t.EvaluateHandle(expression, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustEvalOnSelector(selector string, expression string, options ...interface{}) interface{} {
	result, err := t.EvalOnSelector(selector, expression, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustEvalOnSelectorAll(selector string, expression string, options ...interface{}) interface{} {
	result, err := t.EvalOnSelectorAll(selector, expression, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustExpectConsoleMessage(cb func() error) ConsoleMessage {
	result, err := t.ExpectConsoleMessage(cb)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustExpectDownload(cb func() error) Download {
	result, err := t.ExpectDownload(cb)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustExpectEvent(event string, cb func() error, predicates ...interface{}) interface{} {
	result, err := t.ExpectEvent(event, cb, predicates...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustExpectFileChooser(cb func() error) FileChooser {
	result, err := t.ExpectFileChooser(cb)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustExpectLoadState(state string, cb func() error) {
	err := t.ExpectLoadState(state, cb)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustExpectNavigation(cb func() error, options ...PageWaitForNavigationOptions) Response {
	result, err := t.ExpectNavigation(cb, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustExpectPopup(cb func() error) Page {
	result, err := t.ExpectPopup(cb)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustExpectRequest(url interface{}, cb func() error, options ...interface{}) Request {
	result, err := t.ExpectRequest(url, cb, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustExpectResponse(url interface{}, cb func() error, options ...interface{}) Response {
	result, err := t.ExpectResponse(url, cb, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustExpectWorker(cb func() error) Worker {
	result, err := t.ExpectWorker(cb)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustExpectedDialog(cb func() error) Dialog {
	result, err := t.ExpectedDialog(cb)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustFill(selector, text string, options ...FrameFillOptions) {
	err := t.Fill(selector, text, options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustFocus(expression string, options ...FrameFocusOptions) {
	err := t.Focus(expression, options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustGetAttribute(selector string, name string, options ...PageGetAttributeOptions) string {
	result, err := t.GetAttribute(selector, name, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustGoBack(options ...PageGoBackOptions) Response {
	result, err := t.GoBack(options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustGoForward(options ...PageGoForwardOptions) Response {
	result, err := t.GoForward(options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustGoto(url string, options ...PageGotoOptions) Response {
	result, err := t.Goto(url, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustHover(selector string, options ...PageHoverOptions) {
	err := t.Hover(selector, options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustInnerHTML(selector string, options ...PageInnerHTMLOptions) string {
	result, err := t.InnerHTML(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustInnerText(selector string, options ...PageInnerTextOptions) string {
	result, err := t.InnerText(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustIsChecked(selector string, options ...FrameIsCheckedOptions) bool {
	result, err := t.IsChecked(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustIsDisabled(selector string, options ...FrameIsDisabledOptions) bool {
	result, err := t.IsDisabled(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustIsEditable(selector string, options ...FrameIsEditableOptions) bool {
	result, err := t.IsEditable(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustIsEnabled(selector string, options ...FrameIsEnabledOptions) bool {
	result, err := t.IsEnabled(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustIsHidden(selector string, options ...FrameIsHiddenOptions) bool {
	result, err := t.IsHidden(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustIsVisible(selector string, options ...FrameIsVisibleOptions) bool {
	result, err := t.IsVisible(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustOpener() Page {
	result, err := t.Opener()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustPDF(options ...PagePdfOptions) []byte {
	result, err := t.PDF(options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustPress(selector, key string, options ...PagePressOptions) {
	err := t.Press(selector, key, options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustQuerySelector(selector string) ElementHandle {
	result, err := t.QuerySelector(selector)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustQuerySelectorAll(selector string) []ElementHandle {
	result, err := t.QuerySelectorAll(selector)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustReload(options ...PageReloadOptions) Response {
	result, err := t.Reload(options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustRoute(url interface{}, handler routeHandler) {
	err := t.Route(url, handler)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustScreenshot(options ...PageScreenshotOptions) []byte {
	result, err := t.Screenshot(options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustSelectOption(selector string, values SelectOptionValues, options ...FrameSelectOptionOptions) []string {
	result, err := t.SelectOption(selector, values, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustSetContent(content string, options ...PageSetContentOptions) {
	err := t.SetContent(content, options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustSetExtraHTTPHeaders(headers map[string]string) {
	err := t.SetExtraHTTPHeaders(headers)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustSetInputFiles(selector string, files []InputFile, options ...FrameSetInputFilesOptions) {
	err := t.SetInputFiles(selector, files, options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustSetViewportSize(width, height int) {
	err := t.SetViewportSize(width, height)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustTap(selector string, options ...FrameTapOptions) {
	err := t.Tap(selector, options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustTextContent(selector string, options ...FrameTextContentOptions) string {
	result, err := t.TextContent(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustTitle() string {
	result, err := t.Title()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustType(selector, text string, options ...PageTypeOptions) {
	err := t.Type(selector, text, options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustUncheck(selector string, options ...FrameUncheckOptions) {
	err := t.Uncheck(selector, options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustUnroute(url interface{}, handler ...routeHandler) {
	err := t.Unroute(url, handler...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustWaitForFunction(expression string, arg interface{}, options ...FrameWaitForFunctionOptions) JSHandle {
	result, err := t.WaitForFunction(expression, arg, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustWaitForNavigation(options ...PageWaitForNavigationOptions) Response {
	result, err := t.WaitForNavigation(options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustWaitForSelector(selector string, options ...PageWaitForSelectorOptions) ElementHandle {
	result, err := t.WaitForSelector(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustDragAndDrop(source, target string, options ...FrameDragAndDropOptions) {
	err := t.DragAndDrop(source, target, options...)
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustPause() {
	err := t.Pause()
	if err != nil {
		panic(err)
	}
}

func (t *pageImpl) MustInputValue(selector string, options ...FrameInputValueOptions) string {
	result, err := t.InputValue(selector, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *pageImpl) MustWaitForURL(url string, options ...FrameWaitForURLOptions) {
	err := t.WaitForURL(url, options...)
	if err != nil {
		panic(err)
	}
}

func (t *requestImpl) MustAllHeaders() map[string]string {
	result, err := t.AllHeaders()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *requestImpl) MustHeadersArray() HeadersArray {
	result, err := t.HeadersArray()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *requestImpl) MustHeaderValue(name string) string {
	result, err := t.HeaderValue(name)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *requestImpl) MustHeaderValues(name string) []string {
	result, err := t.HeaderValues(name)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *requestImpl) MustPostData() string {
	result, err := t.PostData()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *requestImpl) MustPostDataBuffer() []byte {
	result, err := t.PostDataBuffer()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *requestImpl) MustPostDataJSON(v interface{}) {
	err := t.PostDataJSON(v)
	if err != nil {
		panic(err)
	}
}

func (t *requestImpl) MustResponse() Response {
	result, err := t.Response()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *requestImpl) MustSizes() *RequestSizesResult {
	result, err := t.Sizes()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *responseImpl) MustAllHeaders() map[string]string {
	result, err := t.AllHeaders()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *responseImpl) MustHeadersArray() HeadersArray {
	result, err := t.HeadersArray()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *responseImpl) MustHeaderValue(name string) string {
	result, err := t.HeaderValue(name)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *responseImpl) MustHeaderValues(name string) []string {
	result, err := t.HeaderValues(name)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *responseImpl) MustBody() []byte {
	result, err := t.Body()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *responseImpl) MustJSON(v interface{}) {
	err := t.JSON(v)
	if err != nil {
		panic(err)
	}
}

func (t *responseImpl) MustText() string {
	result, err := t.Text()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *responseImpl) MustSecurityDetails() *ResponseSecurityDetailsResult {
	result, err := t.SecurityDetails()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *responseImpl) MustServerAddr() *ResponseServerAddrResult {
	result, err := t.ServerAddr()
	if err != nil {
		panic(err)
	}
	return result
}

func (t *routeImpl) MustAbort(errorCode ...string) {
	err := t.Abort(errorCode...)
	if err != nil {
		panic(err)
	}
}

func (t *routeImpl) MustContinue(options ...RouteContinueOptions) {
	err := t.Continue(options...)
	if err != nil {
		panic(err)
	}
}

func (t *routeImpl) MustFulfill(options RouteFulfillOptions) {
	err := t.Fulfill(options)
	if err != nil {
		panic(err)
	}
}

func (t *touchscreenImpl) MustTap(x int, y int) {
	err := t.Tap(x, y)
	if err != nil {
		panic(err)
	}
}

func (t *videoImpl) MustDelete() {
	err := t.Delete()
	if err != nil {
		panic(err)
	}
}

func (t *videoImpl) MustSaveAs(path string) {
	err := t.SaveAs(path)
	if err != nil {
		panic(err)
	}
}

func (t *workerImpl) MustEvaluate(expression string, options ...interface{}) interface{} {
	result, err := t.Evaluate(expression, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *workerImpl) MustEvaluateHandle(expression string, options ...interface{}) JSHandle {
	result, err := t.EvaluateHandle(expression, options...)
	if err != nil {
		panic(err)
	}
	return result
}

func (t *workerImpl) MustExpectEvent(event string, cb func() error, predicates ...interface{}) interface{} {
	result, err := t.ExpectEvent(event, cb, predicates...)
	if err != nil {
		panic(err)
	}
	return result
}
