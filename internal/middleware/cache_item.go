package middleware

import "time"

// A cache item storing any data and a TTL (Time-to-Live) duration.
type cacheItem struct {
	data any
	ttl  time.Time
}

// Constructor.
func newCacheItem(data any, ttl time.Time) cacheItem {
	return cacheItem{
		data: data,
		ttl:  ttl,
	}
}

// Returns true if the item has expired.
func (i cacheItem) expired() bool {
	return time.Now().After(i.ttl)
}
