package packages

import "gitlab.com/gitlab-org/api/client-go"

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
