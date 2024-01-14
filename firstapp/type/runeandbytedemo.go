package main

import (
	"fmt"
)

func main() {
	// Using byte for ASCII character
	var b byte = 'A'
	fmt.Printf("Byte: %c, Unicode Code Point: %d\n", b, b)

	// Using rune for a Unicode character
	var r rune = '世'
	fmt.Printf("Rune: %c, Unicode Code Point: %d\n", r, r)
}