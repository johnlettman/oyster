package field

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestStructure_StringInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(Structure))
	assert.Implements(t, (*fmt.GoStringer)(nil), new(Structure))
}

func TestStructure_String(t *testing.T) {
	type TestCase struct {
		name string
		s    Structure
		want string
	}

	cases := []TestCase{
		{
			name: "full",
			s:    Structure{Type: TypeUint8, Offset: 1, ValueMask: 0xFF, Shift: 32},
			want: "{Type: uint8, Offset: 1, ValueMask: 0x00000000000000FF, Shift: 32}",
		},
		{
			name: "no Shift",
			s:    Structure{Type: TypeUint8, Offset: 1, ValueMask: 0xFF, Shift: 0},
			want: "{Type: uint8, Offset: 1, ValueMask: 0x00000000000000FF}",
		},
		{
			name: "no ValueMask",
			s:    Structure{Type: TypeUint8, Offset: 1, Shift: 32},
			want: "{Type: uint8, Offset: 1, Shift: 32}",
		},
		{
			name: "no Shift no ValueMask",
			s:    Structure{Type: TypeUint8, Offset: 1},
			want: "{Type: uint8, Offset: 1}",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.s.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})

	}
}

func TestStructure_GoString(t *testing.T) {
	type TestCase struct {
		name string
		s    Structure
		want string
	}

	cases := []TestCase{
		{
			name: "full",
			s:    Structure{Type: TypeUint8, Offset: 1, ValueMask: 0xFF, Shift: 32},
			want: "channel.Structure{Type: TypeUint8, Offset: 1, ValueMask: 0x00000000000000FF, Shift: 32}",
		},
		{
			name: "no Shift",
			s:    Structure{Type: TypeUint8, Offset: 1, ValueMask: 0xFF, Shift: 0},
			want: "channel.Structure{Type: TypeUint8, Offset: 1, ValueMask: 0x00000000000000FF}",
		},
		{
			name: "no ValueMask",
			s:    Structure{Type: TypeUint8, Offset: 1, Shift: 32},
			want: "channel.Structure{Type: TypeUint8, Offset: 1, Shift: 32}",
		},
		{
			name: "no Shift no ValueMask",
			s:    Structure{Type: TypeUint8, Offset: 1},
			want: "channel.Structure{Type: TypeUint8, Offset: 1}",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.s.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestStructure_Mask(t *testing.T) {
	type TestCase struct {
		name string
		s    Structure
		want uint64
	}

	cases := []TestCase{
		{
			name: "normal",
			s:    Structure{Type: TypeUint8, Offset: 1, ValueMask: 0xFF, Shift: 0},
			want: 0xFF,
		},
		{
			name: "ValueMask greater than type",
			s:    Structure{Type: TypeUint8, Offset: 1, ValueMask: 0xFFFFFFFF, Shift: 0},
			want: 0xFF,
		},
		{
			name: "positive shift",
			s:    Structure{Type: TypeUint8, Offset: 1, ValueMask: 0xF0, Shift: 4},
			want: 0x0F,
		},
		{
			name: "negative shift",
			s:    Structure{Type: TypeUint8, Offset: 1, ValueMask: 0x0F, Shift: -4},
			want: 0xF0,
		},
		{
			name: "no ValueMask",
			s:    Structure{Type: TypeUint8, Offset: 1, ValueMask: 0, Shift: 0},
			want: math.MaxUint8,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.s.Mask()
			assert.Equal(t, c.want, got, "it should return the correct mask")
		})
	}
}

func TestStructure_MaskBits(t *testing.T) {
	type TestCase struct {
		name string
		s    Structure
		want int
	}

	cases := []TestCase{
		{
			name: "normal count",
			s:    Structure{Type: TypeUint8, Offset: 1, ValueMask: 0b00001111, Shift: 0},
			want: 4,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.s.MaskBits()
			assert.Equal(t, c.want, got, "it should return the correct mask bit count")
		})
	}
}
