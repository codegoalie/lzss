package main

import (
	"fmt"

	"github.com/codegoalie/lzss"
)

func main() {
	// fmt.Println(lzss.Encode("Hello, world!"))
	fmt.Println(lzss.Encode("abcdefabcdef"))
}
