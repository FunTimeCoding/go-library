package client

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) ensureProvider() {
	c.providerOnce.Do(
		func() {
			p, e := oidc.NewProvider(context.Background(), c.issuer)
			errors.PanicOnError(e)
			c.provider = p
			c.verifier = p.Verifier(
				&oidc.Config{ClientID: c.identifier},
			)
		},
	)
}
