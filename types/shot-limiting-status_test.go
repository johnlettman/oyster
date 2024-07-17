package types

import (
	"encoding"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestShotLimitingStatus_TextInterfaces(t *testing.T) {
	assert.Implements(t, (*encoding.TextMarshaler)(nil), new(ShotLimitingStatus))
	assert.Implements(t, (*encoding.TextUnmarshaler)(nil), new(ShotLimitingStatus))
}

func TestShotLimitingStatus_StringInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(ShotLimitingStatus))
	assert.Implements(t, (*fmt.GoStringer)(nil), new(ShotLimitingStatus))
}

func TestShotLimitingStatus_String(t *testing.T) {
	type TestCase struct {
		name string
		s    ShotLimitingStatus
		want string
	}

	cases := []TestCase{
		{"ShotLimitingNormal", ShotLimitingNormal, "normal"},
		{"ShotLimitingImminent", ShotLimitingImminent, "imminent"},
		{"ShotLimitingReduction0to10", ShotLimitingReduction0to10, "reduction of 0 to 10%"},
		{"ShotLimitingReduction10to20", ShotLimitingReduction10to20, "reduction of 10 to 20%"},
		{"ShotLimitingReduction20to30", ShotLimitingReduction20to30, "reduction of 20 to 30%"},
		{"ShotLimitingReduction30to40", ShotLimitingReduction30to40, "reduction of 30 to 40%"},
		{"ShotLimitingReduction40to50", ShotLimitingReduction40to50, "reduction of 40 to 50%"},
		{"ShotLimitingReduction50to60", ShotLimitingReduction50to60, "reduction of 50 to 60%"},
		{"ShotLimitingReduction60to70", ShotLimitingReduction60to70, "reduction of 60 to 70%"},
		{"ShotLimitingReduction70to75", ShotLimitingReduction70to75, "reduction of 70 to 75%"},
		{"unknown value", ShotLimitingStatus(math.MaxUint8), "normal"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.s.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestShotLimitingStatus_GoString(t *testing.T) {
	type TestCase struct {
		name string
		s    ShotLimitingStatus
		want string
	}

	cases := []TestCase{
		{"ShotLimitingNormal", ShotLimitingNormal, "ShotLimitingNormal"},
		{"ShotLimitingImminent", ShotLimitingImminent, "ShotLimitingImminent"},
		{"ShotLimitingReduction0to10", ShotLimitingReduction0to10, "ShotLimitingReduction0to10"},
		{"ShotLimitingReduction10to20", ShotLimitingReduction10to20, "ShotLimitingReduction10to20"},
		{"ShotLimitingReduction20to30", ShotLimitingReduction20to30, "ShotLimitingReduction20to30"},
		{"ShotLimitingReduction30to40", ShotLimitingReduction30to40, "ShotLimitingReduction30to40"},
		{"ShotLimitingReduction40to50", ShotLimitingReduction40to50, "ShotLimitingReduction40to50"},
		{"ShotLimitingReduction50to60", ShotLimitingReduction50to60, "ShotLimitingReduction50to60"},
		{"ShotLimitingReduction60to70", ShotLimitingReduction60to70, "ShotLimitingReduction60to70"},
		{"ShotLimitingReduction70to75", ShotLimitingReduction70to75, "ShotLimitingReduction70to75"},
		{"unknown value", ShotLimitingStatus(math.MaxUint8), "ShotLimitingNormal"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.s.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestShotLimitingStatus_MarshalText(t *testing.T) {
	type TestCase struct {
		name string
		s    ShotLimitingStatus
		want string
	}

	cases := []TestCase{
		{"ShotLimitingNormal", ShotLimitingNormal, "NORMAL"},
		{"ShotLimitingImminent", ShotLimitingImminent, "SHOT_LIMITING_IMMINENT"},
		{"ShotLimitingReduction0to10", ShotLimitingReduction0to10, "SHOT_LIMITING_0_TO_10"},
		{"ShotLimitingReduction10to20", ShotLimitingReduction10to20, "SHOT_LIMITING_10_TO_20"},
		{"ShotLimitingReduction20to30", ShotLimitingReduction20to30, "SHOT_LIMITING_20_TO_30"},
		{"ShotLimitingReduction30to40", ShotLimitingReduction30to40, "SHOT_LIMITING_30_TO_40"},
		{"ShotLimitingReduction40to50", ShotLimitingReduction40to50, "SHOT_LIMITING_40_TO_50"},
		{"ShotLimitingReduction50to60", ShotLimitingReduction50to60, "SHOT_LIMITING_50_TO_60"},
		{"ShotLimitingReduction60to70", ShotLimitingReduction60to70, "SHOT_LIMITING_60_TO_70"},
		{"ShotLimitingReduction70to75", ShotLimitingReduction70to75, "SHOT_LIMITING_70_TO_75"},
		{"unknown value", ShotLimitingStatus(math.MaxUint8), "NORMAL"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.s.MarshalText()
			assert.NoError(t, err, "it should never error")
			assert.Equal(t, []byte(c.want), got, "it should return the correct representation")
		})
	}
}

func TestShotLimitingStatus_UnmarshalText(t *testing.T) {
	type TestCase struct {
		name string
		text string
		want ShotLimitingStatus
	}

	cases := []TestCase{
		{"NORMAL", "NORMAL", ShotLimitingNormal},
		{"SHOT_LIMITING_IMMINENT", "SHOT_LIMITING_IMMINENT", ShotLimitingImminent},
		{"SHOT_LIMITING_0_TO_10", "SHOT_LIMITING_0_TO_10", ShotLimitingReduction0to10},
		{"SHOT_LIMITING_10_TO_20", "SHOT_LIMITING_10_TO_20", ShotLimitingReduction10to20},
		{"SHOT_LIMITING_20_TO_30", "SHOT_LIMITING_20_TO_30", ShotLimitingReduction20to30},
		{"SHOT_LIMITING_30_TO_40", "SHOT_LIMITING_30_TO_40", ShotLimitingReduction30to40},
		{"SHOT_LIMITING_40_TO_50", "SHOT_LIMITING_40_TO_50", ShotLimitingReduction40to50},
		{"SHOT_LIMITING_50_TO_60", "SHOT_LIMITING_50_TO_60", ShotLimitingReduction50to60},
		{"SHOT_LIMITING_60_TO_70", "SHOT_LIMITING_60_TO_70", ShotLimitingReduction60to70},
		{"SHOT_LIMITING_70_TO_75", "SHOT_LIMITING_70_TO_75", ShotLimitingReduction70to75},
		{"random text", gofakeit.LoremIpsumSentence(4), ShotLimitingNormal},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var s ShotLimitingStatus
			err := s.UnmarshalText([]byte(c.text))
			assert.NoError(t, err, "it should never error")
			assert.Equal(t, c.want, s, "it should assign the correct value")
		})
	}
}
