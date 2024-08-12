use std::fmt::{Display, Formatter, Result};

/// Represents the horizontal resolution and rotation rate of the sensor.
/// The effective range of the sensor is increased by 15-20% for every halving of the number of points gathered.
/// For example, [LidarMode::Scan512x10] has a 15-20% longer range than [LidarMode::Scan512x20].
///
/// For additional information, refer to [`lidar_mode`].
///
/// [`lidar_mode`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html?highlight=512x10#lidar-mode
#[derive(Debug, serde::Serialize, serde::Deserialize, Clone, Copy, PartialEq, Eq, Hash)]
pub enum LidarMode {
    /// 10 scans of 512 columns per second.
    #[serde(rename = "512x10")]
    Scan512x10,

    /// 20 scans of 512 columns per second.
    #[serde(rename = "512x20")]
    Scan512x20,

    /// 10 scans of 1024 columns per second.
    #[serde(rename = "1024x10")]
    Scan1024x10,

    /// 20 scans of 1024 columns per second.
    #[serde(rename = "1024x20")]
    Scan1024x20,

    /// 10 scans of 2048 columns per second.
    #[serde(rename = "2048x10")]
    Scan2048x10,

    /// 5 scans of 4096 columns per second.
    #[serde(rename = "4096x5")]
    Scan4096x5,
}

impl LidarMode {
    /// Returns the number of columns based on the enum variant.
    pub const fn columns(self) -> usize {
        match self {
            Self::Scan512x10 | Self::Scan512x20 => 512,
            Self::Scan1024x10 | Self::Scan1024x20 => 1024,
            Self::Scan2048x10 => 2048,
            Self::Scan4096x5 => 4096,
        }
    }

    /// Returns the frequency corresponding to the specified mode.
    pub const fn frequency(self) -> usize {
        match self {
            Self::Scan512x20 | Self::Scan1024x20 => 20,
            Self::Scan512x10 | Self::Scan1024x10 | Self::Scan2048x10 => 10,
            Self::Scan4096x5 => 5,
        }
    }
}

impl Default for LidarMode {
    fn default() -> Self {
        Self::Scan1024x10
    }
}

impl Display for LidarMode {
    fn fmt(&self, f: &mut Formatter) -> Result {
        write!(
            f,
            "{}",
            match self {
                Self::Scan512x10 => "512x10",
                Self::Scan512x20 => "512x20",
                Self::Scan1024x10 => "1024x10",
                Self::Scan1024x20 => "1024x20",
                Self::Scan2048x10 => "2048x10",
                Self::Scan4096x5 => "4096x5",
            }
        )
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    use test_log::test;
    use log::info;

    #[test]
    fn test_columns() {
        let cases = vec![
            (LidarMode::Scan512x10, 512),
            (LidarMode::Scan512x20, 512),
            (LidarMode::Scan1024x10, 1024),
            (LidarMode::Scan1024x20, 1024),
            (LidarMode::Scan2048x10, 2048),
            (LidarMode::Scan4096x5, 4096)
        ];

        for (mode, want) in cases {
            info!("Getting columns for {mode:?}, expecting {want:?}");
            let got = mode.columns();
            assert_eq!(want, got);
        }
    }

    #[test]
    fn test_frequency() {
        let cases = vec![
            (LidarMode::Scan512x10, 10),
            (LidarMode::Scan512x20, 20),
            (LidarMode::Scan1024x10, 10),
            (LidarMode::Scan1024x20, 20),
            (LidarMode::Scan2048x10, 10),
            (LidarMode::Scan4096x5, 5)
        ];

        for (mode, want) in cases {
            info!("Getting frequency for {mode:?}, expecting {want:?}");
            let got = mode.frequency();
            assert_eq!(want, got);
        }
    }

}


