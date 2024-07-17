package types

import (
	"encoding"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestTimeSource_TextInterfaces(t *testing.T) {
	assert.Implements(t, (*encoding.TextMarshaler)(nil), new(TimeSource))
	assert.Implements(t, (*encoding.TextUnmarshaler)(nil), new(TimeSource))
}

func TestTimeSource_StringInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(TimeSource))
	assert.Implements(t, (*fmt.GoStringer)(nil), new(TimeSource))
}

func TestTimeSource_String(t *testing.T) {
	type TestCase struct {
		name string
		ts   TimeSource
		want string
	}

	cases := []TestCase{
		{"TimeFromUnspecified", TimeFromUnspecified, "unspecified"},
		{"TimeFromInternalOscillator", TimeFromInternalOscillator, "time from internal oscillator"},
		{"TimeFromSyncPulseIn", TimeFromSyncPulseIn, "time from sync pulse in"},
		{"TimeFromPTP1588", TimeFromPTP1588, "time from PTP 1588"},
		{"unknown value", TimeSource(math.MaxUint8), "unspecified"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.ts.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestTimeSource_GoString(t *testing.T) {
	type TestCase struct {
		name string
		ts   TimeSource
		want string
	}

	cases := []TestCase{
		{"TimeFromUnspecified", TimeFromUnspecified, "TimeFromUnspecified"},
		{"TimeFromInternalOscillator", TimeFromInternalOscillator, "TimeFromInternalOscillator"},
		{"TimeFromSyncPulseIn", TimeFromSyncPulseIn, "TimeFromSyncPulseIn"},
		{"TimeFromPTP1588", TimeFromPTP1588, "TimeFromPTP1588"},
		{"unknown value", TimeSource(math.MaxUint8), "TimeFromUnspecified"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.ts.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestTimeSource_MarshalText(t *testing.T) {
	type TestCase struct {
		name string
		ts   TimeSource
		want string
	}

	cases := []TestCase{
		{"TimeFromUnspecified", TimeFromUnspecified, "UNSPECIFIED"},
		{"TimeFromInternalOscillator", TimeFromInternalOscillator, "TIME_FROM_INTERNAL_OSC"},
		{"TimeFromSyncPulseIn", TimeFromSyncPulseIn, "TIME_FROM_SYNC_PULSE_IN"},
		{"TimeFromPTP1588", TimeFromPTP1588, "TIME_FROM_PTP_1588"},
		{"unknown value", TimeSource(math.MaxUint8), "UNSPECIFIED"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.ts.MarshalText()
			assert.NoError(t, err, "it should never error")
			assert.Equal(t, []byte(c.want), got, "it should return the correct representation")
		})
	}
}

func TestTimeSource_UnmarshalText(t *testing.T) {
	type TestCase struct {
		name string
		text string
		want TimeSource
	}

	cases := []TestCase{
		{"NORMAL", "UNSPECIFIED", TimeFromUnspecified},
		{"TIME_FROM_INTERNAL_OSC", "TIME_FROM_INTERNAL_OSC", TimeFromInternalOscillator},
		{"TIME_FROM_SYNC_PULSE_IN", "TIME_FROM_SYNC_PULSE_IN", TimeFromSyncPulseIn},
		{"TIME_FROM_PTP_1588", "TIME_FROM_PTP_1588", TimeFromPTP1588},
		{"random text", gofakeit.LoremIpsumSentence(4), TimeFromUnspecified},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var ts TimeSource
			err := ts.UnmarshalText([]byte(c.text))
			assert.NoError(t, err, "it should never error")
			assert.Equal(t, c.want, ts, "it should assign the correct value")
		})
	}
}
