package client

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/location"
	"net/url"
)

func New() *Client {
	return &Client{
		locator: url.URL{
			Scheme: web.Socket,
			Host:   constant.Address,
			Path:   location.Monitor,
		},
		done: make(chan struct{}),
	}
}
