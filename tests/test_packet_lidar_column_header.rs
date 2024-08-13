use bilge::arbitrary_int::u96;
use fake::{Fake, Faker};
use oyster::packet::lidar::column::Header;
use oyster::packet::packing::Packer;

#[test]
fn test_try_from_u96() {
    let timestamp = Faker.fake::<u64>();
    let measurement_id = Faker.fake::<u16>();
    let status = 0b1000_0000_0000_0000u16;
    let want = Header::new(true, measurement_id, timestamp);

    let input: u128 =
        (timestamp as u128) << (16 + 16) | (measurement_id as u128) << 16 | status as u128;

    let got = Header::try_from(u96::new(input));

    assert!(got.is_ok());
    assert_eq!(got.unwrap(), want);
}

#[test]
fn test_try_from_bytes() {
    let timestamp = Faker.fake::<u64>();
    let measurement_id = Faker.fake::<u16>();
    let status = 0b1000_0000_0000_0000u16;
    let want = Header::new(true, measurement_id, timestamp);

    let input: u128 =
        (timestamp as u128) << (16 + 16) | (measurement_id as u128) << 16 | status as u128;

    let got = Header::unpack(&input.to_le_bytes());
    assert!(got.is_ok());

    let got_unwrapped = got.unwrap();
    assert_eq!(got_unwrapped.status(), true);
    assert_eq!(got_unwrapped.measurement_id(), measurement_id);
    assert_eq!(got_unwrapped.timestamp(), timestamp);
}
