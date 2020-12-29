package playwright_test

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestFrameWaitForNavigationShouldWork(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	response, err := helper.Page.ExpectNavigation(func() error {
		_, err := helper.Page.Evaluate("url => window.location.href = url", helper.server.PREFIX+"/grid.html")
		return err
	})
	require.NoError(t, err)
	require.True(t, response.Ok())
	require.Contains(t, response.URL(), "grid.html")
}

func TestFrameWaitForNavigationAnchorLinks(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.SetContent(`<a href="#foobar">foobar</a>`))
	response, err := helper.Page.ExpectNavigation(func() error {
		return helper.Page.Click("a")
	})
	require.NoError(t, err)
	require.Nil(t, response)
	require.Equal(t, helper.server.EMPTY_PAGE+"#foobar", helper.Page.URL())
}

func TestFrameInnerHTML(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.PREFIX + "/dom.html")
	require.NoError(t, err)
	handle, err := helper.Page.QuerySelector("#outer")
	require.NoError(t, err)
	innerHTML, err := handle.InnerHTML()
	require.NoError(t, err)
	require.Equal(t, `<div id="inner">Text,
more text</div>`, innerHTML)
	innerHTML, err = helper.Page.InnerHTML("#outer")
	require.NoError(t, err)
	require.Equal(t, `<div id="inner">Text,
more text</div>`, innerHTML)
}

func TestFrameSetInputFiles(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.SetContent("<input type=file>"))

	require.NoError(t, helper.Page.SetInputFiles("input", []playwright.InputFile{
		{
			Name:     "file-to-upload.txt",
			MimeType: "text/plain",
			Buffer:   []byte("123"),
		},
	}))
	fileName, err := helper.Page.EvalOnSelector("input", "e => e.files[0].name")
	require.NoError(t, err)
	require.Equal(t, "file-to-upload.txt", fileName)
}
