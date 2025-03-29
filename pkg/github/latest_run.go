package github

import "github.com/funtimecoding/go-library/pkg/github/run"

func LatestRun(v []*run.Run) *run.Run {
	latest := v[0]

	for _, element := range v {
		if element.Status == run.Completed {
			if element.CreatedAt.After(latest.CreatedAt) {
				latest = element
			}
		}
	}

	return latest
}
