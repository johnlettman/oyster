use std::fmt::{Display, Formatter};

#[derive(Debug, serde::Serialize, serde::Deserialize, Eq, PartialEq, Copy, Clone)]
pub enum TimestampMode {
    /// A free-running counter based on the sensor's internal oscillator.
    /// It counts seconds and nanoseconds since the sensor was turned on.
    /// The time is reported at a nanosecond resolution, with a minimum increment of around 10 ns.
    #[serde(rename = "TIME_FROM_INTERNAL_OSC")]
    TimeFromInternalOscillator,

    /// A counter synced to the `SYNC_PULSE_IN` input. It also counts seconds (number of pulses)
    /// and nanoseconds since sensor turn on.
    /// If the [MultipurposeIoMode](crate::types::MultipurposeIoMode) is set to
    /// [InputFromNmeaUart](crate::types::MultipurposeIoMode::InputFromNmeaUart),
    /// the seconds register jumps to the time extracted from a NMEA `$GPRMC` message read on the
    /// `MULTIPURPOSE_IO` port.
    /// The time is also reported at a nanosecond resolution with a minimum increment of
    /// around 10 ns.
    #[serde(rename = "TIME_FROM_SYNC_PULSE_IN")]
    TimeFromSyncPulseIn,

    /// Allows synchronization with an external PTP master. It's a monotonically increasing counter
    /// that starts counting seconds and nanoseconds since startup.
    /// As soon as a 1588 sync event happens, the time will be updated to seconds and nanoseconds
    /// since 1970. The counter must always count forward in time. If another 1588 sync event
    /// happens the counter will either jump forward to match the new time, or slow itself down.
    /// It is reported at nanosecond resolution, but the minimum increment varies.
    #[serde(rename = "TIME_FROM_PTP_1588")]
    TimeFromFromPtp1588,
}

impl Display for TimestampMode {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", match *self {
            Self::TimeFromInternalOscillator => "time from internal oscillator",
            Self::TimeFromSyncPulseIn => "time from SYNC_PULSE_IN",
            Self::TimeFromFromPtp1588 => "time from PTP-1588"
        })
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    use test_log::test;
    use log::info;

    #[test]
    fn test_display() {
        let cases = vec![
            (TimestampMode::TimeFromInternalOscillator, "time from internal oscillator"),
            (TimestampMode::TimeFromSyncPulseIn, "time from SYNC_PULSE_IN"),
            (TimestampMode::TimeFromFromPtp1588, "time from PTP-1588")
        ];

        for (mode, want) in cases {
            info!("Displaying {mode:?}, expecting {want:?}");
            let got = format!("{mode}");
            assert_eq!(want, got);
        }
    }
}
