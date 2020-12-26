package playwright

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRouteContinue(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	require.NoError(t, helper.Page.SetExtraHTTPHeaders(map[string]string{
		"extra-http": "42",
	}))
	intercepted := make(chan bool, 1)
	err := helper.Page.Route("**/empty.html", func(route *Route, request *Request) {
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
		require.Equal(t, request.Frame(), helper.Page.MainFrame())
		require.Equal(t, "about:blank", request.Frame().URL())
		require.NoError(t, route.Continue())
		intercepted <- true
	})
	require.NoError(t, err)
	response, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.True(t, response.Ok())
	<-intercepted
}

func TestRouteContinueOverwrite(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	serverRequestChan := helper.server.WaitForRequestChan("/sleep.zzz")
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.Route("**/*", func(route *Route, request *Request) {
		headers := request.Headers()
		headers["Foo"] = "bar"
		require.NoError(t, route.Continue(RouteContinueOptions{
			Method:   String("POST"),
			Headers:  headers,
			PostData: "foobar",
		}))
	}))
	_, err = helper.Page.Evaluate(`() => fetch("/sleep.zzz")`)
	require.NoError(t, err)
	serverRequest := <-serverRequestChan
	require.Equal(t, "POST", serverRequest.Method)
	require.Equal(t, "bar", serverRequest.Header.Get("Foo"))
	respData, err := ioutil.ReadAll(serverRequest.Body)
	require.NoError(t, err)
	require.Equal(t, "foobar", string(respData))
}

func TestRouteContinueOverwriteBodyBytes(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	serverRequestChan := helper.server.WaitForRequestChan("/sleep.zzz")
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.Route("**/*", func(route *Route, request *Request) {
		require.NoError(t, route.Continue(RouteContinueOptions{
			Method:   String("POST"),
			PostData: []byte("foobar"),
		}))
	}))
	_, err = helper.Page.Evaluate(`() => fetch("/sleep.zzz")`)
	require.NoError(t, err)
	serverRequest := <-serverRequestChan
	require.Equal(t, "POST", serverRequest.Method)
	respData, err := ioutil.ReadAll(serverRequest.Body)
	require.NoError(t, err)
	require.Equal(t, "foobar", string(respData))
}

func TestRouteFulfill(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	requestsChan := make(chan *Request, 1)
	err := helper.Page.Route("**/empty.html", func(route *Route, request *Request) {
		require.Equal(t, route.Request(), request)
		require.Contains(t, request.URL(), "empty.html")
		require.True(t, len(request.Headers()["user-agent"]) > 5)
		require.Equal(t, "GET", request.Method())

		postData, err := request.PostData()
		require.NoError(t, err)
		require.Equal(t, "", postData)
		require.True(t, request.IsNavigationRequest())
		require.Equal(t, "document", request.ResourceType())
		require.Equal(t, helper.Page, helper.Page.MainFrame().Page())
		require.Equal(t, request.Frame(), helper.Page.MainFrame())
		require.Equal(t, "about:blank", request.Frame().URL())
		require.NoError(t, route.Fulfill(RouteFulfillOptions{
			Body:        "123",
			ContentType: String("text/plain"),
			Headers: map[string]string{
				"Foo": "bar",
			},
		}))
		requestsChan <- request
	})
	require.NoError(t, err)
	response, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
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
	helper := BeforeEach(t)
	defer helper.AfterEach()
	intercepted := make(chan bool, 1)
	err := helper.Page.Route("**/empty.html", func(route *Route, request *Request) {
		require.NoError(t, route.Fulfill(RouteFulfillOptions{
			Body:        []byte("123"),
			ContentType: String("text/plain"),
		}))
		intercepted <- true
	})
	require.NoError(t, err)
	response, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
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
	helper := BeforeEach(t)
	defer helper.AfterEach()
	intercepted := make(chan bool, 1)
	err := helper.Page.Route("**/empty.html", func(route *Route, request *Request) {
		require.NoError(t, route.Fulfill(RouteFulfillOptions{
			Path: String(filepath.Join(helper.assetDir, "pptr.png")),
		}))
		intercepted <- true
	})
	require.NoError(t, err)
	response, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.True(t, response.Ok())
	<-intercepted
	body, err := response.Body()
	require.NoError(t, err)
	require.True(t, len(body) > 5000)
	require.Equal(t, "image/png", response.Headers()["content-type"])
}

func TestRequestFinished(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	eventsStorage := newSyncSlice()
	var request *Request
	helper.Page.Once("request", func(r *Request) {
		request = r
		eventsStorage.Append("request")
	})
	helper.Page.Once("response", func() {
		eventsStorage.Append("response")
	})
	response, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, response.Finished())
	eventsStorage.Append("requestfinished")
	require.Equal(t, []interface{}{"request", "response", "requestfinished"}, eventsStorage.Get())
	require.Equal(t, response.Request(), request)
	require.Equal(t, response.Frame(), helper.Page.MainFrame())
}

func TestResponsePostData(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	requestData := map[string]interface{}{
		"foo": "bar123",
		"kek": true,
	}
	helper.server.SetRoute("/foobar", func(w http.ResponseWriter, r *http.Request) {
		require.NoError(t, json.NewEncoder(w).Encode(requestData))
	})
	response, err := helper.Page.Goto(helper.server.PREFIX + "/foobar")
	require.NoError(t, err)
	var actualResponse map[string]interface{}
	require.NoError(t, response.JSON(&actualResponse))
	require.Equal(t, requestData, actualResponse)
}

func TestRouteAbort(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	failedRequests := make(chan *Request, 1)
	helper.Page.Once("requestfailed", func(request *Request) {
		failedRequests <- request
	})
	err := helper.Page.Route("**/empty.html", func(route *Route, request *Request) {
		require.NoError(t, route.Abort(String("aborted")))
	})
	require.NoError(t, err)
	_, err = helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.Error(t, err)
	request := <-failedRequests
	require.True(t, len(request.Failure().ErrorText) > 5)
}

func TestRequestPostData(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	helper.server.SetRoute("/foobar", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.Route("**/foobar", func(route *Route, request *Request) {
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
	_, err = helper.Page.Evaluate(`url => fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify({
			"foo": true,
			"kek": 123,
		})
	})`, helper.server.PREFIX+"/foobar")
	require.NoError(t, err)
}
