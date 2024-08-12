use num_derive::{FromPrimitive, ToPrimitive};
use std::fmt::{Display, Formatter};
use num_traits::ToPrimitive;
use serde::{Deserialize, Serialize};

/// Represents the expected baud rate the sensor uses to decode NMEA UART input `$GPRMC` messages.
///
/// For additional information, refer to [`nmea_baud_rate`].
///
/// [`nmea_baud_rate`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#nmea-baud-rate
#[derive(
    Debug,
    Serialize,
    Deserialize,
    Eq,
    PartialEq,
    Copy,
    Clone,
    FromPrimitive,
    ToPrimitive,
)]
#[repr(u32)]
pub enum BaudRate {
    /// 9600 baud.
    #[serde(rename = "BAUD_9600")]
    X9600 = 9600,

    /// 115200 baud.
    #[serde(rename = "BAUD_115200")]
    X115200 = 115200,
}

impl BaudRate {
    pub fn baud(&self) -> u32 {
        self.to_u32().expect("It should convert to u32")
    }
}

impl Display for BaudRate {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "{} baud",
            match self {
                BaudRate::X9600 => "9600 baud",
                BaudRate::X115200 => "115200 baud",
            }
        )
    }
}

impl Default for BaudRate {
    fn default() -> Self {
        Self::X9600
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    use test_log::test;
    use log::info;

    #[test]
    fn test_baud() {
        let cases = vec![
            (BaudRate::X9600, 9600),
            (BaudRate::X115200, 115200)
        ];

        for (rate, want) in cases {
            info!("Getting baud for {rate:?}, expecting {want:?}");
            let got = rate.baud();
            assert_eq!(want, got);
        }
    }

    #[test]
    fn test_display() {
        let cases = vec![
            (BaudRate::X9600, "9600 baud"),
            (BaudRate::X115200, "115200 baud")
        ];

        for (rate, want) in cases {
            info!("Getting baud for {rate:?}, expecting {want:?}");
            let got = format!("{rate}");
            assert_eq!(want, got);
        }
    }

    #[test]
    fn test_default() {
        let want = BaudRate::X9600;
        let got = BaudRate::default();
        assert_eq!(want, got);
    }
}
