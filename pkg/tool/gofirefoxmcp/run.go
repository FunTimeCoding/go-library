package gofirefoxmcp

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/firefox"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/tool"
	"github.com/mark3labs/mcp-go/server"
	"net/http"
)

func Run(
	address string,
	bridgePort int,
) {
	c := firefox.NewEnvironment()
	s := server.NewMCPServer(constant.Name, constant.Version)
	t := tool.New(c)
	addTool(s, t)
	v := generative.New(s)
	bridgeAddress := fmt.Sprintf(":%d", bridgePort)
	b := lifecycle.New(
		lifecycle.WithServer(
			bridgeAddress,
			func(mx *http.ServeMux) {
				mx.Handle("/", c)
			},
		),
		lifecycle.WithServer(
			address,
			func(mx *http.ServeMux) {
				v.Setup(mx)
			},
		),
	)
	b.RunUntilSignal()
}
