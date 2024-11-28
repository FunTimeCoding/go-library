package job

import "github.com/xanzy/go-gitlab"

func Deduplicate(v []*gitlab.Job) []*gitlab.Job {
	var result []*gitlab.Job

	for _, element := range v {
		var found bool

		for _, existing := range result {
			if element.ID == existing.ID {
				found = true

				break
			}
		}

		if !found {
			result = append(result, element)
		}
	}

	return result
}
