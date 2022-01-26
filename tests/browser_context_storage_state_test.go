package playwright_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestBrowserContextStorageStateShouldCaptureLocalStorage(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
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

func TestBrowserContextStorageStateSetLocalStorage(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	context, err := browser.NewContext(
		playwright.BrowserNewContextOptions{
			StorageState: &playwright.BrowserNewContextOptionsStorageState{
				Origins: []playwright.BrowserNewContextOptionsStorageStateOrigins{{
					Origin: playwright.String("https://www.example.com"),
					LocalStorage: []playwright.BrowserNewContextOptionsStorageStateOriginsLocalStorage{
						{
							Name:  playwright.String("name1"),
							Value: playwright.String("value1"),
						},
					},
				},
				},
			},
		},
	)
	require.NoError(t, err)
	defer context.Close()
	page, err := context.NewPage()
	require.NoError(t, err)
	defer page.Close()
	require.NoError(t, page.Route("**/*", func(route playwright.Route, request playwright.Request) {

		require.NoError(t, route.Fulfill(playwright.RouteFulfillOptions{
			Body: "<html></html>",
		}))
	}))
	_, err = page.Goto("https://www.example.com")
	require.NoError(t, err)
	localStorage, err := page.Evaluate("window.localStorage")
	require.NoError(t, err)
	require.Equal(t, map[string]interface{}{"name1": "value1"}, localStorage)
}

func TestBrowserContextStorageStateRoundTripThroughTheFile(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	page1, err := context.NewPage()
	require.NoError(t, err)
	defer page1.Close()
	require.NoError(t, page1.Route("**/*", func(route playwright.Route, request playwright.Request) {
		require.NoError(t, route.Fulfill(playwright.RouteFulfillOptions{
			Body: "<html></html>",
		}))
	}))
	_, err = page1.Goto("https://www.example.com")
	require.NoError(t, err)
	_, err = page1.Evaluate(`
	() => {
		localStorage["name1"] = "value1"
		document.cookie = "username=John Doe"
		return document.cookie
	}
	`)
	require.NoError(t, err)
	tempfile, err := ioutil.TempFile(os.TempDir(), "storage-state*.json")
	require.NoError(t, err)
	state, err := context.StorageState(tempfile.Name())
	require.NoError(t, err)
	stateWritten, err := ioutil.ReadFile(tempfile.Name())
	require.NoError(t, err)
	var storageState *playwright.StorageState
	err = json.Unmarshal(stateWritten, &storageState)
	require.NoError(t, err)
	require.Equal(t, state, storageState)

	context2, err := browser.NewContext(
		playwright.BrowserNewContextOptions{
			StorageStatePath: playwright.String(tempfile.Name()),
		})
	require.NoError(t, err)
	defer context2.Close()
	page2, err := context2.NewPage()
	require.NoError(t, err)
	defer page2.Close()
	require.NoError(t, page2.Route("**/*", func(route playwright.Route, request playwright.Request) {
		require.NoError(t, route.Fulfill(playwright.RouteFulfillOptions{
			Body: "<html></html>",
		}))
	}))
	_, err = page2.Goto("https://www.example.com")
	require.NoError(t, err)
	cookie, err := page2.Evaluate("document.cookie")
	require.NoError(t, err)
	require.Equal(t, "username=John Doe", cookie)
	localStorage, err := page2.Evaluate("window.localStorage")
	require.NoError(t, err)
	require.Equal(t, map[string]interface{}{"name1": "value1"}, localStorage)
}
