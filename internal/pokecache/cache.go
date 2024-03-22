package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data       map[string]cacheEntry
	mu         sync.RWMutex
	defaultTTL time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(defaultTTL time.Duration) (cache *Cache) {
	// Do something with Duration here
	cache = &Cache{
		data:       make(map[string]cacheEntry),
		defaultTTL: defaultTTL,
	}

	cache.StartCleanup(defaultTTL)

	return cache
}

func (c *Cache) Cleanup() {
	now := time.Now()
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.data {
		if now.Sub(entry.createdAt) >= c.defaultTTL {
			delete(c.data, key)
		}
	}
}

func (c *Cache) StartCleanup(internal time.Duration) {
	ticker := time.NewTicker(internal)
	go func() {
		for {
			select {
			case <-ticker.C:
				c.Cleanup()
			}
		}
	}()
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) (val []byte, found bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, found := c.data[key]
	if !found {
		return entry.val, false
	}
	return entry.val, true
}
