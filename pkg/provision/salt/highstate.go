package salt

import (
	"github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/go-library/pkg/provision/salt/constant"
)

func (c *Client) Highstate(
	target string,
) (map[string]response.LocalReturn, error) {
	return c.Local(target, constant.Highstate, nil)
}
