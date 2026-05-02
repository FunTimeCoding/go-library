package gofix

import "go/token"

func runImportAliasFix(
	patterns []string,
	diff bool,
) {
	runImportAliasFixWithDirectory(patterns, "", diff)
}

func runImportAliasFixWithDirectory(
	patterns []string,
	directory string,
	diff bool,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	fileSet := token.NewFileSet()
	all := load(fileSet, directory, patterns)
	edits := findImportAliasEdits(all)

	if len(edits) == 0 {
		return
	}

	applyEdits(fileSet, edits, directory, diff)
}
