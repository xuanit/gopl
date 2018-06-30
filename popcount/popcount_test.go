package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCount24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount24(uint64(i))
	}
}

func BenchmarkPopCount23(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount23(uint64(i))
	}
}

func BenchmarkPopCount25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount25(uint64(i))
	}
}
