package anonymous_struct

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "anonymous_struct",
	Doc:  "flags anonymous struct declarations; extract a named type",
	Run:  run,
}
