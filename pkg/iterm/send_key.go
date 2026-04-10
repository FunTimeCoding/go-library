package iterm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
)

func (c *Client) SendKey(identifier string, key string) error {
	l := c.base.Copy().Path("/sessions/%s/key", identifier).String()
	body, e := json.Marshal(map[string]string{"key": key})

	if e != nil {
		return fmt.Errorf("send key: %w", e)
	}

	r, e := c.client.Post(l, "application/json", bytes.NewReader(body))

	if e != nil {
		return fmt.Errorf("send key: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		b := system.ReadAll(r.Body)

		return fmt.Errorf("send key: %d: %s", r.StatusCode, b)
	}

	return nil
}
