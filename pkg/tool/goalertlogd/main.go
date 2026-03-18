package goalertlogd

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("goalertlog", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	pflag.String(
		argument.Path,
		"tmp/goalertlog.db",
		"Database path",
	)
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.DatabasePath = viper.GetString(argument.Path)
	Run(o)
}
