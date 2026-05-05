package gomcp

import (
	"github.com/funtimecoding/go-library/pkg/argument"
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
	r := reporter.New("gomcp", version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	pflag.String(argument.Token, "", "Bearer token for authorization")
	monitor.ParseBind(version, gitHash, buildDate)
	probe(
		argument.RequiredPositional(0, "URL"),
		viper.GetString(argument.Token),
	)
}
