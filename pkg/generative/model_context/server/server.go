package server

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/mark3labs/mcp-go/server"
	"github.com/mark3labs/mcp-go/util"
	"sync"
)

type Server struct {
	server  *server.MCPServer
	logger  util.Logger
	context context.Context

	openAuthentication   bool
	verifier             *oidc.IDTokenVerifier
	once                 sync.Once
	serverLocator        string
	authorizationLocator string
	clientIdentifier     string

	tokenAuthentication bool
	token               string
}
