package pseudo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListTupleJSON_JSONInterfaces(t *testing.T) {
	assert.Implements(t, (*json.Marshaler)(nil), new(ListTupleJSON[interface{}]))
	assert.Implements(t, (*json.Unmarshaler)(nil), new(ListTupleJSON[interface{}]))
}

func TestListTupleJSON_String(t *testing.T) {
	type TestCase struct {
		name string
		lt   ListTupleJSON[interface{}]
		want string
	}

	cases := []TestCase{
		{"int", ListTupleJSON[interface{}]{1, 2}, "[1, 2]"},
		{"string", ListTupleJSON[interface{}]{"hello", "world"}, "[hello, world]"},
		{"nil", ListTupleJSON[interface{}]{nil, nil}, "[<nil>, <nil>]"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.lt.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestListTupleJSON_GoString(t *testing.T) {
	type TestCase struct {
		name string
		lt   ListTupleJSON[interface{}]
		want string
	}

	cases := []TestCase{
		{"int", ListTupleJSON[interface{}]{1, 2}, "ListTupleJSON[int]{1, 2}"},
		{"string", ListTupleJSON[interface{}]{"hello", "world"}, `ListTupleJSON[string]{"hello", "world"}`},
		{"nil", ListTupleJSON[interface{}]{nil, nil}, "ListTupleJSON[nil]{nil, nil}"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.lt.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestListTupleJSON_Slice(t *testing.T) {
	type TestCase struct {
		name string
		lt   ListTupleJSON[interface{}]
		want [2]interface{}
	}

	cases := []TestCase{
		{"int", ListTupleJSON[interface{}]{1, 2}, [2]interface{}{1, 2}},
		{"string", ListTupleJSON[interface{}]{"hello", "world"}, [2]interface{}{"hello", "world"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.lt.Slice()
			assert.Equal(t, c.want, got, "it should return the correct slice")
		})
	}
}

func TestListTupleJSON_ReadSlice(t *testing.T) {
	type TestCase struct {
		name  string
		slice [2]interface{}
		want  ListTupleJSON[interface{}]
	}

	cases := []TestCase{
		{"int", [2]interface{}{1, 2}, ListTupleJSON[interface{}]{1, 2}},
		{"string", [2]interface{}{"hello", "world"}, ListTupleJSON[interface{}]{"hello", "world"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ListTupleJSON[interface{}]{}
			got.ReadSlice(c.slice)
			assert.Equal(t, c.want, got, "it should change the values to match the slice")
		})
	}
}

func TestListTupleJSON_MarshalJSON(t *testing.T) {
	type TestCase struct {
		name string
		lt   ListTupleJSON[interface{}]
		want string
	}

	cases := []TestCase{
		{"int", ListTupleJSON[interface{}]{1, 2}, "[1,2]"},
		{"string", ListTupleJSON[interface{}]{"hello", "world"}, `["hello","world"]`},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.lt.MarshalJSON()
			assert.NoError(t, err, "it should not error")
			assert.NotNil(t, got, "it should return a value")
			assert.Equal(t, []byte(c.want), got, "it should marshal the correct JSON")
		})
	}
}

func TestListTupleJSON_UnmarshalJSON(t *testing.T) {
	type TestCase struct {
		name    string
		data    string
		want    ListTupleJSON[interface{}]
		wantErr bool
	}

	cases := []TestCase{
		{"number", "[1,2]", ListTupleJSON[interface{}]{float64(1), float64(2)}, false},
		{"string", `["hello","world"]`, ListTupleJSON[interface{}]{"hello", "world"}, false},
		{"invalid JSON", `'[21}[34""'`, ListTupleJSON[interface{}]{}, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := new(ListTupleJSON[interface{}])
			err := got.UnmarshalJSON([]byte(c.data))
			if c.wantErr {
				assert.Error(t, err, "it should error")
			} else {
				assert.NoError(t, err, "it should not error")
				assert.Equal(t, c.want, *got, "it should unmarshal the JSON correctly")
			}
		})
	}
}
