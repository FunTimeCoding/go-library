package goraidparsed

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidparsed/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goraidparsed/option"
	"github.com/funtimecoding/go-library/pkg/tool/goraidparsed/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Parser,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	lifecycle.New(
		l,
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					server.New(
						o.ParserPath,
						o.TemplatePath,
						o.OutputPath,
						l,
						r,
					),
					m,
				)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
