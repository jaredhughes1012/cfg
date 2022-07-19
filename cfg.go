package cfg

import (
	"fmt"
	"strconv"

	"github.com/jaredhughes1012/cfg/internal/mapconvert"
)

// Manages loading and access of external configuration data
type Config struct {
	data    map[string]any
	loaders []Loader
}

// Creates a new configuration instance
func New() *Config {
	return &Config{
		data:    nil,
		loaders: make([]Loader, 0),
	}
}

// Adds a new config loader to this configuration instance
func (cfg *Config) Add(l Loader) {
	cfg.loaders = append(cfg.loaders, l)
}

// Loads from all registered loaders. If any loaders fail, will return the error
func (cfg *Config) Load() error {
	data := make(map[string]any)

	for _, loader := range cfg.loaders {
		d, err := loader.Load()
		if err != nil {
			return err
		}

		data = mapconvert.Fold(d, data)
	}

	cfg.data = data
	return nil
}

func (cfg Config) getVal(key string) (any, error) {
	v := cfg.data[key]
	if v == nil {
		return nil, fmt.Errorf("%s not found", key)
	}

	return v, nil
}

// Gets a string config value, returns an error if the value is not found
func (cfg Config) GetString(key string) (string, error) {
	if v, err := cfg.getVal(key); err != nil {
		return "", err
	} else if vStr, ok := v.(string); ok {
		return vStr, nil
	} else {
		return "", fmt.Errorf("key %s does not have a valid string value", key)
	}
}

// Gets a string config value, panics if value is not found
func (cfg Config) MustGetString(key string) string {
	data, err := cfg.GetString(key)
	if err != nil {
		panic(err)
	}

	return data
}

// Gets an integer config value, returns an error if the value is not found
func (cfg Config) GetInt(key string) (int, error) {
	if v, err := cfg.getVal(key); err != nil {
		return 0, err
	} else if vInt, ok := v.(int); ok {
		return vInt, nil
	} else if vStr, ok := v.(string); ok {
		return strconv.Atoi(vStr)
	} else {
		return 0, fmt.Errorf("key %s does not have a valid integer value", key)
	}
}

// Gets a integer config value, panics if value is not found
func (cfg Config) MustGetInt(key string) int {
	data, err := cfg.GetInt(key)
	if err != nil {
		panic(err)
	}

	return data
}

// Gets an integer config value, returns an error if the value is not found
func (cfg Config) GetFloat64(key string) (float64, error) {
	if v, err := cfg.getVal(key); err != nil {
		return 0, err
	} else if vInt, ok := v.(int); ok {
		return float64(vInt), nil
	} else if vFloat32, ok := v.(float32); ok {
		return float64(vFloat32), nil
	} else if vFloat64, ok := v.(float64); ok {
		return vFloat64, nil
	} else if vStr, ok := v.(string); ok {
		return strconv.ParseFloat(vStr, 64)
	} else {
		return 0, fmt.Errorf("key %s does not have a valid float64 value", key)
	}
}

// Gets a integer config value, panics if value is not found
func (cfg Config) MustGetFloat64(key string) float64 {
	data, err := cfg.GetFloat64(key)
	if err != nil {
		panic(err)
	}

	return data
}

// Binds multiple configuration values simultaneously. The binder registers pointers for configuration
// values, which are all resolved and set simultaneously. If any bound values are not found, the error
// will be returned and none of the pointers will be modified
func (cfg Config) Bind(bindFunc func(*Binder)) error {
	binder := newBinder(&cfg)
	bindFunc(binder)
	return binder.execute()
}

// Generic structure used to load configuration from a source into an object. cfg will process and
// flatten this map internally. Loaders should not modify any names of config values, this should
// be manged internally by cfg
type Loader interface {
	// Loads configuration from a source into a map
	Load() (map[string]any, error)
}
