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

	text, err := page.GetByTestId("Hello").TextContent()
	require.NoError(t, err)
	require.Equal(t, "Hello world", text)

	text, err = page.MainFrame().GetByTestId("Hello").TextContent()
	require.NoError(t, err)
	require.Equal(t, "Hello world", text)

	text, err = page.Locator("div").GetByTestId("Hello").TextContent()
	require.NoError(t, err)
	require.Equal(t, "Hello world", text)
}

func TestGetByTestIdEscapeId(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`<div><div data-testid='He"llo'>Hello world</div></div>`))

	text, err := page.GetByTestId("He\"llo").TextContent()
	require.NoError(t, err)
	require.Equal(t, "Hello world", text)
	count, err := page.GetByTestId(regexp.MustCompile(`[Hh]e.llo`)).Count()
	require.NoError(t, err)
	require.Equal(t, 1, count)
}

func TestGetByText(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`<div><div>yo</div><div>ya</div><div>\nye  </div></div>`))
	require.NoError(t, expect.Locator(page.GetByText("yo")).ToHaveCount(1))
	require.NoError(t, expect.Locator(page.Locator("div").GetByText("yo")).ToHaveCount(1))
}

func TestGetByLabel(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`<div><label for=target>Name</label><input id=target type=text></div>`))

	require.NoError(t, expect.Locator(page.GetByLabel("Name")).ToHaveCount(1))
	require.NoError(t, expect.Locator(page.GetByLabel(regexp.MustCompile(`N?me`))).ToHaveCount(1))
	locator := page.Locator("div")
	require.NoError(t, expect.Locator(locator.GetByLabel("Name")).ToHaveCount(1))

	ret, err := locator.GetByLabel("Name").Evaluate("e => e.nodeName", nil)
	require.NoError(t, err)
	require.Equal(t, "INPUT", ret)
}

func TestGetByPlaceholder(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`
	<div>
    <input placeholder="Hello">
    <input placeholder="Hello World">
  </div>`))

	require.NoError(t, expect.Locator(page.GetByPlaceholder("hello")).ToHaveCount(2))
	locator := page.Locator("div").GetByPlaceholder("Hello", playwright.LocatorGetByPlaceholderOptions{
		Exact: playwright.Bool(true),
	})
	require.NoError(t, expect.Locator(locator).ToHaveCount(1))
}

func TestGetByAltText(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`
	<div>
    <input alt="Hello">
    <input alt="Hello World">
  </div>`))
	require.NoError(t, expect.Locator(page.GetByAltText("hello")).ToHaveCount(2))
	require.NoError(t, expect.Locator(page.GetByAltText(regexp.MustCompile(`Hello.+d`))).ToHaveCount(1))
}

func TestGetByTitle(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`
	<div>
    <input title="Hello">
    <input title="Hello World">
  </div>`))
	require.NoError(t, expect.Locator(page.GetByTitle("hello")).ToHaveCount(2))
}

func TestGetByRole(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`
	<button>Hello</button>
	<button>Hel"lo</button>
	<div role="dialog">I am a dialog</div>
	`))

	count, err := page.GetByRole("button", playwright.LocatorGetByRoleOptions{
		Name: "hello",
	}).Count()
	require.NoError(t, err)
	require.Equal(t, 1, count)

	count, err = page.GetByRole("button", playwright.LocatorGetByRoleOptions{
		Name: "Hel\"lo",
	}).Count()
	require.NoError(t, err)
	require.Equal(t, 1, count)

	count, err = page.GetByRole("button", playwright.LocatorGetByRoleOptions{
		Name: regexp.MustCompile(`(?i)he`),
	}).Count()
	require.NoError(t, err)
	require.Equal(t, 2, count)

	count, err = page.GetByRole("dialog").Count()
	require.NoError(t, err)
	require.Equal(t, 1, count)
}
