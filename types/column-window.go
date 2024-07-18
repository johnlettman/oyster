package types

import (
	"fmt"
	"github.com/johnlettman/oyster/types/pseudo"
)

// ColumnWindow is the window over which the sensor fires in columns.
type ColumnWindow pseudo.ListTupleJSON[int]

// String returns a string representation of the ColumnWindow in the format "V1xV2".
// The value V1 is the first element in the ColumnWindow and V2 is the second element.
func (c ColumnWindow) String() string {
	return fmt.Sprintf("%dx%d", c.V1, c.V2)
}

// GoString returns a Go syntax representation of the ColumnWindow in the format ColumnWindow{V1, V2}.
// The value V1 is the first element in the ColumnWindow and V2 is the second element.
func (c ColumnWindow) GoString() string {
	return fmt.Sprintf("ColumnWindow{%d, %d}", c.V1, c.V2)
}

// MarshalJSON returns the JSON encoding of the ColumnWindow by marshaling
// it as a ListTupleJSON[int].
func (c ColumnWindow) MarshalJSON() ([]byte, error) {
	return (pseudo.ListTupleJSON[int])(c).MarshalJSON()
}

// UnmarshalJSON unmarshals the JSON-encoded data into the ColumnWindow by
// unmarshaling it as a ListTupleJSON[int].
func (c *ColumnWindow) UnmarshalJSON(data []byte) error {
	return (*pseudo.ListTupleJSON[int])(c).UnmarshalJSON(data)
}

// Zero returns true if the values of V1 and V2 in the ColumnWindow are both equal to 0.
func (c ColumnWindow) Zero() bool {
	return c.V1 == 0 && c.V2 == 0
}
