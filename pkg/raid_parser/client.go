package raid_parser

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goraidparsed/generated/client"
)

type Client struct {
	context context.Context
	client  *client.ClientWithResponses
}
