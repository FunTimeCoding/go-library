package iterm

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/iterm/session"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
)

func (c *Client) CreateTab() (session.Session, error) {
	l := c.base.Copy().Path("/tabs").String()
	r, e := c.client.Post(l, "application/json", nil)

	if e != nil {
		return session.Session{}, fmt.Errorf("create tab: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusCreated {
		b := system.ReadAll(r.Body)

		return session.Session{}, fmt.Errorf(
			"create tab: %d: %s",
			r.StatusCode,
			b,
		)
	}

	var result session.Session

	if e = json.NewDecoder(r.Body).Decode(&result); e != nil {
		return session.Session{}, fmt.Errorf("create tab: %w", e)
	}

	return result, nil
}
