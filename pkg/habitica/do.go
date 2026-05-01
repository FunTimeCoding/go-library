package habitica

import (
	"bytes"
	"fmt"
	constant2 "github.com/funtimecoding/go-library/pkg/habitica/constant"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"io"
	"net/http"
)

func (c *Client) do(
	method string,
	path string,
	body any,
) (*http.Response, error) {
	var reader io.Reader

	if body != nil {
		reader = bytes.NewReader(notation.Marshal(body))
	}

	r, e := http.NewRequest(method, join.Empty(c.base, path), reader)

	if e != nil {
		return nil, e
	}

	r.Header.Set(constant2.UserHeader, c.userIdentifier)
	r.Header.Set(constant2.TokenHeader, c.token)
	r.Header.Set(constant.ContentType, constant.Object)
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
				"habitica %s %s: %d (body unreadable: %w)",
				method,
				path,
				result.StatusCode,
				g,
			)
		}

		return nil, fmt.Errorf(
			"habitica %s %s: %d %s",
			method,
			path,
			result.StatusCode,
			string(b),
		)
	}

	return result, nil
}
