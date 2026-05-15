package model_context_client

import (
	"context"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"testing"
)

type Client struct {
	t       *testing.T
	context context.Context
	session *mcp.ClientSession
}
