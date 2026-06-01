package common

import (
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/usage_entry"
	"time"
)

type Timestamped struct {
	Time  time.Time
	Entry *usage_entry.Entry
}
