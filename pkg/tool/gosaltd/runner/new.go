package runner

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/provision/salt"
	"github.com/funtimecoding/go-library/pkg/tool/gosaltd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gosaltd/store"
)

func New(
	o *option.Salt,
	c *salt.Client,
	s *store.Store,
	l *logger.Logger,
	r face.Reporter,
) *Runner {
	return &Runner{
		repository: o.Repository,
		clonePath:  o.ClonePath,
		salt:       c,
		logger:     l,
		recovery:   recovery.New(l, r),
		store:      s,
		trigger:    make(chan TriggerRequest, 1),
		stop:       make(chan struct{}),
	}
}
