use binrw::prelude::*;
use modular_bitfield::prelude::*;

use crate::packet::lidar::header::{AlertFlags, PacketType, SensorInfo, Status};

// #[derive(Debug, BinRead, BinWrite, Clone, Eq, PartialEq)]
// #[brw(little)]
// pub struct Header {
//     /// Identifies LIDAR data vs. other packets in stream.
//     /// The value is `0x1` for LIDAR packets.
//     pub packet_type: PacketType,
//
//     /// Index of the lidar scan, increments every time the sensor completes a rotation,
//     /// crossing the zero azimuth angle.
//     pub frame_id: u16,
//
//     /// Contains the Initialization ID and Serial Number of the LIDAR.
//     pub sensor_info: SensorInfo,
//
//     /// Contains the Shot Limiting and Thermal Shutdown status.
//     pub status: Status,
//
//     /// Countdown from 30s to indicate when shot limiting is imminent.
//     pub shot_limiting_countdown: u8,
//
//     /// Countdown from 30s to indicate that thermal shutdown is imminent.
//     pub thermal_shutdown_countdown: u8,
//
//     /// Contains flags pertaining to the various alerts the sensor can issue.
//     pub alert_flags: AlertFlags,
// }

// impl Default for Header {
//     fn default() -> Self {
//         Self {
//             packet_type: PacketType::default(),
//             frame_id: 0,
//             sensor_info: SensorInfo::default(),
//             status: Status::default(),
//             shot_limiting_countdown: 0,
//             thermal_shutdown_countdown: 0,
//             alert_flags: AlertFlags::default(),
//         }
//     }
// }
