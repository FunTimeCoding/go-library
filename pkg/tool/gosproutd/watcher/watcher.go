package watcher

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/service"
)

type Watcher struct {
	service       *service.Service
	recovery      *recovery.Recovery
	seedDirectory string
	stop          chan struct{}
}
