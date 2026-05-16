package value_return

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"strings"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	if p.Module == nil {
		return
	}

	module := p.Module.Path

	for _, file := range p.Syntax {
		name := p.Fset.File(file.Pos()).Name()

		if filepath.Base(name) == constant.GeneratedFile {
			continue
		}

		if strings.HasSuffix(name, "_test.go") {
			continue
		}

		ast.Inspect(
			file,
			func(n ast.Node) bool {
				f, okay := n.(*ast.FuncDecl)

				if !okay || f.Type.Results == nil {
					return true
				}

				for _, result := range f.Type.Results.List {
					tv := p.TypesInfo.TypeOf(result.Type)

					if tv == nil {
						continue
					}

					named, okay := tv.(*types.Named)

					if !okay {
						continue
					}

					_, isStruct := named.Underlying().(*types.Struct)

					if !isStruct {
						continue
					}

					a := named.Obj().Pkg()

					if a == nil {
						continue
					}

					path := a.Path()

					if path != module && !strings.HasPrefix(
						path,
						fmt.Sprintf("%s/", module),
					) {
						continue
					}

					results.AddBlocked(
						p.Fset.Position(f.Pos()).Filename,
						fmt.Sprintf(
							"%s returns %s.%s by value",
							f.Name.Name,
							a.Name(),
							named.Obj().Name(),
						),
					)
				}

				return true
			},
		)
	}
}
