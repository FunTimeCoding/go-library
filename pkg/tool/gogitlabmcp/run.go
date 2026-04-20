package gogitlabmcp

import (
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/tool"
	"github.com/mark3labs/mcp-go/server"
	"net/http"
)

func Run(address string) {
	s := server.NewMCPServer(constant.Name, constant.Version)
	addTool(s, tool.New(gitlab.NewEnvironment().Nested()))
	lifecycle.New(
		lifecycle.WithServer(
			address,
			func(m *http.ServeMux) {
				generative.New(s).Setup(m)
			},
		),
	).RunUntilSignal()
}
