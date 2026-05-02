package gofix

import "go/token"

func runFix(
	patterns []string,
	diff bool,
	r *results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	fileSet := token.NewFileSet()
	all := load(fileSet, "", patterns)
	violations := findViolations(all)

	if len(violations) == 0 {
		return
	}

	edits := buildAllEdits(fileSet, all, violations, r)
	applyEdits(fileSet, edits, "", diff)

	if !diff {
		loadedFiles := buildLoadedFiles(all)
		fixUnloadedReferences(violations, loadedFiles, "", r)
	}
}
