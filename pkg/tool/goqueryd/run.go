package goqueryd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/ollama"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	generated "github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/server"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/store"
	queryWeb "github.com/funtimecoding/go-library/pkg/tool/goqueryd/web"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/worker"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"time"
)

func Run(
	o *option.Query,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	s := store.New(o.DatabasePath)
	defer s.Close()
	a, e := rerank.New(o.RerankModel, o.RerankTokenizer)
	errors.PanicOnError(e)
	defer errors.LogClose(a)
	v := service.New(s, ollama.NewEnvironment(), a)
	lifecycle.New(
		l,
		lifecycle.WithWorker(
			worker.New(v, 10*time.Minute, l, r),
		),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				generated.HandlerFromMux(server.New(v), m)
				model_context.New(v, r, o.Version).Mount(m)
				queryWeb.New(v).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
