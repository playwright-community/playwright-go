package playwright

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/h2non/filetype"
	"github.com/stretchr/testify/require"
)

func TestPageURL(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	require.Equal(t, "about:blank", helper.Page.URL())
	require.NoError(t, helper.Page.Goto("https://example.com"))
	require.Equal(t, "https://example.com/", helper.Page.URL())
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
		return helper.Page.Goto(helper.server.PREFIX + "/worker/worker.html")
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(helper.Page.Workers()))
	require.Equal(t, worker, helper.Page.Workers()[0])
	worker = helper.Page.Workers()[0]
	require.Contains(t, worker.URL(), "worker.js")
	res, err := worker.Evaluate(`() => self["workerFunction"]()`)
	require.NoError(t, err)
	require.Equal(t, "worker function result", res)
	require.NoError(t, helper.Page.Goto(helper.server.EMPTY_PAGE))
	require.Equal(t, 0, len(helper.Page.Workers()))
}

func TestPageExpectRequest(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	request, err := helper.Page.ExpectRequest("**/*", func() error {
		return helper.Page.Goto(helper.server.EMPTY_PAGE)
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
		return helper.Page.Goto(helper.server.EMPTY_PAGE)
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
	require.NoError(t, helper.Page.Goto(helper.server.EMPTY_PAGE))
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
