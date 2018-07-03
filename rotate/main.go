package main

import (
	"fmt"
)

func rotate(slice []int, by int) []int {
	dest := make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		srcIndex := (i + by) % len(slice)
		dest[i] = slice[srcIndex]
	}
	return dest
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	arr = rotate(arr, 3)
	fmt.Println(arr)
}
