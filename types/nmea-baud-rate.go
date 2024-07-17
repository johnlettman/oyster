package types

// NMEABaudRate represents the expected baud rate the sensor uses to decode
// NMEA UART input $GPRMC messages.
// For additional information, refer to [nmea_baud_rate].
//
// [nmea_baud_rate]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#nmea-baud-rate
type NMEABaudRate int

const (
	NMEABaudRateUnknown NMEABaudRate = iota // An unknown or undefined baud rate.
	NMEABaudRate9600                        // 9600 baud for NMEA UART.
	NMEABaudRate115200                      // 115200 baud for NMEA UART.
)

// String returns the string representation of a NMEABaudRate value.
func (r NMEABaudRate) String() string {
	switch r {
	default:
		fallthrough
	case NMEABaudRateUnknown:
		return "unknown"
	case NMEABaudRate9600:
		return "9600 baud"
	case NMEABaudRate115200:
		return "115200 baud"
	}
}

// GoString returns the Go syntax representation of a NMEABaudRate value.
func (r NMEABaudRate) GoString() string {
	switch r {
	default:
		fallthrough
	case NMEABaudRateUnknown:
		return "NMEABaudRateUnknown"
	case NMEABaudRate9600:
		return "NMEABaudRate9600"
	case NMEABaudRate115200:
		return "NMEABaudRate115200"
	}
}

// MarshalText returns the text representation of a NMEABaudRate value.
func (r NMEABaudRate) MarshalText() ([]byte, error) {
	switch r {
	default:
		fallthrough
	case NMEABaudRateUnknown:
		return []byte("UNKNOWN"), nil
	case NMEABaudRate9600:
		return []byte("BAUD_9600"), nil
	case NMEABaudRate115200:
		return []byte("BAUD_115200"), nil
	}
}

// UnmarshalText updates the value of NMEABaudRate based on the input text.
// It maps the input text to the corresponding NMEABaudRate value.
// If the input text is not recognized, NMEABaudRateUnknown is set.
func (r *NMEABaudRate) UnmarshalText(text []byte) error {
	switch string(text) {
	default:
		*r = NMEABaudRateUnknown
	case "BAUD_9600":
		*r = NMEABaudRate9600
	case "BAUD_115200":
		*r = NMEABaudRate115200
	}

	return nil
}
