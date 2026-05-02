package gofix

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strings"
)

func findImportAliasEdits(all []*packages.Package) []edit {
	var result []edit
	seen := make(map[token.Pos]bool)

	for _, p := range all {
		for _, file := range p.Syntax {
			if isGenerated(file) {
				continue
			}

			for _, spec := range file.Imports {
				if spec.Name == nil {
					continue
				}

				if seen[spec.Pos()] {
					continue
				}

				seen[spec.Pos()] = true
				alias := spec.Name.Name

				if alias == "_" || alias == "." {
					continue
				}

				path := strings.Trim(spec.Path.Value, `"`)
				imported := findImportedPackage(p, path)

				if imported == nil {
					continue
				}

				declaredName := imported.Name()

				if alias == declaredName {
					result = append(
						result,
						buildImportAliasEdits(
							p,
							file,
							spec,
							alias,
							declaredName,
						)...,
					)

					continue
				}

				if hasImportCollision(file, p, path, declaredName) {
					continue
				}

				if hasLocalCollision(p, file, declaredName) {
					errors.Printf(
						"%s: cannot de-alias %s %q (local collision with %q)\n",
						p.Fset.Position(spec.Pos()),
						alias,
						path,
						declaredName,
					)

					continue
				}

				result = append(
					result,
					buildImportAliasEdits(
						p,
						file,
						spec,
						alias,
						declaredName,
					)...,
				)
			}
		}
	}

	return result
}
