package util

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndentAfter(t *testing.T) {
	b := `
This
is
a
test`
	want := `
This
is
XXXXa
XXXXtest`

	start := 3
	prefix := "XXXX"
	got := IndentAfter(b, prefix, start)

	assert.Equal(t, want, got, "it should properly indent after N start lines")
}

func TestIndentBytesAfter(t *testing.T) {
	t.Run("normal buffers", func(t *testing.T) {
		b := []byte(`
This
is
a
test`)
		want := []byte(`
This
is
XXXXa
XXXXtest`)

		start := 3
		prefix := []byte("XXXX")
		got := IndentBytesAfter(b, prefix, start)

		assert.Equal(t, want, got, "it should properly indent after N start lines")
	})

	t.Run("empty buffers", func(t *testing.T) {
		b := []byte("")
		want := []byte("")
		start := gofakeit.Int()
		prefix := []byte(gofakeit.Username())
		got := IndentBytesAfter(b, prefix, start)

		assert.Equal(t, want, got, "it should return an empty buffer for an empty input")
	})

}
