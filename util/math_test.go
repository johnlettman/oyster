package util

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"reflect"
	"testing"
)

import "github.com/brianvoe/gofakeit/v7"

type TestCaseAbs[T constraints.Integer | constraints.Float] struct {
	Input T
	Want  T
}

const CaseCountAbs = 10

var testAbsFakeFunctions = map[string]func() interface{}{
	"uint8":   func() interface{} { return gofakeit.Uint8() },
	"uint16":  func() interface{} { return gofakeit.Uint16() },
	"uint32":  func() interface{} { return gofakeit.Uint32() },
	"uint64":  func() interface{} { return gofakeit.Uint64() },
	"int8":    func() interface{} { return gofakeit.Int8() },
	"int16":   func() interface{} { return gofakeit.Int16() },
	"int32":   func() interface{} { return gofakeit.Int32() },
	"int64":   func() interface{} { return gofakeit.Int64() },
	"float32": func() interface{} { return gofakeit.Float32() },
	"float64": func() interface{} { return gofakeit.Float64() },
}

func testAbsCase[T constraints.Integer | constraints.Float](t *testing.T, c TestCaseAbs[T]) {
	got := Abs(c.Input)
	assert.Equalf(t, got, c.Want, "it should return %v for %v with generic of type %T", c.Want, got, *new(T))
}

func testAbsGeneric[T constraints.Integer | constraints.Float](t *testing.T) {
	n := reflect.TypeFor[T]().Name()
	f, ok := testAbsFakeFunctions[n]
	if !ok {
		t.Fatalf("can not locate randomizer function for type %s", n)
		return
	}

	t.Run(n, func(t *testing.T) {
		for i := 0; i < CaseCountAbs; i++ {
			test := TestCaseAbs[T]{}
			input := f().(T)

			test.Input = input

			if input < 0 {
				test.Want = -input
			} else {
				test.Want = input
			}

			testAbsCase(t, test)
		}
	})
}

func TestAbs(t *testing.T) {
	testAbsGeneric[uint8](t)
	testAbsGeneric[uint16](t)
	testAbsGeneric[uint32](t)
	testAbsGeneric[uint64](t)
	testAbsGeneric[int8](t)
	testAbsGeneric[int16](t)
	testAbsGeneric[int32](t)
	testAbsGeneric[int64](t)
	testAbsGeneric[float32](t)
	testAbsGeneric[float64](t)
}
