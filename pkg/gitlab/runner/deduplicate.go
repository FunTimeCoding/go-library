package runner

import "gitlab.com/gitlab-org/api/client-go"

func Deduplicate(v []*gitlab.Runner) []*gitlab.Runner {
	var result []*gitlab.Runner
	seen := make(map[int]bool)

	for _, e := range v {
		if seen[e.ID] {
			continue
		}

		seen[e.ID] = true
		result = append(result, e)
	}

	return result
}
