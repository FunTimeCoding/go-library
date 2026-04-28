package ticker

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"time"
)

type Ticker struct {
	interval time.Duration
	function func()
	recovery *recovery.Recovery
	ticker   *time.Ticker
	done     chan struct{}
}
