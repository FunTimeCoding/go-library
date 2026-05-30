package gokubernetesd

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/telemetry"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/store"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"os"
)

func Run(
	o *option.Server,
	r face.Reporter,
) {
	svc := service.New(store.New(o.Path))

	if !svc.HasClusters() {
		fmt.Fprintln(os.Stderr, "no kubernetes clusters available")
		os.Exit(1)
	}

	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				model_context.New(
					svc,
					o.ReadOnly,
					r,
					telemetry.NewEnvironment(),
					o.Version,
				).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
