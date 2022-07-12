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

}

func Test_varToMap(t *testing.T) {
	cases := []struct {
		name   string
		envVar string
		delim  string
		output map[string]interface{}
	}{
		{
			name:   "String val",
			envVar: "",
		},
		{
			name: "Int val",
		},
		{
			name: "Float64 val",
		},
		{
			name: "Bool true",
		},
		{
			name: "Bool false",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data := varToMap(c.envVar, c.delim)
			compareMaps(t, c.output, data)
		})
	}
}
