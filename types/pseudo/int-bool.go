package pseudo

import "encoding/json"

// IntBool is a custom type that represents a boolean value as an integer.
// It is used for JSON marshaling and unmarshaling purposes.
type IntBool bool

// MarshalJSON is a method that serializes the IntBool value to JSON format.
// It converts the boolean value to an integer (1 or 0), then marshals the integer to JSON.
// The method returns the marshaled JSON bytes and an error, if any.
func (b IntBool) MarshalJSON() ([]byte, error) {
	var v int
	if b {
		v = 1
	} else {
		v = 0
	}

	return json.Marshal(v)
}

// UnmarshalJSON is a method that deserializes the JSON data into the IntBool value.
// It unmarshals the received JSON data into an integer value, then assigns the IntBool
// value based on whether the integer is non-zero or zero. The IntBool value is set to
// true if the integer is non-zero, and false if the integer is zero.
// The method returns an error if the unmarshaling process fails.
func (b *IntBool) UnmarshalJSON(data []byte) error {
	var v int
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	*b = v != 0
	return nil
}
