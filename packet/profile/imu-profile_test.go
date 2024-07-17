package profile

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIMUProfile_String(t *testing.T) {
	p := IMUProfile(gofakeit.Uint8())
	want := "legacy"
	got := p.String()

	assert.Equal(t, want, got, "it should always return 'legacy'")
}

func TestIMUProfile_GoString(t *testing.T) {
	p := IMUProfile(gofakeit.Uint8())
	want := "IMUProfileLegacy"
	got := p.GoString()

	assert.Equal(t, want, got, "it should always return 'IMUProfileLegacy'")
}

func TestIMUProfile_MarshalText(t *testing.T) {
	p := IMUProfile(gofakeit.Uint8())
	want := []byte("LEGACY")
	got, err := p.MarshalText()

	assert.NoError(t, err, "it should never error")
	assert.Equal(t, want, got, "it should always return 'LEGACY'")
}

func TestIMUProfile_UnmarshalText(t *testing.T) {
	var p IMUProfile
	text := []byte(gofakeit.LoremIpsumSentence(4))
	want := IMUProfileLegacy
	err := p.UnmarshalText(text)

	assert.NoError(t, err, "it should never error")
	assert.Equal(t, want, p, "it should always return IMUProfileLegacy")
}
