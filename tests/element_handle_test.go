package playwright_test

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestElementHandleInnerText(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/dom.html")
	require.NoError(t, err)
	handle, err := page.QuerySelector("#inner")
	require.NoError(t, err)
	t1, err := handle.InnerText()
	require.NoError(t, err)
	require.Equal(t, t1, "Text, more text")
	t2, err := page.InnerText("#inner")
	require.NoError(t, err)
	require.Equal(t, t2, "Text, more text")
}

func TestElementHandleOwnerFrame(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = utils.AttachFrame(page, "iframe1", server.EMPTY_PAGE)
	require.NoError(t, err)
	frame := page.Frames()[1]
	elementHandle, err := frame.EvaluateHandle("document.body")
	require.NoError(t, err)
	ownerFrame, err := elementHandle.(playwright.ElementHandle).OwnerFrame()
	require.NoError(t, err)
	require.Equal(t, ownerFrame, frame)
	require.Equal(t, "iframe1", ownerFrame.Name())
}
func TestElementHandleContentFrame(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = utils.AttachFrame(page, "frame1", server.EMPTY_PAGE)
	require.NoError(t, err)
	elementHandle, err := page.QuerySelector("#frame1")
	require.NoError(t, err)
	frame, err := elementHandle.ContentFrame()
	require.NoError(t, err)
	require.Equal(t, frame, page.Frames()[1])
}
func TestElementHandleGetAttribute(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/dom.html")
	require.NoError(t, err)
	handle, err := page.QuerySelector("#outer")
	require.NoError(t, err)
	a1, err := handle.GetAttribute("name")
	require.NoError(t, err)
	require.Equal(t, "value", a1)
	a2, err := page.GetAttribute("#outer", "name")
	require.NoError(t, err)
	require.Equal(t, "value", a2)
}

func TestElementHandleDispatchEvent(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/input/button.html")
	require.NoError(t, err)
	element, err := page.QuerySelector("button")
	require.NoError(t, err)
	require.NoError(t, element.DispatchEvent("click"))
	result, err := page.Evaluate("() => result")
	require.NoError(t, err)
	require.Equal(t, "Clicked", result)
}

func TestElementHandleDispatchEventInitObject(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := page.SetContent(`
	<button onclick="window.eventBubbles = event.bubbles">ok</button>`)
	require.NoError(t, err)
	element, err := page.QuerySelector("button")
	require.NoError(t, err)
	require.NoError(t, element.DispatchEvent("click", map[string]interface{}{
		"bubbles": true,
	}))
	result, err := page.Evaluate("() => window.eventBubbles")
	require.NoError(t, err)
	require.Equal(t, true, result)
}

func TestElementHandleHover(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/input/scrollable.html")
	require.NoError(t, err)
	btn, err := page.QuerySelector("#button-6")
	require.NoError(t, err)
	require.NoError(t, btn.Hover())
	result, err := page.Evaluate(`document.querySelector("button:hover").id`)
	require.NoError(t, err)
	require.Equal(t, "button-6", result)
}

func TestElementHandleClick(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/input/button.html")
	require.NoError(t, err)
	btn, err := page.QuerySelector("button")
	require.NoError(t, err)
	require.NoError(t, btn.Click())
	result, err := page.Evaluate(`result`)
	require.NoError(t, err)
	require.Equal(t, "Clicked", result)
}

func TestElementHandleDblclick(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/input/button.html")
	require.NoError(t, err)
	_, err = page.Evaluate(`() => {
            window.double = false;
            button = document.querySelector('button');
            button.addEventListener('dblclick', event => {
            window.double = true;
            });
	}`)
	require.NoError(t, err)
	btn, err := page.QuerySelector("button")
	require.NoError(t, err)
	require.NoError(t, btn.Dblclick())
	result, err := page.Evaluate("double")
	require.NoError(t, err)
	require.Equal(t, true, result)

	result, err = page.Evaluate(`result`)
	require.NoError(t, err)
	require.Equal(t, "Clicked", result)
}

