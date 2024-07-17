package types

import (
	"encoding"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperatingMode_StringInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(OperatingMode))
	assert.Implements(t, (*fmt.GoStringer)(nil), new(OperatingMode))
}

func TestOperatingMode_TextInterfaces(t *testing.T) {
	assert.Implements(t, (*encoding.TextMarshaler)(nil), new(OperatingMode))
	assert.Implements(t, (*encoding.TextUnmarshaler)(nil), new(OperatingMode))
}

func TestOperatingMode_String(t *testing.T) {
	type TestCase struct {
		name string
		om   OperatingMode
		want string
	}

	cases := []TestCase{
		{"OperatingNormal", OperatingNormal, "normal"},
		{"OperatingStandby", OperatingStandby, "standby"},
		{"unknown value", OperatingStandby + 1, "normal"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.om.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestOperatingMode_GoString(t *testing.T) {
	type TestCase struct {
		name string
		om   OperatingMode
		want string
	}

	cases := []TestCase{
		{"OperatingNormal", OperatingNormal, "OperatingNormal"},
		{"OperatingStandby", OperatingStandby, "OperatingStandby"},
		{"unknown value", OperatingStandby + 1, "OperatingNormal"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.om.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestOperatingMode_MarshalText(t *testing.T) {
	type TestCase struct {
		name string
		om   OperatingMode
		want string
	}

	cases := []TestCase{
		{"OperatingNormal", OperatingNormal, "NORMAL"},
		{"OperatingStandby", OperatingStandby, "STANDBY"},
		{"unknown value", OperatingStandby + 1, "NORMAL"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.om.MarshalText()
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, c.want, string(got), "it should return the correct representation")
		})
	}
}

func TestOperatingMode_UnmarshalText(t *testing.T) {
	type TestCase struct {
		name string
		text string
		want OperatingMode
	}

	cases := []TestCase{
		{"PolarityUnspecified", "NORMAL", OperatingNormal},
		{"PolarityActiveLow", "STANDBY", OperatingStandby},
		{"unknown value", gofakeit.LoremIpsumSentence(4), OperatingNormal},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var got OperatingMode
			err := got.UnmarshalText([]byte(c.text))
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}
