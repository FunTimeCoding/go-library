package goraidparsed

import (
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/tool/goraidparsed/option"
	"github.com/funtimecoding/go-library/pkg/tool/goraidparsed/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidparsed/server"
	"net/http"
)

func Run(o *option.Parser) {
	l := lifecycle.New(
		lifecycle.WithServer(
			":8081",
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					route.New(o.ParserPath, o.TemplatePath, o.OutputPath),
					m,
				)
			},
		),
	)
	l.RunUntilSignal()
}
