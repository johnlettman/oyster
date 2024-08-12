use crate::packet::lidar;
use chrono::{DateTime, Utc};

pub struct LidarData {
    datetime: DateTime<Utc>,
    packet: Utc,
}
