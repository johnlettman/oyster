package types

import (
	"strings"
)

// FullScaleRange represents whether modification of the onboard gyroscope or accelerometer
// has been enabled with an extended programmable scale.
// For additional information, refer to [gyro_fsr] and [accel_fsr].
//
// [gyro_fsr]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#gyro-fsr
// [accel_fsr]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#accel-fsr
type FullScaleRange uint8

const (
	FullScaleRangeNormal   FullScaleRange = iota // Normal range.
	FullScaleRangeExtended                       // Extended programmable scale.
)

// String returns the string representation of a FullScaleRange value.
func (r FullScaleRange) String() string {
	switch r {
	default:
		fallthrough
	case FullScaleRangeNormal:
		return "normal"
	case FullScaleRangeExtended:
		return "extended"
	}
}

// GoString returns the Go syntax representation of a FullScaleRange value.
func (r FullScaleRange) GoString() string {
	switch r {
	default:
		fallthrough
	case FullScaleRangeNormal:
		return "FullScaleRangeNormal"
	case FullScaleRangeExtended:
		return "FullScaleRangeExtended"
	}
}

// MarshalText returns the text representation of an FullScaleRange.
//   - If the FullScaleRange is FullScaleRangeExtended, it returns "EXTENDED";
//   - otherwise, it returns "NORMAL".
//
// It always returns nil error, indicating no error occurred.
func (r FullScaleRange) MarshalText() ([]byte, error) {
	switch r {
	default:
		fallthrough
	case FullScaleRangeNormal:
		return []byte("NORMAL"), nil
	case FullScaleRangeExtended:
		return []byte("EXTENDED"), nil
	}
}

// UnmarshalText parses the provided text and assigns the corresponding value to the receiver.
// The method converts the input text to uppercase.
//   - If the text is 'EXTENDED', it assigns FullScaleRangeNormal to the receiver.
//   - For any other text, it assigns AutoStartOn to the receiver.
//
// It always returns nil error, indicating no error occurred.
func (r *FullScaleRange) UnmarshalText(text []byte) error {
	switch strings.ToUpper(string(text)) {
	default:
		fallthrough
	case "NORMAL":
		*r = FullScaleRangeNormal
	case "EXTENDED":
		*r = FullScaleRangeExtended
	}

	return nil
}
