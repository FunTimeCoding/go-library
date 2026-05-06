package runner

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/option"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/store"
)

func New(
	o *option.Ansible,
	s *store.Store,
	l *logger.Logger,
	r face.Reporter,
) *Runner {
	return &Runner{
		repository:  o.Repository,
		clonePath:   o.ClonePath,
		ansiblePath: o.AnsiblePath,
		playbook:    o.Playbook,
		logger:      l,
		recovery:    recovery.New(l, r),
		store:       s,
		trigger:     make(chan TriggerRequest, 1),
		stop:        make(chan struct{}),
	}
}
