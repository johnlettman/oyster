package profile

import (
	"github.com/johnlettman/oyster/packet/column"
	"github.com/johnlettman/oyster/packet/column/field"
	"github.com/johnlettman/oyster/util"
)

// LIDARProfile represents how sensor data from the LIDAR sensor
// is packaged and sent over the wire, e.g., via UDP packets.
type LIDARProfile uint8

const (
	// LIDARProfileLegacy is the legacy LIDAR profile (deprecated).
	LIDARProfileLegacy LIDARProfile = iota

	// LIDARProfileDualReturns represents the profile for dual returns from a LIDAR sensor.
	// It is encoded as "RNG19_RFL8_SIG16_NIR16_DUAL".
	// For additional information, refer to [RNG19_RFL8_SIG16_NIR16_DUAL Return Profile].
	//
	// [RNG19_RFL8_SIG16_NIR16_DUAL Return Profile]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/sensor_data/sensor-data.html#dual-return-v3-x
	LIDARProfileDualReturns

	// LIDARProfileSingleReturns represents the profile for single returns from a LIDAR sensor (default).
	// It is encoded as "RNG19_RFL8_SIG16_NIR16".
	// For additional information, refer to [RNG19_RFL8_SIG16_NIR16 Return Profile].
	//
	// [RNG19_RFL8_SIG16_NIR16 Return Profile]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/sensor_data/sensor-data.html#rng19-rfl8-sig16-nir16-return-profile
	LIDARProfileSingleReturns

	// LIDARProfileSingleReturnsLowDataRate is the profile for single returns at a reduced data rate from the LIDAR sensor.
	// It is encoded as "RNG15_RFL8_NIR8".
	// For additional information refer to [RNG15_RFL8_NIR8 Return Profile].
	//
	// [RNG15_RFL8_NIR8 Return Profile]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/sensor_data/sensor-data.html#rng15-rfl8-nir8-return-profile
	LIDARProfileSingleReturnsLowDataRate

	// LIDARProfileFuSaTwoWordPixel is the profile for the Functional Safety data packet format from the LIDAR sensor.
	// It is encoded as "FUSA_RNG15_RFL8_NIR8_DUAL".
	// For additional information, refer to [FUSA_RNG15_RFL8_NIR8_DUAL Return Profile].
	//
	// [FUSA_RNG15_RFL8_NIR8_DUAL Return Profile]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/sensor_data/sensor-data.html#fusa-rng15-rfl8-nir8-dual-return-profile
	LIDARProfileFuSaTwoWordPixel
)

var lidarProfileChannelProfileVK = map[LIDARProfile]column.Profile{
	LIDARProfileLegacy: {
		Fields: column.Entries{
			field.Flags:        {field.TypeUint8, 3, 0, 4},
			field.Reflectivity: {field.TypeUint16, 4, 0, 0},
			field.Signal:       {field.TypeUint16, 6, 0, 0},
			field.NearIR:       {field.TypeUint16, 8, 0, 0},
			field.Raw32Word1:   {field.TypeUint32, 0, 0, 0},
			field.Raw32Word2:   {field.TypeUint32, 4, 0, 0},
			field.Raw32Word3:   {field.TypeUint32, 8, 0, 0},
		},
		DataSize: 12,
	},
	LIDARProfileDualReturns: {
		Fields: column.Entries{
			field.Range:         {field.TypeUint32, 0, 0x0007FFFF, 0},
			field.Flags:         {field.TypeUint8, 2, 0b11111000, 3},
			field.Reflectivity:  {field.TypeUint8, 3, 0, 0},
			field.Range2:        {field.TypeUint32, 4, 0x0007ffff, 0},
			field.Flags2:        {field.TypeUint8, 6, 0b11111000, 3},
			field.Reflectivity2: {field.TypeUint8, 7, 0, 0},
			field.Signal:        {field.TypeUint16, 8, 0, 0},
			field.Signal2:       {field.TypeUint16, 10, 0, 0},
			field.NearIR:        {field.TypeUint16, 12, 0, 0},
			field.Raw32Word1:    {field.TypeUint32, 0, 0, 0},
			field.Raw32Word2:    {field.TypeUint32, 4, 0, 0},
			field.Raw32Word3:    {field.TypeUint32, 8, 0, 0},
			field.Raw32Word4:    {field.TypeUint32, 12, 0, 0},
		},
		DataSize: 16,
	},
	LIDARProfileSingleReturns: {
		Fields: column.Entries{
			field.Range:        {field.TypeUint32, 0, 0x0007ffff, 0},
			field.Flags:        {field.TypeUint8, 2, 0b11111000, 3},
			field.Reflectivity: {field.TypeUint8, 4, 0, 0},
			field.Signal:       {field.TypeUint16, 6, 0, 0},
			field.NearIR:       {field.TypeUint16, 8, 0, 0},
			field.Raw32Word1:   {field.TypeUint32, 0, 0, 0},
			field.Raw32Word2:   {field.TypeUint32, 4, 0, 0},
			field.Raw32Word3:   {field.TypeUint32, 8, 0, 0},
		},
		DataSize: 12,
	},
	LIDARProfileSingleReturnsLowDataRate: {
		Fields: column.Entries{
			field.Range:        {field.TypeUint32, 0, 0x7fff, -3},
			field.Flags:        {field.TypeUint8, 1, 0b10000000, 7},
			field.Reflectivity: {field.TypeUint8, 2, 0, 0},
			field.NearIR:       {field.TypeUint16, 2, 0xff00, 4},
			field.Raw32Word1:   {field.TypeUint32, 0, 0, 0},
		},
		DataSize: 4,
	},
	LIDARProfileFuSaTwoWordPixel: {
		Fields: column.Entries{
			field.Range:         {field.TypeUint32, 0, 0x7fff, -3},
			field.Flags:         {field.TypeUint8, 1, 0b10000000, 7},
			field.Reflectivity:  {field.TypeUint8, 2, 0xff, 0},
			field.NearIR:        {field.TypeUint16, 3, 0xff, -4},
			field.Range2:        {field.TypeUint32, 4, 0x7fff, -3},
			field.Flags2:        {field.TypeUint8, 5, 0b10000000, 7},
			field.Reflectivity2: {field.TypeUint8, 6, 0xff, 0},
			field.Raw32Word1:    {field.TypeUint32, 0, 0, 0},
			field.Raw32Word2:    {field.TypeUint32, 4, 0, 0},
		},
		DataSize: 8,
	},
}

