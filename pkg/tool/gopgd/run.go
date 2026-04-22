package gopgd

import (
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/gopgd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/store"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(o *option.Postgres) {
	s := store.New(o.Inventory)
	lifecycle.New(
		lifecycle.WithServer(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				generated.HandlerFromMux(route.New(s), m)
				generative.New(
					model_context.New(s).Nested(),
				).Setup(m)
			},
		),
	).RunUntilSignal()
}
