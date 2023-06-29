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

func TestTracingStartStop(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, context.Tracing().Start())
	require.NoError(t, context.Tracing().Stop())
}

func TestBrowserContextShouldNoErrorWhenStoppingWithoutStart(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	context, err := browser.NewContext()
	require.NoError(t, err)
	defer context.Close()
	require.NoError(t, context.Tracing().Stop())
}

func TestBrowserContextOutputTraceChunk(t *testing.T) {
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

	button := page.Locator(".box").First()

	err = context.Tracing().StartChunk(playwright.TracingStartChunkOptions{
		Title: playwright.String("foo"),
	})
	require.NoError(t, err)
	err = button.Click()
	require.NoError(t, err)
	err = context.Tracing().StopChunk(playwright.TracingStopChunkOptions{
		Path: playwright.String(filepath.Join(dir, "trace1.zip")),
	})
	require.NoError(t, err)
	require.FileExists(t, filepath.Join(dir, "trace1.zip"))

	err = context.Tracing().StartChunk(playwright.TracingStartChunkOptions{
		Title: playwright.String("foo"),
	})
	require.NoError(t, err)
	err = button.Click()
	require.NoError(t, err)
	err = context.Tracing().StopChunk(playwright.TracingStopChunkOptions{
		Path: playwright.String(filepath.Join(dir, "trace2.zip")),
	})
	require.NoError(t, err)
	require.FileExists(t, filepath.Join(dir, "trace2.zip"))
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
