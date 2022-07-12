package cfgjson

import (
	"encoding/json"
	"os"

	"github.com/jaredhughes1012/cfg"
)

// Config loader designed to load from JSON files
type JsonLoader struct {
	path string
}

// Loads configuration from a source into a map
func (loader JsonLoader) Load() (map[string]any, error) {
	f, err := os.Open(loader.path)
	if err != nil {
		return nil, err
	}

	var data map[string]any
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

var _ cfg.Loader = (*JsonLoader)(nil)

func NewLoader(path string) *JsonLoader {
	return &JsonLoader{
		path: path,
	}
}
