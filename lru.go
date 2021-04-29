package main

import "math"

type LRUCache struct {
	store   map[string]string
	maxSize int
	freq    map[string]int
}

func New(maxSize int) *LRUCache {
	return &LRUCache{store: make(map[string]string), maxSize: maxSize, freq: make(map[string]int)}
}

func (c *LRUCache) Size() int {
	return len(c.store)
}

func (c *LRUCache) Get(key string) string {
	value, ok := c.store[key]
	if ok {
		c.freq[key]++
		return value
	}
	return ""
}

func (c *LRUCache) Put(key string, value string) {
	if c.Size() >= c.maxSize {
		var min int = math.MaxInt64
		var toBeEvicted string
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
