package packages

import "gitlab.com/gitlab-org/api/client-go"

func FindVersionOrLatest(
	v []*gitlab.Package,
	version string,
) *gitlab.Package {
	if len(v) == 0 {
		return nil
	}

	for _, e := range v {
		if e.Version == version {
			return e
		}
	}

	latest := Latest(v)

	if len(latest) == 1 {
		return latest[v[0].Name]
	}

	return nil
}
