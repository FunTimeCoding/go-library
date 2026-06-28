package technitium

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"net/http"
)

func New(
	host string,
	token string,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(token, "token")

	return &Client{
		base:   join.Empty(locator.New(host).String(), "/api"),
		token:  token,
		client: &http.Client{},
	}
}
