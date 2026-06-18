package imports

import "go/ast"

func UsedBy(
	file *ast.File,
	node ast.Node,
) []*ast.ImportSpec {
	names := referencedPackages(node)

	if len(names) == 0 {
		return nil
	}

	var result []*ast.ImportSpec

	for _, spec := range file.Imports {
		local := localName(spec)

		if names[local] {
			result = append(result, spec)
		}
	}

	return result
}
