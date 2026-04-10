package sublime

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/sublime/view"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
)

func (c *Client) View(identifier int) (view.View, error) {
	l := c.base.Copy().Path("/views/%d", identifier).String()
	r, e := c.client.Get(l)

	if e != nil {
		return view.View{}, fmt.Errorf("read view: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		b := system.ReadAll(r.Body)

		return view.View{}, fmt.Errorf("read view: %d: %s", r.StatusCode, b)
	}

	var result view.View

	if e = json.NewDecoder(r.Body).Decode(&result); e != nil {
		return view.View{}, fmt.Errorf("read view: %w", e)
	}

	return result, nil
}
