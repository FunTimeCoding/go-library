package basic

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/go-library/pkg/provision/salt/constant"
)

func (c *Client) DeleteKey(minion string) error {
	b, e := c.Post(
		"",
		commandRequest{
			Client:   constant.WheelClient,
			Function: constant.KeyDelete,
			Match:    minion,
		},
	)

	if e != nil {
		return e
	}

	var r response.Wheel

	if f := json.Unmarshal(b, &r); f != nil {
		return f
	}

	if len(r.Return) == 0 || !r.Return[0].Result.Success {
		return fmt.Errorf("failed to delete key for %s", minion)
	}

	return nil
}
