package gonetbd

import (
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Run(o *option.Netbox) {
	lifecycle.New(
		lifecycle.WithServer(
			web.Listen,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(route.New(o.Client), m)
				generative.New(
					model_context.New(o.Client).Nested(),
				).Setup(m)
			},
		),
	).RunUntilSignal()
}
