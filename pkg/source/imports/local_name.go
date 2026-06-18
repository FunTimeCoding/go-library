package imports

import (
	"go/ast"
	"path"
	"strings"
)

func localName(spec *ast.ImportSpec) string {
	if spec.Name != nil {
		return spec.Name.Name
	}

	importPath := strings.Trim(spec.Path.Value, "\"")

	return path.Base(importPath)
}
