package cfg

import (
	"errors"
	"testing"
)

var (
	errTest  = errors.New("test")
	testData = map[string]any{
		"key": "value",
	}
)

type testLoader struct {
	data map[string]any
	err  error
}

func (loader testLoader) Load() (map[string]any, error) {
	return loader.data, loader.err
}

func newTestLoader(data map[string]any, err error) *testLoader {
	return &testLoader{
		data: data,
		err:  err,
	}
}

func newConfigAndLoad(loaders ...*testLoader) (*Config, error) {
	cfg := New()
	for _, l := range loaders {
		cfg.Add(l)
	}

	err := cfg.Load()
	return cfg, err
}

func Test_Config_Load(t *testing.T) {
	cases := []struct {
		name    string
		err     error
		loaders []*testLoader
	}{
		{
			name: "One loader",
			err:  nil,
			loaders: []*testLoader{
				newTestLoader(testData, nil),
			},
		},
		{
			name: "Three loaders",
			err:  nil,
			loaders: []*testLoader{
				newTestLoader(testData, nil),
				newTestLoader(testData, nil),
				newTestLoader(testData, nil),
			},
		},
		{
			name: "One failing loader",
			err:  errTest,
			loaders: []*testLoader{
				newTestLoader(nil, errTest),
			},
		},
		{
			name: "One successful one failing",
			err:  errTest,
			loaders: []*testLoader{
				newTestLoader(testData, nil),
				newTestLoader(nil, errTest),
			},
		},
		{
			name: "Prioritize first error",
			err:  errTest,
			loaders: []*testLoader{
				newTestLoader(nil, errTest),
				newTestLoader(nil, errors.New("another test")),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := newConfigAndLoad(c.loaders...)
			if c.err != err {
				t.Errorf("%v != %v", c.err, err)
			}
		})
	}
}

func Test_Config_GetString(t *testing.T) {
	cases := []struct {
		name     string
		key      string
		loader   *testLoader
		expected string
		isErr    bool
	}{
		{
			name:     "Happy path",
			loader:   newTestLoader(map[string]any{"key": "val"}, nil),
			key:      "key",
			expected: "val",
			isErr:    false,
		},
		{
			name:     "Value not found",
			loader:   newTestLoader(map[string]any{"key": "val"}, nil),
			key:      "key2",
			expected: "",
			isErr:    true,
		},
		{
			name:     "Not string",
			loader:   newTestLoader(map[string]any{"key": 1}, nil),
			key:      "key",
			expected: "",
			isErr:    true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cfg, err := newConfigAndLoad(c.loader)
			if err != nil {
				t.Fatalf("%v", err)
			}

			actual, err := cfg.GetString(c.key)
			if c.isErr && err == nil {
				t.Error("No error when error expected")
			} else if !c.isErr {
				mustVal := cfg.MustGetString(c.key)
				if c.expected != actual {
					t.Errorf("Value %s != %s", c.expected, actual)
				}
				if actual != mustVal {
					t.Errorf("Must Value %s != %s", actual, mustVal)
				}
				if err != nil {
					t.Errorf("Unexpected error %v", err)
				}
			}
		})
	}
}

func Test_Config_GetInt(t *testing.T) {
	cases := []struct {
		name     string
		key      string
		loader   *testLoader
		expected int
		isErr    bool
	}{
		{
			name:     "Happy path int",
			loader:   newTestLoader(map[string]any{"key": 5}, nil),
			key:      "key",
			expected: 5,
			isErr:    false,
		},
		{
			name:     "Happy path string",
			loader:   newTestLoader(map[string]any{"key": "7"}, nil),
			key:      "key",
			expected: 7,
			isErr:    false,
		},
		{
			name:   "Value not found",
			loader: newTestLoader(map[string]any{"key": "val"}, nil),
			key:    "key2",
			isErr:  true,
		},
		{
			name:   "Non integer string",
			loader: newTestLoader(map[string]any{"key": "test"}, nil),
			key:    "key",
			isErr:  true,
		},
		{
			name:   "Non integer",
			loader: newTestLoader(map[string]any{"key": 35.02}, nil),
			key:    "key",
			isErr:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cfg, err := newConfigAndLoad(c.loader)
			if err != nil {
				t.Fatalf("%v", err)
			}

			actual, err := cfg.GetInt(c.key)
			if c.isErr && err == nil {
				t.Error("No error when error expected")
			} else if !c.isErr {
				mustVal := cfg.MustGetInt(c.key)
				if c.expected != actual {
					t.Errorf("Value %d != %d", c.expected, actual)
				}
				if actual != mustVal {
					t.Errorf("Must Value %d != %d", actual, mustVal)
				}
				if err != nil {
					t.Errorf("Unexpected error %v", err)
				}
			}
		})
	}
}

func Test_Config_GetFloat64(t *testing.T) {
	cases := []struct {
		name     string
		key      string
		loader   *testLoader
		expected float64
		isErr    bool
	}{
		{
			name:     "Happy path float64",
			loader:   newTestLoader(map[string]any{"key": float64(7)}, nil),
			key:      "key",
			expected: 7,
			isErr:    false,
		},
		{
			name:     "Happy path float32",
			loader:   newTestLoader(map[string]any{"key": float32(9)}, nil),
			key:      "key",
			expected: 9,
			isErr:    false,
		},
		{
			name:     "Happy path int",
			loader:   newTestLoader(map[string]any{"key": 2}, nil),
			key:      "key",
			expected: 2,
			isErr:    false,
		},
		{
			name:     "Happy path string",
			loader:   newTestLoader(map[string]any{"key": "14"}, nil),
			key:      "key",
			expected: 14,
			isErr:    false,
		},
		{
			name:   "Value not found",
			loader: newTestLoader(map[string]any{"key": "val"}, nil),
			key:    "key2",
			isErr:  true,
		},
		{
			name:   "Non float64 string",
			loader: newTestLoader(map[string]any{"key": "test"}, nil),
			key:    "key",
			isErr:  true,
		},
		{
			name:   "Non float64",
			loader: newTestLoader(map[string]any{"key": false}, nil),
			key:    "key",
			isErr:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cfg, err := newConfigAndLoad(c.loader)
			if err != nil {
				t.Fatalf("%v", err)
			}

			actual, err := cfg.GetFloat64(c.key)
			if c.isErr && err == nil {
				t.Error("No error when error expected")
			} else if !c.isErr {
				mustVal := cfg.MustGetFloat64(c.key)
				if c.expected != actual {
					t.Errorf("Value %f != %f", c.expected, actual)
				}
				if actual != mustVal {
					t.Errorf("Must Value %f != %f", actual, mustVal)
				}
				if err != nil {
					t.Errorf("Unexpected error %v", err)
				}
			}
		})
	}
}
