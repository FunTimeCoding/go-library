package gosentry

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/check/issue"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New("gosentry", version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.CopyableArgument()
	monitor.NotationArgument()
	monitor.AllArgument()
	monitor.VerboseArgument()
	pflag.String(
		argument.Issue,
		"",
		"Show details for a specific issue (e.g. GO-1B)",
	)
	monitor.ParseBind(version, gitHash, buildDate)

	if i := viper.GetString(argument.Issue); i != "" {
		showIssue(i)

		return
	}

	p := option.New()
	p.Notation = viper.GetBool(argument.Notation)
	p.All = viper.GetBool(argument.All)
	p.Verbose = viper.GetBool(argument.Verbose)
	p.Copyable = viper.GetBool(argument.Copyable)
	issue.Check(p)
}
