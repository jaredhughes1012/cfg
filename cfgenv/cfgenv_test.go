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
