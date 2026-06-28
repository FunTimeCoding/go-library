package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	atlassianFace "github.com/funtimecoding/go-library/pkg/tool/goatlassiand/face"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/service"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server     *server.MCPServer
	jira       atlassianFace.JiraSource
	confluence atlassianFace.ConfluenceSource
	service    *service.Service
	reporter   face.Reporter
}
