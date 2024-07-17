package types

import (
	"encoding"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolarity_StringInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(Polarity))
	assert.Implements(t, (*fmt.GoStringer)(nil), new(Polarity))
}

func TestPolarity_TextInterfaces(t *testing.T) {
	assert.Implements(t, (*encoding.TextMarshaler)(nil), new(Polarity))
	assert.Implements(t, (*encoding.TextUnmarshaler)(nil), new(Polarity))
}

func TestPolarity_String(t *testing.T) {
	type TestCase struct {
		name string
		p    Polarity
		want string
	}

	cases := []TestCase{
		{"PolarityUnspecified", PolarityUnspecified, "unspecified"},
		{"PolarityActiveLow", PolarityActiveLow, "active low"},
		{"PolarityActiveHigh", PolarityActiveHigh, "active high"},
		{"unknown value", PolarityActiveHigh + 1, "unspecified"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.p.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestPolarity_GoString(t *testing.T) {
	type TestCase struct {
		name string
		p    Polarity
		want string
	}

	cases := []TestCase{
		{"PolarityUnspecified", PolarityUnspecified, "PolarityUnspecified"},
		{"PolarityActiveLow", PolarityActiveLow, "PolarityActiveLow"},
		{"PolarityActiveHigh", PolarityActiveHigh, "PolarityActiveHigh"},
		{"unknown value", PolarityActiveHigh + 1, "PolarityUnspecified"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.p.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestPolarity_MarshalText(t *testing.T) {
	type TestCase struct {
		name string
		p    Polarity
		want string
	}

	cases := []TestCase{
		{"PolarityUnspecified", PolarityUnspecified, "UNSPECIFIED"},
		{"PolarityActiveLow", PolarityActiveLow, "ACTIVE_LOW"},
		{"PolarityActiveHigh", PolarityActiveHigh, "ACTIVE_HIGH"},
		{"unknown value", PolarityActiveHigh + 1, "UNSPECIFIED"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.p.MarshalText()
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, c.want, string(got), "it should return the correct representation")
		})
	}
}

func TestPolarity_UnmarshalText(t *testing.T) {
	type TestCase struct {
		name string
		text string
		want Polarity
	}

	cases := []TestCase{
		{"PolarityUnspecified", "UNSPECIFIED", PolarityUnspecified},
		{"PolarityActiveLow", "ACTIVE_LOW", PolarityActiveLow},
		{"PolarityActiveHigh", "ACTIVE_HIGH", PolarityActiveHigh},
		{"unknown value", gofakeit.LoremIpsumSentence(4), PolarityUnspecified},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var got Polarity
			err := got.UnmarshalText([]byte(c.text))
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}
