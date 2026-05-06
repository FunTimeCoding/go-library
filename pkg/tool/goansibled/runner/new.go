package runner

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/option"
)

func New(
	o *option.Ansible,
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
		stop:        make(chan struct{}),
	}
}
