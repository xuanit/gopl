package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	r := bufio.NewScanner(bytes.NewReader(p))
	r.Split(bufio.ScanWords)
	for r.Scan() {
		fmt.Println(r.Text())
		*c += 1
	}
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	r := bufio.NewScanner(bytes.NewReader(p))
	for r.Scan() {
		fmt.Println(r.Text())
		*c++
	}
	return len(p), nil
}

func main() {
	var counter LineCounter

	fmt.Fprint(&counter, "hello world 123\n4")

	fmt.Printf("counter %v\n", counter)
}
