use crate::packet::{ImuProfile, LidarProfile};
use crate::types::*;
use serde_with::{serde_as, BoolFromInt};
use std::net::Ipv4Addr;

/// Represents the configuration parameters for a sensor.
///
/// Ouster documentation: [Sensor Configuration].
///
/// [Sensor Configuration]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sensor-configuration
#[serde_as]
#[derive(Debug, serde::Serialize, serde::Deserialize, PartialEq, Clone)]
pub struct ConfigParams {
    /// The IP address destination to which the sensor sends UDP traffic.
    ///
    /// Ouster documentation: [`udp_dest`].
    ///
    /// [`udp_dest`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#udp-dest
    #[serde(rename = "udp_dest", alias = "udp_ip", default = "ConfigParams::default_udp_ip_dest")]
    udp_ip_dest: Ipv4Addr,

    /// The UDP port of the LIDAR component of the sensor.
    ///
    /// Ouster documentation: [`udp_port_lidar`].
    ///
    /// [`udp_port_lidar`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html?highlight=nmea_in_polarity#udp-port-lidar
    #[serde(default = "ConfigParams::default_udp_port_lidar")]
    udp_port_lidar: u16,

    /// The UDP port of the IMU component of the sensor.
    ///
    /// Ouster documentation: [`udp_port_imu`].
    ///
    /// [`udp_port_imu`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html?highlight=nmea_in_polarity#udp-port-imu
    #[serde(default = "ConfigParams::default_udp_port_imu")]
    udp_port_imu: u16,

    /// The polarity of the `SYNC_PULSE_IN` input, which controls polarity of the `SYNC_PULSE_IN`
    /// pin when timestmap_mopde is set to [TimestampMode::TimeFromSyncPulseIn].
    /// The default is [Polarity::ActiveHigh].
    ///
    /// Ouster documentation: [`sync_pulse_in_polarity`].
    ///
    /// [Polarity::ActiveHigh]: crate::types::Polarity::ActiveHigh
    /// [TimestampMode::TimeFromSyncPulseIn]: crate::types::TimestampMode::TimeFromSyncPulseIn
    /// [`sync_pulse_in_polarity`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sync-pulse-in-polarity
    #[serde(default = "ConfigParams::default_sync_pulse_in_polarity")]
    sync_pulse_in_polarity: Polarity,

    /// The polarity of the `SYNC_PULSE_OUT` output, if the sensor is set as the leader sensor
    /// used for time synchronization.
    /// The default is [Polarity::ActiveHigh].
    ///
    /// Ouster documentation: [`sync_pulse_out_polarity`].
    ///
    /// [`sync_pulse_out_polarity`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sync-pulse-out-polarity
    #[serde(default = "ConfigParams::default_sync_pulse_out_polarity")]
    sync_pulse_out_polarity: Polarity,

    /// The `SYNC_PULSE_OUT` pulse rate in Hz. Valid values are integers >0 Hz, but also limited
    /// by the criteria described in [Time Synchronization].
    /// The default is 1.
    ///
    /// Ouster documentation: [`sync_pulse_out_frequency`].
    ///
    /// [`sync_pulse_out_frequency`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sync-pulse-out-frequency
    /// [Time Synchronization]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/time_sync/time-sync.html#time-synchronization
    #[serde(default = "ConfigParams::default_sync_pulse_out_frequency")]
    sync_pulse_out_frequency: u32,

    /// The angle in degrees that the sensor traverses between each `SYNC_PULSE_OUT` pulse.
    /// For example, a value of 180 means a sync pulse is sent out every 180° for a total of
    /// two pulses per revolution and an angular frequency of 20 Hz if the sensor is
    /// in [LidarMode::Scan1024x10].
    /// This a value ranging from 0 to 360°, but also limited by criteria in [Time Synchronization].
    /// The default is 360°.
    ///
    /// Ouster documentation: [`sync_pulse_out_angle`].
    ///
    /// [LidarMode::Scan1024x10]: crate::types::LidarMode::Scan1024x10
    /// [`sync_pulse_out_angle`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#sync-pulse-out-angle
    /// [Time Synchronization]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/time_sync/time-sync.html#time-synchronization
    #[serde(default = "ConfigParams::default_sync_pulse_out_angle")]
    sync_pulse_out_angle: u16,

