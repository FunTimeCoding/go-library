package iterm

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/iterm/session"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
)

func (c *Client) ReadHistory(
	identifier string,
	count int,
) (session.History, error) {
	l := c.base.Copy().Path("/sessions/%s/history", identifier).
		Set("count", fmt.Sprintf("%d", count)).String()
	r, e := c.client.Get(l)

	if e != nil {
		return session.History{}, fmt.Errorf("read history: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		b := system.ReadAll(r.Body)

		return session.History{}, fmt.Errorf(
			"read history: %d: %s",
			r.StatusCode,
			b,
		)
	}

	var result session.History

	if e = json.NewDecoder(r.Body).Decode(&result); e != nil {
		return session.History{}, fmt.Errorf("read history: %w", e)
	}

	return result, nil
}
