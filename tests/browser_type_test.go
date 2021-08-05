package playwright_test

import (
	"testing"

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
	remote_server := newRemoteServer()
	defer remote_server.Close()
	browser, err := browserType.Connect(remote_server.url)
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
	remote_server := newRemoteServer()
	defer remote_server.Close()
	browser, err := browserType.Connect(remote_server.url)
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

	browser, err = browserType.Connect(remote_server.url)
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
	remote_server := newRemoteServer()
	disconnected1 := newSyncSlice()
	disconnected2 := newSyncSlice()
	browser1, err := browserType.Connect(remote_server.url)
	require.NoError(t, err)
	require.NotNil(t, browser1)
	browser2, err := browserType.Connect(remote_server.url)
	require.NoError(t, err)
	require.NotNil(t, browser2)
	browser1.On("disconnected", func() {
		disconnected1.Append(true)
	})
	browser2.On("disconnected", func() {
		disconnected2.Append(true)
	})
	page, err := browser2.NewPage()
	require.NoError(t, err)
	require.NoError(t, browser1.Close())
	require.False(t, browser1.IsConnected())
	require.Len(t, disconnected1.Get(), 1)
	require.Len(t, disconnected2.Get(), 0)
	remote_server.Close()

	require.Panics(t, func() {
		_, err = page.Title()
	})
	require.False(t, browser2.IsConnected())
	require.Len(t, disconnected2.Get(), 1)
}
