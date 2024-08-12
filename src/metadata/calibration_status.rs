use chrono::{DateTime, TimeDelta, Utc};
use log::debug;
use serde::{Deserialize, Serialize};
use std::str::FromStr;
use void::Void;

/// Contains the calibration status of the sensor.
#[derive(Debug, Serialize, Deserialize, Eq, PartialEq, Clone)]
pub struct CalibrationStatus {
    #[serde(skip)]
    pub message: String,

    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub reflectivity: Option<ReflectivityCalibrationStatus>,
}

/// The calibration data field from the sensor.
/// Currently, this is solely used for reflectivity calibration details.
/// For additional information, refer to [Calibrated Reflectivity].
///
/// Ouster recommends contacting [Ouster Support] if you have questions on whether
/// your sensor is hardware-enabled for calibrated reflectivity.
///
/// [Ouster Support]: mailto://support@ouster.io
/// [Calibrated Reflectivity]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/sensor_data/sensor-data.html#calibrated-reflectivity
#[derive(Debug, Serialize, Deserialize, Eq, PartialEq, Clone)]
pub struct ReflectivityCalibrationStatus {
    /// The date and time when the calibration has been performed.
    pub timestamp: DateTime<Utc>,

    /// true if the sensor is factory-calibrated for better accuracy;
    /// otherwise, the sensor is using default values and likely has less accuracy.
    pub valid: bool,
}

impl ReflectivityCalibrationStatus {
    pub const fn new(timestamp: DateTime<Utc>, valid: bool) -> Self {
        Self { timestamp, valid }
    }

    /// Calculates the age of an object based on its timestamp.
    #[inline]
    pub fn age(&self) -> TimeDelta {
        return Utc::now() - self.timestamp;
    }
}

impl FromStr for CalibrationStatus {
    type Err = Void;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        Ok(CalibrationStatus { message: s.to_string(), reflectivity: None })
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_reflectivity_calibration_status_de() {
        let json_str = r#"{ "timestamp": "2022-11-29T02:00:08Z", "valid": true }"#;

        let want_timestamp: DateTime<Utc> =
            DateTime::parse_from_rfc3339("2022-11-29T02:00:08Z").expect("should parse").to_utc();
        let want_valid = true;
        let want = ReflectivityCalibrationStatus::new(want_timestamp, want_valid);

        let got: ReflectivityCalibrationStatus =
            serde_json::from_str(json_str).expect("should parse");
        assert_eq!(got, want);
    }
}
