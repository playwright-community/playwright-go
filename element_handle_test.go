package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestElementHandleInnerText(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.PREFIX + "/dom.html")
	require.NoError(t, err)
	handle, err := helper.Page.QuerySelector("#inner")
	require.NoError(t, err)
	t1, err := handle.InnerText()
	require.NoError(t, err)
	require.Equal(t, t1, "Text, more text")
	t2, err := helper.Page.InnerText("#inner")
	require.NoError(t, err)
	require.Equal(t, t2, "Text, more text")
}

func TestElementHandleOwnerFrame(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = helper.utils.AttachFrame(helper.Page, "iframe1", helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	frame := helper.Page.Frames()[1]
	elementHandle, err := frame.EvaluateHandle("document.body")
	require.NoError(t, err)
	ownerFrame, err := elementHandle.(*ElementHandle).OwnerFrame()
	require.NoError(t, err)
	require.Equal(t, ownerFrame, frame)
	require.Equal(t, "iframe1", ownerFrame.Name())
}
func TestElementHandleContentFrame(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = helper.utils.AttachFrame(helper.Page, "frame1", helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	elementHandle, err := helper.Page.QuerySelector("#frame1")
	require.NoError(t, err)
	frame, err := elementHandle.ContentFrame()
	require.NoError(t, err)
	require.Equal(t, frame, helper.Page.Frames()[1])
}
func TestElementHandleGetAttribute(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.PREFIX + "/dom.html")
	require.NoError(t, err)
	handle, err := helper.Page.QuerySelector("#outer")
	require.NoError(t, err)
	a1, err := handle.GetAttribute("name")
	require.NoError(t, err)
	require.Equal(t, "value", a1)
	a2, err := helper.Page.GetAttribute("#outer", "name")
	require.NoError(t, err)
	require.Equal(t, "value", a2)
}

func TestElementHandleDispatchEvent(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.PREFIX + "/input/button.html")
	require.NoError(t, err)
	require.NoError(t, helper.Page.DispatchEvent("button", "click"))
	result, err := helper.Page.Evaluate("result")
	require.NoError(t, err)
	require.Equal(t, "Clicked", result)
}

func TestElementHandleHover(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.PREFIX + "/input/scrollable.html")
	require.NoError(t, err)
	btn, err := helper.Page.QuerySelector("#button-6")
	require.NoError(t, err)
	require.NoError(t, btn.Hover())
	result, err := helper.Page.Evaluate(`document.querySelector("button:hover").id`)
	require.NoError(t, err)
	require.Equal(t, "button-6", result)
}

func TestElementHandleClick(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.PREFIX + "/input/button.html")
	require.NoError(t, err)
	btn, err := helper.Page.QuerySelector("button")
	require.NoError(t, err)
	require.NoError(t, btn.Click())
	result, err := helper.Page.Evaluate(`result`)
	require.NoError(t, err)
	require.Equal(t, "Clicked", result)
}

func TestElementHandleDblClick(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.PREFIX + "/input/button.html")
	require.NoError(t, err)
	_, err = helper.Page.Evaluate(`() => {
            window.double = false;
            button = document.querySelector('button');
            button.addEventListener('dblclick', event => {
            window.double = true;
            });
	}`)
	require.NoError(t, err)
	btn, err := helper.Page.QuerySelector("button")
	require.NoError(t, err)
	require.NoError(t, btn.DblClick())
	result, err := helper.Page.Evaluate("double")
	require.NoError(t, err)
	require.Equal(t, true, result)

	result, err = helper.Page.Evaluate(`result`)
	require.NoError(t, err)
	require.Equal(t, "Clicked", result)
}
