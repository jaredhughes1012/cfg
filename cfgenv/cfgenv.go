package cfgenv

import (
	"os"
	"strings"

	"github.com/jaredhughes1012/cfg"
)

// Special options used to control how environment variable configuration is loaded
type Options struct {
	// If set, only variables that start with the given prefix will be opened
	Prefix string

	// If set, all matched variables will be split using this string and then nested. Used to group
	// variables without having access to nesting like in object-based files
	Delimiter string
}

// Standard options that are used if none is provided
var StandardOptions = Options{
	Prefix:    "",
	Delimiter: "__",
}

// Processes environment variables with a specific pattern and loads them as configuration
type EnvLoader struct {
	opts *Options
}

func fold(from, onto map[string]interface{}) map[string]interface{} {
	return nil
}

func varToMap(envVar, delim string) map[string]interface{} {
	return nil
}

// Loads configuration from a source into a map
func (loader EnvLoader) Load() (map[string]any, error) {
	data := make(map[string]any)

	for _, v := range os.Environ() {
		i := strings.Index(v, loader.opts.Prefix)
		if i >= 0 {
			vTrim := v[i+len(loader.opts.Prefix) : len(v)-1]
			envData := varToMap(vTrim, loader.opts.Delimiter)
			data = fold(envData, data)
		}
	}

	return data, nil
}

var _ cfg.Loader = (*EnvLoader)(nil)

// Creates a new cfg loader designed to load from environment variables
func NewLoader(opts *Options) *EnvLoader {
	return &EnvLoader{
		opts: opts,
	}
}
