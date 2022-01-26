package playwright_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/h2non/filetype"
	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestPageURL(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.Equal(t, "about:blank", page.URL())
	_, err := page.Goto("https://example.com")
	require.NoError(t, err)
	require.Equal(t, "https://example.com/", page.URL())
	require.Equal(t, context, page.Context())
	require.Equal(t, 1, len(page.Frames()))
}

func TestPageSetContent(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent("<h1>foo</h1>",
		playwright.PageSetContentOptions{
			WaitUntil: playwright.WaitUntilStateNetworkidle,
		}))
	content, err := page.Content()
	require.NoError(t, err)
	require.Equal(t, content, "<html><head></head><body><h1>foo</h1></body></html>")
}

func TestPageScreenshot(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)

	require.NoError(t, page.SetContent("<h1>foobar</h1>"))
	tmpfile, err := ioutil.TempDir("", "screenshot")
	require.NoError(t, err)
	screenshotPath := filepath.Join(tmpfile, "image.png")
	screenshot, err := page.Screenshot()
	require.NoError(t, err)
	require.True(t, filetype.IsImage(screenshot))
	require.Greater(t, len(screenshot), 50)

	screenshot, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String(screenshotPath),
	})
	require.NoError(t, err)
	require.True(t, filetype.IsImage(screenshot))
	require.Greater(t, len(screenshot), 50)

	_, err = os.Stat(screenshotPath)
	require.NoError(t, err)
}

func TestPagePDF(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	if !isChromium {
		t.Skip("Skipping")
	}
	require.NoError(t, page.SetContent("<h1>foobar</h1>"))
	tmpfile, err := ioutil.TempDir("", "pdf")
	require.NoError(t, err)
	screenshotPath := filepath.Join(tmpfile, "image.png")
	screenshot, err := page.PDF()
	require.NoError(t, err)
	require.Equal(t, "application/pdf", http.DetectContentType(screenshot))
	require.Greater(t, len(screenshot), 50)

	screenshot, err = page.PDF(playwright.PagePdfOptions{
		Path: playwright.String(screenshotPath),
	})
	require.NoError(t, err)
	require.Equal(t, "application/pdf", http.DetectContentType(screenshot))
	require.Greater(t, len(screenshot), 50)

	_, err = os.Stat(screenshotPath)
	require.NoError(t, err)
}

func TestPageQuerySelector(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`<div id="one">
		<span id="two">
			<div id="three">
				<span id="four">
					foobar
				</span>
			</div>
		</span>
	</div>`))
	one, err := page.QuerySelector("div#one")
	require.NoError(t, err)
	two, err := one.QuerySelector("span#two")
	require.NoError(t, err)
	three, err := two.QuerySelector("div#three")
	require.NoError(t, err)
	four, err := three.QuerySelector("span#four")
	require.NoError(t, err)
	textContent, err := four.TextContent()
	require.NoError(t, err)
	require.Equal(t, strings.TrimSpace(textContent), "foobar")
}

func TestPageQuerySelectorAll(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`
	<div class="foo">0</div>
	<div class="foo">1</div>
	<div class="foo">2</div>
	`))
	elements, err := page.QuerySelectorAll("div.foo")
	require.NoError(t, err)
	require.Equal(t, 3, len(elements))
	for i := 0; i < 3; i++ {
		textContent, err := elements[i].TextContent()
		require.NoError(t, err)
		require.Equal(t, strconv.Itoa(i), textContent)
	}
}

func TestPageEvaluate(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	val, err := page.Evaluate(`() => 123`)
	require.NoError(t, err)
	require.Equal(t, val, 123)
}

func TestPageEvalOnSelectorAll(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
		<div class="foo">1</div>
		<div class="foo">2</div>
		<div class="foo">3</div>
	`)
	require.NoError(t, err)
	val, err := page.EvalOnSelectorAll(".foo", `(elements) => elements.map(el => el.textContent)`)
	require.NoError(t, err)
	require.Equal(t, val, []interface{}([]interface{}{"1", "2", "3"}))
}

func TestPageEvalOnSelector(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
		<div class="foo">bar</div>
	`)
	require.NoError(t, err)
	val, err := page.EvalOnSelector(".foo", `(element) => element.textContent`)
	require.NoError(t, err)
	require.Equal(t, val, "bar")
}

