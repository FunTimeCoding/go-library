package gitlab

import "github.com/xanzy/go-gitlab"

func LatestPackages(v []*gitlab.Package) []*gitlab.Package {
	var result []*gitlab.Package

	for _, p := range v {
		// Add first entry either way
		if len(result) == 0 {
			result = append(result, p)

			continue
		}

		var found bool

		// If name is found, replace
		for i, r := range result {
			if r.Name == p.Name {
				if r.Version < p.Version {
					result[i] = p
				}

				found = true

				break
			}
		}

		// If name not found, append as well
		if !found {
			result = append(result, p)
		}
	}

	return result
}
