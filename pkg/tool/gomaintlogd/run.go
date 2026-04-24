package gomaintlogd

import (
	"context"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/web"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Run(o *option.Log) {
	g := logger.New(context.Background())
	s := newStore(o)
	defer s.Close()
	l := lifecycle.New(
		lifecycle.WithServer(
			webConstant.ListenAddress,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(route.New(s), m)
				generative.New(model_context.New(s).Nested()).Setup(m)
				web.New(s).Mount(m)
			},
		),
	)
	g.Structured("starting")
	l.RunUntilSignal()
}
