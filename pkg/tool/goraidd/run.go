package goraidd

import (
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidd/server"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/web"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Run(o *option.Raid) {
	l := lifecycle.New(
		lifecycle.WithServer(
			webConstant.Listen,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					route.New(o.LogCachePath, o.OutputPath),
					m,
				)
				web.NewServer(
					o.LogCachePath,
					o.EliteInsightsPath,
					o.OutputPath,
				).Mount(m)
			},
		),
	)
	l.RunUntilSignal()
}
