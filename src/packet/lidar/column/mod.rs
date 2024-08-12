#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

mod column;
mod data_block;
mod header;
mod status;

pub use column::Column;
pub use data_block::{
    DataBlock, DualReturns, FuSaTwoWordPixel, SingleReturns, SingleReturnsLowDataRate,
};
pub use header::HeaderBlock;

#[cfg(feature = "pyo3")]
pub(crate) fn pyo3_register_module_packet_lidar_column(
    parent_module: &Bound<'_, PyModule>,
) -> PyResult<()> {
    let module = PyModule::new_bound(parent_module.py(), "column")?;
    module.add_class::<DualReturns>()?;
    module.add_class::<FuSaTwoWordPixel>()?;
    module.add_class::<SingleReturns>()?;
    module.add_class::<SingleReturnsLowDataRate>()?;
    module.add_class::<HeaderBlock>()?;
    module.add_class::<Column>()?;

    parent_module.add("column", module)?;
    Ok(())
}
