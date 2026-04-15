package goanalyze

import "go/token"

func runCallFormatFixWithDirectory(
	patterns []string,
	directory string,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	fileSet := token.NewFileSet()
	all := load(fileSet, directory, patterns)
	edits := findCallFormatEdits(all)

	if len(edits) == 0 {
		return
	}

	applyEdits(fileSet, edits, directory, false)
}
