package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseMap(t *testing.T) {
	input := map[string]string{"a": "z"}
	got := ReverseMap(input)
	assert.Equal(t, got["z"], "a", "it should reverse the map")
}
