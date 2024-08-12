use crate::types::FullScaleRange;
use serde::{Deserialize, Serialize};

/// Represents the format of IMU data from an Ouster sensor.
/// For additional information, refer to [`imu_data_format`].
///
/// [`imu_data_format`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/http-api-v1.html#get-api-v1-sensor-metadata-imu-data-format
#[derive(Debug, Serialize, Deserialize, Eq, PartialEq)]
pub struct ImuDataFormat {
    /// The full-scale range of the accelerometer.
    ///
    /// Settings:
    ///   - [FullScaleRange::Normal] (default):
    ///     digital output X-, Y-, Z-axis with a range fixed at ±2g.
    ///   - [FullScaleRange::Extended]:
    ///     digital-output X-, Y-, Z-axis with an expanded full-scale range of ±16g.
    accel_fsr: FullScaleRange,

    /// The full-scale range of the gyroscope.
    ///
    /// Settings:
    ///   - [FullScaleRange::Normal] (default):
    ///     digital output X-, Y-, Z-axis with a range fixed at ±250°/sec.
    ///   - [FullScaleRange::Extended]:
    ///     digital-output X-, Y-, Z-axis with a programmable full-scale range of ±2000°/sec.
    gyro_fsr: FullScaleRange,
}
