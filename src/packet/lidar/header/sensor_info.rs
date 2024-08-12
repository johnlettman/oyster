use binrw::prelude::*;
use modular_bitfield::prelude::*;

/// Represents the header metadata, which contains Initialization ID and Serial Number.
#[bitfield(bits = 64)]
#[derive(Debug, BinRead, BinWrite, Clone, Copy, Eq, PartialEq)]
#[br(map = Self::from_bytes)]
#[bw(map = |&x| Self::into_bytes(x))]
#[cfg_attr(feature = "pyo3", pyclass)]
pub struct SensorInfo {
    /// Updates on every reinitialization, which may be triggered by the
    /// user or an error, and every reboot.
    /// `HeaderSensorInfo.serial_number`
    pub initialization_id: B24,

    /// Serial number of the sensor. This value is unique to each sensor and can be found
    /// on a sticker affixed to the top of the sensor.
    pub serial_number: B40,
}

impl Default for SensorInfo {
    fn default() -> Self {
        Self { bytes: [0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00] }
    }
}
