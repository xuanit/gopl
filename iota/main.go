package main

import (
	"fmt"
)

const delta = 1000

const (
	_ = delta * iota
	KB
	MB
)

func main() {
	fmt.Printf("%d", MB)
}
