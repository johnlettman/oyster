use std::fmt::{Display, Formatter};
use binrw::{BinRead, BinWrite};
use num_derive::{FromPrimitive, ToPrimitive};

#[derive(Debug, BinRead, BinWrite, Eq, PartialEq, Copy, Clone, FromPrimitive, ToPrimitive)]
#[repr(u16)]
#[brw(repr = u16)]
pub enum Status {
    Valid = 0b1000_0000_0000_0000,
    Invalid = 0b0000_0000_0000_0000,
}

impl Status {
    #[inline]
    pub fn is_valid(&self) -> bool {
        *self == Self::Valid
    }
}

impl Default for Status {
    fn default() -> Self {
        Self::Valid
    }
}

impl Display for Status {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", match *self {
            Self::Valid => "valid",
            Self::Invalid => "invalid"
        })
    }
}


#[cfg(test)]
mod tests {
    use super::*;

    use test_log::test;
    use log::info;

    #[test]
    fn test_is_valid() {
       let cases = vec![
           (Status::Valid, true),
           (Status::Invalid, false)
       ];

        for (status, want) in cases {
            info!("Checking is_valid for {status:?}, want {want:?}");
            let got = status.is_valid();
            assert_eq!(want, got);
        }
    }

    #[test]
    fn test_default() {
        let want = Status::Valid;
        let got = Status::default();
        assert_eq!(want, got);
    }

    #[test]
    fn test_display() {
        let cases = vec![
            (Status::Valid, "valid"),
            (Status::Invalid, "invalid")
        ];

        for (status, want) in cases {
            info!("Displaying {status:?}, want {want:?}");
            let got = format!("{status}");
            assert_eq!(want, got);
        }
    }
}
