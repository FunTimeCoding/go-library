package build_tag

import "go/build/constraint"

func collect(
	x constraint.Expr,
	result *[]string,
) {
	switch v := x.(type) {
	case *constraint.TagExpr:
		if isCustom(v.Tag) {
			*result = append(*result, v.Tag)
		}
	case *constraint.NotExpr:
		return
	case *constraint.AndExpr:
		collect(v.X, result)
		collect(v.Y, result)
	case *constraint.OrExpr:
		collect(v.X, result)
		collect(v.Y, result)
	}
}
