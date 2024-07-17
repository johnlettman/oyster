package types

import "github.com/johnlettman/oyster/util"

// ShotLimitingStatus represents the operating state of the Ouster sensor under
// high temperatures. It is used to manage the sensor's performance and lifespan.
//
// States include ShotLimitingNormal (0x00), ShotLimitingImminent (0x01), and
// ShotLimitingReduction0to10 (0x02 and greater).
//
// In ShotLimitingNormal state, sensor operates within range and precision
// specifications. In ShotLimitingImminent state, sensor is preparing to limit its
// shootings due to temperature increase in 30 seconds. After 30 seconds have elapsed
// and the temperature remains elevated, the sensor issues alert 0x0100000F and enters
// a ShotLimitingReduction0to10 and above state.
//
// In ShotLimitingReduction0to10 and above states, the sensor reduces laser power to
// manage thermal load, possibly degrading range and precision by  up to 30%.
// An adjacent state machine oversees thermal shutdown.
// Recovery to ShotLimitingNormal occurs if temperature drops during
// ShotLimitingImminent or ShotLimitingReduction0to10 and above states.
//
// For additional information, refer to [Shot Limiting].
//
// [Shot Limiting]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/sensor_operations/sensor-operations.html#shot-limiting
type ShotLimitingStatus uint8

const (
	ShotLimitingNormal          ShotLimitingStatus = 0x0 // Normal operation of the LIDAR.
	ShotLimitingImminent        ShotLimitingStatus = 0x1 // Approaching the shot limit threshold.
	ShotLimitingReduction0to10  ShotLimitingStatus = 0x2 // Number of shots emitted limited by 0 to 10%.
	ShotLimitingReduction10to20 ShotLimitingStatus = 0x3 // Number of shots emitted limited by 10 to 20%.
	ShotLimitingReduction20to30 ShotLimitingStatus = 0x4 // Number of shots emitted limited by 20 to 30%.
	ShotLimitingReduction30to40 ShotLimitingStatus = 0x5 // Number of shots emitted limited by 30 to 40%.
	ShotLimitingReduction40to50 ShotLimitingStatus = 0x6 // Number of shots emitted limited by 40 to 50%.
	ShotLimitingReduction50to60 ShotLimitingStatus = 0x7 // Number of shots emitted limited by 50 to 60%.
	ShotLimitingReduction60to70 ShotLimitingStatus = 0x8 // Number of shots emitted limited by 60 to 70%.
	ShotLimitingReduction70to75 ShotLimitingStatus = 0x9 // Number of shots emitted limited by 70 to 75%.
)

var (
	// shotLimitingStatusStringKV maps ShotLimitingStatus values to their string representations.
	shotLimitingStatusStringKV = map[ShotLimitingStatus]string{
		ShotLimitingNormal:          "normal",
		ShotLimitingImminent:        "imminent",
		ShotLimitingReduction0to10:  "reduction of 0 to 10%",
		ShotLimitingReduction10to20: "reduction of 10 to 20%",
		ShotLimitingReduction20to30: "reduction of 20 to 30%",
		ShotLimitingReduction30to40: "reduction of 30 to 40%",
		ShotLimitingReduction40to50: "reduction of 40 to 50%",
		ShotLimitingReduction50to60: "reduction of 50 to 60%",
		ShotLimitingReduction60to70: "reduction of 60 to 70%",
		ShotLimitingReduction70to75: "reduction of 70 to 75%",
	}

	// shotLimitingStatusGoStringKV maps ShotLimitingStatus values to their Go syntax representations.
	shotLimitingStatusGoStringKV = map[ShotLimitingStatus]string{
		ShotLimitingNormal:          "ShotLimitingNormal",
		ShotLimitingImminent:        "ShotLimitingImminent",
		ShotLimitingReduction0to10:  "ShotLimitingReduction0to10",
		ShotLimitingReduction10to20: "ShotLimitingReduction10to20",
		ShotLimitingReduction20to30: "ShotLimitingReduction20to30",
		ShotLimitingReduction30to40: "ShotLimitingReduction30to40",
		ShotLimitingReduction40to50: "ShotLimitingReduction40to50",
		ShotLimitingReduction50to60: "ShotLimitingReduction50to60",
		ShotLimitingReduction60to70: "ShotLimitingReduction60to70",
		ShotLimitingReduction70to75: "ShotLimitingReduction70to75",
	}

	// shotLimitingStatusTextKV maps ShotLimitingStatus values to their text representations.
	shotLimitingStatusTextKV = map[ShotLimitingStatus]string{
		ShotLimitingNormal:          "NORMAL",
		ShotLimitingImminent:        "SHOT_LIMITING_IMMINENT",
		ShotLimitingReduction0to10:  "SHOT_LIMITING_0_TO_10",
		ShotLimitingReduction10to20: "SHOT_LIMITING_10_TO_20",
		ShotLimitingReduction20to30: "SHOT_LIMITING_20_TO_30",
		ShotLimitingReduction30to40: "SHOT_LIMITING_30_TO_40",
		ShotLimitingReduction40to50: "SHOT_LIMITING_40_TO_50",
		ShotLimitingReduction50to60: "SHOT_LIMITING_50_TO_60",
		ShotLimitingReduction60to70: "SHOT_LIMITING_60_TO_70",
		ShotLimitingReduction70to75: "SHOT_LIMITING_70_TO_75",
	}

	shotLimitingStatusTextVK = util.ReverseMap(shotLimitingStatusTextKV)
)

// String returns the string representation of a ShotLimitingStatus value.
func (s ShotLimitingStatus) String() string {
	if str, ok := shotLimitingStatusStringKV[s]; ok {
		return str
	}

	return shotLimitingStatusStringKV[ShotLimitingNormal]
}

// GoString returns the Go syntax representation of a ShotLimitingStatus value.
func (s ShotLimitingStatus) GoString() string {
	if str, ok := shotLimitingStatusGoStringKV[s]; ok {
		return str
	}

	return shotLimitingStatusGoStringKV[ShotLimitingNormal]
}

// MarshalText returns the text representation of a ShotLimitingStatus value.
func (s ShotLimitingStatus) MarshalText() ([]byte, error) {
	if text, ok := shotLimitingStatusTextKV[s]; ok {
		return []byte(text), nil
	}

	return []byte(shotLimitingStatusTextKV[ShotLimitingNormal]), nil
}

// UnmarshalText unmarshals the text representation of a ShotLimitingStatus value.
// If the text matches a known status, the corresponding value is assigned to the
// ShotLimitingStatus pointer. Otherwise, the pointer is assigned ShotLimitingNormal.
// This method always returns nil error.
func (s *ShotLimitingStatus) UnmarshalText(text []byte) error {
	if shotLimitingStatus, ok := shotLimitingStatusTextVK[string(text)]; ok {
		*s = shotLimitingStatus
	} else {
		*s = ShotLimitingNormal
	}

	return nil
}
