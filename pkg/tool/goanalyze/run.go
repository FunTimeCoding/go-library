package goanalyze

import (
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"github.com/funtimecoding/go-library/pkg/tool/goanalyze/option"
	"go/token"
	"os"
)

func Run(o *option.Analyze) {
	patterns := o.Patterns

	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	fileSet := token.NewFileSet()
	loaded := load(fileSet, patterns)
	results := output.NewResults()

	for _, p := range loaded {
		check(p, &results)
	}

	hasBlocked := output.PrintResults(results.Entries, o.Summary)

	if hasBlocked {
		os.Exit(1)
	}
}
