package gomaintlogd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	generated "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	maintenanceWeb "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/web"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Run(
	o *option.Log,
	r face.Reporter,
) {
	g := logger.New(context.Background())
	s := store.New(o)
	defer s.Close()
	lifecycle.New(
		g,
		lifecycle.WithServerMiddleware(
			constant.ListenAddress,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(server.New(s), m)
				model_context.New(s, r, o.Version).Mount(m)
				maintenanceWeb.New(s).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
