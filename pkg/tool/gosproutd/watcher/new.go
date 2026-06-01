package watcher

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/store"
)

func New(
	s *store.Store,
	n face.EventNotifier,
	l *logger.Logger,
	r face.Reporter,
	seedDirectory string,
) *Watcher {
	return &Watcher{
		store:         s,
		notifier:      n,
		recovery:      recovery.New(l, r),
		seedDirectory: seedDirectory,
		stop:          make(chan struct{}),
	}
}
