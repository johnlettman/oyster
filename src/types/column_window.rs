use std::fmt::{Display, Formatter, Result};

/// The window over which the sensor fires in columns.
#[derive(Debug, serde::Serialize, serde::Deserialize, Eq, PartialEq, Copy, Clone)]
pub struct ColumnWindow(usize, usize);

impl ColumnWindow {
    pub const fn default_for(columns_per_frame: usize) -> ColumnWindow {
        ColumnWindow(0, columns_per_frame - 1)
    }
}

impl Display for ColumnWindow {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        write!(f, "({}, {})", self.0, self.1)
    }
}

impl Default for ColumnWindow {
    #[inline]
    fn default() -> Self {
        Self::default_for(1024)
    }
}
