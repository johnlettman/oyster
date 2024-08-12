use serde::{Deserialize, Serialize};

/// Represents how sensor data from the LIDAR sensor is packaged and sent over the wire,
/// e.g., via UDP packets.
#[derive(Debug, Serialize, Deserialize, Eq, PartialEq, Copy, Clone)]
pub enum Profile {
    /// The legacy LIDAR profile (deprecated).
    #[serde(rename = "LEGACY")]
    Legacy,

    /// Profile for dual returns from a LIDAR sensor.
    /// It is encoded as `RNG19_RFL8_SIG16_NIR16_DUAL`.
    ///
    /// For additional information, refer to [`RNG19_RFL8_SIG16_NIR16_DUAL` Return Profile].
    ///
    /// [`RNG19_RFL8_SIG16_NIR16_DUAL` Return Profile]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/sensor_data/sensor-data.html#dual-return-v3-x
    #[serde(rename = "RNG19_RFL8_SIG16_NIR16_DUAL")]
    DualReturns,

    /// The profile for single returns from a LIDAR sensor (default).
    /// It is encoded as `RNG19_RFL8_SIG16_NIR16`.
    ///
    /// For additional information, refer to [`RNG19_RFL8_SIG16_NIR16` Return Profile].
    ///
    /// [`RNG19_RFL8_SIG16_NIR16` Return Profile]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/sensor_data/sensor-data.html#rng19-rfl8-sig16-nir16-return-profile
    #[serde(rename = "RNG19_RFL8_SIG16_NIR16")]
    SingleReturns,

    /// The profile for single returns at a reduced data rate from the LIDAR sensor.
    /// It is encoded as `RNG15_RFL8_NIR8`.
    ///
    /// For additional information refer to [`RNG15_RFL8_NIR8` Return Profile].
    ///
    /// [`RNG15_RFL8_NIR8` Return Profile]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/sensor_data/sensor-data.html#rng15-rfl8-nir8-return-profile
    #[serde(rename = "RNG15_RFL8_NIR8")]
    SingleReturnsLowDataRate,

    /// The profile for the Functional Safety data packet format from the LIDAR sensor.
    /// It is encoded as `FUSA_RNG15_RFL8_NIR8_DUAL`.
    ///
    /// For additional information, refer to [`FUSA_RNG15_RFL8_NIR8_DUAL` Return Profile].
    ///
    /// [`FUSA_RNG15_RFL8_NIR8_DUAL` Return Profile]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/sensor_data/sensor-data.html#fusa-rng15-rfl8-nir8-dual-return-profile
    #[serde(rename = "FUSA_RNG15_RFL8_NIR8_DUAL")]
    FuSaTwoWordPixel,
}


impl Default for Profile {
    fn default() -> Self {
        Self::Legacy
    }
}
