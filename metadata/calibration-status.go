package metadata

import (
	"time"
)

// ReflectivityCalibrationStatus contains the calibration data field from the sensor.
// Currently, this is solely used for reflectivity calibration details.
// For additional information, refer to [Calibrated Reflectivity].
//
// Ouster recommends contacting [Ouster Support] if you have questions on whether
// your sensor is hardware-enabled for calibrated reflectivity.
//
// [Ouster Support]: mailto://support@ouster.io
// [Calibrated Reflectivity]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/sensor_data/sensor-data.html#calibrated-reflectivity
type ReflectivityCalibrationStatus struct {
	// Valid is true if the sensor is factory-calibrated for better accuracy;
	// otherwise, the sensor is using default values and likely has less accuracy.
	Valid bool `json:"valid"`

	// Timestamp is the date and time when the calibration has been performed.
	Timestamp time.Time `json:"timestamp"`
}

// Age returns the duration since the calibration was performed.
func (r *ReflectivityCalibrationStatus) Age() time.Duration {
	return time.Since(r.Timestamp)
}

// CalibrationStatus contains the calibration status of the sensor.
type CalibrationStatus struct {
	Reflectivity *ReflectivityCalibrationStatus `json:"reflectivity,omitempty"`
}
