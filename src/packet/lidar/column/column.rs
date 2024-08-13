use crate::packet::lidar::column::{Block, Header};
use binrw::{BinRead, BinWrite};

#[derive(Debug, BinRead, BinWrite, Eq, PartialEq, Clone)]
pub struct Column {
    header: Header,
    block: Block,
}

impl Column {
    #[inline]
    pub fn is_valid(&self) -> bool {
        self.header.is_valid()
    }
}

impl Default for Column {
    fn default() -> Self {
        Self {
            header: Header::default(),
            block: Block::SingleReturn(SingleReturns {})
        }
    }
}
