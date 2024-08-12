#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

pub mod imu;
pub mod lidar;
pub mod profile;

pub use imu::Packet as ImuPacket;
pub use imu::Profile as ImuProfile;

pub use lidar::Packet as LidarPacket;
pub use lidar::Profile as LidarProfile;

use lidar::pyo3_register_module_packet_lidar;

#[cfg(feature = "pyo3")]
pub(crate) fn pyo3_register_module_packet(parent_module: &Bound<'_, PyModule>) -> PyResult<()> {
    let module = PyModule::new_bound(parent_module.py(), "packet")?;
    pyo3_register_module_packet_lidar(&module)?;

    parent_module.add("packet", &module)?;
    Ok(())
}
