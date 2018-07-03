package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freqs := make(map[string]int)

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
	}

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		freqs[word]++
	}

	fmt.Printf("word\tcount\n")
	for w, n := range freqs {
		fmt.Printf("%s\t%d\n", w, n)
	}
}
