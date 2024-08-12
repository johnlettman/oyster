#![crate_type = "lib"]

use pyo3::prelude::*;

mod build;
mod client;
pub mod metadata;
pub mod packet;
pub mod ros;
mod scan;
pub mod types;
mod util;

#[cfg(feature = "pyo3")]
use types::pyo3_register_module_types;

#[cfg(feature = "pyo3")]
use packet::pyo3_register_module_packet;

#[cfg(feature = "pyo3")]
#[pymodule]
fn oyster(m: &Bound<'_, PyModule>) -> PyResult<()> {
    pyo3_register_module_types(m)?;
    pyo3_register_module_packet(m)?;
    Ok(())
}
