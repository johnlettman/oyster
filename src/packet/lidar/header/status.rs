use crate::types::{ShotLimitingStatus, ThermalShutdownStatus};
use binrw::{
    meta::{EndianKind, ReadEndian, WriteEndian},
    BinRead, BinResult, BinWrite, Endian, Error,
};
use num_traits::{FromPrimitive, ToPrimitive};
use std::io::{Error as IOError, ErrorKind::InvalidData, Read, Seek, Write};

const MASK_SHOT_LIMITING: u8 = 0b1111_0000;
const SHIFT_SHOT_LIMITING: usize = 4;
const ERR_SHOT_LIMITING: &'static str = "Invalid shot limiting value";

const MASK_THERMAL_SHUTDOWN: u8 = 0b0000_1111;
const ERR_THERMAL_SHUTDOWN: &'static str = "Invalid thermal shutdown value";

/// The header status, which contains Shot Limiting and Thermal Shutdown Status.
#[derive(Debug, Clone, Copy, Eq, PartialEq, derive_new::new)]
pub struct Status {
    /// Indicates the shot limiting status of the sensor.
    pub shot_limiting: ShotLimitingStatus,

    /// Indicates whether thermal shutdown is imminent.
    pub thermal_shutdown: ThermalShutdownStatus,
}

impl Default for Status {
    /// The default value for [Status]:
    /// - [Status::shot_limiting]: [ShotLimitingStatus::Normal]
    /// - [Status::thermal_shutdown]: [ThermalShutdownStatus::Normal]
    fn default() -> Self {
        Self::new(ShotLimitingStatus::Normal, ThermalShutdownStatus::Normal)
    }
}

impl ReadEndian for Status {
    const ENDIAN: EndianKind = EndianKind::None;
}

impl WriteEndian for Status {
    const ENDIAN: EndianKind = EndianKind::None;
}

impl BinRead for Status {
    type Args<'a> = ();

    fn read<R: Read + Seek>(reader: &mut R) -> BinResult<Self>
    where
        Self: ReadEndian,
    {
        let raw = u8::read(reader)?;
        let pos = reader.stream_position()?;

        let raw_shot_limiting = (raw & MASK_SHOT_LIMITING) >> SHIFT_SHOT_LIMITING;
        let shot_limiting =
            ShotLimitingStatus::from_u8(raw_shot_limiting).ok_or(Error::Custom {
                pos,
                err: Box::new(IOError::new(InvalidData, ERR_SHOT_LIMITING)),
            })?;

        let raw_thermal_shutdown = raw & MASK_THERMAL_SHUTDOWN;
        let thermal_shutdown =
            ThermalShutdownStatus::from_u8(raw_thermal_shutdown).ok_or(Error::Custom {
                pos,
                err: Box::new(IOError::new(InvalidData, ERR_THERMAL_SHUTDOWN)),
            })?;

        Ok(Self { shot_limiting, thermal_shutdown })
    }

    #[inline]
    fn read_options<R: Read + Seek>(
        reader: &mut R,
        _: Endian,
        _: Self::Args<'_>,
    ) -> BinResult<Self> {
        Self::read(reader)
    }
}

impl BinWrite for Status {
    type Args<'a> = ();

    fn write<W: Write + Seek>(&self, writer: &mut W) -> BinResult<()>
    where
        Self: WriteEndian,
    {
        let pos = writer.stream_position()?;

        let raw_shot_limiting = self.shot_limiting.to_u8().ok_or(Error::Custom {
            pos,
            err: Box::new(IOError::new(InvalidData, ERR_SHOT_LIMITING)),
        })? << SHIFT_SHOT_LIMITING;

        let raw_thermal_shutdown = self.thermal_shutdown.to_u8().ok_or(Error::Custom {
            pos,
            err: Box::new(IOError::new(InvalidData, ERR_THERMAL_SHUTDOWN)),
        })?;

        let raw = raw_shot_limiting | raw_thermal_shutdown;
        raw.write(writer)
    }

    #[inline]
    fn write_options<W: Write + Seek>(
        &self,
        writer: &mut W,
        _: Endian,
        _: Self::Args<'_>,
    ) -> BinResult<()> {
        self.write(writer)
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::io::Cursor;

    use log::info;
    use test_log::test;

    #[test]
    fn test_default() {
        let want = Status::new(ShotLimitingStatus::Normal, ThermalShutdownStatus::Normal);
        let got = Status::default();
        assert_eq!(want, got);
    }

    const BINARY_CASES: [(Status, [u8; 1]); 3] = [
        (
            Status {
                shot_limiting: ShotLimitingStatus::Normal,
                thermal_shutdown: ThermalShutdownStatus::Normal,
            },
            [0x00],
        ),
        (
            Status {
                shot_limiting: ShotLimitingStatus::Reduction30to40,
                thermal_shutdown: ThermalShutdownStatus::Normal,
            },
            [0x50],
        ),
        (
            Status {
                shot_limiting: ShotLimitingStatus::Normal,
                thermal_shutdown: ThermalShutdownStatus::Imminent,
            },
            [0x01],
        ),
    ];

    #[test]
    fn test_parse() {
        for (want, bytes) in BINARY_CASES {
            info!("Parsing {bytes:?}, expecting {want:?}");
            let mut cursor = Cursor::new(bytes);

            let got = Status::read(&mut cursor).expect("It should not return an error");
            assert_eq!(want, got);
        }
    }

    #[test]
    fn test_write() {
        for (status, want) in BINARY_CASES {
            info!("Writing {status:?}, expecting {want:?}");
            let mut cursor = Cursor::new(Vec::new());
            status.write(&mut cursor).expect("It should not return an error");

            let inner = cursor.into_inner();
            let got = inner.as_slice();
            assert_eq!(want, got);
        }
    }
}
