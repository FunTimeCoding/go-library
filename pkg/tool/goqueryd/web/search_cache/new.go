package search_cache

import "time"

func New(maxSize int) *Cache {
	return &Cache{
		entries: map[string]*entry{},
		maxSize: maxSize,
		ttl:     15 * time.Minute,
	}
}
