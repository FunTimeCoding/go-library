package type_receiver

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "type_receiver",
	Doc:  "flags packages with more than one type that has method receivers",
	Run:  run,
}
