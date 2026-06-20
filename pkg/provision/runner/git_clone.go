package runner

import (
	"github.com/funtimecoding/go-library/pkg/system/run"
	"os"
)

func (r *Runner) gitClone() {
	if _, e := os.Stat(r.clonePath); e == nil {
		r.gitPull()

		return
	}

	r.logger.Structured("git_clone", "repository", r.repository)
	run.New().Start("git", "clone", r.repository, r.clonePath)
}
