use crate::packet::lidar::column::{DataBlock, HeaderBlock};
use binrw::{BinRead, BinWrite};
use modular_bitfield::bitfield;
use pyo3::pyclass;

#[derive(Debug, BinRead, BinWrite, Eq, PartialEq, Clone)]
#[br(map = Self::from_bytes)]
#[cfg_attr(feature = "pyo3", pyclass)]
pub struct Column {
    header_block: HeaderBlock,
    data_block: DataBlock,
}
