package forbidden_call

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var banned = map[string]bool{
	"Command":        true,
	"CommandContext": true,
}

var Analyzer = &analysis.Analyzer{
	Name:     "forbidden_call",
	Doc:      "flags direct calls to exec.Command and exec.CommandContext; use pkg/system/run instead",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}
