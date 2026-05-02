package gomatmcp

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/option"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Mattermost,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	c := mattermost.NewEnvironment()
	var m *monitor.Monitor
	var p []lifecycle.Option

	if v := monitor.LoadConfiguration(); v.Enabled {
		m = monitor.New(c, v, l, r)
		p = append(p, lifecycle.WithWorker(m))
	}

	lifecycle.New(
		l,
		append(
			p,
			lifecycle.WithServerMiddleware(
				web.AddressPort(o.Port),
				func(u *http.ServeMux) {
					model_context.New(c, m, r).Mount(u)
				},
				web.RecoveryMiddleware(r),
			),
		)...,
	).RunUntilSignal()
}
