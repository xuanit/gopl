package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount23(x uint64) int {
	var bitCount = 0
	for i := 0; i < 8; i++ {
		bitCount += int(pc[byte(x>>(uint64(i)*8))])
	}
	return bitCount
}

func PopCount24(x uint64) int {
	var bitCount = 0
	for i := 0; i < 64; i++ {
		bitCount += int((x >> uint64(i))) & 1
	}
	return bitCount
}

func PopCount25(x uint64) int {
	var bitCount = 0
	var remainingValue = x
	for remainingValue != 0 {
		remainingValue = remainingValue & (remainingValue - 1)
		bitCount++
	}
	return bitCount
}

func main() {
	fmt.Printf("%d", PopCount(4))
}
