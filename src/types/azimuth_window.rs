use std::cmp::{max, min, Ordering};
use std::fmt::{Display, Formatter, Result};

#[derive(Debug, serde::Serialize, serde::Deserialize, Eq, Ord, Copy, Clone)]
pub struct AzimuthWindow(usize, usize);

impl AzimuthWindow {
    pub const MIN_MILLIDEGREES: usize = 0;
    pub const MAX_MILLIDEGREES: usize = 360000;

    #[inline]
    pub const fn new(a: usize, b: usize) -> AzimuthWindow {
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
    pub fn size(&self) -> usize {
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
    pub fn start(&self) -> usize {
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
    pub fn end(&self) -> usize {
        max(self.0, self.1)
    }

    #[inline]
    pub fn zero(&self) -> bool {
        self.0 == 0 && self.1 == 0
    }

    pub fn valid(&self) -> bool {
        return self.0 >= Self::MIN_MILLIDEGREES
            && self.0 <= Self::MAX_MILLIDEGREES
            && self.1 >= Self::MIN_MILLIDEGREES
            && self.1 <= Self::MAX_MILLIDEGREES;
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

impl Display for AzimuthWindow {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        return write!(f, "({}m° -> {}m°)", self.start(), self.end());
    }
}

impl Default for AzimuthWindow {
    fn default() -> Self {
        Self(Self::MIN_MILLIDEGREES, Self::MAX_MILLIDEGREES)
    }
}

impl From<[usize; 2]> for AzimuthWindow {
    fn from(value: [usize; 2]) -> Self {
        Self::new(value[0], value[1])
    }
}

impl Into<[usize; 2]> for AzimuthWindow {
    fn into(self) -> [usize; 2] {
        [self.start(), self.end()]
    }
}

impl From<Vec<usize>> for AzimuthWindow {
    fn from(value: Vec<usize>) -> Self {
        Self::new(value[0], value[1])
    }
}

impl Into<Vec<usize>> for AzimuthWindow {
    fn into(self) -> Vec<usize> {
        vec![self.start(), self.end()]
    }
}
