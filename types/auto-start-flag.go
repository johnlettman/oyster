package types

import (
	"github.com/johnlettman/oyster/types/pseudo"
	"strings"
)

// AutoStartFlag represents a boolean flag used for enabling or disabling auto-start functionality.
//
// For additional information, refer to [Ouster docs: Standby Operating Mode Examples].
//
// Warning:
//
// AutoStartFlag has been deprecated with firmware v2.4 and later.
// Usage of AutoStartFlag in firmware prior to v2.0.0 has unexpected behavior.
//
// [Ouster docs: Standby Operating Mode Examples]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/sensor_operations/sensor-operations.html#standby-operating-mode-examples
type AutoStartFlag pseudo.IntBool

const (
	AutoStartOn  AutoStartFlag = true  // Equivalent to OperatingMode "NORMAL"
	AutoStartOff AutoStartFlag = false // Equivalent to OperatingMode "STANDBY"
)

// String returns the string representation of an AutoStartFlag.
func (a AutoStartFlag) String() string {
	switch a {
	default:
		fallthrough
	case AutoStartOn:
		return "auto start on"
	case AutoStartOff:
		return "auto start off"
	}
}

// GoString returns the Go syntax representation of an AutoStartFlag.
func (a AutoStartFlag) GoString() string {
	switch a {
	default:
		fallthrough
	case AutoStartOn:
		return "AutoStartOn"
	case AutoStartOff:
		return "AutoStartOff"
	}
}

// MarshalText returns the text representation of an AutoStartFlag.
//
// - If the AutoStartFlag is AutoStartOff, it returns "off".
// - Otherwise, it returns "on".
//
// It always returns nil, indicating no error occurred.
func (a AutoStartFlag) MarshalText() ([]byte, error) {
	switch a {
	default:
		fallthrough
	case AutoStartOn:
		return []byte("on"), nil
	case AutoStartOff:
		return []byte("off"), nil
	}
}

// UnmarshalText parses the provided text and assigns the corresponding value to the receiver.
// The method converts the input text to lowercase.
//
//   - If the text is 'off', it assigns AutoStartOff to the receiver.
//   - For any other text, it assigns AutoStartOn to the receiver.
//
// It always returns nil, indicating no error occurred.
func (a *AutoStartFlag) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	default:
		fallthrough
	case "on":
		*a = AutoStartOn
	case "off":
		*a = AutoStartOff
	}

	return nil
}

// MarshalJSON converts AutoStartFlag to JSON format using pseudo.IntBool.MarshalJSON method.
// It returns the JSON bytes and any occurred error.
func (a AutoStartFlag) MarshalJSON() ([]byte, error) {
	return (*pseudo.IntBool)(&a).MarshalJSON()
}

// UnmarshalJSON converts the JSON data into the AutoStartFlag value.
// It leverages the pseudo.IntBool.UnmarshalJSON method to perform the actual unmarshaling.
// It returns any occurred error.
func (a *AutoStartFlag) UnmarshalJSON(data []byte) error {
	return (*pseudo.IntBool)(a).UnmarshalJSON(data)
}
