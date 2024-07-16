package types

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMultipurposeIOMode_String(t *testing.T) {
	type TestCase struct {
		name string
		mode MultipurposeIOMode
		want string
	}

	cases := []TestCase{
		{"MultipurposeOff", MultipurposeOff, "off"},
		{"MultipurposeInputNMEAUART", MultipurposeInputNMEAUART, "input from NMEA UART"},
		{"MultipurposeOutputFromInternalOscillator", MultipurposeOutputFromInternalOscillator, "output from Internal Oscillator"},
		{"MultipurposeOutputFromSyncPulseIn", MultipurposeOutputFromSyncPulseIn, "output from Sync Pulse in"},
		{"MultipurposeOutputFromPTP1588", MultipurposeOutputFromPTP1588, "output from PTP 1588"},
		{"MultipurposeOutputFromEncoderAngle", MultipurposeOutputFromEncoderAngle, "output from encoder angle"},
		{"unknown value", math.MaxUint8, "off"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.mode.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestMultipurposeIOMode_GoString(t *testing.T) {
	type TestCase struct {
		name string
		mode MultipurposeIOMode
		want string
	}

	cases := []TestCase{
		{"MultipurposeOff", MultipurposeOff, "MultipurposeOff"},
		{"MultipurposeInputNMEAUART", MultipurposeInputNMEAUART, "MultipurposeInputNMEAUART"},
		{"MultipurposeOutputFromInternalOscillator", MultipurposeOutputFromInternalOscillator, "MultipurposeOutputFromInternalOscillator"},
		{"MultipurposeOutputFromSyncPulseIn", MultipurposeOutputFromSyncPulseIn, "MultipurposeOutputFromSyncPulseIn"},
		{"MultipurposeOutputFromPTP1588", MultipurposeOutputFromPTP1588, "MultipurposeOutputFromPTP1588"},
		{"MultipurposeOutputFromEncoderAngle", MultipurposeOutputFromEncoderAngle, "MultipurposeOutputFromEncoderAngle"},
		{"unknown value", math.MaxUint8, "MultipurposeOff"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.mode.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestMultipurposeIOMode_MarshalText(t *testing.T) {
	type TestCase struct {
		name string
		mode MultipurposeIOMode
		want string
	}

	cases := []TestCase{
		{"MultipurposeOff", MultipurposeOff, "OFF"},
		{"MultipurposeInputNMEAUART", MultipurposeInputNMEAUART, "INPUT_NMEA_UART"},
		{"MultipurposeOutputFromInternalOscillator", MultipurposeOutputFromInternalOscillator, "OUTPUT_FROM_INTERNAL_OSC"},
		{"MultipurposeOutputFromSyncPulseIn", MultipurposeOutputFromSyncPulseIn, "OUTPUT_FROM_SYNC_PULSE_IN"},
		{"MultipurposeOutputFromPTP1588", MultipurposeOutputFromPTP1588, "OUTPUT_FROM_PTP_1588"},
		{"MultipurposeOutputFromEncoderAngle", MultipurposeOutputFromEncoderAngle, "OUTPUT_FROM_ENCODER_ANGLE"},
		{"unknown value", math.MaxUint8, "OFF"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.mode.MarshalText()
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, []byte(c.want), got, "it should return the correct representation")
		})
	}
}

func TestMultipurposeIOMode_UnmarshalText(t *testing.T) {
	type TestCase struct {
		name string
		text string
		want MultipurposeIOMode
	}

	cases := []TestCase{
		{"MultipurposeOff", "OFF", MultipurposeOff},
		{"MultipurposeInputNMEAUART", "INPUT_NMEA_UART", MultipurposeInputNMEAUART},
		{"MultipurposeOutputFromInternalOscillator", "OUTPUT_FROM_INTERNAL_OSC", MultipurposeOutputFromInternalOscillator},
		{"MultipurposeOutputFromSyncPulseIn", "OUTPUT_FROM_SYNC_PULSE_IN", MultipurposeOutputFromSyncPulseIn},
		{"MultipurposeOutputFromPTP1588", "OUTPUT_FROM_PTP_1588", MultipurposeOutputFromPTP1588},
		{"MultipurposeOutputFromEncoderAngle", "OUTPUT_FROM_ENCODER_ANGLE", MultipurposeOutputFromEncoderAngle},
		{"unknown text", gofakeit.LoremIpsumSentence(4), MultipurposeOff},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var m MultipurposeIOMode
			err := m.UnmarshalText([]byte(c.text))
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, c.want, m, "it should return the correct representation")
		})
	}
}
