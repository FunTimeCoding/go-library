package runner

import "github.com/funtimecoding/go-library/pkg/provision/runner"

func (r *Runner) Sync() (*runner.SyncResult, error) {
	return r.provision.Sync()
}
