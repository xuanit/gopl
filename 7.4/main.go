package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Reader struct {
	p []byte
}

func (reader *Reader) Read(p []byte) (int, error) {
	var count int
	for len(reader.p) > 0 {
		_, size := utf8.DecodeRune(reader.p)

		if count+size > len(p) {
			return count, nil
		}

		copy(p[count:], reader.p[:size])
		reader.p = reader.p[size:]
		count += size
	}

	return count, nil
}

func NewReader(s string) *Reader {
	r := Reader{}
	return &(r)
}

func main() {
	r := strings.NewReader("helloagain")

	w1 := make([]byte, 5)
	w2 := make([]byte, 5)

	r.Read(w1)
	r.Read(w2)

	fmt.Printf("word1 %s\n", string(w1))
	fmt.Printf("word2 %s", string(w2))
}
