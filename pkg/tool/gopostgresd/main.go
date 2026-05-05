package gopostgresd

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/inventory"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/option"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	pflag.Int(argument.Port, web.ListenPort, web.PortUsage)
	pflag.String(
		argument.Inventory,
		"gopostgresd.yaml",
		"Inventory file path",
	)
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Port = argument.RequiredInteger(argument.Port)
	o.Inventory = inventory.Load(viper.GetString(argument.Inventory))
	o.Version = version
	Run(o, r)
}
