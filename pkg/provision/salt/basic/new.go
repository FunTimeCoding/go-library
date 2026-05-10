package basic

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"net/http"
	"net/http/cookiejar"
)

func New(
	host string,
	port int,
	user string,
	password string,
	eauth string,
) *Client {
	jar, e := cookiejar.New(nil)
	errors.PanicOnError(e)

	result := &Client{
		base:   locator.New(host).Port(port).Insecure().String(),
		client: &http.Client{Jar: jar},
	}
	result.login(user, password, eauth)

	return result
}
