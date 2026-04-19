package hub

import (
	"github.com/funtimecoding/go-library/pkg/docker/hub/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func New() *Client {
	return &Client{
		base: locator.New(constant.Host).Base(constant.BasePath),
	}
}
