package util

import (
	"github.com/barweiss/go-tuple"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createRandomTuple() tuple.Pair[string, string] {
	return tuple.Pair[string, string]{
		V1: gofakeit.FirstName(),
		V2: gofakeit.LastName(),
	}
}

func TestMapTuples(t *testing.T) {
	var pair tuple.Pair[string, string]

	count := 10
	list := make([]tuple.Pair[string, string], count)
	want := make(map[string]string, count)

	for i := 0; i < count; i++ {
		pair = createRandomTuple()
		list[i] = pair
		want[pair.V1] = pair.V2
	}

	got := MapTuples[string, string](list)
	assert.Equal(t, want, got, "it should map the tuple")
}
