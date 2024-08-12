use nalgebra::Matrix4;
use serde::{Deserialize, Serialize};

#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

/// Represents the intrinsic parameters of the LIDAR beam.
#[derive(Debug, Serialize, Deserialize, PartialEq)]
#[cfg_attr(feature = "pyo3", pyclass)]
pub struct BeamIntrinsics {
    /// The beam altitude angle offset, measured in degrees.
    pub(crate) beam_altitude_angles: Vec<f64>,

    /// The beam azimuth angle, measured in degrees.
    pub(crate) beam_azimuth_angles: Vec<f64>,

    /// The offset distance between the LIDAR origin and the beam origin, measured in millimeters.
    #[serde(rename = "lidar_origin_to_beam_origin_mm", skip_serializing_if = "Option::is_none")]
    pub(crate) lidar_origin_to_beam_origin: Option<f64>,

    /// The transformation matrix from the LiDAR origin coordinate frame to the LIDAR front optics.
    /// It is a 4x4 matrix stored in row-major order in a 16-element float64 array.
    pub(crate) beam_to_lidar_transform: Option<Matrix4<f64>>,
}

impl BeamIntrinsics {
    pub const fn new(
        beam_altitude_angles: Vec<f64>,
        beam_azimuth_angles: Vec<f64>,
        lidar_origin_to_beam_origin: Option<f64>,
        beam_to_lidar_transform: Option<Matrix4<f64>>,
    ) -> Self {
        Self {
            beam_altitude_angles,
            beam_azimuth_angles,
            lidar_origin_to_beam_origin,
            beam_to_lidar_transform,
        }
    }

    pub fn default_lidar_origin_to_beam_origin_for(product_line: &str) -> f64 {
        let mut lidar_origin_to_beam_origin: f64 = 12.163;

        if product_line.starts_with("OS-0-") {
            lidar_origin_to_beam_origin = 27.67;
        } else if product_line.starts_with("OS-1-") {
            lidar_origin_to_beam_origin = 15.806;
        } else if product_line.starts_with("OS-2-") {
            lidar_origin_to_beam_origin = 13.762;
        }

        lidar_origin_to_beam_origin
    }

    pub fn default_beam_to_lidar_transform_for(product_line: &str) -> Matrix4<f64> {
        let mut beam_to_lidar_transform: Matrix4<f64> = Matrix4::identity();
        beam_to_lidar_transform[(0, 3)] =
            BeamIntrinsics::default_lidar_origin_to_beam_origin_for(product_line);
        beam_to_lidar_transform
    }
}
