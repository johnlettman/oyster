use std::io::{Read, Seek, SeekFrom, Write, Error as IOError, ErrorKind::InvalidData};
use binrw::{BinRead, BinResult, BinWrite, Endian, Error};
use modular_bitfield::bitfield;
use modular_bitfield::prelude::{B13, B16, B19, B8};
use pyo3::pyclass;

const MASK_RANGE: u32 = 0b1111_1111_1111_1111_1110_0000_0000_0000;
const SHIFT_RANGE: usize = 13;
const MAX_RANGE: u32 = 0b111_1111_1111_1111_1111;
const ERR_RANGE_INVALID: &'static str = "Invalid range value";

#[inline]
fn valid_range(range: u32) -> bool {
    range <= MAX_RANGE
}

#[derive(Debug, BinRead, Eq, PartialEq, Clone)]
pub struct SingleReturnBlock {
    pub range: u32,
    pub reflectivity: u8,
    pub signal: u16,
    pub near_ir: u16
}

impl BinRead for SingleReturnBlock {
    type Args<'a> = ();

    fn read_options<R: Read + Seek>(reader: &mut R, endian: Endian, _: Self::Args<'_>) -> BinResult<Self> {
        let range = (u32::read_options(reader, endian, ())? & MASK_RANGE) >> SHIFT_RANGE;
        let reflectivity = u8::read(reader)?;

        // skip the reserved 8-bit chunk
        let pos = reader.stream_position()?;
        reader.seek(SeekFrom::Current(1)).map_err(|e| Error::Custom { pos, err: Box::new(e) })?;

        let signal = u16::read_options(reader, endian, ())?;
        let near_ir = u16::read_options(reader, endian, ())?;

        // skip the reserved 16-bit chunk
        let pos = reader.stream_position()?;
        reader.seek(SeekFrom::Current(2)).map_err(|e| Error::Custom { pos, err: Box::new(e) })?;

        Ok(Self { range, reflectivity, signal, near_ir })
    }
}

impl BinWrite for SingleReturnBlock {
    type Args<'a> = ();

    fn write_options<W: Write + Seek>(&self, writer: &mut W, endian: Endian, _: Self::Args<'_>) -> BinResult<()> {
        if !valid_range(self.range) {
            let pos = writer.stream_position()?;
            return Err(Error::Custom {
                pos,
                err: Box::new(IOError::new(InvalidData, ERR_RANGE_INVALID))
            })
        }

        let raw_range = (self.range & MAX_RANGE) << SHIFT_RANGE;
        raw_range.write_options(writer, endian, ())?;
        self.reflectivity.write(writer)?;

        // skip the reserved 8-bit chunk
        let pos = writer.stream_position()?;
        writer.seek(SeekFrom::Current(1)).map_err(|e| Error::Custom { pos, err: Box::new(e) })?;

        self.signal.write_options(writer, endian, ())?;
        self.near_ir.write_options(writer, endian, ())?;

        // skip the reserved 16-bit chunk
        let pos = writer.stream_position()?;
        writer.seek(SeekFrom::Current(2)).map_err(|e| Error::Custom { pos, err: Box::new(e) })?;

        Ok(())
    }

}
