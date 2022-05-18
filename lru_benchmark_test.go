package main

import (
	"testing"
)

func BenchmarkLRUCache_Get(b *testing.B) {
	lru := New[int, int](5000)
	for i := 0; i < b.N; i++ {
		lru.Put(i, i)
	}
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		lru.Get(i)
	}
}

func BenchmarkLRUCache_Put(b *testing.B) {
	lru := New[int, int](5000)
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		lru.Put(i, i)
	}
}
