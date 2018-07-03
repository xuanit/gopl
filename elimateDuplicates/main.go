package main

import (
	"fmt"
)

func elimateDuplicates(strings []string) []string {
	for i := 0; i < len(strings)-1; i++ {
		j := i + 1
		for ; j < len(strings) && strings[j] == strings[i]; j++ {
		}
		dif := j - i
		if dif > 1 && j < len(strings) {
			copy(strings[i+1:], strings[j:])
		}
		if dif > 1 {
			strings = strings[:len(strings)-(dif-1)]
		}
	}
	return strings
}

func main() {
	strings := []string{"a", "ab", "ab", "ab", "b"}
	strings2 := []string{"ab", "ab", "ab"}
	strings = elimateDuplicates(strings)
	fmt.Println(strings)
	strings2 = elimateDuplicates(strings2)
	fmt.Println(strings2)

}
