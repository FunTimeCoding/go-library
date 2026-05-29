package worker

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/service"
	"time"
)

type Worker struct {
	service  *service.Service
	interval time.Duration
	recovery *recovery.Recovery
	stop     chan struct{}
}
