package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) error {
	// TODO
	return nil
}

func (c Cache) Add(key string, val []byte) {

}

func (c Cache) Get(key string) ([]byte, bool) {

}

func (c Cache) reapLoop() {

}
