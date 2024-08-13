use binrw::{BinRead, BinResult, BinWrite, Endian, Error};
use std::io::{Error as IOError, ErrorKind::InvalidData, Read, Seek, Write};

const MASK_RANGE: u16 = 0b1111_1111_1111_1110;
const SHIFT_RANGE: usize = 1;
const MAX_RANGE: u16 = 0b0111_1111_1111_1111;
const ERR_RANGE_INVALID: &'static str = "Invalid range value";

#[inline]
fn valid_range(range: u16) -> bool {
    range <= MAX_RANGE
}

#[derive(Debug, Eq, PartialEq, Clone, derive_new::new)]
pub struct SingleLowDataRateReturnBlock {
    pub range: u16,
    pub reflectivity: u8,
    pub near_ir: u8,
}

impl BinRead for SingleLowDataRateReturnBlock {
    type Args<'a> = ();

    fn read_options<R: Read + Seek>(
        reader: &mut R,
        endian: Endian,
        _: Self::Args<'_>,
    ) -> BinResult<Self> {
        let range = (u16::read_options(reader, endian, ())? & MASK_RANGE) >> SHIFT_RANGE;
        let reflectivity = u8::read_options(reader, endian, ())?;
        let near_ir = u8::read_options(reader, endian, ())?;
        Ok(Self { range, reflectivity, near_ir })
    }
}

impl BinWrite for SingleLowDataRateReturnBlock {
    type Args<'a> = ();

    fn write_options<W: Write + Seek>(
        &self,
        writer: &mut W,
        endian: Endian,
        _: Self::Args<'_>,
    ) -> BinResult<()> {
        if !valid_range(self.range) {
            let pos = writer.stream_position()?;
            return Err(Error::Custom {
                pos,
                err: Box::new(IOError::new(InvalidData, ERR_RANGE_INVALID)),
            });
        }

        let raw_range = (self.range & MAX_RANGE) << SHIFT_RANGE;
        raw_range.write_options(writer, endian, ())?;
        self.reflectivity.write_options(writer, endian, ())?;
        self.near_ir.write_options(writer, endian, ())?;
        Ok(())
    }
}
