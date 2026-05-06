package runner

import "github.com/funtimecoding/go-library/pkg/system/run"

func (r *Runner) gitFetch() {
	c := run.New()
	c.Directory = r.clonePath
	c.Start("git", "fetch", "origin")
}