func TestElementBoundingBox(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetViewportSize(500, 500))
	_, err := page.Goto(server.PREFIX + "/grid.html")
	require.NoError(t, err)
	element_handle, err := page.QuerySelector(".box:nth-of-type(13)")
	require.NoError(t, err)
	box, err := element_handle.BoundingBox()
	require.NoError(t, err)
	require.Equal(t, 100, box.X)
	require.Equal(t, 50, box.Y)
	require.Equal(t, 50, box.Width)
	require.Equal(t, 50, box.Height)
}

func TestElementHandleTap(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent("<input id='checkbox' type='checkbox'></input>"))
	value, err := page.EvalOnSelector("input", "el => el.checked")
	require.NoError(t, err)
	require.Equal(t, false, value)

	elemHandle, err := page.QuerySelector("#checkbox")
	require.NoError(t, err)
	require.NoError(t, elemHandle.Tap())
	value, err = page.EvalOnSelector("input", "el => el.checked")
	require.NoError(t, err)
	require.Equal(t, true, value)
}

func TestElementHandleQuerySelectorAll(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`
	<div id="a1">
		<div class="foobar">
		</div>
		<div class="foobar">
		</div>
	</div>
	`))
	rootElement, err := page.QuerySelector("#a1")
	require.NoError(t, err)
	elements, err := rootElement.QuerySelectorAll(".foobar")
	require.NoError(t, err)
	require.Equal(t, 2, len(elements))
	className, err := elements[0].GetAttribute("class")
	require.NoError(t, err)
	require.Equal(t, "foobar", className)
}

func TestElementHandleEvalOnSelector(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`
	<div id="a1">
		<div id="a2">
			foobar
		</div>
	</div>
	`))
	rootElement, err := page.QuerySelector("#a1")
	require.NoError(t, err)
	innerText, err := rootElement.EvalOnSelector("#a2", "e => e.innerText")
	require.NoError(t, err)
	require.Equal(t, "foobar", innerText)
}

func TestElementHandleEvalOnSelectorAll(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`
	<div id="a1">
		<div class="foobar">
		</div>
		<div class="foobar">
		</div>
	</div>
	`))
	rootElement, err := page.QuerySelector("#a1")
	require.NoError(t, err)
	classNames, err := rootElement.EvalOnSelectorAll(".foobar", "elements => [...elements].map(e => e.getAttribute('class'))")
	require.NoError(t, err)
	require.Equal(t, []interface{}([]interface{}{"foobar", "foobar"}), classNames)
}

func TestElementHandleString(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	numberHandle, err := page.EvaluateHandle("() => 2")
	require.NoError(t, err)
	require.Equal(t, "JSHandle@2", numberHandle.String())

	stringHandle, err := page.EvaluateHandle("() => 'a'")
	require.NoError(t, err)
	require.Equal(t, "JSHandle@a", stringHandle.String())
}

func TestElementHandleCheck(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`
		<input type="checkbox"/>
	`))
	inputElement, err := page.QuerySelector("input")
	require.NoError(t, err)
	isChecked, err := inputElement.Evaluate("e => e.checked")
	require.NoError(t, err)
	require.Equal(t, false, isChecked)
	require.NoError(t, inputElement.Check())
	isChecked, err = inputElement.Evaluate("e => e.checked")
	require.NoError(t, err)
	require.Equal(t, true, isChecked)
}

func TestElementHandleUnCheck(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetContent(`
		<input type="checkbox" checked/>
	`))
	inputElement, err := page.QuerySelector("input")
	require.NoError(t, err)
	require.NoError(t, inputElement.Uncheck())
	isChecked, err := inputElement.Evaluate("e => e.checked")
	require.NoError(t, err)
	require.Equal(t, false, isChecked)
}

func TestElementHandleSelectOption(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent("<select id='lang'><option value='go'>go</option><option value='python'>python</option></select>"))
	elemHandle, err := page.QuerySelector("#lang")
	require.NoError(t, err)
	selected, err := elemHandle.SelectOption(playwright.SelectOptionValues{
		Value: &[]string{"go"},
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(selected))
	require.Equal(t, "go", selected[0])
}
