package ollama

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/constant"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/ollama/ollama/api"
	"net"
	"net/http"
	"net/url"
)

func New(host string) *Client {
	// https://github.com/ollama/ollama/blob/main/docs/api.md
	return &Client{
		context: context.Background(),
		client: api.NewClient(
			&url.URL{
				Scheme: web.InsecureScheme,
				Host: net.JoinHostPort(
					host,
					fmt.Sprintf("%d", constant.Port),
				),
			},
			http.DefaultClient,
		),
	}
}
