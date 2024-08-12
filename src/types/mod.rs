mod azimuth_window;
mod column_window;
mod full_scale_range;
mod lidar_mode;
mod multipurpose_io_mode;
mod baud_rate;
mod operating_mode;
mod polarity;
mod return_order;
mod shot_limiting_status;
mod thermal_shutdown_status;
mod timestamp_mode;

pub use azimuth_window::AzimuthWindow;
pub use column_window::ColumnWindow;
pub use full_scale_range::FullScaleRange;
pub use lidar_mode::LidarMode;
pub use multipurpose_io_mode::MultipurposeIoMode;
pub use baud_rate::BaudRate;
pub use operating_mode::OperatingMode;
pub use polarity::Polarity;
pub use return_order::ReturnOrder;
pub use shot_limiting_status::ShotLimitingStatus;
pub use thermal_shutdown_status::ThermalShutdownStatus;
pub use timestamp_mode::TimestampMode;

#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "pyo3")]
pub(crate) fn pyo3_register_module_types(parent_module: &Bound<'_, PyModule>) -> PyResult<()> {
    let types_module = PyModule::new_bound(parent_module.py(), "types")?;
    parent_module.add("types", types_module)?;
    Ok(())
}
