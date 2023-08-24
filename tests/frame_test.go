package playwright_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestFrameWaitForNavigationShouldWork(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	//nolint:staticcheck
	response, err := page.ExpectNavigation(func() error {
		_, err := page.Evaluate("url => window.location.href = url", server.PREFIX+"/grid.html")
		return err
	})
	require.NoError(t, err)
	require.True(t, response.Ok())
	require.Contains(t, response.URL(), "grid.html")
}

func TestFrameWaitForNavigationShouldRespectTimeout(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	timeout := 500.0
	//nolint:staticcheck
	_, err := page.ExpectNavigation(func() error {
		_, err := page.Evaluate("url => window.location.href = url", server.EMPTY_PAGE)
		return err
	}, playwright.PageExpectNavigationOptions{
		URL:     "**/frame.html",
		Timeout: playwright.Float(timeout),
	})
	require.ErrorContains(t, err, fmt.Sprintf(`Timeout %.2fms exceeded.`, timeout))
}

func TestFrameWaitForURLShouldWork(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	require.NoError(t, page.SetContent(`<a href="grid.html">foobar</a>`))
	go func() {
		time.Sleep(2 * time.Second)
		require.NoError(t, page.Locator("a").Click())
	}()

	err = page.MainFrame().WaitForURL(server.PREFIX + "/grid.html")
	require.NoError(t, err)
	require.Equal(t, server.PREFIX+"/grid.html", page.URL())
}

func TestFrameWaitForNavigationAnchorLinks(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<a href="#foobar">foobar</a>`))
	//nolint:staticcheck
	response, err := page.ExpectNavigation(func() error {
		return page.Locator("a").Click()
	})
	require.NoError(t, err)
	require.Nil(t, response)
	require.Equal(t, server.EMPTY_PAGE+"#foobar", page.URL())
}

func TestFrameInnerHTML(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/dom.html")
	require.NoError(t, err)
	//nolint:staticcheck
	handle, err := page.QuerySelector("#outer")
	require.NoError(t, err)
	innerHTML, err := handle.InnerHTML()
	require.NoError(t, err)
	require.Equal(t, `<div id="inner">Text,
more text</div>`, innerHTML)
	//nolint:staticcheck
	innerHTML, err = page.InnerHTML("#outer")
	require.NoError(t, err)
	require.Equal(t, `<div id="inner">Text,
more text</div>`, innerHTML)
}

func TestFrameSetInputFiles(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent("<input type=file>"))
	//nolint:staticcheck
	require.NoError(t, page.SetInputFiles("input", []playwright.InputFile{
		{
			Name:     "file-to-upload.txt",
			MimeType: "text/plain",
			Buffer:   []byte("123"),
		},
	}))
	fileName, err := page.Locator("input").Evaluate("e => e.files[0].name", nil)
	require.NoError(t, err)
	require.Equal(t, "file-to-upload.txt", fileName)
}
