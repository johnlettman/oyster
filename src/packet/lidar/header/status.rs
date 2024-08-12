use binrw::prelude::*;
use modular_bitfield::prelude::*;

use crate::types::{ShotLimitingStatus, ThermalShutdownStatus};

/// The header status, which contains Shot Limiting and Thermal Shutdown Status.
#[bitfield(bits = 8)]
#[derive(Debug, BinRead, BinWrite, Clone, Copy, Eq, PartialEq)]
#[br(map = Self::from_bytes)]
#[bw(map = |&x| Self::into_bytes(x))]
#[cfg_attr(feature = "pyo3", pyclass)]
pub struct Status {
    /// Indicates the shot limiting status of the sensor.
    #[bits = 4]
    pub shot_limiting: ShotLimitingStatus,

    /// Indicates whether thermal shutdown is imminent.
    #[bits = 4]
    pub thermal_shutdown: ThermalShutdownStatus,
}

impl Default for Status {
    /// The default value for [Status]:
    /// - [Status::shot_limiting]: [ShotLimitingStatus::Normal]
    /// - [Status::thermal_shutdown]: [ThermalShutdownStatus::Normal]
    ///
    /// ```rust
    /// use oyster::packet::lidar::header::Status;
    /// use oyster::types::{ShotLimitingStatus, ThermalShutdownStatus};
    ///
    /// let status = Status::default();
    /// assert_eq!(status.shot_limiting(), ShotLimitingStatus::Normal);
    /// assert_eq!(status.thermal_shutdown(), ThermalShutdownStatus::Normal);
    /// ```
    fn default() -> Self {
        Self { bytes: [0x00] }
    }
}
