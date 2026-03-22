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
	"req":  {"r", "request"},
	"doc":  {"d", "document"},
	"pos":  {"p", "position"},
	"buf":  {"b", "buffer"},
	"ptr":  {"p", "pointer"},
	"addr": {"a", "address"},
	"ref":  {"r", "reference"},
	"nav":  {"n", "navigation"},
	"prev": {"past"},
	"decl": {"declaration"},
	"yaml": {"m", "markup"},
	"xml":  {"m", "markup"},
	"html": {"m", "markup"},
	"json": {"j", "notation"},
	"config": {"c", "configuration"},
	"cfg":    {"c", "configuration"},
	"llm": {"m", "model"},
	"tmp":    {"t"}, // temporary is too long; t suffices
	"href":   {"link", "reference", "locator"},
	"def":    {"definition"},
	"concat": {"concatenate"},
	"obj":    {"o", "object"},
	"stmt":   {"s", "statement"},
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
