use crate::types::OperatingMode;
use chrono::{DateTime, Utc};
use serde;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, Eq, PartialEq, Clone)]
pub struct SensorInfo {
    /// The current status of the sensor.
    status: Option<String>,

    /// The current startup initialization number of the sensor.
    #[serde(default = "SensorInfo::default_initialization_id")]
    initialization_id: u32,

    /// The market product line of the sensor.
    #[serde(rename = "prod_line")]
    product_line: Option<String>,

    /// The part number of the sensor.
    #[serde(rename = "prod_pn")]
    part_number: Option<String>,

    /// The serial number of the sensor.
    #[serde(rename = "prod_sn")]
    serial_number: Option<String>,

    /// The build date of the sensor.
    build_date: Option<DateTime<Utc>>,

    /// The build revision of the sensor.
    #[serde(rename = "build_rev")]
    build_revision: Option<String>,

    /// The revision of the firmware on the sensor.
    #[serde(rename = "image_rev")]
    image_revision: Option<String>,

    /// The revision of the communications protocol.
    #[serde(rename = "proto_rev")]
    protocol_revision: Option<String>,
}

impl SensorInfo {
    pub const DEFAULT: Self =
        Self::new(None, Self::DEFAULT_INITIALIZATION_ID, None, None, None, None, None, None, None);

    pub const DEFAULT_INITIALIZATION_ID: u32 = 0;

    pub const fn new(
        status: Option<String>,
        initialization_id: u32,
        product_line: Option<String>,
        part_number: Option<String>,
        serial_number: Option<String>,
        build_date: Option<DateTime<Utc>>,
        build_revision: Option<String>,
        image_revision: Option<String>,
        protocol_revision: Option<String>,
    ) -> Self {
        Self {
            status,
            initialization_id,
            product_line,
            part_number,
            serial_number,
            build_date,
            build_revision,
            image_revision,
            protocol_revision,
        }
    }

    const fn default_initialization_id() -> u32 {
        Self::DEFAULT_INITIALIZATION_ID
    }
}

impl Default for SensorInfo {
    fn default() -> Self {
        Self::DEFAULT
    }
}
