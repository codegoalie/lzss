package main

import (
	"fmt"

	"github.com/codegoalie/lzss"
)

const IAmSam = `I AM SAM. I AM SAM. SAM I AM.

THAT SAM-I-AM! THAT SAM-I-AM! I DO NOT LIKE THAT SAM-I-AM!

DO WOULD YOU LIKE GREEN EGGS AND HAM?

I DO NOT LIKE THEM,SAM-I-AM.
I DO NOT LIKE GREEN EGGS AND HAM.`

func main() {
	// fmt.Println(lzss.Encode("Hello, world!"))
	// fmt.Printf("\n\n")
	// fmt.Println(lzss.Encode("abcdefabcdef"))
	// fmt.Printf("\n\n")
	fmt.Println(lzss.Encode("This stringer and that stringer"))
	// fmt.Printf("\n\n")
	// fmt.Println(lzss.Encode(IAmSam))
}
