package gofix

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/gofix/constant"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean("diff", false, "Show diff without applying")
	a.Boolean("survey", false, "Print violations without fixing")
	a.Boolean("rename", false, "Variable letter rename mode")
	a.Boolean("summary", false, "One line per modified file")
	a.Parse(version, gitHash, buildDate)
	patterns := a.Positionals()

	if a.GetBoolean("survey") {
		runSurvey(patterns)

		return
	}

	diff := a.GetBoolean("diff")
	s := newResults()

	if a.GetBoolean("rename") {
		runVariableNamingFix(patterns, diff, &s)
	} else {
		runFix(patterns, diff, &s)
		runCallFormatFix(patterns, diff, &s)
		runImportAliasFix(patterns, diff, &s)
	}

	hasBlocked := printResults(s.Entries, a.GetBoolean("summary"))

	if hasBlocked {
		os.Exit(1)
	}
}
