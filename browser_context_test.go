package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBrowserContextNewPage(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
}

func TestBrowserContextClose(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	context, err := helper.Browser.NewContext()
	require.NoError(t, err)
	require.Equal(t, 2, len(helper.Browser.Contexts()))
	require.Equal(t, helper.Context, helper.Browser.Contexts()[0])
	require.Equal(t, context, helper.Browser.Contexts()[1])
	require.NoError(t, helper.Context.Close())
	require.Equal(t, 1, len(helper.Browser.Contexts()))
	require.NoError(t, context.Close())
	require.Equal(t, 0, len(helper.Browser.Contexts()))
}

func TestBrowserContextOffline(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	offline, err := helper.Page.Evaluate("window.navigator.onLine")
	require.NoError(t, err)
	require.True(t, offline.(bool))

	require.NoError(t, helper.Context.SetOffline(true))
	offline, err = helper.Page.Evaluate("window.navigator.onLine")
	require.NoError(t, err)
	require.False(t, offline.(bool))

	require.NoError(t, helper.Context.SetOffline(false))
	offline, err = helper.Page.Evaluate("window.navigator.onLine")
	require.NoError(t, err)
	require.True(t, offline.(bool))
}

func TestBrowserContextSetExtraHTTPHeaders(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	require.NoError(t, helper.Context.SetExtraHTTPHeaders(map[string]string{
		"extra-http": "42",
	}))
	intercepted := make(chan bool, 1)
	err := helper.Page.Route("**/empty.html", func(route *Route, request *Request) {
		require.NoError(t, route.Continue())
		intercepted <- true
	})
	require.NoError(t, err)
	response, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.True(t, response.Ok())
	<-intercepted
}

func TestBrowserContextSetGeolocation(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	require.NoError(t, helper.Context.GrantPermissions([]string{"geolocation"}))
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Context.SetGeolocation(&SetGeolocationOptions{
		Longitude: 10,
		Latitude:  10,
	}))
	geolocation, err := helper.Page.Evaluate(`() => new Promise(resolve => navigator.geolocation.getCurrentPosition(position => {
      resolve({latitude: position.coords.latitude, longitude: position.coords.longitude});
    }))`)
	require.NoError(t, err)
	require.Equal(t, geolocation, map[string]interface{}{
		"latitude":  10,
		"longitude": 10,
	})
	require.NoError(t, helper.Context.ClearPermissions())
}

func TestBrowserContextAddCookies(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Context.AddCookies(SetNetworkCookieParam{
		URL:   String(helper.server.EMPTY_PAGE),
		Name:  "password",
		Value: "123456",
	}))
	cookie, err := helper.Page.Evaluate("() => document.cookie")
	require.NoError(t, err)
	require.Equal(t, "password=123456", cookie)

	cookies, err := helper.Context.Cookies()
	require.NoError(t, err)
	require.Equal(t, []*NetworkCookie{
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

	require.NoError(t, helper.Page.browserContext.ClearCookies())

	cookie, err = helper.Page.Evaluate("() => document.cookie")
	require.NoError(t, err)
	require.Equal(t, "", cookie)
}

func TestBrowserContextAddInitScript(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	require.NoError(t, helper.Page.AddInitScript(BrowserContextAddInitScriptOptions{
		Script: String(`window['injected'] = 123;`),
	}))
	_, err := helper.Page.Goto(helper.server.PREFIX + "/tamperable.html")
	require.NoError(t, err)
	result, err := helper.Page.Evaluate(`() => window['result']`)
	require.NoError(t, err)
	require.Equal(t, 123, result)
}
