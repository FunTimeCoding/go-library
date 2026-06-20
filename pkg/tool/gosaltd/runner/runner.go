package runner

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/provision/runner"
	"github.com/funtimecoding/go-library/pkg/provision/salt"
	"github.com/funtimecoding/go-library/pkg/provision/store"
)

type Runner struct {
	provision     *runner.Runner
	store         *store.Store
	saltConnector func() *salt.Client
	salt          *salt.Client
	connected     chan struct{}
	logger        *logger.Logger
	recovery      *recovery.Recovery
}
