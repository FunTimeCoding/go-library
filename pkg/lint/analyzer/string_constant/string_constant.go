package string_constant

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "string_constant",
	Doc:  "flags string literals used as function arguments where a constant with the same value exists nearby",
	Run:  run,
}
