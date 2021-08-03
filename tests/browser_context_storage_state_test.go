package playwright_test

import (
	"runtime"
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestBrowserContextStorageStateShouldCaptureLocalStorage(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	if isWebKit && runtime.GOOS == "windows" {
		t.Skip()
	}
	page1, err := context.NewPage()
	require.NoError(t, err)
	require.NoError(t, page1.Route("**/*", func(route playwright.Route, request playwright.Request) {
		require.NoError(t, route.Fulfill(playwright.RouteFulfillOptions{
			Body: "<html></html>",
		}))
	}))
	_, err = page1.Goto("https://www.example.com")
	require.NoError(t, err)
	_, err = page1.Evaluate("localStorage['name1'] = 'value1'")
	require.NoError(t, err)
	_, err = page1.Goto("https://www.domain.com")
	require.NoError(t, err)
	_, err = page1.Evaluate("localStorage['name2'] = 'value2'")
	require.NoError(t, err)

	state, err := context.StorageState()
	require.NoError(t, err)
	origins := state.Origins
	require.Equal(t, 2, len(origins))
	require.Equal(t, origins[0], playwright.OriginsState{
		Origin: "https://www.example.com",
		LocalStorage: []playwright.LocalStorageEntry{
			{
				Name:  "name1",
				Value: "value1",
			},
		},
	})
	require.Equal(t, origins[1], playwright.OriginsState{
		Origin: "https://www.domain.com",
		LocalStorage: []playwright.LocalStorageEntry{
			{
				Name:  "name2",
				Value: "value2",
			},
		},
	})
}
