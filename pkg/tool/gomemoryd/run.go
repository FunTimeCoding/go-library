package gomemoryd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	generated "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/memory_indexer"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"
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
	q := connect.Wait(l)
	i := memory_indexer.New(q)
	svc := service.New(s, i)
	reconcileMemories(svc)
	lifecycle.New(
		l,
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				generated.HandlerFromMux(server.New(s), m)
				model_context.New(svc, r, o.Version).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
