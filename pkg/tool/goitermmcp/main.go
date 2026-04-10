package goitermmcp

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("goitermmcp", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	pflag.Int(argument.Port, 8080, "Listen port")
	monitor.ParseBind(version, gitHash, buildDate)
	Run(fmt.Sprintf(":%d", viper.GetInt(argument.Port)))
}
