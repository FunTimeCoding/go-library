package gofix

import (
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/token"
)

func runCallFormatFixWithDirectory(
	patterns []string,
	directory string,
	r *output.Results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	fileSet := token.NewFileSet()
	all := load(fileSet, directory, patterns)
	edits := findCallFormatEdits(all, r)

	if len(edits) == 0 {
		return
	}

	applyEdits(fileSet, edits, directory, false)
}
