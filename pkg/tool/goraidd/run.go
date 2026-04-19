package goraidd

import (
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/raid_parser"
	"github.com/funtimecoding/go-library/pkg/relational"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidd/server"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/web"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Run(o *option.Raid) {
	p := raid_parser.New("localhost:8081", true)
	r := relational.New(o.PostgresLocator)
	s := store.New(r.Mapper(), o.LogCachePath, o.EliteInsightsPath)
	stop := make(chan struct{})
	go s.RunBackground(stop)
	l := lifecycle.New(
		lifecycle.WithServer(
			webConstant.Listen,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					route.New(
						s,
						o.EliteInsightsPath,
						o.OutputPath,
						p,
					),
					m,
				)
				web.NewServer(
					s,
					o.EliteInsightsPath,
					o.OutputPath,
					p,
				).Mount(m)
			},
		),
	)
	l.RunUntilSignal()
	close(stop)
}
