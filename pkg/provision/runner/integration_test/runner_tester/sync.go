package runner_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/provision/runner"
)

func (o *Tester) Sync() *runner.SyncResult {
	o.t.Helper()
	result, e := o.Runner.Sync()
	assert.FatalOnError(o.t, e)

	return result
}
