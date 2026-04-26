package gofix

import "golang.org/x/tools/go/packages"

func buildGeneratedSet(p *packages.Package) map[string]bool {
	result := make(map[string]bool)

	for _, file := range p.Syntax {
		if isGenerated(file) {
			result[p.Fset.File(file.Pos()).Name()] = true
		}
	}

	return result
}
