package client

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/location"
	"net/url"
)

func New() *Client {
	return &Client{
		locator: url.URL{
			Scheme: web.SocketScheme,
			Host:   constant.Address,
			Path:   location.Monitor,
		},
		done: make(chan struct{}),
	}
}
