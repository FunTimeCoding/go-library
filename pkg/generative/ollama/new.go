package ollama

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/ollama/constant"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"github.com/ollama/ollama/api"
	"net/http"
	"net/url"
)

func New(o ...Option) *Client {
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
		scheme = web.Secure
	} else {
		scheme = web.Insecure
	}

	l := locator.New(result.host).Port(result.port)

	if !result.secure {
		l.Insecure()
	}

	// https://github.com/ollama/ollama/blob/main/docs/api.md
	result.client = api.NewClient(
		&url.URL{Scheme: scheme, Host: l.String()},
		http.DefaultClient,
	)

	return result
}
