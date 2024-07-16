package types

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullScaleRange_String(t *testing.T) {
	type TestCase struct {
		name string
		r    FullScaleRange
		want string
	}

	cases := []TestCase{
		{"FullScaleRangeNormal", FullScaleRangeNormal, "normal"},
		{"FullScaleRangeExtended", FullScaleRangeExtended, "extended"},
		{"unknown value", FullScaleRangeExtended + 1, "normal"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.r.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestFullScaleRange_GoString(t *testing.T) {
	type TestCase struct {
		name string
		r    FullScaleRange
		want string
	}

	cases := []TestCase{
		{"FullScaleRangeNormal", FullScaleRangeNormal, "FullScaleRangeNormal"},
		{"FullScaleRangeExtended", FullScaleRangeExtended, "FullScaleRangeExtended"},
		{"unknown value", FullScaleRangeExtended + 1, "FullScaleRangeNormal"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.r.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestFullScaleRange_MarshalText(t *testing.T) {
	type TestCase struct {
		name string
		r    FullScaleRange
		want string
	}

	cases := []TestCase{
		{"FullScaleRangeNormal", FullScaleRangeNormal, "NORMAL"},
		{"FullScaleRangeExtended", FullScaleRangeExtended, "EXTENDED"},
		{"unknown value", FullScaleRangeExtended + 1, "NORMAL"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.r.MarshalText()
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, []byte(c.want), got, "it should return the correct representation")
		})
	}
}

func TestFullScaleRange_UnmarshalText(t *testing.T) {
	type TestCase struct {
		name string
		text string
		want FullScaleRange
	}

	cases := []TestCase{
		{"lowercase 'normal'", "normal", FullScaleRangeNormal},
		{"lowercase 'extended'", "extended", FullScaleRangeExtended},
		{"uppercase 'NORMAL'", "NORMAL", FullScaleRangeNormal},
		{"uppercase 'EXTENDED'", "EXTENDED", FullScaleRangeExtended},
		{"random text", gofakeit.LoremIpsumSentence(4), FullScaleRangeNormal},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var a FullScaleRange
			err := a.UnmarshalText([]byte(c.text))
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, c.want, a, "it should assign the correct value")
		})
	}
}
