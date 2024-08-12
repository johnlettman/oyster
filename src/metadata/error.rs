use std::{error, fmt, io};

#[derive(Debug)]
pub enum Error {
    IoError(io::Error),
    ParseError(serde_json::Error),
}

impl fmt::Display for Error {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match *self {
            Error::IoError(ref err) => write!(f, "IO error: {}", err),
            Error::ParseError(ref err) => write!(f, "Parse error: {}", err),
        }
    }
}

impl error::Error for Error {
    fn cause(&self) -> Option<&dyn error::Error> {
        match *self {
            Error::IoError(ref err) => Some(err),
            Error::ParseError(ref err) => Some(err),
        }
    }
}

impl From<io::Error> for Error {
    fn from(err: io::Error) -> Error {
        Error::IoError(err)
    }
}

impl From<serde_json::Error> for Error {
    fn from(err: serde_json::Error) -> Error {
        Error::ParseError(err)
    }
}
