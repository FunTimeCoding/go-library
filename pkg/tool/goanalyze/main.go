package goanalyze

import (
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/naming"
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

	multichecker.Main(naming.Analyzer)
}
