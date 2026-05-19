package goqueryd

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store"
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
	a.String(
		"database",
		store.DefaultDatabasePath(),
		"Path to SQLite database file",
	)
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Port = a.RequiredInteger(argument.Port)
	o.DatabasePath = a.GetString("database")
	o.RerankModel = environment.Required(constant.ModelEnvironment)
	o.RerankTokenizer = environment.Required(constant.TokenizerEnvironment)
	o.Version = version
	Run(o, r)
}
