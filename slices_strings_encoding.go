package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "世界 means world"
	var buf [utf8.UTFMax]byte

	// NOTE: i doesn't go 0,1,2,3,4,
	// it'll go on the string code points indexes
	// 0, 3, 6, 7, 8...17 because of the chinese characters
	for i, r := range s {
		fmt.Printf("[%d] i\n", i)

		// bytes required to encode the rune
		rl := utf8.RuneLen(r)

		// slice offset
		si := i + rl

		// copy rune from string to buffer
		copy(buf[:], s[i:si])

		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}
}