func TestPageExpectWorker(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	worker, err := page.ExpectWorker(func() error {
		_, err := page.Goto(server.PREFIX + "/worker/worker.html")
		return err
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(page.Workers()))
	require.Equal(t, worker, page.Workers()[0])
	worker = page.Workers()[0]
	require.Contains(t, worker.URL(), "worker.js")
	res, err := worker.Evaluate(`() => self["workerFunction"]()`)
	require.NoError(t, err)
	require.Equal(t, "worker function result", res)
	_, err = page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.Equal(t, 0, len(page.Workers()))
}

func TestPageExpectRequest(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	request, err := page.ExpectRequest("**/*", func() error {
		_, err := page.Goto(server.EMPTY_PAGE)
		return err
	})
	require.NoError(t, err)
	require.Equal(t, server.EMPTY_PAGE, request.URL())
	require.Equal(t, "document", request.ResourceType())
	require.Equal(t, "GET", request.Method())
}

func TestPageExpectRequestRegexp(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	request, err := page.ExpectRequest(regexp.MustCompile(".*/empty.html"), func() error {
		_, err := page.Goto(server.EMPTY_PAGE)
		return err
	})
	require.NoError(t, err)
	require.Equal(t, server.EMPTY_PAGE, request.URL())
	require.Equal(t, "document", request.ResourceType())
	require.Equal(t, "GET", request.Method())
}

func TestPageExpectRequestFunc(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	request, err := page.ExpectRequest(func(url string) bool {
		return strings.HasSuffix(url, "empty.html")
	}, func() error {
		_, err := page.Goto(server.EMPTY_PAGE)
		return err
	})
	require.NoError(t, err)
	require.Equal(t, server.EMPTY_PAGE, request.URL())
	require.Equal(t, "document", request.ResourceType())
	require.Equal(t, "GET", request.Method())
}

func TestPageExpectResponse(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	response, err := page.ExpectResponse("**/*", func() error {
		_, err := page.Goto(server.EMPTY_PAGE)
		return err
	})
	require.NoError(t, err)
	require.Equal(t, server.EMPTY_PAGE, response.URL())
	require.True(t, response.Ok())
	require.Equal(t, 200, response.Status())
	require.Equal(t, "OK", response.StatusText())
}

func TestPageExpectPopup(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	popup, err := page.ExpectPopup(func() error {
		_, err := page.Evaluate(`window._popup = window.open(document.location.href)`)
		return err
	})
	require.NoError(t, err)
	require.Equal(t, popup.URL(), server.EMPTY_PAGE)
}
func TestPageExpectNavigation(t *testing.T) {
	t.Skip()
	BeforeEach(t)
	defer AfterEach(t)
}

func TestPageExpectLoadState(t *testing.T) {
	t.Skip()
	BeforeEach(t)
	defer AfterEach(t)
}

func TestPageExpectFileChooser(t *testing.T) {
	t.Skip()
	BeforeEach(t)
	defer AfterEach(t)
}

func TestPageExpectDialog(t *testing.T) {
	t.Skip()
	BeforeEach(t)
	defer AfterEach(t)
}

func TestPageExpectConsoleMessage(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	message, err := page.ExpectConsoleMessage(func() error {
		_, err := page.Evaluate(`console.log(123, "abc")`)
		return err
	})
	require.NoError(t, err)
	require.Equal(t, message.Text(), "123 abc")
}

func TestPageExpectEvent(t *testing.T) {
	t.Skip()
	BeforeEach(t)
	defer AfterEach(t)
}

func TestPageOpener(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	eventPage, err := context.ExpectEvent("page", func() error {
		_, err := page.Goto(server.PREFIX + "/popup/window-open.html")
		return err
	})
	require.NoError(t, err)
	popup := eventPage.(playwright.Page)

	opener, err := popup.Opener()
	require.NoError(t, err)
	require.Equal(t, opener, page)

	opener, err = page.Opener()
	require.NoError(t, err)
	require.Nil(t, opener)
}

func TestPageTitle(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`<title>abc</title>`))
	title, err := page.Title()
	require.NoError(t, err)
	require.Equal(t, "abc", title)
}

func TestPageWaitForSelector(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`<h1>myElement</h1>`))
	element, err := page.WaitForSelector("text=myElement")
	require.NoError(t, err)
	textContent, err := element.TextContent()
	require.NoError(t, err)
	require.Equal(t, "myElement", textContent)

	_, err = page.WaitForSelector("h1")
	require.NoError(t, err)
}

