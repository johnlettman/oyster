#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

mod column;
mod header;
mod status;
mod block;

pub use block::{
    Block, SingleReturnBlock, SingleLowDataRateReturnBlock, DualReturnBlock, FuSaDualReturnBlock
};
pub use column::Column;
pub use header::Header;

#[cfg(feature = "pyo3")]
pub(crate) fn pyo3_register_module_packet_lidar_column(
    parent_module: &Bound<'_, PyModule>,
) -> PyResult<()> {
    let module = PyModule::new_bound(parent_module.py(), "column")?;
    module.add_class::<DualReturnBlock>()?;
    module.add_class::<FuSaDualReturnBlock>()?;
    module.add_class::<SingleReturnBlock>()?;
    module.add_class::<SingleLowDataRateReturnBlock>()?;
    module.add_class::<Header>()?;

    parent_module.add("column", module)?;
    Ok(())
}
