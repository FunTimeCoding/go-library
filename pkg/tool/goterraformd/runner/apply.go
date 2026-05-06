package runner

import (
	"github.com/funtimecoding/go-library/pkg/system/run"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/store"
	"path/filepath"
	"time"
)

func (r *Runner) apply(triggerSource string) {
	directory := filepath.Join(r.clonePath, r.terraformPath)
	head := r.currentHead()
	record := store.NewTerraformRun()
	record.TriggerSource = triggerSource
	record.Status = store.StatusRunning
	record.GitHead = head
	r.store.Create(record)
	r.logger.Structured("terraform_apply_start")
	start := time.Now()
	c := run.New().NoPanic()
	c.Directory = directory
	c.Start("terraform", "apply", "-auto-approve")
	record.DurationMillisecond = time.Since(start).Milliseconds()
	record.Output = c.OutputString
	record.ErrorOutput = c.ErrorString

	if c.Error != nil {
		record.Status = store.StatusError
		r.logger.Structured(
			"terraform_apply_error",
			"error",
			c.Error.Error(),
		)
	} else {
		record.Status = store.StatusSuccess
		r.logger.Structured("terraform_apply_done")
	}

	r.store.Update(record)
}
