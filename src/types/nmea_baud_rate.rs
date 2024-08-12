use num_derive::{FromPrimitive, ToPrimitive};
use std::fmt::{Display, Formatter};

/// Represents the expected baud rate the sensor uses to decode NMEA UART input `$GPRMC` messages.
///
/// For additional information, refer to [`nmea_baud_rate`].
///
/// [`nmea_baud_rate`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#nmea-baud-rate
#[derive(
    Debug,
    serde::Serialize,
    serde::Deserialize,
    Eq,
    PartialEq,
    Copy,
    Clone,
    FromPrimitive,
    ToPrimitive,
)]
#[repr(u32)]
pub enum NmeaBaudRate {
    /// 9600 baud.
    #[serde(rename = "BAUD_9600")]
    Baud9600 = 9600,

    /// 115200 baud.
    #[serde(rename = "BAUD_115200")]
    Baud115200 = 115200,
}

impl Display for NmeaBaudRate {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "{} baud",
            match self {
                NmeaBaudRate::Baud9600 => "9600 baud",
                NmeaBaudRate::Baud115200 => "115200 baud",
            }
        )
    }
}

impl Default for NmeaBaudRate {
    fn default() -> Self {
        Self::Baud9600
    }
}
