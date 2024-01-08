package playwright_test

import (
	"errors"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestBrowserTypeBrowserName(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.Equal(t, pw.Chromium.Name(), "chromium")
	require.Equal(t, pw.Firefox.Name(), "firefox")
	require.Equal(t, pw.WebKit.Name(), "webkit")
}

func TestBrowserTypeExecutablePath(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.Greater(t, len(pw.Chromium.ExecutablePath()), 0)
}

func TestBrowserTypeLaunchPersistentContext(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	user_data_dir1 := t.TempDir()
	browser_context, err := browserType.LaunchPersistentContext(user_data_dir1)
	require.NoError(t, err)
	page, err := browser_context.NewPage()
	require.NoError(t, err)
	_, err = page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = page.Evaluate("() => localStorage.hey = 'hello'")
	require.NoError(t, err)
	require.NoError(t, browser_context.Close())

	browser_context2, err := browserType.LaunchPersistentContext(user_data_dir1)
	require.NoError(t, err)
	page2, err := browser_context2.NewPage()
	require.NoError(t, err)
	_, err = page2.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	result, err := page2.Evaluate("() => localStorage.hey")
	require.NoError(t, err)
	require.Equal(t, "hello", result)
	require.NoError(t, browser_context2.Close())

	user_data_dir2 := t.TempDir()
	browser_context3, err := browserType.LaunchPersistentContext(user_data_dir2)
	require.NoError(t, err)
	page3, err := browser_context3.NewPage()
	require.NoError(t, err)
	_, err = page3.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	result, err = page3.Evaluate("() => localStorage.hey")
	require.NoError(t, err)
	require.NotEqual(t, "hello", result)
	require.NoError(t, browser_context3.Close())
}

func TestBrowserTypeConnect(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	remoteServer, err := newRemoteServer()
	require.NoError(t, err)
	defer remoteServer.Close()
	browser, err := browserType.Connect(remoteServer.url)
	require.NoError(t, err)
	require.NotNil(t, browser)
	browser_context, err := browser.NewContext()
	require.NoError(t, err)
	page, err := browser_context.NewPage()
	require.NoError(t, err)
	result, err := page.Evaluate("11 * 11")
	require.NoError(t, err)
	require.Equal(t, result, 121)
	require.NoError(t, browser.Close())
}

func TestBrowserTypeConnectShouldBeAbleToReconnectToBrowser(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	remoteServer, err := newRemoteServer()
	require.NoError(t, err)
	defer remoteServer.Close()
	browser, err := browserType.Connect(remoteServer.url)
	require.NoError(t, err)
	require.NotNil(t, browser)
	require.Len(t, browser.Contexts(), 0)
	browser_context, err := browser.NewContext()
	require.NoError(t, err)
	require.Len(t, browser.Contexts(), 1)
	require.Len(t, browser_context.Pages(), 0)
	page, err := browser_context.NewPage()
	require.Len(t, browser_context.Pages(), 1)
	require.NoError(t, err)
	result, err := page.Evaluate("11 * 11")
	require.NoError(t, err)
	require.Equal(t, result, 121)
	require.NoError(t, browser.Close())

	browser, err = browserType.Connect(remoteServer.url)
	require.NoError(t, err)
	require.NotNil(t, browser)
	require.Len(t, browser.Contexts(), 0)
	browser_context, err = browser.NewContext()
	require.NoError(t, err)
	require.Len(t, browser.Contexts(), 1)
	require.Len(t, browser_context.Pages(), 0)
	page, err = browser_context.NewPage()
	require.Len(t, browser_context.Pages(), 1)
	require.NoError(t, err)
	result, err = page.Evaluate("11 * 11")
	require.NoError(t, err)
	require.Equal(t, result, 121)
	require.NoError(t, browser.Close())
}

func TestBrowserTypeConnectShouldEmitDisconnectedEvent(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	remoteServer, err := newRemoteServer()
	require.NoError(t, err)
	disconnected1 := newSyncSlice[bool]()
	disconnected2 := newSyncSlice[bool]()
	browser1, err := browserType.Connect(remoteServer.url)
	require.NoError(t, err)
	require.NotNil(t, browser1)
	browser2, err := browserType.Connect(remoteServer.url)
	require.NoError(t, err)
	require.NotNil(t, browser2)
	browser1.OnDisconnected(func(playwright.Browser) {
		disconnected1.Append(true)
	})
	browser2.OnDisconnected(func(playwright.Browser) {
		disconnected2.Append(true)
	})
	page, err := browser2.NewPage()
	require.NoError(t, err)
	require.NoError(t, browser1.Close())
	require.False(t, browser1.IsConnected())
	require.Len(t, disconnected1.Get(), 1)
	require.Len(t, disconnected2.Get(), 0)
	remoteServer.Close()

	_, err = page.Title()
	require.Error(t, err)
	require.False(t, browser2.IsConnected())
	require.Len(t, disconnected2.Get(), 1)
}

func TestBrowserTypeConnectSlowMo(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	remoteServer, err := newRemoteServer()
	require.NoError(t, err)
	defer remoteServer.Close()
	browser, err := browserType.Connect(remoteServer.url, playwright.BrowserTypeConnectOptions{
		SlowMo: playwright.Float(100),
	})
	require.NoError(t, err)
	require.NotNil(t, browser)
	browser_context, err := browser.NewContext()
	require.NoError(t, err)
	t1 := time.Now()
	page, err := browser_context.NewPage()
	require.NoError(t, err)
	result, err := page.Evaluate("11 * 11")
	require.NoError(t, err)
	require.Equal(t, result, 121)
	require.GreaterOrEqual(t, time.Since(t1), time.Duration(time.Millisecond*200))
	require.NoError(t, browser.Close())
}

func TestBrowserTypeConnectArtifactPath(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	remoteServer, err := newRemoteServer()
	require.NoError(t, err)
	defer remoteServer.Close()
	browser, err := browserType.Connect(remoteServer.url)
	require.NoError(t, err)
	require.NotNil(t, browser)
	defer browser.Close()
	recordVideoDir := t.TempDir()
	browserContext, err := browser.NewContext(playwright.BrowserNewContextOptions{
		RecordVideo: &playwright.RecordVideo{
			Dir: recordVideoDir,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, browserContext)
	defer browserContext.Close()
	page, err := browserContext.NewPage()
	require.NoError(t, err)
	defer page.Close()
	_, err = page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = page.Video().Path()
	require.Error(t, err)
	require.Equal(t, err, errors.New("Path is not available when connecting remotely. Use SaveAs() to save a local copy."))
}
func TestBrowserTypeConnectOverCDP(t *testing.T) {
	if !isChromium {
		t.Skip("CDP is only supported on Chromium")
	}
	BeforeEach(t)
	defer AfterEach(t)
	port, err := getFreePort()
	require.NoError(t, err)
	browserServer, err := browserType.Launch(playwright.BrowserTypeLaunchOptions{
		Args: []string{fmt.Sprintf("--remote-debugging-port=%d", port)},
	})
	require.NoError(t, err)
	defer browserServer.Close()
	browser, err := browserType.ConnectOverCDP(fmt.Sprintf("http://localhost:%d", port))
	require.NoError(t, err)
	require.NotNil(t, browser)
	defer browser.Close()
	require.Len(t, browser.Contexts(), 1)
}

func TestBrowserTypeConnectOverCDPTwice(t *testing.T) {
	if !isChromium {
		t.Skip("CDP is only supported on Chromium")
	}
	BeforeEach(t)
	defer AfterEach(t)
	port, err := getFreePort()
	require.NoError(t, err)
	browserServer, err := browserType.Launch(playwright.BrowserTypeLaunchOptions{
		Args: []string{fmt.Sprintf("--remote-debugging-port=%d", port)},
	})
	require.NoError(t, err)
	defer browserServer.Close()
	browser1, err := browserType.ConnectOverCDP(fmt.Sprintf("http://localhost:%d", port))
	require.NoError(t, err)
	require.NotNil(t, browser1)
	browser2, err := browserType.ConnectOverCDP(fmt.Sprintf("http://localhost:%d", port))
	require.NoError(t, err)
	require.NotNil(t, browser2)
	defer browser1.Close()
	defer browser2.Close()
	require.Len(t, browser1.Contexts(), 1)
	page1, err := browser1.Contexts()[0].NewPage()
	require.NoError(t, err)
	_, err = page1.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.Len(t, browser2.Contexts(), 1)
	page2, err := browser2.Contexts()[0].NewPage()
	require.NoError(t, err)
	_, err = page2.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	require.Len(t, browser1.Contexts()[0].Pages(), 2)
	require.Len(t, browser2.Contexts()[0].Pages(), 2)
}

func TestSetInputFilesShouldPreserveLastModifiedTimestamp(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	remoteServer, err := newRemoteServer()
	require.NoError(t, err)
	defer remoteServer.Close()

	browser, err := browserType.Connect(remoteServer.url)
	require.NoError(t, err)
	require.NotNil(t, browser)
	browser_context, err := browser.NewContext()
	require.NoError(t, err)
	page, err := browser_context.NewPage()
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<input type=file multiple=true/>`))
	input := page.Locator("input")
	filenames := []string{
		"file-to-upload.txt",
		"file-to-upload-2.txt",
	}
	files := make([]string, len(filenames))
	for i, filename := range filenames {
		files[i] = Asset(filename)
	}
	require.NoError(t, input.SetInputFiles(files))
	result, err := input.Evaluate(`input => [...input.files].map(f => f.name)`, nil)
	require.NoError(t, err)
	names, ok := result.([]interface{})
	require.True(t, ok)
	for i, name := range names {
		require.Equal(t, filenames[i], name.(string))
	}

	results, err := input.Evaluate(`input => [...input.files].map(f => f.lastModified)`, nil)
	require.NoError(t, err)
	timestamps, ok := results.([]interface{})
	require.True(t, ok)
	for i, timestamp := range timestamps {
		expected, err := getFileLastModifiedTimeMs(files[i])
		require.NoError(t, err)
		// On Linux browser sometimes reduces the timestamp by 1ms: 1696272058110.0715  -> 1696272058109 or even
		// rounds it to seconds in WebKit: 1696272058110 -> 1696272058000.
		require.Less(t, math.Abs(float64(expected-int64(timestamp.(int)))), 1000.0)
	}
}
