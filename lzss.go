package lzss

import (
	"fmt"
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
			// fmt.Println("never seen", string(char), pos, seen, activeChains, length)
			out.WriteString(orig.String())
			orig.Reset()
			activeChains = []int{}
			length = 0
			continue
		}
		// fmt.Println("seen", string(char), pos, seen, activeChains, length)

		if len(activeChains) < 1 {
			activeChains = seen
			start = pos
			length = 1
			continue
		}

		newChains := make(map[int]struct{})
		// fmt.Println("cmp", activeChains, seen)
		for _, seenAt := range seen {
			for _, startAt := range activeChains {
				// fmt.Printf("%s == %s", string(in[seenAt]), string(in[startAt+length]))
				if in[seenAt] == in[startAt+length] {
					// fmt.Printf(" used")
					newChains[startAt] = struct{}{}
				}
				// fmt.Println()
			}
		}
		// fmt.Println("chains", activeChains, newChains)
		if len(newChains) > 0 {
			chainSlice := make([]int, len(newChains))
			i := 0
			for startAt := range newChains {
				chainSlice[i] = startAt
				i++
			}
			activeChains = chainSlice
			length++
			// fmt.Println("keep going", string(char), pos, seen, activeChains, length)
			// fmt.Println()
			continue
		}

		if len(seen) > 0 {
			origs := orig.String()
			out.WriteString(tokenOrOrig(start-activeChains[len(activeChains)-1], length, origs[:len(origs)-1]))
			orig.Reset()
			orig.WriteByte(char)
			activeChains = seen
			start = pos
			length = 1
			// fmt.Println("stopped, but started", origs[:len(origs)-1], orig.String())
			continue
		}

		out.WriteString(tokenOrOrig(start-activeChains[len(activeChains)-1], length, orig.String()))

		// fmt.Println("reset")
		// fmt.Println()
		orig.Reset()
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
		out.WriteString(tokenOrOrig(start-activeChains[len(activeChains)-1], length, orig.String()))
	}

	// for char, loc := range mem {
	// 	locs := make([]string, len(loc))
	// 	for i, pos := range loc {
	// 		locs[i] = strconv.Itoa(pos)
	// 	}
	// 	fmt.Printf("%s: %s\n", string(char), strings.Join(locs, ","))
	// 	// fmt.Printf("%s: %d\n", string(char), loc)
	// }
	// fmt.Printf("%+v\n", mem)
	return out.String()
}

func tokenOrOrig(start, length int, orig string) string {
	token := fmt.Sprintf("<%d,%d>", start, length)
	if len(token) < length {
		// fmt.Println("using token")
		return token
	}

	// fmt.Println("using orig")
	return orig
}
