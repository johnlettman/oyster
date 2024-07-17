package field

import (
	"encoding"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestType_StringInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(Type))
	assert.Implements(t, (*fmt.GoStringer)(nil), new(Type))
}

func TestType_TextInterfaces(t *testing.T) {
	assert.Implements(t, (*encoding.TextMarshaler)(nil), new(Type))
	assert.Implements(t, (*encoding.TextUnmarshaler)(nil), new(Type))
}

func TestType_String(t *testing.T) {
	type TestCase struct {
		name string
		t    Type
		want string
	}

	cases := []TestCase{
		{"TypeVoid", TypeVoid, "void"},
		{"TypeUint8", TypeUint8, "uint8"},
		{"TypeUint16", TypeUint16, "uint16"},
		{"TypeUint32", TypeUint32, "uint32"},
		{"TypeUint64", TypeUint64, "uint64"},
		{"unknown value", math.MaxUint8, "void"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.t.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestType_GoString(t *testing.T) {
	type TestCase struct {
		name string
		t    Type
		want string
	}

	cases := []TestCase{
		{"TypeVoid", TypeVoid, "TypeVoid"},
		{"TypeUint8", TypeUint8, "TypeUint8"},
		{"TypeUint16", TypeUint16, "TypeUint16"},
		{"TypeUint32", TypeUint32, "TypeUint32"},
		{"TypeUint64", TypeUint64, "TypeUint64"},
		{"unknown value", math.MaxUint8, "TypeVoid"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.t.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestType_MarshalText(t *testing.T) {
	type TestCase struct {
		name string
		t    Type
		want string
	}

	cases := []TestCase{
		{"TypeVoid", TypeVoid, "VOID"},
		{"TypeUint8", TypeUint8, "UINT8"},
		{"TypeUint16", TypeUint16, "UINT16"},
		{"TypeUint32", TypeUint32, "UINT32"},
		{"TypeUint64", TypeUint64, "UINT64"},
		{"unknown value", math.MaxUint8, "VOID"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.t.MarshalText()
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, c.want, string(got), "it should return the correct representation")
		})
	}
}

func TestType_UnmarshalText(t *testing.T) {
	type TestCase struct {
		name string
		text string
		want Type
	}

	cases := []TestCase{
		{"TypeVoid", "VOID", TypeVoid},
		{"TypeUint8", "UINT8", TypeUint8},
		{"TypeUint16", "UINT16", TypeUint16},
		{"TypeUint32", "UINT32", TypeUint32},
		{"TypeUint64", "UINT64", TypeUint64},
		{"unknown text", gofakeit.LoremIpsumSentence(4), TypeVoid},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var got Type
			err := got.UnmarshalText([]byte(c.text))
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, c.want, got, "it should unmarshal the JSON correctly")
		})
	}
}

func TestType_Size(t *testing.T) {
	type TestCase struct {
		name string
		t    Type
		want int
	}

	cases := []TestCase{
		{"TypeVoid", TypeVoid, 0},
		{"TypeUint8", TypeUint8, 1},
		{"TypeUint16", TypeUint16, 2},
		{"TypeUint32", TypeUint32, 4},
		{"TypeUint64", TypeUint64, 8},
		{"unknown value", math.MaxUint8, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.t.Size()
			assert.Equal(t, c.want, got, "it should return the correct size")
		})
	}
}

func TestType_Mask(t *testing.T) {
	type TestCase struct {
		name string
		t    Type
		want uint64
	}

	cases := []TestCase{
		{"TypeVoid", TypeVoid, 0},
		{"TypeUint8", TypeUint8, math.MaxUint8},
		{"TypeUint16", TypeUint16, math.MaxUint16},
		{"TypeUint32", TypeUint32, math.MaxUint32},
		{"TypeUint64", TypeUint64, math.MaxUint64},
		{"unknown value", math.MaxUint8, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.t.Mask()
			assert.Equal(t, c.want, got, "it should return the correct mask")
		})
	}
}
