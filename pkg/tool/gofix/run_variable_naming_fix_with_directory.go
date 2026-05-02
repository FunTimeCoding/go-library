package gofix

import "go/token"

func runVariableNamingFixWithDirectory(
	patterns []string,
	directory string,
	diff bool,
	r *results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	fileSet := token.NewFileSet()
	all := load(fileSet, directory, patterns)
	edits := findVariableNamingEdits(fileSet, all, r)

	if len(edits) == 0 {
		return
	}

	applyEdits(fileSet, edits, directory, diff)
}
