package sublime

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/sublime/response"
	"github.com/funtimecoding/go-library/pkg/sublime/view"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
)

func (c *Client) Views() ([]view.View, error) {
	l := c.base.Copy().Path("/views").String()
	r, e := c.client.Get(l)

	if e != nil {
		return nil, fmt.Errorf("list views: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		b := system.ReadAll(r.Body)

		return nil, fmt.Errorf("list views: %d: %s", r.StatusCode, b)
	}

	var result response.Views

	if e = json.NewDecoder(r.Body).Decode(&result); e != nil {
		return nil, fmt.Errorf("list views: %w", e)
	}

	return result.Views, nil
}
