use std::fmt::{Display, Formatter};
use serde::{Deserialize, Serialize};

/// Represents the polarity of a signal or an electrical voltage.
/// It is an enumerated type that provides constants for two different polarities:
///
///   - [Polarity::ActiveHigh]:
///     the signal is defined as a signal that is
///     true when it is at a high voltage (aka high true)
///   - [Polarity::ActiveLow]:
///     the signal is defined as a signal that is true
///     when it is at a low voltage (aka low true)
///
/// For additional information about polarities, refer to [Signal Polarity].
/// For additional information regarding polarities as applied to Ouster sensors,
/// refer to the following:
///
///   - [`nmea_in_polarity`]
///   - [`sync_pulse_in_polarity`]
///   - [`sync_pulse_out_polarity`]
///
/// [Signal Polarity]: https://engineering.purdue.edu/~meyer/DDU270/Refs/Pld/pal_polarity.pdf
/// [`nmea_in_polarity`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#nmea-in-polarity
/// [`sync_pulse_in_polarity`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sync-pulse-in-polarity
/// [`sync_pulse_out_polarity`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sync-pulse-out-polarity
#[derive(Debug, Serialize, Deserialize, Eq, PartialEq, Clone, Copy)]
#[serde(rename_all = "SCREAMING_SNAKE_CASE")]
pub enum Polarity {
    ActiveHigh,
    ActiveLow,
}

impl Default for Polarity {
    fn default() -> Self {
        Self::ActiveHigh
    }
}

impl Display for Polarity {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", match *self {
            Self::ActiveHigh => "active high",
            Self::ActiveLow => "active low"
        })
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    use test_log::test;
    use log::info;

    #[test]
    fn test_default() {
        let want = Polarity::ActiveHigh;
        let got = Polarity::default();
        assert_eq!(want, got);
    }

    #[test]
    fn test_display() {
        let cases = vec![
            (Polarity::ActiveHigh, "active high"),
            (Polarity::ActiveLow, "active low")
        ];

        for (polarity, want) in cases {
            info!("Displaying {polarity:?}, want {want:?}");
            let got = format!("{polarity}");
            assert_eq!(want, got);
        }
    }
}
