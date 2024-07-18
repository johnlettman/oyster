package metadata

import (
	"fmt"
	"github.com/johnlettman/oyster/packet/profile"
	"github.com/johnlettman/oyster/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLIDARDataFormat_String(t *testing.T) {
	type TestCase struct {
		name       string
		dataFormat LIDARDataFormat
		want       string
	}

	columnWindow := types.ColumnWindow{0, 1024}
	lidarProfile := profile.LIDARProfileDualReturns
	imuProfile := profile.IMUProfileLegacy
	pixelShiftByRow := []int{1, 2, 3}

	cases := []TestCase{
		{
			name: "full",
			dataFormat: LIDARDataFormat{
				LIDARProfile:     lidarProfile,
				IMUProfile:       imuProfile,
				ColumnWindow:     columnWindow,
				ColumnsPerFrame:  32,
				ColumnsPerPacket: 32,
				PixelShiftByRow:  pixelShiftByRow,
				PixelsPerColumn:  32,
			},
			want: fmt.Sprintf("LIDARDataFormat:\n"+
				"\tLIDARProfile:     %s\n"+
				"\tIMUProfile:       %s\n"+
				"\tColumnWindow:     %s\n"+
				"\tColumnsPerFrame:  32\n"+
				"\tColumnsPerPacket: 32\n"+
				"\tPixelShiftByRow:  %v\n"+
				"\tPixelsPerColumn:  32\n",
				lidarProfile, imuProfile, columnWindow, pixelShiftByRow),
		},
		{
			name: "zero ColumnWindow",
			dataFormat: LIDARDataFormat{
				LIDARProfile:     lidarProfile,
				IMUProfile:       imuProfile,
				ColumnWindow:     types.ColumnWindow{},
				ColumnsPerFrame:  32,
				ColumnsPerPacket: 32,
				PixelShiftByRow:  pixelShiftByRow,
				PixelsPerColumn:  32,
			},
			want: fmt.Sprintf("LIDARDataFormat:\n"+
				"\tLIDARProfile:     %s\n"+
				"\tIMUProfile:       %s\n"+
				"\tColumnsPerFrame:  32\n"+
				"\tColumnsPerPacket: 32\n"+
				"\tPixelShiftByRow:  %v\n"+
				"\tPixelsPerColumn:  32\n",
				lidarProfile, imuProfile, pixelShiftByRow),
		},
		{
			name: "zero ColumnsPerFrame",
			dataFormat: LIDARDataFormat{
				LIDARProfile:     lidarProfile,
				IMUProfile:       imuProfile,
				ColumnWindow:     columnWindow,
				ColumnsPerFrame:  0,
				ColumnsPerPacket: 32,
				PixelShiftByRow:  pixelShiftByRow,
				PixelsPerColumn:  32,
			},
			want: fmt.Sprintf("LIDARDataFormat:\n"+
				"\tLIDARProfile:     %s\n"+
				"\tIMUProfile:       %s\n"+
				"\tColumnWindow:     %s\n"+
				"\tColumnsPerPacket: 32\n"+
				"\tPixelShiftByRow:  %v\n"+
				"\tPixelsPerColumn:  32\n",
				lidarProfile, imuProfile, columnWindow, pixelShiftByRow),
		},
		{
			name: "zero ColumnsPerPacket",
			dataFormat: LIDARDataFormat{
				LIDARProfile:     lidarProfile,
				IMUProfile:       imuProfile,
				ColumnWindow:     columnWindow,
				ColumnsPerFrame:  32,
				ColumnsPerPacket: 0,
				PixelShiftByRow:  pixelShiftByRow,
				PixelsPerColumn:  32,
			},
			want: fmt.Sprintf("LIDARDataFormat:\n"+
				"\tLIDARProfile:     %s\n"+
				"\tIMUProfile:       %s\n"+
				"\tColumnWindow:     %s\n"+
				"\tColumnsPerFrame:  32\n"+
				"\tPixelShiftByRow:  %v\n"+
				"\tPixelsPerColumn:  32\n",
				lidarProfile, imuProfile, columnWindow, pixelShiftByRow),
		},
		{
			name: "zero PixelShiftByRow",
			dataFormat: LIDARDataFormat{
				LIDARProfile:     lidarProfile,
				IMUProfile:       imuProfile,
				ColumnWindow:     columnWindow,
				ColumnsPerFrame:  32,
				ColumnsPerPacket: 32,
				PixelShiftByRow:  []int{},
				PixelsPerColumn:  32,
			},
			want: fmt.Sprintf("LIDARDataFormat:\n"+
				"\tLIDARProfile:     %s\n"+
				"\tIMUProfile:       %s\n"+
				"\tColumnWindow:     %s\n"+
				"\tColumnsPerFrame:  32\n"+
				"\tColumnsPerPacket: 32\n"+
				"\tPixelsPerColumn:  32\n",
				lidarProfile, imuProfile, columnWindow),
		},
		{
			name: "zero PixelsPerColumn",
			dataFormat: LIDARDataFormat{
				LIDARProfile:     lidarProfile,
				IMUProfile:       imuProfile,
				ColumnWindow:     columnWindow,
				ColumnsPerFrame:  32,
				ColumnsPerPacket: 32,
				PixelShiftByRow:  pixelShiftByRow,
				PixelsPerColumn:  0,
			},
			want: fmt.Sprintf("LIDARDataFormat:\n"+
				"\tLIDARProfile:     %s\n"+
				"\tIMUProfile:       %s\n"+
				"\tColumnWindow:     %s\n"+
				"\tColumnsPerFrame:  32\n"+
				"\tColumnsPerPacket: 32\n"+
				"\tPixelShiftByRow:  %v\n",
				lidarProfile, imuProfile, columnWindow, pixelShiftByRow),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.dataFormat.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestLIDARDataFormat_GoString(t *testing.T) {
	type TestCase struct {
		name       string
		dataFormat LIDARDataFormat
		want       string
	}

	columnWindow := types.ColumnWindow{0, 1024}
	lidarProfile := profile.LIDARProfileDualReturns
	imuProfile := profile.IMUProfileLegacy
	pixelShiftByRow := []int{1, 2, 3}

	cases := []TestCase{
		{
			name: "full",
			dataFormat: LIDARDataFormat{
				LIDARProfile:     lidarProfile,
				IMUProfile:       imuProfile,
				ColumnWindow:     columnWindow,
				ColumnsPerFrame:  32,
				ColumnsPerPacket: 32,
				PixelShiftByRow:  pixelShiftByRow,
				PixelsPerColumn:  32,
			},
			want: fmt.Sprintf("LIDARDataFormat{\n"+
				"\tLIDARProfile:     %#v\n"+
				"\tIMUProfile:       %#v\n"+
				"\tColumnWindow:     %#v\n"+
				"\tColumnsPerFrame:  32\n"+
				"\tColumnsPerPacket: 32\n"+
				"\tPixelShiftByRow:  %v\n"+
				"\tPixelsPerColumn:  32\n}",
				lidarProfile, imuProfile, columnWindow, pixelShiftByRow),
		},
		{
			name: "zero ColumnWindow",
			dataFormat: LIDARDataFormat{
				LIDARProfile:     lidarProfile,
				IMUProfile:       imuProfile,
				ColumnWindow:     types.ColumnWindow{},
				ColumnsPerFrame:  32,
				ColumnsPerPacket: 32,
				PixelShiftByRow:  pixelShiftByRow,
				PixelsPerColumn:  32,
			},
			want: fmt.Sprintf("LIDARDataFormat{\n"+
				"\tLIDARProfile:     %#v\n"+
				"\tIMUProfile:       %#v\n"+
				"\tColumnsPerFrame:  32\n"+
				"\tColumnsPerPacket: 32\n"+
				"\tPixelShiftByRow:  %v\n"+
				"\tPixelsPerColumn:  32\n}",
				lidarProfile, imuProfile, pixelShiftByRow),
		},
		{
			name: "zero ColumnsPerFrame",
			dataFormat: LIDARDataFormat{
				LIDARProfile:     lidarProfile,
				IMUProfile:       imuProfile,
				ColumnWindow:     columnWindow,
				ColumnsPerFrame:  0,
				ColumnsPerPacket: 32,
				PixelShiftByRow:  pixelShiftByRow,
				PixelsPerColumn:  32,
			},
			want: fmt.Sprintf("LIDARDataFormat{\n"+
				"\tLIDARProfile:     %#v\n"+
				"\tIMUProfile:       %#v\n"+
				"\tColumnWindow:     %#v\n"+
				"\tColumnsPerPacket: 32\n"+
				"\tPixelShiftByRow:  %v\n"+
				"\tPixelsPerColumn:  32\n}",
				lidarProfile, imuProfile, columnWindow, pixelShiftByRow),
		},
		{
			name: "zero ColumnsPerPacket",
			dataFormat: LIDARDataFormat{
				LIDARProfile:     lidarProfile,
				IMUProfile:       imuProfile,
				ColumnWindow:     columnWindow,
				ColumnsPerFrame:  32,
				ColumnsPerPacket: 0,
				PixelShiftByRow:  pixelShiftByRow,
				PixelsPerColumn:  32,
			},
			want: fmt.Sprintf("LIDARDataFormat{\n"+
				"\tLIDARProfile:     %#v\n"+
				"\tIMUProfile:       %#v\n"+
				"\tColumnWindow:     %#v\n"+
				"\tColumnsPerFrame:  32\n"+
				"\tPixelShiftByRow:  %v\n"+
				"\tPixelsPerColumn:  32\n}",
				lidarProfile, imuProfile, columnWindow, pixelShiftByRow),
		},
		{
			name: "zero PixelShiftByRow",
			dataFormat: LIDARDataFormat{
				LIDARProfile:     lidarProfile,
				IMUProfile:       imuProfile,
				ColumnWindow:     columnWindow,
				ColumnsPerFrame:  32,
				ColumnsPerPacket: 32,
				PixelShiftByRow:  []int{},
				PixelsPerColumn:  32,
			},
			want: fmt.Sprintf("LIDARDataFormat{\n"+
				"\tLIDARProfile:     %#v\n"+
				"\tIMUProfile:       %#v\n"+
				"\tColumnWindow:     %#v\n"+
				"\tColumnsPerFrame:  32\n"+
				"\tColumnsPerPacket: 32\n"+
				"\tPixelsPerColumn:  32\n}",
				lidarProfile, imuProfile, columnWindow),
		},
		{
			name: "zero PixelsPerColumn",
			dataFormat: LIDARDataFormat{
				LIDARProfile:     lidarProfile,
				IMUProfile:       imuProfile,
				ColumnWindow:     columnWindow,
				ColumnsPerFrame:  32,
				ColumnsPerPacket: 32,
				PixelShiftByRow:  pixelShiftByRow,
				PixelsPerColumn:  0,
			},
			want: fmt.Sprintf("LIDARDataFormat{\n"+
				"\tLIDARProfile:     %#v\n"+
				"\tIMUProfile:       %#v\n"+
				"\tColumnWindow:     %#v\n"+
				"\tColumnsPerFrame:  32\n"+
				"\tColumnsPerPacket: 32\n"+
				"\tPixelShiftByRow:  %v\n}",
				lidarProfile, imuProfile, columnWindow, pixelShiftByRow),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.dataFormat.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}
