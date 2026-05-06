package runner

import (
	"github.com/funtimecoding/go-library/pkg/system/run"
	"strings"
)

func (r *Runner) gitRevision(reference string) string {
	c := run.New()
	c.Directory = r.clonePath

	return strings.TrimSpace(c.Start("git", "rev-parse", reference))
}
