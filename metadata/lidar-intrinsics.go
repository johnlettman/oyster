package metadata

// LIDARIntrinsics represents the intrinsic parameters of the LiDAR.
//
// For additional information, refer to [Ouster docs: lidar_intrinsics].
//
// [Ouster docs: lidar_intrinsics]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/http-api-v1.html#get-api-v1-sensor-metadata-lidar-intrinsics
type LIDARIntrinsics struct {
	// LIDARToSensorTransform is the transformation matrix from the LiDAR coordinate frame to the sensor coordinate frame.
	// It is a 4x4 matrix stored in row-major order in a 16-element float64 array.
	LIDARToSensorTransform [16]float64 `json:"lidar_to_sensor_transform"`
}
