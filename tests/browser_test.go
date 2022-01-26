package playwright_test

import (
	"testing"

	"github.com/playwright-community/playwright-go"
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
	intercepted := make(chan bool, 1)
	err := page.Route("**/empty.html", func(route playwright.Route, request playwright.Request) {
		v, ok := request.Headers()["extra-http"]
		require.True(t, ok)
		require.Equal(t, "42", v)
		require.NoError(t, route.Continue())
		intercepted <- true
	})
	require.NoError(t, err)
	_, err = page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	<-intercepted
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

func TestBrowserNewPageWithExtraHTTPHeaders(t *testing.T) {
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

	intercepted := make(chan bool, 1)
	err = page.Route("**/empty.html", func(route playwright.Route, request playwright.Request) {
		v, ok := request.Headers()["extra-http"]
		require.True(t, ok)
		require.Equal(t, "42", v)
		require.NoError(t, route.Continue())
		intercepted <- true
	})
	require.NoError(t, err)
	_, err = page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	<-intercepted

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

func TestNewBrowserCDPSession(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	cdpSession, err := browser.NewBrowserCDPSession()
	if isChromium {
		require.NoError(t, err)
		require.NoError(t, cdpSession.Detach())
	} else {
		require.Error(t, err)
	}
}

func TestBrowserClose(t *testing.T) {
	browser, err := browserType.Launch()
	require.NoError(t, err)
	onCloseWasCalled := make(chan bool, 1)
	onClose := func() {
		onCloseWasCalled <- true
	}
	browser.On("disconnected", onClose)
	require.True(t, browser.IsConnected())
	require.NoError(t, browser.Close())
	<-onCloseWasCalled
	require.False(t, browser.IsConnected())
}
