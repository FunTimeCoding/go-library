package basic

import (
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"net/http"
)

func New(
	host string,
	port int,
	user string,
	password string,
	eauth string,
) *Client {
	result := &Client{
		base:   locator.New(host).Port(port).Insecure().String(),
		client: &http.Client{},
	}
	result.login(user, password, eauth)

	return result
}
