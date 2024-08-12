use std::fmt::{Display, Formatter};
use serde::{Deserialize, Serialize};

/// Represents the mode of the `MULTIPURPOSE_IO` pin.
#[derive(Debug, Serialize, Deserialize, Eq, PartialEq, Copy, Clone)]
pub enum MultipurposeIoMode {
    /// Do not output a `SYNC_PULSE_OUT` signal.
    #[serde(rename = "OFF")]
    Off,

    /// Reconfigure the `MULTIPURPOSE_IO` port as an input.
    /// See [Setting Ouster Sensors Time Source] for more information.
    ///
    /// [Setting Ouster Sensors Time Source]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/time_sync/time-sync.html#setting-sensor-time
    #[serde(rename = "INPUT_NMEA_UART")]
    InputFromNmeaUart,

    /// Output a `SYNC_PULSE_OUT` signal synchronized with the internal clock.
    #[serde(rename = "OUTPUT_FROM_INTERNAL_OSC")]
    OutputFromInternalOscillator,

    /// Output a `SYNC_PULSE_OUT` signal synchronized with a `SYNC_PULSE_IN` provided to the unit.
    #[serde(rename = "OUTPUT_FROM_SYNC_PULSE_IN")]
    OutputFromSyncPulseIn,

    /// Output a `SYNC_PULSE_OUT` signal synchronized with an external PTP IEEE 1588 master.
    #[serde(rename = "OUTPUT_FROM_PTP_1588")]
    OutputFromPtp1588,

    /// Output a `SYNC_PULSE_OUT` signal with a user defined rate in an integer number of degrees.
    #[serde(rename = "OUTPUT_FROM_ENCODER_ANGLE")]
    OutputFromEncoderAngle,
}

impl Default for MultipurposeIoMode {
    fn default() -> Self {
        Self::Off
    }
}

impl Display for MultipurposeIoMode {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", match *self {
            MultipurposeIoMode::Off => "off",
            MultipurposeIoMode::InputFromNmeaUart => "input from NMEA UART",
            MultipurposeIoMode::OutputFromInternalOscillator => "output from internal oscillator",
            MultipurposeIoMode::OutputFromSyncPulseIn => "output from SYNC_PULSE_IN",
            MultipurposeIoMode::OutputFromPtp1588 => "output from IEEE PTP 1588",
            MultipurposeIoMode::OutputFromEncoderAngle => "output from encoder angle"
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
        let want = MultipurposeIoMode::Off;
        let got = MultipurposeIoMode::default();
        assert_eq!(want, got);
    }

    #[test]
    fn test_display() {
        let cases = vec![
            (MultipurposeIoMode::Off, "off"),
            (MultipurposeIoMode::InputFromNmeaUart, "input from NMEA UART"),
            (MultipurposeIoMode::OutputFromInternalOscillator, "output from internal oscillator"),
            (MultipurposeIoMode::OutputFromSyncPulseIn, "output from SYNC_PULSE_IN"),
            (MultipurposeIoMode::OutputFromPtp1588, "output from IEEE PTP 1588"),
            (MultipurposeIoMode::OutputFromEncoderAngle, "output from encoder angle")
        ];

        for (mode, want) in cases {
            info!("Displaying {mode:?}, expecting {want:?}");
            let got = format!("{mode}");
            assert_eq!(want, got);
        }
    }


}
