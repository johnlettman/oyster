package types

import (
	"encoding"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLIDARMode_TextInterfaces(t *testing.T) {
	assert.Implements(t, (*encoding.TextMarshaler)(nil), new(LIDARMode))
	assert.Implements(t, (*encoding.TextUnmarshaler)(nil), new(LIDARMode))
}

func TestLIDARMode_String(t *testing.T) {
	type TestCase struct {
		name string
		m    LIDARMode
		want string
	}

	cases := []TestCase{
		{"LIDARModeUnknown", LidarModeUnknown, "unknown"},
		{"LidarMode512x10", LidarMode512x10, "512x10"},
		{"LidarMode512x20", LidarMode512x20, "512x20"},
		{"LidarMode1024x10", LidarMode1024x10, "1024x10"},
		{"LidarMode1024x20", LidarMode1024x20, "1024x20"},
		{"LidarMode2048x10", LidarMode2048x10, "2048x10"},
		{"LidarMode4096x5", LidarMode4096x5, "4096x5"},
		{"unknown value", LidarMode4096x5 + 1, "unknown"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.m.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestLIDARMode_GoString(t *testing.T) {
	type TestCase struct {
		name string
		m    LIDARMode
		want string
	}

	cases := []TestCase{
		{"LIDARModeUnknown", LidarModeUnknown, "LidarModeUnknown"},
		{"LidarMode512x10", LidarMode512x10, "LidarMode512x10"},
		{"LidarMode512x20", LidarMode512x20, "LidarMode512x20"},
		{"LidarMode1024x10", LidarMode1024x10, "LidarMode1024x10"},
		{"LidarMode1024x20", LidarMode1024x20, "LidarMode1024x20"},
		{"LidarMode2048x10", LidarMode2048x10, "LidarMode2048x10"},
		{"LidarMode4096x5", LidarMode4096x5, "LidarMode4096x5"},
		{"unknown value", LidarMode4096x5 + 1, "LidarModeUnknown"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.m.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestLIDARMode_MarshalText(t *testing.T) {
	type TestCase struct {
		name string
		m    LIDARMode
		want string
	}

	cases := []TestCase{
		{"LIDARModeUnknown", LidarModeUnknown, "unknown"},
		{"LidarMode512x10", LidarMode512x10, "512x10"},
		{"LidarMode512x20", LidarMode512x20, "512x20"},
		{"LidarMode1024x10", LidarMode1024x10, "1024x10"},
		{"LidarMode1024x20", LidarMode1024x20, "1024x20"},
		{"LidarMode2048x10", LidarMode2048x10, "2048x10"},
		{"LidarMode4096x5", LidarMode4096x5, "4096x5"},
		{"unknown value", LidarMode4096x5 + 1, ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.m.MarshalText()
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, []byte(c.want), got, "it should return the correct representation")
		})
	}
}

func TestLIDARMode_UnmarshalText(t *testing.T) {
	type TestCase struct {
		name string
		text string
		want LIDARMode
	}

	cases := []TestCase{
		{"LIDARModeUnknown", "unknown", LidarModeUnknown},
		{"LidarMode512x10", "512x10", LidarMode512x10},
		{"LidarMode512x20", "512x20", LidarMode512x20},
		{"LidarMode1024x10", "1024x10", LidarMode1024x10},
		{"LidarMode1024x20", "1024x20", LidarMode1024x20},
		{"LidarMode2048x10", "2048x10", LidarMode2048x10},
		{"LidarMode4096x5", "4096x5", LidarMode4096x5},
		{"random text", gofakeit.LoremIpsumSentence(4), LidarModeUnknown},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var m LIDARMode
			err := m.UnmarshalText([]byte(c.text))
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, c.want, m, "it should assign the correct value")
		})
	}
}

func TestLIDARMode_Columns(t *testing.T) {
	type TestCase struct {
		name string
		m    LIDARMode
		want int
	}

	cases := []TestCase{
		{"LIDARModeUnknown", LidarModeUnknown, 0},
		{"LIDARMode512x10", LidarMode512x10, 512},
		{"LIDARMode512x20", LidarMode512x20, 512},
		{"LIDARMode1024x10", LidarMode1024x10, 1024},
		{"LIDARMode1024x20", LidarMode1024x20, 1024},
		{"LIDARMode2048x10", LidarMode2048x10, 2048},
		{"LIDARMode4096x5", LidarMode4096x5, 4096},
		{"unknown value", LidarMode4096x5 + 1, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.m.Columns()
			assert.Equal(t, c.want, got, "it should return the correct number of columns")
		})
	}
}

func TestLIDARMode_Frequency(t *testing.T) {
	type TestCase struct {
		name string
		m    LIDARMode
		want int
	}

	cases := []TestCase{
		{"LIDARModeUnknown", LidarModeUnknown, 0},
		{"LIDARMode512x10", LidarMode512x10, 10},
		{"LIDARMode512x20", LidarMode512x20, 20},
		{"LIDARMode1024x10", LidarMode1024x10, 10},
		{"LIDARMode1024x20", LidarMode1024x20, 20},
		{"LIDARMode2048x10", LidarMode2048x10, 10},
		{"LIDARMode4096x5", LidarMode4096x5, 5},
		{"unknown value", LidarMode4096x5 + 1, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.m.Frequency()
			assert.Equal(t, c.want, got, "it should return the correct number of columns")
		})
	}
}
