package playwright_test

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/playwright-community/playwright-go"
)

func routeIframe(t *testing.T, page playwright.Page) {
	t.Helper()

	err := page.Route("**/empty.html", func(route playwright.Route) {
		err := route.Fulfill(playwright.RouteFulfillOptions{
			Body:        `<iframe id="frame1" src="iframe.html"></iframe>`,
			ContentType: playwright.String("text/html"),
		})
		require.NoError(t, err)
	})
	require.NoError(t, err)

	err = page.Route("**/iframe.html", func(route playwright.Route) {
		err = route.Fulfill(playwright.RouteFulfillOptions{
			Body: `
	        <html>
	          <div>
	            <button>Hello iframe</button>
	            <iframe id="frame2" src="iframe-2.html"></iframe>
	          </div>
	          <span>1</span>
	          <span>2</span>
	        </html>`,
			ContentType: playwright.String("text/html"),
		})
		require.NoError(t, err)
	})
	require.NoError(t, err)

	err = page.Route("**/iframe-2.html", func(route playwright.Route) {
		err = route.Fulfill(playwright.RouteFulfillOptions{
			Body:        "<html><button>Hello nested iframe</button></html>",
			ContentType: playwright.String("text/html"),
		})
		require.NoError(t, err)
	})
	require.NoError(t, err)
}

func routeAmbiguous(t *testing.T, page playwright.Page) {
	t.Helper()

	err := page.Route("**/empty.html", func(route playwright.Route) {
		err := route.Fulfill(playwright.RouteFulfillOptions{
			Body: `<iframe src="iframe-1.html"></iframe>
             <iframe src="iframe-2.html"></iframe>
             <iframe src="iframe-3.html"></iframe>`,
			ContentType: playwright.String("text/html"),
		})
		require.NoError(t, err)
	})
	require.NoError(t, err)

	err = page.Route("**/iframe-*", func(route playwright.Route) {
		// const path = new URL(route.request().url()).pathname.slice(1);
		u, err := url.Parse(route.Request().URL())
		require.NoError(t, err)
		path := strings.TrimLeft(u.Path, "/")
		err = route.Fulfill(playwright.RouteFulfillOptions{
			Body:        fmt.Sprintf("<html><button>Hello from %s</button></html>", path),
			ContentType: playwright.String("text/html"),
		})
		require.NoError(t, err)
	})
	require.NoError(t, err)
}

func TestFrameLocatorFirst(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	routeAmbiguous(t, page)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	body, err := page.Locator("body")
	require.NoError(t, err)
	button, err := body.FrameLocator("iframe").First().Locator("button")
	require.NoError(t, err)
	innerText, err := button.InnerText()
	require.NoError(t, err)
	require.Equal(t, "Hello from iframe-1.html", innerText)
}

func TestFrameLocatorNth(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	routeAmbiguous(t, page)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	body, err := page.Locator("body")
	require.NoError(t, err)
	button, err := body.FrameLocator("iframe").Nth(1).Locator("button")
	require.NoError(t, err)
	innerText, err := button.InnerText()
	require.NoError(t, err)
	require.Equal(t, "Hello from iframe-2.html", innerText)
}

func TestFrameLocatorLast(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	routeAmbiguous(t, page)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	body, err := page.Locator("body")
	require.NoError(t, err)
	button, err := body.FrameLocator("iframe").Last().Locator("button")
	require.NoError(t, err)
	innerText, err := button.InnerText()
	require.NoError(t, err)
	require.Equal(t, "Hello from iframe-3.html", innerText)
}

func TestFrameLocatorLocator(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	routeIframe(t, page)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	body, err := page.Locator("body")
	require.NoError(t, err)
	frame1 := body.FrameLocator("#frame1")
	span, err := frame1.Locator("span")
	require.NoError(t, err)
	innerText, err := span.InnerText()
	require.NoError(t, err)
	require.Equal(t, "1", innerText)
}

func TestFrameLocatorFrameLocator(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	routeIframe(t, page)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	body, err := page.Locator("body")
	require.NoError(t, err)
	frame1 := body.FrameLocator("#frame1")
	frame2 := frame1.FrameLocator("#frame2")
	button, err := frame2.Locator("button")
	require.NoError(t, err)
	innerText, err := button.InnerText()
	require.NoError(t, err)
	require.Equal(t, "Hello nested iframe", innerText)
}
