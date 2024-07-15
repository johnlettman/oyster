package pseudo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageStruct_MarshalJSON(t *testing.T) {
	type TestCase struct {
		name string
		ms   MessageStruct[int]
		want string
	}

	value := 5

	tests := []TestCase{
		{name: "has Value has Message", ms: MessageStruct[int]{Value: &value, Message: "Test"}, want: `"Test"`},
		{name: "has Value no Message", ms: MessageStruct[int]{Value: &value, Message: ""}, want: `5`},
		{name: "no Value and Message", ms: MessageStruct[int]{Value: nil, Message: ""}, want: `null`},
		{name: "no Value has Message", ms: MessageStruct[int]{Value: nil, Message: "Test"}, want: `"Test"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ms.MarshalJSON()
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, []byte(tt.want), got)
		})
	}
}

func TestMessageStruct_UnmarshalJSON(t *testing.T) {
	type TestCase struct {
		name    string
		json    string
		want    MessageStruct[int]
		wantErr bool
	}

	value := 5

	tests := []TestCase{
		{name: "string message", json: `"here is a message"`, want: MessageStruct[int]{Message: "here is a message"}, wantErr: false},
		{name: "expected value", json: `5`, want: MessageStruct[int]{Value: &value, Message: ""}, wantErr: false},
		{name: "unexpected value", json: `{"not an int": true}`, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got MessageStruct[int]
			err := got.UnmarshalJSON([]byte(tt.json))

			if tt.wantErr {
				assert.Error(t, err, "it should return an error")
			} else {
				assert.NoError(t, err, "it should not return an error")
				assert.Equal(t, got, tt.want)
			}
		})
	}
}
