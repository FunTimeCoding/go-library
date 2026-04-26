package gofix

import "go/token"

func runVariableNamingFixWithDirectory(
	patterns []string,
	directory string,
	diff bool,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	fileSet := token.NewFileSet()
	all := load(fileSet, directory, patterns)
	edits := findVariableNamingEdits(all)

	if len(edits) == 0 {
		return
	}

	applyEdits(fileSet, edits, directory, diff)
}
