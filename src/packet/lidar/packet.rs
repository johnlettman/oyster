use crate::packet::lidar::Header;
use crate::types::{ShotLimitingStatus, ThermalShutdownStatus};
use binrw::prelude::*;
use modular_bitfield::prelude::*;
use pyo3::pyclass;

/// Represents the data packet header.
#[derive(Debug, BinRead, BinWrite, Clone, Eq, PartialEq)]
#[brw(little)]
#[cfg_attr(feature = "pyo3", pyclass)]
pub struct Packet {
    header: Header,
}
