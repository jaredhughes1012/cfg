package cfg

// Manages loading and access of external configuration data
type Config struct {
	maps    []map[string]any
	flatMap map[string]string
	loaders []Loader
}

// Creates a new configuration instance
func New() *Config {
	return &Config{
		maps:    make([]map[string]any, 0),
		flatMap: map[string]string{},
		loaders: make([]Loader, 0),
	}
}

// Generic structure used to load configuration from a source into an object. cfg will process and
// flatten this map internally. Loaders should not modify any names of config values, this should
// be manged internally by cfg
type Loader interface {
	// Loads configuration from a source into a map
	Load() (map[string]any, error)
}
