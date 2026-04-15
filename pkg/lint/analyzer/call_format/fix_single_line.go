package call_format

import "go/ast"

func fixSingleLine(
	call *ast.CallExpr,
	argIndent string,
	indent string,
) []Fix {
	var result []Fix
	result = append(
		result,
		Fix{
			Position: call.Lparen + 1,
			End:      call.Args[0].Pos(),
			NewText:  "\n" + argIndent,
		},
	)

	for i := 1; i < len(call.Args); i++ {
		result = append(
			result,
			Fix{
				Position: call.Args[i-1].End(),
				End:      call.Args[i].Pos(),
				NewText:  ",\n" + argIndent,
			},
		)
	}

	result = append(
		result,
		Fix{
			Position: call.Args[len(call.Args)-1].End(),
			End:      call.Rparen,
			NewText:  ",\n" + indent,
		},
	)

	return result
}
