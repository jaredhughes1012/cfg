package mapconvert

import (
	"encoding/json"
	"strings"
	"testing"
)

type object map[string]any

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
		from   object
		onto   object
		result object
	}{
		{
			name: "Onto Empty",
			from: object{
				"test": "val",
			},
			onto: object{},
			result: object{
				"test": "val",
			},
		},
		{
			name: "Two non empty",
			from: object{
				"test1": "val1",
			},
			onto: object{
				"test2": "val2",
			},
			result: object{
				"test1": "val1",
				"test2": "val2",
			},
		},
		{
			name: "Overwrite",
			from: object{
				"test": "val1",
			},
			onto: object{
				"test": "val2",
			},
			result: object{
				"test": "val1",
			},
		},
		{
			name: "Map onto primitive",
			from: object{
				"test": object{
					"test2": "val",
				},
			},
			onto: object{
				"test": "string",
			},
			result: object{
				"test": object{
					"test2": "val",
				},
			},
		},
		{
			name: "Primitive onto map",
			from: object{
				"test": "string",
			},
			onto: object{
				"test": object{
					"test2": "val",
				},
			},
			result: object{
				"test": "string",
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
		input    object
		expected object
		delim    string
	}{
		{
			name:  "One deep",
			delim: ".",
			input: object{
				"one": "test",
			},
			expected: object{
				"one": "test",
			},
		},
		{
			name:  "Two deep",
			delim: ",",
			input: object{
				"one": object{
					"two": "test",
				},
			},
			expected: object{
				"one,two": "test",
			},
		},
		{
			name:  "Three deep",
			delim: ".",
			input: object{
				"one": object{
					"two": object{
						"three": "test",
					},
				},
			},
			expected: object{
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
	input := object{
		"KEY": "value",
	}
	expected := object{
		"key": "value",
	}

	actual := ConvertKeys(input, func(s string) string {
		return strings.ToLower(s)
	})

	compareMaps(t, expected, actual)
}
