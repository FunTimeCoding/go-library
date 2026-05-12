package gofirefoxd

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxd/option"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Integer(argument.Port, web.ListenPort, web.PortUsage)
	a.Integer(
		constant.BridgePortFlag,
		6125,
		"WebSocket bridge port for extension",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Port = a.RequiredInteger(argument.Port)
	o.BridgePort = a.RequiredInteger(constant.BridgePortFlag)
	o.Version = version
	Run(o, r)
}
