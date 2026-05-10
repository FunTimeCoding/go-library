package basic

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/go-library/pkg/provision/salt/constant"
)

func (c *Client) ListKeys() (*response.KeysReturn, error) {
	b, e := c.Get(constant.KeysPath)

	if e != nil {
		return nil, e
	}

	var r response.Keys
	f := json.Unmarshal(b, &r)

	if f != nil {
		return nil, f
	}

	return &r.Return, nil
}
