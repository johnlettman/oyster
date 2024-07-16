package util

// ReverseMap reverses the key-value pairs of the input map and returns a new map.
func ReverseMap[K, V comparable](in map[K]V) map[V]K {
	out := make(map[V]K, len(in))
	for k, v := range in {
		out[v] = k
	}
	return out
}