    /// The polarity of `SYNC_PULSE_OUT` output, if the sensor is set as the
    /// master sensor used for time synchronization.
    /// Output `SYNC_PULSE_OUT` pulse width is in milliseconds, increments in 1 ms.
    /// Valid inputs are integers greater than 0 ms, but also limited by the criteria
    /// described in [Time Synchronization].
    /// The default is 10 ms.
    ///
    /// Ouster documentation: [`sync_pulse_out_pulse_width`].
    ///
    /// [`sync_pulse_out_pulse_width`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html?highlight=nmea_in_polarity#sync-pulse-out-pulse-width
    /// [Time Synchronization]: https://static.ouster.dev/sensor-docs/image_route1/image_route3/time_sync/time-sync.html#time-synchronization
    #[serde(default = "ConfigParams::default_sync_pulse_out_pulse_width")]
    sync_pulse_out_pulse_width: u32,

    /// The polarity of NMEA UART input `$GPRMC` messages.
    ///
    /// Ouster documentation: [`nmea_in_polarity`].
    ///
    /// [`nmea_in_polarity`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#nmea-in-polarity
    #[serde(default = "ConfigParams::default_nmea_in_polarity")]
    nmea_in_polarity: Polarity,

    /// The configuration parameter that specifies whether to ignore `$GPRMC` messages if
    /// the valid character is not set.
    ///   - Set false to ignore the messages, and
    ///   - true to use them for time synchronization regardless of the valid character.
    ///
    /// Ouster documentation: [`nmea_ignore_valid_char`].
    ///
    /// [`nmea_ignore_valid_char`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#nmea-ignore-valid-char
    #[serde(default = "ConfigParams::default_nmea_ignore_valid_char")]
    #[serde_as(as = "BoolFromInt<serde_with::formats::Flexible>")]
    nmea_ignore_valid_char: bool,

    /// The expected baud rate the sensor is attempting to decode for
    /// NMEA UART input `$GPRMC` messages.
    ///
    /// Ouster documentation: [`nmea_baud_rate`].
    ///
    /// [`nmea_baud_rate`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#nmea-baud-rate
    #[serde(default = "ConfigParams::default_nmea_baud_rate")]
    nmea_baud_rate: NmeaBaudRate,

    /// Leap seconds that will be added to the UDP timestamp.
    ///
    /// Ouster documentation: [`nmea_leap_seconds`].
    ///
    /// [`nmea_leap_seconds`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#nmea-leap-seconds
    #[serde(default = "ConfigParams::default_nmea_leap_seconds")]
    nmea_leap_seconds: u32,

    /// The visible region of interest of the sensor in millidegrees.
    /// Only the data from within the specified bounds is sent.
    ///
    /// Ouster documentation: [`azimuth_window`].
    ///
    /// [`azimuth_window`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#azimuth-window
    azimuth_window: AzimuthWindow,

    /// The configured value of signal multiplier.
    /// This is of [f64] type value ranging between 0.25x and 3x.
    ///
    /// For 2x and 3x multipliers, the [azimuth_window] parameter sets the azimuth window that
    /// the lasers will be enabled in.
    /// Higher signal multiplier values infer smaller maximum [azimuth_window].
    ///
    /// The [signal_multiplier] value maximum [azimuth_window]:
    ///   - for 0.25x, 0.5x and 1x is 360° (default),
    ///   - for 2x: it's 180°, and
    ///   - for 3x: it's 120°.
    ///
    /// Ouster documentation: [`signal_multiplier`].
    ///
    /// # Note
    /// All sensors have equivalent power draw and thermal output when operating at the maximum
    /// [azimuth_window] for a particular signal multiplier value.
    /// Therefore, using an [azimuth_window] that is smaller than the maximum allowable
    /// [azimuth_window] with a particular signal multiplier value (excluding 1x) can reduce
    /// the power draw and thermal output of the sensor.
    ///
    /// However, while this can increase the max operating temperature of the sensor,
    /// it can also degrade the performance at low temperatures.
    /// This discrepancy will be resolved in a future firmware.
    ///
    /// [signal_multiplier]: ConfigParams::signal_multiplier
    /// [azimuth_window]: ConfigParams::azimuth_window
    /// [`signal_multiplier`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#signal-multiplier
    #[serde(default = "ConfigParams::default_signal_multiplier")]
    signal_multiplier: f64,

