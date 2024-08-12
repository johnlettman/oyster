use std::time::{SystemTime, SystemTimeError, UNIX_EPOCH};

pub fn nanos_to_timeval(nanos: u64) -> libc::timeval {
    let seconds = nanos / 1_000_000_000;
    let microseconds = (nanos % 1_000_000_000) / 1_000;

    libc::timeval { tv_sec: seconds as libc::time_t, tv_usec: microseconds as libc::suseconds_t }
}

pub fn timeval_now() -> Result<libc::timeval, SystemTimeError> {
    let now = SystemTime::now();
    let since_epoch = now.duration_since(UNIX_EPOCH)?;
    Ok(nanos_to_timeval(since_epoch.as_nanos() as u64))
}
