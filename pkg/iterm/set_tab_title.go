package iterm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
)

func (c *Client) SetTabTitle(tabIdentifier string, title string) error {
	l := c.base.Copy().Path("/tabs/%s/title", tabIdentifier).String()
	body, e := json.Marshal(map[string]string{"title": title})

	if e != nil {
		return fmt.Errorf("set tab title: %w", e)
	}

	r, e := c.client.Post(l, "application/json", bytes.NewReader(body))

	if e != nil {
		return fmt.Errorf("set tab title: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		b := system.ReadAll(r.Body)

		return fmt.Errorf("set tab title: %d: %s", r.StatusCode, b)
	}

	return nil
}
