package client

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func New(host string) *Client {
	c, e := client.NewClientWithResponses(locator.New(host).String())
	errors.PanicOnError(e)

	return &Client{context: context.Background(), client: c}
}