var (
	// lidarProfileStringKV maps LIDARProfile values to their string representations.
	lidarProfileStringKV = map[LIDARProfile]string{
		LIDARProfileLegacy:                   "legacy",
		LIDARProfileDualReturns:              "dual-returns",
		LIDARProfileSingleReturns:            "single-returns",
		LIDARProfileSingleReturnsLowDataRate: "single-returns low-data-rate",
		LIDARProfileFuSaTwoWordPixel:         "FuSa two-word pixel",
	}

	// lidarProfileGoStringKV maps LIDARProfile values to their Go syntax representations.
	lidarProfileGoStringKV = map[LIDARProfile]string{
		LIDARProfileLegacy:                   "LIDARProfileLegacy",
		LIDARProfileDualReturns:              "LIDARProfileDualReturns",
		LIDARProfileSingleReturns:            "LIDARProfileSingleReturns",
		LIDARProfileSingleReturnsLowDataRate: "LIDARProfileSingleReturnsLowDataRate",
		LIDARProfileFuSaTwoWordPixel:         "LIDARProfileFuSaTwoWordPixel",
	}

	// lidarProfileTextKV maps LIDARProfile values to their text representations.
	lidarProfileTextKV = map[LIDARProfile]string{
		LIDARProfileLegacy:                   "LEGACY",
		LIDARProfileDualReturns:              "RNG19_RFL8_SIG16_NIR16_DUAL",
		LIDARProfileSingleReturns:            "RNG19_RFL8_SIG16_NIR16",
		LIDARProfileSingleReturnsLowDataRate: "RNG15_RFL8_NIR8",
		LIDARProfileFuSaTwoWordPixel:         "FUSA_RNG15_RFL8_NIR8_DUAL",
	}

	// lidarProfileTextVK maps string representations to LIDARProfile values.
	lidarProfileTextVK = util.ReverseMap(lidarProfileTextKV)
)

// String returns the string representation of an LIDARProfile value.
func (p LIDARProfile) String() string {
	if s, ok := lidarProfileStringKV[p]; ok {
		return s
	} else {
		return lidarProfileStringKV[LIDARProfileLegacy]
	}
}

// GoString returns the Go syntax representation of an LIDARProfile value.
func (p LIDARProfile) GoString() string {
	if s, ok := lidarProfileGoStringKV[p]; ok {
		return s
	} else {
		return lidarProfileGoStringKV[LIDARProfileLegacy]
	}
}

// MarshalText returns the text representation of an LIDARProfile value.
func (p LIDARProfile) MarshalText() ([]byte, error) {
	if s, ok := lidarProfileTextKV[p]; ok {
		return []byte(s), nil
	} else {
		return []byte(lidarProfileTextKV[p]), nil
	}
}

// UnmarshalText sets the value of the LIDARProfile receiver based on the given text.
// It expects the text to be a valid representation of a LIDARProfile value.
//   - If the text is recognized as a valid representation of a LIDARProfile value,
//     the LIDARProfile receiver is updated accordingly.
//   - If the text is not recognized, the LIDARProfile receiver is set to LIDARProfileLegacy.
//
// The function always returns nil.
func (p *LIDARProfile) UnmarshalText(text []byte) error {
	if profile, ok := lidarProfileTextVK[string(text)]; ok {
		*p = profile
	} else {
		*p = LIDARProfileLegacy
	}

	return nil
}

// ColumnProfile returns the column profile associated with the LIDARProfile value.
// If the LIDARProfile value is recognized, the corresponding column profile is returned.
// Otherwise, the column profile for LIDARProfileLegacy is returned.
func (p LIDARProfile) ColumnProfile() column.Profile {
	if c, ok := lidarProfileChannelProfileVK[p]; ok {
		return c
	} else {
		return lidarProfileChannelProfileVK[LIDARProfileLegacy]
	}
}

// ColumnFields returns the corresponding Field entries based on the value of LIDARProfile.
// If the LIDARProfile value is invalid, it returns nil.
func (p LIDARProfile) ColumnFields() column.Entries {
	return p.ColumnProfile().Fields
}

// ColumnFieldCount returns the number of field entries associated with the LIDARProfile value.
// It calculates the count by getting the length of the column fields obtained from the LIDARProfile.
func (p LIDARProfile) ColumnFieldCount() int {
	return len(p.ColumnFields())
}

func (p LIDARProfile) ColumnDataSize() int {
	return p.ColumnProfile().DataSize
}
