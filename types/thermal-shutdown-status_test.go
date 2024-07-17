package types

import (
	"encoding"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThermalShutdownStatus_TextInterfaces(t *testing.T) {
	assert.Implements(t, (*encoding.TextMarshaler)(nil), new(ThermalShutdownStatus))
	assert.Implements(t, (*encoding.TextUnmarshaler)(nil), new(ThermalShutdownStatus))
}

func TestThermalShutdownStatus_StringInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(ThermalShutdownStatus))
	assert.Implements(t, (*fmt.GoStringer)(nil), new(ThermalShutdownStatus))
}

func TestThermalShutdownStatus_String(t *testing.T) {
	type TestCase struct {
		name string
		t    ThermalShutdownStatus
		want string
	}

	cases := []TestCase{
		{"ThermalShutdownNormal", ThermalShutdownNormal, "normal"},
		{"ThermalShutdownImminent", ThermalShutdownImminent, "imminent"},
		{"unknown value", ThermalShutdownImminent + 1, "normal"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.t.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestThermalShutdownStatus_GoString(t *testing.T) {
	type TestCase struct {
		name string
		t    ThermalShutdownStatus
		want string
	}

	cases := []TestCase{
		{"ThermalShutdownNormal", ThermalShutdownNormal, "ThermalShutdownNormal"},
		{"ThermalShutdownImminent", ThermalShutdownImminent, "ThermalShutdownImminent"},
		{"unknown value", ThermalShutdownImminent + 1, "ThermalShutdownNormal"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.t.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestThermalShutdownStatus_MarshalText(t *testing.T) {
	type TestCase struct {
		name string
		t    ThermalShutdownStatus
		want string
	}

	cases := []TestCase{
		{"ThermalShutdownNormal", ThermalShutdownNormal, "NORMAL"},
		{"ThermalShutdownImminent", ThermalShutdownImminent, "IMMINENT"},
		{"unknown value", ThermalShutdownImminent + 1, "NORMAL"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.t.MarshalText()
			assert.NoError(t, err, "it should never error")
			assert.Equal(t, []byte(c.want), got, "it should return the correct representation")
		})
	}
}

func TestThermalShutdownStatus_UnmarshalText(t *testing.T) {
	type TestCase struct {
		name string
		text string
		want ThermalShutdownStatus
	}

	cases := []TestCase{
		{"text 'NORMAL'", "NORMAL", ThermalShutdownNormal},
		{"text 'IMMINENT'", "IMMINENT", ThermalShutdownImminent},
		{"unknown value", gofakeit.LoremIpsumSentence(4), ThermalShutdownNormal},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var got ThermalShutdownStatus
			err := got.UnmarshalText([]byte(c.text))
			assert.NoError(t, err, "it should never error")
			assert.Equal(t, c.want, got, "it should assign the correct value")
		})
	}
}
