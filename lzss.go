package lzss

import (
	"fmt"
	"strings"
)

func Encode(in string) string {
	mem := map[byte]int{}
	out := strings.Builder{}

	orig := strings.Builder{}
	start := -1
	length := 0

	for pos := 0; pos < len(in); pos++ {
		char := in[pos]

		orig.WriteByte(char)

		seen, ok := mem[char]
		if !ok {
			mem[char] = pos
			out.WriteString(orig.String())
			orig = strings.Builder{}
			start = -1
			length = 0
			continue
		}

		if start == -1 {
			start = seen
		}

		if in[seen] == in[start+length] {
			length++
			continue
		}

		token := fmt.Sprintf("<%d,%d>", start, length)
		if len(token) < length {
			out.WriteString(token)
		} else {
			out.WriteString(orig.String())
		}

		orig = strings.Builder{}
		start = -1
		length = 0
	}
	if length > 0 {
		token := fmt.Sprintf("<%d,%d>", start, length)
		if len(token) < length {
			out.WriteString(token)
		} else {
			out.WriteString(orig.String())
		}
	}

	return out.String()
}
