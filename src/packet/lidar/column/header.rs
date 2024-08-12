use chrono::{DateTime, Utc};
#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

use crate::packet::lidar::Header;
use binrw::{BinRead, BinWrite};
use modular_bitfield::prelude::*;

#[bitfield(bits = 96)]
#[derive(Debug, BinRead, BinWrite, Eq, PartialEq, Clone)]
#[br(map = Self::from_bytes)]
#[cfg_attr(feature = "pyo3", pyclass)]
pub struct HeaderBlock {
    #[skip]
    __: B15,
    pub status: bool,
    pub measurement_id: B16,
    pub timestamp: B64,
}

impl HeaderBlock {
    pub fn datetime(&self) -> DateTime<Utc> {
        DateTime::from_timestamp_nanos(self.timestamp() as i64)
    }

    pub fn set_datetime(&mut self, datetime: DateTime<Utc>) {
        self.set_timestamp(datetime.timestamp_nanos_opt().unwrap() as u64)
    }

    pub fn with_datetime(mut self, datetime: DateTime<Utc>) -> Self {
        self.set_timestamp(datetime);
        self
    }
}

impl Default for Header {
    fn default() -> Self {
        Self::new().with_datetime(Utc::now())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_datetime() {
        let timestamp: u64 = 1722737231000000000;
        let datetime: DateTime<Utc> =
            DateTime::parse_from_rfc3339("2024-08-04T02:07:11Z").expect("should parse").to_utc();

        let mut header = HeaderBlock::new().with_timestamp(timestamp);
        assert_eq!(header.datetime(), datetime);

        header.set_datetime(datetime);
        assert_eq!(header.timestamp(), timestamp);
    }
}
