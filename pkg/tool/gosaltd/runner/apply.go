package runner

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gosaltd/store"
	"time"
)

func (r *Runner) apply(triggerSource string) {
	head := r.currentHead()
	record := store.NewHighstateRun()
	record.TriggerSource = triggerSource
	record.Status = store.StatusRunning
	record.GitHead = head
	r.store.Create(record)
	r.logger.Structured("highstate_start")
	start := time.Now()
	result := r.salt.Highstate(AllMinions)
	record.DurationMillisecond = time.Since(start).Milliseconds()
	output, e := json.MarshalIndent(result, "", "  ")
	errors.PanicOnError(e)
	record.Output = string(output)
	record.Status = store.StatusSuccess
	r.logger.Structured("highstate_done")
	r.store.Update(record)
}
