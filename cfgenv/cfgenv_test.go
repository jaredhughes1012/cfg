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
		key    string
		val    string
		delim  string
		output object
	}{
		{
			name:  "Nest one",
			key:   "ONE",
			val:   "test",
			delim: "__",
			output: object{
				"ONE": "test",
			},
		},
		{
			name:  "Nest two",
			key:   "ONE__TWO",
			val:   "test",
			delim: "__",
			output: object{
				"ONE": object{
					"TWO": "test",
				},
			},
		},
		{
			name:  "Nest three",
			key:   "ONE__TWO__THREE",
			val:   "test",
			delim: "__",
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
			data := varToMap(c.key, c.val, c.delim)
			compareMaps(t, c.output, data)
		})
	}
}

func Test_prepareVar(t *testing.T) {
	cases := []struct {
		name   string
		envVar string
		prefix string
		key    string
		val    string
	}{
		{
			name:   "Prefix match",
			envVar: "PREFIX_KEY=val",
			prefix: "PREFIX_",
			key:    "KEY",
			val:    "val",
		},
		{
			name:   "Prefix substring",
			envVar: "STUFF_PREFIX_KEY=val",
			prefix: "PREFIX_",
			key:    "",
			val:    "",
		},
		{
			name:   "Prefix no match",
			envVar: "TEST_KEY=val",
			prefix: "PREFIX_",
			key:    "",
			val:    "",
		},
		{
			name:   "'=' in val",
			envVar: "PREFIX_KEY=val=",
			prefix: "PREFIX_",
			key:    "KEY",
			val:    "val=",
		},
		{
			name:   "Empty prefix",
			envVar: "KEY=val",
			prefix: "",
			key:    "KEY",
			val:    "val",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			key, val := prepareVar(c.prefix, c.envVar)
			if c.key != key {
				t.Errorf("Key %s != %s", c.key, key)
			}
			if c.val != val {
				t.Errorf("Val %s != %s", c.val, val)
			}
		})
	}
}
