package playwright_test

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestBrowserIsConnected(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.True(t, browser.IsConnected())
}

func TestBrowserVersion(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.Greater(t, len(browser.Version()), 2)
}

func TestBrowserNewContext(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.Equal(t, 1, len(context.Pages()))
}

func TestBrowserNewContextWithExtraHTTPHeaders(t *testing.T) {
	newContextWithOptions(t, playwright.BrowserNewContextOptions{
		ExtraHttpHeaders: map[string]string{"extra-http": "42"},
	})
	defer AfterEach(t)
	require.Equal(t, 1, len(context.Pages()))
}

func TestBrowserNewPage(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.Equal(t, 1, len(browser.Contexts()))
	page, err := browser.NewPage()
	require.NoError(t, err)
	require.Equal(t, 2, len(browser.Contexts()))
	require.False(t, page.IsClosed())
	require.NoError(t, page.Close())
	require.True(t, page.IsClosed())
	require.Equal(t, 1, len(browser.Contexts()))
}

func TestBrwoserNewPageWithExtraHTTPHeaders(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.Equal(t, 1, len(browser.Contexts()))
	page, err := browser.NewPage(playwright.BrowserNewContextOptions{
		ExtraHttpHeaders: map[string]string{
			"extra-http": "42",
		},
	})
	require.NoError(t, err)
	require.Equal(t, 2, len(browser.Contexts()))
	require.False(t, page.IsClosed())
	require.NoError(t, page.Close())
	require.True(t, page.IsClosed())
	require.Equal(t, 1, len(browser.Contexts()))
}

func TestBrowserShouldErrorUponSecondCreateNewPage(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	page, err := browser.NewPage()
	require.NoError(t, err)
	_, err = page.Context().NewPage()
	require.Error(t, err)
	require.Equal(t, "Please use browser.NewContext()", err.Error())
	require.NoError(t, page.Close())
}

func TestBrowserClose(t *testing.T) {
	pw, err := playwright.Run()
	require.NoError(t, err)
	browser, err := pw.Chromium.Launch()
	require.NoError(t, err)
	onCloseWasCalled := make(chan bool, 1)
	onClose := func() {
		onCloseWasCalled <- true
	}
	browser.On("disconnected", onClose)
	require.True(t, browser.IsConnected())
	require.NoError(t, browser.Close())
	<-onCloseWasCalled
	require.NoError(t, pw.Stop())
	require.False(t, browser.IsConnected())
}
