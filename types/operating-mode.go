package types

// OperatingMode represents the power consumption and activity level of the Ouster sensor.
// It can be either OperatingNormal, which is the default mode where the sensor performs
// its regular operations and consumes standard power, or OperatingStandby, a low-power
// mode that is useful for power, battery, or thermal-conscious applications.
type OperatingMode uint8

const (
	// OperatingNormal is the default operating mode of the Ouster sensor. In this mode,
	// the sensor performs its regular operations and consumes standard power.
	OperatingNormal OperatingMode = iota

	// OperatingStandby is a low-power mode available from firmware version v2.0.0 and onward.
	// It can be used in power, battery, or thermal-conscious applications. Without
	// undergoing a standard operation, the sensor consumes less power in this mode.
	OperatingStandby
)

// String returns the string representation of a OperatingMode value.
func (m OperatingMode) String() string {
	if m == OperatingStandby {
		return "standby"
	} else {
		return "normal"
	}
}

// GoString returns the Go syntax representation of a OperatingMode value.
func (m OperatingMode) GoString() string {
	if m == OperatingStandby {
		return "OperatingStandby"
	} else {
		return "OperatingNormal"
	}
}

// MarshalText returns the text representation of a OperatingMode value.
func (m OperatingMode) MarshalText() ([]byte, error) {
	if m == OperatingStandby {
		return []byte("STANDBY"), nil
	} else {
		return []byte("NORMAL"), nil
	}
}

// UnmarshalText updates the value of an OperatingMode by unmarshaling the provided text.
//   - If the text is "STANDBY", the OperatingMode is set to OperatingStandby.
//   - Otherwise, it is set to OperatingNormal.
//
// This method does not return an error.
func (m *OperatingMode) UnmarshalText(text []byte) error {
	if string(text) == "STANDBY" {
		*m = OperatingStandby
	} else {
		*m = OperatingNormal
	}

	return nil
}