    /// The configuration of LIDAR packets.
    ///
    /// For additional information, [`udp_profile_lidar`].
    ///
    /// [`udp_profile_lidar`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#udp-profile-lidar
    #[serde(rename = "udp_profile_lidar", default = "ConfigParams::default_lidar_profile")]
    lidar_profile: LidarProfile,

    /// The configuration of IMU packets.
    ///
    /// For additional information, [`udp_profile_imu`].
    ///
    /// [`udp_profile_imu`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#udp-profile-imu
    #[serde(rename = "udp_profile_imu", default = "ConfigParams::default_imu_profile")]
    imu_profile: ImuProfile,

    /// Determines whether phase locking is enabled.
    ///
    /// Ouster documentation: [`phase_lock_enable`].
    ///
    /// [`phase_lock_enable`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#phase-lock-enable
    #[serde(default = "ConfigParams::default_phase_lock_enable")]
    phase_lock_enable: bool,

    /// The angle in the LIDAR Coordinate Frame that sensors are locked to in millidegrees if
    /// PhaseLockEnable is enabled. It ranges from 0...360000 and defaults to 0.
    /// Angle is traversed at the top of the second.
    ///
    /// For more information, refer to [`phase_lock_offset`].
    ///
    /// [`phase_lock_offset`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#phase-lock-offset
    #[serde(default = "ConfigParams::default_phase_lock_offset")]
    phase_lock_offset: u32,

    /// The horizontal resolution and rotation rate of the sensor.
    /// The effective range of the sensor is increased by 15-20% for every halving of the number of
    /// points gathered.
    /// For example, [LidarMode::Scan512x10] has a 15-20% longer range than [LidarMode::Scan512x20].
    ///
    /// Ouster documentation: [`lidar_mode`].
    ///
    /// [`lidar_mode`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html?highlight=512x10#lidar-mode
    lidar_mode: LidarMode,

    /// The methnod used for timestamp measurements.
    ///
    /// Ouster documentation: [`timestamp_mode`].
    ///
    /// [`timestamp_mode`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#timestamp-mode
    #[serde(skip_serializing_if = "Option::is_none")]
    timestamp_mode: Option<TimestampMode>,

    /// The mode of the `MULTIPURPOSE_IO` pin.
    ///
    /// Ouster documentation: [`multipurpose_io_mode`].
    ///
    /// [`multipurpose_io_mode`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#multipurpose-io-mode
    multipurpose_io_mode: MultipurposeIoMode,

    /// The power consumption and activity level of the Ouster sensor.
    ///
    /// Ouster documentation: [`operating_mode`].
    ///
    /// [`operating_mode`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#operating-mode
    operating_mode: OperatingMode,

    #[serde(default = "ConfigParams::default_auto_start_flag")]
    #[serde_as(as = "BoolFromInt<serde_with::formats::Flexible>")]
    auto_start_flag: bool,

    /// The minimum detection range of the LIDAR in centimeters.
    ///
    /// Ouster documentation: [`min_range_threshold_cm`].
    ///
    /// [`min_range_threshold_cm`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#min-range-threshold-cm
    #[serde(
        rename = "min_range_threshold_cm",
        default = "ConfigParams::default_min_range_threshold"
    )]
    min_range_threshold: usize,

    /// The priority of returns for the LIDAR to output.
    ///
    /// Ouster documentation: [`return_order`].
    ///
    /// [`return_order`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html?highlight=return_order#return-order
    return_order: ReturnOrder,

    /// The full-scale range for the gyroscope.
    /// It indicates whether modification of the onboard gyroscope has been enabled with an
    /// extended programmable scale.
    ///
    /// Ouster documentation: [`gyro_fsr`].
    ///
    /// [`gyro_fsr`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#gyro-fsr
    gyro_fsr: FullScaleRange,

    /// The full-scale range for the accelerometer.
    /// It indicates whether modification of the onboard accelerometer has been enabled with an
    /// extended programmable scale.
    ///
    /// Ouster documentation: [`accel_fsr`].
    ///
    /// [`accel_fsr`]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#accel-fsr
    accel_fsr: FullScaleRange,
}

