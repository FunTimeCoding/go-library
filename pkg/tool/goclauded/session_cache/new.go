package session_cache

import "github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/tracker"

func New() *Cache {
	return &Cache{entries: make(map[string]*tracker.State)}
}
