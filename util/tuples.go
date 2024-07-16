package util

import "github.com/barweiss/go-tuple"

// MapTuples creates a map from a slice of tuples.Pair.
// The keys are the values in V1 field of the pairs,
// and the values are the values in V2 field of the pairs.
// The function returns the created map.
func MapTuples[A comparable, B any](t []tuple.Pair[A, B]) map[A]B {
	m := make(map[A]B)

	for _, pair := range t {
		m[pair.V1] = pair.V2
	}

	return m
}
