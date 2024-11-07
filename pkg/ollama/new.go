package ollama

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama/constant"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/ollama/ollama/api"
	"net"
	"net/http"
	"net/url"
)

func New() *Client {
	return &Client{
		context: context.Background(),
		client: api.NewClient(
			&url.URL{
				Scheme: web.InsecureScheme,
				Host: net.JoinHostPort(
					web.Localhost,
					fmt.Sprintf("%d", constant.Port),
				),
			},
			http.DefaultClient,
		),
	}
}
