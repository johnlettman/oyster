use nalgebra::Matrix4;
use serde::{Deserialize, Serialize};

pub const DEFAULT_LIDAR_TO_SENSOR_TRANSFORM: Matrix4<f64> = Matrix4::new(
    1.0, 0.0, 0.0, 6.253, 0.0, 1.0, 0.0, -11.775, 0.0, 0.0, 1.0, 7.645, 0.0, 0.0, 0.0, 1.0,
);

/// Represents the intrinsic parameters of the LIDAR.
/// For additional information, refer to [`lidar_intrinsics`].
///
/// [`lidar_intrinsics`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/http-api-v1.html#get-api-v1-sensor-metadata-lidar-intrinsics
#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct LidarIntrinsics {
    /// The transformation matrix from the LIDAR coordinate frame to the sensor coordinate frame.
    /// It is a 4x4 matrix stored in row-major order in a 16-element float64 array.
    #[serde(default = "LidarIntrinsics::default_lidar_to_sensor_transform")]
    lidar_to_sensor_transform: Matrix4<f64>,
}

impl LidarIntrinsics {
    fn default_lidar_to_sensor_transform() -> Matrix4<f64> {
        DEFAULT_LIDAR_TO_SENSOR_TRANSFORM
    }
}
