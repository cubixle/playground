package lru_test

import (
	"testing"

	"github.com/cubixle/lru"
)

func TestCache(t *testing.T) {
	cache := lru.NewCache(2)
	cache.Set("usa", "washington")
	cache.Set("uk", "london")

	city := cache.Get("usa")
	if city != "washington" {
		t.Fatal("didn't get the correct city for usa")
	}

	city = cache.Get("uk")
	if city != "london" {
		t.Fatal("didn't get the correct city for uk")
	}

	cache.Set("france", "paris")

	city = cache.Get("france")
	if city != "paris" {
		t.Fatal("didn't get the correct city for uk")
	}

	city = cache.Get("usa")
	if city != "" {
		t.Fatal("usa is still in the cache but should have been removed")
	}
}
