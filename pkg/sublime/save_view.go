package sublime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (c *Client) SaveView(
	identifier int,
	path string,
) error {
	l := c.base.Copy().Path("/views/%d/save", identifier).String()
	body, e := json.Marshal(
		map[string]string{
			"file_path": path,
		},
	)

	if e != nil {
		return fmt.Errorf("save view: %w", e)
	}

	r, f := c.client.Post(l, constant.Object, bytes.NewReader(body))

	if f != nil {
		return fmt.Errorf("save view: %w", f)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"save view: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	return nil
}
