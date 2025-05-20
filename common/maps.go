package common

// ContainsKey checks if a map contains a given key.
// K represents the type of the map keys, and must be comparable.
// V represents the type of the map values.
func ContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}
