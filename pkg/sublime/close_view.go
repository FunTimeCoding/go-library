package sublime

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
)

func (c *Client) CloseView(identifier int) error {
	l := c.base.Copy().Path("/views/%d", identifier).String()
	q, e := http.NewRequest(http.MethodDelete, l, nil)

	if e != nil {
		return fmt.Errorf("close view: %w", e)
	}

	r, e := c.client.Do(q)

	if e != nil {
		return fmt.Errorf("close view: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		b := system.ReadAll(r.Body)

		return fmt.Errorf("close view: %d: %s", r.StatusCode, b)
	}

	return nil
}
