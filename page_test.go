package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPageURL(t *testing.T) {
	pw, err := Run()
	require.NoError(t, err)
	browser, err := pw.Chromium.Launch()
	require.NoError(t, err)
	context, err := browser.NewContext()
	require.NoError(t, err)
	page, err := context.NewPage()
	require.NoError(t, err)
	require.Equal(t, "about:blank", page.URL())
	require.NoError(t, page.Goto("https://example.com"))
	require.Equal(t, "https://example.com/", page.URL())

}

func TestPageSetContent(t *testing.T) {
	pw, err := Run()
	require.NoError(t, err)
	browser, err := pw.Chromium.Launch()
	require.NoError(t, err)
	context, err := browser.NewContext()
	require.NoError(t, err)
	page, err := context.NewPage()
	require.NoError(t, err)
	require.NoError(t, page.SetContent("<h1>foo</h1>"))
	content, err := page.Content()
	require.NoError(t, err)
	require.Equal(t, content, "<html><head></head><body><h1>foo</h1></body></html>")
}
