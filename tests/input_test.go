package playwright_test

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestMouseMove(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	if isWebKit {
		_, err := page.Evaluate(`() => new Promise(requestAnimationFrame)`)
		require.NoError(t, err)
	}
	require.NoError(t, page.Mouse().Move(100, 100))
	_, err := page.Evaluate(`() => {
    window['result'] = [];
    document.addEventListener('mousemove', event => {
      window['result'].push([event.clientX, event.clientY]);
    });
  }`)
	require.NoError(t, err)
	require.NoError(t, page.Mouse().Move(200, 300, playwright.MouseMoveOptions{
		Steps: playwright.Int(5),
	}))
	result, err := page.Evaluate("result")
	require.NoError(t, err)
	require.Equal(t, []interface{}([]interface{}{[]interface{}{120, 140}, []interface{}{140, 180}, []interface{}{160, 220}, []interface{}{180, 260}, []interface{}{200, 300}}), result)
}

func TestMouseDown(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<button onmousedown="window.clicked=true"/>`))
	require.NoError(t, page.Hover("button"))
	require.NoError(t, page.Mouse().Down())
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestMouseUp(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<button onmouseup="window.clicked=true"/>`))
	require.NoError(t, page.Hover("button"))
	require.NoError(t, page.Mouse().Down())
	require.NoError(t, page.Mouse().Up())
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestMouseClick(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<button onclick="window.clicked=true" style="width: 500px; height: 500px;"/>`))
	require.NoError(t, page.Hover("button"))
	require.NoError(t, page.Mouse().Click(100, 100))
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestMouseDblclick(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<button ondblclick="window.clicked=true" style="width: 500px; height: 500px;"/>`))
	require.NoError(t, page.Hover("button"))
	require.NoError(t, page.Mouse().Dblclick(100, 100))
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestKeyboardDown(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<input onkeydown="window.clicked=true"/>`))
	require.NoError(t, page.Click("input"))
	require.NoError(t, page.Keyboard().Down("Enter"))
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestKeyboardUp(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<input onkeyup="window.clicked=true"/>`))
	require.NoError(t, page.Click("input"))
	require.NoError(t, page.Keyboard().Up("Enter"))
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestKeyboardInsertText(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<input oninput="window.clicked=true"/>`))
	require.NoError(t, page.Click("input"))
	require.NoError(t, page.Keyboard().InsertText("abc123"))
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestKeyboardType(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<input oninput="window.clicked=true"/>`))
	require.NoError(t, page.Click("input"))
	require.NoError(t, page.Keyboard().Type("abc123"))
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestElementHandleType(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<input oninput="window.clicked=true"/>`))
	require.NoError(t, page.Click("input"))
	inputElement, err := page.QuerySelector("input")
	require.NoError(t, err)
	require.NoError(t, inputElement.Type("abc123"))
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestElementHandleFill(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<input oninput="window.clicked=true"/>`))
	require.NoError(t, page.Click("input"))
	inputElement, err := page.QuerySelector("input")
	require.NoError(t, err)
	require.NoError(t, inputElement.Fill("abc123"))
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestKeyboardInsertPress(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<input onkeydown="window.clicked=true"/>`))
	require.NoError(t, page.Click("input"))
	require.NoError(t, page.Keyboard().Press("A"))
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestElementHandlePress(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<input onkeydown="window.clicked=true"/>`))
	require.NoError(t, page.Click("input"))
	inputElement, err := page.QuerySelector("input")
	require.NoError(t, err)
	require.NoError(t, inputElement.Press("A"))
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestTouchscreenTap(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<button onclick="window.clicked=true" style="width: 500px; height: 500px;"/>`))
	require.NoError(t, page.Touchscreen().Tap(100, 100))
	result, err := page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}
