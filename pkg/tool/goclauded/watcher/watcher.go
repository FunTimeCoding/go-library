package watcher

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service"
)

type Watcher struct {
	service  *service.Service
	recovery *recovery.Recovery
	logger   *logger.Logger
	harbor   string
	changed  map[string]string
}
