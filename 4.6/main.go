package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squash(bytes []byte) []byte {
	for i := 0; i < len(bytes); {
		r, size := utf8.DecodeRune(bytes[i:])
		i += size
		fmt.Printf("processing character %s at %d\n", string(r), i)
		if !unicode.IsSpace(r) {
			continue
		}

		j := i
		for j < len(bytes) {
			r2, size2 := utf8.DecodeRune(bytes[j:])
			if !unicode.IsSpace(r2) {
				break
			}
			j += size2
		}
		if j > i && j < len(bytes) {
			copy(bytes[i:], bytes[j:])
		}
		if j > i {
			bytes = bytes[:len(bytes)-(j-i)]
		}
		fmt.Printf("length of bytes is %d\n", len(bytes))
	}
	return bytes
}

func main() {
	value := "    việt       nam    "
	value = string(squash([]byte(value)))
	fmt.Println(value)
	unicode.IsSpace(rune(11))
}
