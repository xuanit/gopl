package main

import (
	"fmt"
	"io"
	"os"
)

type Writer struct {
	io.Writer
	counter *int64
}

func (w Writer) Write(p []byte) (i int, err error) {
	n, err := w.Writer.Write(p)
	*(w.counter) += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var counter int64
	return Writer{
		Writer:  w,
		counter: &counter,
	}, &counter
}

func main() {
	w, count := CountingWriter(os.Stdout)
	fmt.Fprint(w, "h")
	fmt.Printf("counting %d\n", *count)

	fmt.Fprint(w, "world")
	fmt.Printf("counting %d\n", *count)
}
