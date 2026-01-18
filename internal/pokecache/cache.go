package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	// Create a new cache struct
	cache := Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}
	// Start reap method on new cache before returning
	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	// Lock the c.entry map during entry addition
	// Ensure map is unlocked via defer
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}

	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// Lock the c.entry map during entry addition
	// Ensure map is unlocked via defer
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	// Create a ticker based on the interval passed to dictate when values are reaped
	tick := time.NewTicker(interval)

	for range tick.C {
		c.mu.Lock()
		for k, v := range c.entries {
			if time.Since(v.createdAt) > interval {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}
