package model_context_client

import "github.com/modelcontextprotocol/go-sdk/mcp"

func (c *Client) InitializeResult() *mcp.InitializeResult {
	return c.session.InitializeResult()
}
