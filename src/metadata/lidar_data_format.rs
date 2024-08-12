use crate::metadata::CalibrationStatus;
use crate::packet::LidarProfile;
use crate::types::{ColumnWindow, LidarMode};
use serde::{Deserialize, Serialize};
use std::str::FromStr;
use void::Void;

const DEFAULT_COLUMNS_PER_PACKET: usize = 16;

#[derive(Debug, Serialize, Deserialize, Eq, PartialEq)]
pub struct LidarDataFormat {
    #[serde(skip)]
    message: String,

    /// The configuration of LIDAR packets.
    ///
    /// For additional information, [`udp_profile_lidar`].
    ///
    /// [`udp_profile_lidar`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#udp-profile-lidar
    #[serde(rename = "udp_profile_lidar")]
    profile: LidarProfile,

    column_window: ColumnWindow,
    columns_per_frame: usize,
    columns_per_packet: usize,

    pixel_shift_by_row: Vec<isize>,
    pixels_per_column: usize,
}

impl LidarDataFormat {
    pub fn default_for(mode: LidarMode) -> LidarDataFormat {
        let pixels_per_column: usize = 64;
        let columns_per_packet: usize = DEFAULT_COLUMNS_PER_PACKET;
        let columns_per_frame: usize = mode.columns();
        let column_window = ColumnWindow::default_for(columns_per_frame);

        let pixel_shift_by_row: Vec<isize> = match columns_per_frame {
            512 => Vec::from([9, 6, 3, 0]).repeat(16),
            1024 => Vec::from([18, 12, 6, 0]).repeat(16),
            2048 => Vec::from([36, 24, 12, 0]).repeat(16),
            _ => Vec::from([]),
        };

        LidarDataFormat {
            message: "".to_string(),
            profile: LidarProfile::DEFAULT,
            pixels_per_column,
            column_window,
            columns_per_frame,
            columns_per_packet,
            pixel_shift_by_row,
        }
    }

    fn max_frame_id(&self) -> usize {
        match self.profile {
            LidarProfile::Legacy => u32::MAX as usize,
            _ => u16::MAX as usize,
        }
    }

    fn header_size(&self) -> usize {
        match self.profile {
            LidarProfile::Legacy => 0,
            _ => 32,
        }
    }

    fn footer_size(&self) -> usize {
        match self.profile {
            LidarProfile::Legacy => 0,
            _ => 32,
        }
    }

    fn column_size(&self) -> usize {
        return 10;
    }

    fn column_header_size(&self) -> usize {
        match self.profile {
            LidarProfile::Legacy => 16,
            _ => 12,
        }
    }

    fn column_footer_size(&self) -> usize {
        match self.profile {
            LidarProfile::Legacy => 4,
            _ => 0,
        }
    }

    fn column_status_offset(&self) -> usize {
        match self.profile {
            LidarProfile::Legacy => self.column_size() - self.column_footer_size(),
            _ => 10,
        }
    }
}

impl FromStr for LidarDataFormat {
    type Err = Void;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut data_format = Self::default_for(LidarMode::default());
        data_format.message = s.to_string();
        Ok(data_format)
    }
}
