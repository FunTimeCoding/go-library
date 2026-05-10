package basic

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/go-library/pkg/provision/salt/constant"
)

func (c *Client) ListMinions() ([]response.Minion, error) {
	b, e := c.Get(constant.MinionsPath)

	if e != nil {
		return nil, e
	}

	var raw response.MinionList

	if f := json.Unmarshal(b, &raw); f != nil {
		return nil, f
	}

	if len(raw.Return) == 0 {
		return nil, nil
	}

	var result []response.Minion

	for _, v := range raw.Return[0] {
		var m response.Minion

		if f := json.Unmarshal(v, &m); f != nil {
			continue
		}

		result = append(result, m)
	}

	return result, nil
}
