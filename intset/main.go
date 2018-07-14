package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint
}

const SIZE = 32 << (^(uint(0)) >> 63)

func (s *IntSet) Has(x int) bool {
	word, bit := x/SIZE, uint(x%SIZE)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/SIZE, uint(x%SIZE)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", SIZE*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	length := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < SIZE; j++ {
			if (word & (1 << uint(j))) != 0 {
				length++
			}
		}
	}
	return length
}

func (s *IntSet) Remove(x int) {
	word, bit := x/SIZE, uint(x%SIZE)
	if word >= len(s.words) {
		return
	}

	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	dest := IntSet{words: make([]uint, len(s.words))}
	copy(dest.words, s.words)
	return &dest
}

func (s *IntSet) AddAll(args ...int) {
	for _, v := range args {
		s.Add(v)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	if len(t.words) < len(s.words) {
		s.words = s.words[:len(t.words)]
	}

	for i, tword := range t.words {
		if i >= len(s.words) {
			continue
		}

		s.words[i] &= tword
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, _ := range s.words {
		if i >= len(t.words) {
			continue
		}

		s.words[i] &= ^(t.words[i])
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i >= len(s.words) {
			s.words = append(s.words, tword)
		} else {
			s.words[i] ^= tword
		}
	}
}

func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		for j := 0; j < SIZE; j++ {
			if (word & (1 << uint(j))) != 0 {
				elems = append(elems, i*SIZE+j)
			}
		}
	}
	return elems
}

func main() {
	s1 := IntSet{}
	s1.AddAll(1, 2, 7, 9, 10)

	s2 := IntSet{}
	s2.AddAll(1, 2, 4, 5)

	s1.SymmetricDifference(&s2)

	fmt.Println(s1.String())

	for i, v := range s1.Elems() {
		fmt.Printf("element %d at %d\n", v, i)
	}

	fmt.Printf("Expression %v\n", 32<<(^uint(0)>>63))
}
