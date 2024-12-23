package pokecache

import (
	"time"
	"sync"
)
type cacheEntry struct {
	createdAt time.Time
	val []byte
}
type Cache struct {
	mu *sync.Mutex
	data map[string]cacheEntry
	done chan struct{}
	interval time.Duration
}
func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)
	done := make(chan struct{})
	c :=&Cache {
		data: map[string]cacheEntry{},
		mu: &sync.Mutex{},
		done: done,
		interval: interval,
	}

	go func() {
		for {
			select {
				case <-done:
					return
				case <-ticker.C:
					c.cleanup()			
			}
		}
	}()
	return c

}

func (c *Cache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()
	keysToPurge := []string{}
	for k, cacheEntry := range c.data {
		if cacheEntry.createdAt.Before(time.Now().Add(-1 *c.interval)) {
			keysToPurge = append(keysToPurge, k)
		}
	}
	for _, k:= range keysToPurge {
		delete(c.data, k)
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		val: val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	res, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return res.val, true
}



