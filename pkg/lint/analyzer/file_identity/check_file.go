package file_identity

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
)

func checkFile(
	p *packages.Package,
	results *output.Results,
	file *ast.File,
	name string,
) {
	typeMap := map[string]bool{}
	var identities []identity

	for _, d := range file.Decls {
		g, okay := d.(*ast.GenDecl)

		if !okay {
			continue
		}

		for _, s := range g.Specs {
			t, okay1 := s.(*ast.TypeSpec)

			if !okay1 {
				continue
			}

			typeMap[t.Name.Name] = true
			identities = append(
				identities,
				identity{name: t.Name.Name, position: t.Pos()},
			)
		}
	}

	for _, d := range file.Decls {
		f, okay := d.(*ast.FuncDecl)

		if !okay {
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

		if typeMap[receiverName] {
			continue
		}

		identities = append(
			identities,
			identity{name: f.Name.Name, position: f.Pos(), method: f},
		)
	}

	filename := p.Fset.Position(file.Pos()).Filename

	if len(identities) > 1 {
		results.AddBlocked(
			filename,
			fmt.Sprintf(
				"multiple identities in one file: %s and %s",
				identities[0].name,
				identities[1].name,
			),
		)

		return
	}

	if len(identities) != 1 {
		return
	}

	sole := identities[0]

	if sole.method != nil {
		o := p.TypesInfo.ObjectOf(sole.method.Name)

		if f, okay := o.(*types.Func); okay && isInterfaceMethod(p, f) {
			return
		}
	}

	stem := strings.TrimSuffix(name, constant.GoExtension)
	expected := toSnakeCase(sole.name)

	if stem != expected {
		results.AddBlocked(
			filename,
			fmt.Sprintf(
				"filename %s does not match identity %s (expected %s.go)",
				name,
				sole.name,
				expected,
			),
		)
	}
}
