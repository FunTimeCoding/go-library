package runner

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/store"
)

type Runner struct {
	repository  string
	clonePath   string
	ansiblePath string
	playbook    []string
	logger      *logger.Logger
	recovery    *recovery.Recovery
	store       *store.Store
	trigger     chan TriggerRequest
	stop        chan struct{}
}
