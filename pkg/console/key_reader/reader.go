package key_reader

import (
	"sync"
	"time"
)

type Reader struct {
	handlers     map[rune]handler
	states       map[rune]*state
	mutex        sync.RWMutex
	releaseDelay time.Duration
}
