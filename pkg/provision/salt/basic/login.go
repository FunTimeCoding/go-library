package basic

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/go-library/pkg/provision/salt/constant"
)

func (c *Client) login(
	user string,
	password string,
	eauth string,
) {
	b, e := c.Post(
		constant.LoginPath,
		loginRequest{
			Username: user,
			Password: password,
			EAuth:    eauth,
		},
	)
	errors.PanicOnError(e)
	var r response.Login
	errors.PanicOnError(json.Unmarshal(b, &r))
	c.token = r.Return[0].Token
}
