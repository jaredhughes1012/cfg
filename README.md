# cfg

Configuration management for Go applications

# Get Started
```
go get -u github.com/jaredhughes1012/cfg
```

# Load Configuration

All configuration is loaded into a standardized format and optimized behind the scenes. Config data is not loaded until Load() is called

```
func main() {
    config := cfg.New()

    # Loads all environment variables with the given prefix and splits them by the given delimiter. Uses the standard config if one
    # isn't provided
    config.AddEnvironmentVariables(&config.EnvironmentVariableOptions{
        Prefix: "PREFIX_",
        Delimiter: "__"
    })
    config.AddEnvironmentVariables(nil)

    # Loads data from a JSON file. Does not have a default
    config.AddJsonFile("./some/path")

    if err := config.Load(); err != nil {
        panic(err)
    }
}
```

# Access configuration

Configuration is loaded into a standard format for easy access. Nested data is accessed using a standard delimiter token (environment variables
are nested by splitting on a delimiter). All underscores are removed from names and all letters are lowercased for consistency.

Config allows loading for different types and all types have a "Must" variant that returns an error if the path is not found. If a path isn't found
for the standard accessors, a default value is returned.

```
func main() {
    config := cfg.New()
    config.AddEnvironmentVariables(nil)
    _ = config.Load()

    # Will be empty if value is not found
    strVal := config.GetString("parent:child")

    # Returns an error if not found
    str2, err := config.MustGetString("parent:otherchild")
    if err != nil {
        panic(err)
    }

    intVal := config.GetInt("some:int")
    floatVal := config.GetFloat("some:float")
    
    # Boolean config will evaluate 'true', 'false', '0', or '1' and defaults to false
    boolVal := config.GetBool("some:bool")
}
```