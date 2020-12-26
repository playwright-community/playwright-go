package playwright_test

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestMouseMove(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	if helper.IsWebKit {
		_, err := helper.Page.Evaluate(`() => new Promise(requestAnimationFrame)`)
		require.NoError(t, err)
	}
	require.NoError(t, helper.Page.Mouse().Move(100, 100))
	_, err := helper.Page.Evaluate(`() => {
    window['result'] = [];
    document.addEventListener('mousemove', event => {
      window['result'].push([event.clientX, event.clientY]);
    });
  }`)
	require.NoError(t, err)
	require.NoError(t, helper.Page.Mouse().Move(200, 300, playwright.MouseMoveOptions{
		Steps: playwright.Int(5),
	}))
	result, err := helper.Page.Evaluate("result")
	require.NoError(t, err)
	require.Equal(t, []interface{}([]interface{}{[]interface{}{120, 140}, []interface{}{140, 180}, []interface{}{160, 220}, []interface{}{180, 260}, []interface{}{200, 300}}), result)
}

func TestMouseDown(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.SetContent(`<button onmousedown="window.clicked=true"/>`))
	require.NoError(t, helper.Page.Hover("button"))
	require.NoError(t, helper.Page.Mouse().Down())
	result, err := helper.Page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestMouseUp(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.SetContent(`<button onmouseup="window.clicked=true"/>`))
	require.NoError(t, helper.Page.Hover("button"))
	require.NoError(t, helper.Page.Mouse().Down())
	require.NoError(t, helper.Page.Mouse().Up())
	result, err := helper.Page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestMouseClick(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.SetContent(`<button onclick="window.clicked=true" style="width: 500px; height: 500px;"/>`))
	require.NoError(t, helper.Page.Hover("button"))
	require.NoError(t, helper.Page.Mouse().Click(100, 100))
	result, err := helper.Page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestMouseDblClick(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.SetContent(`<button ondblclick="window.clicked=true" style="width: 500px; height: 500px;"/>`))
	require.NoError(t, helper.Page.Hover("button"))
	require.NoError(t, helper.Page.Mouse().DblClick(100, 100))
	result, err := helper.Page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestKeyboardDown(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.SetContent(`<input onkeydown="window.clicked=true"/>`))
	require.NoError(t, helper.Page.Click("input"))
	require.NoError(t, helper.Page.Keyboard().Down("Enter"))
	result, err := helper.Page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestKeyboardUp(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.SetContent(`<input onkeyup="window.clicked=true"/>`))
	require.NoError(t, helper.Page.Click("input"))
	require.NoError(t, helper.Page.Keyboard().Up("Enter"))
	result, err := helper.Page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestKeyboardInsertText(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.SetContent(`<input oninput="window.clicked=true"/>`))
	require.NoError(t, helper.Page.Click("input"))
	require.NoError(t, helper.Page.Keyboard().InsertText("abc123"))
	result, err := helper.Page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestKeyboardType(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.SetContent(`<input oninput="window.clicked=true"/>`))
	require.NoError(t, helper.Page.Click("input"))
	require.NoError(t, helper.Page.Keyboard().Type("abc123"))
	result, err := helper.Page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}

func TestKeyboardInsertPress(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.SetContent(`<input onkeydown="window.clicked=true"/>`))
	require.NoError(t, helper.Page.Click("input"))
	require.NoError(t, helper.Page.Keyboard().Press("A"))
	result, err := helper.Page.Evaluate("window.clicked")
	require.NoError(t, err)
	require.True(t, result.(bool))
}
