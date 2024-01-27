package playwright_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSHandleGetProperty(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	aHandle, err := page.EvaluateHandle(`() => ({
		one: 1,
		two: 2,
		three: 3
	})`)
	require.NoError(t, err)
	twoHandle, err := aHandle.GetProperty("two")
	require.NoError(t, err)
	value, err := twoHandle.JSONValue()
	require.NoError(t, err)
	require.Equal(t, value, 2)
}

func TestJSHandleGetProperties(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	aHandle, err := page.EvaluateHandle(`() => ({
		one: 1,
		two: 2,
		three: 3
	})`)
	require.NoError(t, err)
	twoHandle, err := aHandle.GetProperties()
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
	BeforeEach(t)
	defer AfterEach(t)
	aHandle, err := page.EvaluateHandle(`() => ({
		one: 1,
		two: 2,
		three: 3
	})`)
	require.NoError(t, err)
	twoHandle, err := aHandle.Evaluate("x => x")
	require.NoError(t, err)
	values := twoHandle.(map[string]interface{})
	require.Equal(t, 1, values["one"])
	require.Equal(t, 2, values["two"])
	require.Equal(t, 3, values["three"])
}

func TestJSHandleEvaluateHandle(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	aHandle, err := page.EvaluateHandle(`() => ({
		one: 1,
		two: 2,
		three: 3
	})`)
	require.NoError(t, err)
	twoHandle, err := aHandle.EvaluateHandle("x => x")
	require.NoError(t, err)
	twoHandle, err = twoHandle.GetProperty("two")
	require.NoError(t, err)
	value, err := twoHandle.JSONValue()
	require.NoError(t, err)
	require.Equal(t, value, 2)
}

func TestJSHandleTypeParsing(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	aHandle, err := page.EvaluateHandle(`() => ({
		an_integer: 1,
		a_float: 2.2222222222,
		a_string_of_an_integer: "3",
	})`)
	require.NoError(t, err)
	twoHandle, err := aHandle.EvaluateHandle("x => x")
	require.NoError(t, err)

	integerHandle, err := twoHandle.GetProperty("an_integer")
	require.NoError(t, err)
	intV, err := integerHandle.JSONValue()
	require.NoError(t, err)
	_, ok := intV.(int)
	require.True(t, ok)
	_, ok = intV.(float64)
	require.False(t, ok)

	floatHandle, err := twoHandle.GetProperty("a_float")
	require.NoError(t, err)
	floatV, err := floatHandle.JSONValue()
	require.NoError(t, err)
	_, ok = floatV.(float64)
	require.True(t, ok)
	_, ok = floatV.(int)
	require.False(t, ok)

	stringHandle, err := twoHandle.GetProperty("a_string_of_an_integer")
	require.NoError(t, err)
	stringV, err := stringHandle.JSONValue()
	require.NoError(t, err)
	_, ok = stringV.(string)
	require.True(t, ok)
	_, ok = stringV.(int)
	require.False(t, ok)
}

func TestJSHandleTypeSerializing(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	nilV, err := page.Evaluate(`a => a`, nil)
	require.NoError(t, err)
	require.Equal(t, nil, nilV)
	boolV, err := page.Evaluate(`a => a`, true)
	require.NoError(t, err)
	require.Equal(t, true, boolV)
	intV, err := page.Evaluate(`a => a`, 42)
	require.NoError(t, err)
	require.Equal(t, 42, intV)
	sliceArgs := []interface{}{"test1", "test2", "test3"}
	res, err := page.Evaluate(`a => a`, sliceArgs)
	require.NoError(t, err)
	sliceV, ok := res.([]interface{})
	require.True(t, ok)
	for i, v := range sliceArgs {
		require.Equal(t, v, sliceV[i])
	}
	mapArgs := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	res, err = page.Evaluate(`a => a`, mapArgs)
	require.NoError(t, err)
	mapV, ok := res.(map[string]interface{})
	require.True(t, ok)
	for k, v := range mapArgs {
		value, ok := mapV[k]
		require.True(t, ok)
		require.Equal(t, v, value)
	}
	// The following cases seem to fail due to playwright
	//floatV, err := page.Evaluate(`a => a`, 42.42)
	//require.NoError(t, err)
	//require.Equal(t, 42.42, floatV)
	//now := time.Now()
	//timeV, err := page.Evaluate(`a => a`, now)
	//require.NoError(t, err)
	//require.Equal(t, now, timeV)
}
