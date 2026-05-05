package gosublimed

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/sublime"
	"github.com/funtimecoding/go-library/pkg/tool/gosublimed/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gosublimed/option"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Sublime,
	r face.Reporter,
) {
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				model_context.New(sublime.NewEnvironment(), r, o.Version).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