func TestPageDispatchEvent(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/input/button.html")
	require.NoError(t, err)
	require.NoError(t, page.DispatchEvent("button", "click"))
	clicked, err := page.Evaluate("() => result")
	require.NoError(t, err)
	require.Equal(t, "Clicked", clicked)
}

func TestPageReload(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = page.Evaluate("window._foo = 10")
	require.NoError(t, err)
	_, err = page.Reload()
	require.NoError(t, err)
	v, err := page.Evaluate("window._foo")
	require.NoError(t, err)
	require.Nil(t, v)
}

func TestPageGoBackGoForward(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	resp, err := page.GoBack()
	require.NoError(t, err)
	require.Nil(t, resp)

	_, err = page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = page.Goto(server.PREFIX + "/grid.html")
	require.NoError(t, err)

	resp, err = page.GoBack()
	require.NoError(t, err)
	require.True(t, resp.Ok())
	require.Equal(t, resp.URL(), server.EMPTY_PAGE)

	resp, err = page.GoForward()
	require.NoError(t, err)
	require.True(t, resp.Ok())
	require.Contains(t, resp.URL(), "/grid.html")

	resp, err = page.GoForward()
	require.NoError(t, err)
	require.Nil(t, resp)
}

func TestPageAddScriptTag(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	scriptHandle, err := page.AddScriptTag(playwright.PageAddScriptTagOptions{
		URL: playwright.String("injectedfile.js"),
	})
	require.NoError(t, err)
	require.NotNil(t, scriptHandle.AsElement())
	v, err := page.Evaluate("__injected")
	require.NoError(t, err)
	require.Equal(t, 42, v)
}

func TestPageAddScriptTagFile(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	scriptHandle, err := page.AddScriptTag(playwright.PageAddScriptTagOptions{
		Path: playwright.String(Asset("injectedfile.js")),
	})
	require.NoError(t, err)
	require.NotNil(t, scriptHandle.AsElement())
	v, err := page.Evaluate("__injected")
	require.NoError(t, err)
	require.Equal(t, 42, v)
}

func TestPageAddStyleTag(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	_, err = page.AddStyleTag(playwright.PageAddStyleTagOptions{
		URL: playwright.String("injectedstyle.css"),
	})
	require.NoError(t, err)
	v, err := page.Evaluate("window.getComputedStyle(document.querySelector('body')).getPropertyValue('background-color')")
	require.NoError(t, err)
	require.Equal(t, "rgb(255, 0, 0)", v)
}

func TestPageAddStyleTagFile(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	_, err = page.AddStyleTag(playwright.PageAddStyleTagOptions{
		Path: playwright.String(Asset("injectedstyle.css")),
	})
	require.NoError(t, err)
	v, err := page.Evaluate("window.getComputedStyle(document.querySelector('body')).getPropertyValue('background-color')")
	require.NoError(t, err)
	require.Equal(t, "rgb(255, 0, 0)", v)
}

func TestPageWaitForLoadState(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/one-style.html")
	require.NoError(t, err)
	page.WaitForLoadState()
	page.WaitForLoadState("networkidle")
}

func TestPlaywrightDevices(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.Greater(t, len(pw.Devices), 10)
	for name, device := range pw.Devices {
		require.NotEmpty(t, name)
		require.NotEmpty(t, device.UserAgent)
		require.NotEmpty(t, device.Viewport)
		require.Greater(t, device.DeviceScaleFactor, float64(0))
		require.NotEmpty(t, device.DefaultBrowserType)
	}
}

func TestPageAddInitScript(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.AddInitScript(playwright.PageAddInitScriptOptions{
		Script: playwright.String(`window['injected'] = 123;`),
	}))
	_, err := page.Goto(server.PREFIX + "/tamperable.html")
	require.NoError(t, err)
	result, err := page.Evaluate(`() => window['result']`)
	require.NoError(t, err)
	require.Equal(t, 123, result)
}

func TestPageExpectSelectorTimeout(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	timeoutError := errors.Unwrap(page.Click("foobar", playwright.PageClickOptions{
		Timeout: playwright.Float(500),
	})).(*playwright.TimeoutError)
	require.Contains(t, timeoutError.Message, "Timeout 500ms exceeded.")
}

