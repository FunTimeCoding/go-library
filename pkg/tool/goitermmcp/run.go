package goitermmcp

import (
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/iterm"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/tool/goitermmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goitermmcp/tool"
	"github.com/mark3labs/mcp-go/server"
	"net/http"
)

func Run(address string) {
	c := iterm.NewEnvironment()
	s := server.NewMCPServer(constant.Name, constant.Version)
	t := tool.New(c)
	addTool(s, t)
	v := generative.New(s)
	b := lifecycle.New(
		lifecycle.WithServer(
			address,
			func(mx *http.ServeMux) {
				v.Setup(mx)
			},
		),
	)
	b.RunUntilSignal()
}
