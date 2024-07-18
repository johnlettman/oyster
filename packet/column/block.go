package column

import (
	"fmt"
	"github.com/johnlettman/oyster/packet/column/field"
	"strings"
)

type Block map[field.Field]interface{}

func (b Block) String() string {
	s := new(strings.Builder)
	s.WriteString("Block:\n")

	for k, v := range b {
		s.WriteString(fmt.Sprintf("\t%-16s %T(%v)\n", k.String()+":", v, v))
	}

	return s.String()
}

func (b Block) GoString() string {
	s := new(strings.Builder)
	s.WriteString("field.Block{\n")

	for k, v := range b {
		s.WriteString(fmt.Sprintf("\t%-20s %T(%v),\n", k.GoString()+":", v, v))
	}

	s.WriteRune('}')
	return s.String()
}
