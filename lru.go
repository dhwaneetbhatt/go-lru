package main

import "math"

type LRUCache[K comparable, V any] struct {
	store   map[K]V
	maxSize int
	freq    map[K]int
}

func New[K comparable, V any](maxSize int) *LRUCache[K, V] {
	return &LRUCache[K, V]{store: make(map[K]V), maxSize: maxSize, freq: make(map[K]int)}
}

func (c *LRUCache[K, V]) Size() int {
	return len(c.store)
}

func (c *LRUCache[K, V]) Get(key K) V {
	value, ok := c.store[key]
	if ok {
		c.freq[key]++
		return value
	}
	var noop V
	return noop
}

func (c *LRUCache[K, V]) Put(key K, value V) {
	if c.Size() >= c.maxSize {
		var min int = math.MaxInt64
		var toBeEvicted K
		for k, v := range c.freq {
			if v < min {
				min = v
				toBeEvicted = k
			}
		}
		delete(c.store, toBeEvicted)
		delete(c.freq, toBeEvicted)
	}
	c.store[key] = value
	c.freq[key] = 0
}
