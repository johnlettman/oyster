#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

use binrw::BinRead;
use modular_bitfield::prelude::*;

/// Column data segment from a LIDAR packet.
#[derive(Debug, Clone, Eq, PartialEq)]
pub enum DataBlock {
    DualReturns(DualReturns),
    SingleReturns(SingleReturns),
    SingleReturnsLowDataRate(SingleReturnsLowDataRate),
    FuSaTwoWordPixel(FuSaTwoWordPixel),
}

#[bitfield(bits = 128)]
#[derive(Debug, BinRead, Eq, PartialEq, Clone)]
#[br(map = Self::from_bytes)]
#[cfg_attr(feature = "pyo3", pyclass)]
pub struct DualReturns {
    #[skip]
    __: B16,
    pub near_ir: B16,
    pub signal2: B16,
    pub signal: B16,
    pub reflectivity2: B8,
    #[skip]
    __: B5,
    pub range2: B19,
    pub reflectivity: B8,
    #[skip]
    __: B5,
    pub range: B19,
}

#[bitfield(bits = 96)]
#[derive(Debug, BinRead, Eq, PartialEq, Clone)]
#[br(map = Self::from_bytes)]
#[cfg_attr(feature = "pyo3", pyclass)]
pub struct SingleReturns {
    #[skip]
    __: B16,
    pub near_ir: B16,
    pub signal: B16,
    #[skip]
    __: B8,
    pub reflectivity: B8,
    #[skip]
    __: B13,
    pub range: B19,
}

#[bitfield(bits = 32)]
#[derive(Debug, BinRead, Eq, PartialEq, Clone)]
#[br(map = Self::from_bytes)]
#[cfg_attr(feature = "pyo3", pyclass)]
pub struct SingleReturnsLowDataRate {
    pub near_ir: B8,
    pub reflectivity: B8,
    #[skip]
    __: B1,
    pub range: B15,
}

#[bitfield(bits = 64)]
#[derive(Debug, BinRead, Eq, PartialEq, Clone)]
#[br(map = Self::from_bytes)]
#[cfg_attr(feature = "pyo3", pyclass)]
pub struct FuSaTwoWordPixel {
    #[skip]
    __: B8,
    pub reflectivity2: B8,
    #[skip]
    __: B1,
    pub range2: B15,
    pub near_ir: B8,
    pub reflectivity: B8,
    #[skip]
    __: B1,
    pub range: B15,
}
