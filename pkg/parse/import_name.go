package parse

import (
	"go/ast"
	"strings"
)

func ImportName(
	f *ast.File,
	path string,
) (string, bool) {
	for _, i := range f.Imports {
		if strings.Trim(i.Path.Value, "\"") != path {
			continue
		}

		if i.Name != nil {
			return i.Name.Name, true
		}

		parts := strings.Split(path, "/")

		return parts[len(parts)-1], true
	}

	return "", false
}
