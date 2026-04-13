package sublime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/sublime/view"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/web/constant"
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

	q, f := http.NewRequest(http.MethodPut, l, bytes.NewReader(body))

	if f != nil {
		return view.View{}, fmt.Errorf("edit view: %w", f)
	}

	q.Header.Set("Content-Type", constant.Object)
	r, g := c.client.Do(q)

	if g != nil {
		return view.View{}, fmt.Errorf("edit view: %w", g)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return view.View{}, fmt.Errorf(
			"edit view: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	var result view.View

	if h := json.NewDecoder(r.Body).Decode(&result); h != nil {
		return view.View{}, fmt.Errorf("edit view: %w", h)
	}

	return result, nil
}
