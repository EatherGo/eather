package eather

import (
	"sync"
	"time"

	"github.com/allegro/bigcache"
)

var (
	cache     *bigcache.BigCache
	oncecache sync.Once
)

func loadCache() *bigcache.BigCache {
	var cacheLoad, _ = bigcache.NewBigCache(bigcache.DefaultConfig(1 * time.Minute))

	return cacheLoad
}

// GetCache to get cache instance
func GetCache() *bigcache.BigCache {

	oncecache.Do(func() {
		cache = loadCache()
	})

	return cache
}
