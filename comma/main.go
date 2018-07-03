package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	// n := len(s)
	// if n <= 3 {
	// 	return s
	// }

	var lenghtOfNumber = strings.Index(s, ".")
	if lenghtOfNumber == -1 {
		lenghtOfNumber = len(s)
	}

	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		buf.WriteByte(s[i])
		remaining := lenghtOfNumber - i - 1
		if remaining > 0 && remaining%3 == 0 {
			buf.WriteRune(',')
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("12345") == "12,345")
	fmt.Println(comma("123") == "123")
	fmt.Println(comma("123456") == "123,456")
	fmt.Println(comma("123456.1") == "123,456.1")
	fmt.Println(comma("+12345.12") == "+12,345.12")
	fmt.Println(comma("-12345.") == "-12,345.")
}
