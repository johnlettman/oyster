use std::cmp::{max, min, Ordering};
use std::fmt::{Display, Formatter, Result};
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, Eq, Ord, Copy, Clone)]
pub struct AzimuthWindow(u32, u32);

impl AzimuthWindow {
    pub const MIN_MILLIDEGREES: u32 = 0;
    pub const MAX_MILLIDEGREES: u32 = 360000;

    #[inline]
    pub const fn new(a: u32, b: u32) -> AzimuthWindow {
        Self(a, b)
    }

    /// Returns the size of the window by finding the absolute difference between
    /// the two window boundaries.
    ///
    /// # Examples
    ///
    /// ```rust
    /// let w = oyster::types::AzimuthWindow(3, 7);
    /// let size = w.size();
    /// assert_eq!(size, 7 - 3);
    /// ```
    #[inline]
    pub fn size(&self) -> u32 {
        self.0.abs_diff(self.1)
    }

    #[inline]
    pub fn zero_size(&self) -> bool {
        self.size() == 0
    }

    /// Returns the end of the window by finding the maximum value between
    /// the two window boundaries.
    ///
    /// # Example
    ///
    /// ```rust
    /// let w = oyster::types::AzimuthWindow(3, 7);
    /// let start = w.start();
    /// assert_eq!(start, 3);
    /// ```
    #[inline]
    pub fn start(&self) -> u32 {
        min(self.0, self.1)
    }

    /// Returns the end of the window by finding the maximum value between
    /// the two window boundaries.
    ///
    /// # Example
    ///
    /// ```rust
    /// # #[macro_use] extern crate oyster; fn main() {
    /// let w = oyster::types::AzimuthWindow(3, 7);
    /// let end = w.end();
    /// assert_eq!(end, 7);
    /// # }
    /// ```
    #[inline]
    pub fn end(&self) -> u32 {
        max(self.0, self.1)
    }

    #[inline]
    pub const fn zero(&self) -> bool {
        self.0 == 0 && self.1 == 0
    }

    #[inline]
    pub const fn valid(&self) -> bool {
        Self::valid_millidegrees(self.0) && Self::valid_millidegrees(self.1)
    }

    #[inline]
    pub const fn valid_millidegrees(md: u32) -> bool {
        Self::MIN_MILLIDEGREES <= md && md <= Self::MAX_MILLIDEGREES
    }
}

impl Default for AzimuthWindow {
    fn default() -> Self {
        Self(Self::MIN_MILLIDEGREES, Self::MAX_MILLIDEGREES)
    }
}

impl Display for AzimuthWindow {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        return write!(f, "({}m° -> {}m°)", self.start(), self.end());
    }
}

impl PartialEq for AzimuthWindow {
    /// Checks if this [AzimuthWindow] is equal to another [AzimuthWindow].
    ///
    /// # Example
    /// ```
    /// use oyster::types::AzimuthWindow;
    /// let w1 = AzimuthWindow(1, 2);
    /// let w2 = AzimuthWindow(2, 1);
    ///
    /// assert!(w1.eq(&w2));
    /// ```
    fn eq(&self, other: &Self) -> bool {
        (self.0 == other.0 && self.1 == other.1) || (self.0 == other.1 && self.1 == other.0)
    }
}

impl PartialOrd for AzimuthWindow {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        self.size().partial_cmp(&other.size())
    }
}

impl From<[u32; 2]> for AzimuthWindow {
    fn from(value: [u32; 2]) -> Self {
        Self::new(value[0], value[1])
    }
}

impl From<AzimuthWindow> for [u32; 2] {
    fn from(value: AzimuthWindow) -> [u32; 2] {
        [value.start(), value.end()]
    }
}

impl From<Vec<u32>> for AzimuthWindow {
    fn from(value: Vec<u32>) -> Self {
        Self::new(value[0], value[1])
    }
}

impl From<AzimuthWindow> for Vec<u32> {
    fn from(value: AzimuthWindow) -> Vec<u32> {
        vec![value.start(), value.end()]
    }
}
