package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus"
	"github.com/funtimecoding/go-library/pkg/web"
	transport "github.com/go-openapi/runtime/client"
	"github.com/prometheus/alertmanager/api/v2/client"
)

func New(
	host string,
	user string,
	password string,
	p *prometheus.Client,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(user, "user")
	errors.FatalOnEmpty(password, "password")
	c := client.NewHTTPClient(nil)
	t := transport.New(
		host,
		client.DefaultBasePath,
		[]string{web.SecureScheme},
	)
	t.DefaultAuthentication = transport.BasicAuth(user, password)
	c.SetTransport(t)

	return &Client{client: c, prometheus: p}
}
