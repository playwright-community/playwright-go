package playwright_test

import (
	"os"
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

// TestBrowserContextCloseRace tests for a data race between Close() and route handlers.
// This reproduces a race condition where Close() writes to closeWasCalled while
// route handler goroutines read it during page navigation.
//
// See: https://github.com/playwright-community/playwright-go/issues/XXX
func TestBrowserContextCloseRace(t *testing.T) {
	// Create a minimal HAR file
	harContent := `{
  "log": {
    "version": "1.2",
    "creator": {"name": "test", "version": "1.0"},
    "entries": [
      {
        "request": {
          "method": "GET",
          "url": "https://example.com/",
          "httpVersion": "HTTP/2.0",
          "headers": [],
          "queryString": [],
          "cookies": [],
          "headersSize": -1,
          "bodySize": 0
        },
        "response": {
          "status": 200,
          "statusText": "OK",
          "httpVersion": "HTTP/2.0",
          "headers": [{"name": "Content-Type", "value": "text/html"}],
          "cookies": [],
          "content": {
            "size": 13,
            "mimeType": "text/html",
            "text": "Hello, World!"
          },
          "redirectURL": "",
          "headersSize": -1,
          "bodySize": 13
        },
        "cache": {},
        "timings": {"send": 0, "wait": 0, "receive": 0}
      }
    ]
  }
}`

	harFile, err := os.CreateTemp("", "test-*.har")
	require.NoError(t, err)
	defer os.Remove(harFile.Name())

	_, err = harFile.WriteString(harContent)
	require.NoError(t, err)
	harFile.Close()

	// Create a new context for this test (don't use BeforeEach)
	testContext, err := browser.NewContext()
	require.NoError(t, err)
	defer testContext.Close()

	// Set up HAR replay - registers internal route handlers
	err = testContext.RouteFromHAR(harFile.Name(), playwright.BrowserContextRouteFromHAROptions{
		NotFound: playwright.HarNotFoundAbort,
	})
	require.NoError(t, err)

	// Add custom route handler
	err = testContext.Route("**/version.json*", func(route playwright.Route) {
		time.Sleep(5 * time.Millisecond) // Increase race window
		_ = route.Fulfill(playwright.RouteFulfillOptions{
			Status:      playwright.Int(200),
			ContentType: playwright.String("application/json"),
			Body:        playwright.String(`{"version": "1.0"}`),
		})
	})
	require.NoError(t, err)

	testPage, err := testContext.NewPage()
	require.NoError(t, err)

	// Start navigation in background
	done := make(chan error, 1)
	go func() {
		_, err := testPage.Goto("https://example.com/")
		done <- err
	}()

	// Give route handlers time to start processing
	time.Sleep(20 * time.Millisecond)

	// Close context while route handlers are actively running
	// This triggers the race between Close() and the route handler goroutines
	// Without proper synchronization, this will be detected by -race flag
	err = testContext.Close()
	require.NoError(t, err)

	// Wait for navigation to complete
	<-done
}
