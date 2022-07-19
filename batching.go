package cfg

type receiver[T any] struct {
	pointer *T
	key     string
	value   T
}

func (r *receiver[T]) execute() {
	*r.pointer = r.value
}

// Stores receivers for multiple configuration values to be bound simultaneously
type Binder struct {
	cfg              *Config
	stringRecievers  []*receiver[string]
	intReceivers     []*receiver[int]
	float64Receivers []*receiver[float64]
}

// Binds a string configuration value. Will be set after the binder function is executed
func (binder *Binder) StringVar(dest *string, key string) {
	binder.stringRecievers = append(binder.stringRecievers, &receiver[string]{
		pointer: dest,
		key:     key,
	})
}

// Binds an integer configuration value. Will be set after the binder function is executed
func (binder *Binder) IntVar(dest *int, key string) {
	binder.intReceivers = append(binder.intReceivers, &receiver[int]{
		pointer: dest,
		key:     key,
	})
}

// Binds a float64 configuration value. Will be set after the binder function is executed
func (binder *Binder) Float64Var(dest *float64, key string) {
	binder.float64Receivers = append(binder.float64Receivers, &receiver[float64]{
		pointer: dest,
		key:     key,
	})
}

func newBinder(cfg *Config) *Binder {
	return &Binder{
		cfg:              cfg,
		stringRecievers:  make([]*receiver[string], 0),
		intReceivers:     make([]*receiver[int], 0),
		float64Receivers: make([]*receiver[float64], 0),
	}
}

func (binder *Binder) execute() error {
	// Get config values
	for _, r := range binder.stringRecievers {
		v, err := binder.cfg.GetString(r.key)
		if err != nil {
			return err
		}
		r.value = v
	}

	for _, r := range binder.intReceivers {
		v, err := binder.cfg.GetInt(r.key)
		if err != nil {
			return err
		}
		r.value = v
	}

	for _, r := range binder.float64Receivers {
		v, err := binder.cfg.GetFloat64(r.key)
		if err != nil {
			return err
		}
		r.value = v
	}

	// Assign all values
	for _, r := range binder.stringRecievers {
		r.execute()
	}

	for _, r := range binder.intReceivers {
		r.execute()
	}

	for _, r := range binder.float64Receivers {
		r.execute()
	}

	return nil
}
