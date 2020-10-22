package lzss

import (
	"fmt"
	"strconv"
	"strings"
)

func Encode(in string) string {
	mem := map[byte][]int{}
	out := strings.Builder{}

	orig := strings.Builder{}
	activeChains := []int{}
	var start int
	length := 0

	for pos := 0; pos < len(in); pos++ {
		char := in[pos]

		orig.WriteByte(char)

		seen, ok := mem[char]
		mem[char] = append(mem[char], pos)
		if !ok {
			fmt.Println("never seen", string(char), pos, seen, activeChains, length)
			out.WriteString(orig.String())
			orig = strings.Builder{}
			activeChains = []int{}
			length = 0
			continue
		}

		if len(activeChains) < 1 {
			activeChains = seen
			start = pos
			length = 1
			continue
		}

		newChains := []int{}
		fmt.Println("cmp", activeChains, seen)
		for _, seenAt := range seen {
			for _, startAt := range activeChains {
				if in[seenAt] == in[startAt+length] {
					newChains = append(newChains, startAt)
				}
			}
		}
		fmt.Println("starts", activeChains, newChains)
		if len(newChains) > 0 {
			activeChains = newChains
			length++
			fmt.Println("keep going", string(char), pos, seen, activeChains, length)
			continue
		}

		token := fmt.Sprintf("<%d,%d>", start-activeChains[len(activeChains)-1], length)
		if len(token) < length {
			fmt.Println("using token", string(char), pos, seen, activeChains, length)
			out.WriteString(token)
		} else {
			fmt.Println("using orig", string(char), pos, seen, activeChains, length)
			out.WriteString(orig.String())
		}

		fmt.Println("reset")
		orig = strings.Builder{}
		activeChains = []int{}
		length = 0

		// printed := false
		// for length := 0; pos+length < len(in); length++ {
		// 	if in[pos+length+1] == in[seen+length+1] {
		// 		continue
		// 	}

		// 	token := fmt.Sprintf("<%d,%d>", pos-length, length)
		// 	fmt.Println(len(token), length)
		// 	if len(token) < length {
		// 		out.WriteString(token)
		// 		printed = true
		// 		break
		// 	}
		// }
		// if !printed {
		// 	out.WriteRune(char)
		// }
		// mem[char] = pos
	}
	if length > 0 {
		token := fmt.Sprintf("<%d,%d>", start-activeChains[len(activeChains)-1], length)
		if len(token) < length {
			out.WriteString(token)
		} else {
			out.WriteString(orig.String())
		}
	}

	for char, loc := range mem {
		locs := make([]string, len(loc))
		for i, pos := range loc {
			locs[i] = strconv.Itoa(pos)
		}
		fmt.Printf("%s: %s\n", string(char), strings.Join(locs, ","))
		// fmt.Printf("%s: %d\n", string(char), loc)
	}
	// fmt.Printf("%+v\n", mem)
	return out.String()
}
