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
		ks := fmt.Sprintf("%s:", k.String())
		s.WriteString(fmt.Sprintf("\t%-14s %v,\n", ks, v))
	}

	return s.String()
}

// GoString returns the Go syntax representation of an Entries value.
func (e Entries) GoString() string {
	s := new(strings.Builder)
	s.WriteString("field.Entries{\n")

	for k, v := range e {
		ks := fmt.Sprintf("%#v:", k)
		s.WriteString(fmt.Sprintf("\t%-14s %#v,\n", ks, v))
	}

	s.WriteRune('}')
	return s.String()
}
