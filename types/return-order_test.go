package types

import (
	"encoding"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnOrder_TextInterfaces(t *testing.T) {
	assert.Implements(t, (*encoding.TextMarshaler)(nil), new(ReturnOrder))
	assert.Implements(t, (*encoding.TextUnmarshaler)(nil), new(ReturnOrder))
}

func TestReturnOrder_StringInterfaces(t *testing.T) {
	assert.Implements(t, (*fmt.Stringer)(nil), new(ReturnOrder))
	assert.Implements(t, (*fmt.GoStringer)(nil), new(ReturnOrder))
}

func TestReturnOrder_String(t *testing.T) {
	type TestCase struct {
		name string
		r    ReturnOrder
		want string
	}

	cases := []TestCase{
		{"ReturnOrderUnspecified", ReturnOrderUnspecified, "unspecified"},
		{"ReturnOrderStrongestToWeakest", ReturnOrderStrongestToWeakest, "strongest to weakest"},
		{"ReturnOrderNearestToFarthest", ReturnOrderNearestToFarthest, "nearest to farthest"},
		{"ReturnOrderFarthestToNearest", ReturnOrderFarthestToNearest, "farthest to nearest"},
		{"ReturnOrderStrongestReturnFirst", ReturnOrderStrongestReturnFirst, "strongest return first"},
		{"ReturnOrderLastReturnFirst", ReturnOrderLastReturnFirst, "last return first"},
		{"unknown value", ReturnOrderLastReturnFirst + 1, "unspecified"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.r.String()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestReturnOrder_GoString(t *testing.T) {
	type TestCase struct {
		name string
		r    ReturnOrder
		want string
	}

	cases := []TestCase{
		{"ReturnOrderUnspecified", ReturnOrderUnspecified, "ReturnOrderUnspecified"},
		{"ReturnOrderStrongestToWeakest", ReturnOrderStrongestToWeakest, "ReturnOrderStrongestToWeakest"},
		{"ReturnOrderNearestToFarthest", ReturnOrderNearestToFarthest, "ReturnOrderNearestToFarthest"},
		{"ReturnOrderFarthestToNearest", ReturnOrderFarthestToNearest, "ReturnOrderFarthestToNearest"},
		{"ReturnOrderStrongestReturnFirst", ReturnOrderStrongestReturnFirst, "ReturnOrderStrongestReturnFirst"},
		{"ReturnOrderLastReturnFirst", ReturnOrderLastReturnFirst, "ReturnOrderLastReturnFirst"},
		{"unknown value", ReturnOrderLastReturnFirst + 1, "ReturnOrderUnspecified"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.r.GoString()
			assert.Equal(t, c.want, got, "it should return the correct representation")
		})
	}
}

func TestReturnOrder_MarshalText(t *testing.T) {
	type TestCase struct {
		name string
		t    ReturnOrder
		want string
	}

	cases := []TestCase{
		{"ReturnOrderUnspecified", ReturnOrderUnspecified, "UNSPECIFIED"},
		{"ReturnOrderStrongestToWeakest", ReturnOrderStrongestToWeakest, "STRONGEST_TO_WEAKEST"},
		{"ReturnOrderNearestToFarthest", ReturnOrderNearestToFarthest, "NEAREST_TO_FARTHEST"},
		{"ReturnOrderFarthestToNearest", ReturnOrderFarthestToNearest, "FARTHEST_TO_NEAREST"},
		{"ReturnOrderStrongestReturnFirst", ReturnOrderStrongestReturnFirst, "STRONGEST_RETURN_FIRST"},
		{"ReturnOrderLastReturnFirst", ReturnOrderLastReturnFirst, "LAST_RETURN_FIRST"},
		{"unknown value", ReturnOrderLastReturnFirst + 1, "UNSPECIFIED"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.t.MarshalText()
			assert.NoError(t, err, "it should never error")
			assert.Equal(t, []byte(c.want), got, "it should return the correct representation")
		})
	}
}

func TestReturnOrder_UnmarshalText(t *testing.T) {
	type TestCase struct {
		name string
		text string
		want ReturnOrder
	}

	cases := []TestCase{
		{"ReturnOrderUnspecified", "UNSPECIFIED", ReturnOrderUnspecified},
		{"ReturnOrderStrongestToWeakest", "STRONGEST_TO_WEAKEST", ReturnOrderStrongestToWeakest},
		{"ReturnOrderNearestToFarthest", "NEAREST_TO_FARTHEST", ReturnOrderNearestToFarthest},
		{"ReturnOrderFarthestToNearest", "FARTHEST_TO_NEAREST", ReturnOrderFarthestToNearest},
		{"ReturnOrderStrongestReturnFirst", "STRONGEST_RETURN_FIRST", ReturnOrderStrongestReturnFirst},
		{"ReturnOrderLastReturnFirst", "LAST_RETURN_FIRST", ReturnOrderLastReturnFirst},
		{"unknown text", gofakeit.LoremIpsumSentence(4), ReturnOrderUnspecified},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var got ReturnOrder
			err := got.UnmarshalText([]byte(c.text))
			assert.NoError(t, err, "it should never error")
			assert.Equal(t, c.want, got, "it should assign the correct value")
		})
	}
}
