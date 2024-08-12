#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

pub mod alert_flags;
pub mod column;
pub mod header;
mod packet;
mod profile;

pub use packet::Packet;
pub use profile::Profile;

pub use header::Header;

use column::pyo3_register_module_packet_lidar_column;

#[cfg(feature = "pyo3")]
pub(crate) fn pyo3_register_module_packet_lidar(
    parent_module: &Bound<'_, PyModule>,
) -> PyResult<()> {
    let module = PyModule::new_bound(parent_module.py(), "lidar")?;
    pyo3_register_module_packet_lidar_column(&module)?;
    parent_module.add("lidar", &module)?;
    Ok(())
}
