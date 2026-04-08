package gomatmcp

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/tool"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/mark3labs/mcp-go/server"
	"net/http"
)

func Run() {
	m := mattermost.NewEnvironment()
	var o *monitor.Monitor

	if c := monitor.LoadConfiguration(); c.Enabled {
		o = monitor.New(m, c)
		errors.PanicOnError(o.Start())
		defer o.Stop()
	}

	s := server.NewMCPServer("Mattermost MCP", constant.DefaultVersion)
	addTool(s, tool.New(m, o), o != nil)
	v := generative.New(s)
	b := lifecycle.New(
		lifecycle.WithServer(
			web.Listen,
			func(mx *http.ServeMux) {
				v.Setup(mx)
			},
		),
	)
	b.Run()
	defer b.Stop()
	system.KillSignalBlock()
}
