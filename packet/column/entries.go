package column

import (
	"fmt"
	"github.com/johnlettman/oyster/packet/column/field"
	"strings"
)

type Entries map[field.Field]field.Structure

// String returns the string representation of an Entries value.
func (e Entries) String() string {
	s := new(strings.Builder)
	s.WriteString("Entries:\n")

	for k, v := range e {
		s.WriteString(fmt.Sprintf("\t%-14s %v,\n", k.String()+":", v))
	}

	return s.String()
}

// GoString returns the Go syntax representation of an Entries value.
func (e Entries) GoString() string {
	s := new(strings.Builder)
	s.WriteString("field.Entries{\n")

	for k, v := range e {
		s.WriteString(fmt.Sprintf("\t%-14s %#v,\n", k.GoString()+":", v))
	}

	s.WriteRune('}')
	return s.String()
}
