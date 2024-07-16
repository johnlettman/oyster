package pseudo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageOverrideJSON_MarshalJSON(t *testing.T) {
	type TestCase struct {
		name string
		m    MessageOverrideJSON[int]
		want string
	}

	value := 5

	cases := []TestCase{
		{name: "has Value has Message", m: MessageOverrideJSON[int]{Value: &value, Message: "Test"}, want: `"Test"`},
		{name: "has Value no Message", m: MessageOverrideJSON[int]{Value: &value, Message: ""}, want: `5`},
		{name: "no Value and Message", m: MessageOverrideJSON[int]{Value: nil, Message: ""}, want: `null`},
		{name: "no Value has Message", m: MessageOverrideJSON[int]{Value: nil, Message: "Test"}, want: `"Test"`},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.m.MarshalJSON()
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, []byte(c.want), got)
		})
	}
}

func TestMessageOverrideJSON_UnmarshalJSON(t *testing.T) {
	type TestCase struct {
		name    string
		json    string
		want    MessageOverrideJSON[int]
		wantErr bool
	}

	value := 5

	cases := []TestCase{
		{name: "string message", json: `"here is a message"`, want: MessageOverrideJSON[int]{Message: "here is a message"}, wantErr: false},
		{name: "expected value", json: `5`, want: MessageOverrideJSON[int]{Value: &value, Message: ""}, wantErr: false},
		{name: "unexpected value", json: `{"not an int": true}`, wantErr: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var got MessageOverrideJSON[int]
			err := got.UnmarshalJSON([]byte(c.json))

			if c.wantErr {
				assert.Error(t, err, "it should return an error")
			} else {
				assert.NoError(t, err, "it should not return an error")
				assert.Equal(t, got, c.want)
			}
		})
	}
}
