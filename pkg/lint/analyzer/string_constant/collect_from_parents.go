package string_constant

import "path/filepath"

func collectFromParents(
	result map[string]knownConstant,
	directory string,
) {
	current := filepath.Dir(directory)
	distance := 1

	for {
		collectFromConstantDirectory(result, current, "constant", distance)

		if filepath.Base(current) == "pkg" {
			break
		}

		parent := filepath.Dir(current)

		if parent == current {
			break
		}

		current = parent
		distance++
	}
}
