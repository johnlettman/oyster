use binrw::__private::Required;
use binrw::meta::{EndianKind, ReadEndian, WriteEndian};
use binrw::{BinRead, BinResult};
use std::io::{Read, Seek};

/// Represents a set of alert flags.
#[derive(Debug, Clone, Copy, Eq, PartialEq, derive_new::new)]
pub struct AlertFlags {
    pub cursor: u8,
    pub cursor_overflow: bool,
    pub alerts_active: bool,
}

impl AlertFlags {
    const MIN_CURSOR: u8 = 0b0000_0000;
    const MAX_CURSOR: u8 = 0b0011_1111;
    const MASK_CURSOR: u8 = 0b1111_1100;
    const SHIFT_CURSOR: usize = 2;

    const FLAG_CURSOR_OVERFLOW: u8 = 0b0000_0010;
    const FLAG_ALERTS_ACTIVE: u8 = 0b0000_0001;
}

impl Default for AlertFlags {
    fn default() -> Self {
        Self::new(0, false, false)
    }
}

impl ReadEndian for AlertFlags {
    const ENDIAN: EndianKind = EndianKind::None;
}

impl WriteEndian for AlertFlags {
    const ENDIAN: EndianKind = EndianKind::None;
}

impl BinRead for AlertFlags {
    fn read<R: Read + Seek>(reader: &mut R) -> BinResult<Self>
    where
        Self: ReadEndian,
        for<'a> Self::Args<'a>: Required,
    {
        let raw = u8::read(reader)?;
    }
}
