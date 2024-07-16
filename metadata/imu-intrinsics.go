package metadata

// IMUIntrinsics represents the intrinsic parameters of the IMU (Inertial Measurement Unit).
//
// For additional information, refer to [Ouster docs: imu_intrinsics].
//
// [Ouster docs: imu_intrinsics]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/http-api-v1.html#get-api-v1-sensor-metadata-imu-intrinsics
type IMUIntrinsics struct {
	// IMUToSensorTransform is the transformation matrix from the IMU (Inertial Measurement Unit) coordinate frame to the sensor coordinate frame.
	// It is a 4x4 matrix stored in row-major order in a 16-element float64 array.
	IMUToSensorTransform [16]float64
}
