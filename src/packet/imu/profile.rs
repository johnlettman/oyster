/// Represents how sensor data from the IMU is packaged and sent over the wire,
/// e.g., via UDP packets.
#[derive(Debug, serde::Serialize, serde::Deserialize, Eq, PartialEq, Copy, Clone)]
#[serde(rename_all = "UPPERCASE")]
pub enum Profile {
    Legacy,
}

impl Default for Profile {
    fn default() -> Self {
        Self::default()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_default() {
        let want = Profile::Legacy;
        let got = Profile::default();
        assert_eq!(want, got);
    }
}
