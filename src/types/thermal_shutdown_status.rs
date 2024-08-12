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
#[derive(Debug, serde::Serialize, serde::Deserialize, Eq, PartialEq)]
#[serde(rename_all = "UPPERCASE")]
pub enum ThermalShutdownStatus {
    /// Normal operation of the sensor.
    Normal = 0x0,

    /// Imminent thermal shutdown due to maximum temperature.
    Imminent = 0x1,
}

impl Default for ThermalShutdownStatus {
    fn default() -> Self {
        Self::Normal
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_default() {
        let want = ThermalShutdownStatus::Normal;
        let got = ThermalShutdownStatus::default();
        assert_eq!(want, got);
    }
}
