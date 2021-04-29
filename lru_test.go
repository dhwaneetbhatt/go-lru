package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLRU(t *testing.T) {
	lru := New(5)
	require.Equal(t, 0, lru.freq["key1"])
	require.Equal(t, 0, lru.Size())
	lru.Put("key1", "val1")
	require.Equal(t, "val1", lru.Get("key1"))
	require.Equal(t, 1, lru.freq["key1"])
	require.Equal(t, 1, lru.Size())

	lru.Put("key2", "val2")
	lru.Get("key2")
	lru.Put("key3", "val3")
	lru.Get("key3")
	lru.Put("key4", "val4")
	lru.Get("key4")
	lru.Put("key5", "val5")
	require.Equal(t, 5, lru.Size())

	lru.Put("key6", "val6")
	require.Equal(t, 5, lru.Size())

	require.Equal(t, "val6", lru.Get("key6"))
	require.Equal(t, "", lru.Get("key5"))
}
