package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("num of different bits %d", popCount(c1, c2))
}

func popCount(hash1 [32]byte, hash2 [32]byte) int {
	count := 0
	for i, _ := range hash1 {
		count += int(pc[byte(hash1[i]^hash2[i])])
	}
	return count
}
