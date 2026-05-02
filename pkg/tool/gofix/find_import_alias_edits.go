package gofix

import (
	"fmt"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strings"
)

func findImportAliasEdits(
	fileSet *token.FileSet,
	all []*packages.Package,
	r *results,
) []edit {
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
				filePath := fileSet.File(spec.Pos()).Name()

				if alias == declaredName {
					r.Add(
						filePath,
						fmt.Sprintf("de-aliased %s (redundant)", alias),
					)
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
					r.AddBlocked(
						filePath,
						fmt.Sprintf(
							"cannot de-alias %s → %s (local collision with %q)",
							alias,
							declaredName,
							declaredName,
						),
					)

					continue
				}

				r.Add(
					filePath,
					fmt.Sprintf("de-aliased %s → %s", alias, declaredName),
				)
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
