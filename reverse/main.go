package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(bytes []byte) []byte {
	reversedBytes := make([]byte, len(bytes))
	fmt.Printf("len of bytes %d\n", len(bytes))
	for i := len(bytes); i > 0; {
		r, size := utf8.DecodeLastRune(bytes[:i])
		fmt.Printf("i is %d\n", i)
		copy(reversedBytes[len(bytes)-i:], []byte(string(r)))
		i -= size
	}
	copy(bytes, reversedBytes)
	return reversedBytes
}

func main() {
	bytes := []byte("Việt Nam")
	reverse(bytes)
	fmt.Printf("final solution is %s\n", string(bytes))
}
