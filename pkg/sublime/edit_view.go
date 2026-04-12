package sublime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/sublime/view"
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
)

func (c *Client) EditView(
	identifier int,
	old string,
	new string,
	all bool,
) (view.View, error) {
	l := c.base.Copy().Path("/views/%d", identifier).String()
	body, e := json.Marshal(
		map[string]any{
			"old_string":  old,
			"new_string":  new,
			"replace_all": all,
		},
	)

	if e != nil {
		return view.View{}, fmt.Errorf("edit view: %w", e)
	}

	q, e := http.NewRequest(http.MethodPut, l, bytes.NewReader(body))

	if e != nil {
		return view.View{}, fmt.Errorf("edit view: %w", e)
	}

	q.Header.Set("Content-Type", "application/json")
	r, e := c.client.Do(q)

	if e != nil {
		return view.View{}, fmt.Errorf("edit view: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		b := system.ReadAll(r.Body)

		return view.View{}, fmt.Errorf("edit view: %d: %s", r.StatusCode, b)
	}

	var result view.View

	if e = json.NewDecoder(r.Body).Decode(&result); e != nil {
		return view.View{}, fmt.Errorf("edit view: %w", e)
	}

	return result, nil
}
