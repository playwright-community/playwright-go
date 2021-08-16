package playwright_test

import (
	"path/filepath"
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestExampleRun(t *testing.T) {
	pw, err := playwright.Run()
	require.NoError(t, err)
	browser, err := pw.Chromium.Launch()
	require.NoError(t, err)
	context, err := browser.NewContext()
	require.NoError(t, err)
	page, err := context.NewPage()
	require.NoError(t, err)
	_, err = page.Goto("http://whatsmyuseragent.org/")
	require.NoError(t, err)
	path := filepath.Join(t.TempDir(), "foo.png")
	_, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: &path,
	})
	require.NoError(t, err)
	require.FileExists(t, path)
	require.NoError(t, browser.Close())
	require.NoError(t, pw.Stop())
}
