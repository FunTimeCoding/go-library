package gofirefoxmcp

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("gofirefoxmcp", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	pflag.Int(argument.Port, 8080, "MCP listen port")
	pflag.Int(constant.BridgePortFlag, 6125, "WebSocket bridge port for extension")
	monitor.ParseBind(version, gitHash, buildDate)
	Run(
		fmt.Sprintf(":%d", viper.GetInt(argument.Port)),
		viper.GetInt(constant.BridgePortFlag),
	)
}
