package metadata

import (
	"github.com/johnlettman/oyster/packet/profile"
	"github.com/johnlettman/oyster/types"
	"github.com/johnlettman/oyster/types/pseudo"
	"net"
)

// ConfigParams represents the configuration parameters for a sensor.
// For additional information, refer to [Sensor Configuration].
//
// [Sensor Configuration]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sensor-configuration
type ConfigParams struct {
	AutoStartFlag    types.AutoStartFlag `json:"auto_start_flag,omitempty"`
	AzimuthWindow    types.AzimuthWindow `json:"azimuth_window,omitempty"`
	ColumnsPerPacket int                 `json:"columns_per_packet,omitempty"`

	LIDARMode     types.LIDARMode          `json:"lidar_mode"`
	OperatingMode types.OperatingMode      `json:"operating_mode"`
	IOMode        types.MultipurposeIOMode `json:"multipurpose_io_mode"`

	NMEABaudRate    types.NMEABaudRate `json:"nmea_baud_rate"`    // The expected baud rate the sensor is attempting to decode for NMEA UART input $GPRMC messages.
	NMEAInPolarity  types.Polarity     `json:"nmea_in_polarity"`  // The polarity of NMEA UART input $GPRMC messages.
	NMEALeapSeconds int                `json:"nmea_leap_seconds"` // Leap seconds that will be added to the UDP timestamp.

	// NMEAIgnoreValidChar represents a configuration parameter that specifies whether to ignore
	// $GPRMC messages if the valid character is not set.
	//   - Set false to ignore the messages, and
	//   - true to use them for time synchronization regardless of the valid character.
	// For additional information, refer to [nmea_ignore_valid_char].
	//
	// [nmea_ignore_valid_char]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#nmea-ignore-valid-char
	NMEAIgnoreValidChar pseudo.IntBoolJSON `json:"nmea_ignore_valid_char"`

	// PhaseLockEnable determines whether phase locking is enabled.
	// For additional information, refer to [phase_lock_enable].
	//
	// [phase_lock_enable]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#phase-lock-enable
	PhaseLockEnable bool `json:"phase_lock_enable"`

	// PhaseLockOffset is the angle in the LiDAR Coordinate Frame that sensors are locked to in
	// millidegrees if PhaseLockEnable is enabled. It ranges from 0...360000 and defaults to 0.
	// Angle is traversed at the top of the second.
	// For more information, refer to [phase_lock_offset].
	//
	// [phase_lock_offset]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#phase-lock-offset
	PhaseLockOffset int `json:"phase_lock_offset"`

	// SignalMultiplier is the configured value of signal multiplier.
	// This is of float64 type value ranging between 0.25x and 3x.
	//
	// For 2x and 3x multipliers, the AzimuthWindow parameter sets the azimuth window that the lasers will be enabled in.
	// Higher signal multiplier values infer smaller maximum AzimuthWindow.
	//
	// The SignalMultiplier value maximum AzimuthWindow:
	//   - for 0.25x, 0.5x and 1x is 360° (default),
	//   - for 2x: it's 180°, and
	//   - for 3x: it's 120°.
	//
	// For additional information, refer to [signal_multiplier].
	//
	// Note:
	//
	// All sensors have equivalent power draw and thermal output when operating at the
	// maximum AzimuthWindow for a particular signal multiplier value.
	// Therefore, using an AzimuthWindow that is smaller than the maximum allowable AzimuthWindow
	// with a particular signal multiplier value (excluding 1x) can reduce the power draw
	// and thermal output of the sensor.
	//
	// However, while this can increase the max operating temperature of the sensor,
	// it can also degrade the performance at low temperatures.
	// This discrepancy will be resolved in a future firmware.
	//
	// [signal_multiplier]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#signal-multiplier
	SignalMultiplier float64 `json:"signal_multiplier"`

	// SyncPulseInPolarity is the polarity of the SYNC_PULSE_IN input,
	// which controls polarity of the SYNC_PULSE_IN pin when TimestampMode is set to
	// types.TimeFromSyncPulseIn.
	// The default is types.PolarityActiveHigh.
	// For additional information, refer to [sync_pulse_in_polarity].
	//
	// [sync_pulse_in_polarity]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sync-pulse-in-polarity
	SyncPulseInPolarity types.Polarity `json:"sync_pulse_in_polarity"`

	// SyncPulseOutPolarity is the polarity of the SYNC_PULSE_OUT output,
	// if the sensor is set as the leader sensor used for time synchronization.
	// The default is types.PolarityActiveHigh.
	// For additional information, refer to [sync_pulse_out_polarity].
	//
	// [sync_pulse_out_polarity]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sync-pulse-out-polarity
	SyncPulseOutPolarity types.Polarity `json:"sync_pulse_out_polarity"`

	// SyncPulseOutFrequency is the SYNC_PULSE_OUT pulse rate in Hz.
	// Valid values are integers >0 Hz, but also limited by the criteria described in
	// [Time Synchronization].
	// The default is 1.
	// For additional information, refer to [sync_pulse_out_frequency].
	//
	// [sync_pulse_out_frequency]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sync-pulse-out-frequency
	// [Time Synchronization]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/time_sync/time-sync.html#time-synchronization
	SyncPulseOutFrequency uint `json:"sync_pulse_out_frequency"`

	// SyncPulseOutAngle is the angle in degrees that the sensor traverses between
	// each SYNC_PULSE_OUT pulse. For example, a value of 180 means a sync pulse is
	// sent out every 180° for a total of two pulses per revolution and
	// an angular frequency of 20 Hz if the sensor is types.LidarMode1024x10.
	// This a value ranging from 0 to 360, but also limited by criteria in
	// [Time Synchronization]. The default is 360.
	// For additional information, refer to [sync_pulse_out_angle].
	//
	// [sync_pulse_out_angle]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sync-pulse-out-angle
	// [Time Synchronization]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/time_sync/time-sync.html#time-synchronization
	SyncPulseOutAngle uint `json:"sync_pulse_out_angle"`

	// SyncPulseOutPulseWidth is polarity of SYNC_PULSE_OUT output,
	// if the sensor is set as the leader sensor used for time synchronization.
	// The output SYNC_PULSE_OUT pulse width increments in 1 ms steps.
	// Limits and rules for valid inputs come from the criteria described in
	// [Time Synchronization].
	// The default is 10.
	// For additional information, refer to [sync_pulse_out_pulse_width].
	//
	// [sync_pulse_out_pulse_width]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sync-pulse-out-pulse-width
	// [Time Synchronization]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/time_sync/time-sync.html#time-synchronization
	SyncPulseOutPulseWidth uint `json:"sync_pulse_out_pulse_width"`

	// TimestampMode is the method used to timestamp measurements.
	// For additional information, refer to [timestamp_mode].
	//
	// [timestamp_mode]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#timestamp-mode
	TimestampMode types.TimeSource `json:"timestamp_mode"`

	IMUProfile       profile.IMUProfile   `json:"udp_profile_imu"`
	LIDARProfile     profile.LIDARProfile `json:"udp_profile_lidar"`
	UDPDestinationIP *net.IP              `json:"udp_dest_ip,omitempty"` // The sensor destination IP address over UDP.
	UDPSourceIP      *net.IP              `json:"udp_ip,omitempty"`      // The sensor source IP address over UDP.
	UDPIMUPort       uint16               `json:"udp_port_imu"`          // The UDP port for the IMU on the sensor.
	UDPLidarPort     uint16               `json:"udp_port_lidar"`        // The UDP port for the LiDAR on the sensor.

	// GyroFSR is the full-scale range for the gyroscope.
	// It indicates whether modification of the onboard gyroscope has been
	// enabled with an extended programmable scale.
	// For additional information, refer to [gyro_fsr].
	//
	// [gyro_fsr]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#gyro-fsr
	GyroFSR *types.FullScaleRange `json:"gyro_fsr,omitempty"`

	// AccelerometerFSR is the full-scale range of the accelerometer.
	// It indicates whether modification of the onboard accelerometer
	// has been enabled with an extended programmable scale.
	// For additional information, refer to [accel_fsr].
	//
	// [accel_fsr]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#accel-fsr
	AccelerometerFSR *types.FullScaleRange `json:"accel_fsr,omitempty"`
}
