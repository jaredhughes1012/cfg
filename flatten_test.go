package cfg

import "testing"

func flattenAndCheck(t *testing.T, data object, key, expected string) {
	flatMap := flatten(data)
	actual := flatMap[key]

	if expected != actual {
		t.Errorf("%s != %s (key %s)", expected, actual, key)
	}
}

func testAllNesting(t *testing.T, input any, output string) {
	t.Run("One deep", func(t *testing.T) {
		data := object{
			"one": input,
		}
		flattenAndCheck(t, data, "one", output)
	})

	t.Run("Two deep", func(t *testing.T) {
		data := object{
			"one": object{
				"two": input,
			},
		}
		flattenAndCheck(t, data, "one:two", output)
	})

	t.Run("Three deep", func(t *testing.T) {
		data := object{
			"one": object{
				"two": object{
					"three": output,
				},
			},
		}
		flattenAndCheck(t, data, "one:two:three", output)
	})

	t.Run("Decapitalized", func(t *testing.T) {
		data := object{
			"LOWERCase": output,
		}

		flattenAndCheck(t, data, "lowercase", output)
	})

	t.Run("Remove underscores", func(t *testing.T) {
		data := object{
			"one_word": output,
		}

		flattenAndCheck(t, data, "lowercase", output)
	})
}

func Test_flatten(t *testing.T) {
	cases := []struct {
		name   string
		input  any
		output string
	}{
		{
			name:   "String",
			input:  "test",
			output: "test",
		},
		{
			name:   "Integer",
			input:  15,
			output: "15",
		},
		// {
		// 	name:   "Float32",
		// 	input:  float32(18),
		// 	output: "18",
		// },
		{
			name:   "Float64",
			input:  float64(24),
			output: "24",
		},
		{
			name:   "Bool true",
			input:  true,
			output: "true",
		},
		{
			name:   "Bool false",
			input:  false,
			output: "false",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			testAllNesting(t, c.input, c.output)
		})
	}
}
