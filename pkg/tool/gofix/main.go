package gofix

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(constant.LocatorEnvironment); c != "" {
		r := reporter.New("gofix", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	f := parseFlags()

	if f.survey {
		runSurvey(f.patterns)

		return
	}

	r := newResults()

	if f.rename {
		runVariableNamingFix(f.patterns, f.diff, &r)
	} else {
		runFix(f.patterns, f.diff, &r)
		runCallFormatFix(f.patterns, f.diff, &r)
		runImportAliasFix(f.patterns, f.diff, &r)
	}

	hasBlocked := printResults(r.Entries, f.summary)

	if hasBlocked {
		os.Exit(1)
	}
}
