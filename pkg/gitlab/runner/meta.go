package runner

import "github.com/funtimecoding/go-library/pkg/console/description"

func (r *Runner) Meta() *description.Description {
	return description.New("Runner", "Runner")
}
