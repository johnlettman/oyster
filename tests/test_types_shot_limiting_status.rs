use oyster::types::ShotLimitingStatus;

#[test]
#[allow(deprecated)]
fn test_types_shot_limiting_status_serde() {
    use serde_json::json;
    let test_cases = vec![
        (ShotLimitingStatus::Normal, json!("SHOT_LIMITING_NORMAL")),
        (ShotLimitingStatus::Imminent, json!("SHOT_LIMITING_IMMINENT")),
        (ShotLimitingStatus::Reduction0to10, json!("SHOT_LIMITING_0_TO_10")),
        (ShotLimitingStatus::Reduction10to20, json!("SHOT_LIMITING_10_TO_20")),
        (ShotLimitingStatus::Reduction20to30, json!("SHOT_LIMITING_20_TO_30")),
        (ShotLimitingStatus::Reduction30to40, json!("SHOT_LIMITING_30_TO_40")),
        (ShotLimitingStatus::Reduction40to50, json!("SHOT_LIMITING_40_TO_50")),
        (ShotLimitingStatus::Reduction50to60, json!("SHOT_LIMITING_50_TO_60")),
        (ShotLimitingStatus::Reduction60to70, json!("SHOT_LIMITING_60_TO_70")),
        (ShotLimitingStatus::Reduction70to75, json!("SHOT_LIMITING_70_TO_75")),
    ];

    for (return_order, expected_json) in test_cases {
        let serialized = serde_json::to_value(&return_order).unwrap();
        assert_eq!(serialized, expected_json);
        let deserialized: ShotLimitingStatus = serde_json::from_value(expected_json).unwrap();
        assert_eq!(deserialized, return_order);
    }
}
