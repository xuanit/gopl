package main

import (
	"fmt"
	"sort"
)

func isAnagrams(s1, s2 string) bool {
	s1Runes := []rune(s1)
	s2Runes := []rune(s2)
	sort.Slice(s1Runes, func(i int, j int) bool {
		return s1Runes[i] < s1Runes[j]
	})
	sort.Slice(s2Runes, func(i int, j int) bool {
		return s2Runes[i] < s2Runes[j]
	})

	if len(s1Runes) != len(s2Runes) {
		return false
	}

	for i, _ := range s1Runes {
		if s1Runes[i] != s2Runes[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagrams("việt", "ệtiv") == true)
	fmt.Println(isAnagrams("việt", "việ") == false)
	fmt.Println(isAnagrams("viet", "teiv") == true)
	fmt.Println(isAnagrams("viet", "tev") == false)
}
