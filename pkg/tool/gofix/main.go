package gofix

import (
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("gofix", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	f := parseFlags()

	if f.survey {
		runSurvey(f.patterns)

		return
	}

	if f.rename {
		runVariableNamingFix(f.patterns, false)

		return
	}

	runFix(f.patterns, f.diff)
	runCallFormatFix(f.patterns, f.diff)
}
