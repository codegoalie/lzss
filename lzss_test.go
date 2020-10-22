package lzss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoToken(t *testing.T) {
	expected := "This string does not repeat."
	actual := Encode(expected)
	assert.Equal(t, expected, actual, "Encode doesn't replace without repetition")
}

func TestDoubleString(t *testing.T) {
	expected := "abcdef<0,6>"
	in := "abcdefabcdef"
	actual := Encode(in)
	assert.Equal(t, expected, actual, "Encode replaces one repetition")
}

func TestOneToken(t *testing.T) {
	expected := "This stringer and that <16,6>"
	in := "This stringer and that stringer"
	actual := Encode(in)
	assert.Equal(t, expected, actual, "Encode replaces one repetition")
}
