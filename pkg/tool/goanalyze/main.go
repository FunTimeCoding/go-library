package goanalyze

import (
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/forbidden_call"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/naming"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/string_concatenation"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/struct_literal"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"golang.org/x/tools/go/analysis/multichecker"
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

	f := parseFlags()

	if f.survey {
		runSurvey(f.patterns)

		return
	}

	if f.fix || f.diff {
		runFix(f.patterns, f.diff)

		return
	}

	multichecker.Main(
		naming.Analyzer,
		forbidden_call.Analyzer,
		string_concatenation.Analyzer,
		struct_literal.Analyzer,
	)
}
