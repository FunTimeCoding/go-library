package technitium

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/detail_error"
	"io"
	"net/http"
)

func (c *Client) do(path string) (*http.Response, error) {
	r, e := http.NewRequest(
		http.MethodGet,
		join.Empty(c.base, path),
		nil,
	)

	if e != nil {
		return nil, e
	}

	web.Bearer(r, c.token)
	result, f := c.client.Do(r)

	if f != nil {
		return nil, f
	}

	if result.StatusCode >= http.StatusBadRequest {
		b, g := io.ReadAll(result.Body)

		if h := result.Body.Close(); h != nil {
			return nil, h
		}

		if g != nil {
			return nil, fmt.Errorf(
				"technitium %s: %d (body unreadable: %w)",
				path,
				result.StatusCode,
				g,
			)
		}

		return nil, detail_error.New(string(b), result.Status)
	}

	return result, nil
}
