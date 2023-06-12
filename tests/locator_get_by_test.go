package playwright_test

import (
	"regexp"
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestGetByTestId(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`<div><div data-testid='Hello'>Hello world</div></div>`))
	locator, err := page.GetByTestId("Hello")
	require.NoError(t, err)
	text, err := locator.TextContent()
	require.NoError(t, err)
	require.Equal(t, "Hello world", text)

	locator, err = page.MainFrame().GetByTestId("Hello")
	require.NoError(t, err)
	text, err = locator.TextContent()
	require.NoError(t, err)
	require.Equal(t, "Hello world", text)

	locator, err = page.Locator("div")
	require.NoError(t, err)
	locator, err = locator.GetByTestId("Hello")
	require.NoError(t, err)
	text, err = locator.TextContent()
	require.NoError(t, err)
	require.Equal(t, "Hello world", text)
}

func TestGetByTestIdEscapeId(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`<div><div data-testid='He"llo'>Hello world</div></div>`))
	locator, err := page.GetByTestId("He\"llo")
	require.NoError(t, err)
	text, err := locator.TextContent()
	require.NoError(t, err)
	require.Equal(t, "Hello world", text)
}

func TestGetByRole(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`
	<button>Hello</button>
	<button>Hel"lo</button>
	<div role="dialog">I am a dialog</div>
	`))
	locator, err := page.GetByRole("button", playwright.LocatorGetByRoleOptions{
		Name: "hello",
	})
	require.NoError(t, err)
	count, err := locator.Count()
	require.NoError(t, err)
	require.Equal(t, 1, count)

	locator, err = page.GetByRole("button", playwright.LocatorGetByRoleOptions{
		Name: "Hel\"lo",
	})
	require.NoError(t, err)
	require.NoError(t, err)
	count, err = locator.Count()
	require.NoError(t, err)
	require.Equal(t, 1, count)

	locator, err = page.GetByRole("button", playwright.LocatorGetByRoleOptions{
		Name: regexp.MustCompile(`(?i)he`),
	})
	require.NoError(t, err)
	require.NoError(t, err)
	count, err = locator.Count()
	require.NoError(t, err)
	require.Equal(t, 2, count)

	locator, err = page.GetByRole("dialog")
	require.NoError(t, err)
	require.NoError(t, err)
	count, err = locator.Count()
	require.NoError(t, err)
	require.Equal(t, 1, count)
}
