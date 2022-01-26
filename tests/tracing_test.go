package playwright_test

import (
	"path/filepath"
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestBrowserContextOutputTrace(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	context, err := browser.NewContext()
	require.NoError(t, err)
	defer context.Close()
	require.NoError(t, context.Tracing().Start(playwright.TracingStartOptions{
		Screenshots: playwright.Bool(true),
		Snapshots:   playwright.Bool(true),
	}))
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

func TestBrowserContextMultipleChunks(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	context, err := browser.NewContext()
	require.NoError(t, err)
	defer context.Close()
	require.NoError(t, context.Tracing().Start(playwright.TracingStartOptions{
		Screenshots: playwright.Bool(true),
		Snapshots:   playwright.Bool(true),
	}))
	page, err := context.NewPage()
	require.NoError(t, err)
	defer page.Close()
	_, err = page.Goto(server.PREFIX + "/frames/frame.html")
	require.NoError(t, err)
	require.NoError(t, context.Tracing().StartChunk())
	require.NoError(t, page.SetContent("<button>Click</button>"))
	require.NoError(t, page.Click("button"))
	dir := t.TempDir()
	require.NoError(t, context.Tracing().StopChunk(playwright.TracingStopChunkOptions{
		Path: playwright.String(filepath.Join(dir, "trace.zip")),
	}))
	require.FileExists(t, filepath.Join(dir, "trace.zip"))
}
