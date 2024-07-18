package field

import (
	"encoding"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestField_StringInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(Field))
	assert.Implements(t, (*fmt.GoStringer)(nil), new(Field))
}

func TestField_TextInterfaces(t *testing.T) {
	assert.Implements(t, (*encoding.TextMarshaler)(nil), new(Field))
	assert.Implements(t, (*encoding.TextUnmarshaler)(nil), new(Field))
}

func TestField_String(t *testing.T) {
	type TestCase struct {
		name string
		f    Field
		want string
	}

	cases := []TestCase{
		{"UnknownField", UnknownField, "UNKNOWN"},
		{"Range", Range, "RANGE"},
		{"Range2", Range2, "RANGE2"},
		{"Signal", Signal, "SIGNAL"},
		{"Signal2", Signal2, "SIGNAL2"},
		{"unknown value", Field(math.MaxUint8), "UNKNOWN"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.f.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestField_GoString(t *testing.T) {
	type TestCase struct {
		name string
		f    Field
		want string
	}

	cases := []TestCase{
		{"UnknownField", UnknownField, "UnknownField"},
		{"Range", Range, "Range"},
		{"Range2", Range2, "Range2"},
		{"Signal", Signal, "Signal"},
		{"Signal2", Signal2, "Signal2"},
		{"unknown value", Field(math.MaxUint8), "UnknownField"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.f.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestField_MarshalText(t *testing.T) {
	type TestCase struct {
		name string
		f    Field
		want string
	}

	cases := []TestCase{
		{"UnknownField", UnknownField, "UNKNOWN"},
		{"Range", Range, "RANGE"},
		{"Range2", Range2, "RANGE2"},
		{"Signal", Signal, "SIGNAL"},
		{"Signal2", Signal2, "SIGNAL2"},
		{"unknown value", Field(math.MaxUint8), "UNKNOWN"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.f.MarshalText()
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, c.want, string(got), "it should return the correct representation")
		})
	}
}

func TestField_UnmarshalText(t *testing.T) {
	type TestCase struct {
		name string
		text string
		want Field
	}

	cases := []TestCase{
		{"UNKNOWN", "UNKNOWN", UnknownField},
		{"RANGE", "RANGE", Range},
		{"RANGE2", "RANGE2", Range2},
		{"SIGNAL", "SIGNAL", Signal},
		{"SIGNAL2", "SIGNAL2", Signal2},
		{"unknown text", gofakeit.LoremIpsumSentence(4), UnknownField},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var got Field
			err := got.UnmarshalText([]byte(c.text))
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, c.want, got, "it should unmarshal the JSON correctly")
		})
	}
}
