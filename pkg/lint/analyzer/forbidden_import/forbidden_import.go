package forbidden_import

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "forbidden_import",
	Doc:  "flags imports of banned packages",
	Run:  run,
}
