package gofix

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func fixFileReferences(
	path string,
	renames []exportedRename,
) {
	content, e := os.ReadFile(path)

	if e != nil {
		return
	}

	set := token.NewFileSet()
	file, e := parser.ParseFile(set, path, content, parser.ParseComments)

	if e != nil {
		return
	}

	renameMap := make(map[string]string)

	for _, r := range renames {
		renameMap[r.oldName] = r.newName
	}

	selectorPositions := collectSelectorPositions(file)
	modified := make([]byte, len(content))
	copy(modified, content)
	offset := 0
	ast.Inspect(
		file,
		func(n ast.Node) bool {
			i, ok := n.(*ast.Ident)

			if !ok {
				return true
			}

			newName, exists := renameMap[i.Name]

			if !exists {
				return true
			}

			if selectorPositions[i.Pos()] {
				return true
			}

			start := set.Position(i.Pos()).Offset + offset
			end := start + len(i.Name)
			modified = append(
				modified[:start],
				append([]byte(newName), modified[end:]...)...,
			)
			offset += len(newName) - len(i.Name)
			fmt.Printf(
				"Renamed: %s → %s (unloaded: %s)\n",
				i.Name,
				newName,
				path,
			)

			return true
		},
	)

	if offset != 0 {
		errors.PanicOnError(os.WriteFile(path, modified, 0644))
	}
}
