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
