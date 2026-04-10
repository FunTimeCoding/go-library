package gosublmcp

import (
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/sublime"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/tool"
	"github.com/mark3labs/mcp-go/server"
	"net/http"
)

func Run(address string) {
	c := sublime.NewEnvironment()
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
