/// Represents the power consumption and activity level of the Ouster sensor.
/// It can be either:
/// - `Normal`, which is the default mode where the sensor performs its regular operations
///   and consumes standard power, or
/// - `Standby`, a low-power mode that is useful for power, battery,
///   or thermal-conscious applications.
#[derive(Debug, serde::Serialize, serde::Deserialize, Eq, PartialEq, Copy, Clone)]
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_default() {
        let want = OperatingMode::Normal;
        let got = OperatingMode::default();
        assert_eq!(want, got);
    }
}
