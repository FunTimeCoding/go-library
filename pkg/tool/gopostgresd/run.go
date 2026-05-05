package gopostgresd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	generated "github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/store"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Postgres,
	r face.Reporter,
) {
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				s := store.New(o.Inventory)
				generated.HandlerFromMux(server.New(s), m)
				model_context.New(s, r, o.Version).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
