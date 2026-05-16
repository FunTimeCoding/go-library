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

func (c *Client) OpenFile(path string) (*view.View, error) {
	body, e := json.Marshal(map[string]string{"file_path": path})

	if e != nil {
		return view.Stub(), fmt.Errorf("open file: %w", e)
	}

	r, f := c.client.Post(
		c.base.Copy().Path("/open").String(),
		constant.Object,
		bytes.NewReader(body),
	)

	if f != nil {
		return view.Stub(), fmt.Errorf("open file: %w", f)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return view.Stub(), fmt.Errorf(
			"open file: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	var result *view.View

	if g := json.NewDecoder(r.Body).Decode(&result); g != nil {
		return view.Stub(), fmt.Errorf("open file: %w", g)
	}

	return result, nil
}
