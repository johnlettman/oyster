use oyster::types::ReturnOrder;

#[test]
#[allow(deprecated)]
fn test_types_return_order_serde() {
    use serde_json::json;
    let test_cases = vec![
        (ReturnOrder::StrongestToWeakest, json!("STRONGEST_TO_WEAKEST")),
        (ReturnOrder::NearestToFarthest, json!("NEAREST_TO_FARTHEST")),
        (ReturnOrder::FarthestToNearest, json!("FARTHEST_TO_NEAREST")),
        (ReturnOrder::StrongestReturnFirst, json!("STRONGEST_RETURN_FIRST")),
        (ReturnOrder::LastReturnFirst, json!("LAST_RETURN_FIRST")),
    ];

    for (return_order, expected_json) in test_cases {
        let serialized = serde_json::to_value(&return_order).unwrap();
        assert_eq!(serialized, expected_json);
        let deserialized: ReturnOrder = serde_json::from_value(expected_json).unwrap();
        assert_eq!(deserialized, return_order);
    }
}
