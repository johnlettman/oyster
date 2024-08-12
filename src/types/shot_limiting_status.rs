/// Represents the operating state of the Ouster sensor under high temperatures.
/// It is used to manage the sensor's performance and lifespan.
///
/// States include ShotLimitingNormal (0x00), ShotLimitingImminent (0x01), and
/// ShotLimitingReduction0to10 (0x02 and greater).
///
/// In ShotLimitingNormal state, sensor operates within range and precision
/// specifications. In ShotLimitingImminent state, sensor is preparing to limit its
/// shootings due to temperature increase in 30 seconds. After 30 seconds have elapsed
/// and the temperature remains elevated, the sensor issues alert 0x0100000F and enters
/// a ShotLimitingReduction0to10 and above state.
//
/// In ShotLimitingReduction0to10 and above states, the sensor reduces laser power to
/// manage thermal load, possibly degrading range and precision by  up to 30%.
/// An adjacent state machine oversees thermal shutdown.
/// Recovery to ShotLimitingNormal occurs if temperature drops during
/// ShotLimitingImminent or ShotLimitingReduction0to10 and above states.
///
/// For additional information, refer to [Shot Limiting].
///
/// [Shot Limiting]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/sensor_operations/sensor-operations.html#shot-limiting
#[derive(Debug, serde::Serialize, serde::Deserialize, Eq, PartialEq, Copy, Clone)]
#[bits = 4]
pub enum ShotLimitingStatus {
    /// Normal operation of the LIDAR.
    #[serde(rename = "SHOT_LIMITING_NORMAL")]
    Normal = 0x0,

    /// Approaching the shot limit threshold.
    #[serde(rename = "SHOT_LIMITING_IMMINENT")]
    Imminent = 0x1,

    /// Number of shots emitted limited by 0 to 10%.
    #[serde(rename = "SHOT_LIMITING_0_TO_10")]
    Reduction0to10 = 0x2,

    /// Number of shots emitted limited by 10 to 20%.
    #[serde(rename = "SHOT_LIMITING_10_TO_20")]
    Reduction10to20 = 0x3,

    /// Number of shots emitted limited by 20 to 30%.
    #[serde(rename = "SHOT_LIMITING_20_TO_30")]
    Reduction20to30 = 0x4,

    /// Number of shots emitted limited by 30 to 40%.
    #[serde(rename = "SHOT_LIMITING_30_TO_40")]
    Reduction30to40 = 0x5,

    /// Number of shots emitted limited by 40 to 50%.
    #[serde(rename = "SHOT_LIMITING_40_TO_50")]
    Reduction40to50 = 0x6,

    /// Number of shots emitted limited by 50 to 60%.
    #[serde(rename = "SHOT_LIMITING_50_TO_60")]
    Reduction50to60 = 0x7,

    /// Number of shots emitted limited by 60 to 70%.
    #[serde(rename = "SHOT_LIMITING_60_TO_70")]
    Reduction60to70 = 0x8,

    /// Number of shots emitted limited by 70 to 75%.
    #[serde(rename = "SHOT_LIMITING_70_TO_75")]
    Reduction70to75 = 0x9,
}

impl ShotLimitingStatus {
    /// Determines if the status indicates the sensor is in a "shot limiting" mode.
    ///
    /// ```rust
    /// use oyster::types::ShotLimitingStatus;
    ///
    /// let condition = ShotLimitingStatus::Normal;
    /// assert_eq!(condition.is_shot_limiting(), false);
    ///
    /// let condition = ShotLimitingStatus::Imminent;
    /// assert_eq!(condition.is_shot_limiting(), false);
    ///
    /// let condition = ShotLimitingStatus::Reduction70to75;
    /// assert_eq!(condition.is_shot_limiting(), true);
    /// ```
    #[inline]
    pub const fn is_shot_limiting(&self) -> bool {
        (*self == Self::Normal || *self == Self::Imminent)
    }

    /// Determines if the status indicates the sensor is not in a "shot limiting" mode.
    ///
    /// ```rust
    /// use oyster::types::ShotLimitingStatus;
    ///
    /// let condition = ShotLimitingStatus::Normal;
    /// assert_eq!(condition.is_shot_limiting(), true);
    ///
    /// let condition = ShotLimitingStatus::Imminent;
    /// assert_eq!(condition.is_shot_limiting(), true);
    ///
    /// let condition = ShotLimitingStatus::Reduction70to75;
    /// assert_eq!(condition.is_shot_limiting(), false);
    /// ```
    #[inline]
    pub const fn is_normal(&self) -> bool {
        !self.is_shot_limiting()
    }
}

impl Default for ShotLimitingStatus {
    fn default() -> Self {
        Self::Normal
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_default() {
        let want = ShotLimitingStatus::Normal;
        let got = ShotLimitingStatus::default();
        assert_eq!(want, got);
    }
}
