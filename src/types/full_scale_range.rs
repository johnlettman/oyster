use std::fmt::{Display, Formatter};

#[derive(Debug, Eq, PartialEq, serde::Serialize, serde::Deserialize, Copy, Clone)]
#[serde(rename_all = "UPPERCASE")]
pub enum FullScaleRange {
    Normal,
    Extended,
}

impl Default for FullScaleRange {
    fn default() -> Self {
        Self::Normal
    }
}

impl Display for FullScaleRange {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", match *self {
            Self::Normal => "normal",
            Self::Extended => "extended"
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
        let want = FullScaleRange::Normal;
        let got = FullScaleRange::default();
        assert_eq!(want, got);
    }

    #[test]
    fn test_display() {
        let cases = vec![
            (FullScaleRange::Normal, "normal"),
            (FullScaleRange::Extended, "extended")
        ];

        for (fsr, want) in cases {
            info!("Displaying {fsr:?}, expecting {want:?}");
            let got = format!("{fsr}");
            assert_eq!(want, got);
        }
    }
}
