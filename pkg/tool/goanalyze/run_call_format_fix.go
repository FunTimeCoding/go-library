package goanalyze

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/call_format"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
)

func runCallFormatFix(
	patterns []string,
	diff bool,
) {
	runCallFormatFixWithDirectory(patterns, "")
}

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

func findCallFormatEdits(all []*packages.Package) []edit {
	var result []edit
	sourceCache := make(map[string][]byte)
	seen := make(map[token.Pos]bool)

	for _, p := range all {
		generated := make(map[string]bool)

		for _, file := range p.Syntax {
			name := p.Fset.File(file.Pos()).Name()

			if filepath.Base(name) == constant.GeneratedFile {
				generated[name] = true
			}
		}

		for _, file := range p.Syntax {
			name := p.Fset.File(file.Pos()).Name()

			if generated[name] {
				continue
			}

			source, ok := sourceCache[name]

			if !ok {
				var e error
				source, e = os.ReadFile(name)

				if e != nil {
					continue
				}

				sourceCache[name] = source
			}

			ast.Inspect(
				file,
				func(n ast.Node) bool {
					call, ok := n.(*ast.CallExpr)

					if !ok {
						return true
					}

					if seen[call.Lparen] {
						return true
					}

					if !call_format.IsViolation(p.Fset, call) {
						return true
					}

					seen[call.Lparen] = true
					fixes := call_format.BuildFixes(
						p.Fset,
						call,
						source,
					)

					for _, f := range fixes {
						result = append(
							result,
							edit{
								position: f.Position,
								end:      f.End,
								newText:  f.NewText,
							},
						)
					}

					return true
				},
			)
		}
	}

	return result
}
