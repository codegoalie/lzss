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
	expected := "abcdef<6,6>"
	in := "abcdefabcdef"
	actual := Encode(in)
	assert.Equal(t, expected, actual, "Encode replaces one repetition")
}

func TestDoubleStringWithSuffix(t *testing.T) {
	expected := "abcdef<6,6>g"
	in := "abcdefabcdefg"
	actual := Encode(in)
	assert.Equal(t, expected, actual, "Encode replaces one repetition")
}

func TestOneToken(t *testing.T) {
	expected := "This stringer and that<18,9>"
	in := "This stringer and that stringer"
	actual := Encode(in)
	assert.Equal(t, expected, actual, "Encode replaces one repetition")
}

func TestSamIAm(t *testing.T) {
	expected := `I AM SAM. <10,10>SAM I AM.

THAT SAM-I-AM! <15,15>I DO NOT LIKE<29,15>

DO WOULD YOU LIKE GREEN EGGS AND HAM?

I<69,15>EM,<113,8>.<29,15>GR<64,16>.`
	in := `I AM SAM. I AM SAM. SAM I AM.

THAT SAM-I-AM! THAT SAM-I-AM! I DO NOT LIKE THAT SAM-I-AM!

DO WOULD YOU LIKE GREEN EGGS AND HAM?

I DO NOT LIKE THEM,SAM-I-AM.
I DO NOT LIKE GREEN EGGS AND HAM.`
	actual := Encode(in)
	assert.Equal(t, expected, actual, "Encode handles Sam I Am")
}

func TestSamIAmSmall(t *testing.T) {
	expected := `I AM SAM. <10,10>SAM I AM.`
	in := `I AM SAM. I AM SAM. SAM I AM.`
	actual := Encode(in)
	assert.Equal(t, expected, actual, "Encode handles small Sam I Am")
}
