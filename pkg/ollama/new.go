package ollama

import (
	"context"
	"github.com/ollama/ollama/api"
	"net"
	"net/http"
	"net/url"
)

func New() *Client {
	scheme := "http"
	host := "localhost"
	port := "11434"

	return &Client{
		context: context.Background(),
		client: api.NewClient(
			&url.URL{Scheme: scheme, Host: net.JoinHostPort(host, port)},
			http.DefaultClient,
		),
	}
}