func TestPageType(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent("<input type='text' />"))

	require.NoError(t, page.Type("input", "hello"))
	value, err := page.EvalOnSelector("input", "el => el.value")
	require.NoError(t, err)
	require.Equal(t, "hello", value)
}

func TestPagePress(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent("<input type='text' />"))

	require.NoError(t, page.Press("input", "h"))
	value, err := page.EvalOnSelector("input", "el => el.value")
	require.NoError(t, err)
	require.Equal(t, "h", value)
}

func TestPageCheck(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent("<input id='checkbox' type='checkbox'></input>"))

	require.NoError(t, page.Check("input"))
	value, err := page.EvalOnSelector("input", "el => el.checked")
	require.NoError(t, err)
	require.Equal(t, true, value)
}

func TestPageUncheck(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent("<input id='checkbox' type='checkbox' checked></input>"))

	require.NoError(t, page.Uncheck("input"))
	value, err := page.EvalOnSelector("input", "el => el.checked")
	require.NoError(t, err)
	require.Equal(t, false, value)
}

func TestPageWaitForTimeout(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	before := time.Now()
	page.WaitForTimeout(1000)
	after := time.Now()
	duration := after.Sub(before)
	require.True(t, duration > time.Second)
	require.True(t, duration < 2*time.Second)
}

func TestPageWaitForFunction(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Evaluate(`() => setTimeout(() => window.FOO = true, 500)`)
	require.NoError(t, err)
	_, err = page.WaitForFunction(`window.FOO === true`, nil)
	require.NoError(t, err)
}

func TestPageDblclick(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<button ondblclick="window.clicked=true"/>`))
	require.NoError(t, page.Dblclick("button"))
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestPageFocus(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<button onfocus="window.clicked=true"/>`))
	require.NoError(t, page.Focus("button"))
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestPageTextContent(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/dom.html")
	require.NoError(t, err)
	content, err := page.TextContent("#inner")
	require.NoError(t, err)
	require.Equal(t, "Text,\nmore text", content)
}

func TestPageAddInitScriptWithPath(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.AddInitScript(playwright.PageAddInitScriptOptions{
		Path: playwright.String(Asset("injectedfile.js")),
	}))
	_, err := page.Goto(server.PREFIX + "/tamperable.html")
	require.NoError(t, err)
	result, err := page.Evaluate(`() => window['result']`)
	require.NoError(t, err)
	require.Equal(t, 123, result)
}

func TestPageSupportNetworkEvents(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	eventsChan := make(chan string, 6)
	page.On("request", func(request playwright.Request) {
		eventsChan <- fmt.Sprintf("%s %s", request.Method(), request.URL())
	})
	page.On("response", func(response playwright.Response) {
		eventsChan <- fmt.Sprintf("%d %s", response.Status(), response.URL())
	})
	page.On("requestfinished", func(request playwright.Request) {
		eventsChan <- fmt.Sprintf("DONE %s", request.URL())
	})
	page.On("requestfailed", func(request playwright.Request) {
		eventsChan <- fmt.Sprintf("FAIL %s", request.URL())
	})
	server.SetRedirect("/foo.html", "/empty.html")
	FOO_URL := server.PREFIX + "/foo.html"
	response, err := page.Goto(FOO_URL)
	require.NoError(t, err)
	eventsSlice := ChanToSlice(eventsChan, 6)
	require.Equal(t, []string{
		fmt.Sprintf("GET %s", FOO_URL),
		fmt.Sprintf("302 %s", FOO_URL),
		fmt.Sprintf("DONE %s", FOO_URL),
		fmt.Sprintf("GET %s", server.EMPTY_PAGE),
		fmt.Sprintf("200 %s", server.EMPTY_PAGE),
		fmt.Sprintf("DONE %s", server.EMPTY_PAGE),
	}, eventsSlice)
	redirectedFrom := response.Request().RedirectedFrom()
	require.Contains(t, redirectedFrom.URL(), "foo.html")
	require.Nil(t, redirectedFrom.RedirectedFrom())
	require.Equal(t, redirectedFrom.RedirectedTo(), response.Request())
}

func TestPageSetViewport(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	utils.VerifyViewport(t, page, 1280, 720)
	require.NoError(t, page.SetViewportSize(123, 456))
	utils.VerifyViewport(t, page, 123, 456)
}

