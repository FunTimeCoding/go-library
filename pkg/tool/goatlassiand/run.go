package goatlassiand

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	generated "github.com/funtimecoding/go-library/pkg/tool/goatlassiand/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/option"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Atlassian,
	r face.Reporter,
) {
	j := jira.NewEnvironment()
	c := confluence.NewEnvironment()
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				generated.HandlerFromMux(server.New(j, c), m)
				model_context.New(j, c, r, o.Version).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
