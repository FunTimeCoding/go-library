package gofirefoxmcp

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	sentryConstant "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/option"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/spf13/pflag"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(
		"gofirefoxmcp",
		environment.Optional(sentryConstant.LocatorEnvironment),
		"",
		version,
	)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()

	if c := environment.Optional(sentryConstant.LocatorEnvironment); c != "" {
		r := reporter.New("gofirefoxmcp", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	pflag.Int(argument.Port, web.ListenPort, web.PortUsage)
	pflag.Int(
		constant.BridgePortFlag,
		6125,
		"WebSocket bridge port for extension",
	)
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Port = argument.RequiredInteger(argument.Port)
	o.BridgePort = argument.RequiredInteger(constant.BridgePortFlag)
	Run(o, r)
}
