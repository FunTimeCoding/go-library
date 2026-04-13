package iterm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (c *Client) SendKey(
	identifier string,
	key string,
) error {
	l := c.base.Copy().Path("/sessions/%s/key", identifier).String()
	body, e := json.Marshal(map[string]string{"key": key})

	if e != nil {
		return fmt.Errorf("send key: %w", e)
	}

	r, f := c.client.Post(l, constant.Object, bytes.NewReader(body))

	if f != nil {
		return fmt.Errorf("send key: %w", f)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"send key: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	return nil
}
