package vault

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"github.com/hashicorp/vault-client-go"
	"time"
)

func New(host string) *Client {
	client, e := vault.New(
		vault.WithAddress(locator.New(host).String()),
		vault.WithRequestTimeout(10*time.Second),
	)
	errors.PanicOnError(e)

	return &Client{client: client}
}
