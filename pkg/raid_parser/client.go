package raid_parser

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goraidparsed/client"
)

type Client struct {
	context context.Context
	client  *client.ClientWithResponses
}
