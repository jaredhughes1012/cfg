package cfgenv

import (
	"os"
	"strings"

	"github.com/jaredhughes1012/cfg"
	"github.com/jaredhughes1012/cfg/internal/mapconvert"
)

type object map[string]any

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

func varToMap(key, val, delim string) object {
	segs := strings.Split(key, delim)
	var m, mprev, mroot map[string]any
	m = make(map[string]any)
	mroot = m

	for _, k := range segs[:len(segs)-1] {
		mprev = m
		m = make(map[string]any)
		mprev[k] = m
	}

	m[segs[len(segs)-1]] = val
	return mroot
}

func prepareVar(prefix, envVar string) (string, string) {
	if len(prefix) > 0 {
		iPrefix := strings.Index(envVar, prefix)
		if iPrefix != 0 {
			return "", ""
		}
		envVar = envVar[iPrefix+len(prefix):]
	}

	ikey := strings.Index(envVar, "=")
	key := envVar[0:ikey]
	val := envVar[ikey+1:]

	return key, val
}

// Loads configuration from a source into a map
func (loader EnvLoader) Load() (map[string]any, error) {
	data := make(map[string]any)

	for _, v := range os.Environ() {
		key, val := prepareVar(loader.opts.Prefix, v)
		if key == "" {
			continue
		}

		envData := varToMap(key, val, loader.opts.Delimiter)
		data = mapconvert.Fold(envData, data)
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
