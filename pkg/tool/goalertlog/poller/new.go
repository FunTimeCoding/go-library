package poller

import (
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/store"
	"time"
)

func New(
	a AlertSource,
	s *store.Store,
	l *logger.Logger,
	interval time.Duration,
	retention time.Duration,
) *Poller {
	return &Poller{
		client:    a,
		store:     s,
		logger:    l,
		interval:  interval,
		retention: retention,
		firing:    make(map[string]string),
		stop:      make(chan struct{}),
	}
}
