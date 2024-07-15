package pseudo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntBool_MarshalJSON(t *testing.T) {
	type TestCase struct {
		name    string
		ib      IntBool
		want    string
		wantErr bool
	}

	tests := []TestCase{
		{name: "true", ib: true, want: "1", wantErr: false},
		{name: "false", ib: false, want: "0", wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ib.MarshalJSON()
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, got, []byte(tt.want))
		})
	}
}

func TestIntBool_UnmarshalJSON(t *testing.T) {
	type TestCase struct {
		name    string
		json    string
		want    IntBool
		wantErr bool
	}

	tests := []TestCase{
		{name: "number 1", json: "1", want: true, wantErr: false},
		{name: "number 0", json: "0", want: false, wantErr: false},
		{name: "non-number", json: `"foo"`, want: false, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var i IntBool
			err := i.UnmarshalJSON([]byte(tt.json))

			if tt.wantErr {
				assert.Error(t, err, "it should error", tt.json)
			} else {
				assert.NoError(t, err, "it should not error", tt.json)
				assert.Equal(t, tt.want, i)
			}

		})
	}
}
