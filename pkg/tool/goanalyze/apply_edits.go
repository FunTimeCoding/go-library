package goanalyze

import (
	"fmt"
	"go/token"
	"os"
	"sort"
)

func applyEdits(
	fileSet *token.FileSet,
	edits []edit,
	directory string,
	diff bool,
) {
	grouped := groupByFile(fileSet, edits, directory)

	for path, fileEdits := range grouped {
		sort.Slice(
			fileEdits, func(i, j int) bool {
				return fileEdits[i].offset > fileEdits[j].offset
			},
		)
		original, e := os.ReadFile(path)

		if e != nil {
			fmt.Fprintf(os.Stderr, "read %s: %s\n", path, e)

			continue
		}

		modified := make([]byte, len(original))
		copy(modified, original)

		for _, fe := range fileEdits {
			modified = splice(
				modified,
				fe.offset,
				fe.length,
				[]byte(fe.newText),
			)
		}

		if diff {
			printDiff(path, original, modified)
		} else {
			e = os.WriteFile(path, modified, 0644)

			if e != nil {
				fmt.Fprintf(os.Stderr, "write %s: %s\n", path, e)
			}
		}
	}
}
