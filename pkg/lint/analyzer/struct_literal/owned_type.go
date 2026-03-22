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
	o, ok := p.TypesInfo.Uses[sel.Sel]

	if !ok {
		return nil
	}

	typeName, ok := o.(*types.TypeName)

	if !ok {
		return nil
	}

	_, isStruct := typeName.Type().Underlying().(*types.Struct)

	if !isStruct {
		return nil
	}

	pkg := typeName.Pkg()

	if pkg == nil {
		return nil
	}

	path := pkg.Path()

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
