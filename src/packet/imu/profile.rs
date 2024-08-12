use std::fmt::{Display, Formatter};
use serde::{Deserialize, Serialize};

/// Represents how sensor data from the IMU is packaged and sent over the wire,
/// e.g., via UDP packets.
#[derive(Debug, Serialize, Deserialize, Eq, PartialEq, Copy, Clone)]
#[serde(rename_all = "UPPERCASE")]
pub enum Profile {
    Legacy,
}

impl Default for Profile {
    fn default() -> Self {
        Self::default()
    }
}

impl Display for Profile {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", match *self {
            Self::Legacy => "legacy"
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
        let want = Profile::Legacy;
        let got = Profile::default();
        assert_eq!(want, got);
    }

    #[test]
    fn test_display() {
        let cases = vec![
            (Profile::Legacy, "legacy")
        ];

        for (profile, want) in cases {
            info!("Displaying {profile:?}, expecting {want:?}");
            let got = format!("{profile}");
            assert_eq!(want, got);
        }
    }
}
