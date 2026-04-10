package iterm

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/iterm/session"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
)

func (c *Client) Sessions() ([]session.Session, error) {
	l := c.base.Copy().Path("/sessions").String()
	r, e := c.client.Get(l)

	if e != nil {
		return nil, fmt.Errorf("list sessions: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		b := system.ReadAll(r.Body)

		return nil, fmt.Errorf("list sessions: %d: %s", r.StatusCode, b)
	}

	var result struct {
		Sessions []session.Session `json:"sessions"`
	}

	if e = json.NewDecoder(r.Body).Decode(&result); e != nil {
		return nil, fmt.Errorf("list sessions: %w", e)
	}

	return result.Sessions, nil
}
