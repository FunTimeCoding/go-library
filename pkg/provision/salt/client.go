package salt

import (
	"context"
	"github.com/daixijun/go-salt/v2"
)

type Client struct {
	context context.Context
	client  *salt.Client
}
