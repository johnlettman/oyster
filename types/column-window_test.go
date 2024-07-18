package types

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColumnWindow_StringInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(ColumnWindow))
	assert.Implements(t, (*fmt.GoStringer)(nil), new(ColumnWindow))
}

func TestColumnWindow_JSONInterfaces(t *testing.T) {
	assert.Implements(t, (*json.Marshaler)(nil), new(ColumnWindow))
	assert.Implements(t, (*json.Unmarshaler)(nil), new(ColumnWindow))
}

func TestColumnWindow_String(t *testing.T) {
	a, b := gofakeit.Int(), gofakeit.Int()
	cw := ColumnWindow{a, b}
	want := fmt.Sprintf("%dx%d", a, b)
	got := cw.String()

	assert.Equal(t, want, got, "it should return the correct representation")
}

func TestColumnWindow_GoString(t *testing.T) {
	a, b := gofakeit.Int(), gofakeit.Int()
	cw := ColumnWindow{a, b}
	want := fmt.Sprintf("ColumnWindow{%d, %d}", a, b)
	got := cw.GoString()

	assert.Equal(t, want, got, "it should return the correct representation")
}

func TestColumnWindow_MarshalJSON(t *testing.T) {
	a, b := gofakeit.Int(), gofakeit.Int()
	cw := ColumnWindow{a, b}
	want := fmt.Sprintf("[%d,%d]", a, b)
	got, err := cw.MarshalJSON()

	assert.NoError(t, err, "it should not error")
	assert.Equal(t, want, string(got), "it should return the correct representation")
}

func TestColumnWindow_UnmarshalJSON(t *testing.T) {
	a, b := gofakeit.Int(), gofakeit.Int()
	data := []byte(fmt.Sprintf("[%d,%d]", a, b))
	cw := ColumnWindow{}
	err := cw.UnmarshalJSON(data)

	assert.NoError(t, err, "it should not error")
	assert.Equal(t, ColumnWindow{a, b}, cw, "it should assign the correct values")
}

func TestColumnWindow_Zero(t *testing.T) {
	cw := ColumnWindow{}

	assert.Zero(t, cw, "it should be zero-initialized")
	assert.True(t, cw.Zero(), "it should return true for zero-initialized ColumnWindow")
}
