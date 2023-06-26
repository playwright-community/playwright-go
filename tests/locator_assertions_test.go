package playwright_test

import (
	"regexp"
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestLocatorAssertionsToBeChecked(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`
	<input id='checkbox1' type='checkbox' checked>
	<input id='checkbox2' type='checkbox'>
	`)
	require.NoError(t, err)

	locator, err := page.Locator("#checkbox1")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToBeChecked())
	require.Error(t, expect.Locator(locator).NotToBeChecked())
	require.Error(t, expect.Locator(locator).ToBeChecked(playwright.LocatorAssertionsToBeCheckedOptions{
		Checked: playwright.Bool(false),
	}))

	locator, err = page.Locator("#checkbox2")
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator).ToBeChecked())
	require.NoError(t, expect.Locator(locator).NotToBeChecked())
}

func TestLocatorAssertionsToBeDisabled(t *testing.T) {
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

	locator, err := page.Locator(":text(\"button1\")")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToBeDisabled())
	require.Error(t, expect.Locator(locator).NotToBeDisabled())

	locator, err = page.Locator(":text(\"button2\")")
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator).ToBeDisabled())
	require.NoError(t, expect.Locator(locator).NotToBeDisabled())

	locator, err = page.Locator("div")
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator).ToBeDisabled())
	require.NoError(t, expect.Locator(locator).NotToBeDisabled())
}

func TestLocatorAssertionsToBeEditable(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`
	<input id=input1>
	<input id=input2 disabled>
	<textarea></textarea>
	`)
	require.NoError(t, err)

	locator, err := page.Locator("#input1")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToBeEditable())
	require.Error(t, expect.Locator(locator).NotToBeEditable())

	locator, err = page.Locator("#input2")
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator).ToBeEditable())
	require.NoError(t, expect.Locator(locator).NotToBeEditable())

	locator, err = page.Locator("textarea")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToBeEditable())
	require.Error(t, expect.Locator(locator).NotToBeEditable())
}

func TestLocatorAssertionsToBeEmpty(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`
	<textarea id="textarea1"></textarea>
	<textarea id="textarea2">test</textarea>
	<div id="div1"></div>
	<div id="div2">test</div>
	`)
	require.NoError(t, err)

	locator, err := page.Locator("#textarea1")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToBeEmpty())
	require.Error(t, expect.Locator(locator).NotToBeEmpty())

	locator, err = page.Locator("#textarea2")
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator).ToBeEmpty())
	require.NoError(t, expect.Locator(locator).NotToBeEmpty())

	locator, err = page.Locator("#div1")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToBeEmpty())
	require.Error(t, expect.Locator(locator).NotToBeEmpty())

	locator, err = page.Locator("#div2")
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator).ToBeEmpty())
	require.NoError(t, expect.Locator(locator).NotToBeEmpty())
}

func TestLocatorAssertionsToBeEnabled(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`
	<button>button1</button>
	<button disabled>button2</button>
	<div>div</div>
	`)
	require.NoError(t, err)

	locator, err := page.Locator(`:text("button1")`)
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToBeEnabled())
	require.Error(t, expect.Locator(locator).NotToBeEnabled())

	locator, err = page.Locator(`:text("button2")`)
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator).ToBeEnabled())
	require.NoError(t, expect.Locator(locator).NotToBeEnabled())

	locator, err = page.Locator("div")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToBeEnabled())
	require.Error(t, expect.Locator(locator).NotToBeEnabled())
}

func TestLocatorAssertionsToBeFocused(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`
	<input id=input1>
	<input id=input2>
	`)
	require.NoError(t, err)

	locator, err := page.Locator("#input1")
	require.NoError(t, err)
	require.NoError(t, locator.Focus())
	require.NoError(t, expect.Locator(locator).ToBeFocused())
	require.Error(t, expect.Locator(locator).NotToBeFocused())

	locator, err = page.Locator("#input2")
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator).ToBeFocused())
	require.NoError(t, expect.Locator(locator).NotToBeFocused())
}

