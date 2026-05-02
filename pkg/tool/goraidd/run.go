package goraidd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/raid_parser"
	"github.com/funtimecoding/go-library/pkg/relational"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/server"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
	raidWeb "github.com/funtimecoding/go-library/pkg/tool/goraidd/web"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Raid,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	s := store.New(
		relational.New(o.PostgresLocator).Mapper(),
		o.LogCachePath,
		o.ElitePath,
		l,
		r,
	)
	lifecycle.New(
		l,
		lifecycle.WithWorker(s),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				p := raid_parser.New("localhost:8081", true)
				generated.HandlerFromMux(
					server.New(
						s,
						o.ElitePath,
						o.OutputPath,
						p,
					),
					m,
				)
				raidWeb.New(
					s,
					o.ElitePath,
					o.OutputPath,
					p,
				).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
