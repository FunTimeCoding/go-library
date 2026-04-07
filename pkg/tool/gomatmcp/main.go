package gomatmcp

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/errors"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	gomonitor "github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/mark3labs/mcp-go/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/tool"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("gomatmcp", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	gomonitor.ParseBind(version, gitHash, buildDate)
	m := mattermost.NewEnvironment()
	var o *monitor.Monitor

	if c := monitor.LoadConfiguration(); c.Enabled {
		o = monitor.New(m, c)
		errors.PanicOnError(o.Start())
		defer o.Stop()
	}

	s := server.NewMCPServer("Mattermost MCP", constant.DefaultVersion)
	addTool(s, tool.New(m, o), o != nil)
	errors.PanicOnError(
		server.NewStreamableHTTPServer(s).Start(web.Listen),
	)
}
