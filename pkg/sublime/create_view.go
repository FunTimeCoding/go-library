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

func (c *Client) CreateView(
	title string,
	content string,
	syntax string,
) (view.View, error) {
	l := c.base.Copy().Path("/views").String()
	body, e := json.Marshal(
		map[string]string{
			"title":   title,
			"content": content,
			"syntax":  syntax,
		},
	)

	if e != nil {
		return view.View{}, fmt.Errorf("create view: %w", e)
	}

	r, e := c.client.Post(l, "application/json", bytes.NewReader(body))

	if e != nil {
		return view.View{}, fmt.Errorf("create view: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusCreated {
		b := system.ReadAll(r.Body)

		return view.View{}, fmt.Errorf("create view: %d: %s", r.StatusCode, b)
	}

	var result view.View

	if e = json.NewDecoder(r.Body).Decode(&result); e != nil {
		return view.View{}, fmt.Errorf("create view: %w", e)
	}

	return result, nil
}
