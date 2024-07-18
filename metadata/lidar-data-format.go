package metadata

import (
	"fmt"
	"github.com/johnlettman/oyster/packet/profile"
	"github.com/johnlettman/oyster/types"
	"math"
	"strings"
)

type LIDARDataFormat struct {
	LIDARProfile profile.LIDARProfile `json:"udp_profile_lidar"`
	IMUProfile   profile.IMUProfile   `json:"udp_profile_imu"`

	ColumnWindow     types.ColumnWindow `json:"column_window"`
	ColumnsPerFrame  int                `json:"columns_per_frame"`
	ColumnsPerPacket int                `json:"columns_per_packet"`

	PixelShiftByRow []int `json:"pixel_shift_by_row"`
	PixelsPerColumn int   `json:"pixels_per_column"`
}

func (f LIDARDataFormat) String() string {
	sb := new(strings.Builder)
	sb.WriteString("LIDARDataFormat:\n")
	sb.WriteString(fmt.Sprintf("\tLIDARProfile:     %s\n", f.LIDARProfile.String()))
	sb.WriteString(fmt.Sprintf("\tIMUProfile:       %s\n", f.IMUProfile.String()))

	if !f.ColumnWindow.Zero() {
		sb.WriteString(fmt.Sprintf("\tColumnWindow:     %s\n", f.ColumnWindow.String()))
	}

	if f.ColumnsPerFrame != 0 {
		sb.WriteString(fmt.Sprintf("\tColumnsPerFrame:  %d\n", f.ColumnsPerFrame))
	}

	if f.ColumnsPerPacket != 0 {
		sb.WriteString(fmt.Sprintf("\tColumnsPerPacket: %d\n", f.ColumnsPerPacket))
	}

	if f.PixelShiftByRow == nil || len(f.PixelShiftByRow) != 0 {
		sb.WriteString(fmt.Sprintf("\tPixelShiftByRow:  %v\n", f.PixelShiftByRow))
	}

	if f.PixelsPerColumn != 0 {
		sb.WriteString(fmt.Sprintf("\tPixelsPerColumn:  %d\n", f.PixelsPerColumn))
	}

	return sb.String()
}

func (f LIDARDataFormat) GoString() string {
	sb := new(strings.Builder)
	sb.WriteString("LIDARDataFormat{\n")
	sb.WriteString(fmt.Sprintf("\tLIDARProfile:     %s\n", f.LIDARProfile.GoString()))
	sb.WriteString(fmt.Sprintf("\tIMUProfile:       %s\n", f.IMUProfile.GoString()))

	if !f.ColumnWindow.Zero() {
		sb.WriteString(fmt.Sprintf("\tColumnWindow:     %s\n", f.ColumnWindow.GoString()))
	}

	if f.ColumnsPerFrame != 0 {
		sb.WriteString(fmt.Sprintf("\tColumnsPerFrame:  %d\n", f.ColumnsPerFrame))
	}

	if f.ColumnsPerPacket != 0 {
		sb.WriteString(fmt.Sprintf("\tColumnsPerPacket: %d\n", f.ColumnsPerPacket))
	}

	if len(f.PixelShiftByRow) != 0 {
		sb.WriteString(fmt.Sprintf("\tPixelShiftByRow:  %v\n", f.PixelShiftByRow))
	}

	if f.PixelsPerColumn != 0 {
		sb.WriteString(fmt.Sprintf("\tPixelsPerColumn:  %d\n", f.PixelsPerColumn))
	}

	sb.WriteRune('}')
	return sb.String()
}

// MaxFrameID returns the maximum frame ID based on the LIDAR profile.
//   - If the LIDAR profile is set to profile.LIDARProfileLegacy,
//     it returns math.MaxUint32;
//   - otherwise, it returns math.MaxUint16.
func (f LIDARDataFormat) MaxFrameID() int {
	if f.LIDARProfile == profile.LIDARProfileLegacy {
		return math.MaxUint32
	} else {
		return math.MaxUint16
	}
}

// Size calculates the total size in bytes of the LIDARDataFormat, including header, columns, and footer.
func (f LIDARDataFormat) Size() int {
	return f.HeaderSize() + (f.ColumnsPerPacket * f.ColumnSize()) + f.FooterSize()
}

// HeaderSize returns the size of the header in bytes for the LIDARDataFormat.
//   - If the LIDAR profile is set to profile.LIDARProfileLegacy, it returns 0;
//   - otherwise, it returns 32.
func (f LIDARDataFormat) HeaderSize() int {
	if f.LIDARProfile == profile.LIDARProfileLegacy {
		return 0
	} else {
		return 32
	}
}

// FooterSize returns the size of the footer in bytes for the LIDARDataFormat.
//   - If the LIDAR profile is set to profile.LIDARProfileLegacy, it returns 0;
//   - otherwise, it returns 32.
func (f LIDARDataFormat) FooterSize() int {
	if f.LIDARProfile == profile.LIDARProfileLegacy {
		return 0
	} else {
		return 32
	}
}

// FooterOffset calculates the offset in bytes for the footer in the LIDARDataFormat.
//   - If the footer size is 0, it returns 0;
//   - otherwise, it returns the calculated offset to the start of the footer.
func (f LIDARDataFormat) FooterOffset() int {
	if f.FooterSize() == 0 {
		return 0
	} else {
		return f.HeaderSize() + (f.ColumnsPerPacket * f.ColumnSize())
	}
}

// ColumnSize calculates the total size in bytes of a single column in the LIDARDataFormat.
// It includes the size of the column header, the size of the pixel data for the column,
// and the size of the column footer.
func (f LIDARDataFormat) ColumnSize() int {
	return f.ColumnHeaderSize() +
		(f.PixelsPerColumn * f.LIDARProfile.ColumnDataSize()) +
		f.ColumnFooterSize()
}

// ColumnOffset calculates the offset in bytes for a given column index in the LIDARDataFormat.
func (f LIDARDataFormat) ColumnOffset(column int) int {
	return f.HeaderSize() + (column * f.ColumnSize())
}

// ColumnHeaderSize returns the size of the header in bytes for the LIDARDataFormat's columns.
//   - If the LIDAR profile is set to profile.LIDARProfileLegacy, it returns 16;
//   - otherwise, it returns 12.
func (f LIDARDataFormat) ColumnHeaderSize() int {
	if f.LIDARProfile == profile.LIDARProfileLegacy {
		return 16
	} else {
		return 12
	}
}

// ColumnFooterSize returns the size of the footer in bytes for the LIDARDataFormat's columns.
//   - If the LIDAR profile is set to profile.LIDARProfileLegacy, it returns 4;
//   - otherwise, it returns 0.
func (f LIDARDataFormat) ColumnFooterSize() int {
	if f.LIDARProfile == profile.LIDARProfileLegacy {
		return 4
	} else {
		return 0
	}
}

// ColumnStatusOffset returns the offset in bytes for the status field in a single column of the LIDARDataFormat.
//   - If the LIDAR profile is set to profile.LIDARProfileLegacy,
//     it returns ColumnSize minus ColumnFooterSize;
//   - otherwise, it returns 10.
func (f LIDARDataFormat) ColumnStatusOffset() int {
	if f.LIDARProfile == profile.LIDARProfileLegacy {
		return f.ColumnSize() - f.ColumnFooterSize()
	} else {
		return 10
	}
}
