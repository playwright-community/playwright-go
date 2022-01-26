package playwright_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestPageRouteContinue(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetExtraHTTPHeaders(map[string]string{
		"extra-http": "42",
	}))
	intercepted := make(chan bool, 1)
	err := page.Route("**/empty.html", func(route playwright.Route, request playwright.Request) {
		require.Equal(t, route.Request(), request)
		require.Contains(t, request.URL(), "empty.html")
		require.True(t, len(request.Headers()["user-agent"]) > 5)
		require.Equal(t, "42", request.Headers()["extra-http"])
		require.Equal(t, "GET", request.Method())

		postData, err := request.PostData()
		require.NoError(t, err)
		require.Equal(t, "", postData)
		require.True(t, request.IsNavigationRequest())
		require.Equal(t, "document", request.ResourceType())
		require.Equal(t, request.Frame(), page.MainFrame())
		require.Equal(t, "about:blank", request.Frame().URL())
		require.NoError(t, route.Continue())
		intercepted <- true
	})
	require.NoError(t, err)
	response, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.True(t, response.Ok())
	<-intercepted
}

func TestRouteContinueOverwrite(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	serverRequestChan := server.WaitForRequestChan("/sleep.zzz")
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.Route("**/*", func(route playwright.Route, request playwright.Request) {
		headers := request.Headers()
		headers["Foo"] = "bar"
		require.NoError(t, route.Continue(playwright.RouteContinueOptions{
			Method:   playwright.String("POST"),
			Headers:  headers,
			PostData: "foobar",
		}))
	}))
	_, err = page.Evaluate(`() => fetch("/sleep.zzz")`)
	require.NoError(t, err)
	serverRequest := <-serverRequestChan
	require.Equal(t, "POST", serverRequest.Method)
	require.Equal(t, "bar", serverRequest.Header.Get("Foo"))
	respData, err := ioutil.ReadAll(serverRequest.Body)
	require.NoError(t, err)
	require.Equal(t, "foobar", string(respData))
}

func TestRouteContinueOverwriteBodyBytes(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	serverRequestChan := server.WaitForRequestChan("/sleep.zzz")
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.Route("**/*", func(route playwright.Route, request playwright.Request) {
		require.NoError(t, route.Continue(playwright.RouteContinueOptions{
			Method:   playwright.String("POST"),
			PostData: []byte("foobar"),
		}))
	}))
	_, err = page.Evaluate(`() => fetch("/sleep.zzz")`)
	require.NoError(t, err)
	serverRequest := <-serverRequestChan
	require.Equal(t, "POST", serverRequest.Method)
	respData, err := ioutil.ReadAll(serverRequest.Body)
	require.NoError(t, err)
	require.Equal(t, "foobar", string(respData))
}

func TestRouteFulfill(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	requestsChan := make(chan playwright.Request, 1)
	err := page.Route("**/empty.html", func(route playwright.Route, request playwright.Request) {
		require.Equal(t, route.Request(), request)
		require.Contains(t, request.URL(), "empty.html")
		require.True(t, len(request.Headers()["user-agent"]) > 5)
		require.Equal(t, "GET", request.Method())

		postData, err := request.PostData()
		require.NoError(t, err)
		require.Equal(t, "", postData)
		require.True(t, request.IsNavigationRequest())
		require.Equal(t, "document", request.ResourceType())
		require.Equal(t, page, page.MainFrame().Page())
		require.Equal(t, request.Frame(), page.MainFrame())
		require.Equal(t, "about:blank", request.Frame().URL())
		require.NoError(t, route.Fulfill(playwright.RouteFulfillOptions{
			Body:        "123",
			ContentType: playwright.String("text/plain"),
			Headers: map[string]string{
				"Foo": "bar",
			},
		}))
		requestsChan <- request
	})
	require.NoError(t, err)
	response, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.True(t, response.Ok())
	request := <-requestsChan
	require.Equal(t, request, response.Request())
	text, err := response.Text()
	require.NoError(t, err)
	require.Equal(t, "123", text)
	require.Equal(t, "bar", response.Headers()["foo"])
	require.Equal(t, "text/plain", response.Headers()["content-type"])
}

