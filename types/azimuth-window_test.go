package types

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultAzimuthWindow(t *testing.T) {
	want := AzimuthWindow{MinMillidegrees, MaxMillidegrees}
	got := DefaultAzimuthWindow()
	assert.Equal(t, want, got)
}

func TestAzimuthWindow_StringInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(AzimuthWindow))
	assert.Implements(t, (*fmt.GoStringer)(nil), new(AzimuthWindow))
}

func TestAzimuthWindow_JSONInterfaces(t *testing.T) {
	assert.Implements(t, (*json.Marshaler)(nil), new(AzimuthWindow))
	assert.Implements(t, (*json.Unmarshaler)(nil), new(AzimuthWindow))
}

func TestAzimuthWindow_String(t *testing.T) {
	a, b := gofakeit.Int(), gofakeit.Int()
	aw := AzimuthWindow{a, b}
	want := fmt.Sprintf("%dm°x%dm°", a, b)
	got := aw.String()

	fmt.Println(want, got)

	assert.Equal(t, want, got, "it should return the correct representation")
}

func TestAzimuthWindow_GoString(t *testing.T) {
	a, b := gofakeit.Int(), gofakeit.Int()
	cw := AzimuthWindow{a, b}
	want := fmt.Sprintf("AzimuthWindow{%d, %d}", a, b)
	got := cw.GoString()

	assert.Equal(t, want, got, "it should return the correct representation")
}

func TestAzimuthWindow_MarshalJSON(t *testing.T) {
	a, b := gofakeit.Int(), gofakeit.Int()
	cw := AzimuthWindow{a, b}
	want := fmt.Sprintf("[%d,%d]", a, b)
	got, err := cw.MarshalJSON()

	assert.NoError(t, err, "it should not error")
	assert.Equal(t, want, string(got), "it should return the correct representation")
}

func TestAzimuthWindow_UnmarshalJSON(t *testing.T) {
	a, b := gofakeit.Int(), gofakeit.Int()
	data := []byte(fmt.Sprintf("[%d,%d]", a, b))
	cw := AzimuthWindow{}
	err := cw.UnmarshalJSON(data)

	assert.NoError(t, err, "it should not error")
	assert.Equal(t, AzimuthWindow{a, b}, cw, "it should assign the correct values")
}

func TestAzimuthWindow_Valid(t *testing.T) {
	type TestCase struct {
		name string
		aw   AzimuthWindow
		want bool
	}

	cases := []TestCase{
		{"negative value", AzimuthWindow{-1, -1}, false},
		{"valid value", AzimuthWindow{1, 1024}, true},
		{"too big value", AzimuthWindow{1, MaxMillidegrees + 1}, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.aw.Valid()
			assert.Equal(t, c.want, got, "it should return the correct validity")
		})
	}
}

func TestAzimuthWindow_Size(t *testing.T) {
	a, b := 1024, 24
	cw := AzimuthWindow{a, b}
	want := 1000
	got := cw.Size()

	assert.Equal(t, want, got, "it should return the correct size")
}

func TestAzimuthWindow_Start(t *testing.T) {
	a, b := 1024, 32
	cw := AzimuthWindow{a, b}
	want := 32
	got := cw.Start()

	assert.Equal(t, want, got, "it should return the correct start")
}

func TestAzimuthWindow_End(t *testing.T) {
	a, b := 1024, 32
	cw := AzimuthWindow{a, b}
	want := 1024
	got := cw.End()

	assert.Equal(t, want, got, "it should return the correct representation")
}
