package gotelemetryd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/relational"
	generated "github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/store"
	telemetryWeb "github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/web"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Telemetry,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	s := store.New(relational.New(o.PostgresLocator).Mapper())
	lifecycle.New(
		l,
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				generated.HandlerFromMux(server.New(s), m)
				model_context.New(s, r, o.Version).Mount(m)
				telemetryWeb.New(s).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
