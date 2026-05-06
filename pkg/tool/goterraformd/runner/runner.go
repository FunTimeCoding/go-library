package runner

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/store"
)

type Runner struct {
	repository    string
	clonePath     string
	terraformPath string
	logger        *logger.Logger
	recovery      *recovery.Recovery
	store         *store.Store
	trigger       chan TriggerRequest
	stop          chan struct{}
}
