package types

// ThermalShutdownStatus represents the state of the Ouster sensor when it reaches
// the maximum operating temperature. This is handled by an independent state machine
// that triggers a ThermalShutdownImminent state (with an OVERTEMP alert) at the sensor's
// maximum temperature. If the sensor stays at this temperature for more than 30 seconds,
// it enters a SHUTDOWN state, issuing alert 0x0100006B and ceasing operation.
// Otherwise, during normal operation the sensor will send ThermalShutdownNormal.
//
// For additional information, refer to [Shot Limiting].
//
// [Shot Limiting]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/sensor_operations/sensor-operations.html#shot-limiting
type ThermalShutdownStatus uint8

const (
	ThermalShutdownNormal   ThermalShutdownStatus = iota // Normal operation of the sensor.
	ThermalShutdownImminent                              // Imminent thermal shutdown due to maximum temperature.
)

// String returns the string representation of a ThermalShutdownStatus value.
// If no match is found, it returns "normal" as the default string representation.
func (t ThermalShutdownStatus) String() string {
	if t == ThermalShutdownImminent {
		return "imminent"
	} else {
		return "normal"
	}
}

// GoString returns the Go syntax representation of a ThermalShutdownStatus value.
// If no match is found, it returns "ThermalShutdownNormal" as the default syntax representation.
func (t ThermalShutdownStatus) GoString() string {
	if t == ThermalShutdownImminent {
		return "ThermalShutdownImminent"
	} else {
		return "ThermalShutdownNormal"
	}
}

// MarshalText returns the text representation of a ThermalShutdownStatus value.
// If no match is found, it returns "NORMAL" as the default text representation.
func (t ThermalShutdownStatus) MarshalText() ([]byte, error) {
	if t == ThermalShutdownImminent {
		return []byte("IMMINENT"), nil
	} else {
		return []byte("NORMAL"), nil
	}
}

// UnmarshalText parses the given text and sets the value of the ThermalShutdownStatus receiver.
//   - If the text value is "IMMINENT", it will be updated to ThermalShutdownImminent.
//   - If the text value is anything other than "IMMINENT", it will be updated to ThermalShutdownNormal.
//
// This method always returns nil.
func (t *ThermalShutdownStatus) UnmarshalText(text []byte) error {
	if string(text) == "IMMINENT" {
		*t = ThermalShutdownImminent
	} else {
		*t = ThermalShutdownNormal
	}

	return nil
}
