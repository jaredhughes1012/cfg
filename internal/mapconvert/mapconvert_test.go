package mapconvert

import (
	"encoding/json"
	"strings"
	"testing"
)

func compareMaps(t *testing.T, expected, actual map[string]interface{}) {
	b, _ := json.Marshal(expected)
	expStr := string(b)

	b, _ = json.Marshal(actual)
	actStr := string(b)

	if expStr != actStr {
		t.Errorf("%s != %s", expStr, actStr)
	}
}

func Test_Fold(t *testing.T) {
	cases := []struct {
		name   string
		from   map[string]any
		onto   map[string]any
		result map[string]any
	}{
		{
			name: "Onto Empty",
			from: map[string]any{
				"test": "val",
			},
			onto: map[string]any{},
			result: map[string]any{
				"test": "val",
			},
		},
		{
			name: "Two non empty",
			from: map[string]any{
				"test1": "val1",
			},
			onto: map[string]any{
				"test2": "val2",
			},
			result: map[string]any{
				"test1": "val1",
				"test2": "val2",
			},
		},
		{
			name: "Overwrite",
			from: map[string]any{
				"test": "val1",
			},
			onto: map[string]any{
				"test": "val2",
			},
			result: map[string]any{
				"test": "val1",
			},
		},
		{
			name: "Map onto primitive",
			from: map[string]any{
				"test": map[string]any{
					"test2": "val",
				},
			},
			onto: map[string]any{
				"test": "string",
			},
			result: map[string]any{
				"test": map[string]any{
					"test2": "val",
				},
			},
		},
		{
			name: "Primitive onto map",
			from: map[string]any{
				"test": "string",
			},
			onto: map[string]any{
				"test": map[string]any{
					"test2": "val",
				},
			},
			result: map[string]any{
				"test": "string",
			},
		},
		{
			name: "Nested Flatten",
			from: map[string]any{
				"test": map[string]any{
					"test1": "val1",
				},
			},
			onto: map[string]any{
				"test": map[string]any{
					"test2": "val2",
				},
			},
			result: map[string]any{
				"test": map[string]any{
					"test1": "val1",
					"test2": "val2",
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := Fold(c.from, c.onto)
			compareMaps(t, c.result, result)
		})
	}
}

func Test_Flatten(t *testing.T) {
	cases := []struct {
		name     string
		input    map[string]any
		expected map[string]any
		delim    string
	}{
		{
			name:  "One deep",
			delim: ".",
			input: map[string]any{
				"one": "test",
			},
			expected: map[string]any{
				"one": "test",
			},
		},
		{
			name:  "Two deep",
			delim: ",",
			input: map[string]any{
				"one": map[string]any{
					"two": "test",
				},
			},
			expected: map[string]any{
				"one,two": "test",
			},
		},
		{
			name:  "Three deep",
			delim: ".",
			input: map[string]any{
				"one": map[string]any{
					"two": map[string]any{
						"three": "test",
					},
				},
			},
			expected: map[string]any{
				"one.two.three": "test",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := Flatten(c.input, c.delim)
			compareMaps(t, c.expected, actual)
		})
	}
}

func Test_ConvertKeys(t *testing.T) {
	input := map[string]any{
		"KEY": "value",
	}
	expected := map[string]any{
		"key": "value",
	}

	actual := ConvertKeys(input, func(s string) string {
		return strings.ToLower(s)
	})

	compareMaps(t, expected, actual)
}
