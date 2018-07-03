package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	const (
		LETTER = "letter"
		DIGIT  = "digit"
		OTHER  = "other"
	)

	counts := make(map[string]int)
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(r) {
			counts[LETTER]++
			continue
		}

		if unicode.IsDigit(r) {
			counts[DIGIT]++
			continue
		}

		counts[OTHER]++

	}

	fmt.Printf("category\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}
}
