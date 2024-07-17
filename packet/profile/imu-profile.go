package profile

// IMUProfile represents how sensor data from the Inertial Measurement Unit (IMU)
// is packaged and sent over the wire, e.g., via UDP packets.
// Currently, IMUProfileLegacy is the only valid value.
//
// For additional information, refer to [Ouster docs: udp_profile_imu].
//
// [Ouster docs: udp_profile_imu]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#udp-profile-imu
type IMUProfile uint8

const (
	IMUProfileLegacy IMUProfile = iota // The default and only configuration for IMU data packets.
)

// String returns the string representation of an IMUProfile value.
func (p IMUProfile) String() string {
	return "legacy"
}

// GoString returns the Go syntax representation of an IMUProfile value.
func (p IMUProfile) GoString() string {
	return "IMUProfileLegacy"
}

// MarshalText returns the text representation of an IMUProfile value.
func (p IMUProfile) MarshalText() ([]byte, error) {
	return []byte("LEGACY"), nil
}

// UnmarshalText sets the value of the receiver IMUProfile to IMUProfileLegacy.
// It always returns nil as there are no errors expected in this method.
func (p *IMUProfile) UnmarshalText(_ []byte) error {
	*p = IMUProfileLegacy
	return nil
}
