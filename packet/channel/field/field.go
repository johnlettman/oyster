package field

import (
	"github.com/johnlettman/oyster/util"
)

type Field uint8

const (
	UnknownField  Field = 0
	Range         Field = 1
	Range2        Field = 2
	Signal        Field = 3
	Signal2       Field = 4
	Reflectivity  Field = 5
	Reflectivity2 Field = 6
	NearIR        Field = 7
	Flags         Field = 8
	Flags2        Field = 9

	RawHeaders Field = 40
	Raw32Word1 Field = 60
	Raw32Word2 Field = 61
	Raw32Word3 Field = 62
	Raw32Word4 Field = 63
	Raw32Word5 Field = 45
	Raw32Word6 Field = 46
	Raw32Word7 Field = 47
	Raw32Word8 Field = 48
	Raw32Word9 Field = 49

	Custom0 Field = 50
	Custom1 Field = 51
	Custom2 Field = 52
	Custom3 Field = 53
	Custom4 Field = 54
	Custom5 Field = 55
	Custom6 Field = 56
	Custom7 Field = 57
	Custom8 Field = 58
	Custom9 Field = 59

	MaxField Field = 64
)

var (
	// fieldGoStringVK maps Field values to their Go syntax representations.
	fieldGoStringVK = map[Field]string{
		UnknownField:  "UnknownField",
		Range:         "Range",
		Range2:        "Range2",
		Signal:        "Signal",
		Signal2:       "Signal2",
		Reflectivity:  "Reflectivity",
		Reflectivity2: "Reflectivity2",
		NearIR:        "NearIR",
		Flags:         "Flags",
		Flags2:        "Flags2",
		RawHeaders:    "RawHeaders",
		Raw32Word5:    "Raw32Word5",
		Raw32Word6:    "Raw32Word6",
		Raw32Word7:    "Raw32Word7",
		Raw32Word8:    "Raw32Word8",
		Raw32Word9:    "Raw32Word9",
		Custom0:       "Custom0",
		Custom1:       "Custom1",
		Custom2:       "Custom2",
		Custom3:       "Custom3",
		Custom4:       "Custom4",
		Custom5:       "Custom5",
		Custom6:       "Custom6",
		Custom7:       "Custom7",
		Custom8:       "Custom8",
		Custom9:       "Custom9",
		Raw32Word1:    "Raw32Word1",
		Raw32Word2:    "Raw32Word2",
		Raw32Word3:    "Raw32Word3",
		Raw32Word4:    "Raw32Word4",
		MaxField:      "MaxField",
	}

	// fieldTextKV maps Field values to their text representations.
	fieldTextKV = map[Field]string{
		UnknownField:  "UNKNOWN",
		Range:         "RANGE",
		Range2:        "RANGE2",
		Signal:        "SIGNAL",
		Signal2:       "SIGNAL2",
		Reflectivity:  "REFLECTIVITY",
		Reflectivity2: "REFLECTIVITY2",
		NearIR:        "NEAR_IR",
		Flags:         "FLAGS",
		Flags2:        "FLAGS2",
		RawHeaders:    "RAW_HEADERS",
		Custom0:       "CUSTOM0",
		Custom1:       "CUSTOM1",
		Custom2:       "CUSTOM2",
		Custom3:       "CUSTOM3",
		Custom4:       "CUSTOM4",
		Custom5:       "CUSTOM5",
		Custom6:       "CUSTOM6",
		Custom7:       "CUSTOM7",
		Custom8:       "CUSTOM8",
		Custom9:       "CUSTOM9",
		Raw32Word1:    "RAW32_WORD1",
		Raw32Word2:    "RAW32_WORD2",
		Raw32Word3:    "RAW32_WORD3",
		Raw32Word4:    "RAW32_WORD4",
		Raw32Word5:    "RAW32_WORD5",
		Raw32Word6:    "RAW32_WORD6",
		Raw32Word7:    "RAW32_WORD7",
		Raw32Word8:    "RAW32_WORD8",
		Raw32Word9:    "RAW32_WORD9",
	}

	// fieldTextVK maps string representations to Field values.
	fieldTextVK = util.ReverseMap(fieldTextKV)
)

// String returns the string representation of a Field value.
func (f Field) String() string {
	if s, ok := fieldTextKV[f]; ok {
		return s
	}

	return fieldTextKV[UnknownField]
}

// GoString returns the Go syntax representation of a Field value.
func (f Field) GoString() string {
	if value, ok := fieldGoStringVK[f]; ok {
		return value
	}

	return fieldGoStringVK[UnknownField]
}

// MarshalText returns the text representation of a Field value.
func (f Field) MarshalText() ([]byte, error) {
	return []byte(f.String()), nil
}

// UnmarshalText updates the value of a Field pointer based on the text representation.
// It attempts to match the provided text with a corresponding Field value.
//   - If a match is found, the pointer is updated to the matching Field value.
//   - If no match is found, the pointer is updated to the UnknownField value.
//
// This method does not return an error.
func (f *Field) UnmarshalText(data []byte) error {
	if field, ok := fieldTextVK[string(data)]; ok {
		*f = field
	} else {
		*f = UnknownField
	}

	return nil
}
