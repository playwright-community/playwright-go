package playwright_test

import (
	"regexp"
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestPageAssertionsToHaveTitle(t *testing.T) {
	BeforeEach(t)

	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<title>new title</title>`))

	require.NoError(t, expect.Page(page).ToHaveTitle("new title"))
	require.NoError(t, expect.Page(page).ToHaveTitle(regexp.MustCompile("(?i)new title")))
	require.NoError(t, expect.Page(page).Not().ToHaveTitle("not the current title", playwright.PageAssertionsToHaveTitleOptions{
		Timeout: playwright.Float(750),
	}))

	_, err = page.Evaluate(`setTimeout(() => {
		document.title = 'great title';
	}, 300);
	`)
	require.NoError(t, err)
	require.NoError(t, expect.Page(page).ToHaveTitle("great title"))
	require.NoError(t, expect.Page(page).Not().ToHaveTitle("not the current title"))
}

func TestPageAssertionsToHaveURL(t *testing.T) {
	BeforeEach(t)

	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	require.NoError(t, expect.Page(page).ToHaveURL(server.EMPTY_PAGE))
	require.NoError(t, expect.Page(page).ToHaveURL(regexp.MustCompile(`.*/empty\.html`), playwright.PageAssertionsToHaveURLOptions{
		Timeout: playwright.Float(750),
	}))
	require.NoError(t, expect.Page(page).Not().ToHaveURL("https://playwright.dev"))
}

func TestPageAssertionsToHaveURLWithBaseURL(t *testing.T) {
	BeforeEach(t)

	page, err := browser.NewPage(playwright.BrowserNewPageOptions{
		BaseURL: &server.PREFIX,
	})
	require.NoError(t, err)
	_, err = page.Goto("/empty.html")
	require.NoError(t, err)
	require.NoError(t, expect.Page(page).ToHaveURL("/empty.html"))
	require.NoError(t, expect.Page(page).ToHaveURL(regexp.MustCompile(`.*/empty\.html`)))
	require.NoError(t, expect.Page(page).Not().ToHaveURL("https://playwright.dev"))
	require.NoError(t, page.Close())
}
