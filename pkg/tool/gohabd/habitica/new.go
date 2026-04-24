package habitica

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"net/http"
)

func New(
	host string,
	userID string,
	token string,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(userID, "user")
	errors.FatalOnEmpty(token, "token")

	return &Client{
		baseURL: locator.New(host).String() + apiBase,
		userID:  userID,
		token:   token,
		http:    &http.Client{},
	}
}
