package goanalyze

import (
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/forbidden_call"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/naming"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/string_concatenation"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/struct_literal"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"go/token"
	"golang.org/x/tools/go/analysis/multichecker"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("goanalyze", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	fix, diff, patterns := parseFlags()

	if fix || diff {
		runFix(patterns, diff)

		return
	}

	multichecker.Main(
		naming.Analyzer,
		forbidden_call.Analyzer,
		string_concatenation.Analyzer,
		struct_literal.Analyzer,
	)
}

func parseFlags() (bool, bool, []string) {
	var fix, diff bool
	var patterns []string

	for _, a := range os.Args[1:] {
		switch a {
		case "--fix":
			fix = true
		case "--diff":
			diff = true
		default:
			patterns = append(patterns, a)
		}
	}

	return fix, diff, patterns
}

func runFix(patterns []string, diff bool) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	fileSet := token.NewFileSet()
	all := load(fileSet, "", patterns)
	violations := findViolations(all)

	if len(violations) == 0 {
		return
	}

	edits := buildAllEdits(all, violations)
	applyEdits(fileSet, edits, "", diff)
}
