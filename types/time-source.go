package types

import (
	"github.com/johnlettman/oyster/util"
)

// TimeSource is the core clock source configuration for the Ouster sensor, giving
// timestamps for all LiDAR and IMU data with 10 nanosecond precision. It can be
// configured to use one out of three sources:
//
//   - TimeFromInternalOscillator:
//     An internal clock derived from a precise, low-drift oscillator.
//   - TimeFromSyncPulseIn:
//     An opto-isolated digital input from an external connector, programmable to an
//     external hardware trigger like a GPS or a frame sync from an industrial camera
//   - TimeFromPTP1588:
//     The IEEE 1588 Precision Time Protocol, which allows network-based time
//     configuration without additional hardware signals.
//
// For additional information, refer to [timestamp_mode].
//
// [timestamp_mode]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#timestamp-mode
type TimeSource uint8

const (
	// TimeFromUnspecified indicates that the time source for the Ouster sensor is not defined.
	TimeFromUnspecified TimeSource = iota

	// TimeFromInternalOscillator is a free-running counter based on the sensor's internal
	// oscillator. It counts seconds and nanoseconds since the sensor was turned on. The
	// time is reported at a nanosecond resolution, with a minimum increment of around 10 ns.
	TimeFromInternalOscillator

	// TimeFromSyncPulseIn is a counter synced to the SYNC_PULSE_IN input. It also counts
	// seconds (number of pulses) and nanoseconds since sensor turn on. If the
	// multipurpose_io_mode is set to INPUT_NMEA_UART, the seconds register jumps to the
	// time extracted from a NMEA $GPRMC message read on the multipurpose_io port. The time
	// is also reported at a nanosecond resolution with a minimum increment of around 10 ns.
	TimeFromSyncPulseIn

	// TimeFromPTP1588 allows synchronization with an external PTP master. It's a
	// monotonically increasing counter that starts counting seconds and nanoseconds since
	// startup. As soon as a 1588 sync event happens, the time will be updated to seconds
	// and nanoseconds since 1970. The counter must always count forward in time. If
	// another 1588 sync event happens the counter will either jump forward to match the
	// new time, or slow itself down. It is reported at nanosecond resolution, but the
	// minimum increment varies.
	TimeFromPTP1588
)

var (
	// timeSourceStringKV maps TimeSource values to their string representations.
	timeSourceStringKV = map[TimeSource]string{
		TimeFromUnspecified:        "unspecified",
		TimeFromInternalOscillator: "time from internal oscillator",
		TimeFromSyncPulseIn:        "time from sync pulse in",
		TimeFromPTP1588:            "time from PTP 1588",
	}

	// timeSourceGoStringKV maps TimeSource values to their Go syntax representations.
	timeSourceGoStringKV = map[TimeSource]string{
		TimeFromUnspecified:        "TimeFromUnspecified",
		TimeFromInternalOscillator: "TimeFromInternalOscillator",
		TimeFromSyncPulseIn:        "TimeFromSyncPulseIn",
		TimeFromPTP1588:            "TimeFromPTP1588",
	}

	// timeSourceTextKV maps TimeSource values to their text representations.
	timeSourceTextKV = map[TimeSource]string{
		TimeFromUnspecified:        "UNSPECIFIED",
		TimeFromInternalOscillator: "TIME_FROM_INTERNAL_OSC",
		TimeFromSyncPulseIn:        "TIME_FROM_SYNC_PULSE_IN",
		TimeFromPTP1588:            "TIME_FROM_PTP_1588",
	}

	// timeSourceTextVK maps string representations to TimeSource values.
	timeSourceTextVK = util.ReverseMap(timeSourceTextKV)
)

// String returns the string representation of a TimeSource value.
func (t TimeSource) String() string {
	if s, ok := timeSourceStringKV[t]; ok {
		return s
	}

	return timeSourceStringKV[TimeFromUnspecified]
}

// GoString returns the Go syntax representation of a TimeSource value.
func (t TimeSource) GoString() string {
	if s, ok := timeSourceGoStringKV[t]; ok {
		return s
	}

	return timeSourceGoStringKV[TimeFromUnspecified]
}

// MarshalText returns the text representation of a TimeSource value.
func (t TimeSource) MarshalText() ([]byte, error) {
	if text, ok := timeSourceTextKV[t]; ok {
		return []byte(text), nil
	}

	return []byte(timeSourceTextKV[TimeFromUnspecified]), nil
}

// UnmarshalText updates the value of a TimeSource pointer based on the text representation.
// It attempts to match the provided text with a corresponding TimeSource value.
//   - If a match is found, the pointer is updated to the matching TimeSource value.
//   - If no match is found, the pointer is updated to the TimeFromUnspecified value.
//
// This method does not return an error.
func (t *TimeSource) UnmarshalText(text []byte) error {
	if timeSource, ok := timeSourceTextVK[string(text)]; ok {
		*t = timeSource
	} else {
		*t = TimeFromUnspecified
	}

	return nil
}
