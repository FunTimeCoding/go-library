package runner

import (
	"github.com/funtimecoding/go-library/pkg/system/run"
	"path/filepath"
)

func (r *Runner) terraformInit() {
	directory := filepath.Join(r.clonePath, r.terraformPath)
	r.logger.Structured("terraform_init")
	c := run.New().NoPanic()
	c.Directory = directory
	c.Start("terraform", "init", "-json")

	if c.Error != nil && r.needsUpgrade(c.OutputString) {
		r.logger.Structured("terraform_init_upgrade")
		u := run.New()
		u.Directory = directory
		u.SetReporter(r.reporter, "terraform init -upgrade")
		u.Start("terraform", "init", "-upgrade")

		return
	}

	if c.Error != nil {
		r.reporter.CaptureWithContext(
			c.Error,
			"terraform init",
			map[string]any{
			"output": c.OutputString,
			"stderr": c.ErrorString,
		})
		panic(c.Error)
	}
}
