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
	err := page.SetContent(`
	<input id='checkbox1' type='checkbox' checked>
	<input id='checkbox2' type='checkbox'>
	`)
	require.NoError(t, err)

	locator := page.Locator("#checkbox1")
	require.NoError(t, locator.Err())
	require.NoError(t, expect.Locator(locator).ToBeChecked())
	require.Error(t, expect.Locator(locator).NotToBeChecked())
	require.Error(t, expect.Locator(locator).ToBeChecked(playwright.LocatorAssertionsToBeCheckedOptions{
		Checked: playwright.Bool(false),
	}))

	require.Error(t, expect.Locator(page.Locator("#checkbox2")).ToBeChecked())
	require.NoError(t, expect.Locator(page.Locator("#checkbox2")).NotToBeChecked())
}

func TestLocatorAssertionsToBeDisabled(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
	<button disabled>button1</button>
	<button>button2</button>
	<div>div</div>
	`)
	require.NoError(t, err)

	require.NoError(t, expect.Locator(page.Locator(":text(\"button1\")")).ToBeDisabled())
	require.Error(t, expect.Locator(page.Locator(":text(\"button1\")")).NotToBeDisabled())

	require.Error(t, expect.Locator(page.Locator(":text(\"button2\")")).ToBeDisabled())
	require.NoError(t, expect.Locator(page.Locator(":text(\"button2\")")).NotToBeDisabled())

	require.Error(t, expect.Locator(page.Locator("div")).ToBeDisabled())
	require.NoError(t, expect.Locator(page.Locator("div")).NotToBeDisabled())
}

func TestLocatorAssertionsToBeEditable(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
	<input id=input1>
	<input id=input2 disabled>
	<textarea></textarea>
	`)
	require.NoError(t, err)

	require.NoError(t, expect.Locator(page.Locator("#input1")).ToBeEditable())
	require.Error(t, expect.Locator(page.Locator("#input1")).NotToBeEditable())

	require.Error(t, expect.Locator(page.Locator("#input2")).ToBeEditable())
	require.NoError(t, expect.Locator(page.Locator("#input2")).NotToBeEditable())

	require.NoError(t, expect.Locator(page.Locator("textarea")).ToBeEditable())
	require.Error(t, expect.Locator(page.Locator("textarea")).NotToBeEditable())
}

func TestLocatorAssertionsToBeEmpty(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
	<textarea id="textarea1"></textarea>
	<textarea id="textarea2">test</textarea>
	<div id="div1"></div>
	<div id="div2">test</div>
	`)
	require.NoError(t, err)

	require.NoError(t, expect.Locator(page.Locator("#textarea1")).ToBeEmpty())
	require.Error(t, expect.Locator(page.Locator("#textarea1")).NotToBeEmpty())

	require.Error(t, expect.Locator(page.Locator("#textarea2")).ToBeEmpty())
	require.NoError(t, expect.Locator(page.Locator("#textarea2")).NotToBeEmpty())

	require.NoError(t, expect.Locator(page.Locator("#div1")).ToBeEmpty())
	require.Error(t, expect.Locator(page.Locator("#div1")).NotToBeEmpty())

	require.Error(t, expect.Locator(page.Locator("#div2")).ToBeEmpty())
	require.NoError(t, expect.Locator(page.Locator("#div2")).NotToBeEmpty())
}

func TestLocatorAssertionsToBeEnabled(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
	<button>button1</button>
	<button disabled>button2</button>
	<div>div</div>
	`)
	require.NoError(t, err)

	require.NoError(t, expect.Locator(page.Locator(`:text("button1")`)).ToBeEnabled())
	require.Error(t, expect.Locator(page.Locator(`:text("button1")`)).NotToBeEnabled())

	require.Error(t, expect.Locator(page.Locator(`:text("button2")`)).ToBeEnabled())
	require.NoError(t, expect.Locator(page.Locator(`:text("button2")`)).NotToBeEnabled())

	require.NoError(t, expect.Locator(page.Locator("div")).ToBeEnabled())
	require.Error(t, expect.Locator(page.Locator("div")).NotToBeEnabled())
}

