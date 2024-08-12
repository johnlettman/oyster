use std::fmt::{Display, Formatter};
use serde::{Deserialize, Serialize};

/// Represents the power consumption and activity level of the Ouster sensor.
/// It can be either:
/// - `Normal`, which is the default mode where the sensor performs its regular operations
///   and consumes standard power, or
/// - `Standby`, a low-power mode that is useful for power, battery,
///   or thermal-conscious applications.
#[derive(Debug, Serialize, Deserialize, Eq, PartialEq, Copy, Clone)]
#[serde(rename_all = "UPPERCASE")]
pub enum OperatingMode {
    /// The default operating mode of the Ouster sensor.
    /// In this mode, the sensor performs its regular operations and consumes standard power.
    Normal,

    /// A low-power mode available from firmware version v2.0.0 and onward.
    /// It can be used in power, battery, or thermal-conscious applications.
    /// Without undergoing a standard operation, the sensor consumes less power in this mode.
    Standby,
}

impl Default for OperatingMode {
    fn default() -> Self {
        Self::Normal
    }
}

impl Display for OperatingMode {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", match *self {
            OperatingMode::Normal => "normal",
            OperatingMode::Standby => "standby"
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
        let want = OperatingMode::Normal;
        let got = OperatingMode::default();
        assert_eq!(want, got);
    }

    #[test]
    fn test_display() {
        let cases = vec![
            (OperatingMode::Normal, "normal"),
            (OperatingMode::Standby, "standby")
        ];

        for (mode, want) in cases {
            info!("Displaying {mode:?}, expecting {want:?}");
            let got = format!("{mode}");
            assert_eq!(want, got);
        }
    }
}
