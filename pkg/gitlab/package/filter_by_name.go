package gitlab

import "github.com/xanzy/go-gitlab"

func FilterByName(
	v []*gitlab.Package,
	name string,
) []*gitlab.Package {
	var result []*gitlab.Package

	for _, element := range v {
		if element.Name == name {
			result = append(result, element)
		}
	}

	return result
}
