# LRU Cache 

This is a simple implementation of a Least Recently Used (LRU) Cache in Go.

An LRU Cache will evict the oldest used key whenever the cache gets full.

Example usage
```go
capacity := 100
cache := lru.NewCache(capacity)
cache.Set("im a key", "im a value")
val := cache.Get("im a key")
println(val)
```
