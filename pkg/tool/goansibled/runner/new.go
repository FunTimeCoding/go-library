package runner

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/provision/runner"
	"github.com/funtimecoding/go-library/pkg/provision/store"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/option"
)

func New(
	o *option.Ansible,
	s *store.Store,
	l *logger.Logger,
	r face.Reporter,
) *Runner {
	result := &Runner{
		store:       s,
		playbook:    o.Playbook,
		clonePath:   o.ClonePath,
		ansiblePath: o.AnsiblePath,
		logger:      l,
	}
	result.provision = runner.New(
		runner.Configuration{
			Repository:      o.Repository,
			ClonePath:       o.ClonePath,
			ToolPath:        o.AnsiblePath,
			ApplyFunction:   result.apply,
			CleanupFunction: s.Cleanup,
		},
		l,
		r,
	)

	return result
}
