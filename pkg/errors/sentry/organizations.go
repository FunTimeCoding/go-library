package sentry

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) Organizations() ([]response.Organization, error) {
	b, e := c.basic.Get("organizations", nil)

	if e != nil {
		return nil, e
	}

	var result []response.Organization
	f := json.Unmarshal(b, &result)

	if f != nil {
		return nil, f
	}

	return result, nil
}
