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
	err = context.Tracing().Stop(filepath.Join(dir, "trace.zip"))
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
	err = context.Tracing().StopChunk(filepath.Join(dir, "trace1.zip"))
	require.NoError(t, err)
	require.FileExists(t, filepath.Join(dir, "trace1.zip"))

	err = context.Tracing().StartChunk(playwright.TracingStartChunkOptions{
		Title: playwright.String("foo"),
	})
	require.NoError(t, err)
	err = button.Click()
	require.NoError(t, err)
	err = context.Tracing().StopChunk(filepath.Join(dir, "trace2.zip"))
	require.NoError(t, err)
	require.FileExists(t, filepath.Join(dir, "trace2.zip"))
}

func TestBrowserContextTracingOutputMultipleChunks(t *testing.T) {
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
	require.NoError(t, page.Locator("button").Click())
	dir := t.TempDir()
	require.NoError(t, context.Tracing().StopChunk(filepath.Join(dir, "trace.zip")))
	require.FileExists(t, filepath.Join(dir, "trace.zip"))
}

func TestBrowserContextTracingRemoteConnect(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	remoteServer, err := newRemoteServer()
	require.NoError(t, err)
	defer remoteServer.Close()
	browser1, err := browserType.Connect(remoteServer.url)
	require.NoError(t, err)
	require.NotNil(t, browser1)
	defer browser1.Close()
	context1, err := browser1.NewContext()
	require.NoError(t, err)
	require.NoError(t, context1.Tracing().Start(playwright.TracingStartOptions{
		Screenshots: playwright.Bool(true),
		Snapshots:   playwright.Bool(true),
	}))
	page, err := context1.NewPage()
	require.NoError(t, err)
	defer page.Close()
	_, err = page.Goto(server.PREFIX + "/frames/frame.html")
	require.NoError(t, err)
	require.NoError(t, context1.Tracing().StartChunk())
	require.NoError(t, page.SetContent("<button>Click</button>"))
	require.NoError(t, page.Locator("button").Click())
	dir := t.TempDir()
	require.NoError(t, context1.Tracing().StopChunk(filepath.Join(dir, "trace.zip")))
	require.FileExists(t, filepath.Join(dir, "trace.zip"))
}
