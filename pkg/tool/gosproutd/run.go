package gosproutd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/event/notifier"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/telemetry"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/store"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/watcher"
	sproutWeb "github.com/funtimecoding/go-library/pkg/tool/gosproutd/web"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Sprout,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	s := store.New(store.DefaultDatabasePath())
	defer s.Close()
	v := service.New(s, notifier.New())
	w := watcher.New(v, l, r, o.SeedDirectory)
	lifecycle.New(
		l,
		lifecycle.WithWorker(w),
		lifecycle.WithServer(
			server.New(
				web.AddressPort(o.Port),
				func(m *http.ServeMux) {
					model_context.New(
						v,
						r,
						telemetry.NewEnvironment(),
						o.Version,
					).Mount(m)
					sproutWeb.New(v).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
