package playwright_test

import (
	"net/http"
	"net/url"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/playwright-community/playwright-go"
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
	valueA, err = request.HeaderValue("not-there")
	require.NoError(t, err)
	require.Equal(t, "", valueA)
}

func TestShouldReportResponseHeadersArray(t *testing.T) {
	if isWebKit && runtime.GOOS == "windows" {
		t.Skip("libcurl does not support non-set-cookie multivalue headers")
	}
	BeforeEach(t)
	defer AfterEach(t)
	server.SetRoute("/headers", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("header-a", "value-a")
		rw.Header().Add("header-a", "value-a-1")
		rw.Header().Add("header-a", "value-a-2")
		rw.Header().Add("header-b", "value-b")
		rw.Header().Add("set-cookie", "a=b")
		rw.Header().Add("set-cookie", "c=d")
	})
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	response, err := page.ExpectResponse("*/**", func() error {
		_, err := page.Evaluate(`
		() => fetch('/headers')
		`)
		return err
	})
	require.NoError(t, err)
	actual_headers := make(map[string][]string)
	pw_headers, err := response.HeadersArray()
	require.NoError(t, err)
	for _, header := range pw_headers {
		if _, ok := actual_headers[strings.ToLower(header.Name)]; !ok {
			actual_headers[strings.ToLower(header.Name)] = make([]string, 0)
		}
		actual_headers[strings.ToLower(header.Name)] = append(actual_headers[strings.ToLower(header.Name)], header.Value)
	}
	for _, k := range []string{"Keep-Alive", "Connection", "Date", "Transfer-Encoding"} {
		delete(actual_headers, strings.ToLower(k))
	}
	require.Equal(t, actual_headers, map[string][]string{
		"header-a":       {"value-a", "value-a-1", "value-a-2"},
		"header-b":       {"value-b"},
		"set-cookie":     {"a=b", "c=d"},
		"content-length": {"0"},
	})
	header, err := response.HeaderValue("header-b")
	require.NoError(t, err)
	require.Equal(t, "value-b", header)
	header, err = response.HeaderValue("set-cookie")
	require.NoError(t, err)
	require.Equal(t, "a=b\nc=d", header)
	header, err = response.HeaderValue("not-there")
	require.NoError(t, err)
	require.Equal(t, "", header)
	headers, err := response.HeaderValues("header-a")
	require.NoError(t, err)
	sort.Strings(headers)
	require.Equal(t, []string{"value-a", "value-a-1", "value-a-2"}, headers)
	headers, err = response.HeaderValues("header-b")
	require.NoError(t, err)
	sort.Strings(headers)
	require.Equal(t, []string{"value-b"}, headers)
	headers, err = response.HeaderValues("set-cookie")
	require.NoError(t, err)
	sort.Strings(headers)
	require.Equal(t, []string{"a=b", "c=d"}, headers)
	headers, err = response.HeaderValues("not-there")
	require.NoError(t, err)
	require.Equal(t, []string{}, headers)
}

func TestShouldReportResponseServerAddr(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	response, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	server_addr, err := response.ServerAddr()
	require.NoError(t, err)
	require.NotNil(t, server_addr)
	url, err := url.Parse(server.PREFIX)
	require.NoError(t, err)
	require.Equal(t, url.Port(), strconv.Itoa(server_addr.Port))
}
