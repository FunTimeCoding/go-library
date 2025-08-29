package netbox

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/cache"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/netbox-community/go-netbox/v4"
)

func New(
	host string,
	token string,
	o ...OptionFunc,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(token, "token")
	result := &Client{
		context: context.Background(),
		client: netbox.NewAPIClientFor(
			fmt.Sprintf("%s://%s", constant.SecureScheme, host),
			token,
		),
		cache: cache.New(),
	}

	for _, f := range o {
		f(result)
	}

	if result.verbose {
		result.client.GetConfig().Debug = true
	}

	return result
}
