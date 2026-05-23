package usage

import (
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/usage_entry"
	"time"
)

type timestamped struct {
	time  time.Time
	entry *usage_entry.Entry
}
