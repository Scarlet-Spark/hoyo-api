package middleware

import (
	"sync"
	"time"
)

// Simple in-memory Cache with TTL (Time to Live) support.
type Cache struct {
	items map[string]cacheItem
	mutex sync.Mutex
}

// Creates a new cache instance.
func NewCache() *Cache {
	cache := &Cache{items: make(map[string]cacheItem)}
	return cache
}

// Get an item from the cache.
// Returns the value or nil, and a bool indicating if the value is found, or if it has expired.
func (c *Cache) Get(k string) any {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if item, exists := c.items[k]; exists {
		if !item.expired() {
			return item.data
		}
		delete(c.items, k)
	}

	return nil
}

// Adds an item to the cache with its TTL.
// Replaces existing item if the key exists.
func (c *Cache) Set(k string, v any, d time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items[k] = cacheItem{
		data: v,
		ttl:  time.Now().Add(d),
	}
}

// Removes item from the cache.
func (c *Cache) Remove(k string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.items, k)
}
