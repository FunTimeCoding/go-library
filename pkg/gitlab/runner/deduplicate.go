package runner

import "github.com/xanzy/go-gitlab"

func Deduplicate(v []*gitlab.Runner) []*gitlab.Runner {
	var result []*gitlab.Runner
	seen := make(map[int]bool)

	for _, element := range v {
		if seen[element.ID] {
			continue
		}

		seen[element.ID] = true
		result = append(result, element)
	}

	return result
}