func TestLocatorAssertionsToBeHidden(t *testing.T) {
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
	</details>
	`)
	require.NoError(t, err)

	locator, err := page.Locator("summary")
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator).ToBeHidden())
	require.NoError(t, expect.Locator(locator).NotToBeHidden())

	locator, err = page.Locator("ul")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToBeHidden())
	require.Error(t, expect.Locator(locator).NotToBeHidden())
}

func TestLocatorAssertionsToBeVisible(t *testing.T) {
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
	</details>
	`)
	require.NoError(t, err)

	locator, err := page.Locator("summary")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToBeVisible())
	require.Error(t, expect.Locator(locator).NotToBeVisible())

	locator, err = page.Locator("ul")
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator).ToBeVisible())
	require.NoError(t, expect.Locator(locator).NotToBeVisible())
}

func TestLocatorAssertionsToContainText(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`<div><span style="display: none">foo</span>test1</div>`)
	require.NoError(t, err)

	locator, err := page.Locator("div")
	require.NoError(t, err)

	require.NoError(t, expect.Locator(locator).ToContainText("foo"))
	require.NoError(t, expect.Locator(locator).NotToContainText("foo", playwright.LocatorAssertionsToContainTextOptions{
		UseInnerText: playwright.Bool(true),
	}))
	require.NoError(t, expect.Locator(locator).ToContainText([]string{"test"}))
	require.NoError(t, expect.Locator(locator).ToContainText(regexp.MustCompile(`test\d+`)))
	require.NoError(t, expect.Locator(locator).ToContainText([]*regexp.Regexp{regexp.MustCompile("test")}))
	require.Error(t, expect.Locator(locator).ToContainText("invalid"))
	require.Error(t, expect.Locator(locator).NotToContainText("test"))
}

func TestLocatorAssertionsToHaveAttribute(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`
	<input id="input1" type="text">
	<input id="input2" type="number">
	`)
	require.NoError(t, err)

	locator, err := page.Locator("#input1")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToHaveAttribute("type", "text"))
	require.NoError(t, expect.Locator(locator).ToHaveAttribute("type", regexp.MustCompile("text")))
	require.Error(t, expect.Locator(locator).NotToHaveAttribute("type", "text"))
	require.Error(t, expect.Locator(locator).NotToHaveAttribute("type", regexp.MustCompile("text")))

	locator, err = page.Locator("#input2")
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator).ToHaveAttribute("type", "text"))
	require.Error(t, expect.Locator(locator).ToHaveAttribute("type", regexp.MustCompile("text")))
	require.NoError(t, expect.Locator(locator).NotToHaveAttribute("type", "text"))
	require.NoError(t, expect.Locator(locator).NotToHaveAttribute("type", regexp.MustCompile("text")))
}

func TestLocatorAssertionsToHaveClass(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`
	<div class="test1">test1</div>
	<div class="test2">test2</div>
	`)
	require.NoError(t, err)

	locator, err := page.Locator(".test1")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToHaveClass("test1"))
	require.NoError(t, expect.Locator(locator).ToHaveClass([]string{"test1"}))
	require.NoError(t, expect.Locator(locator).ToHaveClass(regexp.MustCompile("test.{1}")))
	require.NoError(t, expect.Locator(locator).ToHaveClass([]*regexp.Regexp{regexp.MustCompile(`test\d+`)}))
	require.Error(t, expect.Locator(locator).NotToHaveClass("test1"))

	locator, err = page.Locator(".test2")
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator).ToHaveClass("test1"))
	require.Error(t, expect.Locator(locator).ToHaveClass([]string{"test1"}))
	require.Error(t, expect.Locator(locator).ToHaveClass(regexp.MustCompile(`test\d{2}`)))
	require.Error(t, expect.Locator(locator).ToHaveClass([]*regexp.Regexp{regexp.MustCompile(`test123`)}))
	require.NoError(t, expect.Locator(locator).NotToHaveClass("test1"))
}

