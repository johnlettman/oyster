use nalgebra::{Matrix4, Vector3};
use std::error::Error;
use std::f64::consts::PI;

pub struct XyzLut {
    direction: Vec<Vector3<f64>>,
    offset: Vec<Vector3<f64>>,
}

impl XyzLut {
    pub fn generate(
        width: usize,
        height: usize,
        range_unit: f64,
        beam_to_lidar_transform: Matrix4<f64>,
        transform: Matrix4<f64>,
        azimuth_angles: Vec<f64>,
        altitude_angles: Vec<f64>,
    ) -> Result<XyzLut, Box<dyn Error>> {
        if width == 0 || height == 0 {
            return Err("LUT dimensions must be greater than zero".into());
        }

        let size = width * height;
        let azimuth_angles_len = azimuth_angles.len();
        let altitude_angles_len = altitude_angles.len();

        if (azimuth_angles_len != height || altitude_angles_len != height)
            && (azimuth_angles_len != size || altitude_angles_len != size)
        {
            return Err("unexpected scan dimensions".into());
        }

        let beam_to_lidar_euclidean_distance: f64 = {
            let x = beam_to_lidar_transform[(0, 3)];
            let z = beam_to_lidar_transform[(2, 3)];

            if z != 0.0 {
                (x.powi(2) + z.powi(2)).sqrt()
            } else {
                x
            }
        };

        let mut direction = Vec::with_capacity(size);
        let mut offset = Vec::with_capacity(size);

        let mut encoder = vec![0.0; size];
        let mut azimuth = vec![0.0; size];
        let mut altitude = vec![0.0; size];

        if azimuth_angles_len == height && altitude_angles_len == height {
            let azimuth_radians = 2.0 * PI / width as f64;

            for v in 0..width {
                for u in 0..height {
                    let i = u * width + v;
                    encoder[i] = 2.0 * PI - (v as f64 * azimuth_radians);
                    azimuth[i] = -azimuth_angles[u] * PI / 180.0;
                    altitude[i] = altitude_angles[u] * PI / 180.0;
                }
            }
        } else if azimuth_angles_len == size && altitude_angles_len == size {
            for v in 0..width {
                for u in 0..height {
                    let i = u * width + v;
                    encoder[i] = 0.0;
                    azimuth[i] = azimuth_angles[i] * PI / 180.0;
                    altitude[i] = altitude_angles[i] * PI / 180.0;
                }
            }
        }

        for i in 0..size {
            let dir_x = (encoder[i] + azimuth[i]).cos() * altitude[i].cos();
            let dir_y = (encoder[i] + azimuth[i]).sin() * altitude[i].cos();
            let dir_z = altitude[i].sin();

            direction.push(Vector3::new(dir_x, dir_y, dir_z));

            let off_x = encoder[i].cos() * beam_to_lidar_transform[(0, 3)]
                - dir_x * beam_to_lidar_euclidean_distance;
            let off_y = encoder[i].sin() * beam_to_lidar_transform[(0, 3)]
                - dir_y * beam_to_lidar_euclidean_distance;
            let off_z = -dir_z * beam_to_lidar_euclidean_distance + beam_to_lidar_transform[(2, 3)];

            offset.push(Vector3::new(off_x, off_y, off_z));
        }

        let rot = transform.fixed_view::<3, 3>(0, 0).transpose();
        let trans = transform.fixed_view::<3, 1>(0, 3);

        for dir in &mut direction {
            *dir = rot * (*dir);
        }

        for off in &mut offset {
            *off = rot * (*off) + trans;
        }

        for dir in &mut direction {
            *dir *= range_unit;
        }

        for off in &mut offset {
            *off *= range_unit;
        }

        Ok(XyzLut { direction, offset })
    }
}
