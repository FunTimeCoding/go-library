package gohabd

import (
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/gohabd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(o *option.Habitica) {
	c := habitica.NewEnvironment()
	lifecycle.New(
		lifecycle.WithServer(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				generated.HandlerFromMux(route.New(c), m)
				generative.New(model_context.New(c).Nested()).Setup(m)
			},
		),
	).RunUntilSignal()
}
