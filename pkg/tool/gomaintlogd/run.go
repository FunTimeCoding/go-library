package gomaintlogd

import (
	"context"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/web"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Run(o *option.Log) {
	g := logger.New(context.Background())
	locator := environment.Optional(sentry.LocatorEnvironment)

	if locator != "" {
		r := reporter.New(
			"gomaintlog",
			locator,
			"",
			o.Version,
		)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	s := newStore(o)
	defer s.Close()
	l := lifecycle.New(
		lifecycle.WithServer(
			webConstant.Listen,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(route.New(s), m)
				generative.New(model_context.New(s).Nested()).Setup(m)
				web.NewServer(s).Mount(m)
			},
		),
	)
	g.Structured("starting")
	l.Run()
	system.KillSignalBlock()
	l.Stop()
}

func newStore(o *option.Log) *store.Store {
	if o.PostgresLocator != "" {
		return store.NewPostgres(o.PostgresLocator)
	}

	if o.SQLitePath != "" {
		return store.NewSQLite(o.SQLitePath)
	}

	panic("set POSTGRES_LOCATOR or SQLITE_PATH")
}
