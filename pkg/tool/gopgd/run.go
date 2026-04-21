package gopgd

import (
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/config"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/gopgd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/store"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Run(c *config.Configuration) {
	s := store.New(c)
	lifecycle.New(
		lifecycle.WithServer(
			web.ListenAddress,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(route.New(s), m)
				generative.New(
					model_context.New(s).Nested(),
				).Setup(m)
			},
		),
	).RunUntilSignal()
}
