package playwright_test

import (
	"path/filepath"
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestBrowserContextOutputTrace(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	context, err := browser.NewContext()
	require.NoError(t, err)
	defer context.Close()
	err = context.Tracing().Start(playwright.TracingStartOptions{
		Screenshots: playwright.Bool(true),
		Snapshots:   playwright.Bool(true),
	})
	require.NoError(t, err)
	page, err := context.NewPage()
	require.NoError(t, err)
	defer page.Close()
	_, err = page.Goto(server.PREFIX + "/grid.html")
	require.NoError(t, err)
	dir := t.TempDir()
	err = context.Tracing().Stop(playwright.TracingStopOptions{
		Path: playwright.String(filepath.Join(dir, "trace.zip")),
	})
	require.NoError(t, err)
	require.FileExists(t, filepath.Join(dir, "trace.zip"))
}
