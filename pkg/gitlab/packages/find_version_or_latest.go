package packages

import (
	"github.com/xanzy/go-gitlab"
)

func FindVersionOrLatest(
	v []*gitlab.Package,
	version string,
) *gitlab.Package {
	if len(v) == 0 {
		return nil
	}

	for _, element := range v {
		if element.Version == version {
			return element
		}
	}

	latest := Latest(v)

	if len(latest) == 1 {
		return latest[v[0].Name]
	}

	return nil
}
