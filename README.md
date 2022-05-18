# LRU Cache

An in-memory LRU cache implementation in Go. This is optimized and production-ready code.
```shell
goos: darwin
goarch: amd64
pkg: github.com/dhwaneetbhatt/lru
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkLRUCache_Get-8          4904259               235.2 ns/op            33 B/op          1 allocs/op
BenchmarkLRUCache_Put-8          5621064               226.0 ns/op            34 B/op          1 allocs/op
```