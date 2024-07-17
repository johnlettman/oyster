package types

import (
	"encoding"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNMEABaudRate_TextInterfaces(t *testing.T) {
	assert.Implements(t, (*encoding.TextMarshaler)(nil), new(NMEABaudRate))
	assert.Implements(t, (*encoding.TextUnmarshaler)(nil), new(NMEABaudRate))
}

func TestNMEABaudRate_StringInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(NMEABaudRate))
	assert.Implements(t, (*fmt.GoStringer)(nil), new(NMEABaudRate))
}

func TestNMEABaudRate_String(t *testing.T) {
	type TestCase struct {
		name string
		r    NMEABaudRate
		want string
	}

	cases := []TestCase{
		{"NMEABaudRate9600", NMEABaudRate9600, "9600 baud"},
		{"NMEABaudRate115200", NMEABaudRate115200, "115200 baud"},
		{"NMEABaudRateUnknown", NMEABaudRateUnknown, "unknown"},
		{"unknown value", NMEABaudRate115200 + 1, "unknown"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.r.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestNMEABaudRate_GoString(t *testing.T) {
	type TestCase struct {
		name string
		r    NMEABaudRate
		want string
	}

	cases := []TestCase{
		{"NMEABaudRate9600", NMEABaudRate9600, "NMEABaudRate9600"},
		{"NMEABaudRate115200", NMEABaudRate115200, "NMEABaudRate115200"},
		{"NMEABaudRateUnknown", NMEABaudRateUnknown, "NMEABaudRateUnknown"},
		{"unknown value", NMEABaudRate115200 + 1, "NMEABaudRateUnknown"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.r.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestNMEABaudRate_MarshalText(t *testing.T) {
	type TestCase struct {
		name string
		r    NMEABaudRate
		want string
	}

	cases := []TestCase{
		{"NMEABaudRate9600", NMEABaudRate9600, "BAUD_9600"},
		{"NMEABaudRate115200", NMEABaudRate115200, "BAUD_115200"},
		{"NMEABaudRateUnknown", NMEABaudRateUnknown, "UNKNOWN"},
		{"unknown value", NMEABaudRate115200 + 1, "UNKNOWN"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.r.MarshalText()
			assert.NoError(t, err, "it should never error")
			assert.Equal(t, []byte(c.want), got, "it should return the correct representation")
		})
	}
}

func TestNMEABaudRate_UnmarshalText(t *testing.T) {
	type TestCase struct {
		name string
		text string
		want NMEABaudRate
	}

	cases := []TestCase{
		{"text 'BAUD_9600'", "BAUD_9600", NMEABaudRate9600},
		{"text 'BAUD_115200'", "BAUD_115200", NMEABaudRate115200},
		{"text 'UNKNOWN'", "UNKNOWN", NMEABaudRateUnknown},
		{"unknown value", gofakeit.LoremIpsumSentence(4), NMEABaudRateUnknown},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var r NMEABaudRate
			err := r.UnmarshalText([]byte(c.text))
			assert.NoError(t, err, "it should never error")
			assert.Equal(t, c.want, r, "it should assign the correct value")
		})
	}
}
