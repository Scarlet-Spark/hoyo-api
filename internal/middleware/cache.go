package middleware

import (
	"sync"
	"time"
)

// Simple in-memory Cache with TTL (Time to Live) support.
// Intended to be used to store static information e.g., HoYoWiki data.
type Cache struct {
	items      map[string]cacheItem
	defaultTTL time.Duration
	mutex      sync.Mutex
}

// Constructor.
// Cache has a default TTL of 1 minute.
func NewCache() *Cache {
	cache := &Cache{
		items:      make(map[string]cacheItem),
		defaultTTL: time.Minute,
	}
	return cache
}

// Set the default TTL of the cache.
func (c *Cache) SetCacheDuration(d time.Duration) {
	c.defaultTTL = d
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

// Adds an item to the cache using the default TTL.
// Replaces existing item if the key exists.
func (c *Cache) Set(k string, v any) {
	c.SetWithTTL(k, v, c.defaultTTL)
}

// Adds an item to the cache with a custom TTL.
// Replaces existing item if the key exists.
func (c *Cache) SetWithTTL(k string, v any, d time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items[k] = newCacheItem(v, time.Now().Add(d))
}

// Removes item from the cache.
func (c *Cache) Remove(k string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.items, k)
}
