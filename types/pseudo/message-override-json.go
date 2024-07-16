package pseudo

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MessageOverrideJSON represents a generic message structure that may be overridden by a string in the JSON representation.
type MessageOverrideJSON[T any] struct {
	Value   *T
	Message string
}

// MarshalJSON serializes the MessageOverrideJSON type to JSON representation.
// If MessageOverrideJSON has a non-empty Message or HasMessage is true,
// it marshals only the Message field as a string. Otherwise, it marshals the Value field.
// It returns a byte slice and an error.
func (m MessageOverrideJSON[T]) MarshalJSON() ([]byte, error) {
	if strings.TrimSpace(m.Message) != "" {
		return json.Marshal(m.Message)
	}

	return json.Marshal(m.Value)
}

// UnmarshalJSON deserializes JSON data into a MessageOverrideJSON pointer.
// If the JSON is a string, it assigns the string to the Message field.
// If the JSON is T, it assigns the value to the Value field.
// In case the JSON data cannot be unmarshaled into a string or T, it returns an error.
func (m *MessageOverrideJSON[T]) UnmarshalJSON(data []byte) error {
	err := UnmarshalMessageOrStruct(data, &m.Message, &m.Value)
	if err != nil {
		return err
	}

	return nil
}

// UnmarshalMessageOrStruct will unmarshal JSON data into a target message or struct depending on the
// actual representation in the JSON data.
// This aids in conditions where the data overrides an anticipated structure with an error message.
func UnmarshalMessageOrStruct(data []byte, targetMessage *string, targetStruct interface{}) error {
	if err := json.Unmarshal(data, targetMessage); err == nil {
		return nil
	}
	if err := json.Unmarshal(data, targetStruct); err == nil {
		return nil
	}
	return fmt.Errorf("unable to unmarshal data as string or struct")
}
