package runner

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/store"
)

func New(
	o *option.Terraform,
	s *store.Store,
	l *logger.Logger,
	r face.Reporter,
) *Runner {
	return &Runner{
		repository:    o.Repository,
		clonePath:     o.ClonePath,
		terraformPath: o.TerraformPath,
		logger:        l,
		recovery:      recovery.New(l, r),
		store:         s,
		trigger:       make(chan TriggerRequest, 1),
		stop:          make(chan struct{}),
	}
}
