package metadata

import (
	"time"
)

// ReflectivityCalibration contains the calibration data field from the sensor.
// Currently, this is solely used for reflectivity calibration details.
//
// Ouster recommends contacting [support@ouster.io] if you have questions on whether
// your sensor is hardware-enabled for calibrated reflectivity.
//
// For additional information, refer to [Ouster docs: Calibrated Reflectivity].
//
// [support@ouster.io]: mailto://support@ouster.io
// [Ouster docs: Calibrated Reflectivity]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/sensor_data/sensor-data.html#calibrated-reflectivity
type ReflectivityCalibration struct {
	// Valid is true if the sensor is factory-calibrated for better accuracy;
	// otherwise, the sensor is using default values and likely has less accuracy.
	Valid bool `json:"valid"`

	// Timestamp is the date and time when the calibration has been performed.
	Timestamp time.Time `json:"timestamp"`
}

// Age returns the duration since the calibration was performed.
func (r *ReflectivityCalibration) Age() time.Duration {
	return time.Since(r.Timestamp)
}

// Calibration contains the calibration status of the sensor.
type Calibration struct {
	Reflectivity *ReflectivityCalibration `json:"reflectivity,omitempty"`
}
