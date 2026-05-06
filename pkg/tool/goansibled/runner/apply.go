package runner

import (
	"github.com/funtimecoding/go-library/pkg/system/run"
	"path/filepath"
)

func (r *Runner) apply() {
	directory := filepath.Join(r.clonePath, r.ansiblePath)

	for _, p := range r.playbook {
		r.logger.Structured("playbook_start", "playbook", p)
		c := run.New()
		c.Directory = directory
		c.Start("ansible-playbook", p)
		r.logger.Structured("playbook_done", "playbook", p)
	}
}
