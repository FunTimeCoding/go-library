package hetzner

import (
	"context"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

type Client struct {
	context context.Context
	client  *hcloud.Client
}
