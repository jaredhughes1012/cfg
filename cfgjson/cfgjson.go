package cfgjson

import (
	"encoding/json"
	"os"

	"github.com/jaredhughes1012/cfg"
)

// Config loader designed to load from JSON files
type JsonLoader struct {
	path     string
	required bool
}

// Loads configuration from a source into a map
func (loader JsonLoader) Load() (map[string]any, error) {
	f, err := os.Open(loader.path)
	if err != nil {
		if !loader.required {
			return map[string]any{}, nil
		} else {
			return nil, err
		}
	}

	var data map[string]any
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

var _ cfg.Loader = (*JsonLoader)(nil)

func NewLoader(path string, required bool) *JsonLoader {
	return &JsonLoader{
		path:     path,
		required: required,
	}
}