func TestRouteFulfillByteSlice(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	intercepted := make(chan bool, 1)
	err := page.Route("**/empty.html", func(route playwright.Route, request playwright.Request) {
		require.NoError(t, route.Fulfill(playwright.RouteFulfillOptions{
			Body:        []byte("123"),
			ContentType: playwright.String("text/plain"),
		}))
		intercepted <- true
	})
	require.NoError(t, err)
	response, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.True(t, response.Ok())
	<-intercepted
	text, err := response.Text()
	require.NoError(t, err)
	require.Equal(t, "123", text)
	require.Equal(t, "3", response.Headers()["content-length"])
	require.Equal(t, "text/plain", response.Headers()["content-type"])
}

func TestRouteFulfillPath(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	intercepted := make(chan bool, 1)
	err := page.Route("**/empty.html", func(route playwright.Route, request playwright.Request) {
		require.NoError(t, route.Fulfill(playwright.RouteFulfillOptions{
			Path: playwright.String(Asset("pptr.png")),
		}))
		intercepted <- true
	})
	require.NoError(t, err)
	response, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.True(t, response.Ok())
	<-intercepted
	body, err := response.Body()
	require.NoError(t, err)
	require.True(t, len(body) > 5000)
	require.Equal(t, "image/png", response.Headers()["content-type"])
}

func TestRequestFinished(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	eventsStorage := newSyncSlice()
	var request playwright.Request
	page.Once("request", func(r playwright.Request) {
		request = r
		eventsStorage.Append("request")
	})
	page.Once("response", func() {
		eventsStorage.Append("response")
	})
	response, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	response.Finished()
	eventsStorage.Append("requestfinished")
	require.Equal(t, []interface{}{"request", "response", "requestfinished"}, eventsStorage.Get())
	require.Equal(t, response.Request(), request)
	require.Equal(t, response.Frame(), page.MainFrame())
}

func TestResponsePostData(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	requestData := map[string]interface{}{
		"foo": "bar123",
		"kek": true,
	}
	server.SetRoute("/foobar", func(w http.ResponseWriter, r *http.Request) {
		require.NoError(t, json.NewEncoder(w).Encode(requestData))
	})
	response, err := page.Goto(server.PREFIX + "/foobar")
	require.NoError(t, err)
	var actualResponse map[string]interface{}
	require.NoError(t, response.JSON(&actualResponse))
	require.Equal(t, requestData, actualResponse)
}

func TestRouteAbort(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	failedRequests := make(chan playwright.Request, 1)
	page.Once("requestfailed", func(request playwright.Request) {
		failedRequests <- request
	})
	err := page.Route("**/empty.html", func(route playwright.Route, request playwright.Request) {
		require.NoError(t, route.Abort("aborted"))
	})
	require.NoError(t, err)
	_, err = page.Goto(server.EMPTY_PAGE)
	require.Error(t, err)
	request := <-failedRequests
	require.True(t, len(request.Failure().ErrorText) > 5)
}

func TestRequestPostData(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	server.SetRoute("/foobar", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.Route("**/foobar", func(route playwright.Route, request playwright.Request) {
		var postData map[string]interface{}
		require.NoError(t, request.PostDataJSON(&postData))
		require.Equal(t, map[string]interface{}{
			"foo": true,
			"kek": float64(123),
		}, postData)
		raw, err := request.PostDataBuffer()
		require.NoError(t, err)
		require.Equal(t, []byte(`{"foo":true,"kek":123}`), raw)
		require.NoError(t, route.Continue())
	}))
	_, err = page.Evaluate(`url => fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify({
			"foo": true,
			"kek": 123,
		})
	})`, server.PREFIX+"/foobar")
	require.NoError(t, err)
}
