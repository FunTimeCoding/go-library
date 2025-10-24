package kestra

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/kestra-io/client-sdk/go-sdk"
	"net/http"
)

type basicTransport struct {
	username string
	password string
	r        http.RoundTripper
}

func (t *basicTransport) RoundTrip(e *http.Request) (*http.Response, error) {
	e.SetBasicAuth(t.username, t.password)

	return t.r.RoundTrip(e)
}

func New(
	host string,
	o ...Option,
) *Client {
	// https://k.s3n.sh/api
	// https://kestra.io/docs/how-to-guides/api
	// https://kestra.io/docs/api-reference/open-source
	// Token authentication is EE only: https://kestra.io/docs/enterprise/auth/api-tokens
	c := kestra_api_client.NewConfiguration()
	c.Scheme = constant.Secure
	c.Host = host
	result := &Client{context: context.Background()}

	for _, f := range o {
		f(result)
	}

	if result.token != "" {
		c.DefaultHeader[constant.Authorization] = key_value.Space(
			constant.Bearer,
			result.token,
		)
	}

	if result.user != "" && result.password != "" {
		l := web.Client()
		l.Transport = &basicTransport{
			username: result.user,
			password: result.password,
			r:        http.DefaultTransport,
		}
		c.HTTPClient = l
	}

	result.client = kestra_api_client.NewAPIClient(c)

	return result
}