func TestLocatorAssertionsToHaveCount(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`
	<button>button1</button>
	<button disabled>button2</button>
	`)
	require.NoError(t, err)

	locator, err := page.Locator("button")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToHaveCount(2))
	require.Error(t, expect.Locator(locator).NotToHaveCount(2))
}

func TestLocatorAssertionsToHaveCSS(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`
	<button id="button1" style="display: flex">button1</button>
	<button id="button2">button2</button>
	`)
	require.NoError(t, err)

	locator, err := page.Locator("#button1")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToHaveCSS("display", "flex"))
	require.Error(t, expect.Locator(locator).NotToHaveCSS("display", "flex"))

	locator, err = page.Locator("#button2")
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator).ToHaveCSS("display", "flex"))
	require.NoError(t, expect.Locator(locator).NotToHaveCSS("display", "flex"))
}

func TestLocatorAssertionsToHaveId(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`
	<button id="button1">button1</button>
	<div>div</div>
	`)
	require.NoError(t, err)

	locator, err := page.Locator("button")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToHaveId("button1"))
	require.Error(t, expect.Locator(locator).NotToHaveId("button1"))

	locator, err = page.Locator("div")
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator).ToHaveId("button1"))
	require.NoError(t, expect.Locator(locator).NotToHaveId("button1"))
}

func TestLocatorAssertionsToHaveJSProperty(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent("<div></div>"))
	_, err = page.EvalOnSelector("div", "e => e.foo = true")
	require.NoError(t, err)

	locator, err := page.Locator("div")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToHaveJSProperty("foo", true))
	require.Error(t, expect.Locator(locator).NotToHaveJSProperty("foo", true))
}

func TestLocatorAssertionsToHaveText(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`<div><span style="display: none">foo</span>test</div>`)
	require.NoError(t, err)

	locator, err := page.Locator("div")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToHaveText("footest"))
	require.NoError(t, expect.Locator(locator).ToHaveText("Test",
		playwright.LocatorAssertionsToHaveTextOptions{
			UseInnerText: playwright.Bool(true),
			IgnoreCase:   playwright.Bool(true),
		}))
	require.NoError(t, expect.Locator(locator).ToHaveText([]string{"footest"}))
	require.NoError(t, expect.Locator(locator).ToHaveText(regexp.MustCompile("foo.*")))
	require.NoError(t, expect.Locator(locator).ToHaveText([]*regexp.Regexp{regexp.MustCompile("foo.*")}))
	require.Error(t, expect.Locator(locator).ToHaveText("invalid"))
	require.Error(t, expect.Locator(locator).NotToHaveText("footest"))
}

func TestLocatorAssertionsToHaveValue(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`<input type="text" value="test">`)
	require.NoError(t, err)

	locator, err := page.Locator("input")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToHaveValue("test"))
	require.NoError(t, expect.Locator(locator).ToHaveValue(regexp.MustCompile("te.*")))
	require.Error(t, expect.Locator(locator).ToHaveValue("invalid"))
	require.Error(t, expect.Locator(locator).NotToHaveValue("test"))
}

func TestLocatorAssertionsToHaveValues(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	err = page.SetContent(`<select multiple>
	<option value="R">Red</option>
	<option value="G">Green</option>
	<option value="B">Blue</option>
</select>`)
	require.NoError(t, err)

	locator, err := page.Locator("select")
	require.NoError(t, err)
	_, err = locator.SelectOption(playwright.SelectOptionValues{
		Values: &[]string{"R", "G"},
	})
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToHaveValues([]interface{}{"R", "G"}))
	require.NoError(t, expect.Locator(locator).NotToHaveValues([]interface{}{"G", "B"}))
}
