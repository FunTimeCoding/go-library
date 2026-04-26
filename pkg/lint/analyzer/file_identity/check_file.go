package file_identity

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"go/ast"
	goTypes "go/types"
	"golang.org/x/tools/go/analysis"
	"strings"
)

func checkFile(
	p *analysis.Pass,
	file *ast.File,
	name string,
) {
	types := map[string]bool{}
	var identities []identity

	for _, d := range file.Decls {
		g, ok := d.(*ast.GenDecl)

		if !ok {
			continue
		}

		for _, s := range g.Specs {
			t, ok := s.(*ast.TypeSpec)

			if !ok {
				continue
			}

			types[t.Name.Name] = true
			identities = append(
				identities,
				identity{name: t.Name.Name, position: t.Pos()},
			)
		}
	}

	for _, d := range file.Decls {
		f, ok := d.(*ast.FuncDecl)

		if !ok {
			continue
		}

		if f.Recv == nil {
			identities = append(
				identities,
				identity{name: f.Name.Name, position: f.Pos()},
			)

			continue
		}

		receiverName := receiverTypeName(f.Recv.List[0].Type)

		if types[receiverName] {
			continue
		}

		identities = append(
			identities,
			identity{name: f.Name.Name, position: f.Pos(), method: f},
		)
	}

	if len(identities) > 1 {
		p.Reportf(
			identities[1].position,
			"multiple identities in one file: %s and %s",
			identities[0].name,
			identities[1].name,
		)

		return
	}

	if len(identities) != 1 {
		return
	}

	sole := identities[0]

	if sole.method != nil {
		o := p.TypesInfo.ObjectOf(sole.method.Name)

		if f, ok := o.(*goTypes.Func); ok && isInterfaceMethod(p, f) {
			return
		}
	}

	stem := strings.TrimSuffix(name, constant.GoExtension)
	expected := toSnakeCase(sole.name)

	if stem != expected {
		p.Reportf(
			sole.position,
			"filename %s does not match identity %s (expected %s.go)",
			name,
			sole.name,
			expected,
		)
	}
}
