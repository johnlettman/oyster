package types

// Polarity represents the polarity of a signal or an electrical voltage.
// It is an enumerated type that provides constants for two different polarities:
//
//   - PolarityActiveLow:
//     the signal is defined as a signal that is
//     true when it is at a high voltage (aka high true)
//   - PolarityActiveHigh:
//     the signal is defined as a signal that is true
//     when it is at a low voltage (aka low true)
//
// For additional information about polarities, refer to [Signal Polarity].
// For additional information regarding polarities as applied to Ouster sensors,
// refer to the following:
//
//   - [nmea_in_polarity]
//   - [sync_pulse_in_polarity]
//   - [sync_pulse_out_polarity]
//
// [Signal Polarity]: https://engineering.purdue.edu/~meyer/DDU270/Refs/Pld/pal_polarity.pdf
// [nmea_in_polarity]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#nmea-in-polarity
// [sync_pulse_in_polarity]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sync-pulse-in-polarity
// [sync_pulse_out_polarity]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sync-pulse-out-polarity
type Polarity int

const (
	PolarityUnspecified Polarity = iota // Fallback value for erroneous Polarity values.
	PolarityActiveLow                   // "Active low," active when the voltage is low.
	PolarityActiveHigh                  // "Active high," active when the voltage is high.
)

// String returns the string representation of a Polarity value.
func (p Polarity) String() string {
	switch p {
	default:
		fallthrough
	case PolarityUnspecified:
		return "unspecified"
	case PolarityActiveLow:
		return "active low"
	case PolarityActiveHigh:
		return "active high"
	}
}

// GoString returns the Go syntax representation of a Polarity value.
func (p Polarity) GoString() string {
	switch p {
	default:
		fallthrough
	case PolarityUnspecified:
		return "PolarityUnspecified"
	case PolarityActiveLow:
		return "PolarityActiveLow"
	case PolarityActiveHigh:
		return "PolarityActiveHigh"
	}
}

// MarshalText returns the text representation of a Polarity value.
func (p Polarity) MarshalText() ([]byte, error) {
	switch p {
	default:
		fallthrough
	case PolarityUnspecified:
		return []byte("UNSPECIFIED"), nil
	case PolarityActiveLow:
		return []byte("ACTIVE_LOW"), nil
	case PolarityActiveHigh:
		return []byte("ACTIVE_HIGH"), nil
	}
}

// UnmarshalText parses the given text and sets the value of the Polarity receiver.
// It expects the text to be one of "UNKNOWN", "ACTIVE_LOW", or "ACTIVE_HIGH".
// If the text is not recognized, it defaults to PolarityUnspecified.
// It returns an error if the text cannot be parsed or if it is an empty slice.
func (p *Polarity) UnmarshalText(text []byte) error {
	switch string(text) {
	default:
		fallthrough
	case "UNSPECIFIED":
		*p = PolarityUnspecified
	case "ACTIVE_LOW":
		*p = PolarityActiveLow
	case "ACTIVE_HIGH":
		*p = PolarityActiveHigh
	}

	return nil
}
