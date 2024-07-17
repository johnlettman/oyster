package types

import "github.com/johnlettman/oyster/util"

// ReturnOrder represents the priority of returns for the LIDAR to output.
// The LIDAR can have more than 1 or 2 detected "returns."
// ReturnOrder indicates to the LiDAR which ones it should output.
//
// For additional information, refer to [return_order].
//
// [return_order]: https://static.ouster.dev/sensor-docs/image_route1/image_route2/common_sections/API/sensor_configuration_description.html?highlight=return_order#return-order)
type ReturnOrder int

const (
	ReturnOrderUnspecified        ReturnOrder = iota
	ReturnOrderStrongestToWeakest             // The priority of LiDAR returns with the strongest returns first.
	ReturnOrderFarthestToNearest              // The priority of LiDAR returns with the farthest returns first.
	ReturnOrderNearestToFarthest              // The priority of LiDAR returns with the nearest returns first.

	// Deprecated: Only present in old test firmware.
	ReturnOrderStrongestReturnFirst // The priority of LiDAR returns with the strongest returns first.

	// Deprecated: Only present in old test firmware.
	ReturnOrderLastReturnFirst // The priority of LiDAR returns with the last returns first.
)

var (
	// returnOrderStringKV maps ReturnOrder values to their string representations.
	returnOrderStringKV = map[ReturnOrder]string{
		ReturnOrderUnspecified:          "unspecified",
		ReturnOrderStrongestToWeakest:   "strongest to weakest",
		ReturnOrderNearestToFarthest:    "nearest to farthest",
		ReturnOrderFarthestToNearest:    "farthest to nearest",
		ReturnOrderStrongestReturnFirst: "strongest return first",
		ReturnOrderLastReturnFirst:      "last return first",
	}

	// returnOrderGoStringKV maps ReturnOrder values to their Go syntax representations.
	returnOrderGoStringKV = map[ReturnOrder]string{
		ReturnOrderUnspecified:          "ReturnOrderUnspecified",
		ReturnOrderStrongestToWeakest:   "ReturnOrderStrongestToWeakest",
		ReturnOrderNearestToFarthest:    "ReturnOrderNearestToFarthest",
		ReturnOrderFarthestToNearest:    "ReturnOrderFarthestToNearest",
		ReturnOrderStrongestReturnFirst: "ReturnOrderStrongestReturnFirst",
		ReturnOrderLastReturnFirst:      "ReturnOrderLastReturnFirst",
	}

	// returnOrderTextKV maps ReturnOrder values to their text representations.
	returnOrderTextKV = map[ReturnOrder]string{
		ReturnOrderUnspecified:          "UNSPECIFIED",
		ReturnOrderStrongestToWeakest:   "STRONGEST_TO_WEAKEST",
		ReturnOrderNearestToFarthest:    "NEAREST_TO_FARTHEST",
		ReturnOrderFarthestToNearest:    "FARTHEST_TO_NEAREST",
		ReturnOrderStrongestReturnFirst: "STRONGEST_RETURN_FIRST",
		ReturnOrderLastReturnFirst:      "LAST_RETURN_FIRST",
	}

	// returnOrderTextVK maps string representations to ReturnOrder values.
	returnOrderTextVK = util.ReverseMap(returnOrderTextKV)
)

// String returns the string representation of an ReturnOrder value.
func (r ReturnOrder) String() string {
	if s, ok := returnOrderStringKV[r]; ok {
		return s
	}

	return returnOrderStringKV[ReturnOrderUnspecified]
}

// GoString returns the Go syntax representation of an ReturnOrder value.
func (r ReturnOrder) GoString() string {
	if s, ok := returnOrderGoStringKV[r]; ok {
		return s
	}

	return returnOrderGoStringKV[ReturnOrderUnspecified]
}

// MarshalText returns the text representation of an ReturnOrder value.
func (r ReturnOrder) MarshalText() ([]byte, error) {
	if text, ok := returnOrderTextKV[r]; ok {
		return []byte(text), nil
	}

	return []byte(returnOrderTextKV[ReturnOrderUnspecified]), nil
}

// UnmarshalText updates the value of ReturnOrder by unmarshalling the given text.
// It checks if the text matches any string representation in returnOrderTextVK.
//   - If a match is found, the corresponding ReturnOrder value is assigned to r.
//   - Otherwise, ReturnOrderUnspecified is assigned to r.
//
// UnmarshalText always returns nil as error.
func (r *ReturnOrder) UnmarshalText(text []byte) error {
	if returnOrder, ok := returnOrderTextVK[string(text)]; ok {
		*r = returnOrder
	} else {
		*r = ReturnOrderUnspecified
	}

	return nil
}
