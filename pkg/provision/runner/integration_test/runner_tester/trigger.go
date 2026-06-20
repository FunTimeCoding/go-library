package runner_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/provision/runner"
)

func (o *Tester) Trigger(request runner.TriggerRequest) {
	o.t.Helper()
	assert.FatalOnError(o.t, o.Runner.Trigger(request))
}
