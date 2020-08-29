package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRouteContinue(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	intercepted := make(chan bool, 1)
	helper.Page.Route("**/empty.html", func(route *Route, request *Request) {
		require.Equal(t, route.Request(), request)
		require.Contains(t, request.URL(), "empty.html")
		require.Greater(t, 5, len(request.Headers()["User-Agent"]))
		require.Equal(t, "GET", request.Method())

		postData, err := request.PostData()
		require.NoError(t, err)
		require.Equal(t, "", postData)
		require.True(t, request.IsNavigationRequest())
		require.Equal(t, "document", request.ResourceType())
		require.Equal(t, request.Frame(), helper.Page.mainFrame)
		require.Equal(t, "about:blank", request.Frame().URL())
		route.Continue()
		intercepted <- true
	})
	response, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.True(t, response.Ok())
	<-intercepted
}