func TestPageEmulateMedia(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	utils.AssertEval(t, page, "matchMedia('screen').matches", true)
	utils.AssertEval(t, page, "matchMedia('print').matches", false)
	require.NoError(t, page.EmulateMedia(playwright.PageEmulateMediaOptions{
		Media: playwright.MediaPrint,
	}))
	utils.AssertEval(t, page, "matchMedia('screen').matches", false)
	utils.AssertEval(t, page, "matchMedia('print').matches", true)
	require.NoError(t, page.EmulateMedia())
	utils.AssertEval(t, page, "matchMedia('screen').matches", false)
	utils.AssertEval(t, page, "matchMedia('print').matches", true)
	require.NoError(t, page.EmulateMedia(playwright.PageEmulateMediaOptions{
		Media: playwright.MediaNull,
	}))
	utils.AssertEval(t, page, "matchMedia('screen').matches", true)
	utils.AssertEval(t, page, "matchMedia('print').matches", false)
}

func TestPageBringToFront(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	page1, err := browser.NewPage()
	require.NoError(t, err)
	require.NoError(t, page1.SetContent("Page1"))
	page2, err := browser.NewPage()
	require.NoError(t, err)
	require.NoError(t, page2.SetContent("Page2"))

	require.NoError(t, page1.BringToFront())
	utils.AssertEval(t, page1, "document.visibilityState", "visible")
	utils.AssertEval(t, page2, "document.visibilityState", "visible")

	require.NoError(t, page2.BringToFront())
	utils.AssertEval(t, page1, "document.visibilityState", "visible")
	utils.AssertEval(t, page2, "document.visibilityState", "visible")
	require.NoError(t, page1.Close())
	require.NoError(t, page2.Close())
}

func TestPageFrame(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(fmt.Sprintf("<iframe name=target src=%s></iframe>", server.EMPTY_PAGE))
	require.NoError(t, err)

	var name = "target"
	frame1 := page.Frame(playwright.PageFrameOptions{Name: &name})
	require.Equal(t, name, frame1.Name())
	require.Equal(t, server.EMPTY_PAGE, frame1.URL())

	frame2 := page.Frame(playwright.PageFrameOptions{URL: server.EMPTY_PAGE})
	require.Equal(t, name, frame2.Name())
	require.Equal(t, server.EMPTY_PAGE, frame2.URL())

	var badName = "test"
	frame3 := page.Frame(playwright.PageFrameOptions{Name: &badName, URL: server.EMPTY_PAGE})
	require.Equal(t, name, frame3.Name())
	require.Equal(t, server.EMPTY_PAGE, frame3.URL())

	require.Nil(t, page.Frame(playwright.PageFrameOptions{Name: &badName, URL: "https://example.com"}))
	require.Nil(t, page.Frame(playwright.PageFrameOptions{Name: &badName}))
}

func TestPageTap(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent("<input id='checkbox' type='checkbox'></input>"))
	value, err := page.EvalOnSelector("input", "el => el.checked")
	require.NoError(t, err)
	require.Equal(t, false, value)

	require.NoError(t, page.Tap("input"))
	value, err = page.EvalOnSelector("input", "el => el.checked")
	require.NoError(t, err)
	require.Equal(t, true, value)
}

func TestPagePageError(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	url := server.PREFIX + "/error.html"
	errInterface, err := page.ExpectEvent("pageerror", func() error {
		_, err := page.Goto(url)
		return err
	})
	require.NoError(t, err)
	pageError := errInterface.(*playwright.Error)
	require.Equal(t, "Fancy error!", pageError.Message)
	require.Equal(t, "Error", pageError.Name)

	if browserName == "chromium" {
		require.Equal(t, `Error: Fancy error!
    at c (myscript.js:14:11)
    at b (myscript.js:10:5)
    at a (myscript.js:6:5)
    at myscript.js:3:1`, pageError.Stack)
	}
	if browserName == "firefox" {
		require.Equal(t, `Error: Fancy error!
    at c (myscript.js:14:11)
    at b (myscript.js:10:5)
    at a (myscript.js:6:5)
    at  (myscript.js:3:1)`, pageError.Stack)
	}
	if browserName == "webkit" {
		require.Equal(t, fmt.Sprintf(`Error: Fancy error!
    at c (%[1]s:14:36)
    at b (%[1]s:10:6)
    at a (%[1]s:6:6)
    at global code (%[1]s:3:2)`, url), pageError.Stack)
	}
}

