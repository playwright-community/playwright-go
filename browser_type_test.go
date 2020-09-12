package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBrowserTypeBrowserName(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	require.Equal(t, helper.Playwright.Chromium.Name(), "chromium")
	require.Equal(t, helper.Playwright.Firefox.Name(), "firefox")
	require.Equal(t, helper.Playwright.WebKit.Name(), "webkit")
}

func TestBrowserTypeExecutablePath(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	require.Greater(t, len(helper.Playwright.Chromium.ExecutablePath()), 0)
}

func TestBrowserTypeLaunchPersistentContext(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	user_data_dir1 := t.TempDir()
	browser_context, err := helper.BrowserType.LaunchPersistentContext(user_data_dir1)
	require.NoError(t, err)
	page, err := browser_context.NewPage()
	require.NoError(t, err)
	_, err = page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = page.Evaluate("() => localStorage.hey = 'hello'")
	require.NoError(t, err)
	require.NoError(t, browser_context.Close())

	browser_context2, err := helper.BrowserType.LaunchPersistentContext(user_data_dir1)
	require.NoError(t, err)
	page2, err := browser_context2.NewPage()
	require.NoError(t, err)
	_, err = page2.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	result, err := page2.Evaluate("() => localStorage.hey")
	require.NoError(t, err)
	require.Equal(t, "hello", result)
	require.NoError(t, browser_context2.Close())

	user_data_dir2 := t.TempDir()
	browser_context3, err := helper.BrowserType.LaunchPersistentContext(user_data_dir2)
	require.NoError(t, err)
	page3, err := browser_context3.NewPage()
	require.NoError(t, err)
	_, err = page3.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	result, err = page3.Evaluate("() => localStorage.hey")
	require.NoError(t, err)
	require.NotEqual(t, "hello", result)
	require.NoError(t, browser_context3.Close())
}
