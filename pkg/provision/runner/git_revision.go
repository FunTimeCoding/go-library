package runner

import (
	"github.com/funtimecoding/go-library/pkg/system/run"
	"strings"
)

func (r *Runner) gitRevision(reference string) string {
	c := run.New()
	c.Directory = r.clonePath
	c.Start("git", "rev-parse", reference)

	return strings.TrimSpace(c.OutputString)
}
