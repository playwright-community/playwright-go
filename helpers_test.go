package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testOptionsJSONSerialization struct {
	StringPointer  *string `json:"stringPointer"`
	NormalString   string  `json:"normalString"`
	WithoutJSONTag string
	WithJSONTag    string  `json:"withJSONTag"`
	SkipNilPtrs    *string `json:"skipNilPtrs"`
	SkipMe         *int    `json:"skipMe"`
}

func TestTransformOptions(t *testing.T) {
	// test data
	structVar := &testOptionsJSONSerialization{
		StringPointer:  String("1"),
		NormalString:   "2",
		WithoutJSONTag: "3",
		WithJSONTag:    "4",
	}
	var nilStrPtr *string
	testCases := []struct {
		name           string
		baseMap        map[string]interface{}
		optionalStruct interface{}
		expected       interface{}
	}{
		{
			name: "No options supplied",
			baseMap: map[string]interface{}{
				"1234": nilStrPtr,
				"foo":  "bar",
			},
			expected: map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			name: "Options are nil",
			baseMap: map[string]interface{}{
				"foo": "bar",
			},
			optionalStruct: nil,
			expected: map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			name: "JSON serialization works",
			baseMap: map[string]interface{}{
				"foo": "bar",
			},
			optionalStruct: structVar,
			expected: map[string]interface{}{
				"foo":            "bar",
				"stringPointer":  String("1"),
				"normalString":   "2",
				"WithoutJSONTag": "3",
				"withJSONTag":    "4",
			},
		},
		{
			name: "Second overwrites the first one",
			baseMap: map[string]interface{}{
				"foo": "1",
			},
			optionalStruct: map[string]interface{}{
				"foo": "2",
			},
			expected: map[string]interface{}{
				"foo": "2",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, transformOptions(tc.baseMap, tc.optionalStruct))
		})
	}
}

func TestRemapMapToStruct(t *testing.T) {
	ourStruct := struct {
		V1 string `json:"v1"`
	}{}
	inMap := map[string]interface{}{
		"v1": "foobar",
	}
	remapMapToStruct(inMap, &ourStruct)
	require.Equal(t, ourStruct.V1, "foobar")
}

func TestConvertSelectOptionSet(t *testing.T) {
	testCases := []struct {
		name         string
		optionValues SelectOptionValues
		expected     interface{}
	}{
		{
			name:         "SelectOptionValues is nil",
			optionValues: SelectOptionValues{},
			expected:     make(map[string]interface{}),
		},
		{
			name: "SelectOptionValues is supplied",
			optionValues: SelectOptionValues{
				Values:  StringSlice("a", "b"),
				Indexes: IntSlice(1),
				Labels:  StringSlice("x"),
			},
			expected: map[string]interface{}{
				"options": []map[string]interface{}{
					{"value": "a"}, {"value": "b"}, {"index": 1}, {"label": "x"},
				},
			},
		},
		{
			name: "Only value is supplied",
			optionValues: SelectOptionValues{
				Values: StringSlice("a", "b"),
			},
			expected: map[string]interface{}{
				"options": []map[string]interface{}{
					{"value": "a"}, {"value": "b"},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, convertSelectOptionSet(tc.optionValues))
		})
	}
}
