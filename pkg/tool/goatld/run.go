package goatld

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/option"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(o *option.Atlassian) {
	j := jira.NewEnvironment()
	c := confluence.NewEnvironment()
	lifecycle.New(
		lifecycle.WithServer(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				generated.HandlerFromMux(route.New(j, c), m)
				generative.New(
					model_context.New(j, c).Nested(),
				).Setup(m)
			},
		),
	).RunUntilSignal()
}
