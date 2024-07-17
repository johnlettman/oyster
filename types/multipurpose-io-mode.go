package types

import (
	"github.com/johnlettman/oyster/util"
)

// MultipurposeIOMode controls the functionality of the multipurpose I/O port on the
// Ouster sensor, particularly with respect to the SYNC_PULSE_OUT signal. Its
// configuration parameters include:
//
//   - MultipurposeOff:
//     No SYNC_PULSE_OUT signal is output.
//   - MultipurposeInputNMEAUART:
//     The port is reconfigured as an input.
//   - MultipurposeOutputFromInternalOscillator:
//     SYNC_PULSE_OUT synchronized with the internal clock.
//   - MultipurposeOutputFromSyncPulseIn:
//     SYNC_PULSE_OUT synchronized with provided SYNC_PULSE_IN.
//   - MultipurposeOutputFromPTP1588:
//     SYNC_PULSE_OUT synchronized with an external PTP IEEE 1588 master.
//   - MultipurposeOutputFromEncoderAngle:
//     SYNC_PULSE_OUT at a user-defined rate in an integer number of degrees.
//
// Note:
//
//   - When the mode is MultipurposeOutputFromInternalOscillator, MultipurposeOutputFromSyncPulseIn,
//     or MultipurposeOutputFromPTP1588, ConfigurationParameters.SyncPulseOutFrequency can define the output rate
//     (defaults to 1 Hz).
//   - When the sensor is in MultipurposeOutputFromEncoderAngle mode, ConfigurationParameters.SyncPulseOutAngle
//     defines the output pulse rate.
//   - In all modes, ConfigurationParameters.SyncPulseOutPulseWidth defines the output pulse width.
//
// For additional information, refer to [multipurpose_io_mode] and [External Trigger Clock Source].
//
// [External Trigger Clock Source]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/time_sync/time-sync.html#external-trigger-clock-source
// [multipurpose_io_mode]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#multipurpose-io-mode
type MultipurposeIOMode uint8

const (
	MultipurposeOff MultipurposeIOMode = iota // Do not output a SYNC_PULSE_OUT signal.

	// MultipurposeInputNMEAUART reconfigures the MULTIPURPOSE_IO port as an input.
	// See [Setting Ouster Sensors Time Source] for more information.
	//
	// [Setting Ouster Sensors Time Source]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/time_sync/time-sync.html#setting-sensor-time
	MultipurposeInputNMEAUART

	MultipurposeOutputFromInternalOscillator // Output a SYNC_PULSE_OUT signal synchronized with the internal clock.
	MultipurposeOutputFromSyncPulseIn        // Output a SYNC_PULSE_OUT signal synchronized with a SYNC_PULSE_IN provided to the unit.
	MultipurposeOutputFromPTP1588            // Output a SYNC_PULSE_OUT signal synchronized with an external PTP IEEE 1588 master.
	MultipurposeOutputFromEncoderAngle       // Output a SYNC_PULSE_OUT signal with a user defined rate in an integer number of degrees.
)

var (
	// multipurposeIOModeStringKV maps MultipurposeIOMode values to their string representations.
	multipurposeIOModeStringKV = map[MultipurposeIOMode]string{
		MultipurposeOff:                          "off",
		MultipurposeInputNMEAUART:                "input from NMEA UART",
		MultipurposeOutputFromInternalOscillator: "output from Internal Oscillator",
		MultipurposeOutputFromSyncPulseIn:        "output from Sync Pulse in",
		MultipurposeOutputFromPTP1588:            "output from PTP 1588",
		MultipurposeOutputFromEncoderAngle:       "output from encoder angle",
	}

	// multipurposeIOModeGoStringKV maps MultipurposeIOMode values to their Go syntax representations.
	multipurposeIOModeGoStringKV = map[MultipurposeIOMode]string{
		MultipurposeOff:                          "MultipurposeOff",
		MultipurposeInputNMEAUART:                "MultipurposeInputNMEAUART",
		MultipurposeOutputFromInternalOscillator: "MultipurposeOutputFromInternalOscillator",
		MultipurposeOutputFromSyncPulseIn:        "MultipurposeOutputFromSyncPulseIn",
		MultipurposeOutputFromPTP1588:            "MultipurposeOutputFromPTP1588",
		MultipurposeOutputFromEncoderAngle:       "MultipurposeOutputFromEncoderAngle",
	}

	// multipurposeIOModeTextKV maps MultipurposeIOMode values to their text representations.
	multipurposeIOModeTextKV = map[MultipurposeIOMode]string{
		MultipurposeOff:                          "OFF",
		MultipurposeInputNMEAUART:                "INPUT_NMEA_UART",
		MultipurposeOutputFromInternalOscillator: "OUTPUT_FROM_INTERNAL_OSC",
		MultipurposeOutputFromSyncPulseIn:        "OUTPUT_FROM_SYNC_PULSE_IN",
		MultipurposeOutputFromPTP1588:            "OUTPUT_FROM_PTP_1588",
		MultipurposeOutputFromEncoderAngle:       "OUTPUT_FROM_ENCODER_ANGLE",
	}

	// multipurposeIOModeTextVK maps string representations to MultipurposeIOMode values.
	multipurposeIOModeTextVK = util.ReverseMap(multipurposeIOModeTextKV)
)

// String returns the string representation of a MultipurposeIOMode value.
func (m MultipurposeIOMode) String() string {
	if s, ok := multipurposeIOModeStringKV[m]; ok {
		return s
	}

	return multipurposeIOModeStringKV[MultipurposeOff]
}

// GoString returns the Go syntax representation of a MultipurposeIOMode value.
func (m MultipurposeIOMode) GoString() string {
	if s, ok := multipurposeIOModeGoStringKV[m]; ok {
		return s
	}

	return multipurposeIOModeGoStringKV[MultipurposeOff]
}

// MarshalText returns the text representation of a MultipurposeIOMode value.
func (m MultipurposeIOMode) MarshalText() ([]byte, error) {
	if text, ok := multipurposeIOModeTextKV[m]; ok {
		return []byte(text), nil
	}

	return []byte(multipurposeIOModeTextKV[MultipurposeOff]), nil
}

// UnmarshalText updates the value of the MultipurposeIOMode receiver based on the input text.
//   - If the input text matches a valid mapping in the multipurposeIOModeTextVK map,
//     the receiver is updated to the corresponding MultipurposeIOMode value.
//   - If the input text does not match any valid mapping, the receiver is set to MultipurposeOff.
//
// This method does not return an error.
func (m *MultipurposeIOMode) UnmarshalText(text []byte) error {
	if mode, ok := multipurposeIOModeTextVK[string(text)]; ok {
		*m = mode
	} else {
		*m = MultipurposeOff
	}

	return nil
}
