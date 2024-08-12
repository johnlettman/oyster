#[derive(Debug, modular_bitfield::BitfieldSpecifier, Eq, PartialEq, Copy, Clone)]
#[bits(1)]
pub enum Status {
    Valid = 0x1,
    Invalid = 0x0,
}

impl Status {
    #[inline]
    pub const fn valid(&self) -> bool {
        *self == Self::Valid
    }
}

impl Default for Status {
    fn default() -> Self {
        Self::Valid
    }
}