package memory

import (
	cachetest "github.com/pierrre/imageserver/cache/_test"
	"testing"
)

func TestGetSet(t *testing.T) {
	cache := newTestCache()

	cachetest.CacheTestGetSetAllImages(t, cache)
}

func TestGetErrorMiss(t *testing.T) {
	cache := newTestCache()

	cachetest.CacheTestGetErrorMiss(t, cache)
}

func newTestCache() *MemoryCache {
	return New(20 * 1024 * 1024)
}
