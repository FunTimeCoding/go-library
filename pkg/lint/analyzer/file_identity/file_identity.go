package file_identity

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "file_identity",
	Doc:  "enforces one identity per file and filename matches that identity",
	Run:  run,
}
