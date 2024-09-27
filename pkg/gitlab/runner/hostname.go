package runner

import "github.com/funtimecoding/go-library/pkg/system"

func (r *Runner) hostname() string {
	if s := system.LookupFirst(r.Address); s != "" {
		return s
	}

	if r.Address == "" {
		return NoAddress
	}

	return r.Address
}
