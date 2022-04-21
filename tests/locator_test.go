package playwright_test

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestLocatorAllInnerTexts(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<div>A</div><div>B</div><div>C</div>`))

	locator, err := page.Locator("div")
	require.NoError(t, err)
	innerHTML, err := locator.AllInnerTexts()
	require.NoError(t, err)
	require.ElementsMatch(t, []string{"A", "B", "C"}, innerHTML)
}

func TestLocatorAllTextContents(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<div>A</div><div>B</div><div>C</div>`))

	locator, err := page.Locator("div")
	require.NoError(t, err)
	innerHTML, err := locator.AllTextContents()
	require.NoError(t, err)
	require.ElementsMatch(t, []string{"A", "B", "C"}, innerHTML)
}

func TestLocatorFill(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/dom.html")
	require.NoError(t, err)

	locator, err := page.Locator("#input")
	require.NoError(t, err)
	require.NoError(t, locator.Fill("input value"))
	result, err := locator.InputValue()
	require.NoError(t, err)
	require.Equal(t, "input value", result)
}

func TestLocatorGetAttribute(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/dom.html")
	require.NoError(t, err)

	locator, err := page.Locator("#outer")
	require.NoError(t, err)
	result, err := locator.GetAttribute("name")
	require.NoError(t, err)
	require.Equal(t, "value", result)
	result, err = locator.GetAttribute("foo")
	require.NoError(t, err)
	require.Empty(t, result)
}

func TestLocatorInnerHTML(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/dom.html")
	require.NoError(t, err)

	locator, err := page.Locator("#outer")
	require.NoError(t, err)
	result, err := locator.InnerHTML()
	require.NoError(t, err)
	require.Equal(t, "<div id=\"inner\">Text,\nmore text</div>", result)
}

func TestLocatorInnerText(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/dom.html")
	require.NoError(t, err)

	locator, err := page.Locator("#inner")
	require.NoError(t, err)
	result, err := locator.InnerHTML()
	require.NoError(t, err)
	require.Equal(t, "Text,\nmore text", result)
}

func TestLocatorInputValue(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/dom.html")
	require.NoError(t, err)
	require.NoError(t, page.Fill("#input", "input value"))

	locator, err := page.Locator("#input")
	require.NoError(t, err)
	result, err := locator.InputValue()
	require.NoError(t, err)
	require.Equal(t, "input value", result)
}

func TestLocatorIsChecked(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent("<input type='checkbox' checked><div>Not a checkbox</div>"))

	locator, err := page.Locator("input")
	require.NoError(t, err)
	result, err := locator.IsChecked()
	require.NoError(t, err)
	require.True(t, result)
}

func TestLocatorIsDisabled(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`
	<button disabled>button1</button>
	<button>button2</button>
	<div>div</div>
	`)
	require.NoError(t, err)

	locator, err := page.Locator("div")
	require.NoError(t, err)
	result, err := locator.IsDisabled()
	require.NoError(t, err)
	require.False(t, result)

	locator, err = page.Locator(":text(\"button1\")")
	require.NoError(t, err)
	result, err = locator.IsDisabled()
	require.NoError(t, err)
	require.True(t, result)

	locator, err = page.Locator(":text(\"button2\")")
	require.NoError(t, err)
	result, err = locator.IsDisabled()
	require.NoError(t, err)
	require.False(t, result)
}

func TestLocatorIsEditable(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`<input id=input1 disabled><textarea></textarea><input id=input2>
	`)
	require.NoError(t, err)

	locator, err := page.Locator("#input1")
	require.NoError(t, err)
	result, err := locator.IsEditable()
	require.NoError(t, err)
	require.False(t, result)

	locator, err = page.Locator("#input2")
	require.NoError(t, err)
	result, err = locator.IsEditable()
	require.NoError(t, err)
	require.True(t, result)

	locator, err = page.Locator("textarea")
	require.NoError(t, err)
	result, err = locator.IsEditable()
	require.NoError(t, err)
	require.True(t, result)
}

func TestLocatorIsEnabled(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`
	<button disabled>button1</button>
	<button>button2</button>
	<div>div</div>
	`)
	require.NoError(t, err)

	locator, err := page.Locator("div")
	require.NoError(t, err)
	result, err := locator.IsEnabled()
	require.NoError(t, err)
	require.True(t, result)

	locator, err = page.Locator(":text(\"button1\")")
	require.NoError(t, err)
	result, err = locator.IsEnabled()
	require.NoError(t, err)
	require.False(t, result)

	locator, err = page.Locator(":text(\"button2\")")
	require.NoError(t, err)
	result, err = locator.IsEnabled()
	require.NoError(t, err)
	require.True(t, result)
}

func TestLocatorIsHidden(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`
	<details>
		<summary>click to open</summary>
		<ul>
			<li>hidden item 1</li>
			<li>hidden item 2</li>
			<li>hidden item 3</li>
		</ul>
	</details>`)
	require.NoError(t, err)

	locator, err := page.Locator("ul")
	require.NoError(t, err)
	result, err := locator.IsHidden()
	require.NoError(t, err)
	require.True(t, result)

	locator, err = page.Locator("summary")
	require.NoError(t, err)
	result, err = locator.IsHidden()
	require.NoError(t, err)
	require.False(t, result)
}

func TestLocatorIsVisible(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`
	<details>
		<summary>click to open</summary>
		<ul>
			<li>hidden item 1</li>
			<li>hidden item 2</li>
			<li>hidden item 3</li>
		</ul>
	</details>`)
	require.NoError(t, err)

	locator, err := page.Locator("ul")
	require.NoError(t, err)
	result, err := locator.IsVisible()
	require.NoError(t, err)
	require.False(t, result)

	locator, err = page.Locator("summary")
	require.NoError(t, err)
	result, err = locator.IsVisible()
	require.NoError(t, err)
	require.True(t, result)
}

func TestLocatorSelectOption(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/dom.html")
	require.NoError(t, err)

	locator, err := page.Locator("#select")
	require.NoError(t, err)
	values := []string{"foo"}
	result, err := locator.SelectOption(playwright.SelectOptionValues{Values: &values})
	require.NoError(t, err)
	require.ElementsMatch(t, []string{"foo"}, result)
}

func TestLocatorTextContent(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/dom.html")
	require.NoError(t, err)

	locator, err := page.Locator("#inner")
	require.NoError(t, err)
	result, err := locator.TextContent()
	require.NoError(t, err)
	require.Equal(t, "Text,\nmore text", result)
}