impl ConfigParams {
    pub const DEFAULT_UDP_IP_DEST: Ipv4Addr = Ipv4Addr::new(127, 0, 0, 1);
    pub const DEFAULT_UDP_PORT_LIDAR: u16 = 7502;
    pub const DEFAULT_UDP_PORT_IMU: u16 = 7503;

    pub const DEFAULT_SYNC_PULSE_IN_POLARITY: Polarity = Polarity::ActiveHigh;
    pub const DEFAULT_SYNC_PULSE_OUT_POLARITY: Polarity = Polarity::ActiveHigh;
    pub const DEFAULT_SYNC_PULSE_OUT_FREQUENCY: u32 = 1;
    pub const DEFAULT_SYNC_PULSE_OUT_ANGLE: u16 = 360;
    pub const DEFAULT_SYNC_PULSE_OUT_WIDTH: u32 = 10;

    pub const DEFAULT_NMEA_IN_POLARITY: Polarity = Polarity::ActiveHigh;
    pub const DEFAULT_NMEA_IGNORE_VALID_CHAR: bool = false;
    pub const DEFAULT_NMEA_BAUD_RATE: NmeaBaudRate = NmeaBaudRate::Baud9600;
    pub const DEFAULT_NMEA_LEAP_SECONDS: u32 = 0;

    pub const DEFAULT_SIGNAL_MULTIPLIER: f64 = 1.0;

    pub const DEFAULT_PHASE_LOCK_ENABLE: bool = false;
    pub const DEFAULT_PHASE_LOCK_OFFSET: u32 = 0;

    pub const DEFAULT_LIDAR_PROFILE: LidarProfile = LidarProfile::DEFAULT;
    pub const DEFAULT_IMU_PROFILE: ImuProfile = ImuProfile::DEFAULT;

    pub const DEFAULT_MIN_RANGE_THRESHOLD: usize = 50;

    fn default_udp_ip_dest() -> Ipv4Addr {
        Self::DEFAULT_UDP_IP_DEST
    }
    fn default_udp_port_lidar() -> u16 {
        Self::DEFAULT_UDP_PORT_LIDAR
    }
    fn default_udp_port_imu() -> u16 {
        Self::DEFAULT_UDP_PORT_IMU
    }

    fn default_sync_pulse_in_polarity() -> Polarity {
        Self::DEFAULT_SYNC_PULSE_IN_POLARITY
    }
    fn default_sync_pulse_out_polarity() -> Polarity {
        Self::DEFAULT_SYNC_PULSE_OUT_POLARITY
    }
    fn default_sync_pulse_out_frequency() -> u32 {
        Self::DEFAULT_SYNC_PULSE_OUT_FREQUENCY
    }
    fn default_sync_pulse_out_angle() -> u16 {
        Self::DEFAULT_SYNC_PULSE_OUT_ANGLE
    }
    fn default_sync_pulse_out_pulse_width() -> u32 {
        Self::DEFAULT_SYNC_PULSE_OUT_WIDTH
    }

    fn default_nmea_in_polarity() -> Polarity {
        Self::DEFAULT_NMEA_IN_POLARITY
    }
    fn default_nmea_ignore_valid_char() -> bool {
        Self::DEFAULT_NMEA_IGNORE_VALID_CHAR
    }
    fn default_nmea_baud_rate() -> NmeaBaudRate {
        Self::DEFAULT_NMEA_BAUD_RATE
    }
    fn default_nmea_leap_seconds() -> u32 {
        Self::DEFAULT_NMEA_LEAP_SECONDS
    }

    fn default_signal_multiplier() -> f64 {
        Self::DEFAULT_SIGNAL_MULTIPLIER
    }

    fn default_phase_lock_enable() -> bool {
        Self::DEFAULT_PHASE_LOCK_ENABLE
    }
    fn default_phase_lock_offset() -> u32 {
        Self::DEFAULT_PHASE_LOCK_OFFSET
    }

    fn default_lidar_profile() -> LidarProfile {
        Self::DEFAULT_LIDAR_PROFILE
    }
    fn default_imu_profile() -> ImuProfile {
        Self::DEFAULT_IMU_PROFILE
    }
    fn default_auto_start_flag() -> bool {
        true
    }

    fn default_min_range_threshold() -> usize {
        Self::DEFAULT_MIN_RANGE_THRESHOLD
    }
}
