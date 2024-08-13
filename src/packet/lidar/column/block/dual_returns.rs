use binrw::{BinRead, BinResult, BinWrite, Endian, Error};
use std::io::{Error as IOError, ErrorKind::InvalidData, Read, Seek, SeekFrom, Write};

const MASK_RANGE: u32 = 0b1111_1111_1111_1111_1110_0000_0000_0000;
const SHIFT_RANGE: usize = 13;
const MAX_RANGE: u32 = 0b111_1111_1111_1111_1111;
const ERR_RANGE_INVALID: &'static str = "Invalid range value";

const MASK_REFLECTIVITY: u32 = 0b0000_0000_0000_0000_0000_0000_1111_1111;

#[inline]
fn valid_range(range: u32) -> bool {
    range <= MAX_RANGE
}

#[derive(Debug, Eq, PartialEq, Clone, derive_new::new)]
pub struct DualReturnBlock {
    pub range: u32,
    pub reflectivity: u8,

    pub range2: u32,
    pub reflectivity2: u8,

    pub signal: u16,
    pub signal2: u16,

    pub near_ir: u16,
}

impl DualReturnBlock {
    fn read_rr_block<R: Read + Seek>(reader: &mut R, endian: Endian) -> BinResult<(u32, u8)> {
        let raw = u32::read_options(reader, endian, ())?;
        let range = (raw & MASK_RANGE) >> SHIFT_RANGE;
        let reflectivity = (raw & MASK_REFLECTIVITY) as u8;
        Ok((range, reflectivity))
    }

    fn write_rr_block<W: Write + Seek>(
        range: u32,
        reflectivity: u8,
        writer: &mut W,
        endian: Endian,
    ) -> BinResult<()> {
        if !valid_range(range) {
            let pos = writer.stream_position()?;
            return Err(Error::Custom {
                pos,
                err: Box::new(IOError::new(InvalidData, ERR_RANGE_INVALID)),
            });
        }

        let raw = reflectivity as u32 | (range & MASK_RANGE) << SHIFT_RANGE;
        raw.write_options(writer, endian, ())
    }
}

impl Default for DualReturnBlock {
    fn default() -> Self {
        Self {
            range: 0,
            reflectivity: 0,
            range2: 0,
            reflectivity2: 0,
            signal: 0,
            signal2: 0,
            near_ir: 0,
        }
    }
}

impl BinRead for DualReturnBlock {
    type Args<'a> = ();

    fn read_options<R: Read + Seek>(
        reader: &mut R,
        endian: Endian,
        _: Self::Args<'_>,
    ) -> BinResult<Self> {
        // obtain the first range-reflectivity block
        let (range, reflectivity) = Self::read_rr_block(reader, endian)?;

        // obtain the second range-reflectivity block
        let (range2, reflectivity2) = Self::read_rr_block(reader, endian)?;

        // read the first signal
        let signal = u16::read_options(reader, endian, ())?;

        // read the second signal
        let signal2 = u16::read_options(reader, endian, ())?;

        // read near-IR
        let near_ir = u16::read_options(reader, endian, ())?;

        // skip the reserved 16-bit chunk
        let pos = reader.stream_position()?;
        reader.seek(SeekFrom::Current(2)).map_err(|e| Error::Custom { pos, err: Box::new(e) })?;

        Ok(Self { range, reflectivity, range2, reflectivity2, signal, signal2, near_ir })
    }
}

impl BinWrite for DualReturnBlock {
    type Args<'a> = ();

    fn write_options<W: Write + Seek>(
        &self,
        writer: &mut W,
        endian: Endian,
        _: Self::Args<'_>,
    ) -> BinResult<()> {
        // write the first range-reflectivity block
        Self::write_rr_block(self.range, self.reflectivity, writer, endian)?;

        // write the second range-reflectivity block
        Self::write_rr_block(self.range2, self.reflectivity2, writer, endian)?;

        // write the first signal
        self.signal.write_options(writer, endian, ())?;

        // write the second signal
        self.signal2.write_options(writer, endian, ())?;

        // write near-IR
        self.near_ir.write_options(writer, endian, ())?;

        // skip the reserved 16-bit chunk
        let pos = writer.stream_position()?;
        writer.seek(SeekFrom::Current(2)).map_err(|e| Error::Custom { pos, err: Box::new(e) })?;

        Ok(())
    }
}
