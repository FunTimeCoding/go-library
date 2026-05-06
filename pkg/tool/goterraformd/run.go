package goterraformd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/relational"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/runner"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/store"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Terraform,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	s := store.New(relational.New(o.PostgresLocator).Mapper())
	n := runner.New(o, s, l, r)
	lifecycle.New(
		l,
		lifecycle.WithWorker(n),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				model_context.New(n, s, r, o.Version).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
