package go_mod

import "strings"

func UpdateDirectDependencies(exclusiveMatches []string) {
	hasExclusive := len(exclusiveMatches) > 0

	for _, dep := range Read().Require {
		if dep.Indirect {
			continue
		}

		if hasExclusive {
			var found bool

			for _, element := range exclusiveMatches {
				if strings.Contains(dep.Mod.Path, element) {
					found = true

					break
				}
			}

			if !found {
				continue
			}
		}

		Update(dep.Mod.Path, hasExclusive)
	}

	Tidy()
}
