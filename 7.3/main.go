package main

import (
	"bytes"
	"fmt"
	"strconv"
)

//!+
type tree struct {
	value       *int
	left, right *tree
}

func (t *tree) add(value int) {
	// fmt.Printf("adding into tree %v k\n ", *t)
	if t.value == nil {
		newVal := int(value)
		t.value = &newVal
		return
	}

	if value < *(t.value) {
		if t.left == nil {
			t.left = &tree{}
		}
		t.left.add(value)
	} else {
		if t.right == nil {
			t.right = &tree{}
		}
		t.right.add(value)
	}
}

func (t *tree) String() string {
	var buf bytes.Buffer
	if t == nil {
		return ""
	}

	buf.WriteString(t.left.String())
	buf.WriteByte(' ')
	buf.WriteString(strconv.Itoa(*t.value))
	buf.WriteString(t.right.String())
	return buf.String()
}

func main() {
	t := tree{}
	t.add(4)
	t.add(1)
	t.add(2)
	t.add(5)
	t.add(6)
	fmt.Println(t.String())
}

//!-
