use std::io::{Read, Seek};
use binrw::Endian;
use binrw::prelude::*;
use modular_bitfield::prelude::*;

/// Represents the header metadata, which contains Initialization ID and Serial Number. 64 bits.
#[derive(Debug, Clone, Copy, Eq, PartialEq, derive_new::new)]
pub struct SensorInfo {
    /// Updates on every reinitialization, which may be triggered by the
    /// user or an error, and every reboot. 24 bits.
    /// `HeaderSensorInfo.serial_number`
    pub initialization_id: u32,

    /// Serial number of the sensor. This value is unique to each sensor and can be found
    /// on a sticker affixed to the top of the sensor. 40 bits
    pub serial_number: u64,
}

impl SensorInfo {
    const MASK_INITIALIZATION_ID: u64 = 0xFFFF_FF00_0000_0000;
    const SHIFT_INITIALIZATION_ID: usize = 40;

    const MASK_SERIAL_NUMBER: u64     = 0x0000_00FF_FFFF_FFFF;



}

impl Default for SensorInfo {
    fn default() -> Self {
        Self::new(0, 0)
    }
}

impl BinRead for SensorInfo {
    type Args<'a> = ();

    fn read_options<R: Read + Seek>(reader: &mut R, endian: Endian, args: Self::Args<'_>) -> BinResult<Self> {
        let raw = u64::read_options(reader, endian, args)?;

        let initialization_id =
            ((raw & Self::MASK_INITIALIZATION_ID) >> Self::SHIFT_INITIALIZATION_ID) as u32;
        let serial_number = raw & Self::MASK_SERIAL_NUMBER;

        Ok(Self { initialization_id, serial_number })
    }
}

#[cfg(test)]
mod tests {
    use std::io::Cursor;
    use super::*;

    use test_log::test;
    use log::info;

    const ENDIAN: Endian = Endian::Big;

    #[test]
    fn test_parse() {
        let cases: Vec<(SensorInfo, [u8; 8])> = vec![
            (
                SensorInfo::new(0xFFFF_FF, 0xFF_FFFF_FFFF),
                [
                    0xFF, 0xFF, 0xFF,
                    0xFF, 0xFF, 0xFF, 0xFF, 0xFF
                ]
            ),
        ];

        for (want, bytes) in cases {
            info!("Parsing {bytes:?}, expecting {bytes:?}");
            let mut cursor = Cursor::new(bytes);

            let got = SensorInfo::read_options(&mut cursor, ENDIAN, ())
                .expect("It should not return an error");

            assert_eq!(want, got);
        }



    }





}


