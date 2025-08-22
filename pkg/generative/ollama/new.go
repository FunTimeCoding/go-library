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

func New(o ...OptionFunc) *Client {
	result := &Client{context: context.Background()}

	for _, p := range o {
		p(result)
	}

	if result.host == "" {
		result.host = constant.Host
	}

	if result.port == 0 {
		result.port = constant.Port
	}

	var scheme string

	if result.secure {
		scheme = web.SecureScheme
	} else {
		scheme = web.InsecureScheme
	}

	// https://github.com/ollama/ollama/blob/main/docs/api.md
	result.client = api.NewClient(
		&url.URL{
			Scheme: scheme,
			Host: net.JoinHostPort(
				result.host,
				fmt.Sprintf("%d", result.port),
			),
		},
		http.DefaultClient,
	)

	return result
}
