package types

import (
	"fmt"
	"github.com/johnlettman/oyster/types/pseudo"
	"github.com/johnlettman/oyster/util"
)

const (
	MaxMillidegrees int = 360000 // The maximum value of millidegrees (360000).
	MinMillidegrees     = 0      // The minimum value of millidegrees (0).
)

// AzimuthWindow sets the visible region of interest for the LIDAR sensor in millidegrees.
// The sensor only sends data from the specified azimuth window boundary.
// For additional information, refer to [azimuth_window].
//
// [azimuth_window]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html#azimuth-window
type AzimuthWindow pseudo.ListTupleJSON[int]

// DefaultAzimuthWindow returns a new AzimuthWindow with the minimum and maximum millidegrees values.
// The value of V1 is set to MinMillidegrees, and the value of V2 is set to MaxMillidegrees.
func DefaultAzimuthWindow() AzimuthWindow {
	return AzimuthWindow{V1: MinMillidegrees, V2: MaxMillidegrees}
}

// String returns a string representation of the AzimuthWindow in the format "V1xV2".
// The value V1 is the first element in the AzimuthWindow and V2 is the second element.
func (a AzimuthWindow) String() string {
	return fmt.Sprintf("%dm°x%dm°", a.V1, a.V2)
}

// GoString returns a Go syntax representation of the AzimuthWindow in the format AzimuthWindow{V1, V2}.
// The value V1 is the first element in the AzimuthWindow and V2 is the second element.
func (a AzimuthWindow) GoString() string {
	return fmt.Sprintf("AzimuthWindow{%d, %d}", a.V1, a.V2)
}

// MarshalJSON returns the JSON encoding of the AzimuthWindow by marshaling
// it as a ListTupleJSON[int].
func (a AzimuthWindow) MarshalJSON() ([]byte, error) {
	return (pseudo.ListTupleJSON[int])(a).MarshalJSON()
}

// UnmarshalJSON unmarshals the JSON-encoded data into the AzimuthWindow by
// unmarshaling it as a ListTupleJSON[int].
func (a *AzimuthWindow) UnmarshalJSON(data []byte) error {
	return (*pseudo.ListTupleJSON[int])(a).UnmarshalJSON(data)
}

// Valid checks if the values V1 and V2 of AzimuthWindow are within valid bounds.
func (a AzimuthWindow) Valid() bool {
	return a.V1 <= MaxMillidegrees && a.V2 <= MaxMillidegrees &&
		a.V1 >= MinMillidegrees && a.V2 >= MinMillidegrees
}

// Size returns the absolute difference between V2 and V1 in the AzimuthWindow.
// If V2 is greater than V1, it returns V2 - V1. Otherwise, it returns V1 - V2.
func (a AzimuthWindow) Size() int {
	return util.Abs(a.V2 - a.V1)
}

// Start returns the start of the window by finding the
// minimum value between V1 and V2 in the AzimuthWindow.
func (a AzimuthWindow) Start() int {
	return min(a.V1, a.V2)
}

// End returns the end of the window by finding the
// maximum value between V1 and V2 in the AzimuthWindow.
func (a AzimuthWindow) End() int {
	return max(a.V1, a.V2)
}
