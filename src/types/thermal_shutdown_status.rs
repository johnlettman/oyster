use std::fmt::{Display, Formatter};
use num_derive::{FromPrimitive, ToPrimitive};

/// ThermalShutdownStatus represents the state of the Ouster sensor when it reaches
/// the maximum operating temperature. This is handled by an independent state machine
/// that triggers a `IMMINENT` state (with an `OVERTEMP` alert) at the sensor's
/// maximum temperature. If the sensor stays at this temperature for more than 30 seconds,
/// it enters a `SHUTDOWN` state, issuing alert `0x0100006B` and ceasing operation.
/// Otherwise, during normal operation the sensor will send `NORMAL`.
///
/// For additional information, refer to [Shot Limiting].
///
/// [Shot Limiting]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/sensor_operations/sensor-operations.html#shot-limiting
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
#[serde(rename_all = "UPPERCASE")]
pub enum ThermalShutdownStatus {
    /// Normal operation of the sensor.
    Normal = 0,

    /// Imminent thermal shutdown due to maximum temperature.
    Imminent = 1,
}

impl Default for ThermalShutdownStatus {
    fn default() -> Self {
        Self::Normal
    }
}

impl Display for ThermalShutdownStatus {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", match *self {
            ThermalShutdownStatus::Normal => "normal",
            ThermalShutdownStatus::Imminent => "imminent"
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
        let want = ThermalShutdownStatus::Normal;
        let got = ThermalShutdownStatus::default();
        assert_eq!(want, got);
    }

    #[test]
    fn test_display() {
        let cases = vec![
            (ThermalShutdownStatus::Normal, "normal"),
            (ThermalShutdownStatus::Imminent, "imminent")
        ];

        for (status, want) in cases {
            info!("Displaying {status:?}, expecting {want:?}");
            let got = format!("{status}");
            assert_eq!(want, got);
        }
    }
}
