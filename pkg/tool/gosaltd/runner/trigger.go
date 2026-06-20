package runner

import "github.com/funtimecoding/go-library/pkg/provision/runner"

func (r *Runner) Trigger(request runner.TriggerRequest) error {
	return r.provision.Trigger(request)
}
