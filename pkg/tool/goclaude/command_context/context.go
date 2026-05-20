package command_context

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"

type Context struct {
	host   string
	port   int
	client *client.ClientWithResponses
}
