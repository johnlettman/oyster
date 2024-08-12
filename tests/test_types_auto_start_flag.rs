use env_logger::WriteStyle::Auto;
use fake::{Fake, Faker};
use oyster::types::AutoStartFlag;
use serde_json::json;

#[test]
fn test_serde() {
    use serde_json::json;
    let test_cases = vec![(AutoStartFlag::On, json!(1)), (AutoStartFlag::Off, json!(0))];

    for (flag, expected_json) in test_cases {
        let serialized = serde_json::to_value(&flag).unwrap();
        assert_eq!(serialized, expected_json);
        let deserialized: AutoStartFlag = serde_json::from_value(expected_json).unwrap();
        assert_eq!(deserialized, flag);
    }
}

#[test]
fn test_serde_random_on() {
    let random_u8 = json!(Faker.fake::<u8>() + 1);
    let deserialized: AutoStartFlag = serde_json::from_value(random_u8).unwrap();
    assert_eq!(deserialized, AutoStartFlag::On)
}
