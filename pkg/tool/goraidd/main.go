package goraidd

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/relational/postgres"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/option"
	"github.com/getsentry/sentry-go"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	var h *sentry.Hub

	if c := environment.Optional(constant.LocatorEnvironment); c != "" {
		r := reporter.New("goraidd", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
		h = r.Hub()
	}

	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.PostgresLocator = environment.Required(postgres.LocatorEnvironment)
	o.LogCachePath = "/srv/arcdps-config"
	o.ElitePath = "/srv/elite-insights"
	o.OutputPath = "/srv/gw2-report"
	Run(o, h)
}
