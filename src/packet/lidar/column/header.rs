use binrw::{BinRead, BinResult, BinWrite, Endian, Error};
use chrono::{DateTime, Utc};
use std::io::{Error as IOError, ErrorKind::InvalidData, Read, Seek, Write};
use crate::packet::lidar::column::status::Status;

#[derive(Debug, Eq, PartialEq, Clone)]
pub struct Header {
    pub timestamp: DateTime<Utc>,
    pub measurement_id: u16,
    pub status: Status,
}

impl Header {
    #[inline]
    pub fn is_valid(&self) -> bool {
        self.status.is_valid()
    }
}

impl Default for Header {
    fn default() -> Self {
        Self {
            timestamp: Utc::now(),
            measurement_id: 0,
            status: Status::default()
        }
    }
}

impl BinRead for Header {
    type Args<'a> = ();

    fn read_options<R: Read + Seek>(
        reader: &mut R,
        endian: Endian,
        _: Self::Args<'_>,
    ) -> BinResult<Self> {
        let timestamp = DateTime::from_timestamp_nanos(i64::read_options(reader, endian, ())?);
        let measurement_id = u16::read_options(reader, endian, ())?;
        let status = Status::read_options(reader, endian, ())?;

        Ok(Self { timestamp, measurement_id, status })
    }
}

impl BinWrite for Header {
    type Args<'a> = ();

    fn write_options<W: Write + Seek>(
        &self,
        writer: &mut W,
        endian: Endian,
        _: Self::Args<'_>,
    ) -> BinResult<()> {
        let pos = writer.stream_position()?;
        let timestamp_raw = self.timestamp.timestamp_nanos_opt().ok_or(Error::Custom {
            pos,
            err: Box::new(IOError::new(InvalidData, "Unable to convert timestamp")),
        })?;

        timestamp_raw.write_options(writer, endian, ())?;
        self.measurement_id.write_options(writer, endian, ())?;
        self.status.write_options(writer, endian,())?;

        Ok(())
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
            (Header { timestamp: Utc::now(), measurement_id: 0, status: Status::Valid }, true),
            (Header { timestamp: Utc::now(), measurement_id: 0, status: Status::Invalid }, false),
        ];

        for (header, want) in cases {
            info!("Checking is_valid for {header:?}, want {want:?}");
            let got = header.is_valid();
            assert_eq!(want, got);
        }
    }





}