func TestPageSelectOption(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent("<select id='lang'><option value='go'>go</option><option value='python'>python</option></select>"))
	selected, err := page.SelectOption("#lang", playwright.SelectOptionValues{
		Values: playwright.StringSlice("python"),
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(selected))
	require.Equal(t, "python", selected[0])
}

func TestPageUnrouteShouldWork(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	intercepted := []int{}
	handler1 := func(route playwright.Route, request playwright.Request) {
		intercepted = append(intercepted, 1)
		require.NoError(t, route.Continue())
	}
	require.NoError(t, page.Route("**/empty.html", handler1))
	require.NoError(t, page.Route("**/empty.html", func(route playwright.Route, request playwright.Request) {
		intercepted = append(intercepted, 2)
		require.NoError(t, route.Continue())
	}))
	require.NoError(t, page.Route("**/empty.html", func(route playwright.Route, request playwright.Request) {
		intercepted = append(intercepted, 3)
		require.NoError(t, route.Continue())
	}))
	require.NoError(t, page.Route("**/*", func(route playwright.Route, request playwright.Request) {
		intercepted = append(intercepted, 4)
		require.NoError(t, route.Continue())
	}))

	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.Equal(t, []int{1}, intercepted)

	intercepted = []int{}
	require.NoError(t, page.Unroute("**/empty.html", handler1))
	_, err = page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.Equal(t, []int{2}, intercepted)

	intercepted = []int{}
	require.NoError(t, page.Unroute("**/empty.html"))
	_, err = page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.Equal(t, []int{4}, intercepted)
}

func TestPageDragAndDrop(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/drag-n-drop.html")
	require.NoError(t, err)
	require.NoError(t, page.DragAndDrop("#source", "#target"))
	value, err := page.EvalOnSelector("#target", "target => target.contains(document.querySelector('#source'))")
	require.NoError(t, err)
	require.Equal(t, true, value)
}
func TestPageInputValue(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`
		<input></input>
	`))
	require.NoError(t, page.Fill("input", "hello"))
	value, err := page.InputValue("input")
	require.NoError(t, err)
	require.Equal(t, "hello", value)
	require.NoError(t, page.Fill("input", ""))
	value, err = page.InputValue("input")
	require.NoError(t, err)
	require.Equal(t, "", value)
}

func TestPageWaitForFunction2(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = page.AddScriptTag(
		playwright.PageAddScriptTagOptions{
			Path: playwright.String(Asset(filepath.Join("es6", "es6pathimport.js"))),
			Type: playwright.String("module"),
		},
	)
	require.NoError(t, err)
	_, err = page.WaitForFunction("window.__es6injected", nil)
	require.NoError(t, err)
	value, err := page.Evaluate("window.__es6injected")
	require.NoError(t, err)
	require.Equal(t, 42, value)
}

func TestPageShouldSetBodysizeAndHeadersize(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	request, err := page.ExpectRequest("**/*", func() error {
		_, err = page.Evaluate("() => fetch('./get', { method: 'POST', body: '12345'}).then(r => r.text())")
		require.NoError(t, err)
		return nil
	},
	)
	require.NoError(t, err)
	sizes, err := request.Sizes()
	require.NoError(t, err)
	require.Equal(t, 5, sizes.RequestBodySize)
	require.GreaterOrEqual(t, sizes.RequestHeadersSize, 300)
}

func TestPageTestShouldSetBodysizeTo0(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	request, err := page.ExpectRequest("**/*", func() error {
		_, err = page.Evaluate("() => fetch('./get').then(r => r.text())")
		require.NoError(t, err)
		return nil
	},
	)
	require.NoError(t, err)
	sizes, err := request.Sizes()
	require.NoError(t, err)
	require.Equal(t, 0, sizes.RequestBodySize)
	require.GreaterOrEqual(t, sizes.RequestHeadersSize, 200)
}

func TestPageSetChecked(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`<input id='checkbox' type='checkbox'></input>`))
	require.NoError(t, page.SetChecked("input", true))
	isChecked, err := page.Evaluate("checkbox.checked")
	require.NoError(t, err)
	require.True(t, isChecked.(bool))
	require.NoError(t, page.SetChecked("input", false))
	isChecked, err = page.Evaluate("checkbox.checked")
	require.NoError(t, err)
	require.False(t, isChecked.(bool))
}
