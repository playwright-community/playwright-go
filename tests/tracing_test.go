package playwright_test

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"io"
	"path/filepath"
	"slices"
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestBrowserContextOutputTrace(t *testing.T) {
	BeforeEach(t)

	require.NoError(t, context.Tracing().Start(playwright.TracingStartOptions{
		Screenshots: playwright.Bool(true),
		Snapshots:   playwright.Bool(true),
	}))

	_, err := page.Goto(server.PREFIX + "/grid.html")
	require.NoError(t, err)
	dir := t.TempDir()
	err = context.Tracing().Stop(filepath.Join(dir, "trace.zip"))
	require.NoError(t, err)
	require.FileExists(t, filepath.Join(dir, "trace.zip"))
}

func TestTracingStartStop(t *testing.T) {
	BeforeEach(t)

	require.NoError(t, context.Tracing().Start())
	require.NoError(t, context.Tracing().Stop())
}

func TestBrowserContextShouldNoErrorWhenStoppingWithoutStart(t *testing.T) {
	BeforeEach(t)

	require.NoError(t, context.Tracing().Stop())
}

func TestBrowserContextOutputTraceChunk(t *testing.T) {
	BeforeEach(t)

	require.NoError(t, context.Tracing().Start(playwright.TracingStartOptions{
		Screenshots: playwright.Bool(true),
		Snapshots:   playwright.Bool(true),
	}))

	_, err := page.Goto(server.PREFIX + "/grid.html")
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

	require.NoError(t, context.Tracing().Start(playwright.TracingStartOptions{
		Screenshots: playwright.Bool(true),
		Snapshots:   playwright.Bool(true),
	}))

	_, err := page.Goto(server.PREFIX + "/frames/frame.html")
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
	page1, err := context1.NewPage()
	require.NoError(t, err)
	_, err = page1.Goto(server.PREFIX + "/frames/frame.html")
	require.NoError(t, err)
	require.NoError(t, context1.Tracing().StartChunk())
	require.NoError(t, page1.SetContent("<button>Click</button>"))
	require.NoError(t, page1.Locator("button").Click())
	dir := t.TempDir()
	require.NoError(t, context1.Tracing().StopChunk(filepath.Join(dir, "trace.zip")))
	require.FileExists(t, filepath.Join(dir, "trace.zip"))
}

func TestShouldShowTracingGroupInActionList(t *testing.T) {
	BeforeEach(t)

	require.NoError(t, context.Tracing().Start())
	page, err := context.NewPage()
	require.NoError(t, err)

	require.NoError(t, context.Tracing().Group("outer group"))
	_, err = page.Goto(`data:text/html,<!DOCTYPE html><body><div>Hello world</div></body>`)
	require.NoError(t, err)
	require.NoError(t, context.Tracing().Group("inner group 1"))
	require.NoError(t, page.Locator("body").Click())
	require.NoError(t, context.Tracing().GroupEnd())
	require.NoError(t, context.Tracing().Group("inner group 2"))
	visiable, err := page.GetByText("Hello").IsVisible()
	require.NoError(t, err)
	require.True(t, visiable)
	require.NoError(t, context.Tracing().GroupEnd())
	require.NoError(t, context.Tracing().GroupEnd())

	tracePath := filepath.Join(t.TempDir(), "trace.zip")
	require.NoError(t, context.Tracing().Stop(tracePath))
	require.FileExists(t, tracePath)

	_, events := parseTrace(t, tracePath)
	actions := getTraceActions(events)
	require.Equal(t,
		[]string{
			"BrowserContext.NewPage",
			"outer group",
			"Page.Goto",
			"inner group 1",
			"Locator.Click",
			"inner group 2",
			"Locator.IsVisible",
		}, actions)
}

func parseTrace(t *testing.T, tracePath string) (files map[string][]byte, events []interface{}) {
	t.Helper()
	// read and unzip trace
	r, err := zip.OpenReader(tracePath)
	require.NoError(t, err)
	defer r.Close()

	files = make(map[string][]byte)
	events = make([]interface{}, 0)
	actionMap := make(map[string]interface{})
	for _, f := range r.File {
		rc, err := f.Open()
		require.NoError(t, err)
		defer rc.Close()

		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, rc)
		require.NoError(t, err)

		files[f.Name] = buf.Bytes()
		if f.Name == "trace.trace" || f.Name == "trace.network" {
			// read lines
			for _, line := range bytes.Split(buf.Bytes(), []byte("\n")) {
				if len(line) == 0 {
					continue
				}

				var event map[string]interface{}
				err := json.Unmarshal(line, &event)
				require.NoError(t, err)
				switch event["type"].(string) {
				case "before":
					event["type"] = "action"
					actionMap[event["callId"].(string)] = event
					events = append(events, event)
				case "input":
					break
				case "after":
					break
				default:
					events = append(events, event)
				}
			}
		}
	}

	return
}

func getTraceActions(events []interface{}) []string {
	actions := make([]string, 0)
	actionEvents := slices.DeleteFunc(events, func(e interface{}) bool {
		event := e.(map[string]interface{})
		return event["type"].(string) != "action"
	})
	slices.SortFunc(actionEvents, func(a, b interface{}) int {
		eventA := a.(map[string]interface{})
		eventB := b.(map[string]interface{})
		t1 := eventA["startTime"].(float64)
		t2 := eventB["startTime"].(float64)
		if t1 < t2 {
			return -1
		}
		if t1 > t2 {
			return 1
		}
		return 0
	})
	for _, e := range actionEvents {
		event := e.(map[string]interface{})
		actions = append(actions, event["apiName"].(string))
	}
	return actions
}
