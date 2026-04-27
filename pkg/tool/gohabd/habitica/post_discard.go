package habitica

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func (c *Client) postDiscard(path string) {
	r := c.do(http.MethodPost, path, nil)
	errors.PanicClose(r.Body)
}
