package iterm

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/iterm/session"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
)

func (c *Client) ReadScreen(identifier string) (session.Screen, error) {
	l := c.base.Copy().Path("/sessions/%s/screen", identifier).String()
	r, e := c.client.Get(l)

	if e != nil {
		return session.Screen{}, fmt.Errorf("read screen: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		b := system.ReadAll(r.Body)

		return session.Screen{}, fmt.Errorf("read screen: %d: %s", r.StatusCode, b)
	}

	var result session.Screen

	if e = json.NewDecoder(r.Body).Decode(&result); e != nil {
		return session.Screen{}, fmt.Errorf("read screen: %w", e)
	}

	return result, nil
}
