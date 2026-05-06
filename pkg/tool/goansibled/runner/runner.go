package runner

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/log/logger"
)

type Runner struct {
	repository  string
	clonePath   string
	ansiblePath string
	playbook    []string
	logger      *logger.Logger
	recovery    *recovery.Recovery
	stop        chan struct{}
}
