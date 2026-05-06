package runner

import (
	"github.com/funtimecoding/go-library/pkg/system/run"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/store"
	"path/filepath"
	"time"
)

func (r *Runner) apply(
	playbooks []string,
	triggerSource string,
) {
	directory := filepath.Join(r.clonePath, r.ansiblePath)
	head := r.currentHead()

	for _, p := range playbooks {
		record := store.NewPlaybookRun()
		record.Playbook = p
		record.TriggerSource = triggerSource
		record.Status = store.StatusRunning
		record.GitHead = head
		r.store.Create(record)
		r.logger.Structured("playbook_start", constant.Playbook, p)
		start := time.Now()
		c := run.New().NoPanic()
		c.Directory = directory
		c.Start("ansible-playbook", p)
		record.DurationMillisecond = time.Since(start).Milliseconds()
		record.Output = c.OutputString
		record.ErrorOutput = c.ErrorString

		if c.Error != nil {
			record.Status = store.StatusError
			r.logger.Structured(
				"playbook_error",
				constant.Playbook,
				p,
				"error",
				c.Error.Error(),
			)
		} else {
			record.Status = store.StatusSuccess
			r.logger.Structured("playbook_done", constant.Playbook, p)
		}

		r.store.Update(record)
	}
}
