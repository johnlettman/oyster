mod dual_returns;
mod single_returns_low_data_rate;
mod single_returns;
mod fusa_two_word_pixel;

pub use dual_returns::DualReturnBlock;
pub use single_returns_low_data_rate::SingleLowDataRateReturnBlock;
pub use single_returns::SingleReturnBlock;
pub use fusa_two_word_pixel::FuSaDualReturnBlock;

/// Column data segment from a LIDAR packet.
#[derive(Debug, Clone, Eq, PartialEq)]
pub enum Block {
    DualReturn(DualReturnBlock),
    SingleReturn(SingleReturnBlock),
    SingleLowDataRateReturn(SingleLowDataRateReturnBlock),
    FuSaDualReturn(FuSaDualReturnBlock),
}
