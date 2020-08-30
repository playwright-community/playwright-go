package playwright

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/h2non/filetype"
	"github.com/stretchr/testify/require"
)

func TestPageURL(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	require.Equal(t, "about:blank", helper.Page.URL())
	_, err := helper.Page.Goto("https://example.com")
	require.NoError(t, err)
	require.Equal(t, "https://example.com/", helper.Page.URL())
	require.Equal(t, helper.Context, helper.Page.Context())
	require.Equal(t, 1, len(helper.Page.Frames()))
}

func TestPageSetContent(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	require.NoError(t, helper.Page.SetContent("<h1>foo</h1>",
		PageSetContentOptions{
			WaitUntil: String("networkidle"),
		}))
	content, err := helper.Page.Content()
	require.NoError(t, err)
	require.Equal(t, content, "<html><head></head><body><h1>foo</h1></body></html>")
}

func TestPageScreenshot(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()

	require.NoError(t, helper.Page.SetContent("<h1>foobar</h1>"))
	tmpfile, err := ioutil.TempDir("", "screenshot")
	require.NoError(t, err)
	screenshotPath := filepath.Join(tmpfile, "image.png")
	screenshot, err := helper.Page.Screenshot()
	require.NoError(t, err)
	require.True(t, filetype.IsImage(screenshot))
	require.Greater(t, len(screenshot), 50)

	screenshot, err = helper.Page.Screenshot(PageScreenshotOptions{
		Path: String(screenshotPath),
	})
	require.NoError(t, err)
	require.True(t, filetype.IsImage(screenshot))
	require.Greater(t, len(screenshot), 50)

	_, err = os.Stat(screenshotPath)
	require.NoError(t, err)
}

func TestPagePDF(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	if !helper.IsChromium {
		t.Skip("Skipping")
	}
	require.NoError(t, helper.Page.SetContent("<h1>foobar</h1>"))
	tmpfile, err := ioutil.TempDir("", "pdf")
	require.NoError(t, err)
	screenshotPath := filepath.Join(tmpfile, "image.png")
	screenshot, err := helper.Page.PDF()
	require.NoError(t, err)
	require.Equal(t, "application/pdf", http.DetectContentType(screenshot))
	require.Greater(t, len(screenshot), 50)

	screenshot, err = helper.Page.PDF(PagePdfOptions{
		Path: String(screenshotPath),
	})
	require.NoError(t, err)
	require.Equal(t, "application/pdf", http.DetectContentType(screenshot))
	require.Greater(t, len(screenshot), 50)

	_, err = os.Stat(screenshotPath)
	require.NoError(t, err)
}

func TestPageQuerySelector(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	require.NoError(t, helper.Page.SetContent(`<div id="one">
		<span id="two">
			<div id="three">
				<span id="four">
					foobar
				</span>
			</div>
		</span>
	</div>`))
	one, err := helper.Page.QuerySelector("div#one")
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
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	require.NoError(t, helper.Page.SetContent(`
	<div class="foo">0</div>
	<div class="foo">1</div>
	<div class="foo">2</div>
	`))
	elements, err := helper.Page.QuerySelectorAll("div.foo")
	require.NoError(t, err)
	require.Equal(t, 3, len(elements))
	for i := 0; i < 3; i++ {
		textContent, err := elements[i].TextContent()
		require.NoError(t, err)
		require.Equal(t, strconv.Itoa(i), textContent)
	}
}

func TestPageEvaluate(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	val, err := helper.Page.Evaluate(`() => 123`)
	require.NoError(t, err)
	require.Equal(t, val, 123)
}

func TestPageEvaluateOnSelectorAll(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	err := helper.Page.SetContent(`
		<div class="foo">1</div>
		<div class="foo">2</div>
		<div class="foo">3</div>
	`)
	require.NoError(t, err)
	val, err := helper.Page.EvaluateOnSelectorAll(".foo", `(elements) => elements.map(el => el.textContent)`)
	require.NoError(t, err)
	require.Equal(t, val, []interface{}([]interface{}{"1", "2", "3"}))
}

func TestPageEvaluateOnSelector(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	err := helper.Page.SetContent(`
		<div class="foo">bar</div>
	`)
	require.NoError(t, err)
	val, err := helper.Page.EvaluateOnSelector(".foo", `(element) => element.textContent`)
	require.NoError(t, err)
	require.Equal(t, val, "bar")
}

