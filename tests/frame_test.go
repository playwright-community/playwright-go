package playwright_test

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestFrameWaitForNavigationShouldWork(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	response, err := page.ExpectNavigation(func() error {
		_, err := page.Evaluate("url => window.location.href = url", server.PREFIX+"/grid.html")
		return err
	})
	require.NoError(t, err)
	require.True(t, response.Ok())
	require.Contains(t, response.URL(), "grid.html")
}

func TestFrameWaitForNavigationAnchorLinks(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<a href="#foobar">foobar</a>`))
	response, err := page.ExpectNavigation(func() error {
		return page.Click("a")
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
	handle, err := page.QuerySelector("#outer")
	require.NoError(t, err)
	innerHTML, err := handle.InnerHTML()
	require.NoError(t, err)
	require.Equal(t, `<div id="inner">Text,
more text</div>`, innerHTML)
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

	require.NoError(t, page.SetInputFiles("input", []playwright.InputFile{
		{
			Name:     "file-to-upload.txt",
			MimeType: "text/plain",
			Buffer:   []byte("123"),
		},
	}))
	fileName, err := page.EvalOnSelector("input", "e => e.files[0].name")
	require.NoError(t, err)
	require.Equal(t, "file-to-upload.txt", fileName)
}
