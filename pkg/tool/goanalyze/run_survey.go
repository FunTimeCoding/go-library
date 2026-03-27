package goanalyze

import (
	"fmt"
	"go/token"
	"path/filepath"
	"sort"
	"strings"
)

func runSurvey(patterns []string) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	fileSet := token.NewFileSet()
	all := load(fileSet, "", patterns)
	counts := make(map[string]int)
	examples := make(map[string][]string)

	for _, p := range all {
		generatedFiles := buildGeneratedSet(p)

		for ident, o := range p.TypesInfo.Defs {
			if o == nil {
				continue
			}

			file := p.Fset.File(ident.Pos()).Name()

			if generatedFiles[file] {
				continue
			}

			if strings.HasSuffix(filepath.Base(file), "_test.go") {
				continue
			}

			for _, segment := range segments(ident.Name) {
				if whitelist[segment] {
					continue
				}

				counts[segment]++

				if len(examples[segment]) < 3 {
					examples[segment] = append(examples[segment], ident.Name)
				}
			}
		}
	}

	printSurvey(counts, examples)
}

func printSurvey(
	counts map[string]int,
	examples map[string][]string,
) {
	type entry struct {
		segment string
		count   int
	}

	var entries []entry

	for segment, count := range counts {
		entries = append(entries, entry{segment, count})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].count > entries[j].count
	})

	for _, e := range entries {
		fmt.Printf(
			"%4d  %-20s  %s\n",
			e.count,
			e.segment,
			strings.Join(examples[e.segment], ", "),
		)
	}
}
