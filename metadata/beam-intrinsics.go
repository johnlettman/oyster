package metadata

// BeamIntrinsics represents the intrinsic parameters of the LiDAR beam.
type BeamIntrinsics struct {
	// BeamAltitudeAngles is the beam altitude angle offset, measured in degrees.
	BeamAltitudeAngles []float64 `json:"beam_altitude_angles,omitempty"`

	// BeamAzimuthAngles is the beam azimuth angle offsets, measured in degrees.
	BeamAzimuthAngles []float64 `json:"beam_azimuth_angles,omitempty"`

	// LIDAROriginToBeamOrigin is the offset distance between the LiDAR origin and the beam origin, measured in millimeters.
	LIDAROriginToBeamOrigin float64 `json:"lidar_origin_to_beam_origin_mm,omitempty"`

	// BeamToLIDARTransform represents the transformation matrix from the LiDAR origin coordinate frame to the LiDAR front optics.
	// It is a 4x4 matrix stored in row-major order in a 16-element float64 array.
	BeamToLIDARTransform [16]float64 `json:"beam_to_lidar_transform,omitempty"`
}
