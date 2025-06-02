//go:build local

package chroma

import (
	"context"
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func New() *Client {
	client, e := v2.NewHTTPClient()
	errors.PanicOnError(e)

	return &Client{context: context.Background(), client: client}
}