func TestLocatorAssertionsToBeFocused(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
	<input id=input1>
	<input id=input2>
	`)
	require.NoError(t, err)

	locator := page.Locator("#input1")
	require.NoError(t, locator.Focus())
	require.NoError(t, expect.Locator(locator).ToBeFocused())
	require.Error(t, expect.Locator(locator).NotToBeFocused())

	locator2 := page.Locator("#input2")
	require.Error(t, expect.Locator(locator2).ToBeFocused())
	require.NoError(t, expect.Locator(locator2).NotToBeFocused())
}

func TestLocatorAssertionsToBeHidden(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
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

	locator := page.Locator("summary")
	require.Error(t, expect.Locator(locator).ToBeHidden())
	require.NoError(t, expect.Locator(locator).NotToBeHidden())

	locator2 := page.Locator("ul")
	require.NoError(t, expect.Locator(locator2).ToBeHidden())
	require.Error(t, expect.Locator(locator2).NotToBeHidden())
}

func TestLocatorAssertionsToBeVisible(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
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

	locator := page.Locator("summary")
	require.NoError(t, expect.Locator(locator).ToBeVisible())
	require.Error(t, expect.Locator(locator).NotToBeVisible())

	locator2 := page.Locator("ul")
	require.NoError(t, err)
	require.Error(t, expect.Locator(locator2).ToBeVisible())
	require.NoError(t, expect.Locator(locator2).NotToBeVisible())
}

func TestLocatorAssertionsToContainText(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`<div><span style="display: none">foo</span>test1</div>`)
	require.NoError(t, err)

	locator := page.Locator("div")

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
	err := page.SetContent(`
	<input id="input1" type="text">
	<input id="input2" type="number">
	`)
	require.NoError(t, err)

	input1 := page.Locator("#input1")
	require.NoError(t, expect.Locator(input1).ToHaveAttribute("type", "text"))
	require.NoError(t, expect.Locator(input1).ToHaveAttribute("type", regexp.MustCompile("text")))
	require.Error(t, expect.Locator(input1).NotToHaveAttribute("type", "text"))
	require.Error(t, expect.Locator(input1).NotToHaveAttribute("type", regexp.MustCompile("text")))

	input2 := page.Locator("#input2")
	require.Error(t, expect.Locator(input2).ToHaveAttribute("type", "text"))
	require.Error(t, expect.Locator(input2).ToHaveAttribute("type", regexp.MustCompile("text")))
	require.NoError(t, expect.Locator(input2).NotToHaveAttribute("type", "text"))
	require.NoError(t, expect.Locator(input2).NotToHaveAttribute("type", regexp.MustCompile("text")))
}

func TestLocatorAssertionsToHaveClass(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
	<div class="test1">test1</div>
	<div class="test2">test2</div>
	`)
	require.NoError(t, err)

	locator := page.Locator(".test1")
	require.NoError(t, locator.Err())
	require.NoError(t, expect.Locator(locator).ToHaveClass("test1"))
	require.NoError(t, expect.Locator(locator).ToHaveClass([]string{"test1"}))
	require.NoError(t, expect.Locator(locator).ToHaveClass(regexp.MustCompile("test.{1}")))
	require.NoError(t, expect.Locator(locator).ToHaveClass([]*regexp.Regexp{regexp.MustCompile(`test\d+`)}))
	require.Error(t, expect.Locator(locator).NotToHaveClass("test1"))

	locator2 := page.Locator(".test2")
	require.NoError(t, locator2.Err())
	require.Error(t, expect.Locator(locator2).ToHaveClass("test1"))
	require.Error(t, expect.Locator(locator2).ToHaveClass([]string{"test1"}))
	require.Error(t, expect.Locator(locator2).ToHaveClass(regexp.MustCompile(`test\d{2}`)))
	require.Error(t, expect.Locator(locator2).ToHaveClass([]*regexp.Regexp{regexp.MustCompile(`test123`)}))
	require.NoError(t, expect.Locator(locator2).NotToHaveClass("test1"))
}

func TestLocatorAssertionsToHaveCount(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
	<button>button1</button>
	<button disabled>button2</button>
	`)
	require.NoError(t, err)

	locator := page.Locator("button")
	require.NoError(t, locator.Err())
	require.NoError(t, expect.Locator(locator).ToHaveCount(2))
	require.Error(t, expect.Locator(locator).NotToHaveCount(2))
}

func TestLocatorAssertionsToHaveCSS(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
	<button id="button1" style="display: flex">button1</button>
	<button id="button2">button2</button>
	`)
	require.NoError(t, err)

	button1 := page.Locator("#button1")
	require.NoError(t, err)
	require.NoError(t, expect.Locator(button1).ToHaveCSS("display", "flex"))
	require.Error(t, expect.Locator(button1).NotToHaveCSS("display", "flex"))

	button2 := page.Locator("#button2")
	require.Error(t, expect.Locator(button2).ToHaveCSS("display", "flex"))
	require.NoError(t, expect.Locator(button2).NotToHaveCSS("display", "flex"))
}

