package playwright_test

import (
	"fmt"
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

func TestLocatorLocatorHas(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	expText := " First item 1 First item 1A1"

	err = page.SetContent(`
	<section>
		<ul>
			<li>
				<div>
					<input class="r1" name="r1" type="checkbox"/>
					<span> First item 1</span>
					<span> First item 1A1</span>
				</div>
			</li>
			<li>
				<div>
					<input name="r2" type="checkbox"/>
					<span> Second item 1</span>
					<span> Second item 1<span>A1</span></span>
				</div>
			</li>
		</ul>
	</section>`)
	require.NoError(t, err)

	inputLocator, err := page.Locator("input[name='r1']")
	require.NoError(t, err)

	listLocator, err := page.Locator("ul", playwright.PageLocatorOptions{Has: inputLocator})
	require.NoError(t, err)

	spanLocator, err := page.Locator("span", playwright.PageLocatorOptions{HasText: "First item 1A"})
	require.NoError(t, err)

	targetLocator, err := listLocator.Locator("li div", playwright.LocatorLocatorOptions{Has: spanLocator})
	require.NoError(t, err)

	targetText, err := targetLocator.InnerText()
	require.NoError(t, err)
	require.Equal(t, expText, targetText)
}

func TestLocatorLocatorHasText(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	expText := "A1 B2"

	err = page.SetContent(`
	<section>
		<ul>
			<li>
				<div>
					<input name="r2" type="checkbox"/>
					<span> Second item a1</span>
					<span> Second item 1<span>A1 B2</span></span>
				</div>
			</li>
		</ul>
	</section>`)
	require.NoError(t, err)

	inputLocator, err := page.Locator("input[name='r2']")
	require.NoError(t, err)

	listLocator, err := page.Locator("ul", playwright.PageLocatorOptions{Has: inputLocator})
	require.NoError(t, err)

	wrongTargetLocator, err := listLocator.Locator("li div span", playwright.LocatorLocatorOptions{HasText: "A1"})
	require.NoError(t, err)

	count, err := wrongTargetLocator.Count()
	require.NoError(t, err)
	require.Equal(t, 3, count, "Locator count should be equal 3")

	targetParentLocator, err := listLocator.Locator("li div span", playwright.LocatorLocatorOptions{HasText: "1A1"})
	require.NoError(t, err)

	targetLocator, err := targetParentLocator.Locator("span")
	require.NoError(t, err)

	targetText, err := targetLocator.InnerText()
	require.NoError(t, err)
	require.Equal(t, expText, targetText)
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

func TestLocatorShouldFocusAndBlurButton(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/input/button.html")
	require.NoError(t, err)
	button, err := page.Locator("button")
	require.NoError(t, err)
	ret, err := button.Evaluate(`button => document.activeElement === button`, nil)
	require.NoError(t, err)
	require.False(t, ret.(bool))

	var (
		focused = false
		blurred = false
	)
	require.NoError(t, page.ExposeFunction("focusEvent", func(args ...interface{}) interface{} {
		focused = true
		return nil
	}))
	require.NoError(t, page.ExposeFunction("blurEvent", func(args ...interface{}) interface{} {
		blurred = true
		return nil
	}))
	_, err = button.Evaluate(`button => {
		button.addEventListener('focus', window['focusEvent']);
		button.addEventListener('blur', window['blurEvent']);
}`, nil)
	require.NoError(t, err)

	require.NoError(t, button.Focus())
	ret, err = button.Evaluate(`button => document.activeElement === button`, nil)
	require.NoError(t, err)
	require.True(t, ret.(bool))
	require.True(t, focused)
	require.False(t, blurred)

	require.NoError(t, button.Blur())
	ret, err = button.Evaluate(`button => document.activeElement === button`, nil)
	require.NoError(t, err)
	require.False(t, ret.(bool))
	require.True(t, focused)
	require.True(t, blurred)
}

func TestLocatorAllShouldWork(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<div><p>A</p><p>B</p><p>C</p></div>`))
	expected := []string{"A", "B", "C"}
	texts := make([]string, 0)
	p, err := page.Locator("p")
	require.NoError(t, err)
	locators, err := p.All()
	require.NoError(t, err)
	for _, locator := range locators {
		content, err := locator.TextContent()
		require.NoError(t, err)
		texts = append(texts, content)
	}
	require.ElementsMatch(t, expected, texts)
}

func TestLocatorsClearShouldWork(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(fmt.Sprintf("%s/input/textarea.html", server.PREFIX))
	require.NoError(t, err)
	button, err := page.Locator("input")
	require.NoError(t, err)
	require.NoError(t, button.Fill("some value"))
	ret, err := page.Evaluate(`result`)
	require.NoError(t, err)
	require.Equal(t, "some value", ret)
	require.NoError(t, button.Clear())
	ret, err = page.Evaluate(`result`)
	require.NoError(t, err)
	require.Equal(t, "", ret)
}
