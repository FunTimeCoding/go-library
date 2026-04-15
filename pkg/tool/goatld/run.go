package goatld

import (
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/option"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Run(o *option.Atlassian) {
	lifecycle.New(
		lifecycle.WithServer(
			web.Listen,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(route.New(o.Jira, o.Confluence), m)
				generative.New(
					model_context.New(o.Jira, o.Confluence).Nested(),
				).Setup(m)
			},
		),
	).RunUntilSignal()
}
