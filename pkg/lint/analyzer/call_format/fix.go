package call_format

import (
	"go/ast"
	"go/token"
	"strings"
)

// Fix represents a text replacement for fixing call formatting.
type Fix struct {
	Position token.Pos
	End      token.Pos
	NewText  string
}

// BuildFixes returns the edits needed to fix a call_format violation.
func BuildFixes(
	fileSet *token.FileSet,
	call *ast.CallExpr,
	source []byte,
) []Fix {
	indent := callIndent(fileSet, call, source)
	argIndent := indent + "\t"

	openLine := fileSet.Position(call.Lparen).Line
	closeLine := fileSet.Position(call.Rparen).Line

	if openLine == closeLine {
		return fixSingleLine(call, argIndent, indent)
	}

	return fixMultiLine(fileSet, call, argIndent)
}

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

func fixMultiLine(
	fileSet *token.FileSet,
	call *ast.CallExpr,
	argIndent string,
) []Fix {
	var result []Fix

	openLine := fileSet.Position(call.Lparen).Line
	firstArgLine := fileSet.Position(call.Args[0].Pos()).Line

	if openLine == firstArgLine {
		result = append(
			result,
			Fix{
				Position: call.Lparen + 1,
				End:      call.Args[0].Pos(),
				NewText:  "\n" + argIndent,
			},
		)
	}

	for i := 1; i < len(call.Args); i++ {
		previousEnd := fileSet.Position(call.Args[i-1].End()).Line
		currentStart := fileSet.Position(call.Args[i].Pos()).Line

		if previousEnd == currentStart {
			result = append(
				result,
				Fix{
					Position: call.Args[i-1].End(),
					End:      call.Args[i].Pos(),
					NewText:  ",\n" + argIndent,
				},
			)
		}
	}

	return result
}

func callIndent(
	fileSet *token.FileSet,
	call *ast.CallExpr,
	source []byte,
) string {
	pos := fileSet.Position(call.Pos())
	lineStart := pos.Offset - (pos.Column - 1)
	count := 0

	for i := lineStart; i < len(source) && source[i] == '\t'; i++ {
		count++
	}

	return strings.Repeat("\t", count)
}
