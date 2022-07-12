package cfgenv

import (
	"encoding/json"
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

func Test_fold(t *testing.T) {
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
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := fold(c.from, c.onto)
			compareMaps(t, c.result, result)
		})
	}
}

func Test_varToMap(t *testing.T) {
	cases := []struct {
		name   string
		envVar string
		delim  string
		output object
	}{
		{
			name:   "Nest one",
			envVar: "ONE=test",
			delim:  "__",
			output: object{
				"ONE": object{
					"TWO": object{
						"THREE": "test",
					},
				},
			},
		},
		{
			name:   "Nest two",
			envVar: "ONE__TWO=test",
			delim:  "__",
			output: object{
				"ONE": object{
					"TWO": object{
						"THREE": "test",
					},
				},
			},
		},
		{
			name:   "Nest three",
			envVar: "ONE__TWO__THREE=test",
			delim:  "__",
			output: object{
				"ONE": object{
					"TWO": object{
						"THREE": "test",
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data := varToMap(c.envVar, c.delim)
			compareMaps(t, c.output, data)
		})
	}
}
