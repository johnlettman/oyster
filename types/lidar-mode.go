package types

import (
	"github.com/johnlettman/oyster/util"
)

// LIDARMode represents the horizontal resolution and rotation rate of the sensor.
// The effective range of the sensor is increased by 15-20% for every halving of the number of points gathered.
// For example, LidarMode512x10 has a 15-20% longer range than LidarMode512x20.
// For additional information, refer to [lidar_mode].
//
// [lidar_mode]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html?highlight=512x10#lidar-mode
type LIDARMode int

const (
	LidarModeUnknown LIDARMode = iota // unspecified
	LidarMode512x10                   // 10 scans of 512 columns per second
	LidarMode512x20                   // 20 scans of 512 columns per second
	LidarMode1024x10                  // 10 scans of 1024 columns per second
	LidarMode1024x20                  // 20 scans of 1024 columns per second
	LidarMode2048x10                  // 10 scans of 2048 columns per second
	LidarMode4096x5                   // 5 scans of 4096 columns per second
)

var (
	// lidarModeStringKV maps LIDARMode values to their respective string representations.
	lidarModeStringKV = map[LIDARMode]string{
		LidarModeUnknown: "unknown",
		LidarMode512x10:  "512x10",
		LidarMode512x20:  "512x20",
		LidarMode1024x10: "1024x10",
		LidarMode1024x20: "1024x20",
		LidarMode2048x10: "2048x10",
		LidarMode4096x5:  "4096x5",
	}

	// lidarModeGoStringKV maps LIDARMode values to their respective Go syntax representations.
	lidarModeGoStringKV = map[LIDARMode]string{
		LidarModeUnknown: "LidarModeUnknown",
		LidarMode512x10:  "LidarMode512x10",
		LidarMode512x20:  "LidarMode512x20",
		LidarMode1024x10: "LidarMode1024x10",
		LidarMode1024x20: "LidarMode1024x20",
		LidarMode2048x10: "LidarMode2048x10",
		LidarMode4096x5:  "LidarMode4096x5",
	}

	// lidarModeTextKV maps LIDARMode values to their respective text representations.
	lidarModeTextKV = lidarModeStringKV

	// lidarModeTextVK is a variable that stores the reverse mapping of the lidarModeStringKV map.
	// It maps string representations of LIDARMode values to their respective LIDARMode values.
	lidarModeTextVK = util.ReverseMap(lidarModeTextKV)
)

// String returns the string representation of a LIDARMode value.
// If no match is found, it returns "unknown" as the default string representation.
func (m LIDARMode) String() string {
	if s, ok := lidarModeStringKV[m]; ok {
		return s
	}

	return lidarModeStringKV[LidarModeUnknown]
}

// GoString returns the Go syntax representation of a LIDARMode value.
// If no match is found, it returns "LidarModeUnknown" as the default string representation.
func (m LIDARMode) GoString() string {
	if s, ok := lidarModeGoStringKV[m]; ok {
		return s
	}

	return lidarModeGoStringKV[LidarModeUnknown]
}

// MarshalText returns the text representation of a LIDARMode value.
//   - If the LIDARMode has a matching string representation in the lidarModeStringKV map,
//     it returns the byte slice of that string representation.
//   - If no match is found, it returns nil.
//
// The error returned is always nil.
func (m LIDARMode) MarshalText() ([]byte, error) {
	if s, ok := lidarModeTextKV[m]; ok {
		return []byte(s), nil
	}

	return []byte{}, nil
}

// UnmarshalText unmarshals the given text into a LIDARMode value.
//   - If the string representation of the text exists in the lidarModeTextVK map,
//     it assigns the corresponding LIDARMode value to the receiver pointer.
//   - Otherwise, it assigns LidarModeUnknown to the receiver pointer.
//
// The error returned is always nil.
func (m *LIDARMode) UnmarshalText(text []byte) error {
	if mode, ok := lidarModeTextVK[string(text)]; ok {
		*m = mode
	} else {
		*m = LidarModeUnknown
	}

	return nil
}

// Columns returns the number of columns for a given LIDARMode value.
// It returns 0 if the LIDARMode is unknown or not specified.
func (m LIDARMode) Columns() int {
	switch m {
	default:
		fallthrough
	case LidarModeUnknown:
		return 0
	case LidarMode512x10, LidarMode512x20:
		return 512
	case LidarMode1024x10, LidarMode1024x20:
		return 1024
	case LidarMode2048x10:
		return 2048
	case LidarMode4096x5:
		return 4096
	}
}

// Frequency returns the frequency (number of scans per second) for a given LIDARMode value.
// It returns 0 if the LIDARMode is unknown or not specified.
func (m LIDARMode) Frequency() int {
	switch m {
	default:
		fallthrough
	case LidarModeUnknown:
		return 0
	case LidarMode512x20, LidarMode1024x20:
		return 20
	case LidarMode512x10, LidarMode1024x10, LidarMode2048x10:
		return 10
	case LidarMode4096x5:
		return 5
	}
}
