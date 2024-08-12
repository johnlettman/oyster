use std::mem::MaybeUninit;

pub trait ByteArray: AsRef<[u8]> + AsMut<[u8]> {
    const LEN: usize;
    fn uninit() -> Self;
    fn zeroed() -> Self;
}

impl<const N: usize> ByteArray for [u8; N] {
    const LEN: usize = N;
    fn uninit() -> Self {
        unsafe { MaybeUninit::uninit().assume_init() }
    }
    fn zeroed() -> Self {
        unsafe { MaybeUninit::zeroed().assume_init() }
    }
}

#[derive(Copy, Clone, Debug)]
pub enum PackingError {
    BadSize(usize, &'static str),
    BadAlignment(usize, &'static str),
    InvalidValue(&'static str),
}

pub type PackingResult<T> = Result<T, PackingError>;

pub trait Packer: Sized {
    type Packed: ByteArray;

    fn pack(&self, dest: &mut [u8]) -> PackingResult<()>;
    fn unpack(src: &[u8]) -> PackingResult<Self>;

    /// convenient getter for [ByteArray::LEN]
    fn packed_size() -> usize {
        Self::Packed::LEN
    }
    /// convenient getter for [ByteArray::LEN] * 8
    fn packed_bitsize() -> usize {
        Self::Packed::LEN * 8
    }

    /// like [Self::pack] but to a byte array instead of a slice
    fn packed(&self) -> PackingResult<Self::Packed> {
        let mut buffer = Self::Packed::uninit();
        self.pack(buffer.as_mut())?;
        Ok(buffer)
    }
}

#[macro_export]
macro_rules! bilge_packer {
    ($t: ty, $id: ident) => {
        impl $crate::packet::packing::Packer for $t {
            type Packed = [u8; ($id::BITS as usize + 7) / 8];

            fn pack(&self, dst: &mut [u8]) -> $crate::packet::packing::PackingResult<()> {
                use $crate::packet::packing::ByteArray;
                if dst.len() < Self::Packed::LEN {
                    return Err($crate::packet::packing::PackingError::BadSize(
                        dst.len(),
                        "bilge struct needs exact size",
                    ));
                }

                let common = Self::Packed::LEN.min(core::mem::size_of::<Self>());
                let src = unsafe { core::mem::transmute::<&Self, &Self::Packed>(self) };
                dst[..common].copy_from_slice(&src[..common]);
                dst[common..Self::Packed::LEN].fill(0);
                Ok(())
            }

            fn unpack(src: &[u8]) -> $crate::packet::packing::PackingResult<Self> {
                use $crate::packet::packing::ByteArray;
                if src.len() < Self::Packed::LEN {
                    return Err($crate::packet::packing::PackingError::BadSize(
                        src.len(),
                        "bilge struct needs exact size",
                    ));
                }

                let mut tmp = [0; core::mem::size_of::<Self>()];
                let common = Self::Packed::LEN.min(core::mem::size_of::<Self>());
                tmp[..common].copy_from_slice(&src[..common]);
                Ok(unsafe { core::mem::transmute::<[u8; core::mem::size_of::<Self>()], Self>(tmp) })
            }
        }
    };
}