func TestLocatorAssertionsToHaveId(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
	<button id="button1">button1</button>
	<div>div</div>
	`)
	require.NoError(t, err)

	button := page.Locator("button")
	require.NoError(t, expect.Locator(button).ToHaveId("button1"))
	require.Error(t, expect.Locator(button).NotToHaveId("button1"))

	div := page.Locator("div")
	require.Error(t, expect.Locator(div).ToHaveId("button1"))
	require.NoError(t, expect.Locator(div).NotToHaveId("button1"))
}

func TestLocatorAssertionsToHaveJSProperty(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent("<div></div>"))
	_, err = page.EvalOnSelector("div", "e => e.foo = true")
	require.NoError(t, err)

	locator := page.Locator("div")
	require.NoError(t, expect.Locator(locator).ToHaveJSProperty("foo", true))
	require.Error(t, expect.Locator(locator).NotToHaveJSProperty("foo", true))
}

func TestLocatorAssertionsToHaveText(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`<div><span style="display: none">foo</span>test</div>`)
	require.NoError(t, err)

	locator := page.Locator("div")
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
	err := page.SetContent(`<input type="text" value="test">`)
	require.NoError(t, err)

	locator := page.Locator("input")
	require.NoError(t, expect.Locator(locator).ToHaveValue("test"))
	require.NoError(t, expect.Locator(locator).ToHaveValue(regexp.MustCompile("te.*")))
	require.Error(t, expect.Locator(locator).ToHaveValue("invalid"))
	require.Error(t, expect.Locator(locator).NotToHaveValue("test"))
}

func TestLocatorAssertionsToHaveValues(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`<select multiple>
	<option value="R">Red</option>
	<option value="G">Green</option>
	<option value="B">Blue</option>
</select>`)
	require.NoError(t, err)

	locator := page.Locator("select")
	_, err = locator.SelectOption(playwright.SelectOptionValues{
		Values: &[]string{"R", "G"},
	})
	require.NoError(t, err)
	require.NoError(t, expect.Locator(locator).ToHaveValues([]interface{}{"R", "G"}))
	require.NoError(t, expect.Locator(locator).NotToHaveValues([]interface{}{"G", "B"}))
}

func TestToBeInViewportShouldRespectRatioOption(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
	<style>body, div, html { padding: 0; margin: 0; }</style>
      <div id=big style="height: 400vh;"></div>`)
	require.NoError(t, err)
	locator := page.Locator("div")
	require.NoError(t, expect.Locator(locator).ToBeInViewport())
	require.NoError(t, expect.Locator(locator).ToBeInViewport(playwright.LocatorAssertionsToBeInViewportOptions{
		Ratio: playwright.Float(0.1),
	}))
	require.NoError(t, expect.Locator(locator).ToBeInViewport(playwright.LocatorAssertionsToBeInViewportOptions{
		Ratio: playwright.Float(0.2),
	}))
	require.NoError(t, expect.Locator(locator).ToBeInViewport(playwright.LocatorAssertionsToBeInViewportOptions{
		Ratio: playwright.Float(0.25),
	}))
	// In this test, element's ratio is 0.25.
	require.NoError(t, expect.Locator(locator).NotToBeInViewport(playwright.LocatorAssertionsToBeInViewportOptions{
		Ratio: playwright.Float(0.26),
	}))
	require.NoError(t, expect.Locator(locator).NotToBeInViewport(playwright.LocatorAssertionsToBeInViewportOptions{
		Ratio: playwright.Float(0.3),
	}))
	require.NoError(t, expect.Locator(locator).NotToBeInViewport(playwright.LocatorAssertionsToBeInViewportOptions{
		Ratio: playwright.Float(0.7),
	}))
	require.NoError(t, expect.Locator(locator).Not().ToBeInViewport(playwright.LocatorAssertionsToBeInViewportOptions{
		Ratio: playwright.Float(0.8),
	}))
}

func TestLocatorShouldBeAttachedWithHiddenElement(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`<button style="display:none">hello</button>`)
	require.NoError(t, err)
	locator := page.Locator("button")
	require.NoError(t, expect.Locator(locator).ToBeAttached())
	require.NoError(t, expect.Locator(locator).Not().ToBeAttached(playwright.LocatorAssertionsToBeAttachedOptions{
		Attached: playwright.Bool(false),
	}))
	require.NoError(t, expect.Locator(page.Locator("input")).NotToBeAttached())
}
