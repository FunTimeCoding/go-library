package gofix

import (
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/token"
)

func runSingleParameterFixWithDirectory(
	patterns []string,
	directory string,
	r *output.Results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	fileSet := token.NewFileSet()
	all := load(fileSet, directory, patterns)
	edits := findSingleParameterEdits(all, r)

	if len(edits) == 0 {
		return
	}

	applyEdits(fileSet, edits, directory, false)
}
