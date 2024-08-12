use std::fmt::{Display, Formatter};
#[allow(clippy::deprecated, deprecated)]
use serde::{Deserialize, Serialize};

/// Represents the priority of returns for the LIDAR to output.
/// The LIDAR can have more than 1 or 2 detected "returns."
/// This indicates to the LIDAR which ones it should output.
/// For additional information, refer to [`return_order`].
///
/// [`return_order`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html?highlight=return_order#return-order
#[derive(Debug, Serialize, Deserialize, Eq, PartialEq, Copy, Clone)]
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

impl Display for ReturnOrder {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", match *self {
            Self::StrongestToWeakest => "strongest to weakest",
            Self::NearestToFarthest => "nearest to farthest",
            Self::FarthestToNearest => "farthest to nearest",
            Self::StrongestReturnFirst => "strongest return first",
            Self::LastReturnFirst => "last return first"
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
        let want = ReturnOrder::StrongestToWeakest;
        let got = ReturnOrder::default();
        assert_eq!(want, got);
    }

    #[test]
    fn test_display() {
        let cases = vec![
            (ReturnOrder::StrongestToWeakest, "strongest to weakest"),
            (ReturnOrder::NearestToFarthest, "nearest to farthest"),
            (ReturnOrder::FarthestToNearest, "farthest to nearest"),
            (ReturnOrder::StrongestReturnFirst, "strongest return first"),
            (ReturnOrder::LastReturnFirst, "last return first")
        ];

        for (return_order, want) in cases {
            info!("Displaying {return_order:?}, expecting {want:?}");
            let got = format!("{return_order}");
            assert_eq!(want, got);
        }
    }
}
