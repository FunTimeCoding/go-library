package variable_naming

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "variable_naming",
	Doc:  "assigns canonical single-letter names to local variables based on their type",
	Run:  run,
}
