package types

import (
	"encoding"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAutoStartFlag_TextInterfaces(t *testing.T) {
	assert.Implements(t, (*encoding.TextMarshaler)(nil), new(AutoStartFlag))
	assert.Implements(t, (*encoding.TextUnmarshaler)(nil), new(AutoStartFlag))
}

func TestAutoStartFlag_String(t *testing.T) {
	type TestCase struct {
		a    AutoStartFlag
		want string
	}

	cases := []TestCase{
		{AutoStartOn, "auto start on"},
		{AutoStartOff, "auto start off"},
	}

	for _, c := range cases {
		got := c.a.String()
		assert.Equal(t, c.want, got, "it should return the correct representation")
	}
}

func TestAutoStartFlag_GoString(t *testing.T) {
	type TestCase struct {
		a    AutoStartFlag
		want string
	}

	cases := []TestCase{
		{AutoStartOn, "AutoStartOn"},
		{AutoStartOff, "AutoStartOff"},
	}

	for _, c := range cases {
		got := c.a.GoString()
		assert.Equal(t, c.want, got, "it should return the correct representation")
	}
}

func TestAutoStartFlag_MarshalText(t *testing.T) {
	type TestCase struct {
		a    AutoStartFlag
		want string
	}

	cases := []TestCase{
		{AutoStartOn, "on"},
		{AutoStartOff, "off"},
	}

	for _, c := range cases {
		got, err := c.a.MarshalText()
		assert.NoError(t, err, "it should not error")
		assert.Equal(t, []byte(c.want), got, "it should return the correct representation")
	}
}

func TestAutoStartFlag_UnmarshalText(t *testing.T) {
	type TestCase struct {
		name string
		text string
		want AutoStartFlag
	}

	cases := []TestCase{
		{"lowercase 'on'", "on", AutoStartOn},
		{"lowercase 'off'", "off", AutoStartOff},
		{"uppercase 'ON'", "ON", AutoStartOn},
		{"uppercase 'OFF'", "OFF", AutoStartOff},
		{"random text", gofakeit.LoremIpsumSentence(4), AutoStartOn},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var a AutoStartFlag
			err := a.UnmarshalText([]byte(c.text))
			assert.NoError(t, err, "it should not error")
			assert.Equal(t, c.want, a, "it should assign the correct value")
		})
	}

}

func TestAutoStartFlag_MarshalJSON(t *testing.T) {
	type TestCase struct {
		a    AutoStartFlag
		want string
	}

	cases := []TestCase{
		{AutoStartOn, "1"},
		{AutoStartOff, "0"},
	}

	for _, c := range cases {
		got, err := c.a.MarshalJSON()
		assert.NoError(t, err, "it should not error")
		assert.Equal(t, got, []byte(c.want), "it should provide the correct JSON using IntBool")
	}
}

func TestAutoStartFlag_UnmarshalJSON(t *testing.T) {
	type TestCase struct {
		name    string
		json    string
		want    AutoStartFlag
		wantErr bool
	}

	cases := []TestCase{
		{name: "number 1", json: "1", want: true, wantErr: false},
		{name: "number 0", json: "0", want: false, wantErr: false},
		{name: "non-number", json: `"foo"`, want: false, wantErr: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var a AutoStartFlag
			err := a.UnmarshalJSON([]byte(c.json))

			if c.wantErr {
				assert.Error(t, err, "it should error")
			} else {
				assert.NoError(t, err, "it should not error")
				assert.Equal(t, c.want, a, "it should assign the correct value using IntBool")
			}

		})
	}
}
