use binrw::prelude::*;
use modular_bitfield::prelude::*;

use crate::packet::lidar::header::{AlertFlags, PacketType, SensorInfo, Status};

/// A LIDAR packet header.
///
/// ```rust
/// use std::io::Cursor;
/// use oyster::packet::lidar::Header;
/// use binrw::BinRead;
/// use oyster::packet::lidar::header::PacketType;
///
/// let header_buffer: Vec<u8> = vec![
///     0x00, 0x01  // packet_type
///     0x00, 0xFF  // frame_id
///
///     // sensor_info:
///     0x00, 0xFF, 0x00              // initialization_id
///     0x01, 0x23, 0x45, 0x67, 0x89  // serial_number
///
///     // status:
///     // [shot_limiting]_[thermal_shutdown]
///     0x0_0,
///
///     0x00,  // shot_limiting_countdown
///     0x00,  // thermal_shutdown_countdown
///
///     // alert_flags:
///     // [alerts_active]_[cursor_overflow]_[cursor]
///     0b0_0_101010
/// ];
///
/// let mut cursor = Cursor::new(header_buffer);
/// let header: Header = Header::read(&mut cursor).unwrap();
///
/// assert_eq!(header.packet_type, PacketType::LidarData);
/// assert_eq!(header.frame_id, 0x00FF);
/// assert_eq!(header.sensor_info.initialization_id(), 0x00FF00);
/// assert_eq!(header.sensor_info.serial_number(), 0x0123456789);
/// assert!(!header.status.shot_limiting());
/// assert!(!header.status.thermal_shutdown());
/// assert!(!header.alert_flags.alerts_active());
/// assert!(!header.alert_flags.cursor_overflow());
/// assert_eq!(header.alert_flags.cursor(), 0b101010);
/// ```
#[derive(Debug, BinRead, BinWrite, Clone, Eq, PartialEq)]
#[brw(little)]
pub struct Header {
    /// Identifies LIDAR data vs. other packets in stream.
    /// The value is `0x1` for LIDAR packets.
    pub packet_type: PacketType,

    /// Index of the lidar scan, increments every time the sensor completes a rotation,
    /// crossing the zero azimuth angle.
    pub frame_id: u16,

    /// Contains the Initialization ID and Serial Number of the LIDAR.
    pub sensor_info: SensorInfo,

    /// Contains the Shot Limiting and Thermal Shutdown status.
    pub status: Status,

    /// Countdown from 30s to indicate when shot limiting is imminent.
    pub shot_limiting_countdown: u8,

    /// Countdown from 30s to indicate that thermal shutdown is imminent.
    pub thermal_shutdown_countdown: u8,

    /// Contains flags pertaining to the various alerts the sensor can issue.
    pub alert_flags: AlertFlags,
}

impl Default for Header {
    fn default() -> Self {
        Self {
            packet_type: PacketType::default(),
            frame_id: 0,
            sensor_info: SensorInfo::default(),
            status: Status::default(),
            shot_limiting_countdown: 0,
            thermal_shutdown_countdown: 0,
            alert_flags: AlertFlags::default(),
        }
    }
}
