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
}

func TestTransformOptions(t *testing.T) {
	// test data
	var sizeOneNilSlice []interface{}
	sizeOneNilSlice = append(sizeOneNilSlice, nil)
	structVar := &testOptionsJSONSerialization{
		StringPointer:  String("1"),
		NormalString:   "2",
		WithoutJSONTag: "3",
		WithJSONTag:    "4",
	}
	var sizeOneJSONTest []interface{}
	sizeOneJSONTest = append(sizeOneJSONTest, structVar)
	testCases := []struct {
		name           string
		baseMap        map[string]interface{}
		optionalStruct []interface{}
		expected       interface{}
	}{
		{
			name: "No options supplied",
			baseMap: map[string]interface{}{
				"foo": "bar",
			},
			optionalStruct: make([]interface{}, 0),
			expected: map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			name: "Options are nil",
			baseMap: map[string]interface{}{
				"foo": "bar",
			},
			optionalStruct: sizeOneNilSlice,
			expected: map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			name: "JSON serialization works",
			baseMap: map[string]interface{}{
				"foo": "bar",
			},
			optionalStruct: sizeOneJSONTest,
			expected: map[string]interface{}{
				"foo":            "bar",
				"stringPointer":  String("1"),
				"normalString":   "2",
				"WithoutJSONTag": "3",
				"withJSONTag":    "4",
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
