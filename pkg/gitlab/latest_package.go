package gitlab

import "github.com/xanzy/go-gitlab"

func LatestPackages(v []*gitlab.Package) map[string]*gitlab.Package {
	result := make(map[string]*gitlab.Package)

	for _, p := range v {
		existing := result[p.Name]

		if existing == nil || p.Version > existing.Version {
			result[p.Name] = p
		}
	}

	return result
}
