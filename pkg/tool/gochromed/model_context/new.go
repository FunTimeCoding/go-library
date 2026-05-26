package model_context

import (
	"github.com/funtimecoding/go-library/pkg/chromium"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gochromed/constant"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	c *chromium.Client,
	downloadDirectory string,
	r face.Reporter,
	version string,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Identity.Name(),
			version,
			server.WithToolCapabilities(true),
			server.WithInstructions(constant.Identity.Instructions()),
		),
		client:            c,
		downloadDirectory: downloadDirectory,
		snapshotCache:     make(map[string]map[string]int64),
		reporter:          r,
	}
	result.register()

	return result
}
