package runner

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system/run"
)

func (r *Runner) gitDiffLog(
	old string,
	new string,
) string {
	c := run.New()
	c.Directory = r.clonePath
	c.Start(
		"git",
		"log",
		"--oneline",
		fmt.Sprintf("%s..%s", old, new),
		"--",
		r.toolPath,
	)

	return c.OutputString
}
