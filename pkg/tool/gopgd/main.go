package gopgd

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/inventory"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/option"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	var h *sentry.Hub

	if c := environment.Optional(constant.LocatorEnvironment); c != "" {
		r := reporter.New("gopgd", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
		h = r.Hub()
	}

	pflag.Int(argument.Port, web.ListenPort, web.PortUsage)
	pflag.String(
		argument.Inventory,
		"gopgd.yaml",
		"Inventory file path",
	)
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Port = argument.RequiredInteger(argument.Port)
	o.Inventory = inventory.Load(viper.GetString(argument.Inventory))
	Run(o, h)
}
