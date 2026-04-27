package gofix

import (
	"github.com/funtimecoding/go-library/pkg/lint/segment"
	"go/token"
	"path/filepath"
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

			for _, s := range segment.Segments(ident.Name) {
				if whitelist[s] {
					continue
				}

				counts[s]++

				if len(examples[s]) < 3 {
					examples[s] = append(examples[s], ident.Name)
				}
			}
		}
	}

	printSurvey(counts, examples)
}
