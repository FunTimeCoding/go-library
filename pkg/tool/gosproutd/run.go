package gosproutd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/telemetry"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/notifier"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/option"
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
	n := notifier.New()
	s := store.New(store.DefaultDatabasePath())
	defer s.Close()
	w := watcher.New(s, n, l, r, o.SeedDirectory)
	lifecycle.New(
		l,
		lifecycle.WithWorker(w),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				model_context.New(
					s,
					n,
					r,
					telemetry.NewEnvironment(),
					o.Version,
				).Mount(m)
				sproutWeb.New(s, n).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
