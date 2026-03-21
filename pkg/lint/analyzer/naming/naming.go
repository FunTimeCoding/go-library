package naming

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

// suggestions maps a banned name segment to acceptable alternatives.
// Single-character alternatives apply only in variable contexts
// (local variables, parameters, struct fields).
var suggestions = map[string][]string{
	"url":   {"l", "locator"},
	"mcp":   {"c", "model_context"},
	"dir":   {"d", "directory"},
	"dirs":  {"directories"},
	"tx":    {"t"}, // transaction is too long; t suffices
	"ctx":   {"x"}, // context collides with the standard library package name
	"param": {"parameter"},
	"msg":   {"m", "message"},
	"req":   {"r", "request"},
	"doc":   {"d", "document"},
	"config": {
		"c",
		"configuration",
	}, // configuration is long but beats a vague abbreviation in non-variable contexts; often replaceable with something more specific like setting
	"cfg": {"c", "configuration"},
	"llm": {"m", "model"},
	"tmp": {"t"}, // temporary is too long; t suffices
}

// noSuggestion contains banned noise words with no prescribed replacement.
// They are flagged regardless of identifier kind.
var noSuggestion = map[string]bool{
	"handler": true, // too vague; use a name that describes what it handles
	"data":    true, // too vague; everything is data
	"info":    true, // too vague; everything is info
}

var Analyzer = &analysis.Analyzer{
	Name:     "naming",
	Doc:      "flags blacklisted identifier name segments",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}
