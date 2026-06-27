package model_context

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/face"
	atlassianFace "github.com/funtimecoding/go-library/pkg/tool/goatlassiand/face"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	server     *server.MCPServer
	jira       *jira.Client
	confluence atlassianFace.ConfluenceSource
	reporter   face.Reporter
}
