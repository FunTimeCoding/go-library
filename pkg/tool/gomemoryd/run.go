package gomemoryd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/telemetry"
	generated "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/memory_indexer"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
	memoryWeb "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/web"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/connect"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Memory,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	s := store.New(o.DatabasePath)
	defer s.Close()
	idx := memory_indexer.New(connect.Wait(l))
	v := service.New(s, idx, idx, idx)
	reconcileMemories(v)
	lifecycle.New(
		l,
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				generated.HandlerFromMux(server.New(v), m)
				model_context.New(
					v,
					r,
					telemetry.NewEnvironment(),
					o.Version,
				).Mount(m)
				memoryWeb.New(v).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
