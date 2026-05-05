package gofix

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New("gofix", version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	f := parseFlags()

	if f.survey {
		runSurvey(f.patterns)

		return
	}

	s := newResults()

	if f.rename {
		runVariableNamingFix(f.patterns, f.diff, &s)
	} else {
		runFix(f.patterns, f.diff, &s)
		runCallFormatFix(f.patterns, f.diff, &s)
		runImportAliasFix(f.patterns, f.diff, &s)
	}

	hasBlocked := printResults(s.Entries, f.summary)

	if hasBlocked {
		os.Exit(1)
	}
}
