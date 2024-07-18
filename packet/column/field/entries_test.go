package field

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntries_StringInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(Entries))
	assert.Implements(t, (*fmt.GoStringer)(nil), new(Entries))
}

func TestEntries_String(t *testing.T) {
	f := Reflectivity
	fString := fmt.Sprintf("%s:", f)

	s := Structure{
		Type:      TypeUint8,
		Offset:    1,
		ValueMask: 0xF,
		Shift:     32,
	}

	e := Entries{f: s}
	want := fmt.Sprintf("Entries:\n\t%-14s %v,\n", fString, s)
	got := e.String()

	assert.Equal(t, want, got, "it should return the correct representation")
}

func TestEntries_GoString(t *testing.T) {
	f := Reflectivity
	fString := fmt.Sprintf("%#v:", f)

	s := Structure{
		Type:      TypeUint8,
		Offset:    1,
		ValueMask: 0xF,
		Shift:     32,
	}

	e := Entries{f: s}
	want := fmt.Sprintf("field.Entries{\n\t%-14s %#v,\n}", fString, s)
	got := e.GoString()

	assert.Equal(t, want, got, "it should return the correct representation")
}
