use binrw::prelude::*;
use modular_bitfield::prelude::*;

use crate::types::{ShotLimitingStatus, ThermalShutdownStatus};

/// The header status, which contains Shot Limiting and Thermal Shutdown Status.
#[derive(Debug, BinRead, BinWrite, Clone, Copy, Eq, PartialEq)]
pub struct Status {
    /// Indicates the shot limiting status of the sensor.
    pub shot_limiting: ShotLimitingStatus,

    /// Indicates whether thermal shutdown is imminent.
    pub thermal_shutdown: ThermalShutdownStatus,
}

impl Default for Status {
    /// The default value for [Status]:
    /// - [Status::shot_limiting]: [ShotLimitingStatus::Normal]
    /// - [Status::thermal_shutdown]: [ThermalShutdownStatus::Normal]
    fn default() -> Self {
        Self {
            shot_limiting: ShotLimitingStatus::default(),
            thermal_shutdown: ThermalShutdownStatus::default()
        }
    }
}
