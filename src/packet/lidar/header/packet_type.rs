use std::fmt::{Display, Formatter};
use binrw::{BinRead, BinWrite};
use num_derive::{FromPrimitive, ToPrimitive};

#[derive(Debug, BinRead, BinWrite, Eq, PartialEq, Copy, Clone, FromPrimitive, ToPrimitive)]
#[repr(u8)]
#[brw(repr = u8)]
pub enum PacketType {
    LidarData = 1,
}

impl Default for PacketType {
    fn default() -> Self {
        Self::LidarData
    }
}

impl Display for PacketType {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", match *self {
            Self::LidarData => "LIDAR data"
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
        let want = PacketType::LidarData;
        let got = PacketType::default();
        assert_eq!(want, got);
    }

    #[test]
    fn test_display() {
        let cases = vec![
            (PacketType::LidarData, "LIDAR data")
        ];

        for (packet_type, want) in cases {
            info!("Displaying {packet_type:?}, expecting {want:?}");
            let got = format!("{packet_type}");
            assert_eq!(want, got);
        }
    }
}
