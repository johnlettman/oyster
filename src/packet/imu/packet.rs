use serde::{Deserialize, Serialize};
use std::fmt::{Display, Formatter};

use binrw::prelude::*;

/// Represents the data packet from an IMU (Inertial Measurement Unit) sensor.
/// It contains information about timestamps, linear acceleration, and angular velocity.
///
/// For additional information, refer to [IMU Data Format].
///
/// [IMU Data Format]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/sensor_data/sensor-data.html#imu-data-format
#[derive(Debug, BinRead, BinWrite, Serialize, Deserialize, PartialEq)]
pub struct Packet {
    /// Timestamp of the monotonic system time since boot in ns.
    diagnostic_system_time: u64,

    /// Timestamp for the Accelerometer relative to
    /// [TimestampMode](crate::types::TimestampMode) in ns.
    accelerometer_time: u64,

    /// Timestamp for the Gyroscope relative to
    /// [TimestampMode](crate::types::TimestampMode) in ns.
    gyroscope_time: u64,

    /// Measured linear acceleration in g for the X axis.
    linear_acceleration_x: f32,

    /// Measured linear acceleration in g for the Y axis.
    linear_acceleration_y: f32,

    /// Measured linear acceleration in g for the Z axis.
    linear_acceleration_z: f32,

    /// Measured angular velocity in °/sec for the X axis.
    angular_velocity_x: f32,

    /// Measured angular velocity in °/sec for the Y axis.
    angular_velocity_y: f32,

    /// Measured angular velocity in °/sec for the Z axis.
    angular_velocity_z: f32,
}

impl Packet {
    /// Returns the linear acceleration as an array of three `f32` values.
    /// The linear acceleration values correspond to the acceleration along the X, Y, and Z axes, respectively.
    ///
    /// # Example
    ///
    /// ```rust
    /// let sensor = oyster::packet::imu::Packet::new();
    /// let acceleration = sensor.linear_acceleration();
    /// println!("{:?}", acceleration);
    /// ```
    pub fn linear_acceleration(&self) -> [f32; 3] {
        [self.linear_acceleration_x, self.linear_acceleration_y, self.linear_acceleration_z]
    }

    pub fn angular_velocity(&self) -> [f32; 3] {
        [self.angular_velocity_x, self.angular_velocity_y, self.angular_velocity_z]
    }
}

impl Display for Packet {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "(t+{}, →a=[{:.4}, {:.4}, {:.4}], ω=[{:.4}, {:.4}, {:.4}])",
            self.diagnostic_system_time,
            self.linear_acceleration_x,
            self.linear_acceleration_y,
            self.linear_acceleration_z,
            self.angular_velocity_x,
            self.angular_velocity_y,
            self.angular_velocity_z
        )
    }
}
