package struct_literal

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"path/filepath"
	"strings"
)

func ownedType(
	p *analysis.Pass,
	sel *ast.SelectorExpr,
	module string,
) *types.TypeName {
	o, okay := p.TypesInfo.Uses[sel.Sel]

	if !okay {
		return nil
	}

	typeName, okay := o.(*types.TypeName)

	if !okay {
		return nil
	}

	_, isStruct := typeName.Type().Underlying().(*types.Struct)

	if !isStruct {
		return nil
	}

	a := typeName.Pkg()

	if a == nil {
		return nil
	}

	path := a.Path()

	if path != module && !strings.HasPrefix(
		path,
		fmt.Sprintf("%s/", module),
	) {
		return nil
	}

	if filepath.Base(
		p.Fset.Position(typeName.Pos()).Filename,
	) == constant.GeneratedFile {
		return nil
	}

	return typeName
}
