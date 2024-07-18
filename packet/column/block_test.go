package column

import (
	"fmt"
	"github.com/johnlettman/oyster/packet/column/field"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestBlock_StringInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(Block))
	assert.Implements(t, (*fmt.GoStringer)(nil), new(Block))
}

func TestBlock_String(t *testing.T) {
	f := field.Reflectivity
	v := uint16(math.MaxUint16)
	block := Block{f: v}
	want := fmt.Sprintf("Block:\n\t%-16s %T(%v)\n", f.String()+":", v, v)
	got := block.String()

	assert.Equal(t, want, got, "it should return the correct representation")
}

func TestBlock_GoString(t *testing.T) {
	f := field.Reflectivity
	v := uint16(math.MaxUint16)
	block := Block{f: v}
	want := fmt.Sprintf("field.Block{\n\t%-20s %T(%v),\n}", f.GoString()+":", v, v)
	got := block.GoString()

	assert.Equal(t, want, got, "it should return the correct representation")
}
