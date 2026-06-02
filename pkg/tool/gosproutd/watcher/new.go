package watcher

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/service"
)

func New(
	s *service.Service,
	l *logger.Logger,
	r face.Reporter,
	seedDirectory string,
) *Watcher {
	return &Watcher{
		service:       s,
		recovery:      recovery.New(l, r),
		seedDirectory: seedDirectory,
		stop:          make(chan struct{}),
	}
}
