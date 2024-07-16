package pseudo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntBoolJSON_MarshalJSON(t *testing.T) {
	type TestCase struct {
		name    string
		i       IntBoolJSON
		want    string
		wantErr bool
	}

	cases := []TestCase{
		{name: "true", i: IntBoolJSON(true), want: "1", wantErr: false},
		{name: "false", i: IntBoolJSON(false), want: "0", wantErr: false},
		{name: "native true", i: true, want: "1", wantErr: false},
		{name: "native false", i: false, want: "0", wantErr: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.i.MarshalJSON()
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, got, []byte(c.want))
		})
	}
}

func TestIntBoolJSON_UnmarshalJSON(t *testing.T) {
	type TestCase struct {
		name    string
		json    string
		want    IntBoolJSON
		wantErr bool
	}

	cases := []TestCase{
		{name: "number 1", json: "1", want: true, wantErr: false},
		{name: "number 0", json: "0", want: false, wantErr: false},
		{name: "non-number", json: `"foo"`, want: false, wantErr: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var i IntBoolJSON
			err := i.UnmarshalJSON([]byte(c.json))

			if c.wantErr {
				assert.Error(t, err, "it should error", c.json)
			} else {
				assert.NoError(t, err, "it should not error", c.json)
				assert.Equal(t, c.want, i)
			}

		})
	}
}
