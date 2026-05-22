package basic

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/go-library/pkg/provision/salt/constant"
	"net/http"
)

func (c *Client) login() {
	b, code, e := c.exchange(
		http.MethodPost,
		constant.LoginPath,
		loginRequest{
			Username: c.user,
			Password: c.password,
			EAuth:    c.eauth,
		},
	)
	errors.PanicOnError(e)

	if code >= http.StatusBadRequest {
		errors.PanicOnError(parseDetail(b, code))
	}

	var r response.Login
	errors.PanicOnError(json.Unmarshal(b, &r))
	c.token = r.Return[0].Token
}
