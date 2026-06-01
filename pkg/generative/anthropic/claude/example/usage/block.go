package usage

import (
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/example/common"
	"time"
)

type block struct {
	start   time.Time
	end     time.Time
	entries []*common.Timestamped
}
