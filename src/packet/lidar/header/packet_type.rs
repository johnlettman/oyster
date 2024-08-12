use binrw::{BinRead, BinWrite};
use num_derive::{FromPrimitive, ToPrimitive};

#[derive(Debug, BinRead, BinWrite, Eq, PartialEq, Copy, Clone, FromPrimitive, ToPrimitive)]
#[brw(repr = u8)]
pub enum PacketType {
    LidarData = 1,
}

impl Default for PacketType {
    fn default() -> Self {
        Self::LidarData
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_default() {
        let want = PacketType::LidarData;
        let got = PacketType::default();
        assert_eq!(want, got);
    }
}
