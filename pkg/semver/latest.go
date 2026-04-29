package semver

import "golang.org/x/mod/semver"

func Latest(tags []string) string {
	filtered := Filter(tags)

	if len(filtered) == 0 {
		return ""
	}

	semver.Sort(filtered)

	return filtered[len(filtered)-1]
}
