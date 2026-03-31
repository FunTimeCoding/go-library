package goanalyze

import (
	"go/ast"
	"os"
	"path/filepath"
	"strings"
)

func fixUnloadedReferences(
	violations []violation,
	loadedFiles map[string]bool,
	directory string,
) {
	var renames []exportedRename

	for _, v := range violations {
		if v.fix == "" {
			continue
		}

		if !ast.IsExported(v.ident.Name) {
			continue
		}

		renames = append(
			renames, exportedRename{
				oldName: v.ident.Name,
				newName: replaceSegment(v.ident.Name, v.segment, v.fix),
			},
		)
	}

	if len(renames) == 0 {
		return
	}

	root := directory

	if root == "" {
		root = "."
	}

	_ = filepath.Walk(
		root, func(path string, info os.FileInfo, e error) error {
			if e != nil {
				return nil
			}

			if info.IsDir() {
				return nil
			}

			if !strings.HasSuffix(path, ".go") {
				return nil
			}

			if loadedFiles[path] {
				return nil
			}

			fixFileReferences(path, renames)

			return nil
		},
	)
}

type exportedRename struct {
	oldName string
	newName string
}
