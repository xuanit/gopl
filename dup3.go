package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main_dup3() {
	counts := make(map[string]int)
	locations := make(map[string]map[string]bool)
	files := os.Args[1:]
	for _, arg := range files {
		data, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if locations[line] == nil {
				locations[line] = make(map[string]bool)
			}
			locations[line][arg] = true
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)
			for file, _ := range locations[line] {
				fmt.Printf("%s\t", file)
			}
			fmt.Printf("\n")
		}
	}
}
