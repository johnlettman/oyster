#[allow(clippy::deprecated, deprecated)]

/// Represents the priority of returns for the LIDAR to output.
/// The LIDAR can have more than 1 or 2 detected "returns."
/// This indicates to the LIDAR which ones it should output.
/// For additional information, refer to [`return_order`].
///
/// [`return_order`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html?highlight=return_order#return-order
#[derive(Debug, serde::Serialize, serde::Deserialize, Eq, PartialEq, Copy, Clone)]
#[serde(rename_all = "SCREAMING_SNAKE_CASE")]
pub enum ReturnOrder {
    /// The priority of LIDAR returns with the strongest returns first.
    StrongestToWeakest,

    /// The priority of LIDAR returns with the farthest returns first.
    NearestToFarthest,

    /// The priority of LIDAR returns with the nearest returns first.
    FarthestToNearest,

    /// The priority of LIDAR returns with the strongest returns first.
    #[deprecated = "Only present in old test firmware."]
    StrongestReturnFirst,

    /// The priority of LIDAR returns with the last returns first.
    #[deprecated = "Only present in old test firmware."]
    LastReturnFirst,
}

impl Default for ReturnOrder {
    fn default() -> Self {
        Self::StrongestToWeakest
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_default() {
        let want = ReturnOrder::StrongestToWeakest;
        let got = ReturnOrder::default();
        assert_eq!(want, got);
    }
}
