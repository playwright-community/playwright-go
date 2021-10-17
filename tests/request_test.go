package playwright_test

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestRequestFulfill(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	requests := make(chan playwright.Request, 1)
	routes := make(chan playwright.Route, 1)
	err := page.Route("**/empty.html", func(route playwright.Route, request playwright.Request) {
		requests <- request
		routes <- route
		err := route.Fulfill(playwright.RouteFulfillOptions{
			Body: "Hello World",
		})
		require.NoError(t, err)
	})
	require.NoError(t, err)
	response, err := page.Goto(server.EMPTY_PAGE)
	response.Finished()
	require.NoError(t, err)
	require.True(t, response.Ok())
	text, err := response.Text()
	require.NoError(t, err)
	require.Equal(t, "Hello World", text)
	request := <-requests
	route := <-routes
	headers, err := route.Request().AllHeaders()
	require.NoError(t, err)
	require.Contains(t, headers, "accept")
	require.Equal(t, route.Request(), request)
	require.Contains(t, request.URL(), "empty.html")
	require.Equal(t, "GET", request.Method())
	require.Contains(t, headers, "user-agent")
	postData, err := request.PostData()
	require.NoError(t, err)
	require.Equal(t, "", postData)
	require.True(t, request.IsNavigationRequest())
	require.Equal(t, request.ResourceType(), "document")
	require.Equal(t, request.Frame(), page.MainFrame())
}

func TestRequestContinue(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	intercepted := make(chan bool, 1)
	err := page.Route("**/empty.html", func(route playwright.Route, request playwright.Request) {
		intercepted <- true
		err := route.Continue()
		require.NoError(t, err)
	})
	require.NoError(t, err)
	response, err := page.Goto(server.EMPTY_PAGE)
	response.Finished()
	require.NoError(t, err)
	require.True(t, response.Ok())
	require.True(t, <-intercepted)
}

func TestRequestShouldFireForNavigationRequests(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	requests := make(chan playwright.Request, 1)
	page.On("request", func(request playwright.Request) {
		requests <- request
	})
	response, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	response.Finished()
	require.True(t, response.Ok())
	<-requests
}

func TestShouldReportRequestHeadersArray(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	request, err := page.ExpectRequest("*/**", func() error {
		_, err := page.Evaluate(`
		() => fetch('/headers', {
            headers: [
                ['header-a', 'value-a'],
                ['header-a', 'value-a-1'],
                ['header-a', 'value-a-2'],
                ['header-b', 'value-b'],
            ]
            })
		`)
		return err
	})
	require.NoError(t, err)
	headers, err := request.AllHeaders()
	require.NoError(t, err)
	require.Contains(t, headers, "header-a")
	valueA, err := request.HeaderValue("header-a")
	require.NoError(t, err)
	require.Equal(t, "value-a, value-a-1, value-a-2", valueA)
	valuesA, err := request.HeaderValue("not-there")
	require.NoError(t, err)
	require.Equal(t, "", valuesA)
}
