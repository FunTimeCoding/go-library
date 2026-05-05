package goraidd

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/relational/postgres"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/option"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/spf13/pflag"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	pflag.Int(argument.Port, web.ListenPort, web.PortUsage)
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.Port = argument.RequiredInteger(argument.Port)
	o.PostgresLocator = environment.Required(postgres.LocatorEnvironment)
	o.LogCachePath = "/srv/arcdps-config"
	o.ElitePath = "/srv/elite-insights"
	o.OutputPath = "/srv/gw2-report"
	Run(o, r)
}
