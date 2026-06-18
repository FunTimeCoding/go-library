package service

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"os"
)

func writeFile(
	fileSet *token.FileSet,
	file *ast.File,
	path string,
) error {
	var buffer bytes.Buffer
	e := format.Node(&buffer, fileSet, file)

	if e != nil {
		return e
	}

	return os.WriteFile(path, buffer.Bytes(), 0644)
}
