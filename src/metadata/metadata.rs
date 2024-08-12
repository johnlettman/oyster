use crate::metadata::{
    BeamIntrinsics, CalibrationStatus, ConfigParams, ImuIntrinsics, LidarDataFormat,
    LidarIntrinsics, SensorInfo,
};
use crate::util::string_or_struct;
use serde::{Deserialize, Serialize};
use std::error::Error;
use std::fs::File;
use std::io::BufReader;

#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct Metadata {
    beam_intrinsics: BeamIntrinsics,

    //#[serde(deserialize_with = "string_or_struct")]
    //calibration_status: CalibrationStatus,
    client_version: Option<String>,

    config_params: ConfigParams,
    imu_intrinsics: ImuIntrinsics,

    #[serde(deserialize_with = "string_or_struct")]
    lidar_data_format: LidarDataFormat,
    lidar_intrinsics: LidarIntrinsics,

    sensor_info: SensorInfo,
}

impl Metadata {
    pub fn from_file(path: &str) -> Result<Metadata, Box<dyn Error>> {
        let file = File::open(path)?;
        let reader = BufReader::new(file);
        let metadata = serde_json::from_reader(reader)?;
        Ok(metadata)
    }
}

#[test]
fn test() {
    //let path = "/home/jlettman/repos/oyster/samples/metadata/1_12_os1-991913000010-64.json";
    let path = "/home/jlettman/repos/oyster/samples/metadata/3_0_1_os-122246000293-128.json";
    let metadata = Metadata::from_file(path);

    match metadata {
        Err(e) => eprintln!("{}", e),
        Ok(m) => println!("{:#?}", m),
    }
}
