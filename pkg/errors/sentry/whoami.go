package sentry

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) Whoami() (*response.User, error) {
	b, e := c.basic.Get("auth", nil)

	if e != nil {
		return nil, e
	}

	var result response.User
	f := json.Unmarshal(b, &result)

	if f != nil {
		return nil, f
	}

	return &result, nil
}
