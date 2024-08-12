use nalgebra::{ArrayStorage, Matrix4};
use serde::{Deserialize, Serialize};

pub const DEFAULT_IMU_TO_SENSOR_TRANSFORM: Matrix4<f64> = Matrix4::new(
    1.0, 0.0, 0.0, 6.253, 0.0, 1.0, 0.0, -11.775, 0.0, 0.0, 1.0, 7.645, 0.0, 0.0, 0.0, 1.0,
);

/// IMUIntrinsics represents the intrinsic parameters of the IMU (Inertial Measurement Unit).
/// For additional information, refer to [`imu_intrinsics`].
///
/// [`imu_intrinsics`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/http-api-v1.html#get-api-v1-sensor-metadata-imu-intrinsics
#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct ImuIntrinsics {
    /// The transformation matrix from the IMU (Inertial Measurement Unit)
    /// coordinate frame to the sensor coordinate frame.
    /// It is a 4x4 matrix stored in row-major order in a 16-element float64 array.
    #[serde(default = "ImuIntrinsics::default_imu_to_sensor_transform")]
    imu_to_sensor_transform: Matrix4<f64>,
}

impl ImuIntrinsics {
    fn default_imu_to_sensor_transform() -> Matrix4<f64> {
        DEFAULT_IMU_TO_SENSOR_TRANSFORM
    }
}
