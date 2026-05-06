package runner

import "github.com/funtimecoding/go-library/pkg/system/run"

func (r *Runner) gitPull() {
	c := run.New()
	c.Directory = r.clonePath
	c.Start("git", "pull", "origin", Branch)
}
