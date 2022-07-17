package mapconvert

import "fmt"

// Moves all keys from the "from" map into the "onto" map, recursively. Any duplicate keys
// will be overwritten by the "from" map
func Fold(from, onto map[string]any) map[string]any {
	newMap := make(map[string]any)

	for k, v := range onto {
		newMap[k] = v
	}

	for k, v := range from {
		vMap1, ok := v.(map[string]any)
		if ok {
			vMap2, ok := newMap[k].(map[string]any)
			if ok {
				v = Fold(vMap1, vMap2)
			}
		}

		newMap[k] = v
	}

	return newMap
}

func prefixFlatten(target, m map[string]any, prefix, delim string) {
	if prefix != "" {
		prefix = fmt.Sprintf("%s%s", prefix, delim)
	}

	for k, v := range m {
		k := fmt.Sprintf("%s%s", prefix, k)
		if vMap, ok := v.(map[string]any); ok {
			prefixFlatten(target, vMap, k, delim)
		} else {
			target[k] = v
		}
	}
}

// Converts the given map into a single map. All child maps have their keys appended to the
// parent key and separated by a given delimiter e.g. { "child": { "key": "value"}} with
// delim ":" becomes { "child:key": "value" }
func Flatten(m map[string]any, delim string) map[string]any {
	target := make(map[string]any)
	prefixFlatten(target, m, "", delim)

	return target
}

// Converts all keys in the map using the processor function
func ConvertKeys(m map[string]any, converter func(string) string) map[string]any {
	target := make(map[string]any)

	for k, v := range m {
		target[converter(k)] = v
	}

	return target
}