func TestPageExpectWorker(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	worker, err := helper.Page.ExpectWorker(func() error {
		_, err := helper.Page.Goto(helper.server.PREFIX + "/worker/worker.html")
		return err
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(helper.Page.Workers()))
	require.Equal(t, worker, helper.Page.Workers()[0])
	worker = helper.Page.Workers()[0]
	require.Contains(t, worker.URL(), "worker.js")
	res, err := worker.Evaluate(`() => self["workerFunction"]()`)
	require.NoError(t, err)
	require.Equal(t, "worker function result", res)
	_, err = helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.Equal(t, 0, len(helper.Page.Workers()))
}

func TestPageExpectRequest(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	request, err := helper.Page.ExpectRequest("**/*", func() error {
		_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
		return err
	})
	require.NoError(t, err)
	require.Equal(t, helper.server.EMPTY_PAGE, request.URL())
	require.Equal(t, "document", request.ResourceType())
	require.Equal(t, "GET", request.Method())
}

func TestPageExpectResponse(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	response, err := helper.Page.ExpectResponse("**/*", func() error {
		_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
		return err
	})
	require.NoError(t, err)
	require.Equal(t, helper.server.EMPTY_PAGE, response.URL())
	require.True(t, response.Ok())
	require.Equal(t, 200, response.Status())
	require.Equal(t, "OK", response.StatusText())
}

func TestPageExpectPopup(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	popup, err := helper.Page.ExpectPopup(func() error {
		_, err := helper.Page.Evaluate(`window._popup = window.open(document.location.href)`)
		return err
	})
	require.NoError(t, err)
	require.Equal(t, popup.URL(), helper.server.EMPTY_PAGE)
}

func TestPageExpectNavigation(t *testing.T) {
	t.Skip()
	helper := NewTestHelper(t)
	defer helper.AfterEach()
}

func TestPageExpectLoadState(t *testing.T) {
	t.Skip()
	helper := NewTestHelper(t)
	defer helper.AfterEach()
}

func TestPageExpectFileChooser(t *testing.T) {
	t.Skip()
	helper := NewTestHelper(t)
	defer helper.AfterEach()
}

func TestPageExpectDialog(t *testing.T) {
	t.Skip()
	helper := NewTestHelper(t)
	defer helper.AfterEach()
}

func TestPageExpectConsoleMessage(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	message, err := helper.Page.ExpectConsoleMessage(func() error {
		_, err := helper.Page.Evaluate(`console.log(123, "abc")`)
		return err
	})
	require.NoError(t, err)
	require.Equal(t, message.Text(), "123 abc")
}

func TestPageExpectEvent(t *testing.T) {
	t.Skip()
	helper := NewTestHelper(t)
	defer helper.AfterEach()
}

func TestPageOpener(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	page, err := helper.Context.ExpectEvent("page", func() error {
		_, err := helper.Page.Goto(helper.server.PREFIX + "/popup/window-open.html")
		return err
	})
	require.NoError(t, err)
	popup := page.(*Page)

	opener, err := popup.Opener()
	require.NoError(t, err)
	require.Equal(t, opener, helper.Page)

	opener, err = helper.Page.Opener()
	require.NoError(t, err)
	require.Nil(t, opener)
}

func TestPageTitle(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	require.NoError(t, helper.Page.SetContent(`<title>abc</title>`))
	title, err := helper.Page.Title()
	require.NoError(t, err)
	require.Equal(t, "abc", title)
}

func TestPageWaitForSelector(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	require.NoError(t, helper.Page.SetContent(`<h1>myElement</h1>`))
	element, err := helper.Page.WaitForSelector("text=myElement")
	require.NoError(t, err)
	textContent, err := element.TextContent()
	require.NoError(t, err)
	require.Equal(t, "myElement", textContent)

	_, err = helper.Page.WaitForSelector("h1")
	require.NoError(t, err)
}

func TestPageDispatchEvent(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.PREFIX + "/input/button.html")
	require.NoError(t, err)
	require.NoError(t, helper.Page.DispatchEvent("button", "click"))
	clicked, err := helper.Page.Evaluate("() => result")
	require.NoError(t, err)
	require.Equal(t, "Clicked", clicked)
}

func TestPageReload(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = helper.Page.Evaluate("window._foo = 10")
	require.NoError(t, err)
	_, err = helper.Page.Reload()
	require.NoError(t, err)
	v, err := helper.Page.Evaluate("window._foo")
	require.NoError(t, err)
	require.Nil(t, v)
}

func TestPageGoBackGoForward(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()

	resp, err := helper.Page.GoBack()
	require.NoError(t, err)
	require.Nil(t, resp)

	_, err = helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = helper.Page.Goto(helper.server.PREFIX + "/grid.html")
	require.NoError(t, err)

	resp, err = helper.Page.GoBack()
	require.NoError(t, err)
	require.True(t, resp.Ok())
	require.Equal(t, resp.URL(), helper.server.EMPTY_PAGE)

	resp, err = helper.Page.GoForward()
	require.NoError(t, err)
	require.True(t, resp.Ok())
	require.Contains(t, resp.URL(), "/grid.html")

	resp, err = helper.Page.GoForward()
	require.NoError(t, err)
	require.Nil(t, resp)
}

func TestPageAddScriptTag(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)

	_, err = helper.Page.AddScriptTag(PageAddScriptTagOptions{
		Url: String("injectedfile.js"),
	})
	require.NoError(t, err)
	v, err := helper.Page.Evaluate("__injected")
	require.NoError(t, err)
	require.Equal(t, 42, v)
}

func TestPageAddStyleTag(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)

	_, err = helper.Page.AddStyleTag(PageAddStyleTagOptions{
		Url: String("injectedstyle.css"),
	})
	require.NoError(t, err)
	v, err := helper.Page.Evaluate("window.getComputedStyle(document.querySelector('body')).getPropertyValue('background-color')")
	require.NoError(t, err)
	require.Equal(t, "rgb(255, 0, 0)", v)
}

func TestPageWaitForLoadState(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.PREFIX + "/one-style.html")
	require.NoError(t, err)
	helper.Page.WaitForLoadState()
	helper.Page.WaitForLoadState("networkidle")
}
