package habitica

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/habitica/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
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
		base:           join.Empty(locator.New(host).String(), constant.Base),
		userIdentifier: userID,
		token:          token,
		client:         &http.Client{},
	}
}
