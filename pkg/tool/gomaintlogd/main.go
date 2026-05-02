package gomaintlogd

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/relational/postgres"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(
		"gomaintlog",
		environment.Optional(constant.LocatorEnvironment),
		"",
		version,
	)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.PostgresLocator = environment.Optional(postgres.LocatorEnvironment)
	o.SQLitePath = environment.Optional(store.PathEnvironment)
	Run(o, r)
}
