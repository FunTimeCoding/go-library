package gofix

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"go/ast"
	"go/token"
	"go/types"
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
						buildImportAliasRemovalEdits(
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

func findImportedPackage(
	p *packages.Package,
	path string,
) *types.Package {
	for _, i := range p.Types.Imports() {
		if i.Path() == path {
			return i
		}
	}

	return nil
}

func hasImportCollision(
	file *ast.File,
	p *packages.Package,
	currentPath string,
	declaredName string,
) bool {
	count := 0

	for _, spec := range file.Imports {
		path := strings.Trim(spec.Path.Value, `"`)

		if path == currentPath {
			continue
		}

		if spec.Name != nil && spec.Name.Name != "_" && spec.Name.Name != "." {
			if spec.Name.Name == declaredName {
				return true
			}

			imported := findImportedPackage(p, path)

			if imported != nil && imported.Name() == declaredName {
				count++
			}

			continue
		}

		imported := findImportedPackage(p, path)

		if imported != nil && imported.Name() == declaredName {
			return true
		}
	}

	return count > 0
}

func hasLocalCollision(
	p *packages.Package,
	file *ast.File,
	name string,
) bool {
	for ident, o := range p.TypesInfo.Defs {
		if o == nil {
			continue
		}

		if ident.Name != name {
			continue
		}

		if !positionInFile(p.Fset, file, ident.Pos()) {
			continue
		}

		if _, okay := o.(*types.PkgName); okay {
			continue
		}

		return true
	}

	return false
}

func positionInFile(
	fileSet *token.FileSet,
	file *ast.File,
	position token.Pos,
) bool {
	f := fileSet.File(file.Pos())

	return position >= file.Pos() && int(position) <= f.Base()+f.Size()
}

func buildImportAliasRemovalEdits(
	p *packages.Package,
	file *ast.File,
	spec *ast.ImportSpec,
	alias string,
	declaredName string,
) []edit {
	return buildImportAliasEdits(p, file, spec, alias, declaredName)
}

func buildImportAliasEdits(
	p *packages.Package,
	file *ast.File,
	spec *ast.ImportSpec,
	alias string,
	declaredName string,
) []edit {
	var result []edit

	result = append(
		result,
		edit{
			position: spec.Name.Pos(),
			end:      spec.Path.Pos(),
			newText:  "",
		},
	)

	ast.Inspect(
		file,
		func(n ast.Node) bool {
			selector, okay := n.(*ast.SelectorExpr)

			if !okay {
				return true
			}

			ident, okay := selector.X.(*ast.Ident)

			if !okay {
				return true
			}

			if ident.Name != alias {
				return true
			}

			o := p.TypesInfo.ObjectOf(ident)

			if o == nil {
				return true
			}

			if _, okay := o.(*types.PkgName); !okay {
				return true
			}

			if alias != declaredName {
				result = append(
					result,
					edit{
						position: ident.Pos(),
						end:      ident.End(),
						newText:  declaredName,
					},
				)
			}

			return true
		},
	)

	return result
}
