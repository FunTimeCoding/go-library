package search_cache

import (
	"sync"
	"time"
)

type Cache struct {
	mu      sync.Mutex
	entries map[string]*entry
	order   []string
	maxSize int
	ttl     time.Duration
}
