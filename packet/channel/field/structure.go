package field

import (
	"fmt"
	"github.com/johnlettman/oyster/util"
	"math/bits"
	"strings"
)

// Structure represents a data structure with its type, offset, mask, and shift.
type Structure struct {
	Type      Type
	Offset    int
	ValueMask uint64
	Shift     int
}

// String returns the string representation of a Structure value.
func (s Structure) String() string {
	sb := new(strings.Builder)
	sb.WriteString(fmt.Sprintf("{Type: %s, Offset: %d", s.Type.String(), s.Offset))

	if s.ValueMask != 0 {
		sb.WriteString(fmt.Sprintf(", ValueMask: 0x%.16X", s.ValueMask))
	}

	if s.Shift != 0 {
		sb.WriteString(fmt.Sprintf(", Shift: %d", s.Shift))
	}

	sb.WriteRune('}')
	return sb.String()
}

// GoString returns the Go syntax representation of a Structure value.
func (s Structure) GoString() string {
	sb := new(strings.Builder)
	sb.WriteString(fmt.Sprintf("channel.Structure{Type: %s, Offset: %d", s.Type.GoString(), s.Offset))

	if s.ValueMask != 0 {
		sb.WriteString(fmt.Sprintf(", ValueMask: 0x%.16X", s.ValueMask))
	}

	if s.Shift != 0 {
		sb.WriteString(fmt.Sprintf(", Shift: %d", s.Shift))
	}

	sb.WriteRune('}')
	return sb.String()
}

// Mask returns the masked value based on the Type, ValueMask, and Shift fields of the Structure.
//   - If the ValueMask is zero, it uses the mask derived from the Type field.
//   - If the Shift field is greater than zero, it right-shifts the mask by the Shift amount.
//   - If the Shift field is less than zero, it left-shifts the mask by the absolute value of the Shift amount.
//
// Finally, it performs a bitwise AND operation between the shifted mask and the type-specific mask.
func (s Structure) Mask() uint64 {
	typeMask := s.Type.Mask()

	mask := s.ValueMask
	if mask == 0 {
		mask = typeMask
	}

	if s.Shift > 0 {
		mask >>= uint64(s.Shift)
	} else if s.Shift < 0 {
		mask <<= uint64(util.Abs(s.Shift))
	}

	return mask & typeMask
}

// MaskBits returns the number of bits set to 1 in the masked value obtained from the Mask method of the Structure.
func (s Structure) MaskBits() int {
	return bits.OnesCount64(s.Mask())
}
