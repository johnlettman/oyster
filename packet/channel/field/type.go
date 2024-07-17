package field

import "strings"

type NativeType interface {
	uint8 | uint16 | uint32 | uint64
}

type Type uint8

const (
	TypeVoid Type = iota
	TypeUint8
	TypeUint16
	TypeUint32
	TypeUint64
)

// String returns the string representation of a Type value.
func (t Type) String() string {
	switch t {
	default:
		fallthrough
	case TypeVoid:
		return "void"
	case TypeUint8:
		return "uint8"
	case TypeUint16:
		return "uint16"
	case TypeUint32:
		return "uint32"
	case TypeUint64:
		return "uint64"
	}
}

// GoString returns the Go syntax representation of a Type value.
func (t Type) GoString() string {
	switch t {
	default:
		fallthrough
	case TypeVoid:
		return "TypeVoid"
	case TypeUint8:
		return "TypeUint8"
	case TypeUint16:
		return "TypeUint16"
	case TypeUint32:
		return "TypeUint32"
	case TypeUint64:
		return "TypeUint64"
	}
}

// MarshalText returns the text representation of a Type value.
func (t Type) MarshalText() ([]byte, error) {
	return []byte(strings.ToUpper(t.String())), nil
}

// UnmarshalText updates the value of Type by unmarshalling the given text.
// It checks if the text matches any string representation in returnOrderTextVK.
//   - If a match is found, the corresponding Type value is assigned.
//   - Otherwise, TypeVoid is assigned.
//
// The function always returns nil.
func (t *Type) UnmarshalText(text []byte) error {
	switch string(text) {
	default:
		fallthrough
	case "VOID":
		*t = TypeVoid
	case "UINT8":
		*t = TypeUint8
	case "UINT16":
		*t = TypeUint16
	case "UINT32":
		*t = TypeUint32
	case "UINT64":
		*t = TypeUint64
	}

	return nil
}

// Size returns the size of the Type in bytes.
func (t Type) Size() int {
	switch t {
	default:
		fallthrough
	case TypeVoid:
		return 0
	case TypeUint8:
		return 1
	case TypeUint16:
		return 2
	case TypeUint32:
		return 4
	case TypeUint64:
		return 8
	}
}

// Mask returns a bitmask based on the size of the Type.
func (t Type) Mask() uint64 {
	return (1 << (t.Size() * 8)) - 1
}
