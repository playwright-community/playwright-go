package playwright_test

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestBrowserContextExposeBinding(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	bindingSource := []*playwright.BindingSource{}
	binding := func(source *playwright.BindingSource, a, b int) int {
		bindingSource = append(bindingSource, source)
		return a + b
	}
	err := context.ExposeBinding("add", func(source *playwright.BindingSource, args ...interface{}) interface{} {
		return binding(source, args[0].(int), args[1].(int))
	})
	require.NoError(t, err)
	page, err := context.NewPage()
	require.NoError(t, err)
	result, err := page.Evaluate("add(5, 6)")
	require.NoError(t, err)
	require.Equal(t, bindingSource[0].Context, context)
	require.Equal(t, bindingSource[0].Page, page)
	require.Equal(t, bindingSource[0].Frame, page.MainFrame())
	require.Equal(t, result, 11)
}

func TestBrowserContextExposeFunction(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	err := context.ExposeFunction("add", func(args ...interface{}) interface{} {
		return args[0].(int) + args[1].(int)
	})
	require.NoError(t, err)
	page, err := context.NewPage()
	require.NoError(t, err)
	err = page.ExposeFunction("mul", func(args ...interface{}) interface{} {
		return args[0].(int) * args[1].(int)
	})
	require.NoError(t, err)
	err = context.ExposeFunction("sub", func(args ...interface{}) interface{} {
		return args[0].(int) - args[1].(int)
	})
	require.NoError(t, err)
	result, err := page.Evaluate(`async function() {
		return { mul: await mul(9, 4), add: await add(9, 4), sub: await sub(9, 4) }
	  }`)
	require.NoError(t, err)
	resultMap := result.(map[string]interface{})
	require.Equal(t, 36, resultMap["mul"])
	require.Equal(t, 13, resultMap["add"])
	require.Equal(t, 5, resultMap["sub"])
}
