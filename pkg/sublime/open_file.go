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

func (c *Client) OpenFile(path string) (view.View, error) {
	l := c.base.Copy().Path("/open").String()
	body, e := json.Marshal(map[string]string{
		"file_path": path,
	})

	if e != nil {
		return view.View{}, fmt.Errorf("open file: %w", e)
	}

	r, e := c.client.Post(l, "application/json", bytes.NewReader(body))

	if e != nil {
		return view.View{}, fmt.Errorf("open file: %w", e)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		b := system.ReadAll(r.Body)

		return view.View{}, fmt.Errorf("open file: %d: %s", r.StatusCode, b)
	}

	var result view.View

	if e = json.NewDecoder(r.Body).Decode(&result); e != nil {
		return view.View{}, fmt.Errorf("open file: %w", e)
	}

	return result, nil
}
