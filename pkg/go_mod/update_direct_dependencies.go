package go_mod

import (
	"fmt"
	"strings"
)

func UpdateDirectDependencies(exclusiveMatches []string) {
	fmt.Printf("Exclusive matches: %+v\n", exclusiveMatches)

	skipProxy := len(exclusiveMatches) > 0

	for _, dep := range Read().Require {
		if dep.Indirect {
			continue
		}

		if len(exclusiveMatches) > 0 {
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

		Update(dep.Mod.Path, skipProxy)
	}

	Tidy()
}
