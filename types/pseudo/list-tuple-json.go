package pseudo

import (
	"encoding/json"
	"fmt"
	"github.com/barweiss/go-tuple"
	"reflect"
)

// ListTupleJSON is a type that represents a tuple with two values stored in a List structure for JSON.
type ListTupleJSON[T any] tuple.Pair[T, T]

// String returns a string representation of the ListTupleJSON in the format [v1, v2].
// The first value v1 is stored at index 0 and the second value v2 is stored at index 1.
func (lt ListTupleJSON[T]) String() string {
	return fmt.Sprintf("[%v, %v]", lt.V1, lt.V2)
}

// GoString returns a string representation of the ListTupleJSON in the format ListTupleJSON[T]{V1, V2}.
// The value T is determined using reflection.
//   - If T cannot be inferred or if it is an interface{}, traditional reflection is used with the type of V1.
//   - If T is still nil, it indicates a rare nil tuple and returns "ListTupleJSON[nil]{nil, nil}".
//
// The V1 and V2 values are represented using the %#v formatting verb.
func (lt ListTupleJSON[T]) GoString() string {
	t := reflect.TypeFor[T]() // first try reflecting the generic
	if t == nil || t.Kind() == reflect.Interface {
		// where we are unable to infer from T (e.g., interface{}),
		// try using traditional reflection
		t = reflect.TypeOf(lt.V1)
	}

	if t == nil {
		// if we still have a nil, we have a rare nil tuple
		// why would you ever do this?
		return "ListTupleJSON[nil]{nil, nil}"
	}

	return fmt.Sprintf("ListTupleJSON[%s]{%#v, %#v}", t, lt.V1, lt.V2)
}

// Slice returns a new array containing the two values stored in the ListTupleJSON.
// The first value is stored at index 0 and the second value is stored at index 1.
func (lt ListTupleJSON[T]) Slice() [2]T {
	return [2]T{lt.V1, lt.V2}
}

// ReadSlice assigns the values in the given slice to the V1 and V2 fields of the ListTupleJSON.
// The first value in the slice is assigned to the V1 field, and the second value is assigned to the V2 field.
func (lt *ListTupleJSON[T]) ReadSlice(slice [2]T) {
	lt.V1, lt.V2 = slice[0], slice[1]
}

// MarshalJSON returns the JSON encoding of the ListTupleJSON by marshaling its Slice().
// The first value is stored at index 0 and the second value is stored at index 1.
func (lt ListTupleJSON[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(lt.Slice())
}

// UnmarshalJSON unmarshals the JSON-encoded data into the ListTupleJSON by
// unmarshalling it into a fixed-size slice. The first value in the slice is
// assigned to the V1 field, and the second value is assigned to the V2 field.
// See the ReadSlice method for more details.
func (lt *ListTupleJSON[T]) UnmarshalJSON(data []byte) error {
	var slice [2]T
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}

	lt.ReadSlice(slice)
	return nil
}
