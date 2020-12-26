package playwright_test

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestJSHandleGetProperty(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	aHandle, err := helper.Page.EvaluateHandle(`() => ({
		one: 1,
		two: 2,
		three: 3
	})`)
	require.NoError(t, err)
	twoHandle, err := aHandle.(playwright.JSHandle).GetProperty("two")
	require.NoError(t, err)
	value, err := twoHandle.JSONValue()
	require.NoError(t, err)
	require.Equal(t, value, 2)
}

func TestJSHandleGetProperties(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	aHandle, err := helper.Page.EvaluateHandle(`() => ({
		one: 1,
		two: 2,
		three: 3
	})`)
	require.NoError(t, err)
	twoHandle, err := aHandle.(playwright.JSHandle).GetProperties()
	require.NoError(t, err)
	v1, err := twoHandle["one"].JSONValue()
	require.NoError(t, err)
	require.Equal(t, 1, v1)

	v1, err = twoHandle["two"].JSONValue()
	require.NoError(t, err)
	require.Equal(t, 2, v1)

	v1, err = twoHandle["three"].JSONValue()
	require.NoError(t, err)
	require.Equal(t, 3, v1)
}

func TestJSHandleEvaluate(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	aHandle, err := helper.Page.EvaluateHandle(`() => ({
		one: 1,
		two: 2,
		three: 3
	})`)
	require.NoError(t, err)
	twoHandle, err := aHandle.(playwright.JSHandle).Evaluate("x => x")
	require.NoError(t, err)
	values := twoHandle.(map[string]interface{})
	require.Equal(t, 1, values["one"])
	require.Equal(t, 2, values["two"])
	require.Equal(t, 3, values["three"])
}

func TestJSHandleEvaluateHandle(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	aHandle, err := helper.Page.EvaluateHandle(`() => ({
		one: 1,
		two: 2,
		three: 3
	})`)
	require.NoError(t, err)
	twoHandle, err := aHandle.(playwright.JSHandle).EvaluateHandle("x => x")
	require.NoError(t, err)
	twoHandle, err = twoHandle.(playwright.JSHandle).GetProperty("two")
	require.NoError(t, err)
	value, err := twoHandle.(playwright.JSHandle).JSONValue()
	require.NoError(t, err)
	require.Equal(t, value, 2)
}

func TestJSHandleTypeParsing(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	aHandle, err := helper.Page.EvaluateHandle(`() => ({
		an_integer: 1,
		a_float: 2.2222222222,
		a_string_of_an_integer: "3",
	})`)
	require.NoError(t, err)
	twoHandle, err := aHandle.(playwright.JSHandle).EvaluateHandle("x => x")
	require.NoError(t, err)

	integerHandle, err := twoHandle.(playwright.JSHandle).GetProperty("an_integer")
	require.NoError(t, err)
	intV, err := integerHandle.JSONValue()
	require.NoError(t, err)
	_, ok := intV.(int)
	require.True(t, ok)
	_, ok = intV.(float64)
	require.False(t, ok)

	floatHandle, err := twoHandle.(playwright.JSHandle).GetProperty("a_float")
	require.NoError(t, err)
	floatV, err := floatHandle.JSONValue()
	require.NoError(t, err)
	_, ok = floatV.(float64)
	require.True(t, ok)
	_, ok = floatV.(int)
	require.False(t, ok)

	stringHandle, err := twoHandle.(playwright.JSHandle).GetProperty("a_string_of_an_integer")
	require.NoError(t, err)
	stringV, err := stringHandle.JSONValue()
	require.NoError(t, err)
	_, ok = stringV.(string)
	require.True(t, ok)
	_, ok = stringV.(int)
	require.False(t, ok)
}
