package runner

import "github.com/funtimecoding/go-library/pkg/provision/salt"

func (r *Runner) SaltClient() *salt.Client {
	return r.salt
}
