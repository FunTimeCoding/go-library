package gomatmcp

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/option"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/tool"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/mark3labs/mcp-go/server"
	"net/http"
)

func Run(o *option.Mattermost) {
	c := mattermost.NewEnvironment()
	var m *monitor.Monitor

	if v := monitor.LoadConfiguration(); v.Enabled {
		m = monitor.New(c, v)
		errors.PanicOnError(m.Start())
		defer m.Stop()
	}

	s := server.NewMCPServer("Mattermost MCP", constant.DefaultVersion)
	addTool(s, tool.New(c, m), m != nil)
	lifecycle.New(
		lifecycle.WithServer(
			web.AddressPort(o.Port),
			func(u *http.ServeMux) {
				generative.New(s).Setup(u)
			},
		),
	).RunUntilSignal()
}
