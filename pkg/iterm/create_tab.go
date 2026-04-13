package iterm

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/iterm/session"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (c *Client) CreateTab() (session.Session, error) {
	l := c.base.Copy().Path("/tabs").String()
	r, e := c.client.Post(l, constant.Object, nil)

	if e != nil {
		return session.Session{}, fmt.Errorf("create tab: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusCreated {
		return session.Session{}, fmt.Errorf(
			"create tab: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	var result session.Session

	if f := json.NewDecoder(r.Body).Decode(&result); f != nil {
		return session.Session{}, fmt.Errorf("create tab: %w", f)
	}

	return result, nil
}
