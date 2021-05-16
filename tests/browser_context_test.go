package playwright_test

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestBrowserContextNewPage(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.Equal(t, context.Browser(), browser)
}

func TestBrowserContextNewContext(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.Equal(t, 1, len(browser.Contexts()))
	context2, err := browser.NewContext()
	require.NoError(t, err)
	require.Equal(t, 2, len(browser.Contexts()))
	require.Equal(t, browser.Contexts()[1], context2)
	require.Equal(t, context2.Browser(), browser)
	require.NoError(t, context2.Close())
	require.Equal(t, 1, len(browser.Contexts()))
	require.Equal(t, context.Browser(), browser)
}

func TestBrowserContextClose(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t, false)
	context2, err := browser.NewContext()
	require.NoError(t, err)
	require.Equal(t, 2, len(browser.Contexts()))
	require.Equal(t, context, browser.Contexts()[0])
	require.Equal(t, context2, browser.Contexts()[1])
	require.NoError(t, context.Close())
	require.Equal(t, 1, len(browser.Contexts()))
	require.NoError(t, context2.Close())
	require.Equal(t, 0, len(browser.Contexts()))
}

func TestBrowserContextOffline(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	offline, err := page.Evaluate("window.navigator.onLine")
	require.NoError(t, err)
	require.True(t, offline.(bool))

	require.NoError(t, context.SetOffline(true))
	offline, err = page.Evaluate("window.navigator.onLine")
	require.NoError(t, err)
	require.False(t, offline.(bool))

	require.NoError(t, context.SetOffline(false))
	offline, err = page.Evaluate("window.navigator.onLine")
	require.NoError(t, err)
	require.True(t, offline.(bool))
}

func TestBrowserContextSetExtraHTTPHeaders(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, context.SetExtraHTTPHeaders(map[string]string{
		"extra-http": "42",
	}))
	intercepted := make(chan bool, 1)
	err := page.Route("**/empty.html", func(route playwright.Route, request playwright.Request) {
		require.NoError(t, route.Continue())
		intercepted <- true
	})
	require.NoError(t, err)
	response, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.True(t, response.Ok())
	<-intercepted
}

func TestBrowserContextNewCDPSession(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	cdpSession, err := page.Context().NewCDPSession(page)
	require.NoError(t, err)
	require.NoError(t, cdpSession.Detach())
}

func TestBrowserContextSetGeolocation(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, context.GrantPermissions([]string{"geolocation"}))
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, context.SetGeolocation(&playwright.SetGeolocationOptions{
		Longitude: 10,
		Latitude:  10,
	}))
	geolocation, err := page.Evaluate(`() => new Promise(resolve => navigator.geolocation.getCurrentPosition(position => {
      resolve({latitude: position.coords.latitude, longitude: position.coords.longitude});
    }))`)
	require.NoError(t, err)
	require.Equal(t, geolocation, map[string]interface{}{
		"latitude":  10,
		"longitude": 10,
	})
	require.NoError(t, context.ClearPermissions())
}

func TestBrowserContextAddCookies(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, context.AddCookies(playwright.SetNetworkCookieParam{
		URL:   playwright.String(server.EMPTY_PAGE),
		Name:  "password",
		Value: "123456",
	}))
	cookie, err := page.Evaluate("() => document.cookie")
	require.NoError(t, err)
	require.Equal(t, "password=123456", cookie)

	cookies, err := context.Cookies()
	require.NoError(t, err)
	require.Equal(t, []*playwright.NetworkCookie{
		{
			Name:     "password",
			Value:    "123456",
			Domain:   "127.0.0.1",
			Path:     "/",
			Expires:  -1,
			HttpOnly: false,
			Secure:   false,
			SameSite: "None",
		},
	}, cookies)

	require.NoError(t, page.Context().ClearCookies())
	_, err = page.Reload()
	require.NoError(t, err)
	cookie, err = page.Evaluate("() => document.cookie")
	require.NoError(t, err)
	require.Equal(t, "", cookie)
}

func TestBrowserContextAddInitScript(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, context.AddInitScript(playwright.BrowserContextAddInitScriptOptions{
		Script: playwright.String(`window['injected'] = 123;`),
	}))
	_, err := page.Goto(server.PREFIX + "/tamperable.html")
	require.NoError(t, err)
	result, err := page.Evaluate(`() => window['result']`)
	require.NoError(t, err)
	require.Equal(t, 123, result)
}

func TestBrowserContextAddInitScriptWithPath(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, context.AddInitScript(playwright.BrowserContextAddInitScriptOptions{
		Path: playwright.String(Asset("injectedfile.js")),
	}))
	_, err := page.Goto(server.PREFIX + "/tamperable.html")
	require.NoError(t, err)
	result, err := page.Evaluate(`() => window['result']`)
	require.NoError(t, err)
	require.Equal(t, 123, result)
}

func TestBrowserContextWindowOpenshouldUseParentTabContext(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	popupEvent, err := page.ExpectEvent("popup", func() error {
		_, err := page.Evaluate("url => window.open(url)", server.EMPTY_PAGE)
		return err
	})
	require.NoError(t, err)
	popup := popupEvent.(playwright.Page)
	require.Equal(t, popup.Context(), context)
}

func TestBrowserContextUnrouteShouldWork(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)

	intercepted := []int{}
	handler1 := func(route playwright.Route, request playwright.Request) {
		intercepted = append(intercepted, 1)
		require.NoError(t, route.Continue())
	}
	require.NoError(t, context.Route("**/empty.html", handler1))
	require.NoError(t, context.Route("**/empty.html", func(route playwright.Route, request playwright.Request) {
		intercepted = append(intercepted, 2)
		require.NoError(t, route.Continue())
	}))
	require.NoError(t, context.Route("**/empty.html", func(route playwright.Route, request playwright.Request) {
		intercepted = append(intercepted, 3)
		require.NoError(t, route.Continue())
	}))
	require.NoError(t, context.Route("**/*", func(route playwright.Route, request playwright.Request) {
		intercepted = append(intercepted, 4)
		require.NoError(t, route.Continue())
	}))

	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.Equal(t, []int{1}, intercepted)

	intercepted = []int{}
	require.NoError(t, context.Unroute("**/empty.html", handler1))
	_, err = page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.Equal(t, []int{2}, intercepted)

	intercepted = []int{}
	require.NoError(t, context.Unroute("**/empty.html"))
	_, err = page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.Equal(t, []int{4}, intercepted)
}
