package metadata

import "github.com/johnlettman/oyster/types"

// IMUDataFormat represents the format of IMU data from an Ouster sensor.
//
// For additional information, refer to [Ouster docs: imu_data_format].
//
// [Ouster docs: imu_data_format]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/http-api-v1.html#get-api-v1-sensor-metadata-imu-data-format
type IMUDataFormat struct {
	// AccelerationFullScaleRange is the full-scale range of the accelerometer.
	//
	// Settings:
	//   - types.FullScaleRangeNormal (default): digital output X-, Y-, Z-axis with a range fixed at ±2g.
	//   - types.FullScaleRangeExtended: digital-output X-, Y-, Z-axis with an expanded full-scale range of ±16g.
	AccelerationFullScaleRange types.FullScaleRange `json:"accel_fsr"`

	// GyroscopeFullScaleRange is the full-scale range of the gyroscope.
	//
	// Settings:
	//   - types.FullScaleRangeNormal (default): digital output X-, Y-, Z-axis with a range fixed at ±250°/sec.
	//   - types.FullScaleRangeExtended: digital-output X-, Y-, Z-axis with a programmable full-scale range of ±2000°/sec.
	GyroscopeFullScaleRange types.FullScaleRange `json:"gyro_fsr"`
}
