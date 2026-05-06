package runner

import (
	"github.com/funtimecoding/go-library/pkg/system/run"
	"path/filepath"
)

func (r *Runner) terraformInit() {
	directory := filepath.Join(r.clonePath, r.terraformPath)
	r.logger.Structured("terraform_init")
	c := run.New()
	c.Directory = directory
	c.Start("terraform", "init")
}
