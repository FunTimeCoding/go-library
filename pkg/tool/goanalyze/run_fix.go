package goanalyze

import "go/token"

func runFix(patterns []string, diff bool) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	fileSet := token.NewFileSet()
	all := load(fileSet, "", patterns)
	violations := findViolations(all)

	if len(violations) == 0 {
		return
	}

	edits := buildAllEdits(all, violations)
	applyEdits(fileSet, edits, "", diff)
}
