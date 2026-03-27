package go_mod

import "strings"

func UpdateDirectDependencies(
	exclusiveMatches []string,
	continueOnError bool,
) {
	hasExclusive := len(exclusiveMatches) > 0

	for _, d := range Read().Require {
		if d.Indirect {
			continue
		}

		if hasExclusive {
			var found bool

			for _, m := range exclusiveMatches {
				if strings.Contains(d.Mod.Path, m) {
					found = true

					break
				}
			}

			if !found {
				continue
			}
		}

		Update(d.Mod.Path, hasExclusive, continueOnError)
	}
}
