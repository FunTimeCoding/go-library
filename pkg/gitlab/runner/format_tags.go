package runner

import "github.com/funtimecoding/go-library/pkg/strings/join"

func (r *Runner) formatTags() string {
	if len(r.Tags) == 0 {
		return ""
	}

	return join.Comma(r.Tags)
}
